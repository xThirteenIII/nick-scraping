package environdec

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type PDFs []pdf

// Can add more informations based on what you need, check out response
type pdf struct {
  FriendlyUrl string `json:"friendlyUrl"`
  Title string `json:"title"`
}

func GetEnvirondecPDFs(categoryID, limit, categoryName string){

  pdfsUri := "https://api.environdec.com/api/v1/EPDLibrary/EPD?ProductCategoryId="+categoryID+"&OnlySelectorEPDs=false&Limit="+limit

  request, err := http.NewRequest("GET", pdfsUri, nil)
  if err != nil {
    log.Default().Fatalln(err)
  }

  setRequestHeaders(request)

  var parsedResponse PDFs
  // Save response
  response, err := http.DefaultClient.Do(request)
  if response != nil && response.Body != nil && response.StatusCode == http.StatusOK{
    body, err := io.ReadAll(response.Body) 
    defer response.Body.Close()
		if err != nil {
			log.Fatalln(err.Error(), "defaultClitent.Do()")
		}

		err = json.Unmarshal(body, &parsedResponse)
		if err != nil {
			response.Body.Close()
			log.Fatalln(err.Error(), "json.Unmarshal()")
		}
  }

  for requestNumber := 0; requestNumber < len(parsedResponse); requestNumber++{

    // If the id is an empty string, PDFs are over.
    if parsedResponse[requestNumber].FriendlyUrl == ""{
      fmt.Println("Done.")
      return
    }

    getPdfHtmlUri := "https://environdec.com/library/"+parsedResponse[requestNumber].FriendlyUrl
    
    PDFuri := getLinks(getPdfHtmlUri)

    downloadFileFrom(PDFuri, parsedResponse[requestNumber].Title, categoryName)

    fmt.Println(requestNumber+1, "/", len(parsedResponse))

    time.Sleep( time.Second * 2)

  }

  //className:= ".epd__Document-sc-9l6euf-11"
}

func getLinks(url string) string{

    var linkToPdf string

    resp, err := http.Get(url)

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    f := func(i int, s *goquery.Selection) bool {

        link, _ := s.Attr("href")
        return strings.HasPrefix(link, "https")
    }

    doc.Find("body a").FilterFunction(f).Each(func(_ int, tag *goquery.Selection) {

        link, _ := tag.Attr("href")
        //linkText := tag.Text()
        //fmt.Printf("%s %s\n", linkText, link)
        if strings.Contains(link, "/Data"){
          linkToPdf = link
        }
    })
  return linkToPdf
}

// Downloads a file and save it to the current folder if does not exist.
// Parameters: 
//  - uri: link to the page from which to download file
//  - fileName: the name to save the file with 
//  - dirName: name of the folder in which to put fileName
func downloadFileFrom(uri, fileName, dirName string) {

    resp, err := http.Get(uri)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    pwd, err := os.Getwd()
    if err != nil {
     fmt.Println(err.Error())
    }

    fileHandle, err := os.OpenFile(pwd+"/"+dirName+"/"+fileName+".pdf", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
    if err != nil {
        panic(err)
    }
    defer fileHandle.Close()

    _, err = io.Copy(fileHandle, resp.Body)
    if err != nil {
        panic(err)
    }

}

func setRequestHeaders(request *http.Request){

  // set Headers 
  // request.Header.Add()
  request.Header.Add("Host","api.environdec.com") 
  request.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0")
  request.Header.Add("Accept","*/*")
  request.Header.Add("Accept-Language","en-US,en;q=0.5")
  //request.Header.Add("Accept-Encoding","gzip, deflate, br")
  request.Header.Add("Referer","https://environdec.com/")
  request.Header.Add("Origin","https://environdec.com")
  request.Header.Add("DNT","1")
  request.Header.Add("Connection","keep-alive")
  request.Header.Add("Sec-Fetch-Dest" ,"empty")
  request.Header.Add("Sec-Fetch-Mode","cors")
  request.Header.Add("Sec-Fetch-Site","same-site")
  request.Header.Add("Sec-GPC","1")
}

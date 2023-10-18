package maps

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CategoriesMap map[string]string
type CategoriesIntMap map[int]string
type GeoScopesMap map[string]string

type singleCategory struct {
  Id    string `json:"id"`
  Title string `json:"title"`
}

type singleScope struct {
  Id    string `json:"id"`
  Name  string `json:"name"`
}

type Category []singleCategory
type GeoScope []singleScope


func GetEnvirondecCategories() (CategoriesMap, CategoriesIntMap){
  categoriesUri := "https://api.environdec.com/api/v1/EPDLibrary/ProductCategories"

  request, err := http.NewRequest("GET", categoriesUri, nil)
  if err != nil {
    log.Default().Fatalln(err)
  }


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

  var parsedResponse Category
  // Save response
  response, err := http.DefaultClient.Do(request)
  if response != nil && response.Body != nil && response.StatusCode == http.StatusOK{
    body, err := io.ReadAll(response.Body) 
		if err != nil {
			log.Fatalln(err.Error(), "defaultClitent.Do()")
		}

  
    
		err = json.Unmarshal(body, &parsedResponse)
		if err != nil {
			response.Body.Close()
			log.Fatalln(err.Error(), "json.Unmarshal()")
		}
  }

  categoriesMap := make(CategoriesMap)
  categoriesIntMap := make(CategoriesIntMap)

// save into Map
  for index, name := range parsedResponse{
    categoriesMap[name.Title] = name.Id
    categoriesIntMap[index] = name.Title
  }
  return categoriesMap, categoriesIntMap
  
}

func GetEnvirondecGeoScopes() GeoScopesMap{
  geoScopesUri := "https://api.environdec.com/api/v1/EPDLibrary/GeographicalScopes"

  request, err := http.NewRequest("GET", geoScopesUri, nil)
  if err != nil {
    log.Default().Fatalln(err)
  }


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

  var parsedResponse GeoScope
  // Save response
  response, err := http.DefaultClient.Do(request)
  if response != nil && response.Body != nil && response.StatusCode == http.StatusOK{
    body, err := io.ReadAll(response.Body) 
		if err != nil {
			log.Fatalln(err.Error(), "defaultClitent.Do()")
		}

  
    
		err = json.Unmarshal(body, &parsedResponse)
		if err != nil {
			response.Body.Close()
			log.Fatalln(err.Error(), "json.Unmarshal()")
		}
  }

  geoScopesMap := make(GeoScopesMap)

// save into Map
  for _, name := range parsedResponse{
    geoScopesMap[name.Name] = name.Id
  }
  return geoScopesMap
}


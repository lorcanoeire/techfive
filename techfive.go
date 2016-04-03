package main

import (
    "fmt"
    "github.com/spf13/viper"
    "html/template"
    "os"
)

type Slide struct {
  Author string
  Audience string
  Topic  string
  Name string
  Content string
}

func slide_data(slide string) Slide {

  slide_name_id := fmt.Sprintf("slides.%v.name", slide)
  slide_content_id := fmt.Sprintf("slides.%v.content", slide)
  slide_data := Slide {
    Author: viper.GetString("author"),
    Audience: viper.GetString("audience"),
    Topic: viper.GetString("topic"),
    Name: viper.GetString(slide_name_id),
    Content: viper.GetString(slide_content_id),
  }
	return slide_data
}

func main() {

  // Setup configuration
  viper.SetConfigName("techfive")
  viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No configuration file loaded - using defaults")
	}

  // Specify example slide
  var slide_id string = "slidetwo"
  slide_data := slide_data(slide_id)

  // Parse template
  t, err := template.ParseFiles("techfive-template.html")
  if err != nil {
    fmt.Println(err);
  }

  // Create and open presentation
  presentation, err := os.Create("presentation.html")
  if err != nil {
    fmt.Println("create file: ", err)
    return
  }

  // Write presentation using user based slide data
  err = t.Execute(presentation, slide_data)
  if err != nil {
    fmt.Print("execute: ", err)
    return
  }

}

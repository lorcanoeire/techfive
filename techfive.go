package main

import (
    "fmt"
    "github.com/spf13/viper"
    "html/template"
    "os"
)

const techfive_template string = "techfive-template.html"
const presentation string      = "presentation.html"
const config string            = "techfive"

type Presentation struct {
  Author string
  Audience string
  Topic  string
}

type Slide struct {
  Name string
  Content string
}

func slideData(slides []interface{}, index int) Slide {

  fmt.Printf("\nPreparing slide %v content...", index);
  slide := slides[index].(map[interface {}]interface{})
  name := fmt.Sprintf("%v", slide["name"])
  content := fmt.Sprintf("%v", slide["content"])

  slide_data := Slide {
    Name: name,
    Content: content,
  }
	return slide_data
}

func readPresentationConfig() {

  fmt.Printf("Reading user defined presentation configuration file %v.yml ...", config);
  viper.SetConfigName(config)
  viper.SetConfigType("yaml")
  viper.AddConfigPath(".")

  err := viper.ReadInConfig()
  if err != nil {
    fmt.Println("No configuration file loaded")
  }
}

func main() {

  readPresentationConfig()
  slides := viper.Get("slides").([]interface{})
  slide_data := slideData(slides, 0)

  // Parse template
  t, err := template.ParseFiles(fmt.Sprintf("templates/%v", techfive_template))
  if err != nil {
    fmt.Println(err);
  }

  // Create and open presentation file
  presentation_file, err := os.Create(fmt.Sprintf("presentations/%v", presentation))
  if err != nil {
    fmt.Println("create file: ", err)
    return
  }

  // Write presentation using user based slide data
  fmt.Printf("\nGenerating presentation\nWriting to file %v ...", presentation);
  fmt.Println("");
  err = t.Execute(presentation_file, slide_data)
  if err != nil {
    fmt.Print("execute: ", err)
    return
  }

}

# gofig
Golang package for initializing config files.

Simple package that creates and reads from a configuration file. There are Two main functions.

### NewConfig

 Params :
  - path `string` 
      - Essentially this would be what you want to name the directory the config file will be saved in. Utilizes the `HOME` env value appended with `/.config`
      - So the whole path to your config file will be `HOME`/.config/`path`/config.json
  -  data `struct` 
      - Takes the golang struct and marshalls into `JSON` format
      - This is what is stored in the config file
 
  Returns : `error`
 
 ```
  import "github.com/chris-koptional/gofig"
 
 type Config struct {
  Name  string `json:"name"`
  Age   int    `json:"age"`
 }
 
 config := Config{
     Name: "Chris",
     Age: 28
 }
 
 error := gofig.NewConfig("/myconf", config)
 
 ```
 
### GetConfig
  Params :
  - path `string` 
      - Essentially this would be what you want to name the directory the config file will be saved in. Utilizes the `HOME` env value as appended with `/.config`
      - So the whole path to your config file will be `HOME`/.config/`path`/config.json
  - data *`struct`
      -  Pointer to struct variable 
 
 Returns : `error`
 
 ```
 import "github.com/chris-koptional/gofig"
 
 type Config struct {
  Name  string `json:"name"`
  Age   int    `json:"age"`
 }
 
 var config Config
 
 error := gofig.GetConfig("/myconf", &config)
 
 fmt.Printf("My name : %v", config.Name)
 fmt.Printf("My age : %v", config.Age)
 
 ```

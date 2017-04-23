# gomagic
gomagic is a thin go wrapper around libmagic.

Example Usage
-------------

Initialize magic library
```
m, err := gomagic.New(gomagic.NoneFlag)
if err != nil {
  log.Fatal("Failed to initialize magic library")
}
```

Determine textual description of type of file 
```
fileName := "test_file.pdf"
result, err := m.ExamineFile(fileName)
if err != nil {
  log.Fatal("gomagic.ExamineFile: Failed to parse file")
}
fmt.Println("File is of type", result)
```

Determine mime type description of type of file 
```
err := m.SetFlags(gomagic.NodescFlag)
if err != nil {
  log.Fatal("Failed to set mime type flags on libmagic)
}
fileName := "test_file.pdf"
result, err := magic.ExamineFile(fileName)
if err != nil {
  log.Fatal("gomagic.ExamineFile: Failed to parse file")
}
fmt.Println("File is of mime type", result)
```


# ocr
To detect text from image using tesseract ocr and return true false if exist

## Steps
1. docker build . -t ocr
2. docker run -p 8080:8080 ocr
3. `curl --location --request POST 'localhost:8080/upload' \
   --form 'myFile=@"<path to jpeg file>"' \
   --form 'fatherName="xyz"'`
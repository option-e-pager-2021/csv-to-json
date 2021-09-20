# CSV-TO-JSON
Convert a CSV file with inner variables into a json file

## Install
```bash
$ go build
./csv-to-json 
Usage: input.csv delimiter
```
## Demo
```bash
$ cat pizza.csv
Stuff I like on pizza {{{{love, ham, mozza, raclette
Stuff I dislike on za {{{{hate, pineapple,,
$ csv-to-json pizza.csv '{{{{' > pizza.json
$ cat myoutput.json
{
        "hate": [
                "pineapple",
                "",
                ""
        ],
        "love": [
                "ham",
                "mozza",
                "raclette"
        ]
}
```
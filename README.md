# lottery-spreadsheets

[![Go Report Card](https://goreportcard.com/badge/github.com/178inaba/lottery-spreadsheets)](https://goreportcard.com/report/github.com/178inaba/lottery-spreadsheets)

lottery-spreadsheets makes weighted draw lots with data in the scope of the spreadsheet.  
This is an application created for validating the spreadsheet API.

## Usage

```console
$ lottery-spreadsheets -c /path/to/key.json <spreadsheets_id>
```

Targeting the following table.

| name   | weight |
|--------|-------:|
| sword  |     10 |
| spear  |     30 |
| shield |     50 |

`sword` is `A2`.  
Draw by lot and display `name`.

### Options

* `-c <json_key_file>`
  * **Required.** JSON key file of the service account.
* `-r <range>`
  * Range to use. Default: `A2:B`

## Install

```console
$ go get -u github.com/178inaba/lottery-spreadsheets
```

## License

[MIT](LICENSE)

## Author

[178inaba](https://github.com/178inaba)

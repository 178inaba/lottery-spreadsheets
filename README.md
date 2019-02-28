# lottery-spreadsheets

[![Go Report Card](https://goreportcard.com/badge/github.com/178inaba/lottery-spreadsheets)](https://goreportcard.com/report/github.com/178inaba/lottery-spreadsheets)

## Usage

```console
$ lottery-spreadsheets -c /path/to/key.json <spreadsheets_id>
```

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

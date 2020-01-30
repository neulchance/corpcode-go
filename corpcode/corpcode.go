package corpcode

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	This struct which contains the complete.
 	Array of all <list> of <result> in the file.
*/
type Corps struct {
	// XMLName xml.Name `xml:"result"`
	Corps []Corp `xml:"list"`
}

/*
	the <list> struct
*/
type Corp struct {
	// XMLName    xml.Name `xml:"list"`
	CorpName   string `xml:"corp_name"`
	CorpCode   string `xml:"corp_code"`
	StockCode  string `xml:"stock_code"`
	ModifyDate string `xml:"modify_date"`
}

func from_xml(filename string) Corps {

	// Open xmlFile
	xmlFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of the xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// initialize Corps array
	var corps Corps
	// unmarshal byteArray which contains
	// xmlFiles content into 'corps' which defined above
	xml.Unmarshal(byteValue, &corps)

	return corps
}

func UnmarshalXML(filename string) []Corp {
	corps := from_xml(filename)
	return corps.Corps
}

func CarryToMongo(corps []Corp) {
	cumulate(corps)
}

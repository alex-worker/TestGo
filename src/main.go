package main

import (
	"flag"
	"fmt"
)

const fontName string = "CaslonBold"
const mapName string = "swamp.tmx"

const resPath string = "data/"

func main() {

	resDir := resPath

	flag.StringVar(&resDir, "dir", resPath, "directory path")
	flag.Parse()

	//loadInfo := def.LoadInfo{
	//	MapName:        mapName,
	//	FontName:       fontName,
	//	ResourceFolder: resDir,
	//	ScreenSize:     screenSize,
	//}
	fmt.Println("Hello!")
	//engine.Init(loadInfo)
	//engine.Run()
}

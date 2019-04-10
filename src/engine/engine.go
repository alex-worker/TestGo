package engine

import (
	"fmt"
	"./entity"
	"./ui"
	"./def"
)

// Init engine
func Init(info def.LoadInfo){
	fmt.Println("Engine init...")
	ui.Init(info.ScreenSize)
	ui.LoadTiles(info.TileName)
	entity.LoadMap(info.MapName)
	fmt.Println(info)
}

// Run цикл
func Run(){
	
	for true {
		ui.Draw( entity.GetMap(), entity.GetHero() ) 
		if !ui.Update() {
			break
		}
	}

	ui.Destroy()
	
}
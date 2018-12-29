package cache

import (
	"fmt"

	"github.com/spf13/afero"
)

//SaveToCache -  Saving cache to .cache folder
func SaveToCache(folderName string, fileName string) {
	appFS := afero.NewOsFs()

	//folderName:=".cache"
	exists, err := afero.DirExists(appFS, folderName)

	if err != nil {
		fmt.Println("Error occured", err)
	}
	//fmt.Println(exists)
	if exists {
		fmt.Println(folderName + " already exists!!")
		return

	}
	appFS.MkdirAll(folderName, 0755)
	afero.WriteFile(appFS, folderName+"/"+fileName, []byte("file b"), 0644)
	fmt.Println("Cache created")
	//stat, statError := appFS.Stat(name)
	//fmt.Println(stat)

}

// func main() {
// 	saveToCache(".cache", "cache")
// }

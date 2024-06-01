package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"
)

const (
	SPI_SETDESKWALLPAPER = 0x0014
	SPIF_UPDATEINIFILE    = 0x01
	SPIF_SENDCHANGE       = 0x02
)

func main() {
	// Name of the image you want to set as wallpaper
	imageName := "wallpaper.jpg"

	// Get the absolute path of the image
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting the executable path: %v\n", err)
		return
	}
	imagePath := filepath.Join(filepath.Dir(executablePath), imageName)

	// Check if the image exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		fmt.Printf("Image does not exist: %s\n", imagePath)
		return
	}

	fmt.Printf("Image path: %s\n", imagePath)

	// Convert the path to a UTF-16 pointer
	imagePathPtr, err := syscall.UTF16PtrFromString(imagePath)
	if err != nil {
		fmt.Printf("Error converting the path: %v\n", err)
		return
	}

	// Call the Windows SystemParametersInfo function
	ret, _, err := syscall.NewLazyDLL("user32.dll").NewProc("SystemParametersInfoW").Call(
		SPI_SETDESKWALLPAPER,
		0,
		uintptr(unsafe.Pointer(imagePathPtr)),
		SPIF_UPDATEINIFILE|SPIF_SENDCHANGE,
	)

	if ret == 0 {
		fmt.Printf("Error changing the wallpaper: %v\n", err)
	} else {
		fmt.Println("Wallpaper changed successfully!")
	}
}

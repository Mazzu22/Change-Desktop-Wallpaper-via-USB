# Change-Desktop-Wallpaper-via-USB

This Go program changes the desktop wallpaper on a Windows machine without requiring administrator privileges. The program uses the Windows API to set a specified image as the desktop background.

## Features

- Changes the desktop wallpaper to a specified image.
- Requires no administrator privileges.
- Simple and easy to use.

## Prerequisites

- [Go](https://golang.org/dl/) (Make sure Go is installed and set up on your system).

## Usage

1. **Place the Image**: Ensure the image you want to set as the wallpaper is named `image.jpg` and is located in the same directory as the executable.

2. **Build the Program**:
   ```sh
   go build change_wallpaper.go
   

## How to use and create a automatic USB executions

1. **Download and install the USB Autorun Creator**
	https://usb-autorun-creator.it.softonic.com/
2. put the .exe file and the image on the USB
3. use the USB Autorun Creator and execute the program to complete the process

![image](https://github.com/Mazzu22/Change-Desktop-Wallpaper-via-USB/assets/72032162/5914f046-c539-468d-a925-cf8cb48b8bcb)



## Notes
- The image file should be in a format supported by Windows (e.g., JPEG, BMP).
- The program prints detailed error messages if something goes wrong, helping you troubleshoot any issues.

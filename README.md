# Onion Omega 2 + Golang
![demo](https://cdn-images-1.medium.com/max/800/1*9rGPiCljXHLlMU6fOSmxlw.gif)

![demo2](https://cdn-images-1.medium.com/max/800/1*TSsfa3xzsaL_C4k4Q04Ajw.gif)

Code for the blog post on how to get started with Golang and Onion Omega 2.

### Pre-requisites

Hardware
- Onion Omega 2
- Any LED
- 100 ohm resistor
- Jumpers

Software
- Golang installed 

### Schematic

![Fotos](https://github.com/alvarowolfx/onion-golang-blink/blob/master/schematic/schematic.png)

### Running and playing with it

To cross compile this project to run in the Onion Omega 2, follow the steps:
1. Run `make build` to generate the binary file `blink`.
2. Copy the binary to the Omega2.
3. Run it with `./blink` on the device.
4. Access http://omega-XXXX.local:9090, exchanging XXXX for the last 4 digits of the Mac address of your Omega2.
5. Have Fun.



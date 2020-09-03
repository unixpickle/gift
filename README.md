# gift

**gift** is a convenience library for reading and writing GIFs in Go. Go's standard library already has a GIF encoder and decoder, but it doesn't interpret disposal methods (meaning you may get incomplete frames), and it doesn't create paletted images automatically.

# Related packages

The [gogif](https://github.com/andybons/gogif) package was mostly incorporated into the Go standard library, but the median cut implementation was not. This package does not appear to handle disposal methods properly, although it does provide a way of quantizing an image.

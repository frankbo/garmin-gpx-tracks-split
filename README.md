# Garmin GPX Tracks Split
At the moment Garmin Connect for Android is not able to read `.gpx` with mutliple tracks. It just imports the first track and ignores the others. Therefor this tool creates per track a new gpx file `<trackname>.gpx`. Those file can then be used by Garmin Connect for Android.

## Usage
Provide the path to the `.gpx` file as an argument. See the example below.  
Run the example `go run main.go ./example/etappen.gpx`. This should create a folder `example/etappen` and write all tracks in separate files.
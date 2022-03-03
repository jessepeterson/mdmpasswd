# mdmpasswd: Apple MDM AccountConfiguration password generator

The Apple MDM command `AccountConfiguration` requires a specific property list of hashed password to work. This tool generates that Plist. Apple has the format documented [here](https://developer.apple.com/documentation/devicemanagement/passwordhash?language=objc).

## Usage

`mdmpasswd` takes a parameter for the password and outputs the plist (or, optionally, the Base-64 encoded form of the plist) to stdout.

```sh
$ go run mdmpasswd.go -h
Usage of mdmpasswd:
  -b64
    	Output base64-encoded Plist
  -password string
    	Password to hash
```

Example:

```sh
$ go run mdmpasswd.go -password 'Hello, world!'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>SALTED-SHA512-PBKDF2</key>
    <dict>
      <key>entropy</key>
      <data>LRsP3y8L1kwX9JplSvYJTWaUsZLtrpa/7NamZBWEYONqtlt1yAWPC7kd5v1sjhdkPVzrvPESiIXbmBli3RWe5WmVvxKGjN7f3yl1yZVzUgasXrV1aZZ1TVnJBpOBqFJOPGxcDA3i5Krfts1dRgPjATgbi1nPHXk9I3IsE0a02us=</data>
      <key>iterations</key>
      <integer>34501</integer>
      <key>salt</key>
      <data>dCfDCfGDT+qitaiKelndzwbfxivLvGuLG66Of9ZzwxE=</data>
    </dict>
  </dict>
</plist>
```
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
<plist version="1.0"><dict><key>entropy</key><data>rtXuhI91qDAj+ThxzBIpQrWFN2zaQb0O6t+iG3erAp4CyOAkeVrOLTtnsogYmHbBn1lYMyFJo/5dC6M5/dSre9CBBMirGsk7gIsnj/yVR/+5+rmaZxuUof7U+dWCDlBa0JR+MNCzVhmF8YJ+ogftX1K1WVS21TWMbPNFQonuCxQ=</data><key>iterations</key><integer>25327</integer><key>salt</key><data>eRBY3yyo9tG4s0P4WAveNXN2wx5J/q+K9gSKZWgV61w=</data></dict></plist>
```
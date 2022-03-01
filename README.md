# mdmpasswd: Apple MDM AccountConfiguration password generator

The Apple MDM command `AccountConfiguration` requires a specific property list of hashed password to work. This tools generates plist.

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

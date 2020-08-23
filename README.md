# License

License is a library for enforcing software licenses. Currently supports Gumroad license keys. 

[![Build Status](https://ci.j5y.xyz/api/badges/jchenry/license/status.svg)](https://ci.j5y.xyz/jchenry/license)

## Installation

```bash
go get github.com/jchenry/license
```

## Usage

```go
func main(){

  fmt.Println(license.EnforceWith(
   func(scope string) bool {
      return s == "set-license" || s == "show-license"
	},
	license.Gumroad,
	func (f license.Activation) error{
      if !f.Active{
         return fmt.Errorf("software license is not active"
      }
   })(c.Command.Name, "<someproductid>", "<somekey>"))
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)


package license

type Activation struct {
	Active  bool
	Context map[string]interface{}
}
type Provider func(product, key string) (Activation, error)
type Policy func(l Activation) error

func EnforceWith(allowed func(scope string) bool, check Provider, policy Policy) func(scope, productid, key string) error {
	return func(scope, productid, key string) error {
		if !allowed(scope) {
			if l, err := check(productid, key); err == nil {
				return policy(l)
			} else {
				return err
			}
		}
		return nil
	}
}

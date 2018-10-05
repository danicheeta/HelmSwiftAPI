package helm

import (
	"youtab/dashboard/services/config"
	"youtab/dashboard/services/initializer"
)

var url string

type manager struct {
	Name string
}

func (manager) Initialize() func() {
	url = config.Config().GetDefaultString("swift_url", "http://127.0.0.1:9855/tiller/v2/releases/")

	return nil
}

type Helm interface {
	Install(string) error
	Upgrade(string, map[string]string) error
	Values()
}

func Client(name string) Helm {
	return manager{name}
}

func init() {
	initializer.Register(manager{})
}

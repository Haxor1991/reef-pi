package controller

type Module interface {
	Name() string
	Configure(interface{}) error
	Enable() error
	Disable() error
	IsEnabled() bool
}

type LightingAPI interface {
	Enable(interface{}) error
	Disable() error
	IsEnabled() (bool, error)
	Config() interface{}
}

type NullCrudAPI struct{}

func (n *NullCrudAPI) Create(_ interface{}) error {
	return nil
}

func (n *NullCrudAPI) Get(_ string) (interface{}, error) {
	return nil, nil
}

func (n *NullCrudAPI) Update(_ string, _ interface{}) error {
	return nil
}

func (n *NullCrudAPI) Delete(_ string) error {
	return nil
}

func (n *NullCrudAPI) List() (*[]interface{}, error) {
	var ret []interface{}
	return &ret, nil
}

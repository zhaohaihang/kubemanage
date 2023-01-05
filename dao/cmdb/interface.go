package cmdb

import "gorm.io/gorm"

type CMDBFactory interface {
	Host() HostI
	HostGroup() HostGroup
	Secret() SecretI
}

func NewCMDBFactory(db *gorm.DB) CMDBFactory {
	return &cmdbFactory{db: db}
}

type cmdbFactory struct {
	db *gorm.DB
}

func (c *cmdbFactory) Host() HostI {
	return NewHost(c.db)
}

func (c *cmdbFactory) HostGroup() HostGroup {
	return NewHostGroup(c.db)
}

func (c *cmdbFactory) Secret() SecretI {
	return NewSecretI(c.db)
}

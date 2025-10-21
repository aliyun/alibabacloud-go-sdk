// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCatalog interface {
	dara.Model
	String() string
	GoString() string
	SetExtensionConf(v map[string]*string) *Catalog
	GetExtensionConf() map[string]*string
	SetName(v string) *Catalog
	GetName() *string
	SetProperties(v map[string]interface{}) *Catalog
	GetProperties() map[string]interface{}
}

type Catalog struct {
	ExtensionConf map[string]*string `json:"extensionConf,omitempty" xml:"extensionConf,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
}

func (s Catalog) String() string {
	return dara.Prettify(s)
}

func (s Catalog) GoString() string {
	return s.String()
}

func (s *Catalog) GetExtensionConf() map[string]*string {
	return s.ExtensionConf
}

func (s *Catalog) GetName() *string {
	return s.Name
}

func (s *Catalog) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *Catalog) SetExtensionConf(v map[string]*string) *Catalog {
	s.ExtensionConf = v
	return s
}

func (s *Catalog) SetName(v string) *Catalog {
	s.Name = &v
	return s
}

func (s *Catalog) SetProperties(v map[string]interface{}) *Catalog {
	s.Properties = v
	return s
}

func (s *Catalog) Validate() error {
	return dara.Validate(s)
}

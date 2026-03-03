// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCatalog interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Catalog
	GetComment() *string
	SetName(v string) *Catalog
	GetName() *string
	SetProperties(v map[string]interface{}) *Catalog
	GetProperties() map[string]interface{}
	SetProvider(v string) *Catalog
	GetProvider() *string
	SetType(v string) *Catalog
	GetType() *string
}

type Catalog struct {
	Comment    *string                `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Name       *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	Provider   *string                `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Type       *string                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Catalog) String() string {
	return dara.Prettify(s)
}

func (s Catalog) GoString() string {
	return s.String()
}

func (s *Catalog) GetComment() *string {
	return s.Comment
}

func (s *Catalog) GetName() *string {
	return s.Name
}

func (s *Catalog) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *Catalog) GetProvider() *string {
	return s.Provider
}

func (s *Catalog) GetType() *string {
	return s.Type
}

func (s *Catalog) SetComment(v string) *Catalog {
	s.Comment = &v
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

func (s *Catalog) SetProvider(v string) *Catalog {
	s.Provider = &v
	return s
}

func (s *Catalog) SetType(v string) *Catalog {
	s.Type = &v
	return s
}

func (s *Catalog) Validate() error {
	return dara.Validate(s)
}

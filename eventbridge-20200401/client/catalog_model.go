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
	SetConnectionName(v string) *Catalog
	GetConnectionName() *string
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
	// The comment or description of the data catalog
	//
	// example:
	//
	// 测试数据目录
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The connection name associated with a mounted-type Catalog. Only has a value when Provider is MySQL/PostgreSQL/Elasticsearch
	//
	// example:
	//
	// my_connection
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The unique identifier name of the data catalog
	//
	// example:
	//
	// my_catalog
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Extended properties (JSON object). The Elasticsearch type contains information such as IndexPattern
	//
	// example:
	//
	// {"IndexPattern":"my-index-*"}
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The data source provider. EventHouse is the built-in storage; MySQL/PostgreSQL/Elasticsearch are externally mounted
	//
	// example:
	//
	// EventHouse
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The type of the data catalog, such as RELATIONAL
	//
	// example:
	//
	// RELATIONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *Catalog) GetConnectionName() *string {
	return s.ConnectionName
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

func (s *Catalog) SetConnectionName(v string) *Catalog {
	s.ConnectionName = &v
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

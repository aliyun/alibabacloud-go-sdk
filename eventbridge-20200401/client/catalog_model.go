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
	// 数据目录的备注描述信息
	//
	// example:
	//
	// 测试数据目录
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 挂载类型 Catalog 关联的连接名称。仅 Provider 为 MySQL/PostgreSQL/Elasticsearch 时有值
	//
	// example:
	//
	// my_connection
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// 数据目录的唯一标识名称
	//
	// example:
	//
	// my_catalog
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 扩展属性（JSON 对象）。Elasticsearch 类型包含 IndexPattern 等信息
	//
	// example:
	//
	// {"IndexPattern":"my-index-*"}
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 数据源提供方。EventHouse 为内置存储，MySQL/PostgreSQL/Elasticsearch 为外部挂载
	//
	// example:
	//
	// EventHouse
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// 数据目录类型，如 RELATIONAL
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

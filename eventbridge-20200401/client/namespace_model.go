// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNamespace interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *Namespace
	GetCatalog() *string
	SetComment(v string) *Namespace
	GetComment() *string
	SetName(v string) *Namespace
	GetName() *string
	SetProperties(v string) *Namespace
	GetProperties() *string
}

type Namespace struct {
	// 命名空间所属的数据目录名称
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// 命名空间的备注描述信息
	//
	// example:
	//
	// 测试命名空间
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 命名空间的唯一标识名称
	//
	// example:
	//
	// my_namespace
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 命名空间的扩展属性
	//
	// example:
	//
	// {"key":"value"}
	Properties *string `json:"properties,omitempty" xml:"properties,omitempty"`
}

func (s Namespace) String() string {
	return dara.Prettify(s)
}

func (s Namespace) GoString() string {
	return s.String()
}

func (s *Namespace) GetCatalog() *string {
	return s.Catalog
}

func (s *Namespace) GetComment() *string {
	return s.Comment
}

func (s *Namespace) GetName() *string {
	return s.Name
}

func (s *Namespace) GetProperties() *string {
	return s.Properties
}

func (s *Namespace) SetCatalog(v string) *Namespace {
	s.Catalog = &v
	return s
}

func (s *Namespace) SetComment(v string) *Namespace {
	s.Comment = &v
	return s
}

func (s *Namespace) SetName(v string) *Namespace {
	s.Name = &v
	return s
}

func (s *Namespace) SetProperties(v string) *Namespace {
	s.Properties = &v
	return s
}

func (s *Namespace) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalog(v string) *UpdateNamespaceRequest
	GetCatalog() *string
	SetClientToken(v string) *UpdateNamespaceRequest
	GetClientToken() *string
	SetComment(v string) *UpdateNamespaceRequest
	GetComment() *string
	SetName(v string) *UpdateNamespaceRequest
	GetName() *string
}

type UpdateNamespaceRequest struct {
	// 命名空间所属的数据目录名称。可通过 ListCatalogs 接口获取已有目录列表
	//
	// example:
	//
	// my_catalog
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// 用于保证请求幂等性的Token，防止因网络重试导致重复操作。建议使用 UUID
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 命名空间的备注描述信息。传空字符串可清除原有备注，无格式限制
	//
	// example:
	//
	// 更新后的备注
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// 要修改的命名空间名称。名称本身不可修改，此处用于定位目标命名空间。需同时指定所属 Catalog。可通过 ListNamespaces 获取
	//
	// This parameter is required.
	//
	// example:
	//
	// my_namespace
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceRequest) GetCatalog() *string {
	return s.Catalog
}

func (s *UpdateNamespaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateNamespaceRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateNamespaceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateNamespaceRequest) SetCatalog(v string) *UpdateNamespaceRequest {
	s.Catalog = &v
	return s
}

func (s *UpdateNamespaceRequest) SetClientToken(v string) *UpdateNamespaceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateNamespaceRequest) SetComment(v string) *UpdateNamespaceRequest {
	s.Comment = &v
	return s
}

func (s *UpdateNamespaceRequest) SetName(v string) *UpdateNamespaceRequest {
	s.Name = &v
	return s
}

func (s *UpdateNamespaceRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetCatalogRequest
	GetClientToken() *string
	SetName(v string) *GetCatalogRequest
	GetName() *string
}

type GetCatalogRequest struct {
	// 用于保证请求幂等性的Token。建议使用 UUID
	//
	// example:
	//
	// 1e9b8f60-3a2c-4d7e-9f1b-8c3d5e7a2b4f
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 要查询的数据目录名称。可通过 ListCatalogs 接口获取已有目录列表
	//
	// This parameter is required.
	//
	// example:
	//
	// my_catalog
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCatalogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetCatalogRequest) GetName() *string {
	return s.Name
}

func (s *GetCatalogRequest) SetClientToken(v string) *GetCatalogRequest {
	s.ClientToken = &v
	return s
}

func (s *GetCatalogRequest) SetName(v string) *GetCatalogRequest {
	s.Name = &v
	return s
}

func (s *GetCatalogRequest) Validate() error {
	return dara.Validate(s)
}

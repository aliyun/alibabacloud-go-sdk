// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateProductRequest
	GetDescription() *string
	SetProductId(v string) *UpdateProductRequest
	GetProductId() *string
	SetProductName(v string) *UpdateProductRequest
	GetProductName() *string
	SetProviderName(v string) *UpdateProductRequest
	GetProviderName() *string
}

type UpdateProductRequest struct {
	// 产品描述
	//
	// if can be null:
	// true
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// This parameter is required.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// 产品提供方
	//
	// This parameter is required.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s UpdateProductRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProductRequest) GetProductId() *string {
	return s.ProductId
}

func (s *UpdateProductRequest) GetProductName() *string {
	return s.ProductName
}

func (s *UpdateProductRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *UpdateProductRequest) SetDescription(v string) *UpdateProductRequest {
	s.Description = &v
	return s
}

func (s *UpdateProductRequest) SetProductId(v string) *UpdateProductRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateProductRequest) SetProductName(v string) *UpdateProductRequest {
	s.ProductName = &v
	return s
}

func (s *UpdateProductRequest) SetProviderName(v string) *UpdateProductRequest {
	s.ProviderName = &v
	return s
}

func (s *UpdateProductRequest) Validate() error {
	return dara.Validate(s)
}

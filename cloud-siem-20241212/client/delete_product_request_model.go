// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteProductRequest
	GetLang() *string
	SetProductId(v string) *DeleteProductRequest
	GetProductId() *string
	SetRegionId(v string) *DeleteProductRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteProductRequest
	GetRoleFor() *int64
}

type DeleteProductRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeleteProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteProductRequest) GetProductId() *string {
	return s.ProductId
}

func (s *DeleteProductRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteProductRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteProductRequest) SetLang(v string) *DeleteProductRequest {
	s.Lang = &v
	return s
}

func (s *DeleteProductRequest) SetProductId(v string) *DeleteProductRequest {
	s.ProductId = &v
	return s
}

func (s *DeleteProductRequest) SetRegionId(v string) *DeleteProductRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteProductRequest) SetRoleFor(v int64) *DeleteProductRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteProductRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateProductRequest
	GetLang() *string
	SetProductName(v string) *CreateProductRequest
	GetProductName() *string
	SetRegionId(v string) *CreateProductRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateProductRequest
	GetRoleFor() *int64
	SetVendorName(v string) *CreateProductRequest
	GetVendorName() *string
}

type CreateProductRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// alibaba_cloud_sas。
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 111。
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s CreateProductRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateProductRequest) GetProductName() *string {
	return s.ProductName
}

func (s *CreateProductRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateProductRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateProductRequest) GetVendorName() *string {
	return s.VendorName
}

func (s *CreateProductRequest) SetLang(v string) *CreateProductRequest {
	s.Lang = &v
	return s
}

func (s *CreateProductRequest) SetProductName(v string) *CreateProductRequest {
	s.ProductName = &v
	return s
}

func (s *CreateProductRequest) SetRegionId(v string) *CreateProductRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProductRequest) SetRoleFor(v int64) *CreateProductRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateProductRequest) SetVendorName(v string) *CreateProductRequest {
	s.VendorName = &v
	return s
}

func (s *CreateProductRequest) Validate() error {
	return dara.Validate(s)
}

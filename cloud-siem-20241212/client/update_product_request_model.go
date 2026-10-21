// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateProductRequest
	GetLang() *string
	SetProductId(v string) *UpdateProductRequest
	GetProductId() *string
	SetProductName(v string) *UpdateProductRequest
	GetProductName() *string
	SetRegionId(v string) *UpdateProductRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateProductRequest
	GetRoleFor() *int64
	SetVendorName(v string) *UpdateProductRequest
	GetVendorName() *string
}

type UpdateProductRequest struct {
	// The language of the response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English. This parameter is required. If not provided, the service returns a 400 error.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The product ID. Only products with ProductType set to custom support updates. System preset products (ProductType set to preset) cannot be modified. Calling this operation on a preset product returns a 403 error. To create a custom product, use the CreateProduct API first. This parameter is the resource identifier for the Update operation and is required. Only custom products can be updated. System preset products cannot be updated.
	//
	// example:
	//
	// alibaba_cloud_sas
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The product name.
	//
	// example:
	//
	// alibaba_cloud_sas
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The region where the data management center of Cloud Threat Detection and Response (CTDR) is located. Select the management center region based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator uses to switch to another member\\"s view.
	//
	// example:
	//
	// 1733269771123
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The vendor name.
	//
	// example:
	//
	// 111
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s UpdateProductRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProductRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateProductRequest) GetProductId() *string {
	return s.ProductId
}

func (s *UpdateProductRequest) GetProductName() *string {
	return s.ProductName
}

func (s *UpdateProductRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateProductRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateProductRequest) GetVendorName() *string {
	return s.VendorName
}

func (s *UpdateProductRequest) SetLang(v string) *UpdateProductRequest {
	s.Lang = &v
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

func (s *UpdateProductRequest) SetRegionId(v string) *UpdateProductRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateProductRequest) SetRoleFor(v int64) *UpdateProductRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateProductRequest) SetVendorName(v string) *UpdateProductRequest {
	s.VendorName = &v
	return s
}

func (s *UpdateProductRequest) Validate() error {
	return dara.Validate(s)
}

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
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The product ID.
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
	// The region of the Data Management center for threat analysis. Select the region for the Data Management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: The Chinese mainland.
	//
	// - ap-southeast-1: Regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can specify this parameter to switch to the perspective of the member.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVendorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateVendorRequest
	GetLang() *string
	SetRegionId(v string) *UpdateVendorRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateVendorRequest
	GetRoleFor() *int64
	SetVendorId(v string) *UpdateVendorRequest
	GetVendorId() *string
	SetVendorName(v string) *UpdateVendorRequest
	GetVendorName() *string
}

type UpdateVendorRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region of the Data Management center for threat analysis. Select the region based on where your asset is located. Valid values:
	//
	// - cn-hangzhou: The asset is in the Chinese mainland.
	//
	// - ap-southeast-1: The asset is outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. This parameter allows an administrator to switch to the perspective of a member.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The ID of the vendor.
	//
	// example:
	//
	// vd-qlsw5eocx94w9
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
	// The name of the vendor.
	//
	// example:
	//
	// 111
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s UpdateVendorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVendorRequest) GoString() string {
	return s.String()
}

func (s *UpdateVendorRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateVendorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVendorRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateVendorRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *UpdateVendorRequest) GetVendorName() *string {
	return s.VendorName
}

func (s *UpdateVendorRequest) SetLang(v string) *UpdateVendorRequest {
	s.Lang = &v
	return s
}

func (s *UpdateVendorRequest) SetRegionId(v string) *UpdateVendorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVendorRequest) SetRoleFor(v int64) *UpdateVendorRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateVendorRequest) SetVendorId(v string) *UpdateVendorRequest {
	s.VendorId = &v
	return s
}

func (s *UpdateVendorRequest) SetVendorName(v string) *UpdateVendorRequest {
	s.VendorName = &v
	return s
}

func (s *UpdateVendorRequest) Validate() error {
	return dara.Validate(s)
}

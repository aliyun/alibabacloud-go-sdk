// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVendorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateVendorRequest
	GetLang() *string
	SetRegionId(v string) *CreateVendorRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateVendorRequest
	GetRoleFor() *int64
	SetVendorName(v string) *CreateVendorRequest
	GetVendorName() *string
}

type CreateVendorRequest struct {
	// The language of the response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region of the Data Management center. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: The assets are in the Chinese mainland.
	//
	// - ap-southeast-1: The assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can use this parameter to switch to the member\\"s view.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The name of the vendor.
	//
	// example:
	//
	// 111
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s CreateVendorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVendorRequest) GoString() string {
	return s.String()
}

func (s *CreateVendorRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateVendorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVendorRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateVendorRequest) GetVendorName() *string {
	return s.VendorName
}

func (s *CreateVendorRequest) SetLang(v string) *CreateVendorRequest {
	s.Lang = &v
	return s
}

func (s *CreateVendorRequest) SetRegionId(v string) *CreateVendorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVendorRequest) SetRoleFor(v int64) *CreateVendorRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateVendorRequest) SetVendorName(v string) *CreateVendorRequest {
	s.VendorName = &v
	return s
}

func (s *CreateVendorRequest) Validate() error {
	return dara.Validate(s)
}

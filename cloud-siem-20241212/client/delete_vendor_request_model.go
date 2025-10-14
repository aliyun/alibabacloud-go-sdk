// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVendorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteVendorRequest
	GetLang() *string
	SetRegionId(v string) *DeleteVendorRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteVendorRequest
	GetRoleFor() *int64
	SetVendorId(v string) *DeleteVendorRequest
	GetVendorId() *string
	SetVendorName(v string) *DeleteVendorRequest
	GetVendorName() *string
}

type DeleteVendorRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// vd-qlsw5eocx94w9。
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
	// example:
	//
	// 111。
	VendorName *string `json:"VendorName,omitempty" xml:"VendorName,omitempty"`
}

func (s DeleteVendorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVendorRequest) GoString() string {
	return s.String()
}

func (s *DeleteVendorRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteVendorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVendorRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteVendorRequest) GetVendorId() *string {
	return s.VendorId
}

func (s *DeleteVendorRequest) GetVendorName() *string {
	return s.VendorName
}

func (s *DeleteVendorRequest) SetLang(v string) *DeleteVendorRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVendorRequest) SetRegionId(v string) *DeleteVendorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVendorRequest) SetRoleFor(v int64) *DeleteVendorRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteVendorRequest) SetVendorId(v string) *DeleteVendorRequest {
	s.VendorId = &v
	return s
}

func (s *DeleteVendorRequest) SetVendorName(v string) *DeleteVendorRequest {
	s.VendorName = &v
	return s
}

func (s *DeleteVendorRequest) Validate() error {
	return dara.Validate(s)
}

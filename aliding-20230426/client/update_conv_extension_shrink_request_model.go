// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConvExtensionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobileUrl(v string) *UpdateConvExtensionShrinkRequest
	GetMobileUrl() *string
	SetPcUrl(v string) *UpdateConvExtensionShrinkRequest
	GetPcUrl() *string
	SetStaffIdListShrink(v string) *UpdateConvExtensionShrinkRequest
	GetStaffIdListShrink() *string
	SetSystemUid(v string) *UpdateConvExtensionShrinkRequest
	GetSystemUid() *string
	SetTenantContextShrink(v string) *UpdateConvExtensionShrinkRequest
	GetTenantContextShrink() *string
}

type UpdateConvExtensionShrinkRequest struct {
	// example:
	//
	// https://xxx
	MobileUrl *string `json:"MobileUrl,omitempty" xml:"MobileUrl,omitempty"`
	// example:
	//
	// https://xxx
	PcUrl             *string `json:"PcUrl,omitempty" xml:"PcUrl,omitempty"`
	StaffIdListShrink *string `json:"StaffIdList,omitempty" xml:"StaffIdList,omitempty"`
	// example:
	//
	// 546374856
	SystemUid           *string `json:"SystemUid,omitempty" xml:"SystemUid,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s UpdateConvExtensionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConvExtensionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConvExtensionShrinkRequest) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *UpdateConvExtensionShrinkRequest) GetPcUrl() *string {
	return s.PcUrl
}

func (s *UpdateConvExtensionShrinkRequest) GetStaffIdListShrink() *string {
	return s.StaffIdListShrink
}

func (s *UpdateConvExtensionShrinkRequest) GetSystemUid() *string {
	return s.SystemUid
}

func (s *UpdateConvExtensionShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateConvExtensionShrinkRequest) SetMobileUrl(v string) *UpdateConvExtensionShrinkRequest {
	s.MobileUrl = &v
	return s
}

func (s *UpdateConvExtensionShrinkRequest) SetPcUrl(v string) *UpdateConvExtensionShrinkRequest {
	s.PcUrl = &v
	return s
}

func (s *UpdateConvExtensionShrinkRequest) SetStaffIdListShrink(v string) *UpdateConvExtensionShrinkRequest {
	s.StaffIdListShrink = &v
	return s
}

func (s *UpdateConvExtensionShrinkRequest) SetSystemUid(v string) *UpdateConvExtensionShrinkRequest {
	s.SystemUid = &v
	return s
}

func (s *UpdateConvExtensionShrinkRequest) SetTenantContextShrink(v string) *UpdateConvExtensionShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateConvExtensionShrinkRequest) Validate() error {
	return dara.Validate(s)
}

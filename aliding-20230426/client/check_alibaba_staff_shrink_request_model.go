// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAlibabaStaffShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v string) *CheckAlibabaStaffShrinkRequest
	GetMobile() *string
	SetTenantContextShrink(v string) *CheckAlibabaStaffShrinkRequest
	GetTenantContextShrink() *string
}

type CheckAlibabaStaffShrinkRequest struct {
	// example:
	//
	// 156****9665
	Mobile              *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s CheckAlibabaStaffShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAlibabaStaffShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckAlibabaStaffShrinkRequest) GetMobile() *string {
	return s.Mobile
}

func (s *CheckAlibabaStaffShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CheckAlibabaStaffShrinkRequest) SetMobile(v string) *CheckAlibabaStaffShrinkRequest {
	s.Mobile = &v
	return s
}

func (s *CheckAlibabaStaffShrinkRequest) SetTenantContextShrink(v string) *CheckAlibabaStaffShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CheckAlibabaStaffShrinkRequest) Validate() error {
	return dara.Validate(s)
}

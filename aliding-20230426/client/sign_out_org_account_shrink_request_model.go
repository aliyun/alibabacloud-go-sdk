// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignOutOrgAccountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReason(v string) *SignOutOrgAccountShrinkRequest
	GetReason() *string
	SetReasonI18nForEmployeeShrink(v string) *SignOutOrgAccountShrinkRequest
	GetReasonI18nForEmployeeShrink() *string
	SetTenantContextShrink(v string) *SignOutOrgAccountShrinkRequest
	GetTenantContextShrink() *string
}

type SignOutOrgAccountShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 企业安全合规要求，执行账号强制登出
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// {"zh_CN":"管理员已执行登出","en_US":"Signed out by admin"}
	ReasonI18nForEmployeeShrink *string `json:"ReasonI18nForEmployee,omitempty" xml:"ReasonI18nForEmployee,omitempty"`
	TenantContextShrink         *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s SignOutOrgAccountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SignOutOrgAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *SignOutOrgAccountShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *SignOutOrgAccountShrinkRequest) GetReasonI18nForEmployeeShrink() *string {
	return s.ReasonI18nForEmployeeShrink
}

func (s *SignOutOrgAccountShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SignOutOrgAccountShrinkRequest) SetReason(v string) *SignOutOrgAccountShrinkRequest {
	s.Reason = &v
	return s
}

func (s *SignOutOrgAccountShrinkRequest) SetReasonI18nForEmployeeShrink(v string) *SignOutOrgAccountShrinkRequest {
	s.ReasonI18nForEmployeeShrink = &v
	return s
}

func (s *SignOutOrgAccountShrinkRequest) SetTenantContextShrink(v string) *SignOutOrgAccountShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SignOutOrgAccountShrinkRequest) Validate() error {
	return dara.Validate(s)
}

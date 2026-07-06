// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignOutOrgAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReason(v string) *SignOutOrgAccountRequest
	GetReason() *string
	SetReasonI18nForEmployee(v map[string]*string) *SignOutOrgAccountRequest
	GetReasonI18nForEmployee() map[string]*string
	SetTenantContext(v *SignOutOrgAccountRequestTenantContext) *SignOutOrgAccountRequest
	GetTenantContext() *SignOutOrgAccountRequestTenantContext
}

type SignOutOrgAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 企业安全合规要求，执行账号强制登出
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// {"zh_CN":"管理员已执行登出","en_US":"Signed out by admin"}
	ReasonI18nForEmployee map[string]*string                     `json:"ReasonI18nForEmployee,omitempty" xml:"ReasonI18nForEmployee,omitempty"`
	TenantContext         *SignOutOrgAccountRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s SignOutOrgAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s SignOutOrgAccountRequest) GoString() string {
	return s.String()
}

func (s *SignOutOrgAccountRequest) GetReason() *string {
	return s.Reason
}

func (s *SignOutOrgAccountRequest) GetReasonI18nForEmployee() map[string]*string {
	return s.ReasonI18nForEmployee
}

func (s *SignOutOrgAccountRequest) GetTenantContext() *SignOutOrgAccountRequestTenantContext {
	return s.TenantContext
}

func (s *SignOutOrgAccountRequest) SetReason(v string) *SignOutOrgAccountRequest {
	s.Reason = &v
	return s
}

func (s *SignOutOrgAccountRequest) SetReasonI18nForEmployee(v map[string]*string) *SignOutOrgAccountRequest {
	s.ReasonI18nForEmployee = v
	return s
}

func (s *SignOutOrgAccountRequest) SetTenantContext(v *SignOutOrgAccountRequestTenantContext) *SignOutOrgAccountRequest {
	s.TenantContext = v
	return s
}

func (s *SignOutOrgAccountRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SignOutOrgAccountRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SignOutOrgAccountRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SignOutOrgAccountRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SignOutOrgAccountRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SignOutOrgAccountRequestTenantContext) SetTenantId(v string) *SignOutOrgAccountRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SignOutOrgAccountRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

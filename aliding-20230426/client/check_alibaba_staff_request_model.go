// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAlibabaStaffRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobile(v string) *CheckAlibabaStaffRequest
	GetMobile() *string
	SetTenantContext(v *CheckAlibabaStaffRequestTenantContext) *CheckAlibabaStaffRequest
	GetTenantContext() *CheckAlibabaStaffRequestTenantContext
}

type CheckAlibabaStaffRequest struct {
	// example:
	//
	// 156****9665
	Mobile        *string                                `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	TenantContext *CheckAlibabaStaffRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s CheckAlibabaStaffRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAlibabaStaffRequest) GoString() string {
	return s.String()
}

func (s *CheckAlibabaStaffRequest) GetMobile() *string {
	return s.Mobile
}

func (s *CheckAlibabaStaffRequest) GetTenantContext() *CheckAlibabaStaffRequestTenantContext {
	return s.TenantContext
}

func (s *CheckAlibabaStaffRequest) SetMobile(v string) *CheckAlibabaStaffRequest {
	s.Mobile = &v
	return s
}

func (s *CheckAlibabaStaffRequest) SetTenantContext(v *CheckAlibabaStaffRequestTenantContext) *CheckAlibabaStaffRequest {
	s.TenantContext = v
	return s
}

func (s *CheckAlibabaStaffRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckAlibabaStaffRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CheckAlibabaStaffRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CheckAlibabaStaffRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CheckAlibabaStaffRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CheckAlibabaStaffRequestTenantContext) SetTenantId(v string) *CheckAlibabaStaffRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CheckAlibabaStaffRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchMainOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTargetOrgId(v int64) *SwitchMainOrgRequest
	GetTargetOrgId() *int64
	SetTenantContext(v *SwitchMainOrgRequestTenantContext) *SwitchMainOrgRequest
	GetTenantContext() *SwitchMainOrgRequestTenantContext
}

type SwitchMainOrgRequest struct {
	// example:
	//
	// 21001
	TargetOrgId   *int64                             `json:"TargetOrgId,omitempty" xml:"TargetOrgId,omitempty"`
	TenantContext *SwitchMainOrgRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s SwitchMainOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchMainOrgRequest) GoString() string {
	return s.String()
}

func (s *SwitchMainOrgRequest) GetTargetOrgId() *int64 {
	return s.TargetOrgId
}

func (s *SwitchMainOrgRequest) GetTenantContext() *SwitchMainOrgRequestTenantContext {
	return s.TenantContext
}

func (s *SwitchMainOrgRequest) SetTargetOrgId(v int64) *SwitchMainOrgRequest {
	s.TargetOrgId = &v
	return s
}

func (s *SwitchMainOrgRequest) SetTenantContext(v *SwitchMainOrgRequestTenantContext) *SwitchMainOrgRequest {
	s.TenantContext = v
	return s
}

func (s *SwitchMainOrgRequest) Validate() error {
	return dara.Validate(s)
}

type SwitchMainOrgRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SwitchMainOrgRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SwitchMainOrgRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SwitchMainOrgRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SwitchMainOrgRequestTenantContext) SetTenantId(v string) *SwitchMainOrgRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SwitchMainOrgRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

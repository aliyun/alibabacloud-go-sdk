// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchMainOrgShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTargetOrgId(v int64) *SwitchMainOrgShrinkRequest
	GetTargetOrgId() *int64
	SetTenantContextShrink(v string) *SwitchMainOrgShrinkRequest
	GetTenantContextShrink() *string
}

type SwitchMainOrgShrinkRequest struct {
	// example:
	//
	// 21001
	TargetOrgId         *int64  `json:"TargetOrgId,omitempty" xml:"TargetOrgId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s SwitchMainOrgShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchMainOrgShrinkRequest) GoString() string {
	return s.String()
}

func (s *SwitchMainOrgShrinkRequest) GetTargetOrgId() *int64 {
	return s.TargetOrgId
}

func (s *SwitchMainOrgShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SwitchMainOrgShrinkRequest) SetTargetOrgId(v int64) *SwitchMainOrgShrinkRequest {
	s.TargetOrgId = &v
	return s
}

func (s *SwitchMainOrgShrinkRequest) SetTenantContextShrink(v string) *SwitchMainOrgShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SwitchMainOrgShrinkRequest) Validate() error {
	return dara.Validate(s)
}

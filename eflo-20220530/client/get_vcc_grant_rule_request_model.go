// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVccGrantRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *GetVccGrantRuleRequest
	GetErId() *string
	SetGrantRuleId(v string) *GetVccGrantRuleRequest
	GetGrantRuleId() *string
	SetGrantTenantId(v string) *GetVccGrantRuleRequest
	GetGrantTenantId() *string
	SetInstanceId(v string) *GetVccGrantRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetVccGrantRuleRequest
	GetRegionId() *string
}

type GetVccGrantRuleRequest struct {
	// Lingjun HUB Instance ID
	//
	// example:
	//
	// er-aueyxxsy
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Resource Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// grant-rule-jaj34d75h01
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Authorized Instance ID
	//
	// example:
	//
	// vcc-cn-jaj34d75h01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetVccGrantRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVccGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *GetVccGrantRuleRequest) GetErId() *string {
	return s.ErId
}

func (s *GetVccGrantRuleRequest) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *GetVccGrantRuleRequest) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *GetVccGrantRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVccGrantRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVccGrantRuleRequest) SetErId(v string) *GetVccGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *GetVccGrantRuleRequest) SetGrantRuleId(v string) *GetVccGrantRuleRequest {
	s.GrantRuleId = &v
	return s
}

func (s *GetVccGrantRuleRequest) SetGrantTenantId(v string) *GetVccGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *GetVccGrantRuleRequest) SetInstanceId(v string) *GetVccGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVccGrantRuleRequest) SetRegionId(v string) *GetVccGrantRuleRequest {
	s.RegionId = &v
	return s
}

func (s *GetVccGrantRuleRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpdGrantRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *GetVpdGrantRuleRequest
	GetErId() *string
	SetGrantRuleId(v string) *GetVpdGrantRuleRequest
	GetGrantRuleId() *string
	SetGrantTenantId(v string) *GetVpdGrantRuleRequest
	GetGrantTenantId() *string
	SetInstanceId(v string) *GetVpdGrantRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetVpdGrantRuleRequest
	GetRegionId() *string
}

type GetVpdGrantRuleRequest struct {
	// Lingjun HUB Instance Id
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Resource Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// grant-rule-xrgvqazb
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
	// vpd-xxxxxxxx
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

func (s GetVpdGrantRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpdGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *GetVpdGrantRuleRequest) GetErId() *string {
	return s.ErId
}

func (s *GetVpdGrantRuleRequest) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *GetVpdGrantRuleRequest) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *GetVpdGrantRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVpdGrantRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpdGrantRuleRequest) SetErId(v string) *GetVpdGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *GetVpdGrantRuleRequest) SetGrantRuleId(v string) *GetVpdGrantRuleRequest {
	s.GrantRuleId = &v
	return s
}

func (s *GetVpdGrantRuleRequest) SetGrantTenantId(v string) *GetVpdGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *GetVpdGrantRuleRequest) SetInstanceId(v string) *GetVpdGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVpdGrantRuleRequest) SetRegionId(v string) *GetVpdGrantRuleRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpdGrantRuleRequest) Validate() error {
	return dara.Validate(s)
}

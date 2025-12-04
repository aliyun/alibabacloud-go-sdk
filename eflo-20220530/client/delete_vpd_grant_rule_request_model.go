// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpdGrantRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *DeleteVpdGrantRuleRequest
	GetErId() *string
	SetGrantRuleId(v string) *DeleteVpdGrantRuleRequest
	GetGrantRuleId() *string
	SetGrantTenantId(v string) *DeleteVpdGrantRuleRequest
	GetGrantTenantId() *string
	SetInstanceId(v string) *DeleteVpdGrantRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteVpdGrantRuleRequest
	GetRegionId() *string
}

type DeleteVpdGrantRuleRequest struct {
	// Lingjun HUB ID
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorization Entry ID
	//
	// This parameter is required.
	//
	// example:
	//
	// grant-rule-9rgxqazb
	GrantRuleId *string `json:"GrantRuleId,omitempty" xml:"GrantRuleId,omitempty"`
	// Authorized Tenant ID
	//
	// example:
	//
	// 1013666993027780
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// example:
	//
	// vpd-8rgvqazb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteVpdGrantRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpdGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpdGrantRuleRequest) GetErId() *string {
	return s.ErId
}

func (s *DeleteVpdGrantRuleRequest) GetGrantRuleId() *string {
	return s.GrantRuleId
}

func (s *DeleteVpdGrantRuleRequest) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *DeleteVpdGrantRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteVpdGrantRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVpdGrantRuleRequest) SetErId(v string) *DeleteVpdGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *DeleteVpdGrantRuleRequest) SetGrantRuleId(v string) *DeleteVpdGrantRuleRequest {
	s.GrantRuleId = &v
	return s
}

func (s *DeleteVpdGrantRuleRequest) SetGrantTenantId(v string) *DeleteVpdGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *DeleteVpdGrantRuleRequest) SetInstanceId(v string) *DeleteVpdGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteVpdGrantRuleRequest) SetRegionId(v string) *DeleteVpdGrantRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVpdGrantRuleRequest) Validate() error {
	return dara.Validate(s)
}

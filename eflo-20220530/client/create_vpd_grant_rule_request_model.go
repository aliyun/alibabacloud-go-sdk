// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpdGrantRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *CreateVpdGrantRuleRequest
	GetErId() *string
	SetGrantTenantId(v string) *CreateVpdGrantRuleRequest
	GetGrantTenantId() *string
	SetInstanceId(v string) *CreateVpdGrantRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateVpdGrantRuleRequest
	GetRegionId() *string
}

type CreateVpdGrantRuleRequest struct {
	// Lingjun HUB ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// Authorized Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1013666993027780
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// This parameter is required.
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

func (s CreateVpdGrantRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpdGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateVpdGrantRuleRequest) GetErId() *string {
	return s.ErId
}

func (s *CreateVpdGrantRuleRequest) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *CreateVpdGrantRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVpdGrantRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpdGrantRuleRequest) SetErId(v string) *CreateVpdGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *CreateVpdGrantRuleRequest) SetGrantTenantId(v string) *CreateVpdGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *CreateVpdGrantRuleRequest) SetInstanceId(v string) *CreateVpdGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVpdGrantRuleRequest) SetRegionId(v string) *CreateVpdGrantRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpdGrantRuleRequest) Validate() error {
	return dara.Validate(s)
}

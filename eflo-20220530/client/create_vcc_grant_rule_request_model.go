// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVccGrantRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *CreateVccGrantRuleRequest
	GetErId() *string
	SetGrantTenantId(v string) *CreateVccGrantRuleRequest
	GetGrantTenantId() *string
	SetInstanceId(v string) *CreateVccGrantRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateVccGrantRuleRequest
	GetRegionId() *string
}

type CreateVccGrantRuleRequest struct {
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
	// 1620939556166277
	GrantTenantId *string `json:"GrantTenantId,omitempty" xml:"GrantTenantId,omitempty"`
	// Network Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-8rgvqazb
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

func (s CreateVccGrantRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVccGrantRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateVccGrantRuleRequest) GetErId() *string {
	return s.ErId
}

func (s *CreateVccGrantRuleRequest) GetGrantTenantId() *string {
	return s.GrantTenantId
}

func (s *CreateVccGrantRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVccGrantRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVccGrantRuleRequest) SetErId(v string) *CreateVccGrantRuleRequest {
	s.ErId = &v
	return s
}

func (s *CreateVccGrantRuleRequest) SetGrantTenantId(v string) *CreateVccGrantRuleRequest {
	s.GrantTenantId = &v
	return s
}

func (s *CreateVccGrantRuleRequest) SetInstanceId(v string) *CreateVccGrantRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVccGrantRuleRequest) SetRegionId(v string) *CreateVccGrantRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVccGrantRuleRequest) Validate() error {
	return dara.Validate(s)
}

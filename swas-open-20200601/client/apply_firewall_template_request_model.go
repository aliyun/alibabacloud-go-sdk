// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFirewallTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ApplyFirewallTemplateRequest
	GetClientToken() *string
	SetFirewallTemplateId(v string) *ApplyFirewallTemplateRequest
	GetFirewallTemplateId() *string
	SetInstanceIds(v []*string) *ApplyFirewallTemplateRequest
	GetInstanceIds() []*string
	SetRegionId(v string) *ApplyFirewallTemplateRequest
	GetRegionId() *string
}

type ApplyFirewallTemplateRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can generate the value on your client and must ensure that the value is unique among different requests. The **ClientToken*	- parameter supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The instance IDs of the Simple Application Server instances.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApplyFirewallTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyFirewallTemplateRequest) GoString() string {
	return s.String()
}

func (s *ApplyFirewallTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ApplyFirewallTemplateRequest) GetFirewallTemplateId() *string {
	return s.FirewallTemplateId
}

func (s *ApplyFirewallTemplateRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ApplyFirewallTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ApplyFirewallTemplateRequest) SetClientToken(v string) *ApplyFirewallTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyFirewallTemplateRequest) SetFirewallTemplateId(v string) *ApplyFirewallTemplateRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *ApplyFirewallTemplateRequest) SetInstanceIds(v []*string) *ApplyFirewallTemplateRequest {
	s.InstanceIds = v
	return s
}

func (s *ApplyFirewallTemplateRequest) SetRegionId(v string) *ApplyFirewallTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ApplyFirewallTemplateRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteFirewallRuleRequest
	GetClientToken() *string
	SetInstanceId(v string) *DeleteFirewallRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteFirewallRuleRequest
	GetRegionId() *string
	SetRuleId(v string) *DeleteFirewallRuleRequest
	GetRuleId() *string
}

type DeleteFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the specified simple application server.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The firewall rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteFirewallRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFirewallRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteFirewallRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFirewallRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteFirewallRuleRequest) SetClientToken(v string) *DeleteFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallRuleRequest) SetInstanceId(v string) *DeleteFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFirewallRuleRequest) SetRegionId(v string) *DeleteFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFirewallRuleRequest) SetRuleId(v string) *DeleteFirewallRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteFirewallRuleRequest) Validate() error {
	return dara.Validate(s)
}

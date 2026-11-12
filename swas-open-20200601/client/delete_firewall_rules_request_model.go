// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteFirewallRulesRequest
	GetClientToken() *string
	SetInstanceId(v string) *DeleteFirewallRulesRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteFirewallRulesRequest
	GetRegionId() *string
	SetRuleIds(v []*string) *DeleteFirewallRulesRequest
	GetRuleIds() []*string
}

type DeleteFirewallRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the firewall rules that you want to delete.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DeleteFirewallRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFirewallRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteFirewallRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteFirewallRulesRequest) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *DeleteFirewallRulesRequest) SetClientToken(v string) *DeleteFirewallRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetInstanceId(v string) *DeleteFirewallRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetRegionId(v string) *DeleteFirewallRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetRuleIds(v []*string) *DeleteFirewallRulesRequest {
	s.RuleIds = v
	return s
}

func (s *DeleteFirewallRulesRequest) Validate() error {
	return dara.Validate(s)
}

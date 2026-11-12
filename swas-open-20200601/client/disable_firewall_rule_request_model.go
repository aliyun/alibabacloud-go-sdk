// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableFirewallRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DisableFirewallRuleRequest
	GetClientToken() *string
	SetInstanceId(v string) *DisableFirewallRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *DisableFirewallRuleRequest
	GetRegionId() *string
	SetRemark(v string) *DisableFirewallRuleRequest
	GetRemark() *string
	SetRuleId(v string) *DisableFirewallRuleRequest
	GetRuleId() *string
}

type DisableFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. A client-generated parameter value that is unique among different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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
	// The region ID of the specified simple application server. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// Custom
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The firewall rule ID. You can call the ListFirewallRules operation to obtain the firewall rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableFirewallRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableFirewallRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisableFirewallRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableFirewallRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableFirewallRuleRequest) GetRemark() *string {
	return s.Remark
}

func (s *DisableFirewallRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DisableFirewallRuleRequest) SetClientToken(v string) *DisableFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableFirewallRuleRequest) SetInstanceId(v string) *DisableFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableFirewallRuleRequest) SetRegionId(v string) *DisableFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DisableFirewallRuleRequest) SetRemark(v string) *DisableFirewallRuleRequest {
	s.Remark = &v
	return s
}

func (s *DisableFirewallRuleRequest) SetRuleId(v string) *DisableFirewallRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DisableFirewallRuleRequest) Validate() error {
	return dara.Validate(s)
}

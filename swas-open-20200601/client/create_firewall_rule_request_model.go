// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateFirewallRuleRequest
	GetClientToken() *string
	SetInstanceId(v string) *CreateFirewallRuleRequest
	GetInstanceId() *string
	SetPort(v string) *CreateFirewallRuleRequest
	GetPort() *string
	SetRegionId(v string) *CreateFirewallRuleRequest
	GetRegionId() *string
	SetRemark(v string) *CreateFirewallRuleRequest
	GetRemark() *string
	SetRuleProtocol(v string) *CreateFirewallRuleRequest
	GetRuleProtocol() *string
}

type CreateFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the specified simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The port range.
	//
	// - TCP/UDP: Valid values: 1 to 65535. Use a forward slash (/) to separate the start port and the end port. For example, 1024/1055 indicates the port range from 1024 to 1055.
	//
	// - ICMP: -1/-1. This value indicates all ports.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306/3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the specified simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remark for the firewall rule.
	//
	// example:
	//
	// TEST
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// - TCP: TCP
	//
	// - UDP: UDP
	//
	// - TCP+UDP: TCP and UDP
	//
	// - ICMP: ICMP
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
}

func (s CreateFirewallRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFirewallRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFirewallRuleRequest) GetPort() *string {
	return s.Port
}

func (s *CreateFirewallRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFirewallRuleRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateFirewallRuleRequest) GetRuleProtocol() *string {
	return s.RuleProtocol
}

func (s *CreateFirewallRuleRequest) SetClientToken(v string) *CreateFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetInstanceId(v string) *CreateFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetPort(v string) *CreateFirewallRuleRequest {
	s.Port = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetRegionId(v string) *CreateFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetRemark(v string) *CreateFirewallRuleRequest {
	s.Remark = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetRuleProtocol(v string) *CreateFirewallRuleRequest {
	s.RuleProtocol = &v
	return s
}

func (s *CreateFirewallRuleRequest) Validate() error {
	return dara.Validate(s)
}

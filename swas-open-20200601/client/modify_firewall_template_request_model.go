// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyFirewallTemplateRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyFirewallTemplateRequest
	GetDescription() *string
	SetFirewallTemplateId(v string) *ModifyFirewallTemplateRequest
	GetFirewallTemplateId() *string
	SetFirewallTemplateRule(v []*ModifyFirewallTemplateRequestFirewallTemplateRule) *ModifyFirewallTemplateRequest
	GetFirewallTemplateRule() []*ModifyFirewallTemplateRequestFirewallTemplateRule
	SetName(v string) *ModifyFirewallTemplateRequest
	GetName() *string
	SetRegionId(v string) *ModifyFirewallTemplateRequest
	GetRegionId() *string
}

type ModifyFirewallTemplateRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, and you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the firewall template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The firewall rules.
	FirewallTemplateRule []*ModifyFirewallTemplateRequestFirewallTemplateRule `json:"FirewallTemplateRule,omitempty" xml:"FirewallTemplateRule,omitempty" type:"Repeated"`
	// The name of the firewall template.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the Simple Application Server instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyFirewallTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirewallTemplateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyFirewallTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFirewallTemplateRequest) GetFirewallTemplateId() *string {
	return s.FirewallTemplateId
}

func (s *ModifyFirewallTemplateRequest) GetFirewallTemplateRule() []*ModifyFirewallTemplateRequestFirewallTemplateRule {
	return s.FirewallTemplateRule
}

func (s *ModifyFirewallTemplateRequest) GetName() *string {
	return s.Name
}

func (s *ModifyFirewallTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyFirewallTemplateRequest) SetClientToken(v string) *ModifyFirewallTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetDescription(v string) *ModifyFirewallTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetFirewallTemplateId(v string) *ModifyFirewallTemplateRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetFirewallTemplateRule(v []*ModifyFirewallTemplateRequestFirewallTemplateRule) *ModifyFirewallTemplateRequest {
	s.FirewallTemplateRule = v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetName(v string) *ModifyFirewallTemplateRequest {
	s.Name = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetRegionId(v string) *ModifyFirewallTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) Validate() error {
	if s.FirewallTemplateRule != nil {
		for _, item := range s.FirewallTemplateRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyFirewallTemplateRequestFirewallTemplateRule struct {
	// The ID of the firewall template rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	FirewallTemplateRuleId *string `json:"FirewallTemplateRuleId,omitempty" xml:"FirewallTemplateRuleId,omitempty"`
	// The port range. Valid values: 1 to 65535. Use a forward slash (/) to separate the start port and the end port. For example, `1024/1055` indicates a port range from 1024 to 1055.
	//
	// > If you set the protocol to ICMP, the port must be set to -1/-1.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remark of the firewall rule.
	//
	// example:
	//
	// test
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
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source IPv4 CIDR block for which you want to configure access permissions. You can specify an IP address range in CIDR format or IPv4 format.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s ModifyFirewallTemplateRequestFirewallTemplateRule) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallTemplateRequestFirewallTemplateRule) GoString() string {
	return s.String()
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) GetFirewallTemplateRuleId() *string {
	return s.FirewallTemplateRuleId
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) GetPort() *string {
	return s.Port
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) GetRemark() *string {
	return s.Remark
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) GetRuleProtocol() *string {
	return s.RuleProtocol
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetFirewallTemplateRuleId(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.FirewallTemplateRuleId = &v
	return s
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetPort(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.Port = &v
	return s
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetRemark(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.Remark = &v
	return s
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetRuleProtocol(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.RuleProtocol = &v
	return s
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetSourceCidrIp(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) Validate() error {
	return dara.Validate(s)
}

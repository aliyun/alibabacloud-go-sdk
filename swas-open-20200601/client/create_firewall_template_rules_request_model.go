// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallTemplateRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallRule(v []*CreateFirewallTemplateRulesRequestFirewallRule) *CreateFirewallTemplateRulesRequest
	GetFirewallRule() []*CreateFirewallTemplateRulesRequestFirewallRule
	SetFirewallTemplateId(v string) *CreateFirewallTemplateRulesRequest
	GetFirewallTemplateId() *string
	SetRegionId(v string) *CreateFirewallTemplateRulesRequest
	GetRegionId() *string
}

type CreateFirewallTemplateRulesRequest struct {
	// The firewall rules.
	//
	// This parameter is required.
	FirewallRule []*CreateFirewallTemplateRulesRequestFirewallRule `json:"FirewallRule,omitempty" xml:"FirewallRule,omitempty" type:"Repeated"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the available regions for Simple Application Server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateFirewallTemplateRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallTemplateRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesRequest) GetFirewallRule() []*CreateFirewallTemplateRulesRequestFirewallRule {
	return s.FirewallRule
}

func (s *CreateFirewallTemplateRulesRequest) GetFirewallTemplateId() *string {
	return s.FirewallTemplateId
}

func (s *CreateFirewallTemplateRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFirewallTemplateRulesRequest) SetFirewallRule(v []*CreateFirewallTemplateRulesRequestFirewallRule) *CreateFirewallTemplateRulesRequest {
	s.FirewallRule = v
	return s
}

func (s *CreateFirewallTemplateRulesRequest) SetFirewallTemplateId(v string) *CreateFirewallTemplateRulesRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequest) SetRegionId(v string) *CreateFirewallTemplateRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequest) Validate() error {
	if s.FirewallRule != nil {
		for _, item := range s.FirewallRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFirewallTemplateRulesRequestFirewallRule struct {
	// The port range. Valid values: 1 to 65535. Separate the start port and the end port with a forward slash (/). For example, `1024/1055` indicates the port range from 1024 to 1055.
	//
	// > If you set the protocol to ICMP, the port must be set to -1/-1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The description of the firewall rule.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// - TCP: TCP.
	//
	// - UDP: UDP.
	//
	// - TCP+UDP: TCP and UDP.
	//
	// - ICMP: ICMP.
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source IPv4 CIDR address range for which you want to configure access Settings. Address ranges in CIDR format and IPv4 format are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 222.70.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateFirewallTemplateRulesRequestFirewallRule) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallTemplateRulesRequestFirewallRule) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) GetPort() *string {
	return s.Port
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) GetRemark() *string {
	return s.Remark
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) GetRuleProtocol() *string {
	return s.RuleProtocol
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) SetPort(v string) *CreateFirewallTemplateRulesRequestFirewallRule {
	s.Port = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) SetRemark(v string) *CreateFirewallTemplateRulesRequestFirewallRule {
	s.Remark = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) SetRuleProtocol(v string) *CreateFirewallTemplateRulesRequestFirewallRule {
	s.RuleProtocol = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) SetSourceCidrIp(v string) *CreateFirewallTemplateRulesRequestFirewallRule {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallTemplateRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallTemplateRules(v []*CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) *CreateFirewallTemplateRulesResponseBody
	GetFirewallTemplateRules() []*CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules
	SetRequestId(v string) *CreateFirewallTemplateRulesResponseBody
	GetRequestId() *string
}

type CreateFirewallTemplateRulesResponseBody struct {
	// The firewall template rules.
	FirewallTemplateRules []*CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules `json:"FirewallTemplateRules,omitempty" xml:"FirewallTemplateRules,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallTemplateRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallTemplateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesResponseBody) GetFirewallTemplateRules() []*CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	return s.FirewallTemplateRules
}

func (s *CreateFirewallTemplateRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFirewallTemplateRulesResponseBody) SetFirewallTemplateRules(v []*CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) *CreateFirewallTemplateRulesResponseBody {
	s.FirewallTemplateRules = v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBody) SetRequestId(v string) *CreateFirewallTemplateRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBody) Validate() error {
	if s.FirewallTemplateRules != nil {
		for _, item := range s.FirewallTemplateRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules struct {
	// The ID of the firewall template rule.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateRuleId *string `json:"FirewallTemplateRuleId,omitempty" xml:"FirewallTemplateRuleId,omitempty"`
	// The port range. Valid values: 1 to 65535. Separate the start port and the end port with a forward slash (/). For example, `1024/1055` indicates the port range from 1024 to 1055.
	//
	// > If you set the protocol to ICMP, the port must be set to -1/-1.
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
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source IPv4 CIDR address range for which you want to configure access Settings. Address ranges in CIDR format and IPv4 format are supported.
	//
	// example:
	//
	// 171.233.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) GetFirewallTemplateRuleId() *string {
	return s.FirewallTemplateRuleId
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) GetPort() *string {
	return s.Port
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) GetRemark() *string {
	return s.Remark
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) GetRuleProtocol() *string {
	return s.RuleProtocol
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetFirewallTemplateRuleId(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.FirewallTemplateRuleId = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetPort(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.Port = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetRemark(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.Remark = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetRuleProtocol(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.RuleProtocol = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetSourceCidrIp(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) Validate() error {
	return dara.Validate(s)
}

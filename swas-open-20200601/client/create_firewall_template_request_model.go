// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateFirewallTemplateRequest
	GetDescription() *string
	SetFirewallRule(v []*CreateFirewallTemplateRequestFirewallRule) *CreateFirewallTemplateRequest
	GetFirewallRule() []*CreateFirewallTemplateRequestFirewallRule
	SetName(v string) *CreateFirewallTemplateRequest
	GetName() *string
	SetRegionId(v string) *CreateFirewallTemplateRequest
	GetRegionId() *string
}

type CreateFirewallTemplateRequest struct {
	// The description of the firewall template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The firewall rules.
	FirewallRule []*CreateFirewallTemplateRequestFirewallRule `json:"FirewallRule,omitempty" xml:"FirewallRule,omitempty" type:"Repeated"`
	// The name of the firewall template.
	//
	// This parameter is required.
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

func (s CreateFirewallTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFirewallTemplateRequest) GetFirewallRule() []*CreateFirewallTemplateRequestFirewallRule {
	return s.FirewallRule
}

func (s *CreateFirewallTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateFirewallTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFirewallTemplateRequest) SetDescription(v string) *CreateFirewallTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateFirewallTemplateRequest) SetFirewallRule(v []*CreateFirewallTemplateRequestFirewallRule) *CreateFirewallTemplateRequest {
	s.FirewallRule = v
	return s
}

func (s *CreateFirewallTemplateRequest) SetName(v string) *CreateFirewallTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateFirewallTemplateRequest) SetRegionId(v string) *CreateFirewallTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFirewallTemplateRequest) Validate() error {
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

type CreateFirewallTemplateRequestFirewallRule struct {
	// The port range. Valid values: 1 to 65535. Use a forward slash (/) to separate the start port and the end port. For example, `1024/1055` indicates a port range from 1024 to 1055.
	//
	// > If you set the protocol to ICMP, the port must be set to -1/-1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remark for the firewall rule.
	//
	// example:
	//
	// Custom rule
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
	// The source IPv4 CIDR block for which to set access permissions. Both CIDR and IPv4 formats are supported for the IP address range.
	//
	// This parameter is required.
	//
	// example:
	//
	// 223.166.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateFirewallTemplateRequestFirewallRule) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallTemplateRequestFirewallRule) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRequestFirewallRule) GetPort() *string {
	return s.Port
}

func (s *CreateFirewallTemplateRequestFirewallRule) GetRemark() *string {
	return s.Remark
}

func (s *CreateFirewallTemplateRequestFirewallRule) GetRuleProtocol() *string {
	return s.RuleProtocol
}

func (s *CreateFirewallTemplateRequestFirewallRule) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *CreateFirewallTemplateRequestFirewallRule) SetPort(v string) *CreateFirewallTemplateRequestFirewallRule {
	s.Port = &v
	return s
}

func (s *CreateFirewallTemplateRequestFirewallRule) SetRemark(v string) *CreateFirewallTemplateRequestFirewallRule {
	s.Remark = &v
	return s
}

func (s *CreateFirewallTemplateRequestFirewallRule) SetRuleProtocol(v string) *CreateFirewallTemplateRequestFirewallRule {
	s.RuleProtocol = &v
	return s
}

func (s *CreateFirewallTemplateRequestFirewallRule) SetSourceCidrIp(v string) *CreateFirewallTemplateRequestFirewallRule {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateFirewallTemplateRequestFirewallRule) Validate() error {
	return dara.Validate(s)
}

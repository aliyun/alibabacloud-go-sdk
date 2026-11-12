// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateFirewallRulesRequest
	GetClientToken() *string
	SetFirewallRules(v []*CreateFirewallRulesRequestFirewallRules) *CreateFirewallRulesRequest
	GetFirewallRules() []*CreateFirewallRulesRequestFirewallRules
	SetInstanceId(v string) *CreateFirewallRulesRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateFirewallRulesRequest
	GetRegionId() *string
	SetTag(v []*CreateFirewallRulesRequestTag) *CreateFirewallRulesRequest
	GetTag() []*CreateFirewallRulesRequestTag
}

type CreateFirewallRulesRequest struct {
	// The client token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The array of firewall rules.
	FirewallRules []*CreateFirewallRulesRequestFirewallRules `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty" type:"Repeated"`
	// The instance ID of the specified simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the regions where simple application servers are available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of tags. The list can contain up to 20 tags.
	Tag []*CreateFirewallRulesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateFirewallRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFirewallRulesRequest) GetFirewallRules() []*CreateFirewallRulesRequestFirewallRules {
	return s.FirewallRules
}

func (s *CreateFirewallRulesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateFirewallRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFirewallRulesRequest) GetTag() []*CreateFirewallRulesRequestTag {
	return s.Tag
}

func (s *CreateFirewallRulesRequest) SetClientToken(v string) *CreateFirewallRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFirewallRulesRequest) SetFirewallRules(v []*CreateFirewallRulesRequestFirewallRules) *CreateFirewallRulesRequest {
	s.FirewallRules = v
	return s
}

func (s *CreateFirewallRulesRequest) SetInstanceId(v string) *CreateFirewallRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFirewallRulesRequest) SetRegionId(v string) *CreateFirewallRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFirewallRulesRequest) SetTag(v []*CreateFirewallRulesRequestTag) *CreateFirewallRulesRequest {
	s.Tag = v
	return s
}

func (s *CreateFirewallRulesRequest) Validate() error {
	if s.FirewallRules != nil {
		for _, item := range s.FirewallRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFirewallRulesRequestFirewallRules struct {
	// The port number.
	//
	// - TCP or UDP: Valid values are 1 to 65535. Use a forward slash (/) to separate the start port and the end port. Example: 1/200.
	//
	// - ICMP: -1/-1.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remark.
	//
	// example:
	//
	// Custom firewall rule
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
	// The source IP address or CIDR block.
	//
	// example:
	//
	// 47.101.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateFirewallRulesRequestFirewallRules) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRulesRequestFirewallRules) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesRequestFirewallRules) GetPort() *string {
	return s.Port
}

func (s *CreateFirewallRulesRequestFirewallRules) GetRemark() *string {
	return s.Remark
}

func (s *CreateFirewallRulesRequestFirewallRules) GetRuleProtocol() *string {
	return s.RuleProtocol
}

func (s *CreateFirewallRulesRequestFirewallRules) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *CreateFirewallRulesRequestFirewallRules) SetPort(v string) *CreateFirewallRulesRequestFirewallRules {
	s.Port = &v
	return s
}

func (s *CreateFirewallRulesRequestFirewallRules) SetRemark(v string) *CreateFirewallRulesRequestFirewallRules {
	s.Remark = &v
	return s
}

func (s *CreateFirewallRulesRequestFirewallRules) SetRuleProtocol(v string) *CreateFirewallRulesRequestFirewallRules {
	s.RuleProtocol = &v
	return s
}

func (s *CreateFirewallRulesRequestFirewallRules) SetSourceCidrIp(v string) *CreateFirewallRulesRequestFirewallRules {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateFirewallRulesRequestFirewallRules) Validate() error {
	return dara.Validate(s)
}

type CreateFirewallRulesRequestTag struct {
	// The tag key that is specified for the firewall rule. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. It can be up to 64 characters in length, cannot start with aliyun or acs:, and cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that is specified for the firewall rule. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. It can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFirewallRulesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallRulesRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateFirewallRulesRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateFirewallRulesRequestTag) SetKey(v string) *CreateFirewallRulesRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFirewallRulesRequestTag) SetValue(v string) *CreateFirewallRulesRequestTag {
	s.Value = &v
	return s
}

func (s *CreateFirewallRulesRequestTag) Validate() error {
	return dara.Validate(s)
}

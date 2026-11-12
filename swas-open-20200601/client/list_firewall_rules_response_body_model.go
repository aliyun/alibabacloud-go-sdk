// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFirewallRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallRules(v []*ListFirewallRulesResponseBodyFirewallRules) *ListFirewallRulesResponseBody
	GetFirewallRules() []*ListFirewallRulesResponseBodyFirewallRules
	SetPageNumber(v int32) *ListFirewallRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFirewallRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListFirewallRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListFirewallRulesResponseBody
	GetTotalCount() *int32
}

type ListFirewallRulesResponseBody struct {
	// The list of firewall rules.
	FirewallRules []*ListFirewallRulesResponseBodyFirewallRules `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty" type:"Repeated"`
	// The page number of the firewall rule list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The paging Settings for the number of entries per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of firewall rules.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFirewallRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponseBody) GetFirewallRules() []*ListFirewallRulesResponseBodyFirewallRules {
	return s.FirewallRules
}

func (s *ListFirewallRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFirewallRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFirewallRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFirewallRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFirewallRulesResponseBody) SetFirewallRules(v []*ListFirewallRulesResponseBodyFirewallRules) *ListFirewallRulesResponseBody {
	s.FirewallRules = v
	return s
}

func (s *ListFirewallRulesResponseBody) SetPageNumber(v int32) *ListFirewallRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFirewallRulesResponseBody) SetPageSize(v int32) *ListFirewallRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFirewallRulesResponseBody) SetRequestId(v string) *ListFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFirewallRulesResponseBody) SetTotalCount(v int32) *ListFirewallRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFirewallRulesResponseBody) Validate() error {
	if s.FirewallRules != nil {
		for _, item := range s.FirewallRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFirewallRulesResponseBodyFirewallRules struct {
	// The ID of the simple application server to which the firewall rule belongs.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The firewall policy. Valid values:
	//
	// - accept: Allow.
	//
	// - drop: Deny.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remark of the firewall rule.
	//
	// example:
	//
	// test-default port for MySQL server
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the firewall rule.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// - TCP: TCP.
	//
	// - UDP: UDP.
	//
	// - TCP+UDP: TCP and UDP.
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The tags of the firewall rules.
	Tags []*ListFirewallRulesResponseBodyFirewallRulesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListFirewallRulesResponseBodyFirewallRules) String() string {
	return dara.Prettify(s)
}

func (s ListFirewallRulesResponseBodyFirewallRules) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponseBodyFirewallRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFirewallRulesResponseBodyFirewallRules) GetPolicy() *string {
	return s.Policy
}

func (s *ListFirewallRulesResponseBodyFirewallRules) GetPort() *string {
	return s.Port
}

func (s *ListFirewallRulesResponseBodyFirewallRules) GetRemark() *string {
	return s.Remark
}

func (s *ListFirewallRulesResponseBodyFirewallRules) GetRuleId() *string {
	return s.RuleId
}

func (s *ListFirewallRulesResponseBodyFirewallRules) GetRuleProtocol() *string {
	return s.RuleProtocol
}

func (s *ListFirewallRulesResponseBodyFirewallRules) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ListFirewallRulesResponseBodyFirewallRules) GetTags() []*ListFirewallRulesResponseBodyFirewallRulesTags {
	return s.Tags
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetInstanceId(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.InstanceId = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetPolicy(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.Policy = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetPort(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.Port = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetRemark(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.Remark = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetRuleId(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.RuleId = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetRuleProtocol(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.RuleProtocol = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetSourceCidrIp(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.SourceCidrIp = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetTags(v []*ListFirewallRulesResponseBodyFirewallRulesTags) *ListFirewallRulesResponseBodyFirewallRules {
	s.Tags = v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFirewallRulesResponseBodyFirewallRulesTags struct {
	// The tag key of the firewall rule.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the firewall rule.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFirewallRulesResponseBodyFirewallRulesTags) String() string {
	return dara.Prettify(s)
}

func (s ListFirewallRulesResponseBodyFirewallRulesTags) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponseBodyFirewallRulesTags) GetKey() *string {
	return s.Key
}

func (s *ListFirewallRulesResponseBodyFirewallRulesTags) GetValue() *string {
	return s.Value
}

func (s *ListFirewallRulesResponseBodyFirewallRulesTags) SetKey(v string) *ListFirewallRulesResponseBodyFirewallRulesTags {
	s.Key = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRulesTags) SetValue(v string) *ListFirewallRulesResponseBodyFirewallRulesTags {
	s.Value = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRulesTags) Validate() error {
	return dara.Validate(s)
}

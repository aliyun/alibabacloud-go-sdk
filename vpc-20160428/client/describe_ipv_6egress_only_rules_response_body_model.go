// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6EgressOnlyRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6EgressOnlyRules(v *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules) *DescribeIpv6EgressOnlyRulesResponseBody
	GetIpv6EgressOnlyRules() *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules
	SetPageNumber(v int32) *DescribeIpv6EgressOnlyRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpv6EgressOnlyRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIpv6EgressOnlyRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIpv6EgressOnlyRulesResponseBody
	GetTotalCount() *int32
}

type DescribeIpv6EgressOnlyRulesResponseBody struct {
	// The details about the egress-only rules.
	Ipv6EgressOnlyRules *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules `json:"Ipv6EgressOnlyRules,omitempty" xml:"Ipv6EgressOnlyRules,omitempty" type:"Struct"`
	// The number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E16671B7-DEA6-48E0-8E9C-41913DAD44DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIpv6EgressOnlyRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6EgressOnlyRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) GetIpv6EgressOnlyRules() *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules {
	return s.Ipv6EgressOnlyRules
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) SetIpv6EgressOnlyRules(v *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules) *DescribeIpv6EgressOnlyRulesResponseBody {
	s.Ipv6EgressOnlyRules = v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) SetPageNumber(v int32) *DescribeIpv6EgressOnlyRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) SetPageSize(v int32) *DescribeIpv6EgressOnlyRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) SetRequestId(v string) *DescribeIpv6EgressOnlyRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) SetTotalCount(v int32) *DescribeIpv6EgressOnlyRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBody) Validate() error {
	if s.Ipv6EgressOnlyRules != nil {
		if err := s.Ipv6EgressOnlyRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules struct {
	Ipv6EgressOnlyRule []*DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule `json:"Ipv6EgressOnlyRule,omitempty" xml:"Ipv6EgressOnlyRule,omitempty" type:"Repeated"`
}

func (s DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules) GoString() string {
	return s.String()
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules) GetIpv6EgressOnlyRule() []*DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule {
	return s.Ipv6EgressOnlyRule
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules) SetIpv6EgressOnlyRule(v []*DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules {
	s.Ipv6EgressOnlyRule = v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRules) Validate() error {
	if s.Ipv6EgressOnlyRule != nil {
		for _, item := range s.Ipv6EgressOnlyRule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule struct {
	// The description of the egress-only rule.
	//
	// example:
	//
	// ruledescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance to which the egress-only rule applies.
	//
	// example:
	//
	// ipv6gw-bp1rhhs9zjlxukc5e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance to which the egress-only rule applies.
	//
	// example:
	//
	// Ipv6Address
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the egress-only rule.
	//
	// example:
	//
	// ipv6py-bp1rr7fq1md8pbb3k****
	Ipv6EgressOnlyRuleId *string `json:"Ipv6EgressOnlyRuleId,omitempty" xml:"Ipv6EgressOnlyRuleId,omitempty"`
	// The name of the egress-only rule.
	//
	// example:
	//
	// rulename
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the egress-only rule.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) GoString() string {
	return s.String()
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) GetIpv6EgressOnlyRuleId() *string {
	return s.Ipv6EgressOnlyRuleId
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) GetName() *string {
	return s.Name
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) GetStatus() *string {
	return s.Status
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) SetDescription(v string) *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule {
	s.Description = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) SetInstanceId(v string) *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) SetInstanceType(v string) *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule {
	s.InstanceType = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) SetIpv6EgressOnlyRuleId(v string) *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule {
	s.Ipv6EgressOnlyRuleId = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) SetName(v string) *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule {
	s.Name = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) SetStatus(v string) *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule {
	s.Status = &v
	return s
}

func (s *DescribeIpv6EgressOnlyRulesResponseBodyIpv6EgressOnlyRulesIpv6EgressOnlyRule) Validate() error {
	return dara.Validate(s)
}

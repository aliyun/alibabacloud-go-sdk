// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRuleGroupsResponseBody
	GetRequestId() *string
	SetRuleGroups(v []*DescribeRuleGroupsResponseBodyRuleGroups) *DescribeRuleGroupsResponseBody
	GetRuleGroups() []*DescribeRuleGroupsResponseBodyRuleGroups
	SetTotalCount(v int64) *DescribeRuleGroupsResponseBody
	GetTotalCount() *int64
}

type DescribeRuleGroupsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 02E9A4B8-90FB-5F41-A049-C82277EB82FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array of regular expression rule groups.
	RuleGroups []*DescribeRuleGroupsResponseBodyRuleGroups `json:"RuleGroups,omitempty" xml:"RuleGroups,omitempty" type:"Repeated"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 24
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRuleGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRuleGroupsResponseBody) GetRuleGroups() []*DescribeRuleGroupsResponseBodyRuleGroups {
	return s.RuleGroups
}

func (s *DescribeRuleGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRuleGroupsResponseBody) SetRequestId(v string) *DescribeRuleGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBody) SetRuleGroups(v []*DescribeRuleGroupsResponseBodyRuleGroups) *DescribeRuleGroupsResponseBody {
	s.RuleGroups = v
	return s
}

func (s *DescribeRuleGroupsResponseBody) SetTotalCount(v int64) *DescribeRuleGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRuleGroupsResponseBody) Validate() error {
	if s.RuleGroups != nil {
		for _, item := range s.RuleGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRuleGroupsResponseBodyRuleGroups struct {
	// The most recent time when the rule group was modified.
	//
	// example:
	//
	// 1664336364000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the automatic update feature is enabled for the rule group.
	//
	// 	- 1: The automatic update feature is enabled for the rule group.
	//
	// 	- 2: The automatic update feature is disabled for the rule group.
	//
	// example:
	//
	// 1
	IsSubscribe *int32 `json:"IsSubscribe,omitempty" xml:"IsSubscribe,omitempty"`
	// The ID of the rule group.
	//
	// 	- 0: The rule group is created from scratch.
	//
	// 	- 1011: The rule group is a strict rule group.
	//
	// 	- 1012: The rule group is a medium rule group.
	//
	// 	- 1013: The rue group is a loose rule group.
	//
	// example:
	//
	// 1012
	ParentRuleGroupId *int64 `json:"ParentRuleGroupId,omitempty" xml:"ParentRuleGroupId,omitempty"`
	// The ID of the regular expression rule group.
	//
	// example:
	//
	// 115361
	RuleGroupId *int64 `json:"RuleGroupId,omitempty" xml:"RuleGroupId,omitempty"`
	// The name of the rule group.
	//
	// example:
	//
	// ssssss
	RuleGroupName *string `json:"RuleGroupName,omitempty" xml:"RuleGroupName,omitempty"`
	// The number of built-in rules in the rule group.
	//
	// example:
	//
	// 4444
	RuleTotalCount *int32 `json:"RuleTotalCount,omitempty" xml:"RuleTotalCount,omitempty"`
}

func (s DescribeRuleGroupsResponseBodyRuleGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleGroupsResponseBodyRuleGroups) GoString() string {
	return s.String()
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) GetIsSubscribe() *int32 {
	return s.IsSubscribe
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) GetParentRuleGroupId() *int64 {
	return s.ParentRuleGroupId
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) GetRuleGroupId() *int64 {
	return s.RuleGroupId
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) GetRuleGroupName() *string {
	return s.RuleGroupName
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) GetRuleTotalCount() *int32 {
	return s.RuleTotalCount
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetGmtModified(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.GmtModified = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetIsSubscribe(v int32) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.IsSubscribe = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetParentRuleGroupId(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.ParentRuleGroupId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleGroupId(v int64) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleGroupId = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleGroupName(v string) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleGroupName = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) SetRuleTotalCount(v int32) *DescribeRuleGroupsResponseBodyRuleGroups {
	s.RuleTotalCount = &v
	return s
}

func (s *DescribeRuleGroupsResponseBodyRuleGroups) Validate() error {
	return dara.Validate(s)
}

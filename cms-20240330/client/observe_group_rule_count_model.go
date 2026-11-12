// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObserveGroupRuleCount interface {
	dara.Model
	String() string
	GoString() string
	SetMainCount(v int64) *ObserveGroupRuleCount
	GetMainCount() *int64
	SetRelatedCount(v int64) *ObserveGroupRuleCount
	GetRelatedCount() *int64
	SetRuleId(v string) *ObserveGroupRuleCount
	GetRuleId() *string
	SetTotalCount(v int64) *ObserveGroupRuleCount
	GetTotalCount() *int64
}

type ObserveGroupRuleCount struct {
	// The number of deduplicated primary entities.
	//
	// example:
	//
	// 3
	MainCount *int64 `json:"mainCount,omitempty" xml:"mainCount,omitempty"`
	// The number of deduplicated associated child entities.
	//
	// example:
	//
	// 9
	RelatedCount *int64 `json:"relatedCount,omitempty" xml:"relatedCount,omitempty"`
	// The stable ID of the managed rule.
	//
	// example:
	//
	// rule-rocketmq-prod
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The total number of deduplicated entities.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ObserveGroupRuleCount) String() string {
	return dara.Prettify(s)
}

func (s ObserveGroupRuleCount) GoString() string {
	return s.String()
}

func (s *ObserveGroupRuleCount) GetMainCount() *int64 {
	return s.MainCount
}

func (s *ObserveGroupRuleCount) GetRelatedCount() *int64 {
	return s.RelatedCount
}

func (s *ObserveGroupRuleCount) GetRuleId() *string {
	return s.RuleId
}

func (s *ObserveGroupRuleCount) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ObserveGroupRuleCount) SetMainCount(v int64) *ObserveGroupRuleCount {
	s.MainCount = &v
	return s
}

func (s *ObserveGroupRuleCount) SetRelatedCount(v int64) *ObserveGroupRuleCount {
	s.RelatedCount = &v
	return s
}

func (s *ObserveGroupRuleCount) SetRuleId(v string) *ObserveGroupRuleCount {
	s.RuleId = &v
	return s
}

func (s *ObserveGroupRuleCount) SetTotalCount(v int64) *ObserveGroupRuleCount {
	s.TotalCount = &v
	return s
}

func (s *ObserveGroupRuleCount) Validate() error {
	return dara.Validate(s)
}

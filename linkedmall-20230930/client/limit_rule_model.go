// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLimitRule interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *LimitRule
	GetBeginTime() *int64
	SetCondcase(v string) *LimitRule
	GetCondcase() *string
	SetEndTime(v int64) *LimitRule
	GetEndTime() *int64
	SetLimitNum(v int32) *LimitRule
	GetLimitNum() *int32
	SetRuleType(v string) *LimitRule
	GetRuleType() *string
}

type LimitRule struct {
	// example:
	//
	// 1724947200000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// example:
	//
	// day
	Condcase *string `json:"condcase,omitempty" xml:"condcase,omitempty"`
	// example:
	//
	// 1724947200000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	LimitNum *int32 `json:"limitNum,omitempty" xml:"limitNum,omitempty"`
	// example:
	//
	// UpperNumberPerUser
	RuleType *string `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
}

func (s LimitRule) String() string {
	return dara.Prettify(s)
}

func (s LimitRule) GoString() string {
	return s.String()
}

func (s *LimitRule) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *LimitRule) GetCondcase() *string {
	return s.Condcase
}

func (s *LimitRule) GetEndTime() *int64 {
	return s.EndTime
}

func (s *LimitRule) GetLimitNum() *int32 {
	return s.LimitNum
}

func (s *LimitRule) GetRuleType() *string {
	return s.RuleType
}

func (s *LimitRule) SetBeginTime(v int64) *LimitRule {
	s.BeginTime = &v
	return s
}

func (s *LimitRule) SetCondcase(v string) *LimitRule {
	s.Condcase = &v
	return s
}

func (s *LimitRule) SetEndTime(v int64) *LimitRule {
	s.EndTime = &v
	return s
}

func (s *LimitRule) SetLimitNum(v int32) *LimitRule {
	s.LimitNum = &v
	return s
}

func (s *LimitRule) SetRuleType(v string) *LimitRule {
	s.RuleType = &v
	return s
}

func (s *LimitRule) Validate() error {
	return dara.Validate(s)
}

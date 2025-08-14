// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHitRuleTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeHitRuleTrendRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeHitRuleTrendRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeHitRuleTrendRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeHitRuleTrendRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeHitRuleTrendRequest
	GetRegId() *string
	SetRuleStatus(v string) *DescribeHitRuleTrendRequest
	GetRuleStatus() *string
}

type DescribeHitRuleTrendRequest struct {
	// Sets the language type for the request and response messages. Default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Start time, accurate to milliseconds (ms).
	//
	// example:
	//
	// 1737101348000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// example:
	//
	// 1746669075000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event codes, separated by commas (,).
	//
	// example:
	//
	// de_ahqhsw7665,de_agbzfi5134
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Rule status
	//
	// example:
	//
	// DRAFT
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
}

func (s DescribeHitRuleTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeHitRuleTrendRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeHitRuleTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeHitRuleTrendRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeHitRuleTrendRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeHitRuleTrendRequest) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeHitRuleTrendRequest) SetLang(v string) *DescribeHitRuleTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeHitRuleTrendRequest) SetBeginTime(v int64) *DescribeHitRuleTrendRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeHitRuleTrendRequest) SetEndTime(v int64) *DescribeHitRuleTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHitRuleTrendRequest) SetEventCodes(v string) *DescribeHitRuleTrendRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeHitRuleTrendRequest) SetRegId(v string) *DescribeHitRuleTrendRequest {
	s.RegId = &v
	return s
}

func (s *DescribeHitRuleTrendRequest) SetRuleStatus(v string) *DescribeHitRuleTrendRequest {
	s.RuleStatus = &v
	return s
}

func (s *DescribeHitRuleTrendRequest) Validate() error {
	return dara.Validate(s)
}

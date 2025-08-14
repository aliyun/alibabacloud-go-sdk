// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleBarChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRuleBarChartRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeRuleBarChartRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeRuleBarChartRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeRuleBarChartRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeRuleBarChartRequest
	GetRegId() *string
}

type DescribeRuleBarChartRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Query start time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1739841750000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1750904603000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event codes, separated by commas (,).
	//
	// example:
	//
	// account_abuse_pro,account_abuse
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeRuleBarChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleBarChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeRuleBarChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRuleBarChartRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeRuleBarChartRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeRuleBarChartRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeRuleBarChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRuleBarChartRequest) SetLang(v string) *DescribeRuleBarChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRuleBarChartRequest) SetBeginTime(v int64) *DescribeRuleBarChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeRuleBarChartRequest) SetEndTime(v int64) *DescribeRuleBarChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRuleBarChartRequest) SetEventCodes(v string) *DescribeRuleBarChartRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeRuleBarChartRequest) SetRegId(v string) *DescribeRuleBarChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRuleBarChartRequest) Validate() error {
	return dara.Validate(s)
}

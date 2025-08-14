// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHitRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeHitRuleListRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeHitRuleListRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeHitRuleListRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeHitRuleListRequest
	GetEventCodes() *string
	SetEventType(v string) *DescribeHitRuleListRequest
	GetEventType() *string
	SetRegId(v string) *DescribeHitRuleListRequest
	GetRegId() *string
}

type DescribeHitRuleListRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
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
	// 1752027960000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event codes, separated by commas if multiple.
	//
	// example:
	//
	// de_ahqhsw7665,de_ahqhsw7622
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Event type
	//
	// example:
	//
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeHitRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeHitRuleListRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeHitRuleListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeHitRuleListRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeHitRuleListRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeHitRuleListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeHitRuleListRequest) SetLang(v string) *DescribeHitRuleListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeHitRuleListRequest) SetBeginTime(v int64) *DescribeHitRuleListRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeHitRuleListRequest) SetEndTime(v int64) *DescribeHitRuleListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHitRuleListRequest) SetEventCodes(v string) *DescribeHitRuleListRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeHitRuleListRequest) SetEventType(v string) *DescribeHitRuleListRequest {
	s.EventType = &v
	return s
}

func (s *DescribeHitRuleListRequest) SetRegId(v string) *DescribeHitRuleListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeHitRuleListRequest) Validate() error {
	return dara.Validate(s)
}

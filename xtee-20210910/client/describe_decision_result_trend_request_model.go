// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDecisionResultTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDecisionResultTrendRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeDecisionResultTrendRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeDecisionResultTrendRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeDecisionResultTrendRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeDecisionResultTrendRequest
	GetRegId() *string
}

type DescribeDecisionResultTrendRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start timestamp, in milliseconds.
	//
	// example:
	//
	// 1737101348000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The end timestamp, in milliseconds.
	//
	// example:
	//
	// 1747101348000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The event code.
	//
	// example:
	//
	// de_ahqhsw7665,de_agbzfi5134
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeDecisionResultTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDecisionResultTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeDecisionResultTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDecisionResultTrendRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeDecisionResultTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDecisionResultTrendRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeDecisionResultTrendRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeDecisionResultTrendRequest) SetLang(v string) *DescribeDecisionResultTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDecisionResultTrendRequest) SetBeginTime(v int64) *DescribeDecisionResultTrendRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeDecisionResultTrendRequest) SetEndTime(v int64) *DescribeDecisionResultTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDecisionResultTrendRequest) SetEventCodes(v string) *DescribeDecisionResultTrendRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeDecisionResultTrendRequest) SetRegId(v string) *DescribeDecisionResultTrendRequest {
	s.RegId = &v
	return s
}

func (s *DescribeDecisionResultTrendRequest) Validate() error {
	return dara.Validate(s)
}

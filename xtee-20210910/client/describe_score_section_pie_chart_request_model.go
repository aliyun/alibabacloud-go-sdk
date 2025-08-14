// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreSectionPieChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeScoreSectionPieChartRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeScoreSectionPieChartRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeScoreSectionPieChartRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeScoreSectionPieChartRequest
	GetEventCodes() *string
	SetEventType(v string) *DescribeScoreSectionPieChartRequest
	GetEventType() *string
	SetRegId(v string) *DescribeScoreSectionPieChartRequest
	GetRegId() *string
}

type DescribeScoreSectionPieChartRequest struct {
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
	// Start time, accurate to milliseconds (ms).
	//
	// example:
	//
	// 1751249559000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// example:
	//
	// 1740535600000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_afghcf6411
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Event type.
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

func (s DescribeScoreSectionPieChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionPieChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionPieChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeScoreSectionPieChartRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeScoreSectionPieChartRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeScoreSectionPieChartRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeScoreSectionPieChartRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeScoreSectionPieChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeScoreSectionPieChartRequest) SetLang(v string) *DescribeScoreSectionPieChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeScoreSectionPieChartRequest) SetBeginTime(v int64) *DescribeScoreSectionPieChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeScoreSectionPieChartRequest) SetEndTime(v int64) *DescribeScoreSectionPieChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScoreSectionPieChartRequest) SetEventCodes(v string) *DescribeScoreSectionPieChartRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeScoreSectionPieChartRequest) SetEventType(v string) *DescribeScoreSectionPieChartRequest {
	s.EventType = &v
	return s
}

func (s *DescribeScoreSectionPieChartRequest) SetRegId(v string) *DescribeScoreSectionPieChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeScoreSectionPieChartRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsBarChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTagsBarChartRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeTagsBarChartRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeTagsBarChartRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeTagsBarChartRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeTagsBarChartRequest
	GetRegId() *string
	SetResult(v string) *DescribeTagsBarChartRequest
	GetResult() *string
}

type DescribeTagsBarChartRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 1751249559000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1751595912000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_afghcf6411
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Policy execution result
	//
	// example:
	//
	// PASS
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeTagsBarChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsBarChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsBarChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTagsBarChartRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeTagsBarChartRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeTagsBarChartRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeTagsBarChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTagsBarChartRequest) GetResult() *string {
	return s.Result
}

func (s *DescribeTagsBarChartRequest) SetLang(v string) *DescribeTagsBarChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagsBarChartRequest) SetBeginTime(v int64) *DescribeTagsBarChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeTagsBarChartRequest) SetEndTime(v int64) *DescribeTagsBarChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTagsBarChartRequest) SetEventCodes(v string) *DescribeTagsBarChartRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeTagsBarChartRequest) SetRegId(v string) *DescribeTagsBarChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTagsBarChartRequest) SetResult(v string) *DescribeTagsBarChartRequest {
	s.Result = &v
	return s
}

func (s *DescribeTagsBarChartRequest) Validate() error {
	return dara.Validate(s)
}

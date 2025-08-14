// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventResultBarChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventResultBarChartRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeEventResultBarChartRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeEventResultBarChartRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeEventResultBarChartRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeEventResultBarChartRequest
	GetRegId() *string
}

type DescribeEventResultBarChartRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 1737101348000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1744337383000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_ahqhsw7665,de_ahqhsw7622
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeEventResultBarChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultBarChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventResultBarChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventResultBarChartRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeEventResultBarChartRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeEventResultBarChartRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeEventResultBarChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventResultBarChartRequest) SetLang(v string) *DescribeEventResultBarChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventResultBarChartRequest) SetBeginTime(v int64) *DescribeEventResultBarChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeEventResultBarChartRequest) SetEndTime(v int64) *DescribeEventResultBarChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventResultBarChartRequest) SetEventCodes(v string) *DescribeEventResultBarChartRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeEventResultBarChartRequest) SetRegId(v string) *DescribeEventResultBarChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventResultBarChartRequest) Validate() error {
	return dara.Validate(s)
}

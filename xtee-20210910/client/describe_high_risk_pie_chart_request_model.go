// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighRiskPieChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeHighRiskPieChartRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeHighRiskPieChartRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeHighRiskPieChartRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeHighRiskPieChartRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeHighRiskPieChartRequest
	GetRegId() *string
}

type DescribeHighRiskPieChartRequest struct {
	// Sets the language type for requests and received messages. Default value is **zh**. Values:
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
	// 1730453400000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event code.
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
}

func (s DescribeHighRiskPieChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighRiskPieChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeHighRiskPieChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeHighRiskPieChartRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeHighRiskPieChartRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeHighRiskPieChartRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeHighRiskPieChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeHighRiskPieChartRequest) SetLang(v string) *DescribeHighRiskPieChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeHighRiskPieChartRequest) SetBeginTime(v int64) *DescribeHighRiskPieChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeHighRiskPieChartRequest) SetEndTime(v int64) *DescribeHighRiskPieChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHighRiskPieChartRequest) SetEventCodes(v string) *DescribeHighRiskPieChartRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeHighRiskPieChartRequest) SetRegId(v string) *DescribeHighRiskPieChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeHighRiskPieChartRequest) Validate() error {
	return dara.Validate(s)
}

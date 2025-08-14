// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskLineChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeRiskLineChartRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeRiskLineChartRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeRiskLineChartRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeRiskLineChartRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeRiskLineChartRequest
	GetRegId() *string
}

type DescribeRiskLineChartRequest struct {
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
	// 1748491200000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event codes, separated by commas (,).
	//
	// example:
	//
	// de_afghcf6411,de_awkhwh0314
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeRiskLineChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskLineChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskLineChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskLineChartRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeRiskLineChartRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeRiskLineChartRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeRiskLineChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRiskLineChartRequest) SetLang(v string) *DescribeRiskLineChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskLineChartRequest) SetBeginTime(v int64) *DescribeRiskLineChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeRiskLineChartRequest) SetEndTime(v int64) *DescribeRiskLineChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskLineChartRequest) SetEventCodes(v string) *DescribeRiskLineChartRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeRiskLineChartRequest) SetRegId(v string) *DescribeRiskLineChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRiskLineChartRequest) Validate() error {
	return dara.Validate(s)
}

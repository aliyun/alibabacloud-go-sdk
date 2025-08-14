// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskTagsLineChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *DescribeRiskTagsLineChartRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeRiskTagsLineChartRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeRiskTagsLineChartRequest
	GetEventCodes() *string
	SetLang(v string) *DescribeRiskTagsLineChartRequest
	GetLang() *string
	SetRegId(v string) *DescribeRiskTagsLineChartRequest
	GetRegId() *string
}

type DescribeRiskTagsLineChartRequest struct {
	// Start time of the query, in milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729563800605
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1729563800605
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Event code
	//
	// example:
	//
	// de_afghcf6411
	EventCodes *string `json:"EventCodes,omitempty" xml:"EventCodes,omitempty"`
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
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s DescribeRiskTagsLineChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTagsLineChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskTagsLineChartRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeRiskTagsLineChartRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeRiskTagsLineChartRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeRiskTagsLineChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskTagsLineChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeRiskTagsLineChartRequest) SetBeginTime(v int64) *DescribeRiskTagsLineChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeRiskTagsLineChartRequest) SetEndTime(v int64) *DescribeRiskTagsLineChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskTagsLineChartRequest) SetEventCodes(v string) *DescribeRiskTagsLineChartRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeRiskTagsLineChartRequest) SetLang(v string) *DescribeRiskTagsLineChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskTagsLineChartRequest) SetRegId(v string) *DescribeRiskTagsLineChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeRiskTagsLineChartRequest) Validate() error {
	return dara.Validate(s)
}

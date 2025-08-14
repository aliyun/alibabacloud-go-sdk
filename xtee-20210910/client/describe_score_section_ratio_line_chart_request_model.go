// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreSectionRatioLineChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeScoreSectionRatioLineChartRequest
	GetLang() *string
	SetBeginTime(v string) *DescribeScoreSectionRatioLineChartRequest
	GetBeginTime() *string
	SetByPassEventCodes(v string) *DescribeScoreSectionRatioLineChartRequest
	GetByPassEventCodes() *string
	SetEndTime(v string) *DescribeScoreSectionRatioLineChartRequest
	GetEndTime() *string
	SetMainEventCodes(v string) *DescribeScoreSectionRatioLineChartRequest
	GetMainEventCodes() *string
	SetRegId(v string) *DescribeScoreSectionRatioLineChartRequest
	GetRegId() *string
	SetShuntEventCodes(v string) *DescribeScoreSectionRatioLineChartRequest
	GetShuntEventCodes() *string
}

type DescribeScoreSectionRatioLineChartRequest struct {
	// Sets the language type for request and response messages, with a default value of **zh**. Values:
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
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// Bypass event code
	//
	// example:
	//
	// de_amnhke2482
	ByPassEventCodes *string `json:"byPassEventCodes,omitempty" xml:"byPassEventCodes,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1748491200000
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Main event code
	//
	// example:
	//
	// de_avcqzc3714
	MainEventCodes *string `json:"mainEventCodes,omitempty" xml:"mainEventCodes,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Shunt event code
	//
	// example:
	//
	// de_amnhke2488
	ShuntEventCodes *string `json:"shuntEventCodes,omitempty" xml:"shuntEventCodes,omitempty"`
}

func (s DescribeScoreSectionRatioLineChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionRatioLineChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionRatioLineChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeScoreSectionRatioLineChartRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeScoreSectionRatioLineChartRequest) GetByPassEventCodes() *string {
	return s.ByPassEventCodes
}

func (s *DescribeScoreSectionRatioLineChartRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeScoreSectionRatioLineChartRequest) GetMainEventCodes() *string {
	return s.MainEventCodes
}

func (s *DescribeScoreSectionRatioLineChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeScoreSectionRatioLineChartRequest) GetShuntEventCodes() *string {
	return s.ShuntEventCodes
}

func (s *DescribeScoreSectionRatioLineChartRequest) SetLang(v string) *DescribeScoreSectionRatioLineChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartRequest) SetBeginTime(v string) *DescribeScoreSectionRatioLineChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartRequest) SetByPassEventCodes(v string) *DescribeScoreSectionRatioLineChartRequest {
	s.ByPassEventCodes = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartRequest) SetEndTime(v string) *DescribeScoreSectionRatioLineChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartRequest) SetMainEventCodes(v string) *DescribeScoreSectionRatioLineChartRequest {
	s.MainEventCodes = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartRequest) SetRegId(v string) *DescribeScoreSectionRatioLineChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartRequest) SetShuntEventCodes(v string) *DescribeScoreSectionRatioLineChartRequest {
	s.ShuntEventCodes = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartRequest) Validate() error {
	return dara.Validate(s)
}

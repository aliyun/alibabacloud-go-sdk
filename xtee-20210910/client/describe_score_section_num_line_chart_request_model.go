// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreSectionNumLineChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeScoreSectionNumLineChartRequest
	GetLang() *string
	SetBeginTime(v string) *DescribeScoreSectionNumLineChartRequest
	GetBeginTime() *string
	SetByPassEventCodes(v string) *DescribeScoreSectionNumLineChartRequest
	GetByPassEventCodes() *string
	SetEndTime(v string) *DescribeScoreSectionNumLineChartRequest
	GetEndTime() *string
	SetMainEventCodes(v string) *DescribeScoreSectionNumLineChartRequest
	GetMainEventCodes() *string
	SetRegId(v string) *DescribeScoreSectionNumLineChartRequest
	GetRegId() *string
	SetShuntEventCodes(v string) *DescribeScoreSectionNumLineChartRequest
	GetShuntEventCodes() *string
}

type DescribeScoreSectionNumLineChartRequest struct {
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
	// Start timestamp, in milliseconds.
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
	// End timestamp, in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1751249559000
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
	// Diversion event code
	//
	// example:
	//
	// de_amnhke2488
	ShuntEventCodes *string `json:"shuntEventCodes,omitempty" xml:"shuntEventCodes,omitempty"`
}

func (s DescribeScoreSectionNumLineChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionNumLineChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionNumLineChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeScoreSectionNumLineChartRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeScoreSectionNumLineChartRequest) GetByPassEventCodes() *string {
	return s.ByPassEventCodes
}

func (s *DescribeScoreSectionNumLineChartRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeScoreSectionNumLineChartRequest) GetMainEventCodes() *string {
	return s.MainEventCodes
}

func (s *DescribeScoreSectionNumLineChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeScoreSectionNumLineChartRequest) GetShuntEventCodes() *string {
	return s.ShuntEventCodes
}

func (s *DescribeScoreSectionNumLineChartRequest) SetLang(v string) *DescribeScoreSectionNumLineChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartRequest) SetBeginTime(v string) *DescribeScoreSectionNumLineChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartRequest) SetByPassEventCodes(v string) *DescribeScoreSectionNumLineChartRequest {
	s.ByPassEventCodes = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartRequest) SetEndTime(v string) *DescribeScoreSectionNumLineChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartRequest) SetMainEventCodes(v string) *DescribeScoreSectionNumLineChartRequest {
	s.MainEventCodes = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartRequest) SetRegId(v string) *DescribeScoreSectionNumLineChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartRequest) SetShuntEventCodes(v string) *DescribeScoreSectionNumLineChartRequest {
	s.ShuntEventCodes = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartRequest) Validate() error {
	return dara.Validate(s)
}

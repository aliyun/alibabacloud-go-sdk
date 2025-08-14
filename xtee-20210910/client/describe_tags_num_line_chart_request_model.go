// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsNumLineChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTagsNumLineChartRequest
	GetLang() *string
	SetBeginTime(v string) *DescribeTagsNumLineChartRequest
	GetBeginTime() *string
	SetByPassEventCodes(v string) *DescribeTagsNumLineChartRequest
	GetByPassEventCodes() *string
	SetEndTime(v string) *DescribeTagsNumLineChartRequest
	GetEndTime() *string
	SetMainEventCodes(v string) *DescribeTagsNumLineChartRequest
	GetMainEventCodes() *string
	SetRegId(v string) *DescribeTagsNumLineChartRequest
	GetRegId() *string
	SetShuntEventCodes(v string) *DescribeTagsNumLineChartRequest
	GetShuntEventCodes() *string
}

type DescribeTagsNumLineChartRequest struct {
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
	// 1749002991000
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Main event code
	//
	// example:
	//
	// de_amnhke2482
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

func (s DescribeTagsNumLineChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsNumLineChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsNumLineChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTagsNumLineChartRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeTagsNumLineChartRequest) GetByPassEventCodes() *string {
	return s.ByPassEventCodes
}

func (s *DescribeTagsNumLineChartRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTagsNumLineChartRequest) GetMainEventCodes() *string {
	return s.MainEventCodes
}

func (s *DescribeTagsNumLineChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTagsNumLineChartRequest) GetShuntEventCodes() *string {
	return s.ShuntEventCodes
}

func (s *DescribeTagsNumLineChartRequest) SetLang(v string) *DescribeTagsNumLineChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagsNumLineChartRequest) SetBeginTime(v string) *DescribeTagsNumLineChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeTagsNumLineChartRequest) SetByPassEventCodes(v string) *DescribeTagsNumLineChartRequest {
	s.ByPassEventCodes = &v
	return s
}

func (s *DescribeTagsNumLineChartRequest) SetEndTime(v string) *DescribeTagsNumLineChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTagsNumLineChartRequest) SetMainEventCodes(v string) *DescribeTagsNumLineChartRequest {
	s.MainEventCodes = &v
	return s
}

func (s *DescribeTagsNumLineChartRequest) SetRegId(v string) *DescribeTagsNumLineChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTagsNumLineChartRequest) SetShuntEventCodes(v string) *DescribeTagsNumLineChartRequest {
	s.ShuntEventCodes = &v
	return s
}

func (s *DescribeTagsNumLineChartRequest) Validate() error {
	return dara.Validate(s)
}

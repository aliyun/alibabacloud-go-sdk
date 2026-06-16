// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRatioLineChartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTagsRatioLineChartRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeTagsRatioLineChartRequest
	GetBeginTime() *int64
	SetByPassEventCodes(v string) *DescribeTagsRatioLineChartRequest
	GetByPassEventCodes() *string
	SetEndTime(v int64) *DescribeTagsRatioLineChartRequest
	GetEndTime() *int64
	SetMainEventCodes(v string) *DescribeTagsRatioLineChartRequest
	GetMainEventCodes() *string
	SetRegId(v string) *DescribeTagsRatioLineChartRequest
	GetRegId() *string
	SetShuntEventCodes(v string) *DescribeTagsRatioLineChartRequest
	GetShuntEventCodes() *string
}

type DescribeTagsRatioLineChartRequest struct {
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
	// The start time, in milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1751249559000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The bypass event codes.
	//
	// example:
	//
	// de_amnhke2482
	ByPassEventCodes *string `json:"byPassEventCodes,omitempty" xml:"byPassEventCodes,omitempty"`
	// The end time, in milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1751595912000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The main event codes.
	//
	// example:
	//
	// de_amnhke2482
	MainEventCodes *string `json:"mainEventCodes,omitempty" xml:"mainEventCodes,omitempty"`
	// The region code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The diversion event codes.
	//
	// example:
	//
	// de_amnhke2488
	ShuntEventCodes *string `json:"shuntEventCodes,omitempty" xml:"shuntEventCodes,omitempty"`
}

func (s DescribeTagsRatioLineChartRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRatioLineChartRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRatioLineChartRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTagsRatioLineChartRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeTagsRatioLineChartRequest) GetByPassEventCodes() *string {
	return s.ByPassEventCodes
}

func (s *DescribeTagsRatioLineChartRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeTagsRatioLineChartRequest) GetMainEventCodes() *string {
	return s.MainEventCodes
}

func (s *DescribeTagsRatioLineChartRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTagsRatioLineChartRequest) GetShuntEventCodes() *string {
	return s.ShuntEventCodes
}

func (s *DescribeTagsRatioLineChartRequest) SetLang(v string) *DescribeTagsRatioLineChartRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagsRatioLineChartRequest) SetBeginTime(v int64) *DescribeTagsRatioLineChartRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeTagsRatioLineChartRequest) SetByPassEventCodes(v string) *DescribeTagsRatioLineChartRequest {
	s.ByPassEventCodes = &v
	return s
}

func (s *DescribeTagsRatioLineChartRequest) SetEndTime(v int64) *DescribeTagsRatioLineChartRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTagsRatioLineChartRequest) SetMainEventCodes(v string) *DescribeTagsRatioLineChartRequest {
	s.MainEventCodes = &v
	return s
}

func (s *DescribeTagsRatioLineChartRequest) SetRegId(v string) *DescribeTagsRatioLineChartRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTagsRatioLineChartRequest) SetShuntEventCodes(v string) *DescribeTagsRatioLineChartRequest {
	s.ShuntEventCodes = &v
	return s
}

func (s *DescribeTagsRatioLineChartRequest) Validate() error {
	return dara.Validate(s)
}

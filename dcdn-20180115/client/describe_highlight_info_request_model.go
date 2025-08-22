// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighlightInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeHighlightInfoRequest
	GetEndTime() *string
	SetLang(v string) *DescribeHighlightInfoRequest
	GetLang() *string
	SetStartTime(v string) *DescribeHighlightInfoRequest
	GetStartTime() *string
	SetTraceId(v string) *DescribeHighlightInfoRequest
	GetTraceId() *string
}

type DescribeHighlightInfoRequest struct {
	// The end of the time range to query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-04-10T02:43:34Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The access language. Valid values:
	//
	// 	- **en-US*	- (default): English.
	//
	// 	- **zh-CN**: Chinese.
	//
	// This parameter is required.
	//
	// example:
	//
	// en_US
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-02-19T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the trace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 800e749616838513398137319e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeHighlightInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighlightInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeHighlightInfoRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeHighlightInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeHighlightInfoRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeHighlightInfoRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeHighlightInfoRequest) SetEndTime(v string) *DescribeHighlightInfoRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeHighlightInfoRequest) SetLang(v string) *DescribeHighlightInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeHighlightInfoRequest) SetStartTime(v string) *DescribeHighlightInfoRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeHighlightInfoRequest) SetTraceId(v string) *DescribeHighlightInfoRequest {
	s.TraceId = &v
	return s
}

func (s *DescribeHighlightInfoRequest) Validate() error {
	return dara.Validate(s)
}

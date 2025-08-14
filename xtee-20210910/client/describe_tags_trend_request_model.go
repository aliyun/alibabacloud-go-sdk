// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeTagsTrendRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeTagsTrendRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeTagsTrendRequest
	GetEndTime() *int64
	SetEventCodes(v string) *DescribeTagsTrendRequest
	GetEventCodes() *string
	SetRegId(v string) *DescribeTagsTrendRequest
	GetRegId() *string
	SetResult(v string) *DescribeTagsTrendRequest
	GetResult() *string
}

type DescribeTagsTrendRequest struct {
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
	// example:
	//
	// 1751249559000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// example:
	//
	// 1740298093000
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

func (s DescribeTagsTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTagsTrendRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeTagsTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeTagsTrendRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeTagsTrendRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeTagsTrendRequest) GetResult() *string {
	return s.Result
}

func (s *DescribeTagsTrendRequest) SetLang(v string) *DescribeTagsTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTagsTrendRequest) SetBeginTime(v int64) *DescribeTagsTrendRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeTagsTrendRequest) SetEndTime(v int64) *DescribeTagsTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTagsTrendRequest) SetEventCodes(v string) *DescribeTagsTrendRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeTagsTrendRequest) SetRegId(v string) *DescribeTagsTrendRequest {
	s.RegId = &v
	return s
}

func (s *DescribeTagsTrendRequest) SetResult(v string) *DescribeTagsTrendRequest {
	s.Result = &v
	return s
}

func (s *DescribeTagsTrendRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResultCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeResultCountRequest
	GetLang() *string
	SetBeginTime(v int64) *DescribeResultCountRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *DescribeResultCountRequest
	GetEndTime() *int64
	SetRegId(v string) *DescribeResultCountRequest
	GetRegId() *string
}

type DescribeResultCountRequest struct {
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
	// 1739841750000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1739843750000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeResultCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeResultCountRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeResultCountRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeResultCountRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeResultCountRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeResultCountRequest) SetLang(v string) *DescribeResultCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResultCountRequest) SetBeginTime(v int64) *DescribeResultCountRequest {
	s.BeginTime = &v
	return s
}

func (s *DescribeResultCountRequest) SetEndTime(v int64) *DescribeResultCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeResultCountRequest) SetRegId(v string) *DescribeResultCountRequest {
	s.RegId = &v
	return s
}

func (s *DescribeResultCountRequest) Validate() error {
	return dara.Validate(s)
}

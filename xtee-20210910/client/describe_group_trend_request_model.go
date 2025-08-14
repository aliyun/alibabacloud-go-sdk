// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeGroupTrendRequest
	GetLang() *string
	SetDay(v string) *DescribeGroupTrendRequest
	GetDay() *string
	SetRegId(v string) *DescribeGroupTrendRequest
	GetRegId() *string
}

type DescribeGroupTrendRequest struct {
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
	// day
	//
	// example:
	//
	// 1
	Day *string `json:"day,omitempty" xml:"day,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeGroupTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeGroupTrendRequest) GetDay() *string {
	return s.Day
}

func (s *DescribeGroupTrendRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeGroupTrendRequest) SetLang(v string) *DescribeGroupTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeGroupTrendRequest) SetDay(v string) *DescribeGroupTrendRequest {
	s.Day = &v
	return s
}

func (s *DescribeGroupTrendRequest) SetRegId(v string) *DescribeGroupTrendRequest {
	s.RegId = &v
	return s
}

func (s *DescribeGroupTrendRequest) Validate() error {
	return dara.Validate(s)
}

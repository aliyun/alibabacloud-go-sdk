// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnprotectedPortTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeUnprotectedPortTrendRequest
	GetEndTime() *string
	SetLang(v string) *DescribeUnprotectedPortTrendRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeUnprotectedPortTrendRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeUnprotectedPortTrendRequest
	GetStartTime() *string
}

type DescribeUnprotectedPortTrendRequest struct {
	// example:
	//
	// 1751210395
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 61.155.60.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1655778046
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUnprotectedPortTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnprotectedPortTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeUnprotectedPortTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUnprotectedPortTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeUnprotectedPortTrendRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeUnprotectedPortTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUnprotectedPortTrendRequest) SetEndTime(v string) *DescribeUnprotectedPortTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUnprotectedPortTrendRequest) SetLang(v string) *DescribeUnprotectedPortTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUnprotectedPortTrendRequest) SetSourceIp(v string) *DescribeUnprotectedPortTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUnprotectedPortTrendRequest) SetStartTime(v string) *DescribeUnprotectedPortTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUnprotectedPortTrendRequest) Validate() error {
	return dara.Validate(s)
}

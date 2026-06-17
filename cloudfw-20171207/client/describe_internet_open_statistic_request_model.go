// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInternetOpenStatisticRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInternetOpenStatisticRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeInternetOpenStatisticRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInternetOpenStatisticRequest
	GetStartTime() *string
}

type DescribeInternetOpenStatisticRequest struct {
	// The end of the time range to query. This value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1736386501
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the request and response.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the visitor.
	//
	// example:
	//
	// 202.109.244.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The start of the time range to query. This value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1734386501
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInternetOpenStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenStatisticRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInternetOpenStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetOpenStatisticRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetOpenStatisticRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInternetOpenStatisticRequest) SetEndTime(v string) *DescribeInternetOpenStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetOpenStatisticRequest) SetLang(v string) *DescribeInternetOpenStatisticRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetOpenStatisticRequest) SetSourceIp(v string) *DescribeInternetOpenStatisticRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetOpenStatisticRequest) SetStartTime(v string) *DescribeInternetOpenStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetOpenStatisticRequest) Validate() error {
	return dara.Validate(s)
}

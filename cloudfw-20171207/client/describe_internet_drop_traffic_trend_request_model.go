// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetDropTrafficTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeInternetDropTrafficTrendRequest
	GetDirection() *string
	SetEndTime(v string) *DescribeInternetDropTrafficTrendRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInternetDropTrafficTrendRequest
	GetLang() *string
	SetSourceCode(v string) *DescribeInternetDropTrafficTrendRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeInternetDropTrafficTrendRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInternetDropTrafficTrendRequest
	GetStartTime() *string
}

type DescribeInternetDropTrafficTrendRequest struct {
	// The traffic direction.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end time of the query, specified as a timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1756346821
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The traceability code.
	//
	// This parameter is required.
	//
	// example:
	//
	// yundun
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 120.136.21.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The start time, specified as a timestamp in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1749176793
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInternetDropTrafficTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDropTrafficTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetDropTrafficTrendRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeInternetDropTrafficTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInternetDropTrafficTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetDropTrafficTrendRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeInternetDropTrafficTrendRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetDropTrafficTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInternetDropTrafficTrendRequest) SetDirection(v string) *DescribeInternetDropTrafficTrendRequest {
	s.Direction = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendRequest) SetEndTime(v string) *DescribeInternetDropTrafficTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendRequest) SetLang(v string) *DescribeInternetDropTrafficTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendRequest) SetSourceCode(v string) *DescribeInternetDropTrafficTrendRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendRequest) SetSourceIp(v string) *DescribeInternetDropTrafficTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendRequest) SetStartTime(v string) *DescribeInternetDropTrafficTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetDropTrafficTrendRequest) Validate() error {
	return dara.Validate(s)
}

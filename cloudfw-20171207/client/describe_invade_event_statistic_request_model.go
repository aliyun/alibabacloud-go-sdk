// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEventStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInvadeEventStatisticRequest
	GetEndTime() *string
	SetLang(v string) *DescribeInvadeEventStatisticRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeInvadeEventStatisticRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeInvadeEventStatisticRequest
	GetStartTime() *string
}

type DescribeInvadeEventStatisticRequest struct {
	// The end time. This is a UNIX timestamp that is accurate to the second.
	//
	// example:
	//
	// 1774713600
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the visitor.
	//
	// example:
	//
	// 52.130.200.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The start time. This is a UNIX timestamp that is accurate to the second.
	//
	// example:
	//
	// 1746151757
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvadeEventStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEventStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventStatisticRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInvadeEventStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInvadeEventStatisticRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInvadeEventStatisticRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvadeEventStatisticRequest) SetEndTime(v string) *DescribeInvadeEventStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInvadeEventStatisticRequest) SetLang(v string) *DescribeInvadeEventStatisticRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInvadeEventStatisticRequest) SetSourceIp(v string) *DescribeInvadeEventStatisticRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInvadeEventStatisticRequest) SetStartTime(v string) *DescribeInvadeEventStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInvadeEventStatisticRequest) Validate() error {
	return dara.Validate(s)
}

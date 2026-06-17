// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingRiskTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeOutgoingRiskTrendRequest
	GetEndTime() *string
	SetLang(v string) *DescribeOutgoingRiskTrendRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeOutgoingRiskTrendRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeOutgoingRiskTrendRequest
	GetStartTime() *string
}

type DescribeOutgoingRiskTrendRequest struct {
	// The end of the time range to query. Use a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1755051062
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the response message.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 219.145.94.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The start of the time range to query. Use a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1733882648
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOutgoingRiskTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingRiskTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingRiskTrendRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingRiskTrendRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingRiskTrendRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOutgoingRiskTrendRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingRiskTrendRequest) SetEndTime(v string) *DescribeOutgoingRiskTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingRiskTrendRequest) SetLang(v string) *DescribeOutgoingRiskTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingRiskTrendRequest) SetSourceIp(v string) *DescribeOutgoingRiskTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOutgoingRiskTrendRequest) SetStartTime(v string) *DescribeOutgoingRiskTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingRiskTrendRequest) Validate() error {
	return dara.Validate(s)
}

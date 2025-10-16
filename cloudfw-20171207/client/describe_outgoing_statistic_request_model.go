// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeOutgoingStatisticRequest
	GetEndTime() *string
	SetLang(v string) *DescribeOutgoingStatisticRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeOutgoingStatisticRequest
	GetSourceIp() *string
	SetStartTime(v string) *DescribeOutgoingStatisticRequest
	GetStartTime() *string
}

type DescribeOutgoingStatisticRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1734920543
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 117.32.136.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1746554400
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOutgoingStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingStatisticRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeOutgoingStatisticRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOutgoingStatisticRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeOutgoingStatisticRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeOutgoingStatisticRequest) SetEndTime(v string) *DescribeOutgoingStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingStatisticRequest) SetLang(v string) *DescribeOutgoingStatisticRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingStatisticRequest) SetSourceIp(v string) *DescribeOutgoingStatisticRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOutgoingStatisticRequest) SetStartTime(v string) *DescribeOutgoingStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingStatisticRequest) Validate() error {
	return dara.Validate(s)
}

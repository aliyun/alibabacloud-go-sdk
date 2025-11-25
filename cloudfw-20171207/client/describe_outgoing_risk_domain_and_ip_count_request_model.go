// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingRiskDomainAndIpCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeOutgoingRiskDomainAndIpCountRequest
	GetEndTime() *int64
	SetStartTime(v int64) *DescribeOutgoingRiskDomainAndIpCountRequest
	GetStartTime() *int64
}

type DescribeOutgoingRiskDomainAndIpCountRequest struct {
	// example:
	//
	// 1751595213
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1749434787
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOutgoingRiskDomainAndIpCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingRiskDomainAndIpCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingRiskDomainAndIpCountRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeOutgoingRiskDomainAndIpCountRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeOutgoingRiskDomainAndIpCountRequest) SetEndTime(v int64) *DescribeOutgoingRiskDomainAndIpCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingRiskDomainAndIpCountRequest) SetStartTime(v int64) *DescribeOutgoingRiskDomainAndIpCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingRiskDomainAndIpCountRequest) Validate() error {
	return dara.Validate(s)
}

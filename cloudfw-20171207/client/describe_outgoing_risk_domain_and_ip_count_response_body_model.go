// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingRiskDomainAndIpCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOutgoingRiskDomainAndIpCountResponseBody
	GetRequestId() *string
	SetRiskDomainCount(v int64) *DescribeOutgoingRiskDomainAndIpCountResponseBody
	GetRiskDomainCount() *int64
	SetRiskIpCount(v int64) *DescribeOutgoingRiskDomainAndIpCountResponseBody
	GetRiskIpCount() *int64
	SetTotalCount(v int64) *DescribeOutgoingRiskDomainAndIpCountResponseBody
	GetTotalCount() *int64
}

type DescribeOutgoingRiskDomainAndIpCountResponseBody struct {
	// example:
	//
	// EE258AC0-6EDD-5929-AB47-165E9B54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6
	RiskDomainCount *int64 `json:"RiskDomainCount,omitempty" xml:"RiskDomainCount,omitempty"`
	// example:
	//
	// 47
	RiskIpCount *int64 `json:"RiskIpCount,omitempty" xml:"RiskIpCount,omitempty"`
	// example:
	//
	// 6
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingRiskDomainAndIpCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingRiskDomainAndIpCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponseBody) GetRiskDomainCount() *int64 {
	return s.RiskDomainCount
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponseBody) GetRiskIpCount() *int64 {
	return s.RiskIpCount
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponseBody) SetRequestId(v string) *DescribeOutgoingRiskDomainAndIpCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponseBody) SetRiskDomainCount(v int64) *DescribeOutgoingRiskDomainAndIpCountResponseBody {
	s.RiskDomainCount = &v
	return s
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponseBody) SetRiskIpCount(v int64) *DescribeOutgoingRiskDomainAndIpCountResponseBody {
	s.RiskIpCount = &v
	return s
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponseBody) SetTotalCount(v int64) *DescribeOutgoingRiskDomainAndIpCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOutgoingRiskDomainAndIpCountResponseBody) Validate() error {
	return dara.Validate(s)
}

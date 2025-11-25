// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDropTrafficTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeVpcFirewallDropTrafficTrendRequest
	GetEndTime() *int64
	SetOrder(v string) *DescribeVpcFirewallDropTrafficTrendRequest
	GetOrder() *string
	SetSort(v string) *DescribeVpcFirewallDropTrafficTrendRequest
	GetSort() *string
	SetSourceIp(v string) *DescribeVpcFirewallDropTrafficTrendRequest
	GetSourceIp() *string
	SetStartTime(v int64) *DescribeVpcFirewallDropTrafficTrendRequest
	GetStartTime() *int64
	SetTrafficTime(v int64) *DescribeVpcFirewallDropTrafficTrendRequest
	GetTrafficTime() *int64
}

type DescribeVpcFirewallDropTrafficTrendRequest struct {
	// example:
	//
	// 1747792853
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// LastTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 183.237.161.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1656664560
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1739337840
	TrafficTime *int64 `json:"TrafficTime,omitempty" xml:"TrafficTime,omitempty"`
}

func (s DescribeVpcFirewallDropTrafficTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDropTrafficTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) GetTrafficTime() *int64 {
	return s.TrafficTime
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) SetEndTime(v int64) *DescribeVpcFirewallDropTrafficTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) SetOrder(v string) *DescribeVpcFirewallDropTrafficTrendRequest {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) SetSort(v string) *DescribeVpcFirewallDropTrafficTrendRequest {
	s.Sort = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) SetSourceIp(v string) *DescribeVpcFirewallDropTrafficTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) SetStartTime(v int64) *DescribeVpcFirewallDropTrafficTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) SetTrafficTime(v int64) *DescribeVpcFirewallDropTrafficTrendRequest {
	s.TrafficTime = &v
	return s
}

func (s *DescribeVpcFirewallDropTrafficTrendRequest) Validate() error {
	return dara.Validate(s)
}

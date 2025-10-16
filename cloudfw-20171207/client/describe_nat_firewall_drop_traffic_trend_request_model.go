// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallDropTrafficTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeNatFirewallDropTrafficTrendRequest
	GetEndTime() *int64
	SetSourceIp(v string) *DescribeNatFirewallDropTrafficTrendRequest
	GetSourceIp() *string
	SetStartTime(v int64) *DescribeNatFirewallDropTrafficTrendRequest
	GetStartTime() *int64
}

type DescribeNatFirewallDropTrafficTrendRequest struct {
	// example:
	//
	// 1758334822
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 122.190.56.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1740968766
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeNatFirewallDropTrafficTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallDropTrafficTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallDropTrafficTrendRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeNatFirewallDropTrafficTrendRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeNatFirewallDropTrafficTrendRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeNatFirewallDropTrafficTrendRequest) SetEndTime(v int64) *DescribeNatFirewallDropTrafficTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendRequest) SetSourceIp(v string) *DescribeNatFirewallDropTrafficTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendRequest) SetStartTime(v int64) *DescribeNatFirewallDropTrafficTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeNatFirewallDropTrafficTrendRequest) Validate() error {
	return dara.Validate(s)
}

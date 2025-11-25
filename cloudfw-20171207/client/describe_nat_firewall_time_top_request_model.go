// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallTimeTopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInterval(v int64) *DescribeNatFirewallTimeTopRequest
	GetInterval() *int64
	SetLang(v string) *DescribeNatFirewallTimeTopRequest
	GetLang() *string
	SetLimit(v int64) *DescribeNatFirewallTimeTopRequest
	GetLimit() *int64
	SetNatGatewayId(v string) *DescribeNatFirewallTimeTopRequest
	GetNatGatewayId() *string
	SetSort(v string) *DescribeNatFirewallTimeTopRequest
	GetSort() *string
	SetSrcPrivateIP(v string) *DescribeNatFirewallTimeTopRequest
	GetSrcPrivateIP() *string
	SetSrcPublicIP(v string) *DescribeNatFirewallTimeTopRequest
	GetSrcPublicIP() *string
	SetTrafficTime(v string) *DescribeNatFirewallTimeTopRequest
	GetTrafficTime() *string
}

type DescribeNatFirewallTimeTopRequest struct {
	// example:
	//
	// 60
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 20
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// ngw-uf62zzi7000bca7zn****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// example:
	//
	// total_max_bps
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 10.100.134.XXX
	SrcPrivateIP *string `json:"SrcPrivateIP,omitempty" xml:"SrcPrivateIP,omitempty"`
	// example:
	//
	// 47.93.47.XXX
	SrcPublicIP *string `json:"SrcPublicIP,omitempty" xml:"SrcPublicIP,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1749693960
	TrafficTime *string `json:"TrafficTime,omitempty" xml:"TrafficTime,omitempty"`
}

func (s DescribeNatFirewallTimeTopRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallTimeTopRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallTimeTopRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeNatFirewallTimeTopRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNatFirewallTimeTopRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeNatFirewallTimeTopRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallTimeTopRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeNatFirewallTimeTopRequest) GetSrcPrivateIP() *string {
	return s.SrcPrivateIP
}

func (s *DescribeNatFirewallTimeTopRequest) GetSrcPublicIP() *string {
	return s.SrcPublicIP
}

func (s *DescribeNatFirewallTimeTopRequest) GetTrafficTime() *string {
	return s.TrafficTime
}

func (s *DescribeNatFirewallTimeTopRequest) SetInterval(v int64) *DescribeNatFirewallTimeTopRequest {
	s.Interval = &v
	return s
}

func (s *DescribeNatFirewallTimeTopRequest) SetLang(v string) *DescribeNatFirewallTimeTopRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatFirewallTimeTopRequest) SetLimit(v int64) *DescribeNatFirewallTimeTopRequest {
	s.Limit = &v
	return s
}

func (s *DescribeNatFirewallTimeTopRequest) SetNatGatewayId(v string) *DescribeNatFirewallTimeTopRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallTimeTopRequest) SetSort(v string) *DescribeNatFirewallTimeTopRequest {
	s.Sort = &v
	return s
}

func (s *DescribeNatFirewallTimeTopRequest) SetSrcPrivateIP(v string) *DescribeNatFirewallTimeTopRequest {
	s.SrcPrivateIP = &v
	return s
}

func (s *DescribeNatFirewallTimeTopRequest) SetSrcPublicIP(v string) *DescribeNatFirewallTimeTopRequest {
	s.SrcPublicIP = &v
	return s
}

func (s *DescribeNatFirewallTimeTopRequest) SetTrafficTime(v string) *DescribeNatFirewallTimeTopRequest {
	s.TrafficTime = &v
	return s
}

func (s *DescribeNatFirewallTimeTopRequest) Validate() error {
	return dara.Validate(s)
}

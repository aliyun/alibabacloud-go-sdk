// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetTimeTopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeInternetTimeTopRequest
	GetDirection() *string
	SetEndTime(v string) *DescribeInternetTimeTopRequest
	GetEndTime() *string
	SetIPType(v string) *DescribeInternetTimeTopRequest
	GetIPType() *string
	SetInterval(v int64) *DescribeInternetTimeTopRequest
	GetInterval() *int64
	SetLang(v string) *DescribeInternetTimeTopRequest
	GetLang() *string
	SetLimit(v string) *DescribeInternetTimeTopRequest
	GetLimit() *string
	SetNatIP(v string) *DescribeInternetTimeTopRequest
	GetNatIP() *string
	SetOrder(v string) *DescribeInternetTimeTopRequest
	GetOrder() *string
	SetSort(v string) *DescribeInternetTimeTopRequest
	GetSort() *string
	SetSourceCode(v string) *DescribeInternetTimeTopRequest
	GetSourceCode() *string
	SetSourceIp(v string) *DescribeInternetTimeTopRequest
	GetSourceIp() *string
	SetSrcIP(v string) *DescribeInternetTimeTopRequest
	GetSrcIP() *string
	SetStartTime(v string) *DescribeInternetTimeTopRequest
	GetStartTime() *string
	SetTrafficTime(v string) *DescribeInternetTimeTopRequest
	GetTrafficTime() *string
	SetTrafficType(v string) *DescribeInternetTimeTopRequest
	GetTrafficType() *string
}

type DescribeInternetTimeTopRequest struct {
	// The direction of the traffic.
	//
	// This parameter is required.
	//
	// example:
	//
	// out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the query time range, specified as a Unix timestamp in seconds.
	//
	// example:
	//
	// 1733796528
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The network type of the IP address.
	//
	// example:
	//
	// Public
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The time interval in seconds.
	//
	// example:
	//
	// 60
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The language of the response.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return. The value must be between 1 and 50. Default value: 50.
	//
	// example:
	//
	// 5
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The NAT public IP address. If omitted, traffic data for all assets is returned.
	//
	// example:
	//
	// 47.97.66.XXX
	NatIP *string `json:"NatIP,omitempty" xml:"NatIP,omitempty"`
	// The sort order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The metric by which to sort the results.
	//
	// example:
	//
	// in_bps
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The identifier for the trace source.
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
	// 60.179.179.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 8.153.18.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The start of the query time range, specified as a Unix timestamp in seconds.
	//
	// example:
	//
	// 1749434787
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The traffic timestamp to query, specified as a Unix timestamp in seconds.
	//
	// example:
	//
	// 1745222880
	TrafficTime *string `json:"TrafficTime,omitempty" xml:"TrafficTime,omitempty"`
	// The traffic type.
	//
	// example:
	//
	// EIP_TRAFFIC
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
}

func (s DescribeInternetTimeTopRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetTimeTopRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetTimeTopRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeInternetTimeTopRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInternetTimeTopRequest) GetIPType() *string {
	return s.IPType
}

func (s *DescribeInternetTimeTopRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeInternetTimeTopRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeInternetTimeTopRequest) GetLimit() *string {
	return s.Limit
}

func (s *DescribeInternetTimeTopRequest) GetNatIP() *string {
	return s.NatIP
}

func (s *DescribeInternetTimeTopRequest) GetOrder() *string {
	return s.Order
}

func (s *DescribeInternetTimeTopRequest) GetSort() *string {
	return s.Sort
}

func (s *DescribeInternetTimeTopRequest) GetSourceCode() *string {
	return s.SourceCode
}

func (s *DescribeInternetTimeTopRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInternetTimeTopRequest) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeInternetTimeTopRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInternetTimeTopRequest) GetTrafficTime() *string {
	return s.TrafficTime
}

func (s *DescribeInternetTimeTopRequest) GetTrafficType() *string {
	return s.TrafficType
}

func (s *DescribeInternetTimeTopRequest) SetDirection(v string) *DescribeInternetTimeTopRequest {
	s.Direction = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetEndTime(v string) *DescribeInternetTimeTopRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetIPType(v string) *DescribeInternetTimeTopRequest {
	s.IPType = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetInterval(v int64) *DescribeInternetTimeTopRequest {
	s.Interval = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetLang(v string) *DescribeInternetTimeTopRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetLimit(v string) *DescribeInternetTimeTopRequest {
	s.Limit = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetNatIP(v string) *DescribeInternetTimeTopRequest {
	s.NatIP = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetOrder(v string) *DescribeInternetTimeTopRequest {
	s.Order = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetSort(v string) *DescribeInternetTimeTopRequest {
	s.Sort = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetSourceCode(v string) *DescribeInternetTimeTopRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetSourceIp(v string) *DescribeInternetTimeTopRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetSrcIP(v string) *DescribeInternetTimeTopRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetStartTime(v string) *DescribeInternetTimeTopRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetTrafficTime(v string) *DescribeInternetTimeTopRequest {
	s.TrafficTime = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) SetTrafficType(v string) *DescribeInternetTimeTopRequest {
	s.TrafficType = &v
	return s
}

func (s *DescribeInternetTimeTopRequest) Validate() error {
	return dara.Validate(s)
}

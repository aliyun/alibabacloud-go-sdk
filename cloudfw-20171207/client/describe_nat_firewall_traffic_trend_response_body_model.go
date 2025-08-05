// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallTrafficTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeNatFirewallTrafficTrendResponseBodyDataList) *DescribeNatFirewallTrafficTrendResponseBody
	GetDataList() []*DescribeNatFirewallTrafficTrendResponseBodyDataList
	SetMaxInBps(v int64) *DescribeNatFirewallTrafficTrendResponseBody
	GetMaxInBps() *int64
	SetMaxOutBps(v int64) *DescribeNatFirewallTrafficTrendResponseBody
	GetMaxOutBps() *int64
	SetMaxTotalBps(v int64) *DescribeNatFirewallTrafficTrendResponseBody
	GetMaxTotalBps() *int64
	SetRequestId(v string) *DescribeNatFirewallTrafficTrendResponseBody
	GetRequestId() *string
}

type DescribeNatFirewallTrafficTrendResponseBody struct {
	// The statistics on traffic.
	DataList []*DescribeNatFirewallTrafficTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The maximum inbound network throughput, which indicates the maximum number of bits that are sent inbound per second. Unit: bit/s.
	//
	// example:
	//
	// 18038632
	MaxInBps *int64 `json:"MaxInBps,omitempty" xml:"MaxInBps,omitempty"`
	// The maximum outbound network throughput, which indicates the maximum number of bits that are sent outbound per second. Unit: bit/s.
	//
	// example:
	//
	// 122596487
	MaxOutBps *int64 `json:"MaxOutBps,omitempty" xml:"MaxOutBps,omitempty"`
	// The total maximum inbound and outbound network throughput, which indicates the maximum number of bits that are sent inbound and outbound per second. Unit: bit/s.
	//
	// example:
	//
	// 66953194
	MaxTotalBps *int64 `json:"MaxTotalBps,omitempty" xml:"MaxTotalBps,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 87F23A3A-6F57-59C3-8237-A090D0613D71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNatFirewallTrafficTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallTrafficTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) GetDataList() []*DescribeNatFirewallTrafficTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) GetMaxInBps() *int64 {
	return s.MaxInBps
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) GetMaxOutBps() *int64 {
	return s.MaxOutBps
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) GetMaxTotalBps() *int64 {
	return s.MaxTotalBps
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) SetDataList(v []*DescribeNatFirewallTrafficTrendResponseBodyDataList) *DescribeNatFirewallTrafficTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) SetMaxInBps(v int64) *DescribeNatFirewallTrafficTrendResponseBody {
	s.MaxInBps = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) SetMaxOutBps(v int64) *DescribeNatFirewallTrafficTrendResponseBody {
	s.MaxOutBps = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) SetMaxTotalBps(v int64) *DescribeNatFirewallTrafficTrendResponseBody {
	s.MaxTotalBps = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) SetRequestId(v string) *DescribeNatFirewallTrafficTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNatFirewallTrafficTrendResponseBodyDataList struct {
	// The maximum inbound network throughput, which indicates the maximum number of bits that are sent inbound per second. Unit: bit/s.
	//
	// example:
	//
	// 18038632
	MaxInBps *int64 `json:"MaxInBps,omitempty" xml:"MaxInBps,omitempty"`
	// The maximum outbound network throughput, which indicates the maximum number of bits that are sent outbound per second. Unit: bit/s.
	//
	// example:
	//
	// 122596487
	MaxOutBps *int64 `json:"MaxOutBps,omitempty" xml:"MaxOutBps,omitempty"`
	// The total maximum inbound and outbound network throughput, which indicates the maximum number of bits that are sent inbound and outbound per second. Unit: bit/s.
	//
	// example:
	//
	// 66953194
	MaxTotalBps *int64 `json:"MaxTotalBps,omitempty" xml:"MaxTotalBps,omitempty"`
	// The time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1734418980
	TrafficTime *int64 `json:"TrafficTime,omitempty" xml:"TrafficTime,omitempty"`
}

func (s DescribeNatFirewallTrafficTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallTrafficTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallTrafficTrendResponseBodyDataList) GetMaxInBps() *int64 {
	return s.MaxInBps
}

func (s *DescribeNatFirewallTrafficTrendResponseBodyDataList) GetMaxOutBps() *int64 {
	return s.MaxOutBps
}

func (s *DescribeNatFirewallTrafficTrendResponseBodyDataList) GetMaxTotalBps() *int64 {
	return s.MaxTotalBps
}

func (s *DescribeNatFirewallTrafficTrendResponseBodyDataList) GetTrafficTime() *int64 {
	return s.TrafficTime
}

func (s *DescribeNatFirewallTrafficTrendResponseBodyDataList) SetMaxInBps(v int64) *DescribeNatFirewallTrafficTrendResponseBodyDataList {
	s.MaxInBps = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponseBodyDataList) SetMaxOutBps(v int64) *DescribeNatFirewallTrafficTrendResponseBodyDataList {
	s.MaxOutBps = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponseBodyDataList) SetMaxTotalBps(v int64) *DescribeNatFirewallTrafficTrendResponseBodyDataList {
	s.MaxTotalBps = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponseBodyDataList) SetTrafficTime(v int64) *DescribeNatFirewallTrafficTrendResponseBodyDataList {
	s.TrafficTime = &v
	return s
}

func (s *DescribeNatFirewallTrafficTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

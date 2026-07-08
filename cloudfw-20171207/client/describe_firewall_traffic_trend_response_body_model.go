// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTrafficTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeFirewallTrafficTrendResponseBodyDataList) *DescribeFirewallTrafficTrendResponseBody
	GetDataList() []*DescribeFirewallTrafficTrendResponseBodyDataList
	SetMaxBandwidthTime(v int64) *DescribeFirewallTrafficTrendResponseBody
	GetMaxBandwidthTime() *int64
	SetMaxBandwidthTimeBps(v *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) *DescribeFirewallTrafficTrendResponseBody
	GetMaxBandwidthTimeBps() *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps
	SetRequestId(v string) *DescribeFirewallTrafficTrendResponseBody
	GetRequestId() *string
}

type DescribeFirewallTrafficTrendResponseBody struct {
	// The returned data list.
	DataList []*DescribeFirewallTrafficTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The timestamp when the peak total traffic occurred. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1758470400
	MaxBandwidthTime *int64 `json:"MaxBandwidthTime,omitempty" xml:"MaxBandwidthTime,omitempty"`
	// The traffic distribution at the time of peak total traffic.
	MaxBandwidthTimeBps *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps `json:"MaxBandwidthTimeBps,omitempty" xml:"MaxBandwidthTimeBps,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4E7F94C7-781F-5192-86CF-DB085****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFirewallTrafficTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTrafficTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTrafficTrendResponseBody) GetDataList() []*DescribeFirewallTrafficTrendResponseBodyDataList {
	return s.DataList
}

func (s *DescribeFirewallTrafficTrendResponseBody) GetMaxBandwidthTime() *int64 {
	return s.MaxBandwidthTime
}

func (s *DescribeFirewallTrafficTrendResponseBody) GetMaxBandwidthTimeBps() *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps {
	return s.MaxBandwidthTimeBps
}

func (s *DescribeFirewallTrafficTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFirewallTrafficTrendResponseBody) SetDataList(v []*DescribeFirewallTrafficTrendResponseBodyDataList) *DescribeFirewallTrafficTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBody) SetMaxBandwidthTime(v int64) *DescribeFirewallTrafficTrendResponseBody {
	s.MaxBandwidthTime = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBody) SetMaxBandwidthTimeBps(v *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) *DescribeFirewallTrafficTrendResponseBody {
	s.MaxBandwidthTimeBps = v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBody) SetRequestId(v string) *DescribeFirewallTrafficTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MaxBandwidthTimeBps != nil {
		if err := s.MaxBandwidthTimeBps.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFirewallTrafficTrendResponseBodyDataList struct {
	// The Internet firewall traffic.
	//
	// example:
	//
	// 100
	InternetBps *int64 `json:"InternetBps,omitempty" xml:"InternetBps,omitempty"`
	// The NAT firewall traffic.
	//
	// example:
	//
	// 100
	NatBps *int64 `json:"NatBps,omitempty" xml:"NatBps,omitempty"`
	// The time when the traffic occurred. The value is a UNIX timestamp. Unit: seconds.
	//
	// If the data at this point in time has not been processed, the values of all other fields are -1.
	//
	// example:
	//
	// 1758470400
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The total firewall traffic.
	//
	// example:
	//
	// 300
	TotalBps *int64 `json:"TotalBps,omitempty" xml:"TotalBps,omitempty"`
	// The VPC firewall traffic.
	//
	// example:
	//
	// 100
	VpcBps *int64 `json:"VpcBps,omitempty" xml:"VpcBps,omitempty"`
}

func (s DescribeFirewallTrafficTrendResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTrafficTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) GetInternetBps() *int64 {
	return s.InternetBps
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) GetNatBps() *int64 {
	return s.NatBps
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) GetTime() *int64 {
	return s.Time
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) GetTotalBps() *int64 {
	return s.TotalBps
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) GetVpcBps() *int64 {
	return s.VpcBps
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) SetInternetBps(v int64) *DescribeFirewallTrafficTrendResponseBodyDataList {
	s.InternetBps = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) SetNatBps(v int64) *DescribeFirewallTrafficTrendResponseBodyDataList {
	s.NatBps = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) SetTime(v int64) *DescribeFirewallTrafficTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) SetTotalBps(v int64) *DescribeFirewallTrafficTrendResponseBodyDataList {
	s.TotalBps = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) SetVpcBps(v int64) *DescribeFirewallTrafficTrendResponseBodyDataList {
	s.VpcBps = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps struct {
	// The Internet firewall traffic at the time of peak total traffic.
	//
	// example:
	//
	// 100
	InternetBps *int64 `json:"InternetBps,omitempty" xml:"InternetBps,omitempty"`
	// The NAT firewall traffic at the time of peak total traffic.
	//
	// example:
	//
	// 100
	NatBps *int64 `json:"NatBps,omitempty" xml:"NatBps,omitempty"`
	// The peak total traffic.
	//
	// example:
	//
	// 1000
	TotalBps *int64 `json:"TotalBps,omitempty" xml:"TotalBps,omitempty"`
	// The VPC firewall traffic at the time of peak total traffic.
	//
	// example:
	//
	// 100
	VpcBps *int64 `json:"VpcBps,omitempty" xml:"VpcBps,omitempty"`
}

func (s DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) GetInternetBps() *int64 {
	return s.InternetBps
}

func (s *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) GetNatBps() *int64 {
	return s.NatBps
}

func (s *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) GetTotalBps() *int64 {
	return s.TotalBps
}

func (s *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) GetVpcBps() *int64 {
	return s.VpcBps
}

func (s *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) SetInternetBps(v int64) *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps {
	s.InternetBps = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) SetNatBps(v int64) *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps {
	s.NatBps = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) SetTotalBps(v int64) *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps {
	s.TotalBps = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) SetVpcBps(v int64) *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps {
	s.VpcBps = &v
	return s
}

func (s *DescribeFirewallTrafficTrendResponseBodyMaxBandwidthTimeBps) Validate() error {
	return dara.Validate(s)
}

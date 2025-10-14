// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsRequestStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribePdnsRequestStatisticResponseBodyData) *DescribePdnsRequestStatisticResponseBody
	GetData() []*DescribePdnsRequestStatisticResponseBodyData
	SetRequestId(v string) *DescribePdnsRequestStatisticResponseBody
	GetRequestId() *string
}

type DescribePdnsRequestStatisticResponseBody struct {
	// The statistics on the DNS requests.
	Data []*DescribePdnsRequestStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePdnsRequestStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsRequestStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePdnsRequestStatisticResponseBody) GetData() []*DescribePdnsRequestStatisticResponseBodyData {
	return s.Data
}

func (s *DescribePdnsRequestStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePdnsRequestStatisticResponseBody) SetData(v []*DescribePdnsRequestStatisticResponseBodyData) *DescribePdnsRequestStatisticResponseBody {
	s.Data = v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBody) SetRequestId(v string) *DescribePdnsRequestStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePdnsRequestStatisticResponseBodyData struct {
	// The total number of DoH requests, including HTTP and HTTPS requests.
	//
	// example:
	//
	// 0
	DohTotalCount *int64 `json:"DohTotalCount,omitempty" xml:"DohTotalCount,omitempty"`
	// The number of HTTP requests.
	//
	// example:
	//
	// 0
	HttpCount *int64 `json:"HttpCount,omitempty" xml:"HttpCount,omitempty"`
	// The number of HTTPS requests. On the Traffic Analysis tab of the Public DNS console, the value of this parameter includes the number of DNS over HTTPs (DoH) requests. Therefore, the number of DoH requests is not separately displayed in the console.
	//
	// example:
	//
	// 0
	HttpsCount *int64 `json:"HttpsCount,omitempty" xml:"HttpsCount,omitempty"`
	// The number of source IP addresses.
	//
	// example:
	//
	// 10
	IpCount *int64 `json:"IpCount,omitempty" xml:"IpCount,omitempty"`
	// The statistical timestamp. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1706716800000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of UDP requests.
	//
	// example:
	//
	// 5000
	UdpTotalCount *int64 `json:"UdpTotalCount,omitempty" xml:"UdpTotalCount,omitempty"`
	// The number of IPv4-based requests.
	//
	// example:
	//
	// 0
	V4Count *int64 `json:"V4Count,omitempty" xml:"V4Count,omitempty"`
	// The number of IPv4-based HTTP requests.
	//
	// example:
	//
	// 0
	V4HttpCount *int64 `json:"V4HttpCount,omitempty" xml:"V4HttpCount,omitempty"`
	// The number of IPv4-based HTTPS requests.
	//
	// example:
	//
	// 0
	V4HttpsCount *int64 `json:"V4HttpsCount,omitempty" xml:"V4HttpsCount,omitempty"`
	// The number of IPv6-based requests.
	//
	// example:
	//
	// 0
	V6Count *int64 `json:"V6Count,omitempty" xml:"V6Count,omitempty"`
	// The number of IPv6-based HTTP requests.
	//
	// example:
	//
	// 0
	V6HttpCount *int64 `json:"V6HttpCount,omitempty" xml:"V6HttpCount,omitempty"`
	// The number of IPv6-based HTTPS requests.
	//
	// example:
	//
	// 0
	V6HttpsCount *int64 `json:"V6HttpsCount,omitempty" xml:"V6HttpsCount,omitempty"`
}

func (s DescribePdnsRequestStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsRequestStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetDohTotalCount() *int64 {
	return s.DohTotalCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetHttpCount() *int64 {
	return s.HttpCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetHttpsCount() *int64 {
	return s.HttpsCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetIpCount() *int64 {
	return s.IpCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetUdpTotalCount() *int64 {
	return s.UdpTotalCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetV4Count() *int64 {
	return s.V4Count
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetV4HttpCount() *int64 {
	return s.V4HttpCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetV4HttpsCount() *int64 {
	return s.V4HttpsCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetV6Count() *int64 {
	return s.V6Count
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetV6HttpCount() *int64 {
	return s.V6HttpCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) GetV6HttpsCount() *int64 {
	return s.V6HttpsCount
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetDohTotalCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.DohTotalCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetHttpCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.HttpCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetHttpsCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.HttpsCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetIpCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.IpCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetTimestamp(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetTotalCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetUdpTotalCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.UdpTotalCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetV4Count(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.V4Count = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetV4HttpCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.V4HttpCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetV4HttpsCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.V4HttpsCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetV6Count(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.V6Count = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetV6HttpCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.V6HttpCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) SetV6HttpsCount(v int64) *DescribePdnsRequestStatisticResponseBodyData {
	s.V6HttpsCount = &v
	return s
}

func (s *DescribePdnsRequestStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}

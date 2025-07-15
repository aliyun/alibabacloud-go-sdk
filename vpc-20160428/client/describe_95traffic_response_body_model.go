// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribe95TrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *Describe95TrafficResponseBody
	GetRequestId() *string
	SetTraffic95Summary(v *Describe95TrafficResponseBodyTraffic95Summary) *Describe95TrafficResponseBody
	GetTraffic95Summary() *Describe95TrafficResponseBodyTraffic95Summary
}

type Describe95TrafficResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information returned.
	Traffic95Summary *Describe95TrafficResponseBodyTraffic95Summary `json:"Traffic95Summary,omitempty" xml:"Traffic95Summary,omitempty" type:"Struct"`
}

func (s Describe95TrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s Describe95TrafficResponseBody) GoString() string {
	return s.String()
}

func (s *Describe95TrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *Describe95TrafficResponseBody) GetTraffic95Summary() *Describe95TrafficResponseBodyTraffic95Summary {
	return s.Traffic95Summary
}

func (s *Describe95TrafficResponseBody) SetRequestId(v string) *Describe95TrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *Describe95TrafficResponseBody) SetTraffic95Summary(v *Describe95TrafficResponseBodyTraffic95Summary) *Describe95TrafficResponseBody {
	s.Traffic95Summary = v
	return s
}

func (s *Describe95TrafficResponseBody) Validate() error {
	return dara.Validate(s)
}

type Describe95TrafficResponseBodyTraffic95Summary struct {
	// The peak bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s.
	//
	// example:
	//
	// 20000.0
	Bandwidth *int64 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The daily peak bandwidth. Unit: Mbit/s.
	//
	// <props="china"> For more information, see [Daily peak bandwidth](https://help.aliyun.com/document_detail/89729.html).
	//
	// example:
	//
	// 1064.244837773641
	FifthPeakBandwidth *string `json:"FifthPeakBandwidth,omitempty" xml:"FifthPeakBandwidth,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// cbwp-wz9j19xrwf78fvz7*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metering method of the Internet Shared Bandwidth instance. Valid values:
	//
	// 	- PayBy95: pay-by-enhanced-95th-percentile
	//
	// 	- PayByBandwidth: pay-by-bandwidth
	//
	// 	- PayByDominantTraffic: pay-by-dominant-traffic
	//
	// example:
	//
	// PayBy95
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The guaranteed bandwidth of the Internet Shared Bandwidth instance. Unit: Mbit/s.
	//
	// example:
	//
	// 0.0
	MinimumConsumeBandwidth *string `json:"MinimumConsumeBandwidth,omitempty" xml:"MinimumConsumeBandwidth,omitempty"`
	// The average bandwidth every 5 minutes in the inbound and outbound directions.
	Traffic95DetailList *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList `json:"Traffic95DetailList,omitempty" xml:"Traffic95DetailList,omitempty" type:"Struct"`
}

func (s Describe95TrafficResponseBodyTraffic95Summary) String() string {
	return dara.Prettify(s)
}

func (s Describe95TrafficResponseBodyTraffic95Summary) GoString() string {
	return s.String()
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) GetFifthPeakBandwidth() *string {
	return s.FifthPeakBandwidth
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) GetInstanceId() *string {
	return s.InstanceId
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) GetMinimumConsumeBandwidth() *string {
	return s.MinimumConsumeBandwidth
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) GetTraffic95DetailList() *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList {
	return s.Traffic95DetailList
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) SetBandwidth(v int64) *Describe95TrafficResponseBodyTraffic95Summary {
	s.Bandwidth = &v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) SetFifthPeakBandwidth(v string) *Describe95TrafficResponseBodyTraffic95Summary {
	s.FifthPeakBandwidth = &v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) SetInstanceId(v string) *Describe95TrafficResponseBodyTraffic95Summary {
	s.InstanceId = &v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) SetInternetChargeType(v string) *Describe95TrafficResponseBodyTraffic95Summary {
	s.InternetChargeType = &v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) SetMinimumConsumeBandwidth(v string) *Describe95TrafficResponseBodyTraffic95Summary {
	s.MinimumConsumeBandwidth = &v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) SetTraffic95DetailList(v *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList) *Describe95TrafficResponseBodyTraffic95Summary {
	s.Traffic95DetailList = v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95Summary) Validate() error {
	return dara.Validate(s)
}

type Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList struct {
	Traffic95Detail []*Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail `json:"Traffic95Detail,omitempty" xml:"Traffic95Detail,omitempty" type:"Repeated"`
}

func (s Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList) String() string {
	return dara.Prettify(s)
}

func (s Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList) GoString() string {
	return s.String()
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList) GetTraffic95Detail() []*Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail {
	return s.Traffic95Detail
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList) SetTraffic95Detail(v []*Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList {
	s.Traffic95Detail = v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailList) Validate() error {
	return dara.Validate(s)
}

type Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail struct {
	// The sampled bandwidth value, which is the larger bandwidth value in the inbound and outbound directions within a sampling interval. Unit: Mbit/s.
	//
	// example:
	//
	// 118.5090322113037
	BillBandwidth *string `json:"BillBandwidth,omitempty" xml:"BillBandwidth,omitempty"`
	// The inbound bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 118.5090322113037
	InBandwidth *string `json:"InBandwidth,omitempty" xml:"InBandwidth,omitempty"`
	// The outbound bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 96.41217480977376
	OutBandwidth *string `json:"OutBandwidth,omitempty" xml:"OutBandwidth,omitempty"`
	// The statistical time. The value is a string.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) String() string {
	return dara.Prettify(s)
}

func (s Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) GoString() string {
	return s.String()
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) GetBillBandwidth() *string {
	return s.BillBandwidth
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) GetInBandwidth() *string {
	return s.InBandwidth
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) GetOutBandwidth() *string {
	return s.OutBandwidth
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) GetTime() *string {
	return s.Time
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) SetBillBandwidth(v string) *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail {
	s.BillBandwidth = &v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) SetInBandwidth(v string) *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail {
	s.InBandwidth = &v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) SetOutBandwidth(v string) *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail {
	s.OutBandwidth = &v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) SetTime(v string) *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail {
	s.Time = &v
	return s
}

func (s *Describe95TrafficResponseBodyTraffic95SummaryTraffic95DetailListTraffic95Detail) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregateInternetTraffic(v float32) *DescribePostpayBillResponseBody
	GetAggregateInternetTraffic() *float32
	SetAggregateNatTraffic(v float32) *DescribePostpayBillResponseBody
	GetAggregateNatTraffic() *float32
	SetAggregateSdlTraffic(v float32) *DescribePostpayBillResponseBody
	GetAggregateSdlTraffic() *float32
	SetAggregateTotalTraffic(v float32) *DescribePostpayBillResponseBody
	GetAggregateTotalTraffic() *float32
	SetAggregateVpcTraffic(v float32) *DescribePostpayBillResponseBody
	GetAggregateVpcTraffic() *float32
	SetBillList(v []*DescribePostpayBillResponseBodyBillList) *DescribePostpayBillResponseBody
	GetBillList() []*DescribePostpayBillResponseBodyBillList
	SetRequestId(v string) *DescribePostpayBillResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePostpayBillResponseBody
	GetTotalCount() *int32
}

type DescribePostpayBillResponseBody struct {
	// The aggregated Internet traffic, in GB.
	//
	// example:
	//
	// 10
	AggregateInternetTraffic *float32 `json:"AggregateInternetTraffic,omitempty" xml:"AggregateInternetTraffic,omitempty"`
	// The aggregated NAT traffic, in GB.
	//
	// example:
	//
	// 10
	AggregateNatTraffic *float32 `json:"AggregateNatTraffic,omitempty" xml:"AggregateNatTraffic,omitempty"`
	// The aggregated sensitive data detection traffic, in GB.
	//
	// example:
	//
	// 10
	AggregateSdlTraffic *float32 `json:"AggregateSdlTraffic,omitempty" xml:"AggregateSdlTraffic,omitempty"`
	// The aggregated total traffic, in GB.
	//
	// example:
	//
	// 40
	AggregateTotalTraffic *float32 `json:"AggregateTotalTraffic,omitempty" xml:"AggregateTotalTraffic,omitempty"`
	// The aggregated VPC traffic, in GB.
	//
	// example:
	//
	// 10
	AggregateVpcTraffic *float32 `json:"AggregateVpcTraffic,omitempty" xml:"AggregateVpcTraffic,omitempty"`
	// The bill list.
	BillList []*DescribePostpayBillResponseBodyBillList `json:"BillList,omitempty" xml:"BillList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 133173B9-8010-5DF5-8B93-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 132
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePostpayBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayBillResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayBillResponseBody) GetAggregateInternetTraffic() *float32 {
	return s.AggregateInternetTraffic
}

func (s *DescribePostpayBillResponseBody) GetAggregateNatTraffic() *float32 {
	return s.AggregateNatTraffic
}

func (s *DescribePostpayBillResponseBody) GetAggregateSdlTraffic() *float32 {
	return s.AggregateSdlTraffic
}

func (s *DescribePostpayBillResponseBody) GetAggregateTotalTraffic() *float32 {
	return s.AggregateTotalTraffic
}

func (s *DescribePostpayBillResponseBody) GetAggregateVpcTraffic() *float32 {
	return s.AggregateVpcTraffic
}

func (s *DescribePostpayBillResponseBody) GetBillList() []*DescribePostpayBillResponseBodyBillList {
	return s.BillList
}

func (s *DescribePostpayBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePostpayBillResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePostpayBillResponseBody) SetAggregateInternetTraffic(v float32) *DescribePostpayBillResponseBody {
	s.AggregateInternetTraffic = &v
	return s
}

func (s *DescribePostpayBillResponseBody) SetAggregateNatTraffic(v float32) *DescribePostpayBillResponseBody {
	s.AggregateNatTraffic = &v
	return s
}

func (s *DescribePostpayBillResponseBody) SetAggregateSdlTraffic(v float32) *DescribePostpayBillResponseBody {
	s.AggregateSdlTraffic = &v
	return s
}

func (s *DescribePostpayBillResponseBody) SetAggregateTotalTraffic(v float32) *DescribePostpayBillResponseBody {
	s.AggregateTotalTraffic = &v
	return s
}

func (s *DescribePostpayBillResponseBody) SetAggregateVpcTraffic(v float32) *DescribePostpayBillResponseBody {
	s.AggregateVpcTraffic = &v
	return s
}

func (s *DescribePostpayBillResponseBody) SetBillList(v []*DescribePostpayBillResponseBodyBillList) *DescribePostpayBillResponseBody {
	s.BillList = v
	return s
}

func (s *DescribePostpayBillResponseBody) SetRequestId(v string) *DescribePostpayBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayBillResponseBody) SetTotalCount(v int32) *DescribePostpayBillResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePostpayBillResponseBody) Validate() error {
	if s.BillList != nil {
		for _, item := range s.BillList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePostpayBillResponseBodyBillList struct {
	// The end time, expressed as a second-level UNIX timestamp. The value is on the hour or on the day.
	//
	// example:
	//
	// 1733710015
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of Internet instances.
	//
	// example:
	//
	// 3
	InternetInstanceCnt *int64 `json:"InternetInstanceCnt,omitempty" xml:"InternetInstanceCnt,omitempty"`
	// The Internet traffic, in GB.
	//
	// example:
	//
	// 10
	InternetTraffic *float32 `json:"InternetTraffic,omitempty" xml:"InternetTraffic,omitempty"`
	// Indicates whether the bill is deducted. A value of 0 indicates that the bill is not deducted. Any value greater than 0 indicates that the bill is deducted. If the bill is deducted, it is not charged.
	//
	// > This field is meaningful only when you query hourly data.
	//
	// example:
	//
	// 0
	IsDerated *int32 `json:"IsDerated,omitempty" xml:"IsDerated,omitempty"`
	// The log service usage duration, in T × h.
	//
	// example:
	//
	// 24
	LogStorage *int64 `json:"LogStorage,omitempty" xml:"LogStorage,omitempty"`
	// The number of NAT instances.
	//
	// example:
	//
	// 2
	NatInstanceCnt *int64 `json:"NatInstanceCnt,omitempty" xml:"NatInstanceCnt,omitempty"`
	// The NAT traffic, in GB.
	//
	// example:
	//
	// 2
	NatTraffic *float32 `json:"NatTraffic,omitempty" xml:"NatTraffic,omitempty"`
	// The data leak detection usage duration, in hours.
	//
	// example:
	//
	// 0
	Sdl *int64 `json:"Sdl,omitempty" xml:"Sdl,omitempty"`
	// The sensitive data detection traffic, in GB.
	//
	// example:
	//
	// 1
	SdlTraffic *float32 `json:"SdlTraffic,omitempty" xml:"SdlTraffic,omitempty"`
	// The start time, expressed as a second-level UNIX timestamp. The value is on the hour or on the day.
	//
	// example:
	//
	// 1710206070000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The threat intelligence usage duration, in hours.
	//
	// example:
	//
	// 1
	ThreatIntelligence *int64 `json:"ThreatIntelligence,omitempty" xml:"ThreatIntelligence,omitempty"`
	// The number of VPC instances.
	//
	// example:
	//
	// 2
	VpcInstanceCnt *int64 `json:"VpcInstanceCnt,omitempty" xml:"VpcInstanceCnt,omitempty"`
	// The VPC traffic, in GB.
	//
	// example:
	//
	// 4
	VpcTraffic *float32 `json:"VpcTraffic,omitempty" xml:"VpcTraffic,omitempty"`
}

func (s DescribePostpayBillResponseBodyBillList) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayBillResponseBodyBillList) GoString() string {
	return s.String()
}

func (s *DescribePostpayBillResponseBodyBillList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePostpayBillResponseBodyBillList) GetInternetInstanceCnt() *int64 {
	return s.InternetInstanceCnt
}

func (s *DescribePostpayBillResponseBodyBillList) GetInternetTraffic() *float32 {
	return s.InternetTraffic
}

func (s *DescribePostpayBillResponseBodyBillList) GetIsDerated() *int32 {
	return s.IsDerated
}

func (s *DescribePostpayBillResponseBodyBillList) GetLogStorage() *int64 {
	return s.LogStorage
}

func (s *DescribePostpayBillResponseBodyBillList) GetNatInstanceCnt() *int64 {
	return s.NatInstanceCnt
}

func (s *DescribePostpayBillResponseBodyBillList) GetNatTraffic() *float32 {
	return s.NatTraffic
}

func (s *DescribePostpayBillResponseBodyBillList) GetSdl() *int64 {
	return s.Sdl
}

func (s *DescribePostpayBillResponseBodyBillList) GetSdlTraffic() *float32 {
	return s.SdlTraffic
}

func (s *DescribePostpayBillResponseBodyBillList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePostpayBillResponseBodyBillList) GetThreatIntelligence() *int64 {
	return s.ThreatIntelligence
}

func (s *DescribePostpayBillResponseBodyBillList) GetVpcInstanceCnt() *int64 {
	return s.VpcInstanceCnt
}

func (s *DescribePostpayBillResponseBodyBillList) GetVpcTraffic() *float32 {
	return s.VpcTraffic
}

func (s *DescribePostpayBillResponseBodyBillList) SetEndTime(v int64) *DescribePostpayBillResponseBodyBillList {
	s.EndTime = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetInternetInstanceCnt(v int64) *DescribePostpayBillResponseBodyBillList {
	s.InternetInstanceCnt = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetInternetTraffic(v float32) *DescribePostpayBillResponseBodyBillList {
	s.InternetTraffic = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetIsDerated(v int32) *DescribePostpayBillResponseBodyBillList {
	s.IsDerated = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetLogStorage(v int64) *DescribePostpayBillResponseBodyBillList {
	s.LogStorage = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetNatInstanceCnt(v int64) *DescribePostpayBillResponseBodyBillList {
	s.NatInstanceCnt = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetNatTraffic(v float32) *DescribePostpayBillResponseBodyBillList {
	s.NatTraffic = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetSdl(v int64) *DescribePostpayBillResponseBodyBillList {
	s.Sdl = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetSdlTraffic(v float32) *DescribePostpayBillResponseBodyBillList {
	s.SdlTraffic = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetStartTime(v int64) *DescribePostpayBillResponseBodyBillList {
	s.StartTime = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetThreatIntelligence(v int64) *DescribePostpayBillResponseBodyBillList {
	s.ThreatIntelligence = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetVpcInstanceCnt(v int64) *DescribePostpayBillResponseBodyBillList {
	s.VpcInstanceCnt = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) SetVpcTraffic(v float32) *DescribePostpayBillResponseBodyBillList {
	s.VpcTraffic = &v
	return s
}

func (s *DescribePostpayBillResponseBodyBillList) Validate() error {
	return dara.Validate(s)
}

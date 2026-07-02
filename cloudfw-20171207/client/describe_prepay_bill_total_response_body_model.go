// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrepayBillTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillList(v []*DescribePrepayBillTotalResponseBodyBillList) *DescribePrepayBillTotalResponseBody
	GetBillList() []*DescribePrepayBillTotalResponseBodyBillList
	SetRequestId(v string) *DescribePrepayBillTotalResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePrepayBillTotalResponseBody
	GetTotalCount() *int32
}

type DescribePrepayBillTotalResponseBody struct {
	// The bill list, aggregated by day.
	BillList []*DescribePrepayBillTotalResponseBodyBillList `json:"BillList,omitempty" xml:"BillList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 450D47F5-956E-543E-8502-***********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 132
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePrepayBillTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrepayBillTotalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrepayBillTotalResponseBody) GetBillList() []*DescribePrepayBillTotalResponseBodyBillList {
	return s.BillList
}

func (s *DescribePrepayBillTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrepayBillTotalResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePrepayBillTotalResponseBody) SetBillList(v []*DescribePrepayBillTotalResponseBodyBillList) *DescribePrepayBillTotalResponseBody {
	s.BillList = v
	return s
}

func (s *DescribePrepayBillTotalResponseBody) SetRequestId(v string) *DescribePrepayBillTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBody) SetTotalCount(v int32) *DescribePrepayBillTotalResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBody) Validate() error {
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

type DescribePrepayBillTotalResponseBodyBillList struct {
	// The actual billed traffic for sensitive data leak detection.
	//
	// example:
	//
	// 5
	BilledDetectionTraffic *float32 `json:"BilledDetectionTraffic,omitempty" xml:"BilledDetectionTraffic,omitempty"`
	// The sensitive data detection traffic for the day.
	//
	// example:
	//
	// 8
	DailyDetectionTraffic *float32 `json:"DailyDetectionTraffic,omitempty" xml:"DailyDetectionTraffic,omitempty"`
	// The total elastic traffic for the day. Unit: GB.
	//
	// example:
	//
	// 10
	DailyOverflowTraffic *float32 `json:"DailyOverflowTraffic,omitempty" xml:"DailyOverflowTraffic,omitempty"`
	// The default bandwidth of the edition. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	DefaultBandwidth *int64 `json:"DefaultBandwidth,omitempty" xml:"DefaultBandwidth,omitempty"`
	// The elastic bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 200
	ElasticBandwidth *int64 `json:"ElasticBandwidth,omitempty" xml:"ElasticBandwidth,omitempty"`
	// The end time of the day. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1761667200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The extended bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	ExtensionBandwidth *int64 `json:"ExtensionBandwidth,omitempty" xml:"ExtensionBandwidth,omitempty"`
	// The Internet traffic bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 10
	InternetTrafficBandwidth *float32 `json:"InternetTrafficBandwidth,omitempty" xml:"InternetTrafficBandwidth,omitempty"`
	// The monthly free traffic quota for sensitive data detection. Unit: GB.
	//
	// example:
	//
	// 10
	MonthlyRemainingFreeTraffic *float32 `json:"MonthlyRemainingFreeTraffic,omitempty" xml:"MonthlyRemainingFreeTraffic,omitempty"`
	// The NAT traffic bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 10
	NatTrafficBandwidth *float32 `json:"NatTrafficBandwidth,omitempty" xml:"NatTrafficBandwidth,omitempty"`
	// The timestamp when the maximum combined bandwidth (Internet + VPC + NAT) occurred on that day.
	//
	// example:
	//
	// 1761588300
	OverflowTime *int64 `json:"OverflowTime,omitempty" xml:"OverflowTime,omitempty"`
	// The start time of the day. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1761580800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The temporary upgrade bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 0
	TemporaryBandwidth *int64 `json:"TemporaryBandwidth,omitempty" xml:"TemporaryBandwidth,omitempty"`
	// The VPC traffic bandwidth. Unit: Gbit/s.
	//
	// example:
	//
	// 10
	VpcTrafficBandwidth *float32 `json:"VpcTrafficBandwidth,omitempty" xml:"VpcTrafficBandwidth,omitempty"`
}

func (s DescribePrepayBillTotalResponseBodyBillList) String() string {
	return dara.Prettify(s)
}

func (s DescribePrepayBillTotalResponseBodyBillList) GoString() string {
	return s.String()
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetBilledDetectionTraffic() *float32 {
	return s.BilledDetectionTraffic
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetDailyDetectionTraffic() *float32 {
	return s.DailyDetectionTraffic
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetDailyOverflowTraffic() *float32 {
	return s.DailyOverflowTraffic
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetDefaultBandwidth() *int64 {
	return s.DefaultBandwidth
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetElasticBandwidth() *int64 {
	return s.ElasticBandwidth
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetExtensionBandwidth() *int64 {
	return s.ExtensionBandwidth
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetInternetTrafficBandwidth() *float32 {
	return s.InternetTrafficBandwidth
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetMonthlyRemainingFreeTraffic() *float32 {
	return s.MonthlyRemainingFreeTraffic
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetNatTrafficBandwidth() *float32 {
	return s.NatTrafficBandwidth
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetOverflowTime() *int64 {
	return s.OverflowTime
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetTemporaryBandwidth() *int64 {
	return s.TemporaryBandwidth
}

func (s *DescribePrepayBillTotalResponseBodyBillList) GetVpcTrafficBandwidth() *float32 {
	return s.VpcTrafficBandwidth
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetBilledDetectionTraffic(v float32) *DescribePrepayBillTotalResponseBodyBillList {
	s.BilledDetectionTraffic = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetDailyDetectionTraffic(v float32) *DescribePrepayBillTotalResponseBodyBillList {
	s.DailyDetectionTraffic = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetDailyOverflowTraffic(v float32) *DescribePrepayBillTotalResponseBodyBillList {
	s.DailyOverflowTraffic = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetDefaultBandwidth(v int64) *DescribePrepayBillTotalResponseBodyBillList {
	s.DefaultBandwidth = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetElasticBandwidth(v int64) *DescribePrepayBillTotalResponseBodyBillList {
	s.ElasticBandwidth = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetEndTime(v int64) *DescribePrepayBillTotalResponseBodyBillList {
	s.EndTime = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetExtensionBandwidth(v int64) *DescribePrepayBillTotalResponseBodyBillList {
	s.ExtensionBandwidth = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetInternetTrafficBandwidth(v float32) *DescribePrepayBillTotalResponseBodyBillList {
	s.InternetTrafficBandwidth = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetMonthlyRemainingFreeTraffic(v float32) *DescribePrepayBillTotalResponseBodyBillList {
	s.MonthlyRemainingFreeTraffic = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetNatTrafficBandwidth(v float32) *DescribePrepayBillTotalResponseBodyBillList {
	s.NatTrafficBandwidth = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetOverflowTime(v int64) *DescribePrepayBillTotalResponseBodyBillList {
	s.OverflowTime = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetStartTime(v int64) *DescribePrepayBillTotalResponseBodyBillList {
	s.StartTime = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetTemporaryBandwidth(v int64) *DescribePrepayBillTotalResponseBodyBillList {
	s.TemporaryBandwidth = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) SetVpcTrafficBandwidth(v float32) *DescribePrepayBillTotalResponseBodyBillList {
	s.VpcTrafficBandwidth = &v
	return s
}

func (s *DescribePrepayBillTotalResponseBodyBillList) Validate() error {
	return dara.Validate(s)
}

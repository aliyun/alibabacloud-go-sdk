// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrepayDailyBillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBills(v []*DescribePrepayDailyBillsResponseBodyBills) *DescribePrepayDailyBillsResponseBody
	GetBills() []*DescribePrepayDailyBillsResponseBodyBills
	SetRequestId(v string) *DescribePrepayDailyBillsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribePrepayDailyBillsResponseBody
	GetTotalCount() *int64
}

type DescribePrepayDailyBillsResponseBody struct {
	// The bills of the burstable QPS (pay-as-you-go) feature.
	Bills []*DescribePrepayDailyBillsResponseBodyBills `json:"Bills,omitempty" xml:"Bills,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC10C9EA-A367-52D5-***-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePrepayDailyBillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrepayDailyBillsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrepayDailyBillsResponseBody) GetBills() []*DescribePrepayDailyBillsResponseBodyBills {
	return s.Bills
}

func (s *DescribePrepayDailyBillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrepayDailyBillsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePrepayDailyBillsResponseBody) SetBills(v []*DescribePrepayDailyBillsResponseBodyBills) *DescribePrepayDailyBillsResponseBody {
	s.Bills = v
	return s
}

func (s *DescribePrepayDailyBillsResponseBody) SetRequestId(v string) *DescribePrepayDailyBillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBody) SetTotalCount(v int64) *DescribePrepayDailyBillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBody) Validate() error {
	if s.Bills != nil {
		for _, item := range s.Bills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePrepayDailyBillsResponseBodyBills struct {
	// The burstable QPS of the WAF instance.
	//
	// example:
	//
	// 100
	ElasticQpsSetValue *int64 `json:"ElasticQpsSetValue,omitempty" xml:"ElasticQpsSetValue,omitempty"`
	// The billing end time. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1687591200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The status of QPS usage within the current period of time. Valid values:
	//
	// 	- **0**: normal.
	//
	// 	- **1**: excess.
	//
	// 	- **2**: sandbox.
	//
	// example:
	//
	// 0
	ExceedStatus *int32 `json:"ExceedStatus,omitempty" xml:"ExceedStatus,omitempty"`
	// The peak QPS within the current period of time.
	//
	// example:
	//
	// 600
	MaxQps *int64 `json:"MaxQps,omitempty" xml:"MaxQps,omitempty"`
	// The unit price in the bill. The price is measured in CNY for bills at the China site (aliyun.com) and in USD for bills at the international site (alibabacloud.com).
	//
	// example:
	//
	// 0.25
	Price *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	// The extended QPS of the WAF instance.
	//
	// example:
	//
	// 10
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The default QPS of the WAF instance.
	//
	// example:
	//
	// 10
	QpsVersion *int64 `json:"QpsVersion,omitempty" xml:"QpsVersion,omitempty"`
	// Indicates whether risk identification is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	RiskControl *bool `json:"RiskControl,omitempty" xml:"RiskControl,omitempty"`
	// The number of times that risk identification is performed.
	//
	// example:
	//
	// 100
	RiskTraffic *int64 `json:"RiskTraffic,omitempty" xml:"RiskTraffic,omitempty"`
	// The billing start time. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1687822980
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The actual QPS in total.
	//
	// example:
	//
	// 0
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The billing types.
	Type []*string `json:"Type,omitempty" xml:"Type,omitempty" type:"Repeated"`
}

func (s DescribePrepayDailyBillsResponseBodyBills) String() string {
	return dara.Prettify(s)
}

func (s DescribePrepayDailyBillsResponseBodyBills) GoString() string {
	return s.String()
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetElasticQpsSetValue() *int64 {
	return s.ElasticQpsSetValue
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetExceedStatus() *int32 {
	return s.ExceedStatus
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetMaxQps() *int64 {
	return s.MaxQps
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetPrice() *float32 {
	return s.Price
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetQps() *int64 {
	return s.Qps
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetQpsVersion() *int64 {
	return s.QpsVersion
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetRiskControl() *bool {
	return s.RiskControl
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetRiskTraffic() *int64 {
	return s.RiskTraffic
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetTotal() *int64 {
	return s.Total
}

func (s *DescribePrepayDailyBillsResponseBodyBills) GetType() []*string {
	return s.Type
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetElasticQpsSetValue(v int64) *DescribePrepayDailyBillsResponseBodyBills {
	s.ElasticQpsSetValue = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetEndTime(v int64) *DescribePrepayDailyBillsResponseBodyBills {
	s.EndTime = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetExceedStatus(v int32) *DescribePrepayDailyBillsResponseBodyBills {
	s.ExceedStatus = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetMaxQps(v int64) *DescribePrepayDailyBillsResponseBodyBills {
	s.MaxQps = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetPrice(v float32) *DescribePrepayDailyBillsResponseBodyBills {
	s.Price = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetQps(v int64) *DescribePrepayDailyBillsResponseBodyBills {
	s.Qps = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetQpsVersion(v int64) *DescribePrepayDailyBillsResponseBodyBills {
	s.QpsVersion = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetRiskControl(v bool) *DescribePrepayDailyBillsResponseBodyBills {
	s.RiskControl = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetRiskTraffic(v int64) *DescribePrepayDailyBillsResponseBodyBills {
	s.RiskTraffic = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetStartTime(v int64) *DescribePrepayDailyBillsResponseBodyBills {
	s.StartTime = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetTotal(v int64) *DescribePrepayDailyBillsResponseBodyBills {
	s.Total = &v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) SetType(v []*string) *DescribePrepayDailyBillsResponseBodyBills {
	s.Type = v
	return s
}

func (s *DescribePrepayDailyBillsResponseBodyBills) Validate() error {
	return dara.Validate(s)
}

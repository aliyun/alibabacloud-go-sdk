// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserBillPredictionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillPredictionData(v *DescribeLiveUserBillPredictionResponseBodyBillPredictionData) *DescribeLiveUserBillPredictionResponseBody
	GetBillPredictionData() *DescribeLiveUserBillPredictionResponseBodyBillPredictionData
	SetBillType(v string) *DescribeLiveUserBillPredictionResponseBody
	GetBillType() *string
	SetEndTime(v string) *DescribeLiveUserBillPredictionResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeLiveUserBillPredictionResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeLiveUserBillPredictionResponseBody
	GetStartTime() *string
}

type DescribeLiveUserBillPredictionResponseBody struct {
	BillPredictionData *DescribeLiveUserBillPredictionResponseBodyBillPredictionData `json:"BillPredictionData,omitempty" xml:"BillPredictionData,omitempty" type:"Struct"`
	// The billing method. The following billing methods are supported:
	//
	// - hour_flow: Pay-by-traffic on an hourly basis.
	//
	// - day_bandwidth: Pay-by-bandwidth on a daily basis.
	//
	// - month_95: Pay-by-monthly 95th percentile peak bandwidth.
	//
	// - month_avg_day_bandwidth: Pay-by-monthly average of daily peak bandwidth.
	//
	// - month_4th_day_bandwidth: Pay-by-monthly 4th peak bandwidth.
	//
	// - month_avg_day_95: Pay-by-monthly average of daily 95th percentile peak bandwidth.
	//
	// - month_95_night_half: Pay-by-nightly 95th percentile peak bandwidth with a 50% discount.
	//
	// - hour_vas: Pay-for-value-added services on an hourly basis.
	//
	// - day_count: Pay-by-daily request count.
	//
	// example:
	//
	// day_bandwidth
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// The end time of the query. The time is in UTC and follows the ISO 8601 standard.
	//
	// Format: YYYY-MM-DDThh:mm:ssZ. The default value is the current time.
	//
	// example:
	//
	// 2018-10-25T10:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B95BE680-5A6A-1CAD-8AB1-09DFF5D6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. The time is in UTC and follows the ISO 8601 standard.
	//
	// Format: YYYY-MM-DDThh:mm:ssZ. The default value is 00:00 on the first day of the month.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveUserBillPredictionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserBillPredictionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserBillPredictionResponseBody) GetBillPredictionData() *DescribeLiveUserBillPredictionResponseBodyBillPredictionData {
	return s.BillPredictionData
}

func (s *DescribeLiveUserBillPredictionResponseBody) GetBillType() *string {
	return s.BillType
}

func (s *DescribeLiveUserBillPredictionResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveUserBillPredictionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveUserBillPredictionResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveUserBillPredictionResponseBody) SetBillPredictionData(v *DescribeLiveUserBillPredictionResponseBodyBillPredictionData) *DescribeLiveUserBillPredictionResponseBody {
	s.BillPredictionData = v
	return s
}

func (s *DescribeLiveUserBillPredictionResponseBody) SetBillType(v string) *DescribeLiveUserBillPredictionResponseBody {
	s.BillType = &v
	return s
}

func (s *DescribeLiveUserBillPredictionResponseBody) SetEndTime(v string) *DescribeLiveUserBillPredictionResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveUserBillPredictionResponseBody) SetRequestId(v string) *DescribeLiveUserBillPredictionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveUserBillPredictionResponseBody) SetStartTime(v string) *DescribeLiveUserBillPredictionResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveUserBillPredictionResponseBody) Validate() error {
	if s.BillPredictionData != nil {
		if err := s.BillPredictionData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLiveUserBillPredictionResponseBodyBillPredictionData struct {
	BillPredictionDataItem []*DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem `json:"BillPredictionDataItem,omitempty" xml:"BillPredictionDataItem,omitempty" type:"Repeated"`
}

func (s DescribeLiveUserBillPredictionResponseBodyBillPredictionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserBillPredictionResponseBodyBillPredictionData) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionData) GetBillPredictionDataItem() []*DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	return s.BillPredictionDataItem
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionData) SetBillPredictionDataItem(v []*DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) *DescribeLiveUserBillPredictionResponseBodyBillPredictionData {
	s.BillPredictionDataItem = v
	return s
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionData) Validate() error {
	if s.BillPredictionDataItem != nil {
		for _, item := range s.BillPredictionDataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem struct {
	Area    *string  `json:"Area,omitempty" xml:"Area,omitempty"`
	TimeStp *string  `json:"TimeStp,omitempty" xml:"TimeStp,omitempty"`
	Value   *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GetArea() *string {
	return s.Area
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GetTimeStp() *string {
	return s.TimeStp
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GetValue() *float32 {
	return s.Value
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetArea(v string) *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.Area = &v
	return s
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetTimeStp(v string) *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.TimeStp = &v
	return s
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetValue(v float32) *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.Value = &v
	return s
}

func (s *DescribeLiveUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) Validate() error {
	return dara.Validate(s)
}

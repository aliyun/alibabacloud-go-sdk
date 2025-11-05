// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillPredictionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillPredictionData(v *DescribeCdnUserBillPredictionResponseBodyBillPredictionData) *DescribeCdnUserBillPredictionResponseBody
	GetBillPredictionData() *DescribeCdnUserBillPredictionResponseBodyBillPredictionData
	SetBillType(v string) *DescribeCdnUserBillPredictionResponseBody
	GetBillType() *string
	SetEndTime(v string) *DescribeCdnUserBillPredictionResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeCdnUserBillPredictionResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeCdnUserBillPredictionResponseBody
	GetStartTime() *string
}

type DescribeCdnUserBillPredictionResponseBody struct {
	// The estimated bill data.
	BillPredictionData *DescribeCdnUserBillPredictionResponseBodyBillPredictionData `json:"BillPredictionData,omitempty" xml:"BillPredictionData,omitempty" type:"Struct"`
	// The metering method.
	//
	// > If the metering method ends with _overseas, the billable region is outside the Chinese mainland. For example, BillType": "month_avg_day_bandwidth_overseas specifies a billable region outside the Chinese mainland and that the metering method is pay by daily peak bandwidth per month.
	//
	// Valid values:
	//
	// 	- hour_flow: pay by hourly data transfer
	//
	// 	- day_bandwidth: pay by daily bandwidth
	//
	// 	- month_95: pay by monthly 95th percentile bandwidth.
	//
	// 	- month_avg_day_bandwidth: pay by average daily peak bandwidth per month
	//
	// 	- month_4th_day_bandwidth: pay by monthly 4th peak bandwidth
	//
	// 	- month_avg_day_95: pay by average daily 95th percentile bandwidth per month
	//
	// 	- month_95_night_half: pay by 95th percentile bandwidth with 50% off from 00:00 to 08:00.
	//
	// 	- hour_vas: pay by value-added services per hour
	//
	// 	- day_count: pay by daily requests
	//
	// example:
	//
	// month_95
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// The end time of the estimation.
	//
	// example:
	//
	// 2018-10-25T10:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the estimation.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnUserBillPredictionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillPredictionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionResponseBody) GetBillPredictionData() *DescribeCdnUserBillPredictionResponseBodyBillPredictionData {
	return s.BillPredictionData
}

func (s *DescribeCdnUserBillPredictionResponseBody) GetBillType() *string {
	return s.BillType
}

func (s *DescribeCdnUserBillPredictionResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnUserBillPredictionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnUserBillPredictionResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetBillPredictionData(v *DescribeCdnUserBillPredictionResponseBodyBillPredictionData) *DescribeCdnUserBillPredictionResponseBody {
	s.BillPredictionData = v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetBillType(v string) *DescribeCdnUserBillPredictionResponseBody {
	s.BillType = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetEndTime(v string) *DescribeCdnUserBillPredictionResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetRequestId(v string) *DescribeCdnUserBillPredictionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBody) SetStartTime(v string) *DescribeCdnUserBillPredictionResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBody) Validate() error {
	if s.BillPredictionData != nil {
		if err := s.BillPredictionData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnUserBillPredictionResponseBodyBillPredictionData struct {
	BillPredictionDataItem []*DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem `json:"BillPredictionDataItem,omitempty" xml:"BillPredictionDataItem,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserBillPredictionResponseBodyBillPredictionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillPredictionResponseBodyBillPredictionData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionData) GetBillPredictionDataItem() []*DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	return s.BillPredictionDataItem
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionData) SetBillPredictionDataItem(v []*DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) *DescribeCdnUserBillPredictionResponseBodyBillPredictionData {
	s.BillPredictionDataItem = v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionData) Validate() error {
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

type DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem struct {
	// The billable region.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The time when the value used as the estimated value is generated. This parameter is returned only if the metering method is pay by 95th percentile, pay by 95th percentile bandwidth with 50% off from 00:00 to 08:00, or pay by 4th peak bandwidth per month.
	//
	// example:
	//
	// 2018-10-15T16:00:00Z
	TimeStp *string `json:"TimeStp,omitempty" xml:"TimeStp,omitempty"`
	// The estimated value.
	//
	// example:
	//
	// 10000
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GetArea() *string {
	return s.Area
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GetTimeStp() *string {
	return s.TimeStp
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GetValue() *float32 {
	return s.Value
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetArea(v string) *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.Area = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetTimeStp(v string) *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.TimeStp = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetValue(v float32) *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.Value = &v
	return s
}

func (s *DescribeCdnUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) Validate() error {
	return dara.Validate(s)
}

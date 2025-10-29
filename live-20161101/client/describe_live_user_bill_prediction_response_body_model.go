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
	// The estimated bill data.
	BillPredictionData *DescribeLiveUserBillPredictionResponseBodyBillPredictionData `json:"BillPredictionData,omitempty" xml:"BillPredictionData,omitempty" type:"Struct"`
	// The metering method. Valid values:
	//
	// 	- hour_flow: pay by hourly data transfer
	//
	// 	- day_bandwidth: pay by daily bandwidth
	//
	// 	- month_95: pay by monthly 95th percentile bandwidth
	//
	// 	- month_avg_day_bandwidth: pay by average daily peak bandwidth per month
	//
	// 	- month_4th_day_bandwidth: pay by 4th peak bandwidth per month
	//
	// 	- month_avg_day_95: pay by average daily 95th percentile bandwidth per month
	//
	// 	- month_95_night_half: pay by 95th percentile bandwidth with 50% off from 00:00 to 08:00
	//
	// 	- hour_vas: pay by value-added services per hour
	//
	// 	- day_count: pay by daily requests
	//
	// example:
	//
	// day_bandwidth
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// The end time. If you do not specify the request parameter EndTime, the end time is the current time by default. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
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
	// The start time. If you do not specify the request parameter StartTime, the start time is 00:00 on the first day of the month by default. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
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
	// The billable region. Valid values:
	//
	// 	- CN: Chinese mainland
	//
	// 	- OverSeas: countries and regions outside the Chinese mainland
	//
	// 	- AP1: Asia Pacific 1, including Hong Kong (China), Macao (China), Taiwan (China), Japan, and other Southeast Asia countries and regions except Vietnam and Indonesia
	//
	// 	- AP2: Asia Pacific 2, including Indonesia, South Korea, and Vietnam
	//
	// 	- AP3: Asia Pacific 3, including Australia and New Zealand
	//
	// 	- NA: North America, including US and Canada
	//
	// 	- SA: South America, specifically meaning Brazil
	//
	// 	- EU: Europe, including Ukraine, UK, France, Netherlands, Spain, Italy, Sweden, and Germany
	//
	// 	- MEAA: Middle East and Africa, including South Africa, Oman, UAE, and Kuwait
	//
	// By default, data of all regions is aggregated and returned.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The time at which the estimated value occurs. This parameter is returned only if the metering method is pay by 95th percentile bandwidth, pay by 95th percentile bandwidth with 50% off from 00:00 to 08:00, or pay by 4th peak bandwidth per month.
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

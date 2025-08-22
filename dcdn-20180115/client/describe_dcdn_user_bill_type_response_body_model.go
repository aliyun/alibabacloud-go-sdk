// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserBillTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillTypeData(v *DescribeDcdnUserBillTypeResponseBodyBillTypeData) *DescribeDcdnUserBillTypeResponseBody
	GetBillTypeData() *DescribeDcdnUserBillTypeResponseBodyBillTypeData
	SetRequestId(v string) *DescribeDcdnUserBillTypeResponseBody
	GetRequestId() *string
}

type DescribeDcdnUserBillTypeResponseBody struct {
	// The information about the metering method.
	BillTypeData *DescribeDcdnUserBillTypeResponseBodyBillTypeData `json:"BillTypeData,omitempty" xml:"BillTypeData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserBillTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeResponseBody) GetBillTypeData() *DescribeDcdnUserBillTypeResponseBodyBillTypeData {
	return s.BillTypeData
}

func (s *DescribeDcdnUserBillTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserBillTypeResponseBody) SetBillTypeData(v *DescribeDcdnUserBillTypeResponseBodyBillTypeData) *DescribeDcdnUserBillTypeResponseBody {
	s.BillTypeData = v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBody) SetRequestId(v string) *DescribeDcdnUserBillTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserBillTypeResponseBodyBillTypeData struct {
	BillTypeDataItem []*DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem `json:"BillTypeDataItem,omitempty" xml:"BillTypeDataItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserBillTypeResponseBodyBillTypeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillTypeResponseBodyBillTypeData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeData) GetBillTypeDataItem() []*DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	return s.BillTypeDataItem
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeData) SetBillTypeDataItem(v []*DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) *DescribeDcdnUserBillTypeResponseBodyBillTypeData {
	s.BillTypeDataItem = v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem struct {
	// The metering method. Valid values:
	//
	// 	- **hour_flow**: pay by hourly traffic
	//
	// 	- **day_bandwidth**: pay by daily bandwidth
	//
	// 	- **month_95**: pay by monthly 95th percentile
	//
	// 	- **month_avg_day_bandwidth**: pay by average daily peak bandwidth per month
	//
	// 	- **month_4th_day_bandwidth**: pay by 4th peak bandwidth per month
	//
	// 	- **month_avg_day_95**: pay by average daily 95th percentile per month
	//
	// 	- **month_95_night_half**: pay by 95th percentile (50% off during nighttime)
	//
	// 	- **hour_vas**: pay by value-added service per month
	//
	// 	- **quic_hour_count**: pay by QUIC request per hour
	//
	// 	- **hour_count**: pay by request per hour
	//
	// 	- **rtlog_count_day**: pay by the number of real-time logs per day
	//
	// example:
	//
	// month_avg_day_bandwidth_overseas
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// The metering cycle.
	//
	// example:
	//
	// month
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The dimension. Valid values:
	//
	// 	- **flow**: network traffic and bandwidth
	//
	// 	- **vas**: value-added services (HTTPS and requests for dynamic content)
	//
	// 	- **websocket**: WebSocket
	//
	// 	- **quic**: QUIC requests
	//
	// 	- **rtlog2sls**: log entries delivered to Log Service in real time
	//
	// example:
	//
	// flow
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the metering method ends.
	//
	// example:
	//
	// 2018-10-31T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// dcdn
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The time when the metering method takes effect.
	//
	// example:
	//
	// 2018-10-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetBillType() *string {
	return s.BillType
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetProduct() *string {
	return s.Product
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillType(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillType = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillingCycle(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillingCycle = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetDimension(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Dimension = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetEndTime(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetProduct(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Product = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetStartTime(v string) *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) Validate() error {
	return dara.Validate(s)
}

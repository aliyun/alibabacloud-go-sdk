// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillTypeData(v *DescribeCdnUserBillTypeResponseBodyBillTypeData) *DescribeCdnUserBillTypeResponseBody
	GetBillTypeData() *DescribeCdnUserBillTypeResponseBodyBillTypeData
	SetRequestId(v string) *DescribeCdnUserBillTypeResponseBody
	GetRequestId() *string
}

type DescribeCdnUserBillTypeResponseBody struct {
	// Details about the metering methods returned.
	BillTypeData *DescribeCdnUserBillTypeResponseBodyBillTypeData `json:"BillTypeData,omitempty" xml:"BillTypeData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnUserBillTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponseBody) GetBillTypeData() *DescribeCdnUserBillTypeResponseBodyBillTypeData {
	return s.BillTypeData
}

func (s *DescribeCdnUserBillTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnUserBillTypeResponseBody) SetBillTypeData(v *DescribeCdnUserBillTypeResponseBodyBillTypeData) *DescribeCdnUserBillTypeResponseBody {
	s.BillTypeData = v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBody) SetRequestId(v string) *DescribeCdnUserBillTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnUserBillTypeResponseBodyBillTypeData struct {
	BillTypeDataItem []*DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem `json:"BillTypeDataItem,omitempty" xml:"BillTypeDataItem,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeData) GetBillTypeDataItem() []*DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	return s.BillTypeDataItem
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeData) SetBillTypeDataItem(v []*DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) *DescribeCdnUserBillTypeResponseBodyBillTypeData {
	s.BillTypeDataItem = v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeData) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem struct {
	// The metering method.
	//
	// > If the metering method is suffixed with \\*\\*_overseas\\*\\*, the billable region is outside the Chinese mainland. For example, "BillType": "month_avg_day_bandwidth_overseas" indicates that the metering method is pay by average daily peak bandwidth per month in a billable region outside the Chinese mainland.
	//
	// Valid values:
	//
	// 	- hour_flow: pay by hourly data transfer
	//
	// 	- day_bandwidth: pay by daily bandwidth
	//
	// 	- month_95: pay by monthly 95th percentile bandwidth
	//
	// 	- month_avg_day_bandwidth: pay by average daily peak bandwidth per month
	//
	// 	- month_4th_day_bandwidth: pay by monthly 4th peak bandwidth
	//
	// 	- month_avg_day_95: pay by average daily 95th percentile bandwidth per month
	//
	// 	- month_95_night_half: pay by 95th percentile bandwidth with 50% off from 00:00 to 08:00
	//
	// 	- hour_vas: pay by value-added services per hour
	//
	// 	- quic_hour_count: pay by hourly QUIC requests
	//
	// 	- day_count: pay by daily requests
	//
	// 	- hour_count: pay by hourly requests
	//
	// 	- day_95: pay by daily 95th percentile bandwidth
	//
	// example:
	//
	// month_avg_day_bandwidth_overseas
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// The billing cycle.
	//
	// example:
	//
	// month
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The dimension. Valid values:
	//
	// 	- flow: traffic and bandwidth
	//
	// 	- vas: value-added services (HTTPS and requests for dynamic content)
	//
	// 	- quic: the number of QUIC requests
	//
	// 	- websocket: the WebSocket communications protocol
	//
	// 	- rtlog2sls: log entries delivered to Log Service in real time
	//
	// 	- stationflow: traffic over the internal network
	//
	// example:
	//
	// flow
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The time when the metering method ended.
	//
	// example:
	//
	// 2018-10-31T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the product.
	//
	// example:
	//
	// cdn
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The time when the metering method started.
	//
	// example:
	//
	// 2018-10-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetBillType() *string {
	return s.BillType
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetProduct() *string {
	return s.Product
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillType(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillType = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetBillingCycle(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.BillingCycle = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetDimension(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Dimension = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetEndTime(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetProduct(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.Product = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) SetStartTime(v string) *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillTypeResponseBodyBillTypeDataBillTypeDataItem) Validate() error {
	return dara.Validate(s)
}

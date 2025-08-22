// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserBillHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillHistoryData(v *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) *DescribeDcdnUserBillHistoryResponseBody
	GetBillHistoryData() *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData
	SetRequestId(v string) *DescribeDcdnUserBillHistoryResponseBody
	GetRequestId() *string
}

type DescribeDcdnUserBillHistoryResponseBody struct {
	// The billing history returned.
	BillHistoryData *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData `json:"BillHistoryData,omitempty" xml:"BillHistoryData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ED61C6C3-8241-4187-AAA7-5157AE175CEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnUserBillHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBody) GetBillHistoryData() *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData {
	return s.BillHistoryData
}

func (s *DescribeDcdnUserBillHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnUserBillHistoryResponseBody) SetBillHistoryData(v *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) *DescribeDcdnUserBillHistoryResponseBody {
	s.BillHistoryData = v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBody) SetRequestId(v string) *DescribeDcdnUserBillHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserBillHistoryResponseBodyBillHistoryData struct {
	BillHistoryDataItem []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem `json:"BillHistoryDataItem,omitempty" xml:"BillHistoryDataItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) GetBillHistoryDataItem() []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	return s.BillHistoryDataItem
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) SetBillHistoryDataItem(v []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData {
	s.BillHistoryDataItem = v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem struct {
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 2018-09-30T17:00:00Z
	BillTime *string `json:"BillTime,omitempty" xml:"BillTime,omitempty"`
	// The metering method.
	//
	// example:
	//
	// month_4th_day_bandwidth
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// The billable items.
	BillingData *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData `json:"BillingData,omitempty" xml:"BillingData,omitempty" type:"Struct"`
	// The dimension.
	//
	// example:
	//
	// vas
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GetBillTime() *string {
	return s.BillTime
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GetBillType() *string {
	return s.BillType
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GetBillingData() *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData {
	return s.BillingData
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillTime(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillTime = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillType(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillType = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillingData(v *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillingData = v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetDimension(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.Dimension = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData struct {
	BillingDataItem []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem `json:"BillingDataItem,omitempty" xml:"BillingDataItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) GetBillingDataItem() []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	return s.BillingDataItem
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) SetBillingDataItem(v []*DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData {
	s.BillingDataItem = v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem struct {
	// The bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 4839
	Bandwidth *float32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The region for which the billing records are generated. Valid values: **CN**, **OverSeas**, **AP1**, **AP2**, **AP3**, **NA**, **SA**, **EU**, and **MEAA**.
	//
	// example:
	//
	// AP1
	CdnRegion *string `json:"CdnRegion,omitempty" xml:"CdnRegion,omitempty"`
	// The billing method of the disk. Valid values: **StaticHttp**, **DynamicHttp**, and **DynamicHttps**.
	//
	// example:
	//
	// DynamicHttp
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of billing entries.
	//
	// example:
	//
	// 205624
	Count *float32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The amount of network traffic. Unit: bytes.
	//
	// example:
	//
	// 2456
	Flow *float32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetBandwidth() *float32 {
	return s.Bandwidth
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetCdnRegion() *string {
	return s.CdnRegion
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetCount() *float32 {
	return s.Count
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetFlow() *float32 {
	return s.Flow
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetBandwidth(v float32) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Bandwidth = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetCdnRegion(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.CdnRegion = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetChargeType(v string) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.ChargeType = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetCount(v float32) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Count = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetFlow(v float32) *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Flow = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) Validate() error {
	return dara.Validate(s)
}

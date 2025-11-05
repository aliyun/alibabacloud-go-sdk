// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillHistoryData(v *DescribeCdnUserBillHistoryResponseBodyBillHistoryData) *DescribeCdnUserBillHistoryResponseBody
	GetBillHistoryData() *DescribeCdnUserBillHistoryResponseBodyBillHistoryData
	SetRequestId(v string) *DescribeCdnUserBillHistoryResponseBody
	GetRequestId() *string
}

type DescribeCdnUserBillHistoryResponseBody struct {
	// The billing history returned.
	BillHistoryData *DescribeCdnUserBillHistoryResponseBodyBillHistoryData `json:"BillHistoryData,omitempty" xml:"BillHistoryData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ED61C6C3-8241-4187-AAA7-5157AE175CEC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnUserBillHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBody) GetBillHistoryData() *DescribeCdnUserBillHistoryResponseBodyBillHistoryData {
	return s.BillHistoryData
}

func (s *DescribeCdnUserBillHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnUserBillHistoryResponseBody) SetBillHistoryData(v *DescribeCdnUserBillHistoryResponseBodyBillHistoryData) *DescribeCdnUserBillHistoryResponseBody {
	s.BillHistoryData = v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBody) SetRequestId(v string) *DescribeCdnUserBillHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBody) Validate() error {
	if s.BillHistoryData != nil {
		if err := s.BillHistoryData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnUserBillHistoryResponseBodyBillHistoryData struct {
	BillHistoryDataItem []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem `json:"BillHistoryDataItem,omitempty" xml:"BillHistoryDataItem,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryData) GetBillHistoryDataItem() []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	return s.BillHistoryDataItem
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryData) SetBillHistoryDataItem(v []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) *DescribeCdnUserBillHistoryResponseBodyBillHistoryData {
	s.BillHistoryDataItem = v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryData) Validate() error {
	if s.BillHistoryDataItem != nil {
		for _, item := range s.BillHistoryDataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem struct {
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	BillTime *string `json:"BillTime,omitempty" xml:"BillTime,omitempty"`
	// The metering method.
	//
	// example:
	//
	// month_4th_day_bandwidth
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// The billable items.
	BillingData *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData `json:"BillingData,omitempty" xml:"BillingData,omitempty" type:"Struct"`
	// The dimension.
	//
	// example:
	//
	// flow
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GetBillTime() *string {
	return s.BillTime
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GetBillType() *string {
	return s.BillType
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GetBillingData() *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData {
	return s.BillingData
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillTime(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillTime = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillType(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillType = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetBillingData(v *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.BillingData = v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) SetDimension(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem {
	s.Dimension = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItem) Validate() error {
	if s.BillingData != nil {
		if err := s.BillingData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData struct {
	BillingDataItem []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem `json:"BillingDataItem,omitempty" xml:"BillingDataItem,omitempty" type:"Repeated"`
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) GetBillingDataItem() []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	return s.BillingDataItem
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) SetBillingDataItem(v []*DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData {
	s.BillingDataItem = v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingData) Validate() error {
	if s.BillingDataItem != nil {
		for _, item := range s.BillingDataItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem struct {
	// The bandwidth. Unit: bit/s.
	//
	// example:
	//
	// 4041
	Bandwidth *float32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The billable region. Valid values:
	//
	// 	- **CN**: Chinese mainland
	//
	// 	- **OverSeas**: outside the Chinese mainland
	//
	// 	- **AP1**: Asia Pacific 1
	//
	// 	- **AP2**: Asia Pacific 2
	//
	// 	- **AP3**: Asia Pacific 3
	//
	// 	- **NA**: North America
	//
	// 	- **SA**: South America
	//
	// 	- **EU**: Europe
	//
	// 	- **MEAA**: Middle East and Africa
	//
	// example:
	//
	// AP1
	CdnRegion *string `json:"CdnRegion,omitempty" xml:"CdnRegion,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **StaticHttp**: static HTTP requests
	//
	// 	- **DynamicHttp**: dynamic HTTP requests
	//
	// 	- **DynamicHttps**: dynamic HTTPS requests
	//
	// example:
	//
	// DynamicHttp
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// 203601
	Count *float32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The amount of network traffic. Unit: bytes.
	//
	// example:
	//
	// 24567
	Flow *float32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetBandwidth() *float32 {
	return s.Bandwidth
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetCdnRegion() *string {
	return s.CdnRegion
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetCount() *float32 {
	return s.Count
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) GetFlow() *float32 {
	return s.Flow
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetBandwidth(v float32) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Bandwidth = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetCdnRegion(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.CdnRegion = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetChargeType(v string) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.ChargeType = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetCount(v float32) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Count = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) SetFlow(v float32) *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem {
	s.Flow = &v
	return s
}

func (s *DescribeCdnUserBillHistoryResponseBodyBillHistoryDataBillHistoryDataItemBillingDataBillingDataItem) Validate() error {
	return dara.Validate(s)
}

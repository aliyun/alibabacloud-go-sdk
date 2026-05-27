// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumDailyBillsByItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SumDailyBillsByItemResponseBodyData) *SumDailyBillsByItemResponseBody
	GetData() *SumDailyBillsByItemResponseBodyData
	SetHttpCode(v int32) *SumDailyBillsByItemResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *SumDailyBillsByItemResponseBody
	GetRequestId() *string
}

type SumDailyBillsByItemResponseBody struct {
	Data *SumDailyBillsByItemResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0bc3b4ab17217876841756121e1349
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SumDailyBillsByItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SumDailyBillsByItemResponseBody) GoString() string {
	return s.String()
}

func (s *SumDailyBillsByItemResponseBody) GetData() *SumDailyBillsByItemResponseBodyData {
	return s.Data
}

func (s *SumDailyBillsByItemResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SumDailyBillsByItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SumDailyBillsByItemResponseBody) SetData(v *SumDailyBillsByItemResponseBodyData) *SumDailyBillsByItemResponseBody {
	s.Data = v
	return s
}

func (s *SumDailyBillsByItemResponseBody) SetHttpCode(v int32) *SumDailyBillsByItemResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SumDailyBillsByItemResponseBody) SetRequestId(v string) *SumDailyBillsByItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *SumDailyBillsByItemResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SumDailyBillsByItemResponseBodyData struct {
	ItemSummaryBills []*SumDailyBillsByItemResponseBodyDataItemSummaryBills `json:"itemSummaryBills,omitempty" xml:"itemSummaryBills,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 60
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SumDailyBillsByItemResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SumDailyBillsByItemResponseBodyData) GoString() string {
	return s.String()
}

func (s *SumDailyBillsByItemResponseBodyData) GetItemSummaryBills() []*SumDailyBillsByItemResponseBodyDataItemSummaryBills {
	return s.ItemSummaryBills
}

func (s *SumDailyBillsByItemResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SumDailyBillsByItemResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *SumDailyBillsByItemResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SumDailyBillsByItemResponseBodyData) SetItemSummaryBills(v []*SumDailyBillsByItemResponseBodyDataItemSummaryBills) *SumDailyBillsByItemResponseBodyData {
	s.ItemSummaryBills = v
	return s
}

func (s *SumDailyBillsByItemResponseBodyData) SetPageNumber(v int64) *SumDailyBillsByItemResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyData) SetPageSize(v int64) *SumDailyBillsByItemResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyData) SetTotalCount(v int64) *SumDailyBillsByItemResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyData) Validate() error {
	if s.ItemSummaryBills != nil {
		for _, item := range s.ItemSummaryBills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SumDailyBillsByItemResponseBodyDataItemSummaryBills struct {
	// example:
	//
	// RMB
	Currency      *string                                                             `json:"currency,omitempty" xml:"currency,omitempty"`
	DailySumBills []*SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills `json:"dailySumBills,omitempty" xml:"dailySumBills,omitempty" type:"Repeated"`
	// example:
	//
	// DRStorage
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// example:
	//
	// 50
	Percentage *float64 `json:"percentage,omitempty" xml:"percentage,omitempty"`
	// example:
	//
	// OdpsStandard
	SpecCode *string `json:"specCode,omitempty" xml:"specCode,omitempty"`
	// example:
	//
	// 10000
	TotalCost *string `json:"totalCost,omitempty" xml:"totalCost,omitempty"`
}

func (s SumDailyBillsByItemResponseBodyDataItemSummaryBills) String() string {
	return dara.Prettify(s)
}

func (s SumDailyBillsByItemResponseBodyDataItemSummaryBills) GoString() string {
	return s.String()
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) GetCurrency() *string {
	return s.Currency
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) GetDailySumBills() []*SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills {
	return s.DailySumBills
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) GetItemName() *string {
	return s.ItemName
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) GetPercentage() *float64 {
	return s.Percentage
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) GetSpecCode() *string {
	return s.SpecCode
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) GetTotalCost() *string {
	return s.TotalCost
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) SetCurrency(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBills {
	s.Currency = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) SetDailySumBills(v []*SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) *SumDailyBillsByItemResponseBodyDataItemSummaryBills {
	s.DailySumBills = v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) SetItemName(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBills {
	s.ItemName = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) SetPercentage(v float64) *SumDailyBillsByItemResponseBodyDataItemSummaryBills {
	s.Percentage = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) SetSpecCode(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBills {
	s.SpecCode = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) SetTotalCost(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBills {
	s.TotalCost = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBills) Validate() error {
	if s.DailySumBills != nil {
		for _, item := range s.DailySumBills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills struct {
	// example:
	//
	// 31
	Cost *string `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// RMB
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// 20260409
	DateTime  *string                                                                      `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	ItemBills []*SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills `json:"itemBills,omitempty" xml:"itemBills,omitempty" type:"Repeated"`
}

func (s SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) String() string {
	return dara.Prettify(s)
}

func (s SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) GoString() string {
	return s.String()
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) GetCost() *string {
	return s.Cost
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) GetCurrency() *string {
	return s.Currency
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) GetDateTime() *string {
	return s.DateTime
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) GetItemBills() []*SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills {
	return s.ItemBills
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) SetCost(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills {
	s.Cost = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) SetCurrency(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills {
	s.Currency = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) SetDateTime(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills {
	s.DateTime = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) SetItemBills(v []*SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills {
	s.ItemBills = v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBills) Validate() error {
	if s.ItemBills != nil {
		for _, item := range s.ItemBills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills struct {
	// example:
	//
	// 433
	Cost *string `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// RMB
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// empty
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// example:
	//
	// 60
	Percentage *float64 `json:"percentage,omitempty" xml:"percentage,omitempty"`
}

func (s SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) String() string {
	return dara.Prettify(s)
}

func (s SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) GoString() string {
	return s.String()
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) GetCost() *string {
	return s.Cost
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) GetCurrency() *string {
	return s.Currency
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) GetItemName() *string {
	return s.ItemName
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) GetPercentage() *float64 {
	return s.Percentage
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) SetCost(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills {
	s.Cost = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) SetCurrency(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills {
	s.Currency = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) SetItemName(v string) *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills {
	s.ItemName = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) SetPercentage(v float64) *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills {
	s.Percentage = &v
	return s
}

func (s *SumDailyBillsByItemResponseBodyDataItemSummaryBillsDailySumBillsItemBills) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumBillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SumBillsResponseBodyData) *SumBillsResponseBody
	GetData() *SumBillsResponseBodyData
	SetHttpCode(v int32) *SumBillsResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *SumBillsResponseBody
	GetRequestId() *string
}

type SumBillsResponseBody struct {
	Data *SumBillsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// BB66A390-4EF7-557E-9489-7F98D6F44002
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SumBillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SumBillsResponseBody) GoString() string {
	return s.String()
}

func (s *SumBillsResponseBody) GetData() *SumBillsResponseBodyData {
	return s.Data
}

func (s *SumBillsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SumBillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SumBillsResponseBody) SetData(v *SumBillsResponseBodyData) *SumBillsResponseBody {
	s.Data = v
	return s
}

func (s *SumBillsResponseBody) SetHttpCode(v int32) *SumBillsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SumBillsResponseBody) SetRequestId(v string) *SumBillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SumBillsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SumBillsResponseBodyData struct {
	// example:
	//
	// CNY
	Currency  *string                              `json:"currency,omitempty" xml:"currency,omitempty"`
	ItemBills []*SumBillsResponseBodyDataItemBills `json:"itemBills,omitempty" xml:"itemBills,omitempty" type:"Repeated"`
	// example:
	//
	// 123.56
	TotalCost *string `json:"totalCost,omitempty" xml:"totalCost,omitempty"`
}

func (s SumBillsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SumBillsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SumBillsResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *SumBillsResponseBodyData) GetItemBills() []*SumBillsResponseBodyDataItemBills {
	return s.ItemBills
}

func (s *SumBillsResponseBodyData) GetTotalCost() *string {
	return s.TotalCost
}

func (s *SumBillsResponseBodyData) SetCurrency(v string) *SumBillsResponseBodyData {
	s.Currency = &v
	return s
}

func (s *SumBillsResponseBodyData) SetItemBills(v []*SumBillsResponseBodyDataItemBills) *SumBillsResponseBodyData {
	s.ItemBills = v
	return s
}

func (s *SumBillsResponseBodyData) SetTotalCost(v string) *SumBillsResponseBodyData {
	s.TotalCost = &v
	return s
}

func (s *SumBillsResponseBodyData) Validate() error {
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

type SumBillsResponseBodyDataItemBills struct {
	// example:
	//
	// 123.56
	Cost *string `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// projectName
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// example:
	//
	// 56.12
	Percentage *float64 `json:"percentage,omitempty" xml:"percentage,omitempty"`
}

func (s SumBillsResponseBodyDataItemBills) String() string {
	return dara.Prettify(s)
}

func (s SumBillsResponseBodyDataItemBills) GoString() string {
	return s.String()
}

func (s *SumBillsResponseBodyDataItemBills) GetCost() *string {
	return s.Cost
}

func (s *SumBillsResponseBodyDataItemBills) GetCurrency() *string {
	return s.Currency
}

func (s *SumBillsResponseBodyDataItemBills) GetItemName() *string {
	return s.ItemName
}

func (s *SumBillsResponseBodyDataItemBills) GetPercentage() *float64 {
	return s.Percentage
}

func (s *SumBillsResponseBodyDataItemBills) SetCost(v string) *SumBillsResponseBodyDataItemBills {
	s.Cost = &v
	return s
}

func (s *SumBillsResponseBodyDataItemBills) SetCurrency(v string) *SumBillsResponseBodyDataItemBills {
	s.Currency = &v
	return s
}

func (s *SumBillsResponseBodyDataItemBills) SetItemName(v string) *SumBillsResponseBodyDataItemBills {
	s.ItemName = &v
	return s
}

func (s *SumBillsResponseBodyDataItemBills) SetPercentage(v float64) *SumBillsResponseBodyDataItemBills {
	s.Percentage = &v
	return s
}

func (s *SumBillsResponseBodyDataItemBills) Validate() error {
	return dara.Validate(s)
}

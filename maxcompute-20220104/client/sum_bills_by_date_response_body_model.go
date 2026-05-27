// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumBillsByDateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SumBillsByDateResponseBodyData) *SumBillsByDateResponseBody
	GetData() []*SumBillsByDateResponseBodyData
	SetHttpCode(v int32) *SumBillsByDateResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *SumBillsByDateResponseBody
	GetRequestId() *string
}

type SumBillsByDateResponseBody struct {
	Data []*SumBillsByDateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// BB66A390-4EF7-557E-9489-7F98D6F44002
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SumBillsByDateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SumBillsByDateResponseBody) GoString() string {
	return s.String()
}

func (s *SumBillsByDateResponseBody) GetData() []*SumBillsByDateResponseBodyData {
	return s.Data
}

func (s *SumBillsByDateResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SumBillsByDateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SumBillsByDateResponseBody) SetData(v []*SumBillsByDateResponseBodyData) *SumBillsByDateResponseBody {
	s.Data = v
	return s
}

func (s *SumBillsByDateResponseBody) SetHttpCode(v int32) *SumBillsByDateResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SumBillsByDateResponseBody) SetRequestId(v string) *SumBillsByDateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SumBillsByDateResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SumBillsByDateResponseBodyData struct {
	// example:
	//
	// 2000
	Cost *string `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// RMB
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// 20250719
	DateTime  *string                                    `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	ItemBills []*SumBillsByDateResponseBodyDataItemBills `json:"itemBills,omitempty" xml:"itemBills,omitempty" type:"Repeated"`
}

func (s SumBillsByDateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SumBillsByDateResponseBodyData) GoString() string {
	return s.String()
}

func (s *SumBillsByDateResponseBodyData) GetCost() *string {
	return s.Cost
}

func (s *SumBillsByDateResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *SumBillsByDateResponseBodyData) GetDateTime() *string {
	return s.DateTime
}

func (s *SumBillsByDateResponseBodyData) GetItemBills() []*SumBillsByDateResponseBodyDataItemBills {
	return s.ItemBills
}

func (s *SumBillsByDateResponseBodyData) SetCost(v string) *SumBillsByDateResponseBodyData {
	s.Cost = &v
	return s
}

func (s *SumBillsByDateResponseBodyData) SetCurrency(v string) *SumBillsByDateResponseBodyData {
	s.Currency = &v
	return s
}

func (s *SumBillsByDateResponseBodyData) SetDateTime(v string) *SumBillsByDateResponseBodyData {
	s.DateTime = &v
	return s
}

func (s *SumBillsByDateResponseBodyData) SetItemBills(v []*SumBillsByDateResponseBodyDataItemBills) *SumBillsByDateResponseBodyData {
	s.ItemBills = v
	return s
}

func (s *SumBillsByDateResponseBodyData) Validate() error {
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

type SumBillsByDateResponseBodyDataItemBills struct {
	// example:
	//
	// 1000
	Cost *string `json:"cost,omitempty" xml:"cost,omitempty"`
	// example:
	//
	// RMB
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// example:
	//
	// projectName
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// example:
	//
	// 50
	Percentage *float64 `json:"percentage,omitempty" xml:"percentage,omitempty"`
}

func (s SumBillsByDateResponseBodyDataItemBills) String() string {
	return dara.Prettify(s)
}

func (s SumBillsByDateResponseBodyDataItemBills) GoString() string {
	return s.String()
}

func (s *SumBillsByDateResponseBodyDataItemBills) GetCost() *string {
	return s.Cost
}

func (s *SumBillsByDateResponseBodyDataItemBills) GetCurrency() *string {
	return s.Currency
}

func (s *SumBillsByDateResponseBodyDataItemBills) GetItemName() *string {
	return s.ItemName
}

func (s *SumBillsByDateResponseBodyDataItemBills) GetPercentage() *float64 {
	return s.Percentage
}

func (s *SumBillsByDateResponseBodyDataItemBills) SetCost(v string) *SumBillsByDateResponseBodyDataItemBills {
	s.Cost = &v
	return s
}

func (s *SumBillsByDateResponseBodyDataItemBills) SetCurrency(v string) *SumBillsByDateResponseBodyDataItemBills {
	s.Currency = &v
	return s
}

func (s *SumBillsByDateResponseBodyDataItemBills) SetItemName(v string) *SumBillsByDateResponseBodyDataItemBills {
	s.ItemName = &v
	return s
}

func (s *SumBillsByDateResponseBodyDataItemBills) SetPercentage(v float64) *SumBillsByDateResponseBodyDataItemBills {
	s.Percentage = &v
	return s
}

func (s *SumBillsByDateResponseBodyDataItemBills) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpInventorySummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetGwpInventorySummaryResponseBodyData) *GetGwpInventorySummaryResponseBody
	GetData() *GetGwpInventorySummaryResponseBodyData
	SetRequestId(v string) *GetGwpInventorySummaryResponseBody
	GetRequestId() *string
}

type GetGwpInventorySummaryResponseBody struct {
	// The returned results.
	Data *GetGwpInventorySummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGwpInventorySummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGwpInventorySummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryResponseBody) GetData() *GetGwpInventorySummaryResponseBodyData {
	return s.Data
}

func (s *GetGwpInventorySummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGwpInventorySummaryResponseBody) SetData(v *GetGwpInventorySummaryResponseBodyData) *GetGwpInventorySummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetGwpInventorySummaryResponseBody) SetRequestId(v string) *GetGwpInventorySummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGwpInventorySummaryResponseBodyData struct {
	// Top 4 types of carbon footprint contribution.
	Items []*GetGwpInventorySummaryResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The emission quantity.
	//
	// example:
	//
	// 1.0100
	Quantity *float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The time when the result was generated, in the millisecond timestamp format.
	//
	// example:
	//
	// 1709108026000
	ResultGenerateTime *int64 `json:"resultGenerateTime,omitempty" xml:"resultGenerateTime,omitempty"`
	// Emission Unit.
	//
	// example:
	//
	// tCO₂e/Piece(s)
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpInventorySummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGwpInventorySummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryResponseBodyData) GetItems() []*GetGwpInventorySummaryResponseBodyDataItems {
	return s.Items
}

func (s *GetGwpInventorySummaryResponseBodyData) GetQuantity() *float64 {
	return s.Quantity
}

func (s *GetGwpInventorySummaryResponseBodyData) GetResultGenerateTime() *int64 {
	return s.ResultGenerateTime
}

func (s *GetGwpInventorySummaryResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *GetGwpInventorySummaryResponseBodyData) SetItems(v []*GetGwpInventorySummaryResponseBodyDataItems) *GetGwpInventorySummaryResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyData) SetQuantity(v float64) *GetGwpInventorySummaryResponseBodyData {
	s.Quantity = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyData) SetResultGenerateTime(v int64) *GetGwpInventorySummaryResponseBodyData {
	s.ResultGenerateTime = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyData) SetUnit(v string) *GetGwpInventorySummaryResponseBodyData {
	s.Unit = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGwpInventorySummaryResponseBodyDataItems struct {
	// Inventory resource type name.
	//
	// example:
	//
	// Energy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Percentage.
	//
	// example:
	//
	// 99.01
	Percent *string `json:"percent,omitempty" xml:"percent,omitempty"`
	// Quantity.
	//
	// example:
	//
	// 9.9763
	Quantity *float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// The unit.
	//
	// example:
	//
	// kgCO₂e/Piece(s)
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpInventorySummaryResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetGwpInventorySummaryResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) GetPercent() *string {
	return s.Percent
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) GetQuantity() *float64 {
	return s.Quantity
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) GetUnit() *string {
	return s.Unit
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) SetName(v string) *GetGwpInventorySummaryResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) SetPercent(v string) *GetGwpInventorySummaryResponseBodyDataItems {
	s.Percent = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) SetQuantity(v float64) *GetGwpInventorySummaryResponseBodyDataItems {
	s.Quantity = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) SetUnit(v string) *GetGwpInventorySummaryResponseBodyDataItems {
	s.Unit = &v
	return s
}

func (s *GetGwpInventorySummaryResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}

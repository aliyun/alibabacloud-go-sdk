// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpBenchmarkSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetGwpBenchmarkSummaryResponseBodyData) *GetGwpBenchmarkSummaryResponseBody
	GetData() *GetGwpBenchmarkSummaryResponseBodyData
	SetRequestId(v string) *GetGwpBenchmarkSummaryResponseBody
	GetRequestId() *string
}

type GetGwpBenchmarkSummaryResponseBody struct {
	// The response parameters.
	Data *GetGwpBenchmarkSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGwpBenchmarkSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryResponseBody) GetData() *GetGwpBenchmarkSummaryResponseBodyData {
	return s.Data
}

func (s *GetGwpBenchmarkSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGwpBenchmarkSummaryResponseBody) SetData(v *GetGwpBenchmarkSummaryResponseBodyData) *GetGwpBenchmarkSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBody) SetRequestId(v string) *GetGwpBenchmarkSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGwpBenchmarkSummaryResponseBodyData struct {
	// Carbon Reduction Contribution Top4 Details.
	Items []*GetGwpBenchmarkSummaryResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Emission amount is presented with four decimal places. Normally, modeling doesn\\"t result in negative values, but users can represent carbon reductions as negatives. The amount, paired with the unit, defines the emissions. Both are dynamically adjusted. If emissions exceed `1000 kgCO₂e/productUnit`, they convert to `tCO₂e/productUnit`. If they fall below `1 kgCO₂e/productUnit`, they convert to `gCO₂e/productUnit`. Otherwise, they stay in `kgCO₂e/productUnit`.
	//
	// example:
	//
	// 1000.0000
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Unit of emissions. The default value is `kgCO₂e/productUnit.` `productUnit` is the unit selected for the product. The unit value is changed to `tCO₂e/productUnit` or `gCO₂e/productUnit`. For more information, see the remarks in the quantity column.
	//
	// example:
	//
	// kgCO₂e/t
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpBenchmarkSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) GetItems() []*GetGwpBenchmarkSummaryResponseBodyDataItems {
	return s.Items
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) GetQuantity() *int64 {
	return s.Quantity
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) SetItems(v []*GetGwpBenchmarkSummaryResponseBodyDataItems) *GetGwpBenchmarkSummaryResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) SetQuantity(v int64) *GetGwpBenchmarkSummaryResponseBodyData {
	s.Quantity = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) SetUnit(v string) *GetGwpBenchmarkSummaryResponseBodyData {
	s.Unit = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyData) Validate() error {
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

type GetGwpBenchmarkSummaryResponseBodyDataItems struct {
	// Name of carbon reduction details.
	//
	// example:
	//
	// Energy-Replacement
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Percentage of emissions. The value is of the string type. Two decimal places are reserved for numbers. For example, "99.01" indicates the 99.01% of this type of emissions to the total emissions. Note that the returned string itself does not contain a percent sign.
	//
	// example:
	//
	// 99.01
	Percent *string `json:"percent,omitempty" xml:"percent,omitempty"`
	// Emission amount is presented with four decimal places. Normally, modeling doesn\\"t result in negative values, but users can represent carbon reductions as negatives. The amount, paired with the unit, defines the emissions. Both are dynamically adjusted. If emissions exceed `1000 kgCO₂e/productUnit`, they convert to `tCO₂e/productUnit`. If they fall below `1 kgCO₂e/productUnit`, they convert to `gCO₂e/productUnit`. Otherwise, they stay in `kgCO₂e/productUnit`.
	//
	// example:
	//
	// 9.9763
	Quantity *int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Unit of emissions. The default value is `kgCO₂e/productUnit.` `productUnit` is the unit selected for the product. The unit value is changed to `tCO₂e/productUnit` or `gCO₂e/productUnit`. For more information, see the remarks in the quantity column.
	//
	// example:
	//
	// kgCO₂e/kg
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpBenchmarkSummaryResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkSummaryResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) GetPercent() *string {
	return s.Percent
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) GetQuantity() *int64 {
	return s.Quantity
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) GetUnit() *string {
	return s.Unit
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) SetName(v string) *GetGwpBenchmarkSummaryResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) SetPercent(v string) *GetGwpBenchmarkSummaryResponseBodyDataItems {
	s.Percent = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) SetQuantity(v int64) *GetGwpBenchmarkSummaryResponseBodyDataItems {
	s.Quantity = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) SetUnit(v string) *GetGwpBenchmarkSummaryResponseBodyDataItems {
	s.Unit = &v
	return s
}

func (s *GetGwpBenchmarkSummaryResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}

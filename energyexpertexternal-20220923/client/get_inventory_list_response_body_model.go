// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInventoryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetInventoryListResponseBodyData) *GetInventoryListResponseBody
	GetData() *GetInventoryListResponseBodyData
	SetRequestId(v string) *GetInventoryListResponseBody
	GetRequestId() *string
}

type GetInventoryListResponseBody struct {
	// The response parameters.
	Data *GetInventoryListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetInventoryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInventoryListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInventoryListResponseBody) GetData() *GetInventoryListResponseBodyData {
	return s.Data
}

func (s *GetInventoryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInventoryListResponseBody) SetData(v *GetInventoryListResponseBodyData) *GetInventoryListResponseBody {
	s.Data = v
	return s
}

func (s *GetInventoryListResponseBody) SetRequestId(v string) *GetInventoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInventoryListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInventoryListResponseBodyData struct {
	// Inventory detail.
	Items []*GetInventoryListResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Unit of product.
	//
	// example:
	//
	// kg
	ProductUnit *string `json:"productUnit,omitempty" xml:"productUnit,omitempty"`
	// Emission Unit: The default value is kgCO₂ /productUnit. productUnit is the unit selected for the product. The unit value is changed to tCO₂ e/productUnit or gCO₂ e/productUnit based on the emission quantity. For more information, see the quantity column.
	//
	// example:
	//
	// kgCO₂e/kg
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetInventoryListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInventoryListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInventoryListResponseBodyData) GetItems() []*GetInventoryListResponseBodyDataItems {
	return s.Items
}

func (s *GetInventoryListResponseBodyData) GetProductUnit() *string {
	return s.ProductUnit
}

func (s *GetInventoryListResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *GetInventoryListResponseBodyData) SetItems(v []*GetInventoryListResponseBodyDataItems) *GetInventoryListResponseBodyData {
	s.Items = v
	return s
}

func (s *GetInventoryListResponseBodyData) SetProductUnit(v string) *GetInventoryListResponseBodyData {
	s.ProductUnit = &v
	return s
}

func (s *GetInventoryListResponseBodyData) SetUnit(v string) *GetInventoryListResponseBodyData {
	s.Unit = &v
	return s
}

func (s *GetInventoryListResponseBodyData) Validate() error {
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

type GetInventoryListResponseBodyDataItems struct {
	// Emission quantity: may be positive, negative, or 0. To ensure the calculation accuracy, 24 decimal places are reserved for the calculation process. We recommend that you intercept data based on your business requirements.
	//
	// example:
	//
	// 1000.000000000000000000000000000000
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	// Name
	//
	// > The name is related to the request parameters group. Request parameters: resource, return name parameter meaning: list name; request parameters: process, return name parameter meaning: process name; request parameters: resourceType, return name parameter meaning: inventory resource type name; request parameters: processType, return name parameter meaning: flow name.
	//
	// example:
	//
	// Energy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Percentage
	//
	// example:
	//
	// 99.01
	Percent *string `json:"percent,omitempty" xml:"percent,omitempty"`
	// Process Name: It is only meaningful when the request parameters group is resource.
	//
	// example:
	//
	// Process-1
	ProcessName *string `json:"processName,omitempty" xml:"processName,omitempty"`
}

func (s GetInventoryListResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetInventoryListResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetInventoryListResponseBodyDataItems) GetCarbonEmission() *float64 {
	return s.CarbonEmission
}

func (s *GetInventoryListResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *GetInventoryListResponseBodyDataItems) GetPercent() *string {
	return s.Percent
}

func (s *GetInventoryListResponseBodyDataItems) GetProcessName() *string {
	return s.ProcessName
}

func (s *GetInventoryListResponseBodyDataItems) SetCarbonEmission(v float64) *GetInventoryListResponseBodyDataItems {
	s.CarbonEmission = &v
	return s
}

func (s *GetInventoryListResponseBodyDataItems) SetName(v string) *GetInventoryListResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *GetInventoryListResponseBodyDataItems) SetPercent(v string) *GetInventoryListResponseBodyDataItems {
	s.Percent = &v
	return s
}

func (s *GetInventoryListResponseBodyDataItems) SetProcessName(v string) *GetInventoryListResponseBodyDataItems {
	s.ProcessName = &v
	return s
}

func (s *GetInventoryListResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}

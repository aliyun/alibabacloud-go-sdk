// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpInventoryConstituteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetGwpInventoryConstituteResponseBodyData) *GetGwpInventoryConstituteResponseBody
	GetData() *GetGwpInventoryConstituteResponseBodyData
	SetRequestId(v string) *GetGwpInventoryConstituteResponseBody
	GetRequestId() *string
}

type GetGwpInventoryConstituteResponseBody struct {
	// The response parameters.
	Data *GetGwpInventoryConstituteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 06DA2909-7736-5738-AA31-ACD617D828F1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGwpInventoryConstituteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGwpInventoryConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetGwpInventoryConstituteResponseBody) GetData() *GetGwpInventoryConstituteResponseBodyData {
	return s.Data
}

func (s *GetGwpInventoryConstituteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGwpInventoryConstituteResponseBody) SetData(v *GetGwpInventoryConstituteResponseBodyData) *GetGwpInventoryConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetGwpInventoryConstituteResponseBody) SetRequestId(v string) *GetGwpInventoryConstituteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGwpInventoryConstituteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGwpInventoryConstituteResponseBodyData struct {
	// Aggregated by resource type of an inventory.
	ByResourceType []*GwpInventoryConstitute `json:"byResourceType,omitempty" xml:"byResourceType,omitempty" type:"Repeated"`
	// Emission quantity: may be positive, negative, or 0. To ensure the calculation accuracy, 24 decimal places are reserved for the calculation process. We recommend that you intercept data based on your business requirements.
	//
	// example:
	//
	// 1009.976265540000000000000000000000
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	// Organized by hierarchy from high to low, according to the flow-> process-> inventory hierarchy.
	Items []*GwpInventoryConstitute `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// This is not used for displaying
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Emission Unit.
	//
	// example:
	//
	// kgCO₂e/t
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpInventoryConstituteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGwpInventoryConstituteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGwpInventoryConstituteResponseBodyData) GetByResourceType() []*GwpInventoryConstitute {
	return s.ByResourceType
}

func (s *GetGwpInventoryConstituteResponseBodyData) GetCarbonEmission() *float64 {
	return s.CarbonEmission
}

func (s *GetGwpInventoryConstituteResponseBodyData) GetItems() []*GwpInventoryConstitute {
	return s.Items
}

func (s *GetGwpInventoryConstituteResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGwpInventoryConstituteResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetByResourceType(v []*GwpInventoryConstitute) *GetGwpInventoryConstituteResponseBodyData {
	s.ByResourceType = v
	return s
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetCarbonEmission(v float64) *GetGwpInventoryConstituteResponseBodyData {
	s.CarbonEmission = &v
	return s
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetItems(v []*GwpInventoryConstitute) *GetGwpInventoryConstituteResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetName(v string) *GetGwpInventoryConstituteResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGwpInventoryConstituteResponseBodyData) SetUnit(v string) *GetGwpInventoryConstituteResponseBodyData {
	s.Unit = &v
	return s
}

func (s *GetGwpInventoryConstituteResponseBodyData) Validate() error {
	if s.ByResourceType != nil {
		for _, item := range s.ByResourceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

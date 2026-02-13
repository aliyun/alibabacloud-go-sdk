// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGwpInventoryConstitute interface {
	dara.Model
	String() string
	GoString() string
	SetByResourceType(v []*GwpResourceConstitute) *GwpInventoryConstitute
	GetByResourceType() []*GwpResourceConstitute
	SetCarbonEmission(v float64) *GwpInventoryConstitute
	GetCarbonEmission() *float64
	SetItems(v []*GwpInventoryConstitute) *GwpInventoryConstitute
	GetItems() []*GwpInventoryConstitute
	SetName(v string) *GwpInventoryConstitute
	GetName() *string
	SetPercent(v float64) *GwpInventoryConstitute
	GetPercent() *float64
	SetResourceType(v int32) *GwpInventoryConstitute
	GetResourceType() *int32
	SetUnit(v string) *GwpInventoryConstitute
	GetUnit() *string
}

type GwpInventoryConstitute struct {
	// Grouped by inventory resource type, this data reveals emissions details for each category. It serves the "By type" analysis in the composition breakdown. A nested structure is employed: total carbon emissions are organized first by inventory type, forming a two-level hierarchy, with the innermost level \\"byResource\\" currently empty.
	ByResourceType []*GwpResourceConstitute `json:"byResourceType,omitempty" xml:"byResourceType,omitempty" type:"Repeated"`
	// Emission quantity
	//
	// example:
	//
	// 1009.976265540000000000000000000000
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	// Organized hierarchically, it cascades from high to low: flow-> process-> inventory level. Employed for "By inventory" analysis in compositional breakdowns, the innermost layer of this nested structure is empty.
	Items []*GwpInventoryConstitute `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Name
	//
	// example:
	//
	// Acquisition of Raw Materials
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Percentage of emissions, for example 100.00 means 100.00%.
	//
	// example:
	//
	// 100.00
	Percent *float64 `json:"percent,omitempty" xml:"percent,omitempty"`
	// Resouce type of inventory.
	//
	// example:
	//
	// 1
	ResourceType *int32 `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// Unit
	//
	// example:
	//
	// kgCO₂e/kg
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GwpInventoryConstitute) String() string {
	return dara.Prettify(s)
}

func (s GwpInventoryConstitute) GoString() string {
	return s.String()
}

func (s *GwpInventoryConstitute) GetByResourceType() []*GwpResourceConstitute {
	return s.ByResourceType
}

func (s *GwpInventoryConstitute) GetCarbonEmission() *float64 {
	return s.CarbonEmission
}

func (s *GwpInventoryConstitute) GetItems() []*GwpInventoryConstitute {
	return s.Items
}

func (s *GwpInventoryConstitute) GetName() *string {
	return s.Name
}

func (s *GwpInventoryConstitute) GetPercent() *float64 {
	return s.Percent
}

func (s *GwpInventoryConstitute) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *GwpInventoryConstitute) GetUnit() *string {
	return s.Unit
}

func (s *GwpInventoryConstitute) SetByResourceType(v []*GwpResourceConstitute) *GwpInventoryConstitute {
	s.ByResourceType = v
	return s
}

func (s *GwpInventoryConstitute) SetCarbonEmission(v float64) *GwpInventoryConstitute {
	s.CarbonEmission = &v
	return s
}

func (s *GwpInventoryConstitute) SetItems(v []*GwpInventoryConstitute) *GwpInventoryConstitute {
	s.Items = v
	return s
}

func (s *GwpInventoryConstitute) SetName(v string) *GwpInventoryConstitute {
	s.Name = &v
	return s
}

func (s *GwpInventoryConstitute) SetPercent(v float64) *GwpInventoryConstitute {
	s.Percent = &v
	return s
}

func (s *GwpInventoryConstitute) SetResourceType(v int32) *GwpInventoryConstitute {
	s.ResourceType = &v
	return s
}

func (s *GwpInventoryConstitute) SetUnit(v string) *GwpInventoryConstitute {
	s.Unit = &v
	return s
}

func (s *GwpInventoryConstitute) Validate() error {
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

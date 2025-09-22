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
	ByResourceType []*GwpResourceConstitute  `json:"byResourceType,omitempty" xml:"byResourceType,omitempty" type:"Repeated"`
	CarbonEmission *float64                  `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	Items          []*GwpInventoryConstitute `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Name           *string                   `json:"name,omitempty" xml:"name,omitempty"`
	Percent        *float64                  `json:"percent,omitempty" xml:"percent,omitempty"`
	ResourceType   *int32                    `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Unit           *string                   `json:"unit,omitempty" xml:"unit,omitempty"`
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
	return dara.Validate(s)
}

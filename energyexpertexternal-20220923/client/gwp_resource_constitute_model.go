// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGwpResourceConstitute interface {
	dara.Model
	String() string
	GoString() string
	SetCarbonEmission(v float64) *GwpResourceConstitute
	GetCarbonEmission() *float64
	SetName(v string) *GwpResourceConstitute
	GetName() *string
	SetPercent(v string) *GwpResourceConstitute
	GetPercent() *string
	SetResourceType(v int32) *GwpResourceConstitute
	GetResourceType() *int32
	SetUnit(v string) *GwpResourceConstitute
	GetUnit() *string
}

type GwpResourceConstitute struct {
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	Name           *string  `json:"name,omitempty" xml:"name,omitempty"`
	Percent        *string  `json:"percent,omitempty" xml:"percent,omitempty"`
	ResourceType   *int32   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	Unit           *string  `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GwpResourceConstitute) String() string {
	return dara.Prettify(s)
}

func (s GwpResourceConstitute) GoString() string {
	return s.String()
}

func (s *GwpResourceConstitute) GetCarbonEmission() *float64 {
	return s.CarbonEmission
}

func (s *GwpResourceConstitute) GetName() *string {
	return s.Name
}

func (s *GwpResourceConstitute) GetPercent() *string {
	return s.Percent
}

func (s *GwpResourceConstitute) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *GwpResourceConstitute) GetUnit() *string {
	return s.Unit
}

func (s *GwpResourceConstitute) SetCarbonEmission(v float64) *GwpResourceConstitute {
	s.CarbonEmission = &v
	return s
}

func (s *GwpResourceConstitute) SetName(v string) *GwpResourceConstitute {
	s.Name = &v
	return s
}

func (s *GwpResourceConstitute) SetPercent(v string) *GwpResourceConstitute {
	s.Percent = &v
	return s
}

func (s *GwpResourceConstitute) SetResourceType(v int32) *GwpResourceConstitute {
	s.ResourceType = &v
	return s
}

func (s *GwpResourceConstitute) SetUnit(v string) *GwpResourceConstitute {
	s.Unit = &v
	return s
}

func (s *GwpResourceConstitute) Validate() error {
	return dara.Validate(s)
}

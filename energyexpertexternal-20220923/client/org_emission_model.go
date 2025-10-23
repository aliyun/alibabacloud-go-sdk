// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrgEmission interface {
	dara.Model
	String() string
	GoString() string
	SetCarbonEmissionData(v float64) *OrgEmission
	GetCarbonEmissionData() *float64
	SetModuleEmissionList(v []*OrgEmissionModuleEmissionList) *OrgEmission
	GetModuleEmissionList() []*OrgEmissionModuleEmissionList
	SetName(v string) *OrgEmission
	GetName() *string
	SetNameKey(v string) *OrgEmission
	GetNameKey() *string
	SetRatio(v float64) *OrgEmission
	GetRatio() *float64
	SetSubEmissionItems(v []*OrgEmission) *OrgEmission
	GetSubEmissionItems() []*OrgEmission
	SetWeightingCarbonEmissionData(v float64) *OrgEmission
	GetWeightingCarbonEmissionData() *float64
	SetWeightingProportion(v float64) *OrgEmission
	GetWeightingProportion() *float64
	SetWeightingRatio(v float64) *OrgEmission
	GetWeightingRatio() *float64
}

type OrgEmission struct {
	CarbonEmissionData          *float64                         `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	ModuleEmissionList          []*OrgEmissionModuleEmissionList `json:"moduleEmissionList,omitempty" xml:"moduleEmissionList,omitempty" type:"Repeated"`
	Name                        *string                          `json:"name,omitempty" xml:"name,omitempty"`
	NameKey                     *string                          `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	Ratio                       *float64                         `json:"ratio,omitempty" xml:"ratio,omitempty"`
	SubEmissionItems            []*OrgEmission                   `json:"subEmissionItems,omitempty" xml:"subEmissionItems,omitempty" type:"Repeated"`
	WeightingCarbonEmissionData *float64                         `json:"weightingCarbonEmissionData,omitempty" xml:"weightingCarbonEmissionData,omitempty"`
	WeightingProportion         *float64                         `json:"weightingProportion,omitempty" xml:"weightingProportion,omitempty"`
	WeightingRatio              *float64                         `json:"weightingRatio,omitempty" xml:"weightingRatio,omitempty"`
}

func (s OrgEmission) String() string {
	return dara.Prettify(s)
}

func (s OrgEmission) GoString() string {
	return s.String()
}

func (s *OrgEmission) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *OrgEmission) GetModuleEmissionList() []*OrgEmissionModuleEmissionList {
	return s.ModuleEmissionList
}

func (s *OrgEmission) GetName() *string {
	return s.Name
}

func (s *OrgEmission) GetNameKey() *string {
	return s.NameKey
}

func (s *OrgEmission) GetRatio() *float64 {
	return s.Ratio
}

func (s *OrgEmission) GetSubEmissionItems() []*OrgEmission {
	return s.SubEmissionItems
}

func (s *OrgEmission) GetWeightingCarbonEmissionData() *float64 {
	return s.WeightingCarbonEmissionData
}

func (s *OrgEmission) GetWeightingProportion() *float64 {
	return s.WeightingProportion
}

func (s *OrgEmission) GetWeightingRatio() *float64 {
	return s.WeightingRatio
}

func (s *OrgEmission) SetCarbonEmissionData(v float64) *OrgEmission {
	s.CarbonEmissionData = &v
	return s
}

func (s *OrgEmission) SetModuleEmissionList(v []*OrgEmissionModuleEmissionList) *OrgEmission {
	s.ModuleEmissionList = v
	return s
}

func (s *OrgEmission) SetName(v string) *OrgEmission {
	s.Name = &v
	return s
}

func (s *OrgEmission) SetNameKey(v string) *OrgEmission {
	s.NameKey = &v
	return s
}

func (s *OrgEmission) SetRatio(v float64) *OrgEmission {
	s.Ratio = &v
	return s
}

func (s *OrgEmission) SetSubEmissionItems(v []*OrgEmission) *OrgEmission {
	s.SubEmissionItems = v
	return s
}

func (s *OrgEmission) SetWeightingCarbonEmissionData(v float64) *OrgEmission {
	s.WeightingCarbonEmissionData = &v
	return s
}

func (s *OrgEmission) SetWeightingProportion(v float64) *OrgEmission {
	s.WeightingProportion = &v
	return s
}

func (s *OrgEmission) SetWeightingRatio(v float64) *OrgEmission {
	s.WeightingRatio = &v
	return s
}

func (s *OrgEmission) Validate() error {
	if s.ModuleEmissionList != nil {
		for _, item := range s.ModuleEmissionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubEmissionItems != nil {
		for _, item := range s.SubEmissionItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OrgEmissionModuleEmissionList struct {
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	Name               *string  `json:"name,omitempty" xml:"name,omitempty"`
	NameKey            *string  `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	Ratio              *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
}

func (s OrgEmissionModuleEmissionList) String() string {
	return dara.Prettify(s)
}

func (s OrgEmissionModuleEmissionList) GoString() string {
	return s.String()
}

func (s *OrgEmissionModuleEmissionList) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *OrgEmissionModuleEmissionList) GetName() *string {
	return s.Name
}

func (s *OrgEmissionModuleEmissionList) GetNameKey() *string {
	return s.NameKey
}

func (s *OrgEmissionModuleEmissionList) GetRatio() *float64 {
	return s.Ratio
}

func (s *OrgEmissionModuleEmissionList) SetCarbonEmissionData(v float64) *OrgEmissionModuleEmissionList {
	s.CarbonEmissionData = &v
	return s
}

func (s *OrgEmissionModuleEmissionList) SetName(v string) *OrgEmissionModuleEmissionList {
	s.Name = &v
	return s
}

func (s *OrgEmissionModuleEmissionList) SetNameKey(v string) *OrgEmissionModuleEmissionList {
	s.NameKey = &v
	return s
}

func (s *OrgEmissionModuleEmissionList) SetRatio(v float64) *OrgEmissionModuleEmissionList {
	s.Ratio = &v
	return s
}

func (s *OrgEmissionModuleEmissionList) Validate() error {
	return dara.Validate(s)
}

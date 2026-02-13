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
	// Carbon Equivalent
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// Sub-module carbon emission data
	ModuleEmissionList []*OrgEmissionModuleEmissionList `json:"moduleEmissionList,omitempty" xml:"moduleEmissionList,omitempty" type:"Repeated"`
	// The name of the organization.
	//
	// example:
	//
	// EnterpriseZ
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Enterprise type
	//
	// example:
	//
	// Z-20240115-4
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of carbon emissions
	//
	// example:
	//
	// 0.2
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Sub-level data, that is, site data under the organization
	SubEmissionItems []*OrgEmission `json:"subEmissionItems,omitempty" xml:"subEmissionItems,omitempty" type:"Repeated"`
	// Calculate carbon emissions by share ratio
	//
	// example:
	//
	// 2.3
	WeightingCarbonEmissionData *float64 `json:"weightingCarbonEmissionData,omitempty" xml:"weightingCarbonEmissionData,omitempty"`
	// Weight ratio
	//
	// example:
	//
	// 0.3
	WeightingProportion *float64 `json:"weightingProportion,omitempty" xml:"weightingProportion,omitempty"`
	// Share ratio Carbon emissions YoY
	//
	// example:
	//
	// 0.4
	WeightingRatio *float64 `json:"weightingRatio,omitempty" xml:"weightingRatio,omitempty"`
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
	// Carbon emissions
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// module name
	//
	// example:
	//
	// Scope 1: Direct GHG emissions
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Module key
	//
	// example:
	//
	// carbonInventory.check.scope_1_direct_ghg_emissions
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of carbon emissions
	//
	// example:
	//
	// 0.2
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
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

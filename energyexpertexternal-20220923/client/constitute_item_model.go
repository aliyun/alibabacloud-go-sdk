// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConstituteItem interface {
	dara.Model
	String() string
	GoString() string
	SetCarbonEmissionData(v float64) *ConstituteItem
	GetCarbonEmissionData() *float64
	SetDataUnit(v string) *ConstituteItem
	GetDataUnit() *string
	SetEmissionSource(v string) *ConstituteItem
	GetEmissionSource() *string
	SetEmissionSourceKey(v string) *ConstituteItem
	GetEmissionSourceKey() *string
	SetEnterpriseName(v string) *ConstituteItem
	GetEnterpriseName() *string
	SetEnvGasEmissions(v []*ConstituteItemEnvGasEmissions) *ConstituteItem
	GetEnvGasEmissions() []*ConstituteItemEnvGasEmissions
	SetName(v string) *ConstituteItem
	GetName() *string
	SetNameKey(v string) *ConstituteItem
	GetNameKey() *string
	SetRatio(v float64) *ConstituteItem
	GetRatio() *float64
	SetRawData(v float64) *ConstituteItem
	GetRawData() *float64
	SetSubConstituteItems(v []*ConstituteItem) *ConstituteItem
	GetSubConstituteItems() []*ConstituteItem
}

type ConstituteItem struct {
	CarbonEmissionData *float64                         `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	DataUnit           *string                          `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	EmissionSource     *string                          `json:"emissionSource,omitempty" xml:"emissionSource,omitempty"`
	EmissionSourceKey  *string                          `json:"emissionSourceKey,omitempty" xml:"emissionSourceKey,omitempty"`
	EnterpriseName     *string                          `json:"enterpriseName,omitempty" xml:"enterpriseName,omitempty"`
	EnvGasEmissions    []*ConstituteItemEnvGasEmissions `json:"envGasEmissions,omitempty" xml:"envGasEmissions,omitempty" type:"Repeated"`
	Name               *string                          `json:"name,omitempty" xml:"name,omitempty"`
	NameKey            *string                          `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	Ratio              *float64                         `json:"ratio,omitempty" xml:"ratio,omitempty"`
	RawData            *float64                         `json:"rawData,omitempty" xml:"rawData,omitempty"`
	SubConstituteItems []*ConstituteItem                `json:"subConstituteItems,omitempty" xml:"subConstituteItems,omitempty" type:"Repeated"`
}

func (s ConstituteItem) String() string {
	return dara.Prettify(s)
}

func (s ConstituteItem) GoString() string {
	return s.String()
}

func (s *ConstituteItem) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *ConstituteItem) GetDataUnit() *string {
	return s.DataUnit
}

func (s *ConstituteItem) GetEmissionSource() *string {
	return s.EmissionSource
}

func (s *ConstituteItem) GetEmissionSourceKey() *string {
	return s.EmissionSourceKey
}

func (s *ConstituteItem) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *ConstituteItem) GetEnvGasEmissions() []*ConstituteItemEnvGasEmissions {
	return s.EnvGasEmissions
}

func (s *ConstituteItem) GetName() *string {
	return s.Name
}

func (s *ConstituteItem) GetNameKey() *string {
	return s.NameKey
}

func (s *ConstituteItem) GetRatio() *float64 {
	return s.Ratio
}

func (s *ConstituteItem) GetRawData() *float64 {
	return s.RawData
}

func (s *ConstituteItem) GetSubConstituteItems() []*ConstituteItem {
	return s.SubConstituteItems
}

func (s *ConstituteItem) SetCarbonEmissionData(v float64) *ConstituteItem {
	s.CarbonEmissionData = &v
	return s
}

func (s *ConstituteItem) SetDataUnit(v string) *ConstituteItem {
	s.DataUnit = &v
	return s
}

func (s *ConstituteItem) SetEmissionSource(v string) *ConstituteItem {
	s.EmissionSource = &v
	return s
}

func (s *ConstituteItem) SetEmissionSourceKey(v string) *ConstituteItem {
	s.EmissionSourceKey = &v
	return s
}

func (s *ConstituteItem) SetEnterpriseName(v string) *ConstituteItem {
	s.EnterpriseName = &v
	return s
}

func (s *ConstituteItem) SetEnvGasEmissions(v []*ConstituteItemEnvGasEmissions) *ConstituteItem {
	s.EnvGasEmissions = v
	return s
}

func (s *ConstituteItem) SetName(v string) *ConstituteItem {
	s.Name = &v
	return s
}

func (s *ConstituteItem) SetNameKey(v string) *ConstituteItem {
	s.NameKey = &v
	return s
}

func (s *ConstituteItem) SetRatio(v float64) *ConstituteItem {
	s.Ratio = &v
	return s
}

func (s *ConstituteItem) SetRawData(v float64) *ConstituteItem {
	s.RawData = &v
	return s
}

func (s *ConstituteItem) SetSubConstituteItems(v []*ConstituteItem) *ConstituteItem {
	s.SubConstituteItems = v
	return s
}

func (s *ConstituteItem) Validate() error {
	if s.EnvGasEmissions != nil {
		for _, item := range s.EnvGasEmissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubConstituteItems != nil {
		for _, item := range s.SubConstituteItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ConstituteItemEnvGasEmissions struct {
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	GasEmissionData    *float64 `json:"gasEmissionData,omitempty" xml:"gasEmissionData,omitempty"`
	Name               *string  `json:"name,omitempty" xml:"name,omitempty"`
	Type               *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ConstituteItemEnvGasEmissions) String() string {
	return dara.Prettify(s)
}

func (s ConstituteItemEnvGasEmissions) GoString() string {
	return s.String()
}

func (s *ConstituteItemEnvGasEmissions) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *ConstituteItemEnvGasEmissions) GetGasEmissionData() *float64 {
	return s.GasEmissionData
}

func (s *ConstituteItemEnvGasEmissions) GetName() *string {
	return s.Name
}

func (s *ConstituteItemEnvGasEmissions) GetType() *string {
	return s.Type
}

func (s *ConstituteItemEnvGasEmissions) SetCarbonEmissionData(v float64) *ConstituteItemEnvGasEmissions {
	s.CarbonEmissionData = &v
	return s
}

func (s *ConstituteItemEnvGasEmissions) SetGasEmissionData(v float64) *ConstituteItemEnvGasEmissions {
	s.GasEmissionData = &v
	return s
}

func (s *ConstituteItemEnvGasEmissions) SetName(v string) *ConstituteItemEnvGasEmissions {
	s.Name = &v
	return s
}

func (s *ConstituteItemEnvGasEmissions) SetType(v string) *ConstituteItemEnvGasEmissions {
	s.Type = &v
	return s
}

func (s *ConstituteItemEnvGasEmissions) Validate() error {
	return dara.Validate(s)
}

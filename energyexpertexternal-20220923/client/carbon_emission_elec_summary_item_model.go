// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarbonEmissionElecSummaryItem interface {
	dara.Model
	String() string
	GoString() string
	SetCarbonEmissionData(v float64) *CarbonEmissionElecSummaryItem
	GetCarbonEmissionData() *float64
	SetDataUnit(v string) *CarbonEmissionElecSummaryItem
	GetDataUnit() *string
	SetName(v string) *CarbonEmissionElecSummaryItem
	GetName() *string
	SetRatio(v float64) *CarbonEmissionElecSummaryItem
	GetRatio() *float64
	SetRawData(v float64) *CarbonEmissionElecSummaryItem
	GetRawData() *float64
}

type CarbonEmissionElecSummaryItem struct {
	// Carbon Equivalent.
	//
	// example:
	//
	// 1.22
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kg
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The enterprise name.
	//
	// example:
	//
	// xxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The percentage of electricity consumption. Valid values: 0 to 1.
	//
	// example:
	//
	// 0.22
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption in Kwh.
	//
	// example:
	//
	// 1.2
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s CarbonEmissionElecSummaryItem) String() string {
	return dara.Prettify(s)
}

func (s CarbonEmissionElecSummaryItem) GoString() string {
	return s.String()
}

func (s *CarbonEmissionElecSummaryItem) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *CarbonEmissionElecSummaryItem) GetDataUnit() *string {
	return s.DataUnit
}

func (s *CarbonEmissionElecSummaryItem) GetName() *string {
	return s.Name
}

func (s *CarbonEmissionElecSummaryItem) GetRatio() *float64 {
	return s.Ratio
}

func (s *CarbonEmissionElecSummaryItem) GetRawData() *float64 {
	return s.RawData
}

func (s *CarbonEmissionElecSummaryItem) SetCarbonEmissionData(v float64) *CarbonEmissionElecSummaryItem {
	s.CarbonEmissionData = &v
	return s
}

func (s *CarbonEmissionElecSummaryItem) SetDataUnit(v string) *CarbonEmissionElecSummaryItem {
	s.DataUnit = &v
	return s
}

func (s *CarbonEmissionElecSummaryItem) SetName(v string) *CarbonEmissionElecSummaryItem {
	s.Name = &v
	return s
}

func (s *CarbonEmissionElecSummaryItem) SetRatio(v float64) *CarbonEmissionElecSummaryItem {
	s.Ratio = &v
	return s
}

func (s *CarbonEmissionElecSummaryItem) SetRawData(v float64) *CarbonEmissionElecSummaryItem {
	s.RawData = &v
	return s
}

func (s *CarbonEmissionElecSummaryItem) Validate() error {
	return dara.Validate(s)
}

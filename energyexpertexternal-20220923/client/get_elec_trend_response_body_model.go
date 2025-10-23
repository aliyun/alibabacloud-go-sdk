// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElecTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetElecTrendResponseBody
	GetCode() *string
	SetData(v *GetElecTrendResponseBodyData) *GetElecTrendResponseBody
	GetData() *GetElecTrendResponseBodyData
	SetRequestId(v string) *GetElecTrendResponseBody
	GetRequestId() *string
}

type GetElecTrendResponseBody struct {
	// The code returned for the request. A value of Success indicates that the request was successful. Other values indicate that the request failed. You can troubleshoot the error by viewing the error message returned.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetElecTrendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetElecTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetElecTrendResponseBody) GetData() *GetElecTrendResponseBodyData {
	return s.Data
}

func (s *GetElecTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetElecTrendResponseBody) SetCode(v string) *GetElecTrendResponseBody {
	s.Code = &v
	return s
}

func (s *GetElecTrendResponseBody) SetData(v *GetElecTrendResponseBodyData) *GetElecTrendResponseBody {
	s.Data = v
	return s
}

func (s *GetElecTrendResponseBody) SetRequestId(v string) *GetElecTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetElecTrendResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetElecTrendResponseBodyData struct {
	// Photoelectricity monthly electricity consumption and carbon emissions and other data.
	Light []*GetElecTrendResponseBodyDataLight `json:"light,omitempty" xml:"light,omitempty" type:"Repeated"`
	// Monthly electricity consumption and carbon emissions data for nuclear power
	Nuclear []*GetElecTrendResponseBodyDataNuclear `json:"nuclear,omitempty" xml:"nuclear,omitempty" type:"Repeated"`
	// Monthly data on renewable electricity consumption and carbon emissions
	Renewing []*GetElecTrendResponseBodyDataRenewing `json:"renewing,omitempty" xml:"renewing,omitempty" type:"Repeated"`
	// Data such as monthly electricity consumption and carbon emissions from the mains.
	Urban []*GetElecTrendResponseBodyDataUrban `json:"urban,omitempty" xml:"urban,omitempty" type:"Repeated"`
	// Monthly data on electricity consumption and carbon emissions for hydropower.
	Water []*GetElecTrendResponseBodyDataWater `json:"water,omitempty" xml:"water,omitempty" type:"Repeated"`
	// Monthly wind power consumption and carbon emission data
	Wind []*GetElecTrendResponseBodyDataWind `json:"wind,omitempty" xml:"wind,omitempty" type:"Repeated"`
	// Zero electricity monthly electricity consumption and carbon emissions and other data.
	Zero []*GetElecTrendResponseBodyDataZero `json:"zero,omitempty" xml:"zero,omitempty" type:"Repeated"`
}

func (s GetElecTrendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyData) GetLight() []*GetElecTrendResponseBodyDataLight {
	return s.Light
}

func (s *GetElecTrendResponseBodyData) GetNuclear() []*GetElecTrendResponseBodyDataNuclear {
	return s.Nuclear
}

func (s *GetElecTrendResponseBodyData) GetRenewing() []*GetElecTrendResponseBodyDataRenewing {
	return s.Renewing
}

func (s *GetElecTrendResponseBodyData) GetUrban() []*GetElecTrendResponseBodyDataUrban {
	return s.Urban
}

func (s *GetElecTrendResponseBodyData) GetWater() []*GetElecTrendResponseBodyDataWater {
	return s.Water
}

func (s *GetElecTrendResponseBodyData) GetWind() []*GetElecTrendResponseBodyDataWind {
	return s.Wind
}

func (s *GetElecTrendResponseBodyData) GetZero() []*GetElecTrendResponseBodyDataZero {
	return s.Zero
}

func (s *GetElecTrendResponseBodyData) SetLight(v []*GetElecTrendResponseBodyDataLight) *GetElecTrendResponseBodyData {
	s.Light = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetNuclear(v []*GetElecTrendResponseBodyDataNuclear) *GetElecTrendResponseBodyData {
	s.Nuclear = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetRenewing(v []*GetElecTrendResponseBodyDataRenewing) *GetElecTrendResponseBodyData {
	s.Renewing = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetUrban(v []*GetElecTrendResponseBodyDataUrban) *GetElecTrendResponseBodyData {
	s.Urban = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetWater(v []*GetElecTrendResponseBodyDataWater) *GetElecTrendResponseBodyData {
	s.Water = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetWind(v []*GetElecTrendResponseBodyDataWind) *GetElecTrendResponseBodyData {
	s.Wind = v
	return s
}

func (s *GetElecTrendResponseBodyData) SetZero(v []*GetElecTrendResponseBodyDataZero) *GetElecTrendResponseBodyData {
	s.Zero = v
	return s
}

func (s *GetElecTrendResponseBodyData) Validate() error {
	if s.Light != nil {
		for _, item := range s.Light {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nuclear != nil {
		for _, item := range s.Nuclear {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Renewing != nil {
		for _, item := range s.Renewing {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Urban != nil {
		for _, item := range s.Urban {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Water != nil {
		for _, item := range s.Water {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Wind != nil {
		for _, item := range s.Wind {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Zero != nil {
		for _, item := range s.Zero {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetElecTrendResponseBodyDataLight struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power type name.
	//
	// example:
	//
	// Solar Power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.light_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e. 50%).
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataLight) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponseBodyDataLight) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataLight) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecTrendResponseBodyDataLight) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecTrendResponseBodyDataLight) GetMonth() *int32 {
	return s.Month
}

func (s *GetElecTrendResponseBodyDataLight) GetName() *string {
	return s.Name
}

func (s *GetElecTrendResponseBodyDataLight) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecTrendResponseBodyDataLight) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecTrendResponseBodyDataLight) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecTrendResponseBodyDataLight) GetYear() *string {
	return s.Year
}

func (s *GetElecTrendResponseBodyDataLight) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataLight {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetDataUnit(v string) *GetElecTrendResponseBodyDataLight {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetMonth(v int32) *GetElecTrendResponseBodyDataLight {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetName(v string) *GetElecTrendResponseBodyDataLight {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetNameKey(v string) *GetElecTrendResponseBodyDataLight {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetRatio(v float64) *GetElecTrendResponseBodyDataLight {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetRawData(v float64) *GetElecTrendResponseBodyDataLight {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) SetYear(v string) *GetElecTrendResponseBodyDataLight {
	s.Year = &v
	return s
}

func (s *GetElecTrendResponseBodyDataLight) Validate() error {
	return dara.Validate(s)
}

type GetElecTrendResponseBodyDataNuclear struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Nuclear power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.nuclear_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataNuclear) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponseBodyDataNuclear) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataNuclear) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecTrendResponseBodyDataNuclear) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecTrendResponseBodyDataNuclear) GetMonth() *int32 {
	return s.Month
}

func (s *GetElecTrendResponseBodyDataNuclear) GetName() *string {
	return s.Name
}

func (s *GetElecTrendResponseBodyDataNuclear) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecTrendResponseBodyDataNuclear) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecTrendResponseBodyDataNuclear) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecTrendResponseBodyDataNuclear) GetYear() *string {
	return s.Year
}

func (s *GetElecTrendResponseBodyDataNuclear) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataNuclear {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetDataUnit(v string) *GetElecTrendResponseBodyDataNuclear {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetMonth(v int32) *GetElecTrendResponseBodyDataNuclear {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetName(v string) *GetElecTrendResponseBodyDataNuclear {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetNameKey(v string) *GetElecTrendResponseBodyDataNuclear {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetRatio(v float64) *GetElecTrendResponseBodyDataNuclear {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetRawData(v float64) *GetElecTrendResponseBodyDataNuclear {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) SetYear(v string) *GetElecTrendResponseBodyDataNuclear {
	s.Year = &v
	return s
}

func (s *GetElecTrendResponseBodyDataNuclear) Validate() error {
	return dara.Validate(s)
}

type GetElecTrendResponseBodyDataRenewing struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Renewable electricity
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.carbonEmissionAnalysis.components.CarbonDetail.lingTanDianLi
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataRenewing) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponseBodyDataRenewing) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataRenewing) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecTrendResponseBodyDataRenewing) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecTrendResponseBodyDataRenewing) GetMonth() *int32 {
	return s.Month
}

func (s *GetElecTrendResponseBodyDataRenewing) GetName() *string {
	return s.Name
}

func (s *GetElecTrendResponseBodyDataRenewing) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecTrendResponseBodyDataRenewing) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecTrendResponseBodyDataRenewing) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecTrendResponseBodyDataRenewing) GetYear() *string {
	return s.Year
}

func (s *GetElecTrendResponseBodyDataRenewing) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataRenewing {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetDataUnit(v string) *GetElecTrendResponseBodyDataRenewing {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetMonth(v int32) *GetElecTrendResponseBodyDataRenewing {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetName(v string) *GetElecTrendResponseBodyDataRenewing {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetNameKey(v string) *GetElecTrendResponseBodyDataRenewing {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetRatio(v float64) *GetElecTrendResponseBodyDataRenewing {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetRawData(v float64) *GetElecTrendResponseBodyDataRenewing {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) SetYear(v string) *GetElecTrendResponseBodyDataRenewing {
	s.Year = &v
	return s
}

func (s *GetElecTrendResponseBodyDataRenewing) Validate() error {
	return dara.Validate(s)
}

type GetElecTrendResponseBodyDataUrban struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Grid power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.urban_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e. 50%).
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataUrban) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponseBodyDataUrban) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataUrban) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecTrendResponseBodyDataUrban) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecTrendResponseBodyDataUrban) GetMonth() *int32 {
	return s.Month
}

func (s *GetElecTrendResponseBodyDataUrban) GetName() *string {
	return s.Name
}

func (s *GetElecTrendResponseBodyDataUrban) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecTrendResponseBodyDataUrban) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecTrendResponseBodyDataUrban) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecTrendResponseBodyDataUrban) GetYear() *string {
	return s.Year
}

func (s *GetElecTrendResponseBodyDataUrban) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataUrban {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetDataUnit(v string) *GetElecTrendResponseBodyDataUrban {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetMonth(v int32) *GetElecTrendResponseBodyDataUrban {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetName(v string) *GetElecTrendResponseBodyDataUrban {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetNameKey(v string) *GetElecTrendResponseBodyDataUrban {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetRatio(v float64) *GetElecTrendResponseBodyDataUrban {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetRawData(v float64) *GetElecTrendResponseBodyDataUrban {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) SetYear(v string) *GetElecTrendResponseBodyDataUrban {
	s.Year = &v
	return s
}

func (s *GetElecTrendResponseBodyDataUrban) Validate() error {
	return dara.Validate(s)
}

type GetElecTrendResponseBodyDataWater struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Hydro power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.water_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e. 50%).
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataWater) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponseBodyDataWater) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataWater) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecTrendResponseBodyDataWater) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecTrendResponseBodyDataWater) GetMonth() *int32 {
	return s.Month
}

func (s *GetElecTrendResponseBodyDataWater) GetName() *string {
	return s.Name
}

func (s *GetElecTrendResponseBodyDataWater) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecTrendResponseBodyDataWater) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecTrendResponseBodyDataWater) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecTrendResponseBodyDataWater) GetYear() *string {
	return s.Year
}

func (s *GetElecTrendResponseBodyDataWater) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataWater {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetDataUnit(v string) *GetElecTrendResponseBodyDataWater {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetMonth(v int32) *GetElecTrendResponseBodyDataWater {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetName(v string) *GetElecTrendResponseBodyDataWater {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetNameKey(v string) *GetElecTrendResponseBodyDataWater {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetRatio(v float64) *GetElecTrendResponseBodyDataWater {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetRawData(v float64) *GetElecTrendResponseBodyDataWater {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) SetYear(v string) *GetElecTrendResponseBodyDataWater {
	s.Year = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWater) Validate() error {
	return dara.Validate(s)
}

type GetElecTrendResponseBodyDataWind struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Wind power
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.check.wind_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataWind) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponseBodyDataWind) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataWind) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecTrendResponseBodyDataWind) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecTrendResponseBodyDataWind) GetMonth() *int32 {
	return s.Month
}

func (s *GetElecTrendResponseBodyDataWind) GetName() *string {
	return s.Name
}

func (s *GetElecTrendResponseBodyDataWind) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecTrendResponseBodyDataWind) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecTrendResponseBodyDataWind) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecTrendResponseBodyDataWind) GetYear() *string {
	return s.Year
}

func (s *GetElecTrendResponseBodyDataWind) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataWind {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetDataUnit(v string) *GetElecTrendResponseBodyDataWind {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetMonth(v int32) *GetElecTrendResponseBodyDataWind {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetName(v string) *GetElecTrendResponseBodyDataWind {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetNameKey(v string) *GetElecTrendResponseBodyDataWind {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetRatio(v float64) *GetElecTrendResponseBodyDataWind {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetRawData(v float64) *GetElecTrendResponseBodyDataWind {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) SetYear(v string) *GetElecTrendResponseBodyDataWind {
	s.Year = &v
	return s
}

func (s *GetElecTrendResponseBodyDataWind) Validate() error {
	return dara.Validate(s)
}

type GetElecTrendResponseBodyDataZero struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The price unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// Month
	//
	// example:
	//
	// 12
	Month *int32 `json:"month,omitempty" xml:"month,omitempty"`
	// Power Type Name
	//
	// example:
	//
	// Zero carbon electricity
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Power Type Code
	//
	// example:
	//
	// carbonInventory.carbonEmissionAnalysis.components.CarbonDetail.lingTanDianLi
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 3.14
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
	// Year
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecTrendResponseBodyDataZero) String() string {
	return dara.Prettify(s)
}

func (s GetElecTrendResponseBodyDataZero) GoString() string {
	return s.String()
}

func (s *GetElecTrendResponseBodyDataZero) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecTrendResponseBodyDataZero) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecTrendResponseBodyDataZero) GetMonth() *int32 {
	return s.Month
}

func (s *GetElecTrendResponseBodyDataZero) GetName() *string {
	return s.Name
}

func (s *GetElecTrendResponseBodyDataZero) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecTrendResponseBodyDataZero) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecTrendResponseBodyDataZero) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecTrendResponseBodyDataZero) GetYear() *string {
	return s.Year
}

func (s *GetElecTrendResponseBodyDataZero) SetCarbonEmissionData(v float64) *GetElecTrendResponseBodyDataZero {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetDataUnit(v string) *GetElecTrendResponseBodyDataZero {
	s.DataUnit = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetMonth(v int32) *GetElecTrendResponseBodyDataZero {
	s.Month = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetName(v string) *GetElecTrendResponseBodyDataZero {
	s.Name = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetNameKey(v string) *GetElecTrendResponseBodyDataZero {
	s.NameKey = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetRatio(v float64) *GetElecTrendResponseBodyDataZero {
	s.Ratio = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetRawData(v float64) *GetElecTrendResponseBodyDataZero {
	s.RawData = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) SetYear(v string) *GetElecTrendResponseBodyDataZero {
	s.Year = &v
	return s
}

func (s *GetElecTrendResponseBodyDataZero) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElecConstituteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetElecConstituteResponseBodyData) *GetElecConstituteResponseBody
	GetData() *GetElecConstituteResponseBodyData
	SetRequestId(v string) *GetElecConstituteResponseBody
	GetRequestId() *string
}

type GetElecConstituteResponseBody struct {
	// The returned data.
	Data *GetElecConstituteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetElecConstituteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBody) GetData() *GetElecConstituteResponseBodyData {
	return s.Data
}

func (s *GetElecConstituteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetElecConstituteResponseBody) SetData(v *GetElecConstituteResponseBodyData) *GetElecConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetElecConstituteResponseBody) SetRequestId(v string) *GetElecConstituteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetElecConstituteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetElecConstituteResponseBodyData struct {
	// Photoelectric power consumption and carbon emission data of each enterprise.
	Light *GetElecConstituteResponseBodyDataLight `json:"light,omitempty" xml:"light,omitempty" type:"Struct"`
	// Data on nuclear power consumption and carbon emissions at each enterprise.
	Nuclear *GetElecConstituteResponseBodyDataNuclear `json:"nuclear,omitempty" xml:"nuclear,omitempty" type:"Struct"`
	// Data on renewable electricity consumption and carbon emissions at each enterprise.
	Renewing *GetElecConstituteResponseBodyDataRenewing `json:"renewing,omitempty" xml:"renewing,omitempty" type:"Struct"`
	// Data on mains power electricity consumption and carbon emission of each enterprise.
	Urban *GetElecConstituteResponseBodyDataUrban `json:"urban,omitempty" xml:"urban,omitempty" type:"Struct"`
	// Hydropower consumption and carbon emission data of each enterprise.
	Water *GetElecConstituteResponseBodyDataWater `json:"water,omitempty" xml:"water,omitempty" type:"Struct"`
	// Wind power consumption and carbon emission data of each enterprise.
	Wind *GetElecConstituteResponseBodyDataWind `json:"wind,omitempty" xml:"wind,omitempty" type:"Struct"`
	// Data of zero electricity consumption and carbon emission of each enterprise.
	Zero *GetElecConstituteResponseBodyDataZero `json:"zero,omitempty" xml:"zero,omitempty" type:"Struct"`
}

func (s GetElecConstituteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyData) GetLight() *GetElecConstituteResponseBodyDataLight {
	return s.Light
}

func (s *GetElecConstituteResponseBodyData) GetNuclear() *GetElecConstituteResponseBodyDataNuclear {
	return s.Nuclear
}

func (s *GetElecConstituteResponseBodyData) GetRenewing() *GetElecConstituteResponseBodyDataRenewing {
	return s.Renewing
}

func (s *GetElecConstituteResponseBodyData) GetUrban() *GetElecConstituteResponseBodyDataUrban {
	return s.Urban
}

func (s *GetElecConstituteResponseBodyData) GetWater() *GetElecConstituteResponseBodyDataWater {
	return s.Water
}

func (s *GetElecConstituteResponseBodyData) GetWind() *GetElecConstituteResponseBodyDataWind {
	return s.Wind
}

func (s *GetElecConstituteResponseBodyData) GetZero() *GetElecConstituteResponseBodyDataZero {
	return s.Zero
}

func (s *GetElecConstituteResponseBodyData) SetLight(v *GetElecConstituteResponseBodyDataLight) *GetElecConstituteResponseBodyData {
	s.Light = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetNuclear(v *GetElecConstituteResponseBodyDataNuclear) *GetElecConstituteResponseBodyData {
	s.Nuclear = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetRenewing(v *GetElecConstituteResponseBodyDataRenewing) *GetElecConstituteResponseBodyData {
	s.Renewing = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetUrban(v *GetElecConstituteResponseBodyDataUrban) *GetElecConstituteResponseBodyData {
	s.Urban = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetWater(v *GetElecConstituteResponseBodyDataWater) *GetElecConstituteResponseBodyData {
	s.Water = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetWind(v *GetElecConstituteResponseBodyDataWind) *GetElecConstituteResponseBodyData {
	s.Wind = v
	return s
}

func (s *GetElecConstituteResponseBodyData) SetZero(v *GetElecConstituteResponseBodyDataZero) *GetElecConstituteResponseBodyData {
	s.Zero = v
	return s
}

func (s *GetElecConstituteResponseBodyData) Validate() error {
	if s.Light != nil {
		if err := s.Light.Validate(); err != nil {
			return err
		}
	}
	if s.Nuclear != nil {
		if err := s.Nuclear.Validate(); err != nil {
			return err
		}
	}
	if s.Renewing != nil {
		if err := s.Renewing.Validate(); err != nil {
			return err
		}
	}
	if s.Urban != nil {
		if err := s.Urban.Validate(); err != nil {
			return err
		}
	}
	if s.Water != nil {
		if err := s.Water.Validate(); err != nil {
			return err
		}
	}
	if s.Wind != nil {
		if err := s.Wind.Validate(); err != nil {
			return err
		}
	}
	if s.Zero != nil {
		if err := s.Zero.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetElecConstituteResponseBodyDataLight struct {
	// Carbon emission.
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// light
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.light_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.2
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 1.2
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataLight) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataLight) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataLight) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecConstituteResponseBodyDataLight) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecConstituteResponseBodyDataLight) GetName() *string {
	return s.Name
}

func (s *GetElecConstituteResponseBodyDataLight) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecConstituteResponseBodyDataLight) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecConstituteResponseBodyDataLight) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecConstituteResponseBodyDataLight) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataLight {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetDataUnit(v string) *GetElecConstituteResponseBodyDataLight {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetName(v string) *GetElecConstituteResponseBodyDataLight {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetNameKey(v string) *GetElecConstituteResponseBodyDataLight {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetRatio(v float64) *GetElecConstituteResponseBodyDataLight {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) SetRawData(v float64) *GetElecConstituteResponseBodyDataLight {
	s.RawData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataLight) Validate() error {
	return dara.Validate(s)
}

type GetElecConstituteResponseBodyDataNuclear struct {
	// Carbon emission.
	//
	// example:
	//
	// 2.3
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// nuclear
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.nuclear_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.6
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 2
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataNuclear) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataNuclear) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataNuclear) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecConstituteResponseBodyDataNuclear) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecConstituteResponseBodyDataNuclear) GetName() *string {
	return s.Name
}

func (s *GetElecConstituteResponseBodyDataNuclear) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecConstituteResponseBodyDataNuclear) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecConstituteResponseBodyDataNuclear) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataNuclear {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetDataUnit(v string) *GetElecConstituteResponseBodyDataNuclear {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetName(v string) *GetElecConstituteResponseBodyDataNuclear {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetNameKey(v string) *GetElecConstituteResponseBodyDataNuclear {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetRatio(v float64) *GetElecConstituteResponseBodyDataNuclear {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) SetRawData(v float64) *GetElecConstituteResponseBodyDataNuclear {
	s.RawData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataNuclear) Validate() error {
	return dara.Validate(s)
}

type GetElecConstituteResponseBodyDataRenewing struct {
	// Carbon emission.
	//
	// example:
	//
	// 2.3
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// renewing
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.carbonEmissionAnalysis.components.CarbonDetail.keZaiShengZiYuan
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.44
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 4.3
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataRenewing) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataRenewing) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataRenewing) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecConstituteResponseBodyDataRenewing) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecConstituteResponseBodyDataRenewing) GetName() *string {
	return s.Name
}

func (s *GetElecConstituteResponseBodyDataRenewing) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecConstituteResponseBodyDataRenewing) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecConstituteResponseBodyDataRenewing) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataRenewing {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetDataUnit(v string) *GetElecConstituteResponseBodyDataRenewing {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetName(v string) *GetElecConstituteResponseBodyDataRenewing {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetNameKey(v string) *GetElecConstituteResponseBodyDataRenewing {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetRatio(v float64) *GetElecConstituteResponseBodyDataRenewing {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) SetRawData(v float64) *GetElecConstituteResponseBodyDataRenewing {
	s.RawData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataRenewing) Validate() error {
	return dara.Validate(s)
}

type GetElecConstituteResponseBodyDataUrban struct {
	// Carbon emission.
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// urban
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.urban_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.4
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 1.2
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataUrban) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataUrban) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataUrban) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecConstituteResponseBodyDataUrban) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecConstituteResponseBodyDataUrban) GetName() *string {
	return s.Name
}

func (s *GetElecConstituteResponseBodyDataUrban) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecConstituteResponseBodyDataUrban) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecConstituteResponseBodyDataUrban) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecConstituteResponseBodyDataUrban) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataUrban {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetDataUnit(v string) *GetElecConstituteResponseBodyDataUrban {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetName(v string) *GetElecConstituteResponseBodyDataUrban {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetNameKey(v string) *GetElecConstituteResponseBodyDataUrban {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetRatio(v float64) *GetElecConstituteResponseBodyDataUrban {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) SetRawData(v float64) *GetElecConstituteResponseBodyDataUrban {
	s.RawData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataUrban) Validate() error {
	return dara.Validate(s)
}

type GetElecConstituteResponseBodyDataWater struct {
	// Carbon emission.
	//
	// example:
	//
	// 2.1
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// water
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.water_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.4
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 1.2
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataWater) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataWater) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataWater) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecConstituteResponseBodyDataWater) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecConstituteResponseBodyDataWater) GetName() *string {
	return s.Name
}

func (s *GetElecConstituteResponseBodyDataWater) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecConstituteResponseBodyDataWater) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecConstituteResponseBodyDataWater) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecConstituteResponseBodyDataWater) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataWater {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetDataUnit(v string) *GetElecConstituteResponseBodyDataWater {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetName(v string) *GetElecConstituteResponseBodyDataWater {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetNameKey(v string) *GetElecConstituteResponseBodyDataWater {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetRatio(v float64) *GetElecConstituteResponseBodyDataWater {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) SetRawData(v float64) *GetElecConstituteResponseBodyDataWater {
	s.RawData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWater) Validate() error {
	return dara.Validate(s)
}

type GetElecConstituteResponseBodyDataWind struct {
	// Carbon emission.
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// wind
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.check.wind_electricity
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.3
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 1.1
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataWind) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataWind) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataWind) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecConstituteResponseBodyDataWind) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecConstituteResponseBodyDataWind) GetName() *string {
	return s.Name
}

func (s *GetElecConstituteResponseBodyDataWind) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecConstituteResponseBodyDataWind) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecConstituteResponseBodyDataWind) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecConstituteResponseBodyDataWind) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataWind {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetDataUnit(v string) *GetElecConstituteResponseBodyDataWind {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetName(v string) *GetElecConstituteResponseBodyDataWind {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetNameKey(v string) *GetElecConstituteResponseBodyDataWind {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetRatio(v float64) *GetElecConstituteResponseBodyDataWind {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) SetRawData(v float64) *GetElecConstituteResponseBodyDataWind {
	s.RawData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataWind) Validate() error {
	return dara.Validate(s)
}

type GetElecConstituteResponseBodyDataZero struct {
	// Carbon emission.
	//
	// example:
	//
	// 1.2
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// The unit.
	//
	// example:
	//
	// kWh
	DataUnit *string `json:"dataUnit,omitempty" xml:"dataUnit,omitempty"`
	// The name.
	//
	// example:
	//
	// zero
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique identifier of name.
	//
	// example:
	//
	// carbonInventory.carbonEmissionAnalysis.components.CarbonDetail.lingTanDianLi
	NameKey *string `json:"nameKey,omitempty" xml:"nameKey,omitempty"`
	// Proportion of electricity consumption to all electricity consumption in the month: example value: 0.5 (i. e., 50%)
	//
	// example:
	//
	// 0.33
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Electricity consumption
	//
	// example:
	//
	// 444.3
	RawData *float64 `json:"rawData,omitempty" xml:"rawData,omitempty"`
}

func (s GetElecConstituteResponseBodyDataZero) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteResponseBodyDataZero) GoString() string {
	return s.String()
}

func (s *GetElecConstituteResponseBodyDataZero) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetElecConstituteResponseBodyDataZero) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetElecConstituteResponseBodyDataZero) GetName() *string {
	return s.Name
}

func (s *GetElecConstituteResponseBodyDataZero) GetNameKey() *string {
	return s.NameKey
}

func (s *GetElecConstituteResponseBodyDataZero) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetElecConstituteResponseBodyDataZero) GetRawData() *float64 {
	return s.RawData
}

func (s *GetElecConstituteResponseBodyDataZero) SetCarbonEmissionData(v float64) *GetElecConstituteResponseBodyDataZero {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetDataUnit(v string) *GetElecConstituteResponseBodyDataZero {
	s.DataUnit = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetName(v string) *GetElecConstituteResponseBodyDataZero {
	s.Name = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetNameKey(v string) *GetElecConstituteResponseBodyDataZero {
	s.NameKey = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetRatio(v float64) *GetElecConstituteResponseBodyDataZero {
	s.Ratio = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) SetRawData(v float64) *GetElecConstituteResponseBodyDataZero {
	s.RawData = &v
	return s
}

func (s *GetElecConstituteResponseBodyDataZero) Validate() error {
	return dara.Validate(s)
}

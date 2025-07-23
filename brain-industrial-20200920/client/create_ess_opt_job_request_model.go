// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEssOptJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessKey(v string) *CreateEssOptJobRequest
	GetBusinessKey() *string
	SetDuration(v int32) *CreateEssOptJobRequest
	GetDuration() *int32
	SetElecPrice(v []*CreateEssOptJobRequestElecPrice) *CreateEssOptJobRequest
	GetElecPrice() []*CreateEssOptJobRequestElecPrice
	SetFreq(v string) *CreateEssOptJobRequest
	GetFreq() *string
	SetGenPrice(v []*CreateEssOptJobRequestGenPrice) *CreateEssOptJobRequest
	GetGenPrice() []*CreateEssOptJobRequestGenPrice
	SetLocation(v *CreateEssOptJobRequestLocation) *CreateEssOptJobRequest
	GetLocation() *CreateEssOptJobRequestLocation
	SetModelVersion(v string) *CreateEssOptJobRequest
	GetModelVersion() *string
	SetRunDate(v string) *CreateEssOptJobRequest
	GetRunDate() *string
	SetSystemData(v []*CreateEssOptJobRequestSystemData) *CreateEssOptJobRequest
	GetSystemData() []*CreateEssOptJobRequestSystemData
	SetTimeZone(v string) *CreateEssOptJobRequest
	GetTimeZone() *string
	SetTopoType(v string) *CreateEssOptJobRequest
	GetTopoType() *string
}

type CreateEssOptJobRequest struct {
	BusinessKey *string `json:"BusinessKey,omitempty" xml:"BusinessKey,omitempty"`
	// example:
	//
	// 1
	Duration  *int32                             `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ElecPrice []*CreateEssOptJobRequestElecPrice `json:"ElecPrice,omitempty" xml:"ElecPrice,omitempty" type:"Repeated"`
	// example:
	//
	// FIFTEEN_MIN
	Freq     *string                           `json:"Freq,omitempty" xml:"Freq,omitempty"`
	GenPrice []*CreateEssOptJobRequestGenPrice `json:"GenPrice,omitempty" xml:"GenPrice,omitempty" type:"Repeated"`
	Location *CreateEssOptJobRequestLocation   `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-02-12
	RunDate    *string                             `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	SystemData []*CreateEssOptJobRequestSystemData `json:"SystemData,omitempty" xml:"SystemData,omitempty" type:"Repeated"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// example:
	//
	// LOAD_ESS_SOLAR
	TopoType *string `json:"TopoType,omitempty" xml:"TopoType,omitempty"`
}

func (s CreateEssOptJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobRequest) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequest) GetBusinessKey() *string {
	return s.BusinessKey
}

func (s *CreateEssOptJobRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateEssOptJobRequest) GetElecPrice() []*CreateEssOptJobRequestElecPrice {
	return s.ElecPrice
}

func (s *CreateEssOptJobRequest) GetFreq() *string {
	return s.Freq
}

func (s *CreateEssOptJobRequest) GetGenPrice() []*CreateEssOptJobRequestGenPrice {
	return s.GenPrice
}

func (s *CreateEssOptJobRequest) GetLocation() *CreateEssOptJobRequestLocation {
	return s.Location
}

func (s *CreateEssOptJobRequest) GetModelVersion() *string {
	return s.ModelVersion
}

func (s *CreateEssOptJobRequest) GetRunDate() *string {
	return s.RunDate
}

func (s *CreateEssOptJobRequest) GetSystemData() []*CreateEssOptJobRequestSystemData {
	return s.SystemData
}

func (s *CreateEssOptJobRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateEssOptJobRequest) GetTopoType() *string {
	return s.TopoType
}

func (s *CreateEssOptJobRequest) SetBusinessKey(v string) *CreateEssOptJobRequest {
	s.BusinessKey = &v
	return s
}

func (s *CreateEssOptJobRequest) SetDuration(v int32) *CreateEssOptJobRequest {
	s.Duration = &v
	return s
}

func (s *CreateEssOptJobRequest) SetElecPrice(v []*CreateEssOptJobRequestElecPrice) *CreateEssOptJobRequest {
	s.ElecPrice = v
	return s
}

func (s *CreateEssOptJobRequest) SetFreq(v string) *CreateEssOptJobRequest {
	s.Freq = &v
	return s
}

func (s *CreateEssOptJobRequest) SetGenPrice(v []*CreateEssOptJobRequestGenPrice) *CreateEssOptJobRequest {
	s.GenPrice = v
	return s
}

func (s *CreateEssOptJobRequest) SetLocation(v *CreateEssOptJobRequestLocation) *CreateEssOptJobRequest {
	s.Location = v
	return s
}

func (s *CreateEssOptJobRequest) SetModelVersion(v string) *CreateEssOptJobRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateEssOptJobRequest) SetRunDate(v string) *CreateEssOptJobRequest {
	s.RunDate = &v
	return s
}

func (s *CreateEssOptJobRequest) SetSystemData(v []*CreateEssOptJobRequestSystemData) *CreateEssOptJobRequest {
	s.SystemData = v
	return s
}

func (s *CreateEssOptJobRequest) SetTimeZone(v string) *CreateEssOptJobRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateEssOptJobRequest) SetTopoType(v string) *CreateEssOptJobRequest {
	s.TopoType = &v
	return s
}

func (s *CreateEssOptJobRequest) Validate() error {
	return dara.Validate(s)
}

type CreateEssOptJobRequestElecPrice struct {
	// example:
	//
	// 00:00:15
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// example:
	//
	// 0.5
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s CreateEssOptJobRequestElecPrice) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobRequestElecPrice) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequestElecPrice) GetDataTime() *string {
	return s.DataTime
}

func (s *CreateEssOptJobRequestElecPrice) GetPrice() *string {
	return s.Price
}

func (s *CreateEssOptJobRequestElecPrice) SetDataTime(v string) *CreateEssOptJobRequestElecPrice {
	s.DataTime = &v
	return s
}

func (s *CreateEssOptJobRequestElecPrice) SetPrice(v string) *CreateEssOptJobRequestElecPrice {
	s.Price = &v
	return s
}

func (s *CreateEssOptJobRequestElecPrice) Validate() error {
	return dara.Validate(s)
}

type CreateEssOptJobRequestGenPrice struct {
	// example:
	//
	// 00:00:15
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// example:
	//
	// 0.3
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s CreateEssOptJobRequestGenPrice) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobRequestGenPrice) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequestGenPrice) GetDataTime() *string {
	return s.DataTime
}

func (s *CreateEssOptJobRequestGenPrice) GetPrice() *string {
	return s.Price
}

func (s *CreateEssOptJobRequestGenPrice) SetDataTime(v string) *CreateEssOptJobRequestGenPrice {
	s.DataTime = &v
	return s
}

func (s *CreateEssOptJobRequestGenPrice) SetPrice(v string) *CreateEssOptJobRequestGenPrice {
	s.Price = &v
	return s
}

func (s *CreateEssOptJobRequestGenPrice) Validate() error {
	return dara.Validate(s)
}

type CreateEssOptJobRequestLocation struct {
	// example:
	//
	// 10.123
	Altitude *float64 `json:"Altitude,omitempty" xml:"Altitude,omitempty"`
	// example:
	//
	// 40.027
	Latitude *float64 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// 120.042
	Longitude *float64 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
}

func (s CreateEssOptJobRequestLocation) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobRequestLocation) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequestLocation) GetAltitude() *float64 {
	return s.Altitude
}

func (s *CreateEssOptJobRequestLocation) GetLatitude() *float64 {
	return s.Latitude
}

func (s *CreateEssOptJobRequestLocation) GetLongitude() *float64 {
	return s.Longitude
}

func (s *CreateEssOptJobRequestLocation) SetAltitude(v float64) *CreateEssOptJobRequestLocation {
	s.Altitude = &v
	return s
}

func (s *CreateEssOptJobRequestLocation) SetLatitude(v float64) *CreateEssOptJobRequestLocation {
	s.Latitude = &v
	return s
}

func (s *CreateEssOptJobRequestLocation) SetLongitude(v float64) *CreateEssOptJobRequestLocation {
	s.Longitude = &v
	return s
}

func (s *CreateEssOptJobRequestLocation) Validate() error {
	return dara.Validate(s)
}

type CreateEssOptJobRequestSystemData struct {
	HistoryData []map[string]interface{} `json:"HistoryData,omitempty" xml:"HistoryData,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SystemId     *string                `json:"SystemId,omitempty" xml:"SystemId,omitempty"`
	SystemParams map[string]interface{} `json:"SystemParams,omitempty" xml:"SystemParams,omitempty"`
	// example:
	//
	// ess
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
}

func (s CreateEssOptJobRequestSystemData) String() string {
	return dara.Prettify(s)
}

func (s CreateEssOptJobRequestSystemData) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequestSystemData) GetHistoryData() []map[string]interface{} {
	return s.HistoryData
}

func (s *CreateEssOptJobRequestSystemData) GetSystemId() *string {
	return s.SystemId
}

func (s *CreateEssOptJobRequestSystemData) GetSystemParams() map[string]interface{} {
	return s.SystemParams
}

func (s *CreateEssOptJobRequestSystemData) GetSystemType() *string {
	return s.SystemType
}

func (s *CreateEssOptJobRequestSystemData) SetHistoryData(v []map[string]interface{}) *CreateEssOptJobRequestSystemData {
	s.HistoryData = v
	return s
}

func (s *CreateEssOptJobRequestSystemData) SetSystemId(v string) *CreateEssOptJobRequestSystemData {
	s.SystemId = &v
	return s
}

func (s *CreateEssOptJobRequestSystemData) SetSystemParams(v map[string]interface{}) *CreateEssOptJobRequestSystemData {
	s.SystemParams = v
	return s
}

func (s *CreateEssOptJobRequestSystemData) SetSystemType(v string) *CreateEssOptJobRequestSystemData {
	s.SystemType = &v
	return s
}

func (s *CreateEssOptJobRequestSystemData) Validate() error {
	return dara.Validate(s)
}

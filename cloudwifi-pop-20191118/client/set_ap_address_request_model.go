// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApAreaName(v string) *SetApAddressRequest
	GetApAreaName() *string
	SetApBuildingName(v string) *SetApAddressRequest
	GetApBuildingName() *string
	SetApCampusName(v string) *SetApAddressRequest
	GetApCampusName() *string
	SetApCityName(v string) *SetApAddressRequest
	GetApCityName() *string
	SetApFloor(v string) *SetApAddressRequest
	GetApFloor() *string
	SetApGroup(v string) *SetApAddressRequest
	GetApGroup() *string
	SetApName(v string) *SetApAddressRequest
	GetApName() *string
	SetApNationName(v string) *SetApAddressRequest
	GetApNationName() *string
	SetApProvinceName(v string) *SetApAddressRequest
	GetApProvinceName() *string
	SetApUnitId(v int64) *SetApAddressRequest
	GetApUnitId() *int64
	SetApUnitName(v string) *SetApAddressRequest
	GetApUnitName() *string
	SetAppCode(v string) *SetApAddressRequest
	GetAppCode() *string
	SetAppName(v string) *SetApAddressRequest
	GetAppName() *string
	SetDirection(v string) *SetApAddressRequest
	GetDirection() *string
	SetLanguage(v string) *SetApAddressRequest
	GetLanguage() *string
	SetLat(v string) *SetApAddressRequest
	GetLat() *string
	SetLng(v string) *SetApAddressRequest
	GetLng() *string
	SetMac(v string) *SetApAddressRequest
	GetMac() *string
}

type SetApAddressRequest struct {
	ApAreaName     *string `json:"ApAreaName,omitempty" xml:"ApAreaName,omitempty"`
	ApBuildingName *string `json:"ApBuildingName,omitempty" xml:"ApBuildingName,omitempty"`
	ApCampusName   *string `json:"ApCampusName,omitempty" xml:"ApCampusName,omitempty"`
	ApCityName     *string `json:"ApCityName,omitempty" xml:"ApCityName,omitempty"`
	ApFloor        *string `json:"ApFloor,omitempty" xml:"ApFloor,omitempty"`
	ApGroup        *string `json:"ApGroup,omitempty" xml:"ApGroup,omitempty"`
	ApName         *string `json:"ApName,omitempty" xml:"ApName,omitempty"`
	ApNationName   *string `json:"ApNationName,omitempty" xml:"ApNationName,omitempty"`
	ApProvinceName *string `json:"ApProvinceName,omitempty" xml:"ApProvinceName,omitempty"`
	ApUnitId       *int64  `json:"ApUnitId,omitempty" xml:"ApUnitId,omitempty"`
	ApUnitName     *string `json:"ApUnitName,omitempty" xml:"ApUnitName,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// This parameter is required.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	Lat      *string `json:"Lat,omitempty" xml:"Lat,omitempty"`
	Lng      *string `json:"Lng,omitempty" xml:"Lng,omitempty"`
	// This parameter is required.
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
}

func (s SetApAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s SetApAddressRequest) GoString() string {
	return s.String()
}

func (s *SetApAddressRequest) GetApAreaName() *string {
	return s.ApAreaName
}

func (s *SetApAddressRequest) GetApBuildingName() *string {
	return s.ApBuildingName
}

func (s *SetApAddressRequest) GetApCampusName() *string {
	return s.ApCampusName
}

func (s *SetApAddressRequest) GetApCityName() *string {
	return s.ApCityName
}

func (s *SetApAddressRequest) GetApFloor() *string {
	return s.ApFloor
}

func (s *SetApAddressRequest) GetApGroup() *string {
	return s.ApGroup
}

func (s *SetApAddressRequest) GetApName() *string {
	return s.ApName
}

func (s *SetApAddressRequest) GetApNationName() *string {
	return s.ApNationName
}

func (s *SetApAddressRequest) GetApProvinceName() *string {
	return s.ApProvinceName
}

func (s *SetApAddressRequest) GetApUnitId() *int64 {
	return s.ApUnitId
}

func (s *SetApAddressRequest) GetApUnitName() *string {
	return s.ApUnitName
}

func (s *SetApAddressRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *SetApAddressRequest) GetAppName() *string {
	return s.AppName
}

func (s *SetApAddressRequest) GetDirection() *string {
	return s.Direction
}

func (s *SetApAddressRequest) GetLanguage() *string {
	return s.Language
}

func (s *SetApAddressRequest) GetLat() *string {
	return s.Lat
}

func (s *SetApAddressRequest) GetLng() *string {
	return s.Lng
}

func (s *SetApAddressRequest) GetMac() *string {
	return s.Mac
}

func (s *SetApAddressRequest) SetApAreaName(v string) *SetApAddressRequest {
	s.ApAreaName = &v
	return s
}

func (s *SetApAddressRequest) SetApBuildingName(v string) *SetApAddressRequest {
	s.ApBuildingName = &v
	return s
}

func (s *SetApAddressRequest) SetApCampusName(v string) *SetApAddressRequest {
	s.ApCampusName = &v
	return s
}

func (s *SetApAddressRequest) SetApCityName(v string) *SetApAddressRequest {
	s.ApCityName = &v
	return s
}

func (s *SetApAddressRequest) SetApFloor(v string) *SetApAddressRequest {
	s.ApFloor = &v
	return s
}

func (s *SetApAddressRequest) SetApGroup(v string) *SetApAddressRequest {
	s.ApGroup = &v
	return s
}

func (s *SetApAddressRequest) SetApName(v string) *SetApAddressRequest {
	s.ApName = &v
	return s
}

func (s *SetApAddressRequest) SetApNationName(v string) *SetApAddressRequest {
	s.ApNationName = &v
	return s
}

func (s *SetApAddressRequest) SetApProvinceName(v string) *SetApAddressRequest {
	s.ApProvinceName = &v
	return s
}

func (s *SetApAddressRequest) SetApUnitId(v int64) *SetApAddressRequest {
	s.ApUnitId = &v
	return s
}

func (s *SetApAddressRequest) SetApUnitName(v string) *SetApAddressRequest {
	s.ApUnitName = &v
	return s
}

func (s *SetApAddressRequest) SetAppCode(v string) *SetApAddressRequest {
	s.AppCode = &v
	return s
}

func (s *SetApAddressRequest) SetAppName(v string) *SetApAddressRequest {
	s.AppName = &v
	return s
}

func (s *SetApAddressRequest) SetDirection(v string) *SetApAddressRequest {
	s.Direction = &v
	return s
}

func (s *SetApAddressRequest) SetLanguage(v string) *SetApAddressRequest {
	s.Language = &v
	return s
}

func (s *SetApAddressRequest) SetLat(v string) *SetApAddressRequest {
	s.Lat = &v
	return s
}

func (s *SetApAddressRequest) SetLng(v string) *SetApAddressRequest {
	s.Lng = &v
	return s
}

func (s *SetApAddressRequest) SetMac(v string) *SetApAddressRequest {
	s.Mac = &v
	return s
}

func (s *SetApAddressRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKeywordSuggestInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *KeywordSuggestInfo
	GetAddress() *string
	SetBusinessAreaWithCity(v *KeywordSuggestInfo) *KeywordSuggestInfo
	GetBusinessAreaWithCity() *KeywordSuggestInfo
	SetCityCode(v int32) *KeywordSuggestInfo
	GetCityCode() *int32
	SetCityName(v string) *KeywordSuggestInfo
	GetCityName() *string
	SetDisplayName(v string) *KeywordSuggestInfo
	GetDisplayName() *string
	SetHotelId(v string) *KeywordSuggestInfo
	GetHotelId() *string
	SetIcon(v string) *KeywordSuggestInfo
	GetIcon() *string
	SetPoint(v string) *KeywordSuggestInfo
	GetPoint() *string
	SetPrice(v string) *KeywordSuggestInfo
	GetPrice() *string
	SetRegion(v int32) *KeywordSuggestInfo
	GetRegion() *int32
	SetType(v int32) *KeywordSuggestInfo
	GetType() *int32
	SetTypeDesc(v string) *KeywordSuggestInfo
	GetTypeDesc() *string
}

type KeywordSuggestInfo struct {
	Address              *string             `json:"address,omitempty" xml:"address,omitempty"`
	BusinessAreaWithCity *KeywordSuggestInfo `json:"business_area_with_city,omitempty" xml:"business_area_with_city,omitempty"`
	// example:
	//
	// 300100
	CityCode *int32 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// example:
	//
	// 杭州
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 杭州东站
	DisplayName *string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	// example:
	//
	// 53853318
	HotelId *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	Icon    *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// 4.8分
	Point *string `json:"point,omitempty" xml:"point,omitempty"`
	// example:
	//
	// 524
	Price *string `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 0
	Region *int32 `json:"region,omitempty" xml:"region,omitempty"`
	Type   *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 景点
	TypeDesc *string `json:"type_desc,omitempty" xml:"type_desc,omitempty"`
}

func (s KeywordSuggestInfo) String() string {
	return dara.Prettify(s)
}

func (s KeywordSuggestInfo) GoString() string {
	return s.String()
}

func (s *KeywordSuggestInfo) GetAddress() *string {
	return s.Address
}

func (s *KeywordSuggestInfo) GetBusinessAreaWithCity() *KeywordSuggestInfo {
	return s.BusinessAreaWithCity
}

func (s *KeywordSuggestInfo) GetCityCode() *int32 {
	return s.CityCode
}

func (s *KeywordSuggestInfo) GetCityName() *string {
	return s.CityName
}

func (s *KeywordSuggestInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *KeywordSuggestInfo) GetHotelId() *string {
	return s.HotelId
}

func (s *KeywordSuggestInfo) GetIcon() *string {
	return s.Icon
}

func (s *KeywordSuggestInfo) GetPoint() *string {
	return s.Point
}

func (s *KeywordSuggestInfo) GetPrice() *string {
	return s.Price
}

func (s *KeywordSuggestInfo) GetRegion() *int32 {
	return s.Region
}

func (s *KeywordSuggestInfo) GetType() *int32 {
	return s.Type
}

func (s *KeywordSuggestInfo) GetTypeDesc() *string {
	return s.TypeDesc
}

func (s *KeywordSuggestInfo) SetAddress(v string) *KeywordSuggestInfo {
	s.Address = &v
	return s
}

func (s *KeywordSuggestInfo) SetBusinessAreaWithCity(v *KeywordSuggestInfo) *KeywordSuggestInfo {
	s.BusinessAreaWithCity = v
	return s
}

func (s *KeywordSuggestInfo) SetCityCode(v int32) *KeywordSuggestInfo {
	s.CityCode = &v
	return s
}

func (s *KeywordSuggestInfo) SetCityName(v string) *KeywordSuggestInfo {
	s.CityName = &v
	return s
}

func (s *KeywordSuggestInfo) SetDisplayName(v string) *KeywordSuggestInfo {
	s.DisplayName = &v
	return s
}

func (s *KeywordSuggestInfo) SetHotelId(v string) *KeywordSuggestInfo {
	s.HotelId = &v
	return s
}

func (s *KeywordSuggestInfo) SetIcon(v string) *KeywordSuggestInfo {
	s.Icon = &v
	return s
}

func (s *KeywordSuggestInfo) SetPoint(v string) *KeywordSuggestInfo {
	s.Point = &v
	return s
}

func (s *KeywordSuggestInfo) SetPrice(v string) *KeywordSuggestInfo {
	s.Price = &v
	return s
}

func (s *KeywordSuggestInfo) SetRegion(v int32) *KeywordSuggestInfo {
	s.Region = &v
	return s
}

func (s *KeywordSuggestInfo) SetType(v int32) *KeywordSuggestInfo {
	s.Type = &v
	return s
}

func (s *KeywordSuggestInfo) SetTypeDesc(v string) *KeywordSuggestInfo {
	s.TypeDesc = &v
	return s
}

func (s *KeywordSuggestInfo) Validate() error {
	return dara.Validate(s)
}

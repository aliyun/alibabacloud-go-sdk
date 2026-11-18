// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryResult interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryResultData) *QueryResult
	GetData() []*QueryResultData
	SetRequestId(v string) *QueryResult
	GetRequestId() *string
}

type QueryResult struct {
	// The list of returned locations.
	Data []*QueryResultData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The unique request ID for troubleshooting.
	//
	// example:
	//
	// 15A85F99-408F-5BFA-B617-0C7D3E732ABD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryResult) String() string {
	return dara.Prettify(s)
}

func (s QueryResult) GoString() string {
	return s.String()
}

func (s *QueryResult) GetData() []*QueryResultData {
	return s.Data
}

func (s *QueryResult) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryResult) SetData(v []*QueryResultData) *QueryResult {
	s.Data = v
	return s
}

func (s *QueryResult) SetRequestId(v string) *QueryResult {
	s.RequestId = &v
	return s
}

func (s *QueryResult) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryResultData struct {
	// The detailed address.
	//
	// example:
	//
	// No. 22 Fayun Lane, West Lake Scenic Area
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The city code.
	//
	// example:
	//
	// 0571
	CityCode *string `json:"cityCode,omitempty" xml:"cityCode,omitempty"`
	// The city name.
	//
	// example:
	//
	// Hangzhou
	CityName *string `json:"cityName,omitempty" xml:"cityName,omitempty"`
	// The distance from the center point in meters.
	//
	// example:
	//
	// 815
	DistanceMeter *string `json:"distanceMeter,omitempty" xml:"distanceMeter,omitempty"`
	// The district code.
	//
	// example:
	//
	// 330106
	DistrictCode *string `json:"districtCode,omitempty" xml:"districtCode,omitempty"`
	// The district name.
	//
	// example:
	//
	// Xihu District
	DistrictName *string `json:"districtName,omitempty" xml:"districtName,omitempty"`
	// The internal identifier of the location.
	//
	// example:
	//
	// B023B0BI3J
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The list of images.
	Images []*QueryResultDataImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// The latitude of the location.
	//
	// example:
	//
	// 30.234403
	Latitude *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// The longitude of the location.
	//
	// example:
	//
	// 120.097327
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// The related business information.
	Metadata *QueryResultDataMetadata `json:"metadata,omitempty" xml:"metadata,omitempty" type:"Struct"`
	// The location name.
	//
	// example:
	//
	// Amanfayun Hangzhou
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The province code.
	//
	// example:
	//
	// 330000
	ProvinceCode *string `json:"provinceCode,omitempty" xml:"provinceCode,omitempty"`
	// The province name.
	//
	// example:
	//
	// Zhejiang Province
	ProvinceName *string `json:"provinceName,omitempty" xml:"provinceName,omitempty"`
	// The location type code.
	//
	// example:
	//
	// 100102
	TypeCode *string `json:"typeCode,omitempty" xml:"typeCode,omitempty"`
	// The location types. Multiple types are separated by vertical bars.
	//
	// example:
	//
	// Accommodation|Hotel|Five-star hotel
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s QueryResultData) String() string {
	return dara.Prettify(s)
}

func (s QueryResultData) GoString() string {
	return s.String()
}

func (s *QueryResultData) GetAddress() *string {
	return s.Address
}

func (s *QueryResultData) GetCityCode() *string {
	return s.CityCode
}

func (s *QueryResultData) GetCityName() *string {
	return s.CityName
}

func (s *QueryResultData) GetDistanceMeter() *string {
	return s.DistanceMeter
}

func (s *QueryResultData) GetDistrictCode() *string {
	return s.DistrictCode
}

func (s *QueryResultData) GetDistrictName() *string {
	return s.DistrictName
}

func (s *QueryResultData) GetId() *string {
	return s.Id
}

func (s *QueryResultData) GetImages() []*QueryResultDataImages {
	return s.Images
}

func (s *QueryResultData) GetLatitude() *string {
	return s.Latitude
}

func (s *QueryResultData) GetLongitude() *string {
	return s.Longitude
}

func (s *QueryResultData) GetMetadata() *QueryResultDataMetadata {
	return s.Metadata
}

func (s *QueryResultData) GetName() *string {
	return s.Name
}

func (s *QueryResultData) GetProvinceCode() *string {
	return s.ProvinceCode
}

func (s *QueryResultData) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *QueryResultData) GetTypeCode() *string {
	return s.TypeCode
}

func (s *QueryResultData) GetTypes() *string {
	return s.Types
}

func (s *QueryResultData) SetAddress(v string) *QueryResultData {
	s.Address = &v
	return s
}

func (s *QueryResultData) SetCityCode(v string) *QueryResultData {
	s.CityCode = &v
	return s
}

func (s *QueryResultData) SetCityName(v string) *QueryResultData {
	s.CityName = &v
	return s
}

func (s *QueryResultData) SetDistanceMeter(v string) *QueryResultData {
	s.DistanceMeter = &v
	return s
}

func (s *QueryResultData) SetDistrictCode(v string) *QueryResultData {
	s.DistrictCode = &v
	return s
}

func (s *QueryResultData) SetDistrictName(v string) *QueryResultData {
	s.DistrictName = &v
	return s
}

func (s *QueryResultData) SetId(v string) *QueryResultData {
	s.Id = &v
	return s
}

func (s *QueryResultData) SetImages(v []*QueryResultDataImages) *QueryResultData {
	s.Images = v
	return s
}

func (s *QueryResultData) SetLatitude(v string) *QueryResultData {
	s.Latitude = &v
	return s
}

func (s *QueryResultData) SetLongitude(v string) *QueryResultData {
	s.Longitude = &v
	return s
}

func (s *QueryResultData) SetMetadata(v *QueryResultDataMetadata) *QueryResultData {
	s.Metadata = v
	return s
}

func (s *QueryResultData) SetName(v string) *QueryResultData {
	s.Name = &v
	return s
}

func (s *QueryResultData) SetProvinceCode(v string) *QueryResultData {
	s.ProvinceCode = &v
	return s
}

func (s *QueryResultData) SetProvinceName(v string) *QueryResultData {
	s.ProvinceName = &v
	return s
}

func (s *QueryResultData) SetTypeCode(v string) *QueryResultData {
	s.TypeCode = &v
	return s
}

func (s *QueryResultData) SetTypes(v string) *QueryResultData {
	s.Types = &v
	return s
}

func (s *QueryResultData) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryResultDataImages struct {
	// The image name.
	//
	// example:
	//
	// Lobby
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The image URL.
	//
	// example:
	//
	// http://store.example.com/showpic/e76beb3bedca3fe0375cef12742ea23e
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryResultDataImages) String() string {
	return dara.Prettify(s)
}

func (s QueryResultDataImages) GoString() string {
	return s.String()
}

func (s *QueryResultDataImages) GetTitle() *string {
	return s.Title
}

func (s *QueryResultDataImages) GetUrl() *string {
	return s.Url
}

func (s *QueryResultDataImages) SetTitle(v string) *QueryResultDataImages {
	s.Title = &v
	return s
}

func (s *QueryResultDataImages) SetUrl(v string) *QueryResultDataImages {
	s.Url = &v
	return s
}

func (s *QueryResultDataImages) Validate() error {
	return dara.Validate(s)
}

type QueryResultDataMetadata struct {
	// The average spend in CNY.
	//
	// example:
	//
	// 500
	AverageSpend *string `json:"averageSpend,omitempty" xml:"averageSpend,omitempty"`
	// The business area.
	//
	// example:
	//
	// Central Area of West Lake Tourist Zone
	BusinessArea *string `json:"businessArea,omitempty" xml:"businessArea,omitempty"`
	// The business hours for today.
	//
	// example:
	//
	// 09:00-18:00
	DailyOpeningHours *string `json:"dailyOpeningHours,omitempty" xml:"dailyOpeningHours,omitempty"`
	// The main tag.
	//
	// example:
	//
	// Luxury
	MainTag *string `json:"mainTag,omitempty" xml:"mainTag,omitempty"`
	// The contact phone numbers. Multiple numbers are separated by vertical bars.
	//
	// example:
	//
	// 0571-8732****
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The rating.
	//
	// example:
	//
	// 4.7
	Score *string `json:"score,omitempty" xml:"score,omitempty"`
	// The weekly opening days.
	//
	// example:
	//
	// Monday to Sunday
	WeeklyOpeningDays *string `json:"weeklyOpeningDays,omitempty" xml:"weeklyOpeningDays,omitempty"`
}

func (s QueryResultDataMetadata) String() string {
	return dara.Prettify(s)
}

func (s QueryResultDataMetadata) GoString() string {
	return s.String()
}

func (s *QueryResultDataMetadata) GetAverageSpend() *string {
	return s.AverageSpend
}

func (s *QueryResultDataMetadata) GetBusinessArea() *string {
	return s.BusinessArea
}

func (s *QueryResultDataMetadata) GetDailyOpeningHours() *string {
	return s.DailyOpeningHours
}

func (s *QueryResultDataMetadata) GetMainTag() *string {
	return s.MainTag
}

func (s *QueryResultDataMetadata) GetPhone() *string {
	return s.Phone
}

func (s *QueryResultDataMetadata) GetScore() *string {
	return s.Score
}

func (s *QueryResultDataMetadata) GetWeeklyOpeningDays() *string {
	return s.WeeklyOpeningDays
}

func (s *QueryResultDataMetadata) SetAverageSpend(v string) *QueryResultDataMetadata {
	s.AverageSpend = &v
	return s
}

func (s *QueryResultDataMetadata) SetBusinessArea(v string) *QueryResultDataMetadata {
	s.BusinessArea = &v
	return s
}

func (s *QueryResultDataMetadata) SetDailyOpeningHours(v string) *QueryResultDataMetadata {
	s.DailyOpeningHours = &v
	return s
}

func (s *QueryResultDataMetadata) SetMainTag(v string) *QueryResultDataMetadata {
	s.MainTag = &v
	return s
}

func (s *QueryResultDataMetadata) SetPhone(v string) *QueryResultDataMetadata {
	s.Phone = &v
	return s
}

func (s *QueryResultDataMetadata) SetScore(v string) *QueryResultDataMetadata {
	s.Score = &v
	return s
}

func (s *QueryResultDataMetadata) SetWeeklyOpeningDays(v string) *QueryResultDataMetadata {
	s.WeeklyOpeningDays = &v
	return s
}

func (s *QueryResultDataMetadata) Validate() error {
	return dara.Validate(s)
}

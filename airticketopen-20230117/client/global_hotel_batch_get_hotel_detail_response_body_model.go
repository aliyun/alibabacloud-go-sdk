// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelBatchGetHotelDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelBatchGetHotelDetailResponseBodyData) *GlobalHotelBatchGetHotelDetailResponseBody
	GetData() *GlobalHotelBatchGetHotelDetailResponseBodyData
	SetErrorCode(v string) *GlobalHotelBatchGetHotelDetailResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelBatchGetHotelDetailResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelBatchGetHotelDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelBatchGetHotelDetailResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelBatchGetHotelDetailResponseBody
	GetTracerId() *string
}

type GlobalHotelBatchGetHotelDetailResponseBody struct {
	Data *GlobalHotelBatchGetHotelDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// CityCodeRequired
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 城市编码不能为空
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// traceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) GetData() *GlobalHotelBatchGetHotelDetailResponseBodyData {
	return s.Data
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) SetData(v *GlobalHotelBatchGetHotelDetailResponseBodyData) *GlobalHotelBatchGetHotelDetailResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) SetErrorCode(v string) *GlobalHotelBatchGetHotelDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) SetErrorMsg(v string) *GlobalHotelBatchGetHotelDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) SetRequestId(v string) *GlobalHotelBatchGetHotelDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) SetSuccess(v bool) *GlobalHotelBatchGetHotelDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) SetTracerId(v string) *GlobalHotelBatchGetHotelDetailResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelBatchGetHotelDetailResponseBodyData struct {
	Hotels []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotels `json:"Hotels,omitempty" xml:"Hotels,omitempty" type:"Repeated"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyData) GetHotels() []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	return s.Hotels
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyData) SetHotels(v []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) *GlobalHotelBatchGetHotelDetailResponseBodyData {
	s.Hotels = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyData) Validate() error {
	if s.Hotels != nil {
		for _, item := range s.Hotels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelBatchGetHotelDetailResponseBodyDataHotels struct {
	// example:
	//
	// No.33 East Chang An Avenue
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 14:00
	CheckInTime *string `json:"CheckInTime,omitempty" xml:"CheckInTime,omitempty"`
	// example:
	//
	// 12:00
	CheckOutTime *string `json:"CheckOutTime,omitempty" xml:"CheckOutTime,omitempty"`
	// example:
	//
	// Beijing
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// example:
	//
	// China
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	// example:
	//
	// 五星级豪华酒店
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// HOTEL_NOT_FOUND
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 酒店不存在
	ErrorMessage *string                                                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Facilities   []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities `json:"Facilities,omitempty" xml:"Facilities,omitempty" type:"Repeated"`
	// example:
	//
	// Beijing Hotel
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// example:
	//
	// 北京饭店
	HotelNameCn *string `json:"HotelNameCn,omitempty" xml:"HotelNameCn,omitempty"`
	// example:
	//
	// LUXURY
	HotelType *string `json:"HotelType,omitempty" xml:"HotelType,omitempty"`
	// example:
	//
	// 39.9042
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// 116.4074
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// example:
	//
	// 2018
	OpeningTime *int32                                                          `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	Pictures    []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures `json:"Pictures,omitempty" xml:"Pictures,omitempty" type:"Repeated"`
	Policies    []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// GOOGLE
	PositionType *string `json:"PositionType,omitempty" xml:"PositionType,omitempty"`
	// example:
	//
	// 2021
	RenovationTime *int32                                                           `json:"RenovationTime,omitempty" xml:"RenovationTime,omitempty"`
	RoomTypes      []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes `json:"RoomTypes,omitempty" xml:"RoomTypes,omitempty" type:"Repeated"`
	// example:
	//
	// H001
	StandardHotelId *string `json:"StandardHotelId,omitempty" xml:"StandardHotelId,omitempty"`
	// example:
	//
	// 5
	Star *string `json:"Star,omitempty" xml:"Star,omitempty"`
	// example:
	//
	// ONLINE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// +86-10-65137766
	Tel *string `json:"Tel,omitempty" xml:"Tel,omitempty"`
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetAddress() *string {
	return s.Address
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetCheckInTime() *string {
	return s.CheckInTime
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetCheckOutTime() *string {
	return s.CheckOutTime
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetCityName() *string {
	return s.CityName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetCountryName() *string {
	return s.CountryName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetDescription() *string {
	return s.Description
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetFacilities() []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities {
	return s.Facilities
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetHotelName() *string {
	return s.HotelName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetHotelNameCn() *string {
	return s.HotelNameCn
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetHotelType() *string {
	return s.HotelType
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetLatitude() *string {
	return s.Latitude
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetLongitude() *string {
	return s.Longitude
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetOpeningTime() *int32 {
	return s.OpeningTime
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetPictures() []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures {
	return s.Pictures
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetPolicies() []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies {
	return s.Policies
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetPositionType() *string {
	return s.PositionType
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetRenovationTime() *int32 {
	return s.RenovationTime
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetRoomTypes() []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	return s.RoomTypes
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetStandardHotelId() *string {
	return s.StandardHotelId
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetStar() *string {
	return s.Star
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetStatus() *string {
	return s.Status
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetTel() *string {
	return s.Tel
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) GetTimezone() *string {
	return s.Timezone
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetAddress(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Address = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetCheckInTime(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.CheckInTime = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetCheckOutTime(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.CheckOutTime = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetCityName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.CityName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetCountryName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.CountryName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetDescription(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Description = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetErrorCode(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetErrorMessage(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.ErrorMessage = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetFacilities(v []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Facilities = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetHotelName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.HotelName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetHotelNameCn(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.HotelNameCn = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetHotelType(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.HotelType = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetLatitude(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Latitude = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetLongitude(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Longitude = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetOpeningTime(v int32) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.OpeningTime = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetPictures(v []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Pictures = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetPolicies(v []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Policies = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetPositionType(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.PositionType = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetRenovationTime(v int32) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.RenovationTime = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetRoomTypes(v []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.RoomTypes = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetStandardHotelId(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.StandardHotelId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetStar(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Star = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetStatus(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Status = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetTel(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Tel = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) SetTimezone(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels {
	s.Timezone = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotels) Validate() error {
	if s.Facilities != nil {
		for _, item := range s.Facilities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Pictures != nil {
		for _, item := range s.Pictures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RoomTypes != nil {
		for _, item := range s.RoomTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities struct {
	// example:
	//
	// 室外恒温泳池
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// F001
	FacilityId *string `json:"FacilityId,omitempty" xml:"FacilityId,omitempty"`
	// example:
	//
	// 游泳池
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) GetDescription() *string {
	return s.Description
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) GetFacilityId() *string {
	return s.FacilityId
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) GetName() *string {
	return s.Name
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) SetDescription(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities {
	s.Description = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) SetFacilityId(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities {
	s.FacilityId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) SetName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities {
	s.Name = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsFacilities) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures struct {
	// example:
	//
	// 酒店大堂
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// HOTEL
	FirstCategoryCode *string `json:"FirstCategoryCode,omitempty" xml:"FirstCategoryCode,omitempty"`
	// example:
	//
	// 酒店
	FirstCategoryName *string `json:"FirstCategoryName,omitempty" xml:"FirstCategoryName,omitempty"`
	// example:
	//
	// true
	IsHeadPic *bool `json:"IsHeadPic,omitempty" xml:"IsHeadPic,omitempty"`
	// example:
	//
	// PIC001
	PictureId *string `json:"PictureId,omitempty" xml:"PictureId,omitempty"`
	// example:
	//
	// LOBBY
	SecondCategoryCode *string `json:"SecondCategoryCode,omitempty" xml:"SecondCategoryCode,omitempty"`
	// example:
	//
	// 大堂
	SecondCategoryName *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
	// example:
	//
	// https://img.example.com/1.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) GetDescription() *string {
	return s.Description
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) GetFirstCategoryCode() *string {
	return s.FirstCategoryCode
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) GetFirstCategoryName() *string {
	return s.FirstCategoryName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) GetIsHeadPic() *bool {
	return s.IsHeadPic
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) GetPictureId() *string {
	return s.PictureId
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) GetSecondCategoryCode() *string {
	return s.SecondCategoryCode
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) GetSecondCategoryName() *string {
	return s.SecondCategoryName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) GetUrl() *string {
	return s.Url
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) SetDescription(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.Description = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) SetFirstCategoryCode(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.FirstCategoryCode = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) SetFirstCategoryName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.FirstCategoryName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) SetIsHeadPic(v bool) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.IsHeadPic = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) SetPictureId(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.PictureId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) SetSecondCategoryCode(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.SecondCategoryCode = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) SetSecondCategoryName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.SecondCategoryName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) SetUrl(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.Url = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPictures) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies struct {
	// example:
	//
	// 入住政策
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// CHECK_IN
	GroupTypeId *string                                                              `json:"GroupTypeId,omitempty" xml:"GroupTypeId,omitempty"`
	Items       []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) GetGroupName() *string {
	return s.GroupName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) GetGroupTypeId() *string {
	return s.GroupTypeId
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) GetItems() []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	return s.Items
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) SetGroupName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies {
	s.GroupName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) SetGroupTypeId(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies {
	s.GroupTypeId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) SetItems(v []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies {
	s.Items = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPolicies) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems struct {
	Children []interface{} `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	// example:
	//
	// 入住时间
	ItemName *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	// example:
	//
	// CHECK_IN_TIME
	ItemTypeId *string `json:"ItemTypeId,omitempty" xml:"ItemTypeId,omitempty"`
	// example:
	//
	// 14:00
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GetChildren() []interface{} {
	return s.Children
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GetItemName() *string {
	return s.ItemName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GetItemTypeId() *string {
	return s.ItemTypeId
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GetValue() *string {
	return s.Value
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) SetChildren(v []interface{}) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	s.Children = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) SetItemName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	s.ItemName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) SetItemTypeId(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	s.ItemTypeId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) SetValue(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	s.Value = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes struct {
	BedType  []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType  `json:"BedType,omitempty" xml:"BedType,omitempty" type:"Repeated"`
	Pictures []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures `json:"Pictures,omitempty" xml:"Pictures,omitempty" type:"Repeated"`
	// example:
	//
	// Deluxe King Room
	RoomName *string `json:"RoomName,omitempty" xml:"RoomName,omitempty"`
	// example:
	//
	// 豪华大床房
	RoomNameCn *string `json:"RoomNameCn,omitempty" xml:"RoomNameCn,omitempty"`
	// example:
	//
	// 35.0
	RoomSize *float64 `json:"RoomSize,omitempty" xml:"RoomSize,omitempty"`
	// example:
	//
	// SQM
	RoomSizeUnit *string `json:"RoomSizeUnit,omitempty" xml:"RoomSizeUnit,omitempty"`
	// example:
	//
	// R001
	StandardRoomId *string `json:"StandardRoomId,omitempty" xml:"StandardRoomId,omitempty"`
	// example:
	//
	// WINDOW
	WindowType *string `json:"WindowType,omitempty" xml:"WindowType,omitempty"`
	// example:
	//
	// SMALL_WINDOW
	WindowTypeDefect *string `json:"WindowTypeDefect,omitempty" xml:"WindowTypeDefect,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetBedType() []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType {
	return s.BedType
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetPictures() []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	return s.Pictures
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetRoomName() *string {
	return s.RoomName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetRoomNameCn() *string {
	return s.RoomNameCn
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetRoomSize() *float64 {
	return s.RoomSize
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetRoomSizeUnit() *string {
	return s.RoomSizeUnit
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetStandardRoomId() *string {
	return s.StandardRoomId
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetWindowType() *string {
	return s.WindowType
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetWindowTypeDefect() *string {
	return s.WindowTypeDefect
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetBedType(v []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.BedType = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetPictures(v []*GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.Pictures = v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetRoomName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.RoomName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetRoomNameCn(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.RoomNameCn = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetRoomSize(v float64) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.RoomSize = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetRoomSizeUnit(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.RoomSizeUnit = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetStandardRoomId(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.StandardRoomId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetWindowType(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.WindowType = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetWindowTypeDefect(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.WindowTypeDefect = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypes) Validate() error {
	if s.BedType != nil {
		for _, item := range s.BedType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Pictures != nil {
		for _, item := range s.Pictures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType struct {
	// example:
	//
	// 1
	BedCount *int32 `json:"BedCount,omitempty" xml:"BedCount,omitempty"`
	// example:
	//
	// 1.8
	BedSize *string `json:"BedSize,omitempty" xml:"BedSize,omitempty"`
	// example:
	//
	// 大床
	BedType *string `json:"BedType,omitempty" xml:"BedType,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) GetBedCount() *int32 {
	return s.BedCount
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) GetBedSize() *string {
	return s.BedSize
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) GetBedType() *string {
	return s.BedType
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) SetBedCount(v int32) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType {
	s.BedCount = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) SetBedSize(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType {
	s.BedSize = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) SetBedType(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType {
	s.BedType = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) Validate() error {
	return dara.Validate(s)
}

type GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures struct {
	// example:
	//
	// 酒店大堂
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// HOTEL
	FirstCategoryCode *string `json:"FirstCategoryCode,omitempty" xml:"FirstCategoryCode,omitempty"`
	// example:
	//
	// 酒店
	FirstCategoryName *string `json:"FirstCategoryName,omitempty" xml:"FirstCategoryName,omitempty"`
	// example:
	//
	// true
	IsHeadPic *bool `json:"IsHeadPic,omitempty" xml:"IsHeadPic,omitempty"`
	// example:
	//
	// PIC001
	PictureId *string `json:"PictureId,omitempty" xml:"PictureId,omitempty"`
	// example:
	//
	// LOBBY
	SecondCategoryCode *string `json:"SecondCategoryCode,omitempty" xml:"SecondCategoryCode,omitempty"`
	// example:
	//
	// 大堂
	SecondCategoryName *string `json:"SecondCategoryName,omitempty" xml:"SecondCategoryName,omitempty"`
	// example:
	//
	// https://img.example.com/1.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GoString() string {
	return s.String()
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetDescription() *string {
	return s.Description
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetFirstCategoryCode() *string {
	return s.FirstCategoryCode
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetFirstCategoryName() *string {
	return s.FirstCategoryName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetIsHeadPic() *bool {
	return s.IsHeadPic
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetPictureId() *string {
	return s.PictureId
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetSecondCategoryCode() *string {
	return s.SecondCategoryCode
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetSecondCategoryName() *string {
	return s.SecondCategoryName
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetUrl() *string {
	return s.Url
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetDescription(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.Description = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetFirstCategoryCode(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.FirstCategoryCode = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetFirstCategoryName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.FirstCategoryName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetIsHeadPic(v bool) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.IsHeadPic = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetPictureId(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.PictureId = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetSecondCategoryCode(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.SecondCategoryCode = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetSecondCategoryName(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.SecondCategoryName = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetUrl(v string) *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.Url = &v
	return s
}

func (s *GlobalHotelBatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) Validate() error {
	return dara.Validate(s)
}

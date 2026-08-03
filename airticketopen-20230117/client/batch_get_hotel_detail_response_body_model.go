// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetHotelDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BatchGetHotelDetailResponseBodyData) *BatchGetHotelDetailResponseBody
	GetData() *BatchGetHotelDetailResponseBodyData
	SetErrorCode(v string) *BatchGetHotelDetailResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *BatchGetHotelDetailResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *BatchGetHotelDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchGetHotelDetailResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *BatchGetHotelDetailResponseBody
	GetTracerId() *string
}

type BatchGetHotelDetailResponseBody struct {
	Data *BatchGetHotelDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s BatchGetHotelDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBody) GetData() *BatchGetHotelDetailResponseBodyData {
	return s.Data
}

func (s *BatchGetHotelDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchGetHotelDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *BatchGetHotelDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetHotelDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchGetHotelDetailResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *BatchGetHotelDetailResponseBody) SetData(v *BatchGetHotelDetailResponseBodyData) *BatchGetHotelDetailResponseBody {
	s.Data = v
	return s
}

func (s *BatchGetHotelDetailResponseBody) SetErrorCode(v string) *BatchGetHotelDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchGetHotelDetailResponseBody) SetErrorMsg(v string) *BatchGetHotelDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *BatchGetHotelDetailResponseBody) SetRequestId(v string) *BatchGetHotelDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetHotelDetailResponseBody) SetSuccess(v bool) *BatchGetHotelDetailResponseBody {
	s.Success = &v
	return s
}

func (s *BatchGetHotelDetailResponseBody) SetTracerId(v string) *BatchGetHotelDetailResponseBody {
	s.TracerId = &v
	return s
}

func (s *BatchGetHotelDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetHotelDetailResponseBodyData struct {
	Hotels []*BatchGetHotelDetailResponseBodyDataHotels `json:"Hotels,omitempty" xml:"Hotels,omitempty" type:"Repeated"`
}

func (s BatchGetHotelDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBodyData) GetHotels() []*BatchGetHotelDetailResponseBodyDataHotels {
	return s.Hotels
}

func (s *BatchGetHotelDetailResponseBodyData) SetHotels(v []*BatchGetHotelDetailResponseBodyDataHotels) *BatchGetHotelDetailResponseBodyData {
	s.Hotels = v
	return s
}

func (s *BatchGetHotelDetailResponseBodyData) Validate() error {
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

type BatchGetHotelDetailResponseBodyDataHotels struct {
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
	ErrorMessage *string                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Facilities   []*BatchGetHotelDetailResponseBodyDataHotelsFacilities `json:"Facilities,omitempty" xml:"Facilities,omitempty" type:"Repeated"`
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
	OpeningTime *int32                                               `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	Pictures    []*BatchGetHotelDetailResponseBodyDataHotelsPictures `json:"Pictures,omitempty" xml:"Pictures,omitempty" type:"Repeated"`
	Policies    []*BatchGetHotelDetailResponseBodyDataHotelsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// GOOGLE
	PositionType *string `json:"PositionType,omitempty" xml:"PositionType,omitempty"`
	// example:
	//
	// 2021
	RenovationTime *int32                                                `json:"RenovationTime,omitempty" xml:"RenovationTime,omitempty"`
	RoomTypes      []*BatchGetHotelDetailResponseBodyDataHotelsRoomTypes `json:"RoomTypes,omitempty" xml:"RoomTypes,omitempty" type:"Repeated"`
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

func (s BatchGetHotelDetailResponseBodyDataHotels) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBodyDataHotels) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetAddress() *string {
	return s.Address
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetCheckInTime() *string {
	return s.CheckInTime
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetCheckOutTime() *string {
	return s.CheckOutTime
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetCityName() *string {
	return s.CityName
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetCountryName() *string {
	return s.CountryName
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetDescription() *string {
	return s.Description
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetFacilities() []*BatchGetHotelDetailResponseBodyDataHotelsFacilities {
	return s.Facilities
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetHotelName() *string {
	return s.HotelName
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetHotelNameCn() *string {
	return s.HotelNameCn
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetHotelType() *string {
	return s.HotelType
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetLatitude() *string {
	return s.Latitude
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetLongitude() *string {
	return s.Longitude
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetOpeningTime() *int32 {
	return s.OpeningTime
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetPictures() []*BatchGetHotelDetailResponseBodyDataHotelsPictures {
	return s.Pictures
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetPolicies() []*BatchGetHotelDetailResponseBodyDataHotelsPolicies {
	return s.Policies
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetPositionType() *string {
	return s.PositionType
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetRenovationTime() *int32 {
	return s.RenovationTime
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetRoomTypes() []*BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	return s.RoomTypes
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetStandardHotelId() *string {
	return s.StandardHotelId
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetStar() *string {
	return s.Star
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetStatus() *string {
	return s.Status
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetTel() *string {
	return s.Tel
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) GetTimezone() *string {
	return s.Timezone
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetAddress(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Address = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetCheckInTime(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.CheckInTime = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetCheckOutTime(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.CheckOutTime = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetCityName(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.CityName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetCountryName(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.CountryName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetDescription(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Description = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetErrorCode(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.ErrorCode = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetErrorMessage(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.ErrorMessage = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetFacilities(v []*BatchGetHotelDetailResponseBodyDataHotelsFacilities) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Facilities = v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetHotelName(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.HotelName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetHotelNameCn(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.HotelNameCn = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetHotelType(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.HotelType = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetLatitude(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Latitude = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetLongitude(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Longitude = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetOpeningTime(v int32) *BatchGetHotelDetailResponseBodyDataHotels {
	s.OpeningTime = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetPictures(v []*BatchGetHotelDetailResponseBodyDataHotelsPictures) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Pictures = v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetPolicies(v []*BatchGetHotelDetailResponseBodyDataHotelsPolicies) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Policies = v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetPositionType(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.PositionType = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetRenovationTime(v int32) *BatchGetHotelDetailResponseBodyDataHotels {
	s.RenovationTime = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetRoomTypes(v []*BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) *BatchGetHotelDetailResponseBodyDataHotels {
	s.RoomTypes = v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetStandardHotelId(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.StandardHotelId = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetStar(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Star = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetStatus(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Status = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetTel(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Tel = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) SetTimezone(v string) *BatchGetHotelDetailResponseBodyDataHotels {
	s.Timezone = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotels) Validate() error {
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

type BatchGetHotelDetailResponseBodyDataHotelsFacilities struct {
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

func (s BatchGetHotelDetailResponseBodyDataHotelsFacilities) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBodyDataHotelsFacilities) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsFacilities) GetDescription() *string {
	return s.Description
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsFacilities) GetFacilityId() *string {
	return s.FacilityId
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsFacilities) GetName() *string {
	return s.Name
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsFacilities) SetDescription(v string) *BatchGetHotelDetailResponseBodyDataHotelsFacilities {
	s.Description = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsFacilities) SetFacilityId(v string) *BatchGetHotelDetailResponseBodyDataHotelsFacilities {
	s.FacilityId = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsFacilities) SetName(v string) *BatchGetHotelDetailResponseBodyDataHotelsFacilities {
	s.Name = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsFacilities) Validate() error {
	return dara.Validate(s)
}

type BatchGetHotelDetailResponseBodyDataHotelsPictures struct {
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
	IsHeadPic         *bool   `json:"IsHeadPic,omitempty" xml:"IsHeadPic,omitempty"`
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

func (s BatchGetHotelDetailResponseBodyDataHotelsPictures) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBodyDataHotelsPictures) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) GetDescription() *string {
	return s.Description
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) GetFirstCategoryCode() *string {
	return s.FirstCategoryCode
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) GetFirstCategoryName() *string {
	return s.FirstCategoryName
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) GetIsHeadPic() *bool {
	return s.IsHeadPic
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) GetPictureId() *string {
	return s.PictureId
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) GetSecondCategoryCode() *string {
	return s.SecondCategoryCode
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) GetSecondCategoryName() *string {
	return s.SecondCategoryName
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) GetUrl() *string {
	return s.Url
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) SetDescription(v string) *BatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.Description = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) SetFirstCategoryCode(v string) *BatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.FirstCategoryCode = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) SetFirstCategoryName(v string) *BatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.FirstCategoryName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) SetIsHeadPic(v bool) *BatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.IsHeadPic = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) SetPictureId(v string) *BatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.PictureId = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) SetSecondCategoryCode(v string) *BatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.SecondCategoryCode = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) SetSecondCategoryName(v string) *BatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.SecondCategoryName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) SetUrl(v string) *BatchGetHotelDetailResponseBodyDataHotelsPictures {
	s.Url = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPictures) Validate() error {
	return dara.Validate(s)
}

type BatchGetHotelDetailResponseBodyDataHotelsPolicies struct {
	// example:
	//
	// 入住政策
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// CHECK_IN
	GroupTypeId *string                                                   `json:"GroupTypeId,omitempty" xml:"GroupTypeId,omitempty"`
	Items       []*BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s BatchGetHotelDetailResponseBodyDataHotelsPolicies) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBodyDataHotelsPolicies) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPolicies) GetGroupName() *string {
	return s.GroupName
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPolicies) GetGroupTypeId() *string {
	return s.GroupTypeId
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPolicies) GetItems() []*BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	return s.Items
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPolicies) SetGroupName(v string) *BatchGetHotelDetailResponseBodyDataHotelsPolicies {
	s.GroupName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPolicies) SetGroupTypeId(v string) *BatchGetHotelDetailResponseBodyDataHotelsPolicies {
	s.GroupTypeId = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPolicies) SetItems(v []*BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) *BatchGetHotelDetailResponseBodyDataHotelsPolicies {
	s.Items = v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPolicies) Validate() error {
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

type BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems struct {
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

func (s BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GetChildren() []interface{} {
	return s.Children
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GetItemName() *string {
	return s.ItemName
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GetItemTypeId() *string {
	return s.ItemTypeId
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) GetValue() *string {
	return s.Value
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) SetChildren(v []interface{}) *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	s.Children = v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) SetItemName(v string) *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	s.ItemName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) SetItemTypeId(v string) *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	s.ItemTypeId = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) SetValue(v string) *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems {
	s.Value = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsPoliciesItems) Validate() error {
	return dara.Validate(s)
}

type BatchGetHotelDetailResponseBodyDataHotelsRoomTypes struct {
	BedType  []*BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType  `json:"BedType,omitempty" xml:"BedType,omitempty" type:"Repeated"`
	Pictures []*BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures `json:"Pictures,omitempty" xml:"Pictures,omitempty" type:"Repeated"`
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
	// standardRoomId
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

func (s BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetBedType() []*BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType {
	return s.BedType
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetPictures() []*BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	return s.Pictures
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetRoomName() *string {
	return s.RoomName
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetRoomNameCn() *string {
	return s.RoomNameCn
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetRoomSize() *float64 {
	return s.RoomSize
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetRoomSizeUnit() *string {
	return s.RoomSizeUnit
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetStandardRoomId() *string {
	return s.StandardRoomId
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetWindowType() *string {
	return s.WindowType
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) GetWindowTypeDefect() *string {
	return s.WindowTypeDefect
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetBedType(v []*BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.BedType = v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetPictures(v []*BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.Pictures = v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetRoomName(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.RoomName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetRoomNameCn(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.RoomNameCn = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetRoomSize(v float64) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.RoomSize = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetRoomSizeUnit(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.RoomSizeUnit = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetStandardRoomId(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.StandardRoomId = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetWindowType(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.WindowType = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) SetWindowTypeDefect(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes {
	s.WindowTypeDefect = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypes) Validate() error {
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

type BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType struct {
	// example:
	//
	// 1
	BedCount *int32 `json:"BedCount,omitempty" xml:"BedCount,omitempty"`
	// example:
	//
	// 1.9
	BedSize *string `json:"BedSize,omitempty" xml:"BedSize,omitempty"`
	// example:
	//
	// 大床
	BedType *string `json:"BedType,omitempty" xml:"BedType,omitempty"`
}

func (s BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) GetBedCount() *int32 {
	return s.BedCount
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) GetBedSize() *string {
	return s.BedSize
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) GetBedType() *string {
	return s.BedType
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) SetBedCount(v int32) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType {
	s.BedCount = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) SetBedSize(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType {
	s.BedSize = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) SetBedType(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType {
	s.BedType = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesBedType) Validate() error {
	return dara.Validate(s)
}

type BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures struct {
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
	IsHeadPic         *bool   `json:"IsHeadPic,omitempty" xml:"IsHeadPic,omitempty"`
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

func (s BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) String() string {
	return dara.Prettify(s)
}

func (s BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GoString() string {
	return s.String()
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetDescription() *string {
	return s.Description
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetFirstCategoryCode() *string {
	return s.FirstCategoryCode
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetFirstCategoryName() *string {
	return s.FirstCategoryName
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetIsHeadPic() *bool {
	return s.IsHeadPic
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetPictureId() *string {
	return s.PictureId
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetSecondCategoryCode() *string {
	return s.SecondCategoryCode
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetSecondCategoryName() *string {
	return s.SecondCategoryName
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) GetUrl() *string {
	return s.Url
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetDescription(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.Description = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetFirstCategoryCode(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.FirstCategoryCode = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetFirstCategoryName(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.FirstCategoryName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetIsHeadPic(v bool) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.IsHeadPic = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetPictureId(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.PictureId = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetSecondCategoryCode(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.SecondCategoryCode = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetSecondCategoryName(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.SecondCategoryName = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) SetUrl(v string) *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures {
	s.Url = &v
	return s
}

func (s *BatchGetHotelDetailResponseBodyDataHotelsRoomTypesPictures) Validate() error {
	return dara.Validate(s)
}

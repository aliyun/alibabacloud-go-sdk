// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketPageQueryScenicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketPageQueryScenicResponseBodyData) *TicketPageQueryScenicResponseBody
	GetData() *TicketPageQueryScenicResponseBodyData
	SetErrorCode(v string) *TicketPageQueryScenicResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketPageQueryScenicResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketPageQueryScenicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketPageQueryScenicResponseBody
	GetSuccess() *bool
}

type TicketPageQueryScenicResponseBody struct {
	Data *TicketPageQueryScenicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ScenicIdInvalid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// ScenicId不合法
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TicketPageQueryScenicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryScenicResponseBody) GoString() string {
	return s.String()
}

func (s *TicketPageQueryScenicResponseBody) GetData() *TicketPageQueryScenicResponseBodyData {
	return s.Data
}

func (s *TicketPageQueryScenicResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketPageQueryScenicResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketPageQueryScenicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketPageQueryScenicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketPageQueryScenicResponseBody) SetData(v *TicketPageQueryScenicResponseBodyData) *TicketPageQueryScenicResponseBody {
	s.Data = v
	return s
}

func (s *TicketPageQueryScenicResponseBody) SetErrorCode(v string) *TicketPageQueryScenicResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketPageQueryScenicResponseBody) SetErrorMsg(v string) *TicketPageQueryScenicResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketPageQueryScenicResponseBody) SetRequestId(v string) *TicketPageQueryScenicResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketPageQueryScenicResponseBody) SetSuccess(v bool) *TicketPageQueryScenicResponseBody {
	s.Success = &v
	return s
}

func (s *TicketPageQueryScenicResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketPageQueryScenicResponseBodyData struct {
	Scenics []*TicketPageQueryScenicResponseBodyDataScenics `json:"Scenics,omitempty" xml:"Scenics,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s TicketPageQueryScenicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryScenicResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketPageQueryScenicResponseBodyData) GetScenics() []*TicketPageQueryScenicResponseBodyDataScenics {
	return s.Scenics
}

func (s *TicketPageQueryScenicResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *TicketPageQueryScenicResponseBodyData) SetScenics(v []*TicketPageQueryScenicResponseBodyDataScenics) *TicketPageQueryScenicResponseBodyData {
	s.Scenics = v
	return s
}

func (s *TicketPageQueryScenicResponseBodyData) SetTotalSize(v int64) *TicketPageQueryScenicResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyData) Validate() error {
	if s.Scenics != nil {
		for _, item := range s.Scenics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketPageQueryScenicResponseBodyDataScenics struct {
	// example:
	//
	// 杭州市西湖区龙井路1号
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 杭州市
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 中国
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// 国家5A级旅游景区
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ["https://example.com/img1.jpg"]
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 31.138026
	Latitude *float64 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// AAAAA
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 121.658793
	Longitude *float64 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	// example:
	//
	// 西湖风景区
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 08:00-17:30
	OpeningTime *string `json:"OpeningTime,omitempty" xml:"OpeningTime,omitempty"`
	// example:
	//
	// 0571-12345678
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// example:
	//
	// [{"type":"儿童","typeDesc":"3周岁(含)至11周岁(含)享受优惠票"},{"type":"老年人","typeDesc":"65周岁(含)以上享受优惠票"}]
	PreferentialPolicy *string `json:"PreferentialPolicy,omitempty" xml:"PreferentialPolicy,omitempty"`
	// example:
	//
	// 浙江省
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// 1天
	ResidenceTime *string `json:"ResidenceTime,omitempty" xml:"ResidenceTime,omitempty"`
	// example:
	//
	// 123456
	ScenicId *int64 `json:"ScenicId,omitempty" xml:"ScenicId,omitempty"`
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
}

func (s TicketPageQueryScenicResponseBodyDataScenics) String() string {
	return dara.Prettify(s)
}

func (s TicketPageQueryScenicResponseBodyDataScenics) GoString() string {
	return s.String()
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetAddress() *string {
	return s.Address
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetCity() *string {
	return s.City
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetCountry() *string {
	return s.Country
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetDescription() *string {
	return s.Description
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetImages() []*string {
	return s.Images
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetLatitude() *float64 {
	return s.Latitude
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetLevel() *string {
	return s.Level
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetLongitude() *float64 {
	return s.Longitude
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetName() *string {
	return s.Name
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetOpeningTime() *string {
	return s.OpeningTime
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetPhone() *string {
	return s.Phone
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetPreferentialPolicy() *string {
	return s.PreferentialPolicy
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetProvince() *string {
	return s.Province
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetResidenceTime() *string {
	return s.ResidenceTime
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetScenicId() *int64 {
	return s.ScenicId
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) GetTimezone() *string {
	return s.Timezone
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetAddress(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Address = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetCity(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.City = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetCountry(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Country = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetDescription(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Description = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetImages(v []*string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Images = v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetLatitude(v float64) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Latitude = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetLevel(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Level = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetLongitude(v float64) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Longitude = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetName(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Name = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetOpeningTime(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.OpeningTime = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetPhone(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Phone = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetPreferentialPolicy(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.PreferentialPolicy = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetProvince(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Province = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetResidenceTime(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.ResidenceTime = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetScenicId(v int64) *TicketPageQueryScenicResponseBodyDataScenics {
	s.ScenicId = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) SetTimezone(v string) *TicketPageQueryScenicResponseBodyDataScenics {
	s.Timezone = &v
	return s
}

func (s *TicketPageQueryScenicResponseBodyDataScenics) Validate() error {
	return dara.Validate(s)
}

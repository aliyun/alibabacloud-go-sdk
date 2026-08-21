// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryScenicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketQueryScenicResponseBodyData) *TicketQueryScenicResponseBody
	GetData() *TicketQueryScenicResponseBodyData
	SetErrorCode(v string) *TicketQueryScenicResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketQueryScenicResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketQueryScenicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketQueryScenicResponseBody
	GetSuccess() *bool
}

type TicketQueryScenicResponseBody struct {
	Data *TicketQueryScenicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s TicketQueryScenicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryScenicResponseBody) GoString() string {
	return s.String()
}

func (s *TicketQueryScenicResponseBody) GetData() *TicketQueryScenicResponseBodyData {
	return s.Data
}

func (s *TicketQueryScenicResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketQueryScenicResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketQueryScenicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketQueryScenicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketQueryScenicResponseBody) SetData(v *TicketQueryScenicResponseBodyData) *TicketQueryScenicResponseBody {
	s.Data = v
	return s
}

func (s *TicketQueryScenicResponseBody) SetErrorCode(v string) *TicketQueryScenicResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketQueryScenicResponseBody) SetErrorMsg(v string) *TicketQueryScenicResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketQueryScenicResponseBody) SetRequestId(v string) *TicketQueryScenicResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketQueryScenicResponseBody) SetSuccess(v bool) *TicketQueryScenicResponseBody {
	s.Success = &v
	return s
}

func (s *TicketQueryScenicResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryScenicResponseBodyData struct {
	Scenic *TicketQueryScenicResponseBodyDataScenic `json:"Scenic,omitempty" xml:"Scenic,omitempty" type:"Struct"`
}

func (s TicketQueryScenicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryScenicResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketQueryScenicResponseBodyData) GetScenic() *TicketQueryScenicResponseBodyDataScenic {
	return s.Scenic
}

func (s *TicketQueryScenicResponseBodyData) SetScenic(v *TicketQueryScenicResponseBodyDataScenic) *TicketQueryScenicResponseBodyData {
	s.Scenic = v
	return s
}

func (s *TicketQueryScenicResponseBodyData) Validate() error {
	if s.Scenic != nil {
		if err := s.Scenic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryScenicResponseBodyDataScenic struct {
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

func (s TicketQueryScenicResponseBodyDataScenic) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryScenicResponseBodyDataScenic) GoString() string {
	return s.String()
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetAddress() *string {
	return s.Address
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetCity() *string {
	return s.City
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetCountry() *string {
	return s.Country
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetDescription() *string {
	return s.Description
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetImages() []*string {
	return s.Images
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetLatitude() *float64 {
	return s.Latitude
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetLevel() *string {
	return s.Level
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetLongitude() *float64 {
	return s.Longitude
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetName() *string {
	return s.Name
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetOpeningTime() *string {
	return s.OpeningTime
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetPhone() *string {
	return s.Phone
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetPreferentialPolicy() *string {
	return s.PreferentialPolicy
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetProvince() *string {
	return s.Province
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetResidenceTime() *string {
	return s.ResidenceTime
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetScenicId() *int64 {
	return s.ScenicId
}

func (s *TicketQueryScenicResponseBodyDataScenic) GetTimezone() *string {
	return s.Timezone
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetAddress(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.Address = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetCity(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.City = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetCountry(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.Country = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetDescription(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.Description = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetImages(v []*string) *TicketQueryScenicResponseBodyDataScenic {
	s.Images = v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetLatitude(v float64) *TicketQueryScenicResponseBodyDataScenic {
	s.Latitude = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetLevel(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.Level = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetLongitude(v float64) *TicketQueryScenicResponseBodyDataScenic {
	s.Longitude = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetName(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.Name = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetOpeningTime(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.OpeningTime = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetPhone(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.Phone = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetPreferentialPolicy(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.PreferentialPolicy = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetProvince(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.Province = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetResidenceTime(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.ResidenceTime = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetScenicId(v int64) *TicketQueryScenicResponseBodyDataScenic {
	s.ScenicId = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) SetTimezone(v string) *TicketQueryScenicResponseBodyDataScenic {
	s.Timezone = &v
	return s
}

func (s *TicketQueryScenicResponseBodyDataScenic) Validate() error {
	return dara.Validate(s)
}

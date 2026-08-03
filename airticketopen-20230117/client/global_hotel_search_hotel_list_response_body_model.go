// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelSearchHotelListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelSearchHotelListResponseBodyData) *GlobalHotelSearchHotelListResponseBody
	GetData() *GlobalHotelSearchHotelListResponseBodyData
	SetErrorCode(v string) *GlobalHotelSearchHotelListResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelSearchHotelListResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelSearchHotelListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelSearchHotelListResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelSearchHotelListResponseBody
	GetTracerId() *string
}

type GlobalHotelSearchHotelListResponseBody struct {
	Data *GlobalHotelSearchHotelListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GlobalHotelSearchHotelListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchHotelListResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchHotelListResponseBody) GetData() *GlobalHotelSearchHotelListResponseBodyData {
	return s.Data
}

func (s *GlobalHotelSearchHotelListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelSearchHotelListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelSearchHotelListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelSearchHotelListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelSearchHotelListResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelSearchHotelListResponseBody) SetData(v *GlobalHotelSearchHotelListResponseBodyData) *GlobalHotelSearchHotelListResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBody) SetErrorCode(v string) *GlobalHotelSearchHotelListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBody) SetErrorMsg(v string) *GlobalHotelSearchHotelListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBody) SetRequestId(v string) *GlobalHotelSearchHotelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBody) SetSuccess(v bool) *GlobalHotelSearchHotelListResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBody) SetTracerId(v string) *GlobalHotelSearchHotelListResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelSearchHotelListResponseBodyData struct {
	Hotels []*GlobalHotelSearchHotelListResponseBodyDataHotels `json:"Hotels,omitempty" xml:"Hotels,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GlobalHotelSearchHotelListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchHotelListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchHotelListResponseBodyData) GetHotels() []*GlobalHotelSearchHotelListResponseBodyDataHotels {
	return s.Hotels
}

func (s *GlobalHotelSearchHotelListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GlobalHotelSearchHotelListResponseBodyData) SetHotels(v []*GlobalHotelSearchHotelListResponseBodyDataHotels) *GlobalHotelSearchHotelListResponseBodyData {
	s.Hotels = v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBodyData) SetTotal(v int32) *GlobalHotelSearchHotelListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBodyData) Validate() error {
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

type GlobalHotelSearchHotelListResponseBodyDataHotels struct {
	// example:
	//
	// 北京市
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// example:
	//
	// 中国
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	// example:
	//
	// 北京饭店
	HotelName *string `json:"HotelName,omitempty" xml:"HotelName,omitempty"`
	// example:
	//
	// H001
	StandardHotelId *string `json:"StandardHotelId,omitempty" xml:"StandardHotelId,omitempty"`
	// example:
	//
	// ONLINE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GlobalHotelSearchHotelListResponseBodyDataHotels) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchHotelListResponseBodyDataHotels) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) GetCityName() *string {
	return s.CityName
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) GetCountryName() *string {
	return s.CountryName
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) GetHotelName() *string {
	return s.HotelName
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) GetStandardHotelId() *string {
	return s.StandardHotelId
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) GetStatus() *string {
	return s.Status
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) SetCityName(v string) *GlobalHotelSearchHotelListResponseBodyDataHotels {
	s.CityName = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) SetCountryName(v string) *GlobalHotelSearchHotelListResponseBodyDataHotels {
	s.CountryName = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) SetHotelName(v string) *GlobalHotelSearchHotelListResponseBodyDataHotels {
	s.HotelName = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) SetStandardHotelId(v string) *GlobalHotelSearchHotelListResponseBodyDataHotels {
	s.StandardHotelId = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) SetStatus(v string) *GlobalHotelSearchHotelListResponseBodyDataHotels {
	s.Status = &v
	return s
}

func (s *GlobalHotelSearchHotelListResponseBodyDataHotels) Validate() error {
	return dara.Validate(s)
}

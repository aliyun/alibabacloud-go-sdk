// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchHotelListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SearchHotelListResponseBodyData) *SearchHotelListResponseBody
	GetData() *SearchHotelListResponseBodyData
	SetErrorCode(v string) *SearchHotelListResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SearchHotelListResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *SearchHotelListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchHotelListResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *SearchHotelListResponseBody
	GetTracerId() *string
}

type SearchHotelListResponseBody struct {
	Data *SearchHotelListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SearchHotelListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchHotelListResponseBody) GoString() string {
	return s.String()
}

func (s *SearchHotelListResponseBody) GetData() *SearchHotelListResponseBodyData {
	return s.Data
}

func (s *SearchHotelListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchHotelListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SearchHotelListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchHotelListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchHotelListResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *SearchHotelListResponseBody) SetData(v *SearchHotelListResponseBodyData) *SearchHotelListResponseBody {
	s.Data = v
	return s
}

func (s *SearchHotelListResponseBody) SetErrorCode(v string) *SearchHotelListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchHotelListResponseBody) SetErrorMsg(v string) *SearchHotelListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SearchHotelListResponseBody) SetRequestId(v string) *SearchHotelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchHotelListResponseBody) SetSuccess(v bool) *SearchHotelListResponseBody {
	s.Success = &v
	return s
}

func (s *SearchHotelListResponseBody) SetTracerId(v string) *SearchHotelListResponseBody {
	s.TracerId = &v
	return s
}

func (s *SearchHotelListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchHotelListResponseBodyData struct {
	Hotels []*SearchHotelListResponseBodyDataHotels `json:"Hotels,omitempty" xml:"Hotels,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchHotelListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchHotelListResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchHotelListResponseBodyData) GetHotels() []*SearchHotelListResponseBodyDataHotels {
	return s.Hotels
}

func (s *SearchHotelListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *SearchHotelListResponseBodyData) SetHotels(v []*SearchHotelListResponseBodyDataHotels) *SearchHotelListResponseBodyData {
	s.Hotels = v
	return s
}

func (s *SearchHotelListResponseBodyData) SetTotal(v int32) *SearchHotelListResponseBodyData {
	s.Total = &v
	return s
}

func (s *SearchHotelListResponseBodyData) Validate() error {
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

type SearchHotelListResponseBodyDataHotels struct {
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

func (s SearchHotelListResponseBodyDataHotels) String() string {
	return dara.Prettify(s)
}

func (s SearchHotelListResponseBodyDataHotels) GoString() string {
	return s.String()
}

func (s *SearchHotelListResponseBodyDataHotels) GetCityName() *string {
	return s.CityName
}

func (s *SearchHotelListResponseBodyDataHotels) GetCountryName() *string {
	return s.CountryName
}

func (s *SearchHotelListResponseBodyDataHotels) GetHotelName() *string {
	return s.HotelName
}

func (s *SearchHotelListResponseBodyDataHotels) GetStandardHotelId() *string {
	return s.StandardHotelId
}

func (s *SearchHotelListResponseBodyDataHotels) GetStatus() *string {
	return s.Status
}

func (s *SearchHotelListResponseBodyDataHotels) SetCityName(v string) *SearchHotelListResponseBodyDataHotels {
	s.CityName = &v
	return s
}

func (s *SearchHotelListResponseBodyDataHotels) SetCountryName(v string) *SearchHotelListResponseBodyDataHotels {
	s.CountryName = &v
	return s
}

func (s *SearchHotelListResponseBodyDataHotels) SetHotelName(v string) *SearchHotelListResponseBodyDataHotels {
	s.HotelName = &v
	return s
}

func (s *SearchHotelListResponseBodyDataHotels) SetStandardHotelId(v string) *SearchHotelListResponseBodyDataHotels {
	s.StandardHotelId = &v
	return s
}

func (s *SearchHotelListResponseBodyDataHotels) SetStatus(v string) *SearchHotelListResponseBodyDataHotels {
	s.Status = &v
	return s
}

func (s *SearchHotelListResponseBodyDataHotels) Validate() error {
	return dara.Validate(s)
}

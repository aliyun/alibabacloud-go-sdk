// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelSearchCityPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GlobalHotelSearchCityPageResponseBodyData) *GlobalHotelSearchCityPageResponseBody
	GetData() *GlobalHotelSearchCityPageResponseBodyData
	SetErrorCode(v string) *GlobalHotelSearchCityPageResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GlobalHotelSearchCityPageResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GlobalHotelSearchCityPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GlobalHotelSearchCityPageResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *GlobalHotelSearchCityPageResponseBody
	GetTracerId() *string
}

type GlobalHotelSearchCityPageResponseBody struct {
	// The business data.
	Data *GlobalHotelSearchCityPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// CityCodeRequired
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// City code cannot be empty
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 260E4F99-983D-1919-834C-5C42E98E5B2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// TraceId
	//
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelSearchCityPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchCityPageResponseBody) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchCityPageResponseBody) GetData() *GlobalHotelSearchCityPageResponseBodyData {
	return s.Data
}

func (s *GlobalHotelSearchCityPageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GlobalHotelSearchCityPageResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GlobalHotelSearchCityPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GlobalHotelSearchCityPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GlobalHotelSearchCityPageResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelSearchCityPageResponseBody) SetData(v *GlobalHotelSearchCityPageResponseBodyData) *GlobalHotelSearchCityPageResponseBody {
	s.Data = v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBody) SetErrorCode(v string) *GlobalHotelSearchCityPageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBody) SetErrorMsg(v string) *GlobalHotelSearchCityPageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBody) SetRequestId(v string) *GlobalHotelSearchCityPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBody) SetSuccess(v bool) *GlobalHotelSearchCityPageResponseBody {
	s.Success = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBody) SetTracerId(v string) *GlobalHotelSearchCityPageResponseBody {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalHotelSearchCityPageResponseBodyData struct {
	// The list of cities.
	//
	// example:
	//
	// []
	Cities []*GlobalHotelSearchCityPageResponseBodyDataCities `json:"Cities,omitempty" xml:"Cities,omitempty" type:"Repeated"`
	// Indicates whether there is a next page.
	//
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GlobalHotelSearchCityPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchCityPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchCityPageResponseBodyData) GetCities() []*GlobalHotelSearchCityPageResponseBodyDataCities {
	return s.Cities
}

func (s *GlobalHotelSearchCityPageResponseBodyData) GetHasNext() *bool {
	return s.HasNext
}

func (s *GlobalHotelSearchCityPageResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GlobalHotelSearchCityPageResponseBodyData) SetCities(v []*GlobalHotelSearchCityPageResponseBodyDataCities) *GlobalHotelSearchCityPageResponseBodyData {
	s.Cities = v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyData) SetHasNext(v bool) *GlobalHotelSearchCityPageResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyData) SetTotal(v int32) *GlobalHotelSearchCityPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyData) Validate() error {
	if s.Cities != nil {
		for _, item := range s.Cities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GlobalHotelSearchCityPageResponseBodyDataCities struct {
	// The Chinese name.
	//
	// example:
	//
	// 北京市
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	// The city code.
	//
	// example:
	//
	// 110100
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The country code.
	//
	// example:
	//
	// 156
	Country *int32 `json:"Country,omitempty" xml:"Country,omitempty"`
	// The country code (ISO 3166-1 alpha-2).
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The English name.
	//
	// example:
	//
	// Beijing
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// The administrative level.
	//
	// example:
	//
	// 3
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The parent city code.
	//
	// example:
	//
	// 110000
	ParentCode *int32 `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
	// The region.
	//
	// example:
	//
	// 1
	Region *int32 `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GlobalHotelSearchCityPageResponseBodyDataCities) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelSearchCityPageResponseBodyDataCities) GoString() string {
	return s.String()
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) GetCnName() *string {
	return s.CnName
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) GetCode() *int32 {
	return s.Code
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) GetCountry() *int32 {
	return s.Country
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) GetCountryCode() *string {
	return s.CountryCode
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) GetEnName() *string {
	return s.EnName
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) GetLevel() *int32 {
	return s.Level
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) GetParentCode() *int32 {
	return s.ParentCode
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) GetRegion() *int32 {
	return s.Region
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) SetCnName(v string) *GlobalHotelSearchCityPageResponseBodyDataCities {
	s.CnName = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) SetCode(v int32) *GlobalHotelSearchCityPageResponseBodyDataCities {
	s.Code = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) SetCountry(v int32) *GlobalHotelSearchCityPageResponseBodyDataCities {
	s.Country = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) SetCountryCode(v string) *GlobalHotelSearchCityPageResponseBodyDataCities {
	s.CountryCode = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) SetEnName(v string) *GlobalHotelSearchCityPageResponseBodyDataCities {
	s.EnName = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) SetLevel(v int32) *GlobalHotelSearchCityPageResponseBodyDataCities {
	s.Level = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) SetParentCode(v int32) *GlobalHotelSearchCityPageResponseBodyDataCities {
	s.ParentCode = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) SetRegion(v int32) *GlobalHotelSearchCityPageResponseBodyDataCities {
	s.Region = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) Validate() error {
	return dara.Validate(s)
}

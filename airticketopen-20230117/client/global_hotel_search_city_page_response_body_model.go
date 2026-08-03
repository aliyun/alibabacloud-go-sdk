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
	Data *GlobalHotelSearchCityPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// example:
	//
	// []
	Cities []*GlobalHotelSearchCityPageResponseBodyDataCities `json:"Cities,omitempty" xml:"Cities,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
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
	// example:
	//
	// 北京市
	CnName *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	// example:
	//
	// 110100
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 156
	Country *int32 `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// example:
	//
	// Beijing
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// example:
	//
	// 3
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 110000
	ParentCode *int32 `json:"ParentCode,omitempty" xml:"ParentCode,omitempty"`
	// example:
	//
	// 1
	Region *int32 `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) GetType() *int32 {
	return s.Type
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

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) SetType(v int32) *GlobalHotelSearchCityPageResponseBodyDataCities {
	s.Type = &v
	return s
}

func (s *GlobalHotelSearchCityPageResponseBodyDataCities) Validate() error {
	return dara.Validate(s)
}

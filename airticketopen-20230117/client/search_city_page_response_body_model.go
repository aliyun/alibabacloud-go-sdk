// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCityPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SearchCityPageResponseBodyData) *SearchCityPageResponseBody
	GetData() *SearchCityPageResponseBodyData
	SetErrorCode(v string) *SearchCityPageResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SearchCityPageResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *SearchCityPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchCityPageResponseBody
	GetSuccess() *bool
	SetTracerId(v string) *SearchCityPageResponseBody
	GetTracerId() *string
}

type SearchCityPageResponseBody struct {
	Data *SearchCityPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SearchCityPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchCityPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchCityPageResponseBody) GetData() *SearchCityPageResponseBodyData {
	return s.Data
}

func (s *SearchCityPageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchCityPageResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SearchCityPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchCityPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchCityPageResponseBody) GetTracerId() *string {
	return s.TracerId
}

func (s *SearchCityPageResponseBody) SetData(v *SearchCityPageResponseBodyData) *SearchCityPageResponseBody {
	s.Data = v
	return s
}

func (s *SearchCityPageResponseBody) SetErrorCode(v string) *SearchCityPageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchCityPageResponseBody) SetErrorMsg(v string) *SearchCityPageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SearchCityPageResponseBody) SetRequestId(v string) *SearchCityPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchCityPageResponseBody) SetSuccess(v bool) *SearchCityPageResponseBody {
	s.Success = &v
	return s
}

func (s *SearchCityPageResponseBody) SetTracerId(v string) *SearchCityPageResponseBody {
	s.TracerId = &v
	return s
}

func (s *SearchCityPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCityPageResponseBodyData struct {
	// example:
	//
	// []
	Cities []*SearchCityPageResponseBodyDataCities `json:"Cities,omitempty" xml:"Cities,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchCityPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchCityPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchCityPageResponseBodyData) GetCities() []*SearchCityPageResponseBodyDataCities {
	return s.Cities
}

func (s *SearchCityPageResponseBodyData) GetHasNext() *bool {
	return s.HasNext
}

func (s *SearchCityPageResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *SearchCityPageResponseBodyData) SetCities(v []*SearchCityPageResponseBodyDataCities) *SearchCityPageResponseBodyData {
	s.Cities = v
	return s
}

func (s *SearchCityPageResponseBodyData) SetHasNext(v bool) *SearchCityPageResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *SearchCityPageResponseBodyData) SetTotal(v int32) *SearchCityPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *SearchCityPageResponseBodyData) Validate() error {
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

type SearchCityPageResponseBodyDataCities struct {
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

func (s SearchCityPageResponseBodyDataCities) String() string {
	return dara.Prettify(s)
}

func (s SearchCityPageResponseBodyDataCities) GoString() string {
	return s.String()
}

func (s *SearchCityPageResponseBodyDataCities) GetCnName() *string {
	return s.CnName
}

func (s *SearchCityPageResponseBodyDataCities) GetCode() *int32 {
	return s.Code
}

func (s *SearchCityPageResponseBodyDataCities) GetCountry() *int32 {
	return s.Country
}

func (s *SearchCityPageResponseBodyDataCities) GetCountryCode() *string {
	return s.CountryCode
}

func (s *SearchCityPageResponseBodyDataCities) GetEnName() *string {
	return s.EnName
}

func (s *SearchCityPageResponseBodyDataCities) GetLevel() *int32 {
	return s.Level
}

func (s *SearchCityPageResponseBodyDataCities) GetParentCode() *int32 {
	return s.ParentCode
}

func (s *SearchCityPageResponseBodyDataCities) GetRegion() *int32 {
	return s.Region
}

func (s *SearchCityPageResponseBodyDataCities) GetType() *int32 {
	return s.Type
}

func (s *SearchCityPageResponseBodyDataCities) SetCnName(v string) *SearchCityPageResponseBodyDataCities {
	s.CnName = &v
	return s
}

func (s *SearchCityPageResponseBodyDataCities) SetCode(v int32) *SearchCityPageResponseBodyDataCities {
	s.Code = &v
	return s
}

func (s *SearchCityPageResponseBodyDataCities) SetCountry(v int32) *SearchCityPageResponseBodyDataCities {
	s.Country = &v
	return s
}

func (s *SearchCityPageResponseBodyDataCities) SetCountryCode(v string) *SearchCityPageResponseBodyDataCities {
	s.CountryCode = &v
	return s
}

func (s *SearchCityPageResponseBodyDataCities) SetEnName(v string) *SearchCityPageResponseBodyDataCities {
	s.EnName = &v
	return s
}

func (s *SearchCityPageResponseBodyDataCities) SetLevel(v int32) *SearchCityPageResponseBodyDataCities {
	s.Level = &v
	return s
}

func (s *SearchCityPageResponseBodyDataCities) SetParentCode(v int32) *SearchCityPageResponseBodyDataCities {
	s.ParentCode = &v
	return s
}

func (s *SearchCityPageResponseBodyDataCities) SetRegion(v int32) *SearchCityPageResponseBodyDataCities {
	s.Region = &v
	return s
}

func (s *SearchCityPageResponseBodyDataCities) SetType(v int32) *SearchCityPageResponseBodyDataCities {
	s.Type = &v
	return s
}

func (s *SearchCityPageResponseBodyDataCities) Validate() error {
	return dara.Validate(s)
}

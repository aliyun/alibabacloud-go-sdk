// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutboundNumbersOfUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListOutboundNumbersOfUserResponseBody
	GetCode() *string
	SetData(v *ListOutboundNumbersOfUserResponseBodyData) *ListOutboundNumbersOfUserResponseBody
	GetData() *ListOutboundNumbersOfUserResponseBodyData
	SetHttpStatusCode(v int32) *ListOutboundNumbersOfUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListOutboundNumbersOfUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListOutboundNumbersOfUserResponseBody
	GetRequestId() *string
}

type ListOutboundNumbersOfUserResponseBody struct {
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListOutboundNumbersOfUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOutboundNumbersOfUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundNumbersOfUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListOutboundNumbersOfUserResponseBody) GetData() *ListOutboundNumbersOfUserResponseBodyData {
	return s.Data
}

func (s *ListOutboundNumbersOfUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListOutboundNumbersOfUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListOutboundNumbersOfUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOutboundNumbersOfUserResponseBody) SetCode(v string) *ListOutboundNumbersOfUserResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBody) SetData(v *ListOutboundNumbersOfUserResponseBodyData) *ListOutboundNumbersOfUserResponseBody {
	s.Data = v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBody) SetHttpStatusCode(v int32) *ListOutboundNumbersOfUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBody) SetMessage(v string) *ListOutboundNumbersOfUserResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBody) SetRequestId(v string) *ListOutboundNumbersOfUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOutboundNumbersOfUserResponseBodyData struct {
	List []*ListOutboundNumbersOfUserResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOutboundNumbersOfUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundNumbersOfUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserResponseBodyData) GetList() []*ListOutboundNumbersOfUserResponseBodyDataList {
	return s.List
}

func (s *ListOutboundNumbersOfUserResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOutboundNumbersOfUserResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOutboundNumbersOfUserResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOutboundNumbersOfUserResponseBodyData) SetList(v []*ListOutboundNumbersOfUserResponseBodyDataList) *ListOutboundNumbersOfUserResponseBodyData {
	s.List = v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyData) SetPageNumber(v int32) *ListOutboundNumbersOfUserResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyData) SetPageSize(v int32) *ListOutboundNumbersOfUserResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyData) SetTotalCount(v int32) *ListOutboundNumbersOfUserResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListOutboundNumbersOfUserResponseBodyDataList struct {
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 0830019****
	Number   *string `json:"Number,omitempty" xml:"Number,omitempty"`
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// Deprecated
	ProviderCode *string `json:"ProviderCode,omitempty" xml:"ProviderCode,omitempty"`
	// Deprecated
	ProviderDisplayName *string `json:"ProviderDisplayName,omitempty" xml:"ProviderDisplayName,omitempty"`
	ProviderShortName   *string `json:"ProviderShortName,omitempty" xml:"ProviderShortName,omitempty"`
	ProviderType        *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	Province            *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s ListOutboundNumbersOfUserResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListOutboundNumbersOfUserResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) GetCity() *string {
	return s.City
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) GetNumber() *string {
	return s.Number
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) GetProvider() *string {
	return s.Provider
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) GetProviderCode() *string {
	return s.ProviderCode
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) GetProviderDisplayName() *string {
	return s.ProviderDisplayName
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) GetProviderShortName() *string {
	return s.ProviderShortName
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) GetProvince() *string {
	return s.Province
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetCity(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.City = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetNumber(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetProvider(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.Provider = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetProviderCode(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.ProviderCode = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetProviderDisplayName(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.ProviderDisplayName = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetProviderShortName(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.ProviderShortName = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetProviderType(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.ProviderType = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) SetProvince(v string) *ListOutboundNumbersOfUserResponseBodyDataList {
	s.Province = &v
	return s
}

func (s *ListOutboundNumbersOfUserResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

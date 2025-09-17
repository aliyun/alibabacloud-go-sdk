// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessTokensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokens(v *ListAccessTokensResponseBodyAccessTokens) *ListAccessTokensResponseBody
	GetAccessTokens() *ListAccessTokensResponseBodyAccessTokens
	SetPageNumber(v int32) *ListAccessTokensResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAccessTokensResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListAccessTokensResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAccessTokensResponseBody
	GetTotalCount() *int32
}

type ListAccessTokensResponseBody struct {
	// The activation codes returned.
	AccessTokens *ListAccessTokensResponseBodyAccessTokens `json:"AccessTokens,omitempty" xml:"AccessTokens,omitempty" type:"Struct"`
	// The number of entries per page. Valid values:
	//
	// 	- 10
	//
	// 	- 20
	//
	// 	- 50
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E2DA3097-79B9-53AE-B0DF-281DC54F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of migration sources returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccessTokensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessTokensResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponseBody) GetAccessTokens() *ListAccessTokensResponseBodyAccessTokens {
	return s.AccessTokens
}

func (s *ListAccessTokensResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAccessTokensResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAccessTokensResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessTokensResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAccessTokensResponseBody) SetAccessTokens(v *ListAccessTokensResponseBodyAccessTokens) *ListAccessTokensResponseBody {
	s.AccessTokens = v
	return s
}

func (s *ListAccessTokensResponseBody) SetPageNumber(v int32) *ListAccessTokensResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAccessTokensResponseBody) SetPageSize(v int32) *ListAccessTokensResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAccessTokensResponseBody) SetRequestId(v string) *ListAccessTokensResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessTokensResponseBody) SetTotalCount(v int32) *ListAccessTokensResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAccessTokensResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAccessTokensResponseBodyAccessTokens struct {
	AccessToken []*ListAccessTokensResponseBodyAccessTokensAccessToken `json:"AccessToken,omitempty" xml:"AccessToken,omitempty" type:"Repeated"`
}

func (s ListAccessTokensResponseBodyAccessTokens) String() string {
	return dara.Prettify(s)
}

func (s ListAccessTokensResponseBodyAccessTokens) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponseBodyAccessTokens) GetAccessToken() []*ListAccessTokensResponseBodyAccessTokensAccessToken {
	return s.AccessToken
}

func (s *ListAccessTokensResponseBodyAccessTokens) SetAccessToken(v []*ListAccessTokensResponseBodyAccessTokensAccessToken) *ListAccessTokensResponseBodyAccessTokens {
	s.AccessToken = v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokens) Validate() error {
	return dara.Validate(s)
}

type ListAccessTokensResponseBodyAccessTokensAccessToken struct {
	// The ID of the activation code.
	//
	// example:
	//
	// at-bp1akz2zp67r0k6r****
	AccessTokenId *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	// The maximum number of times that the activation code can be used. Valid values: 1 to 1000.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The time when the activation code was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-09T02:35:44Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the activation code.
	//
	// example:
	//
	// This is an activation code
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the activation code.
	//
	// example:
	//
	// test_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of migration sources whose information has been imported to Server Migration Center (SMC) by using the activation code.
	//
	// example:
	//
	// 5
	RegisteredCount *string `json:"RegisteredCount,omitempty" xml:"RegisteredCount,omitempty"`
	// The status of the activation code. Valid values:
	//
	// 	- activated
	//
	// 	- unactivated
	//
	// 	- expired
	//
	// example:
	//
	// activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The validity period of the activation code. Unit: day. Valid values: 1 to 90. Default value: 30.
	//
	// example:
	//
	// 30
	TimeToLiveInDays *string `json:"TimeToLiveInDays,omitempty" xml:"TimeToLiveInDays,omitempty"`
}

func (s ListAccessTokensResponseBodyAccessTokensAccessToken) String() string {
	return dara.Prettify(s)
}

func (s ListAccessTokensResponseBodyAccessTokensAccessToken) GoString() string {
	return s.String()
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) GetAccessTokenId() *string {
	return s.AccessTokenId
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) GetCount() *string {
	return s.Count
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) GetDescription() *string {
	return s.Description
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) GetName() *string {
	return s.Name
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) GetRegisteredCount() *string {
	return s.RegisteredCount
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) GetStatus() *string {
	return s.Status
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) GetTimeToLiveInDays() *string {
	return s.TimeToLiveInDays
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetAccessTokenId(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.AccessTokenId = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetCount(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.Count = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetCreationTime(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.CreationTime = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetDescription(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.Description = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetName(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.Name = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetRegisteredCount(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.RegisteredCount = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetStatus(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.Status = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) SetTimeToLiveInDays(v string) *ListAccessTokensResponseBodyAccessTokensAccessToken {
	s.TimeToLiveInDays = &v
	return s
}

func (s *ListAccessTokensResponseBodyAccessTokensAccessToken) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePurchasedApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribePurchasedApisRequest
	GetApiId() *string
	SetApiName(v string) *DescribePurchasedApisRequest
	GetApiName() *string
	SetGroupId(v string) *DescribePurchasedApisRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribePurchasedApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePurchasedApisRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribePurchasedApisRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribePurchasedApisRequest
	GetStageName() *string
	SetVisibility(v string) *DescribePurchasedApisRequest
	GetVisibility() *string
}

type DescribePurchasedApisRequest struct {
	// The ID of the API.
	//
	// example:
	//
	// 3b81fd160f5645e097cc8855d75a1cf6
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// Cz88IpQuery
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// d27ad517be5f4c95ac3631780a8f4d50
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **PRE**
	//
	// 	- **TEST**
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// Specifies whether the API is public. Valid values:
	//
	// 	- **PUBLIC**: indicates that the API is public. If you set this parameter to PUBLIC, this API is displayed on the API List page in the console for all users after the API is published to the production environment.
	//
	// 	- **PRIVATE**: indicates that the API is private. If you set this parameter to PRIVATE, this API is not displayed in Alibaba Cloud Marketplace after the API group to which this API belongs is made available.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribePurchasedApisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePurchasedApisRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribePurchasedApisRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribePurchasedApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePurchasedApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePurchasedApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePurchasedApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePurchasedApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribePurchasedApisRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *DescribePurchasedApisRequest) SetApiId(v string) *DescribePurchasedApisRequest {
	s.ApiId = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetApiName(v string) *DescribePurchasedApisRequest {
	s.ApiName = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetGroupId(v string) *DescribePurchasedApisRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetPageNumber(v int32) *DescribePurchasedApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetPageSize(v int32) *DescribePurchasedApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetSecurityToken(v string) *DescribePurchasedApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetStageName(v string) *DescribePurchasedApisRequest {
	s.StageName = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetVisibility(v string) *DescribePurchasedApisRequest {
	s.Visibility = &v
	return s
}

func (s *DescribePurchasedApisRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApisRequest
	GetApiId() *string
	SetApiMethod(v string) *DescribeApisRequest
	GetApiMethod() *string
	SetApiName(v string) *DescribeApisRequest
	GetApiName() *string
	SetApiPath(v string) *DescribeApisRequest
	GetApiPath() *string
	SetCatalogId(v string) *DescribeApisRequest
	GetCatalogId() *string
	SetEnableTagAuth(v bool) *DescribeApisRequest
	GetEnableTagAuth() *bool
	SetGroupId(v string) *DescribeApisRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribeApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApisRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApisRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApisRequest
	GetStageName() *string
	SetTag(v []*DescribeApisRequestTag) *DescribeApisRequest
	GetTag() []*DescribeApisRequestTag
	SetUnDeployed(v bool) *DescribeApisRequest
	GetUnDeployed() *bool
	SetVisibility(v string) *DescribeApisRequest
	GetVisibility() *string
}

type DescribeApisRequest struct {
	// The API ID.
	//
	// example:
	//
	// f68c19ee3bd1478fb58aa05ce8ae9b5a
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The HTTP method of the API request.
	//
	// example:
	//
	// GET
	ApiMethod *string `json:"ApiMethod,omitempty" xml:"ApiMethod,omitempty"`
	// The API name. The name is used for fuzzy match.
	//
	// example:
	//
	// weather
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /illegal/query
	ApiPath *string `json:"ApiPath,omitempty" xml:"ApiPath,omitempty"`
	// The category ID.
	//
	// example:
	//
	// 1553414085247362
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	// Specifies whether to enable tag verification.
	//
	// example:
	//
	// true
	EnableTagAuth *bool `json:"EnableTagAuth,omitempty" xml:"EnableTagAuth,omitempty"`
	// The API group ID.
	//
	// example:
	//
	// c4a4d2de657548a2bd485d5d4df42b4a
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment in which you want to perform this operation. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the staging environment
	//
	// 	- **TEST**: the test environment
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The tags of objects that match the rule.
	//
	// example:
	//
	// Key， Value
	Tag []*DescribeApisRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to filter unpublished APIs.
	//
	// example:
	//
	// true
	UnDeployed *bool `json:"UnDeployed,omitempty" xml:"UnDeployed,omitempty"`
	// Specifies whether the API is public. Valid values:
	//
	// 	- **PUBLIC**: The API is public. If you publish the definition of a public API to the production environment, the definition is displayed on the APIs page for all users.
	//
	// 	- **PRIVATE**: The API is private. If you publish an API group that contains a private API in Alibaba Cloud Marketplace, the API is not displayed in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisRequest) GetApiMethod() *string {
	return s.ApiMethod
}

func (s *DescribeApisRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApisRequest) GetApiPath() *string {
	return s.ApiPath
}

func (s *DescribeApisRequest) GetCatalogId() *string {
	return s.CatalogId
}

func (s *DescribeApisRequest) GetEnableTagAuth() *bool {
	return s.EnableTagAuth
}

func (s *DescribeApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApisRequest) GetTag() []*DescribeApisRequestTag {
	return s.Tag
}

func (s *DescribeApisRequest) GetUnDeployed() *bool {
	return s.UnDeployed
}

func (s *DescribeApisRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *DescribeApisRequest) SetApiId(v string) *DescribeApisRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApisRequest) SetApiMethod(v string) *DescribeApisRequest {
	s.ApiMethod = &v
	return s
}

func (s *DescribeApisRequest) SetApiName(v string) *DescribeApisRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApisRequest) SetApiPath(v string) *DescribeApisRequest {
	s.ApiPath = &v
	return s
}

func (s *DescribeApisRequest) SetCatalogId(v string) *DescribeApisRequest {
	s.CatalogId = &v
	return s
}

func (s *DescribeApisRequest) SetEnableTagAuth(v bool) *DescribeApisRequest {
	s.EnableTagAuth = &v
	return s
}

func (s *DescribeApisRequest) SetGroupId(v string) *DescribeApisRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApisRequest) SetPageNumber(v int32) *DescribeApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisRequest) SetPageSize(v int32) *DescribeApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisRequest) SetSecurityToken(v string) *DescribeApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisRequest) SetStageName(v string) *DescribeApisRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApisRequest) SetTag(v []*DescribeApisRequestTag) *DescribeApisRequest {
	s.Tag = v
	return s
}

func (s *DescribeApisRequest) SetUnDeployed(v bool) *DescribeApisRequest {
	s.UnDeployed = &v
	return s
}

func (s *DescribeApisRequest) SetVisibility(v string) *DescribeApisRequest {
	s.Visibility = &v
	return s
}

func (s *DescribeApisRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApisRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeApisRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeApisRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeApisRequestTag) SetKey(v string) *DescribeApisRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeApisRequestTag) SetValue(v string) *DescribeApisRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeApisRequestTag) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeployedApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeDeployedApisRequest
	GetApiId() *string
	SetApiMethod(v string) *DescribeDeployedApisRequest
	GetApiMethod() *string
	SetApiName(v string) *DescribeDeployedApisRequest
	GetApiName() *string
	SetApiPath(v string) *DescribeDeployedApisRequest
	GetApiPath() *string
	SetEnableTagAuth(v bool) *DescribeDeployedApisRequest
	GetEnableTagAuth() *bool
	SetGroupId(v string) *DescribeDeployedApisRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribeDeployedApisRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDeployedApisRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeDeployedApisRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeDeployedApisRequest
	GetStageName() *string
	SetTag(v []*DescribeDeployedApisRequestTag) *DescribeDeployedApisRequest
	GetTag() []*DescribeDeployedApisRequestTag
}

type DescribeDeployedApisRequest struct {
	// The ID of the API.
	//
	// example:
	//
	// c076144d7878437b8f82fb85890ce6a0
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The HTTP method of the API request.
	//
	// example:
	//
	// POST
	ApiMethod *string `json:"ApiMethod,omitempty" xml:"ApiMethod,omitempty"`
	// The name of the API. The name is used for fuzzy match.
	//
	// example:
	//
	// weather
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The request path of the API.
	//
	// example:
	//
	// /st4
	ApiPath *string `json:"ApiPath,omitempty" xml:"ApiPath,omitempty"`
	// Specifies whether to enable tag verification.
	//
	// example:
	//
	// true
	EnableTagAuth *bool `json:"EnableTagAuth,omitempty" xml:"EnableTagAuth,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 63be9002440b4778a61122f14c2b2bbb
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 10.
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
	// 	- **TEST**
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The tags.
	Tag []*DescribeDeployedApisRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeDeployedApisRequest) GetApiMethod() *string {
	return s.ApiMethod
}

func (s *DescribeDeployedApisRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeDeployedApisRequest) GetApiPath() *string {
	return s.ApiPath
}

func (s *DescribeDeployedApisRequest) GetEnableTagAuth() *bool {
	return s.EnableTagAuth
}

func (s *DescribeDeployedApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDeployedApisRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDeployedApisRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeployedApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDeployedApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeDeployedApisRequest) GetTag() []*DescribeDeployedApisRequestTag {
	return s.Tag
}

func (s *DescribeDeployedApisRequest) SetApiId(v string) *DescribeDeployedApisRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetApiMethod(v string) *DescribeDeployedApisRequest {
	s.ApiMethod = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetApiName(v string) *DescribeDeployedApisRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetApiPath(v string) *DescribeDeployedApisRequest {
	s.ApiPath = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetEnableTagAuth(v bool) *DescribeDeployedApisRequest {
	s.EnableTagAuth = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetGroupId(v string) *DescribeDeployedApisRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetPageNumber(v int32) *DescribeDeployedApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetPageSize(v int32) *DescribeDeployedApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetSecurityToken(v string) *DescribeDeployedApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetStageName(v string) *DescribeDeployedApisRequest {
	s.StageName = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetTag(v []*DescribeDeployedApisRequestTag) *DescribeDeployedApisRequest {
	s.Tag = v
	return s
}

func (s *DescribeDeployedApisRequest) Validate() error {
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

type DescribeDeployedApisRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// appname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testapp
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDeployedApisRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApisRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDeployedApisRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDeployedApisRequestTag) SetKey(v string) *DescribeDeployedApisRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDeployedApisRequestTag) SetValue(v string) *DescribeDeployedApisRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDeployedApisRequestTag) Validate() error {
	return dara.Validate(s)
}

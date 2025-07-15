// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApiHistoriesRequest
	GetApiId() *string
	SetApiName(v string) *DescribeApiHistoriesRequest
	GetApiName() *string
	SetGroupId(v string) *DescribeApiHistoriesRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribeApiHistoriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApiHistoriesRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribeApiHistoriesRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeApiHistoriesRequest
	GetStageName() *string
}

type DescribeApiHistoriesRequest struct {
	// The ID of the API.
	//
	// example:
	//
	// c076144d7878437b8f82fb85890ce6a0
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API. The name is used for fuzzy match.
	//
	// example:
	//
	// weather
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1084f9034c744137901057206b39d2b6
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
}

func (s DescribeApiHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiHistoriesRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiHistoriesRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiHistoriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApiHistoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApiHistoriesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeApiHistoriesRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiHistoriesRequest) SetApiId(v string) *DescribeApiHistoriesRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetApiName(v string) *DescribeApiHistoriesRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetGroupId(v string) *DescribeApiHistoriesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetPageNumber(v int32) *DescribeApiHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetPageSize(v int32) *DescribeApiHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetSecurityToken(v string) *DescribeApiHistoriesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetStageName(v string) *DescribeApiHistoriesRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiHistoriesRequest) Validate() error {
	return dara.Validate(s)
}

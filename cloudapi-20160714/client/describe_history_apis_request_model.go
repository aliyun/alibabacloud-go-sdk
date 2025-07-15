// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeHistoryApisRequest
	GetApiId() *string
	SetApiName(v string) *DescribeHistoryApisRequest
	GetApiName() *string
	SetGroupId(v string) *DescribeHistoryApisRequest
	GetGroupId() *string
	SetPageNumber(v string) *DescribeHistoryApisRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeHistoryApisRequest
	GetPageSize() *string
	SetSecurityToken(v string) *DescribeHistoryApisRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribeHistoryApisRequest
	GetStageName() *string
}

type DescribeHistoryApisRequest struct {
	// The ID of the API.
	//
	// example:
	//
	// a12068f555964ca8a0c9c33288f1e5a7
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API. The name is used for fuzzy match.
	//
	// example:
	//
	// getPersonInfo
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0ccb66aadb5345b78a40f57d192d8aa4
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The page number of the current page.
	//
	// example:
	//
	// 2
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s DescribeHistoryApisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeHistoryApisRequest) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeHistoryApisRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeHistoryApisRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeHistoryApisRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeHistoryApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeHistoryApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribeHistoryApisRequest) SetApiId(v string) *DescribeHistoryApisRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetApiName(v string) *DescribeHistoryApisRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetGroupId(v string) *DescribeHistoryApisRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetPageNumber(v string) *DescribeHistoryApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetPageSize(v string) *DescribeHistoryApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetSecurityToken(v string) *DescribeHistoryApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetStageName(v string) *DescribeHistoryApisRequest {
	s.StageName = &v
	return s
}

func (s *DescribeHistoryApisRequest) Validate() error {
	return dara.Validate(s)
}

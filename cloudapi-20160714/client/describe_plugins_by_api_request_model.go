// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginsByApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribePluginsByApiRequest
	GetApiId() *string
	SetGroupId(v string) *DescribePluginsByApiRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribePluginsByApiRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginsByApiRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribePluginsByApiRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribePluginsByApiRequest
	GetStageName() *string
}

type DescribePluginsByApiRequest struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93f4ead4080c4b2da70b7f81f50ae459
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the group to which the API belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3c7a38392e764718ad7673e7b7f535d4
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The environment in which the API is running. Valid values:
	//
	// 	- **RELEASE**: production
	//
	// 	- **PRE**: staging
	//
	// 	- **TEST**: test
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribePluginsByApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByApiRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribePluginsByApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePluginsByApiRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginsByApiRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginsByApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePluginsByApiRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribePluginsByApiRequest) SetApiId(v string) *DescribePluginsByApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribePluginsByApiRequest) SetGroupId(v string) *DescribePluginsByApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePluginsByApiRequest) SetPageNumber(v int32) *DescribePluginsByApiRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginsByApiRequest) SetPageSize(v int32) *DescribePluginsByApiRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePluginsByApiRequest) SetSecurityToken(v string) *DescribePluginsByApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePluginsByApiRequest) SetStageName(v string) *DescribePluginsByApiRequest {
	s.StageName = &v
	return s
}

func (s *DescribePluginsByApiRequest) Validate() error {
	return dara.Validate(s)
}

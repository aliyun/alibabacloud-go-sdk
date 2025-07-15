// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginsByGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribePluginsByGroupRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribePluginsByGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginsByGroupRequest
	GetPageSize() *int32
	SetSecurityToken(v string) *DescribePluginsByGroupRequest
	GetSecurityToken() *string
	SetStageName(v string) *DescribePluginsByGroupRequest
	GetStageName() *string
}

type DescribePluginsByGroupRequest struct {
	// Group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 08ae4aa0f95e4321849ee57f4e0b3077
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Pagination parameter: current page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Pagination parameter: number of items per page.
	//
	// example:
	//
	// 10
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies the environment of the API to operate on.
	//
	// - **RELEASE**: Production
	//
	// - **PRE**: Pre-release
	//
	// - **TEST**: Test
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribePluginsByGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsByGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginsByGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePluginsByGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginsByGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginsByGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePluginsByGroupRequest) GetStageName() *string {
	return s.StageName
}

func (s *DescribePluginsByGroupRequest) SetGroupId(v string) *DescribePluginsByGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePluginsByGroupRequest) SetPageNumber(v int32) *DescribePluginsByGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginsByGroupRequest) SetPageSize(v int32) *DescribePluginsByGroupRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePluginsByGroupRequest) SetSecurityToken(v string) *DescribePluginsByGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePluginsByGroupRequest) SetStageName(v string) *DescribePluginsByGroupRequest {
	s.StageName = &v
	return s
}

func (s *DescribePluginsByGroupRequest) Validate() error {
	return dara.Validate(s)
}

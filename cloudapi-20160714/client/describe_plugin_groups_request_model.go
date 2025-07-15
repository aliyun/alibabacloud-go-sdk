// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribePluginGroupsRequest
	GetDescription() *string
	SetGroupId(v string) *DescribePluginGroupsRequest
	GetGroupId() *string
	SetGroupName(v string) *DescribePluginGroupsRequest
	GetGroupName() *string
	SetPageNumber(v int32) *DescribePluginGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginGroupsRequest
	GetPageSize() *int32
	SetPluginId(v string) *DescribePluginGroupsRequest
	GetPluginId() *string
	SetSecurityToken(v string) *DescribePluginGroupsRequest
	GetSecurityToken() *string
}

type DescribePluginGroupsRequest struct {
	// API group description
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// API group ID
	//
	// example:
	//
	// 8cc2a3cbe3394524b6e71be5db9b02a3
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// API group name
	//
	// example:
	//
	// crm_custom_service
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Pagination parameter: current page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Pagination parameter: number of items per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// API Gateway plugin ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1f3bde29b43d4d53989248327ff737f2
	PluginId      *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePluginGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginGroupsRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribePluginGroupsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribePluginGroupsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribePluginGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginGroupsRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *DescribePluginGroupsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePluginGroupsRequest) SetDescription(v string) *DescribePluginGroupsRequest {
	s.Description = &v
	return s
}

func (s *DescribePluginGroupsRequest) SetGroupId(v string) *DescribePluginGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePluginGroupsRequest) SetGroupName(v string) *DescribePluginGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribePluginGroupsRequest) SetPageNumber(v int32) *DescribePluginGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginGroupsRequest) SetPageSize(v int32) *DescribePluginGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePluginGroupsRequest) SetPluginId(v string) *DescribePluginGroupsRequest {
	s.PluginId = &v
	return s
}

func (s *DescribePluginGroupsRequest) SetSecurityToken(v string) *DescribePluginGroupsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePluginGroupsRequest) Validate() error {
	return dara.Validate(s)
}

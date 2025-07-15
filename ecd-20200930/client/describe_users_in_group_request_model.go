// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersInGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectState(v int32) *DescribeUsersInGroupRequest
	GetConnectState() *int32
	SetDesktopGroupId(v string) *DescribeUsersInGroupRequest
	GetDesktopGroupId() *string
	SetEndUserId(v string) *DescribeUsersInGroupRequest
	GetEndUserId() *string
	SetEndUserIds(v []*string) *DescribeUsersInGroupRequest
	GetEndUserIds() []*string
	SetFilter(v string) *DescribeUsersInGroupRequest
	GetFilter() *string
	SetMaxResults(v int32) *DescribeUsersInGroupRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeUsersInGroupRequest
	GetNextToken() *string
	SetOrgId(v string) *DescribeUsersInGroupRequest
	GetOrgId() *string
	SetQueryUserDetail(v bool) *DescribeUsersInGroupRequest
	GetQueryUserDetail() *bool
	SetRegionId(v string) *DescribeUsersInGroupRequest
	GetRegionId() *string
}

type DescribeUsersInGroupRequest struct {
	// The status of the desktop connection for the end user.
	//
	// Valid values:
	//
	// - 0: Disconnected.
	//
	// - 1: Connected.
	//
	// example:
	//
	// 1
	ConnectState *int32 `json:"ConnectState,omitempty" xml:"ConnectState,omitempty"`
	// The ID of the cloud computer share.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-8ttn55ujj8nj8****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The ID of the authorized user.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The IDs of the authorized users.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The query string for fuzzy match. If you specify this parameter, the system returns all results that contain the string.
	//
	// example:
	//
	// alice
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The number of entries to return on each page.
	//
	// 	- Maximum value: 100.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. If this parameter is left empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the organization to which the end user belongs.
	//
	// example:
	//
	// org-d0fua2oyukw8j****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// Specifies whether to query user details.
	//
	// Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// false
	QueryUserDetail *bool `json:"QueryUserDetail,omitempty" xml:"QueryUserDetail,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUsersInGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersInGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupRequest) GetConnectState() *int32 {
	return s.ConnectState
}

func (s *DescribeUsersInGroupRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeUsersInGroupRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeUsersInGroupRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeUsersInGroupRequest) GetFilter() *string {
	return s.Filter
}

func (s *DescribeUsersInGroupRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeUsersInGroupRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUsersInGroupRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeUsersInGroupRequest) GetQueryUserDetail() *bool {
	return s.QueryUserDetail
}

func (s *DescribeUsersInGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUsersInGroupRequest) SetConnectState(v int32) *DescribeUsersInGroupRequest {
	s.ConnectState = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetDesktopGroupId(v string) *DescribeUsersInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetEndUserId(v string) *DescribeUsersInGroupRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetEndUserIds(v []*string) *DescribeUsersInGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeUsersInGroupRequest) SetFilter(v string) *DescribeUsersInGroupRequest {
	s.Filter = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetMaxResults(v int32) *DescribeUsersInGroupRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetNextToken(v string) *DescribeUsersInGroupRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetOrgId(v string) *DescribeUsersInGroupRequest {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetQueryUserDetail(v bool) *DescribeUsersInGroupRequest {
	s.QueryUserDetail = &v
	return s
}

func (s *DescribeUsersInGroupRequest) SetRegionId(v string) *DescribeUsersInGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUsersInGroupRequest) Validate() error {
	return dara.Validate(s)
}

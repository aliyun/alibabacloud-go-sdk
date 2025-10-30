// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServiceUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVpcEndpointServiceUsersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVpcEndpointServiceUsersRequest
	GetNextToken() *string
	SetRegionId(v string) *ListVpcEndpointServiceUsersRequest
	GetRegionId() *string
	SetServiceId(v string) *ListVpcEndpointServiceUsersRequest
	GetServiceId() *string
	SetUserId(v int64) *ListVpcEndpointServiceUsersRequest
	GetUserId() *int64
	SetUserListType(v string) *ListVpcEndpointServiceUsersRequest
	GetUserListType() *string
}

type ListVpcEndpointServiceUsersRequest struct {
	// The number of entries to return on each page. Valid values: **1 to 50**. Default value: **50**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If this is your first request and no next requests are to be performed, you do not need to specify this parameter.
	//
	// 	- If a next request is to be performed, set the value to the value of **NextToken*	- that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the endpoint service that you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the Alibaba Cloud account in the whitelist of the endpoint service.
	//
	// example:
	//
	// 12345678
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The type of the user list in the whitelist of the endpoint service.
	//
	// example:
	//
	// Users
	UserListType *string `json:"UserListType,omitempty" xml:"UserListType,omitempty"`
}

func (s ListVpcEndpointServiceUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServiceUsersRequest) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVpcEndpointServiceUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVpcEndpointServiceUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListVpcEndpointServiceUsersRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListVpcEndpointServiceUsersRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *ListVpcEndpointServiceUsersRequest) GetUserListType() *string {
	return s.UserListType
}

func (s *ListVpcEndpointServiceUsersRequest) SetMaxResults(v int32) *ListVpcEndpointServiceUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) SetNextToken(v string) *ListVpcEndpointServiceUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) SetRegionId(v string) *ListVpcEndpointServiceUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) SetServiceId(v string) *ListVpcEndpointServiceUsersRequest {
	s.ServiceId = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) SetUserId(v int64) *ListVpcEndpointServiceUsersRequest {
	s.UserId = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) SetUserListType(v string) *ListVpcEndpointServiceUsersRequest {
	s.UserListType = &v
	return s
}

func (s *ListVpcEndpointServiceUsersRequest) Validate() error {
	return dara.Validate(s)
}

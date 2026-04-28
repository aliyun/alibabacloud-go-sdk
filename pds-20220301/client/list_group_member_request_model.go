// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListGroupMemberRequest
	GetGroupId() *string
	SetLimit(v int32) *ListGroupMemberRequest
	GetLimit() *int32
	SetMarker(v string) *ListGroupMemberRequest
	GetMarker() *string
	SetMemberType(v string) *ListGroupMemberRequest
	GetMemberType() *string
}

type ListGroupMemberRequest struct {
	// The ID of the group of which you want to query members.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e5***2c2
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The member type. If you do not specify this parameter, both types of members are returned. Valid values:
	//
	// 	- user
	//
	// 	- group
	//
	// Note: A group can be a member of only one group. It cannot be a member of multiple groups. A user can be a member of multiple groups.
	//
	// example:
	//
	// user
	MemberType *string `json:"member_type,omitempty" xml:"member_type,omitempty"`
}

func (s ListGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *ListGroupMemberRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupMemberRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListGroupMemberRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListGroupMemberRequest) GetMemberType() *string {
	return s.MemberType
}

func (s *ListGroupMemberRequest) SetGroupId(v string) *ListGroupMemberRequest {
	s.GroupId = &v
	return s
}

func (s *ListGroupMemberRequest) SetLimit(v int32) *ListGroupMemberRequest {
	s.Limit = &v
	return s
}

func (s *ListGroupMemberRequest) SetMarker(v string) *ListGroupMemberRequest {
	s.Marker = &v
	return s
}

func (s *ListGroupMemberRequest) SetMemberType(v string) *ListGroupMemberRequest {
	s.MemberType = &v
	return s
}

func (s *ListGroupMemberRequest) Validate() error {
	return dara.Validate(s)
}

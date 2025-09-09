// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsForUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroups(v []*ListHostGroupsForUserGroupResponseBodyHostGroups) *ListHostGroupsForUserGroupResponseBody
	GetHostGroups() []*ListHostGroupsForUserGroupResponseBodyHostGroups
	SetRequestId(v string) *ListHostGroupsForUserGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHostGroupsForUserGroupResponseBody
	GetTotalCount() *int32
}

type ListHostGroupsForUserGroupResponseBody struct {
	// The host groups returned.
	HostGroups []*ListHostGroupsForUserGroupResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of host groups returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostGroupsForUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupResponseBody) GetHostGroups() []*ListHostGroupsForUserGroupResponseBodyHostGroups {
	return s.HostGroups
}

func (s *ListHostGroupsForUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostGroupsForUserGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHostGroupsForUserGroupResponseBody) SetHostGroups(v []*ListHostGroupsForUserGroupResponseBodyHostGroups) *ListHostGroupsForUserGroupResponseBody {
	s.HostGroups = v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBody) SetRequestId(v string) *ListHostGroupsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBody) SetTotalCount(v int32) *ListHostGroupsForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHostGroupsForUserGroupResponseBodyHostGroups struct {
	// The description of the host group.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the host group.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The name of the host group.
	//
	// example:
	//
	// group
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListHostGroupsForUserGroupResponseBodyHostGroups) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsForUserGroupResponseBodyHostGroups) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) GetComment() *string {
	return s.Comment
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) SetComment(v string) *ListHostGroupsForUserGroupResponseBodyHostGroups {
	s.Comment = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) SetHostGroupId(v string) *ListHostGroupsForUserGroupResponseBodyHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) SetHostGroupName(v string) *ListHostGroupsForUserGroupResponseBodyHostGroups {
	s.HostGroupName = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) Validate() error {
	return dara.Validate(s)
}

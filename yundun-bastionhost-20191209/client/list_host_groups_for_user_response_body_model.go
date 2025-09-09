// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroups(v []*ListHostGroupsForUserResponseBodyHostGroups) *ListHostGroupsForUserResponseBody
	GetHostGroups() []*ListHostGroupsForUserResponseBodyHostGroups
	SetRequestId(v string) *ListHostGroupsForUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHostGroupsForUserResponseBody
	GetTotalCount() *int32
}

type ListHostGroupsForUserResponseBody struct {
	// The host groups returned.
	HostGroups []*ListHostGroupsForUserResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The request ID.
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

func (s ListHostGroupsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserResponseBody) GetHostGroups() []*ListHostGroupsForUserResponseBodyHostGroups {
	return s.HostGroups
}

func (s *ListHostGroupsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostGroupsForUserResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHostGroupsForUserResponseBody) SetHostGroups(v []*ListHostGroupsForUserResponseBodyHostGroups) *ListHostGroupsForUserResponseBody {
	s.HostGroups = v
	return s
}

func (s *ListHostGroupsForUserResponseBody) SetRequestId(v string) *ListHostGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsForUserResponseBody) SetTotalCount(v int32) *ListHostGroupsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostGroupsForUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListHostGroupsForUserResponseBodyHostGroups struct {
	// The remarks of the host group.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The host group ID.
	//
	// example:
	//
	// １
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The name of the host group.
	//
	// example:
	//
	// group
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListHostGroupsForUserResponseBodyHostGroups) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsForUserResponseBodyHostGroups) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) GetComment() *string {
	return s.Comment
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) SetComment(v string) *ListHostGroupsForUserResponseBodyHostGroups {
	s.Comment = &v
	return s
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) SetHostGroupId(v string) *ListHostGroupsForUserResponseBodyHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) SetHostGroupName(v string) *ListHostGroupsForUserResponseBodyHostGroups {
	s.HostGroupName = &v
	return s
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) Validate() error {
	return dara.Validate(s)
}

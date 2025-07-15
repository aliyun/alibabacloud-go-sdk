// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupList(v []*ListLiveMessageGroupsResponseBodyGroupList) *ListLiveMessageGroupsResponseBody
	GetGroupList() []*ListLiveMessageGroupsResponseBodyGroupList
	SetHasmore(v bool) *ListLiveMessageGroupsResponseBody
	GetHasmore() *bool
	SetNextpageToken(v string) *ListLiveMessageGroupsResponseBody
	GetNextpageToken() *string
	SetRequestId(v string) *ListLiveMessageGroupsResponseBody
	GetRequestId() *string
}

type ListLiveMessageGroupsResponseBody struct {
	// Details about the groups.
	GroupList []*ListLiveMessageGroupsResponseBodyGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// Indicates whether the current page is followed by another page.
	//
	// example:
	//
	// false
	Hasmore *bool `json:"Hasmore,omitempty" xml:"Hasmore,omitempty"`
	// The starting page number for the next query. This parameter is returned only if the value of Hasmore is true.
	//
	// example:
	//
	// 1001
	NextpageToken *string `json:"NextpageToken,omitempty" xml:"NextpageToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B5D95365-5A46-1A6A-BBF5-C7B6BDED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLiveMessageGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupsResponseBody) GetGroupList() []*ListLiveMessageGroupsResponseBodyGroupList {
	return s.GroupList
}

func (s *ListLiveMessageGroupsResponseBody) GetHasmore() *bool {
	return s.Hasmore
}

func (s *ListLiveMessageGroupsResponseBody) GetNextpageToken() *string {
	return s.NextpageToken
}

func (s *ListLiveMessageGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveMessageGroupsResponseBody) SetGroupList(v []*ListLiveMessageGroupsResponseBodyGroupList) *ListLiveMessageGroupsResponseBody {
	s.GroupList = v
	return s
}

func (s *ListLiveMessageGroupsResponseBody) SetHasmore(v bool) *ListLiveMessageGroupsResponseBody {
	s.Hasmore = &v
	return s
}

func (s *ListLiveMessageGroupsResponseBody) SetNextpageToken(v string) *ListLiveMessageGroupsResponseBody {
	s.NextpageToken = &v
	return s
}

func (s *ListLiveMessageGroupsResponseBody) SetRequestId(v string) *ListLiveMessageGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveMessageGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLiveMessageGroupsResponseBodyGroupList struct {
	// The list of the IDs of the group administrators.
	AdminList []*string `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	// The time when the group was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1698299727
	Createtime *int64 `json:"Createtime,omitempty" xml:"Createtime,omitempty"`
	// The ID of the group creator.
	//
	// example:
	//
	// user_77
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// Indicates whether the group is deleted.
	//
	// example:
	//
	// true
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// cU9MeBqf****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The additional information about the group.
	//
	// example:
	//
	// testgroupinfo
	GroupInfo *string `json:"GroupInfo,omitempty" xml:"GroupInfo,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// mytestgroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ListLiveMessageGroupsResponseBodyGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupsResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) GetAdminList() []*string {
	return s.AdminList
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) GetCreatetime() *int64 {
	return s.Createtime
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) GetDelete() *bool {
	return s.Delete
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) GetGroupInfo() *string {
	return s.GroupInfo
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) SetAdminList(v []*string) *ListLiveMessageGroupsResponseBodyGroupList {
	s.AdminList = v
	return s
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) SetCreatetime(v int64) *ListLiveMessageGroupsResponseBodyGroupList {
	s.Createtime = &v
	return s
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) SetCreatorId(v string) *ListLiveMessageGroupsResponseBodyGroupList {
	s.CreatorId = &v
	return s
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) SetDelete(v bool) *ListLiveMessageGroupsResponseBodyGroupList {
	s.Delete = &v
	return s
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) SetGroupId(v string) *ListLiveMessageGroupsResponseBodyGroupList {
	s.GroupId = &v
	return s
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) SetGroupInfo(v string) *ListLiveMessageGroupsResponseBodyGroupList {
	s.GroupInfo = &v
	return s
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) SetGroupName(v string) *ListLiveMessageGroupsResponseBodyGroupList {
	s.GroupName = &v
	return s
}

func (s *ListLiveMessageGroupsResponseBodyGroupList) Validate() error {
	return dara.Validate(s)
}

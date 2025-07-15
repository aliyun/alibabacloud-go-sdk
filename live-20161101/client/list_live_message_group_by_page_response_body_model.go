// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveMessageGroupByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupList(v []*ListLiveMessageGroupByPageResponseBodyGroupList) *ListLiveMessageGroupByPageResponseBody
	GetGroupList() []*ListLiveMessageGroupByPageResponseBodyGroupList
	SetPageNumber(v int32) *ListLiveMessageGroupByPageResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLiveMessageGroupByPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLiveMessageGroupByPageResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLiveMessageGroupByPageResponseBody
	GetTotalCount() *int32
}

type ListLiveMessageGroupByPageResponseBody struct {
	// The list of groups.
	GroupList []*ListLiveMessageGroupByPageResponseBodyGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B5D95365-5A46-1A6A-BBF5-C7B6BDED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLiveMessageGroupByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupByPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupByPageResponseBody) GetGroupList() []*ListLiveMessageGroupByPageResponseBodyGroupList {
	return s.GroupList
}

func (s *ListLiveMessageGroupByPageResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLiveMessageGroupByPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLiveMessageGroupByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveMessageGroupByPageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLiveMessageGroupByPageResponseBody) SetGroupList(v []*ListLiveMessageGroupByPageResponseBodyGroupList) *ListLiveMessageGroupByPageResponseBody {
	s.GroupList = v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBody) SetPageNumber(v int32) *ListLiveMessageGroupByPageResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBody) SetPageSize(v int32) *ListLiveMessageGroupByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBody) SetRequestId(v string) *ListLiveMessageGroupByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBody) SetTotalCount(v int32) *ListLiveMessageGroupByPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLiveMessageGroupByPageResponseBodyGroupList struct {
	// The list of administrators.
	AdminList []*string `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	// The time when the group was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1698299727
	Createtime *int64 `json:"Createtime,omitempty" xml:"Createtime,omitempty"`
	// The ID of the user who created the group.
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
	// The ID of the interactive messaging group.
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

func (s ListLiveMessageGroupByPageResponseBodyGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListLiveMessageGroupByPageResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) GetAdminList() []*string {
	return s.AdminList
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) GetCreatetime() *int64 {
	return s.Createtime
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) GetDelete() *bool {
	return s.Delete
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) GetGroupInfo() *string {
	return s.GroupInfo
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) SetAdminList(v []*string) *ListLiveMessageGroupByPageResponseBodyGroupList {
	s.AdminList = v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) SetCreatetime(v int64) *ListLiveMessageGroupByPageResponseBodyGroupList {
	s.Createtime = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) SetCreatorId(v string) *ListLiveMessageGroupByPageResponseBodyGroupList {
	s.CreatorId = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) SetDelete(v bool) *ListLiveMessageGroupByPageResponseBodyGroupList {
	s.Delete = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) SetGroupId(v string) *ListLiveMessageGroupByPageResponseBodyGroupList {
	s.GroupId = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) SetGroupInfo(v string) *ListLiveMessageGroupByPageResponseBodyGroupList {
	s.GroupInfo = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) SetGroupName(v string) *ListLiveMessageGroupByPageResponseBodyGroupList {
	s.GroupName = &v
	return s
}

func (s *ListLiveMessageGroupByPageResponseBodyGroupList) Validate() error {
	return dara.Validate(s)
}

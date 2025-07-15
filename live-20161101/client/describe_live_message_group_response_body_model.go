// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdminList(v []*string) *DescribeLiveMessageGroupResponseBody
	GetAdminList() []*string
	SetCreatetime(v int64) *DescribeLiveMessageGroupResponseBody
	GetCreatetime() *int64
	SetCreatorId(v string) *DescribeLiveMessageGroupResponseBody
	GetCreatorId() *string
	SetDeletatime(v int64) *DescribeLiveMessageGroupResponseBody
	GetDeletatime() *int64
	SetDelete(v bool) *DescribeLiveMessageGroupResponseBody
	GetDelete() *bool
	SetDeletor(v string) *DescribeLiveMessageGroupResponseBody
	GetDeletor() *string
	SetGroupId(v string) *DescribeLiveMessageGroupResponseBody
	GetGroupId() *string
	SetGroupInfo(v string) *DescribeLiveMessageGroupResponseBody
	GetGroupInfo() *string
	SetGroupName(v string) *DescribeLiveMessageGroupResponseBody
	GetGroupName() *string
	SetMsgAmount(v map[string]*int64) *DescribeLiveMessageGroupResponseBody
	GetMsgAmount() map[string]*int64
	SetOnlineUserCounts(v int64) *DescribeLiveMessageGroupResponseBody
	GetOnlineUserCounts() *int64
	SetRequestId(v string) *DescribeLiveMessageGroupResponseBody
	GetRequestId() *string
	SetSuperLargeGroup(v bool) *DescribeLiveMessageGroupResponseBody
	GetSuperLargeGroup() *bool
	SetTotalTimes(v int64) *DescribeLiveMessageGroupResponseBody
	GetTotalTimes() *int64
}

type DescribeLiveMessageGroupResponseBody struct {
	// The list of the group administrators.
	AdminList []*string `json:"AdminList,omitempty" xml:"AdminList,omitempty" type:"Repeated"`
	// The time when the group was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1698305471
	Createtime *int64 `json:"Createtime,omitempty" xml:"Createtime,omitempty"`
	// The ID of the group creator.
	//
	// example:
	//
	// uid1
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The time when the group was deleted. This parameter is returned only if the group was deleted.
	//
	// example:
	//
	// 1698299827
	Deletatime *int64 `json:"Deletatime,omitempty" xml:"Deletatime,omitempty"`
	// Indicates whether the group was deleted.
	//
	// example:
	//
	// false
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// The ID of the user who deleted the group. This parameter is returned only if the group was deleted.
	//
	// example:
	//
	// uid1
	Deletor *string `json:"Deletor,omitempty" xml:"Deletor,omitempty"`
	// The group ID.
	//
	// example:
	//
	// grouptest1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Additional information about the group.
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
	// The categorized message statistics. This parameter is returned only if the group exists.
	MsgAmount map[string]*int64 `json:"MsgAmount,omitempty" xml:"MsgAmount,omitempty"`
	// The number of online users in the group. This parameter is returned only if the group exists.
	//
	// example:
	//
	// 2
	OnlineUserCounts *int64 `json:"OnlineUserCounts,omitempty" xml:"OnlineUserCounts,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1815A27D-BAE2-10E6-89FD-D477951C67C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the group is a super group. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// False
	SuperLargeGroup *bool `json:"SuperLargeGroup,omitempty" xml:"SuperLargeGroup,omitempty"`
	// The total number of sessions. This parameter is returned only if the group exists.
	//
	// example:
	//
	// 50
	TotalTimes *int64 `json:"TotalTimes,omitempty" xml:"TotalTimes,omitempty"`
}

func (s DescribeLiveMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveMessageGroupResponseBody) GetAdminList() []*string {
	return s.AdminList
}

func (s *DescribeLiveMessageGroupResponseBody) GetCreatetime() *int64 {
	return s.Createtime
}

func (s *DescribeLiveMessageGroupResponseBody) GetCreatorId() *string {
	return s.CreatorId
}

func (s *DescribeLiveMessageGroupResponseBody) GetDeletatime() *int64 {
	return s.Deletatime
}

func (s *DescribeLiveMessageGroupResponseBody) GetDelete() *bool {
	return s.Delete
}

func (s *DescribeLiveMessageGroupResponseBody) GetDeletor() *string {
	return s.Deletor
}

func (s *DescribeLiveMessageGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeLiveMessageGroupResponseBody) GetGroupInfo() *string {
	return s.GroupInfo
}

func (s *DescribeLiveMessageGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeLiveMessageGroupResponseBody) GetMsgAmount() map[string]*int64 {
	return s.MsgAmount
}

func (s *DescribeLiveMessageGroupResponseBody) GetOnlineUserCounts() *int64 {
	return s.OnlineUserCounts
}

func (s *DescribeLiveMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveMessageGroupResponseBody) GetSuperLargeGroup() *bool {
	return s.SuperLargeGroup
}

func (s *DescribeLiveMessageGroupResponseBody) GetTotalTimes() *int64 {
	return s.TotalTimes
}

func (s *DescribeLiveMessageGroupResponseBody) SetAdminList(v []*string) *DescribeLiveMessageGroupResponseBody {
	s.AdminList = v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetCreatetime(v int64) *DescribeLiveMessageGroupResponseBody {
	s.Createtime = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetCreatorId(v string) *DescribeLiveMessageGroupResponseBody {
	s.CreatorId = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetDeletatime(v int64) *DescribeLiveMessageGroupResponseBody {
	s.Deletatime = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetDelete(v bool) *DescribeLiveMessageGroupResponseBody {
	s.Delete = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetDeletor(v string) *DescribeLiveMessageGroupResponseBody {
	s.Deletor = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetGroupId(v string) *DescribeLiveMessageGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetGroupInfo(v string) *DescribeLiveMessageGroupResponseBody {
	s.GroupInfo = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetGroupName(v string) *DescribeLiveMessageGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetMsgAmount(v map[string]*int64) *DescribeLiveMessageGroupResponseBody {
	s.MsgAmount = v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetOnlineUserCounts(v int64) *DescribeLiveMessageGroupResponseBody {
	s.OnlineUserCounts = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetRequestId(v string) *DescribeLiveMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetSuperLargeGroup(v bool) *DescribeLiveMessageGroupResponseBody {
	s.SuperLargeGroup = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) SetTotalTimes(v int64) *DescribeLiveMessageGroupResponseBody {
	s.TotalTimes = &v
	return s
}

func (s *DescribeLiveMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

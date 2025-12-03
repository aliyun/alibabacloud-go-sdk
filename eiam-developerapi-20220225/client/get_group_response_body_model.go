// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v int64) *GetGroupResponseBody
	GetCreateTime() *int64
	SetDescription(v string) *GetGroupResponseBody
	GetDescription() *string
	SetGroupExternalId(v string) *GetGroupResponseBody
	GetGroupExternalId() *string
	SetGroupId(v string) *GetGroupResponseBody
	GetGroupId() *string
	SetGroupName(v string) *GetGroupResponseBody
	GetGroupName() *string
	SetGroupSourceId(v string) *GetGroupResponseBody
	GetGroupSourceId() *string
	SetGroupSourceType(v string) *GetGroupResponseBody
	GetGroupSourceType() *string
	SetInstanceId(v string) *GetGroupResponseBody
	GetInstanceId() *string
	SetUpdateTime(v int64) *GetGroupResponseBody
	GetUpdateTime() *int64
}

type GetGroupResponseBody struct {
	// The time when the group was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The group description.
	//
	// example:
	//
	// description_demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The external ID of the group.
	//
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupExternalId *string `json:"groupExternalId,omitempty" xml:"groupExternalId,omitempty"`
	// The group ID.
	//
	// example:
	//
	// group_ufdsasn35ea5lmthk267xxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The group name.
	//
	// example:
	//
	// name_test
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The source ID of the group.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupSourceId *string `json:"groupSourceId,omitempty" xml:"groupSourceId,omitempty"`
	// The source type of the group. Valid values: build_in, ding_talk, ad, and ldap.
	//
	// example:
	//
	// build_in
	GroupSourceType *string `json:"groupSourceType,omitempty" xml:"groupSourceType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The time when the group was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetGroupResponseBody) GetGroupExternalId() *string {
	return s.GroupExternalId
}

func (s *GetGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGroupResponseBody) GetGroupSourceId() *string {
	return s.GroupSourceId
}

func (s *GetGroupResponseBody) GetGroupSourceType() *string {
	return s.GroupSourceType
}

func (s *GetGroupResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetGroupResponseBody) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetGroupResponseBody) SetCreateTime(v int64) *GetGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetGroupResponseBody) SetDescription(v string) *GetGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupExternalId(v string) *GetGroupResponseBody {
	s.GroupExternalId = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupId(v string) *GetGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupName(v string) *GetGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupSourceId(v string) *GetGroupResponseBody {
	s.GroupSourceId = &v
	return s
}

func (s *GetGroupResponseBody) SetGroupSourceType(v string) *GetGroupResponseBody {
	s.GroupSourceType = &v
	return s
}

func (s *GetGroupResponseBody) SetInstanceId(v string) *GetGroupResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetGroupResponseBody) SetUpdateTime(v int64) *GetGroupResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

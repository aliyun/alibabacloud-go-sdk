// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody
	GetGroup() *GetGroupResponseBodyGroup
	SetRequestId(v string) *GetGroupResponseBody
	GetRequestId() *string
}

type GetGroupResponseBody struct {
	// The information about the account group.
	Group *GetGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) GetGroup() *GetGroupResponseBodyGroup {
	return s.Group
}

func (s *GetGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGroupResponseBody) SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody {
	s.Group = v
	return s
}

func (s *GetGroupResponseBody) SetRequestId(v string) *GetGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGroupResponseBody) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGroupResponseBodyGroup struct {
	// The time at which the group was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the group.
	//
	// example:
	//
	// test_group
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The external ID of the group, which can be used to associate the group with an external system. By default, the external ID is the group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// The group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// group_name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The source ID of the group. By default, the source ID is the instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupSourceId *string `json:"GroupSourceId,omitempty" xml:"GroupSourceId,omitempty"`
	// The source type of the group. Only build_in may be returned, which indicates that the group was created in IDaaS.
	//
	// *build_in:Create By Self.
	//
	// example:
	//
	// build_in
	GroupSourceType *string `json:"GroupSourceType,omitempty" xml:"GroupSourceType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time at which the group was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s GetGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetGroupResponseBodyGroup) GetDescription() *string {
	return s.Description
}

func (s *GetGroupResponseBodyGroup) GetGroupExternalId() *string {
	return s.GroupExternalId
}

func (s *GetGroupResponseBodyGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *GetGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGroupResponseBodyGroup) GetGroupSourceId() *string {
	return s.GroupSourceId
}

func (s *GetGroupResponseBodyGroup) GetGroupSourceType() *string {
	return s.GroupSourceType
}

func (s *GetGroupResponseBodyGroup) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetGroupResponseBodyGroup) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetGroupResponseBodyGroup) SetCreateTime(v int64) *GetGroupResponseBodyGroup {
	s.CreateTime = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetDescription(v string) *GetGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupExternalId(v string) *GetGroupResponseBodyGroup {
	s.GroupExternalId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupId(v string) *GetGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupName(v string) *GetGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupSourceId(v string) *GetGroupResponseBodyGroup {
	s.GroupSourceId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupSourceType(v string) *GetGroupResponseBodyGroup {
	s.GroupSourceType = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetInstanceId(v string) *GetGroupResponseBodyGroup {
	s.InstanceId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetUpdateTime(v int64) *GetGroupResponseBodyGroup {
	s.UpdateTime = &v
	return s
}

func (s *GetGroupResponseBodyGroup) Validate() error {
	return dara.Validate(s)
}

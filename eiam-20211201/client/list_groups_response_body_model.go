// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*ListGroupsResponseBodyGroups) *ListGroupsResponseBody
	GetGroups() []*ListGroupsResponseBodyGroups
	SetRequestId(v string) *ListGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGroupsResponseBody
	GetTotalCount() *int64
}

type ListGroupsResponseBody struct {
	// The queried account groups.
	Groups []*ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. The maximum number of entries returned at a time depends on the value of PageSize.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) GetGroups() []*ListGroupsResponseBodyGroups {
	return s.Groups
}

func (s *ListGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupsResponseBody) SetGroups(v []*ListGroupsResponseBodyGroups) *ListGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotalCount(v int64) *ListGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGroupsResponseBodyGroups struct {
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
	// test group
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
	// The source ID of the group. If the group was imported from other services, this value indicates the external source ID. By default, the source ID is the instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupSourceId *string `json:"GroupSourceId,omitempty" xml:"GroupSourceId,omitempty"`
	// The source type of the group. Only build_in may be returned, which indicates that the group was created in IDaaS.
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

func (s ListGroupsResponseBodyGroups) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroups) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListGroupsResponseBodyGroups) GetDescription() *string {
	return s.Description
}

func (s *ListGroupsResponseBodyGroups) GetGroupExternalId() *string {
	return s.GroupExternalId
}

func (s *ListGroupsResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsResponseBodyGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGroupsResponseBodyGroups) GetGroupSourceId() *string {
	return s.GroupSourceId
}

func (s *ListGroupsResponseBodyGroups) GetGroupSourceType() *string {
	return s.GroupSourceType
}

func (s *ListGroupsResponseBodyGroups) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsResponseBodyGroups) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListGroupsResponseBodyGroups) SetCreateTime(v int64) *ListGroupsResponseBodyGroups {
	s.CreateTime = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetDescription(v string) *ListGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupExternalId(v string) *ListGroupsResponseBodyGroups {
	s.GroupExternalId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupId(v string) *ListGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupName(v string) *ListGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupSourceId(v string) *ListGroupsResponseBodyGroups {
	s.GroupSourceId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupSourceType(v string) *ListGroupsResponseBodyGroups {
	s.GroupSourceType = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetInstanceId(v string) *ListGroupsResponseBodyGroups {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetUpdateTime(v int64) *ListGroupsResponseBodyGroups {
	s.UpdateTime = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}

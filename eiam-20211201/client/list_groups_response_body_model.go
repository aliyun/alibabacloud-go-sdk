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
	// Account group list.
	Groups []*ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of matched entries. The maximum number of entries returned in a single request is determined by pageSize.
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
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsResponseBodyGroups struct {
	// Group creation time in Unix timestamp format, in milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Group description.
	//
	// example:
	//
	// test group
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Group external ID, used for association with external systems. Defaults to the account group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// Group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Group name.
	//
	// example:
	//
	// group_name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Group source ID. If created by importing from other sources, this is the external source ID. Defaults to the instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupSourceId *string `json:"GroupSourceId,omitempty" xml:"GroupSourceId,omitempty"`
	// Group source type. Currently, only self-built is supported. Valid values:
	//
	// - build_in: self-built.
	//
	// example:
	//
	// build_in
	GroupSourceType *string `json:"GroupSourceType,omitempty" xml:"GroupSourceType,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Group last update time in Unix timestamp format, in milliseconds.
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

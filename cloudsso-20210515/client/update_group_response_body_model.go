// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v *UpdateGroupResponseBodyGroup) *UpdateGroupResponseBody
	GetGroup() *UpdateGroupResponseBodyGroup
	SetRequestId(v string) *UpdateGroupResponseBody
	GetRequestId() *string
}

type UpdateGroupResponseBody struct {
	// The information about the group.
	Group *UpdateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F723DE01-6276-5DC4-9B1F-9CBE3E1748B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBody) GetGroup() *UpdateGroupResponseBodyGroup {
	return s.Group
}

func (s *UpdateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupResponseBody) SetGroup(v *UpdateGroupResponseBodyGroup) *UpdateGroupResponseBody {
	s.Group = v
	return s
}

func (s *UpdateGroupResponseBody) SetRequestId(v string) *UpdateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupResponseBody) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGroupResponseBodyGroup struct {
	// The time when the group was created.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the group.
	//
	// example:
	//
	// This is a group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// g-00jqzghi2n3o5hkh****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// NewTestGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the group. Valid values:
	//
	// - Manual: The group is manually created.
	//
	// - Synchronized: The group is synchronized from an external identity provider (IdP).
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The time when the group was modified.
	//
	// example:
	//
	// 2021-11-01T06:06:11Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBodyGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateGroupResponseBodyGroup) GetDescription() *string {
	return s.Description
}

func (s *UpdateGroupResponseBodyGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateGroupResponseBodyGroup) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *UpdateGroupResponseBodyGroup) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateGroupResponseBodyGroup) SetCreateTime(v string) *UpdateGroupResponseBodyGroup {
	s.CreateTime = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetDescription(v string) *UpdateGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetGroupId(v string) *UpdateGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetGroupName(v string) *UpdateGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetProvisionType(v string) *UpdateGroupResponseBodyGroup {
	s.ProvisionType = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) SetUpdateTime(v string) *UpdateGroupResponseBodyGroup {
	s.UpdateTime = &v
	return s
}

func (s *UpdateGroupResponseBodyGroup) Validate() error {
	return dara.Validate(s)
}

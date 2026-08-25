// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroup(v *CreateGroupResponseBodyGroup) *CreateGroupResponseBody
	GetGroup() *CreateGroupResponseBodyGroup
	SetRequestId(v string) *CreateGroupResponseBody
	GetRequestId() *string
}

type CreateGroupResponseBody struct {
	// The information about the group.
	Group *CreateGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 20E9650E-EC23-593E-933F-EA0D280D040C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) GetGroup() *CreateGroupResponseBodyGroup {
	return s.Group
}

func (s *CreateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupResponseBody) SetGroup(v *CreateGroupResponseBodyGroup) *CreateGroupResponseBody {
	s.Group = v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupResponseBody) Validate() error {
	if s.Group != nil {
		if err := s.Group.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateGroupResponseBodyGroup struct {
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
	// TestGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the group. The value is fixed as Manual, which indicates that the group is manually created.
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The time when the information about the group was modified.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateGroupResponseBodyGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBodyGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateGroupResponseBodyGroup) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupResponseBodyGroup) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupResponseBodyGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupResponseBodyGroup) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *CreateGroupResponseBodyGroup) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateGroupResponseBodyGroup) SetCreateTime(v string) *CreateGroupResponseBodyGroup {
	s.CreateTime = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetDescription(v string) *CreateGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetGroupId(v string) *CreateGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetGroupName(v string) *CreateGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetProvisionType(v string) *CreateGroupResponseBodyGroup {
	s.ProvisionType = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) SetUpdateTime(v string) *CreateGroupResponseBodyGroup {
	s.UpdateTime = &v
	return s
}

func (s *CreateGroupResponseBodyGroup) Validate() error {
	return dara.Validate(s)
}

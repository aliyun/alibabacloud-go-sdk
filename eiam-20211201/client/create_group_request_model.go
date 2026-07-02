// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateGroupRequest
	GetDescription() *string
	SetGroupExternalId(v string) *CreateGroupRequest
	GetGroupExternalId() *string
	SetGroupName(v string) *CreateGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *CreateGroupRequest
	GetInstanceId() *string
}

type CreateGroupRequest struct {
	// Description. The maximum length is 256 characters.
	//
	// example:
	//
	// this is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Group external ID, used for association with external systems. The default value is the group ID. The maximum length is 64 characters.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// Group name. The maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// name_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupRequest) GetGroupExternalId() *string {
	return s.GroupExternalId
}

func (s *CreateGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetGroupExternalId(v string) *CreateGroupRequest {
	s.GroupExternalId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) SetInstanceId(v string) *CreateGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateGroupRequest) Validate() error {
	return dara.Validate(s)
}

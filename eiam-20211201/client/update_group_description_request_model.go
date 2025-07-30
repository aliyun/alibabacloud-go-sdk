// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateGroupDescriptionRequest
	GetDescription() *string
	SetGroupId(v string) *UpdateGroupDescriptionRequest
	GetGroupId() *string
	SetInstanceId(v string) *UpdateGroupDescriptionRequest
	GetInstanceId() *string
}

type UpdateGroupDescriptionRequest struct {
	// The description of the account group.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateGroupDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateGroupDescriptionRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateGroupDescriptionRequest) SetDescription(v string) *UpdateGroupDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateGroupDescriptionRequest) SetGroupId(v string) *UpdateGroupDescriptionRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupDescriptionRequest) SetInstanceId(v string) *UpdateGroupDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateGroupDescriptionRequest) Validate() error {
	return dara.Validate(s)
}

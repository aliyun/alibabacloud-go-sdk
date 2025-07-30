// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyGroupRequest
	GetDescription() *string
	SetGroupId(v string) *ModifyGroupRequest
	GetGroupId() *string
	SetNewGroupName(v string) *ModifyGroupRequest
	GetNewGroupName() *string
}

type ModifyGroupRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	NewGroupName *string `json:"NewGroupName,omitempty" xml:"NewGroupName,omitempty"`
}

func (s ModifyGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyGroupRequest) GetNewGroupName() *string {
	return s.NewGroupName
}

func (s *ModifyGroupRequest) SetDescription(v string) *ModifyGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyGroupRequest) SetGroupId(v string) *ModifyGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyGroupRequest) SetNewGroupName(v string) *ModifyGroupRequest {
	s.NewGroupName = &v
	return s
}

func (s *ModifyGroupRequest) Validate() error {
	return dara.Validate(s)
}

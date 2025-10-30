// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateUserGroupRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateUserGroupRequestUpdateCommand) *UpdateUserGroupRequest
	GetUpdateCommand() *UpdateUserGroupRequestUpdateCommand
}

type UpdateUserGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId    *int64                               `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	UpdateCommand *UpdateUserGroupRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateUserGroupRequest) GetUpdateCommand() *UpdateUserGroupRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateUserGroupRequest) SetOpTenantId(v int64) *UpdateUserGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUpdateCommand(v *UpdateUserGroupRequestUpdateCommand) *UpdateUserGroupRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateUserGroupRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserGroupRequestUpdateCommand struct {
	AdminUserIdList []*string `json:"AdminUserIdList,omitempty" xml:"AdminUserIdList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13423
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateUserGroupRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequestUpdateCommand) GetAdminUserIdList() []*string {
	return s.AdminUserIdList
}

func (s *UpdateUserGroupRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserGroupRequestUpdateCommand) GetId() *string {
	return s.Id
}

func (s *UpdateUserGroupRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateUserGroupRequestUpdateCommand) SetAdminUserIdList(v []*string) *UpdateUserGroupRequestUpdateCommand {
	s.AdminUserIdList = v
	return s
}

func (s *UpdateUserGroupRequestUpdateCommand) SetDescription(v string) *UpdateUserGroupRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateUserGroupRequestUpdateCommand) SetId(v string) *UpdateUserGroupRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateUserGroupRequestUpdateCommand) SetName(v string) *UpdateUserGroupRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateUserGroupRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}

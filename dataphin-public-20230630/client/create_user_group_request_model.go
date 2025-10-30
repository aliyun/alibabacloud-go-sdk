// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateUserGroupRequestCreateCommand) *CreateUserGroupRequest
	GetCreateCommand() *CreateUserGroupRequestCreateCommand
	SetOpTenantId(v int64) *CreateUserGroupRequest
	GetOpTenantId() *int64
}

type CreateUserGroupRequest struct {
	CreateCommand *CreateUserGroupRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) GetCreateCommand() *CreateUserGroupRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateUserGroupRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateUserGroupRequest) SetCreateCommand(v *CreateUserGroupRequestCreateCommand) *CreateUserGroupRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateUserGroupRequest) SetOpTenantId(v int64) *CreateUserGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateUserGroupRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserGroupRequestCreateCommand struct {
	// example:
	//
	// true
	Active          *bool     `json:"Active,omitempty" xml:"Active,omitempty"`
	AdminUserIdList []*string `json:"AdminUserIdList,omitempty" xml:"AdminUserIdList,omitempty" type:"Repeated"`
	// example:
	//
	// xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateUserGroupRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequestCreateCommand) GetActive() *bool {
	return s.Active
}

func (s *CreateUserGroupRequestCreateCommand) GetAdminUserIdList() []*string {
	return s.AdminUserIdList
}

func (s *CreateUserGroupRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateUserGroupRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateUserGroupRequestCreateCommand) SetActive(v bool) *CreateUserGroupRequestCreateCommand {
	s.Active = &v
	return s
}

func (s *CreateUserGroupRequestCreateCommand) SetAdminUserIdList(v []*string) *CreateUserGroupRequestCreateCommand {
	s.AdminUserIdList = v
	return s
}

func (s *CreateUserGroupRequestCreateCommand) SetDescription(v string) *CreateUserGroupRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateUserGroupRequestCreateCommand) SetName(v string) *CreateUserGroupRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateUserGroupRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *UpdateProjectMemberRequest
	GetId() *int64
	SetOpTenantId(v int64) *UpdateProjectMemberRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateProjectMemberRequestUpdateCommand) *UpdateProjectMemberRequest
	GetUpdateCommand() *UpdateProjectMemberRequestUpdateCommand
}

type UpdateProjectMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 711833
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *UpdateProjectMemberRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateProjectMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateProjectMemberRequest) GetUpdateCommand() *UpdateProjectMemberRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateProjectMemberRequest) SetId(v int64) *UpdateProjectMemberRequest {
	s.Id = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetOpTenantId(v int64) *UpdateProjectMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateProjectMemberRequest) SetUpdateCommand(v *UpdateProjectMemberRequestUpdateCommand) *UpdateProjectMemberRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateProjectMemberRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateProjectMemberRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// This parameter is required.
	UserList []*UpdateProjectMemberRequestUpdateCommandUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s UpdateProjectMemberRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectMemberRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRequestUpdateCommand) GetEnv() *string {
	return s.Env
}

func (s *UpdateProjectMemberRequestUpdateCommand) GetUserList() []*UpdateProjectMemberRequestUpdateCommandUserList {
	return s.UserList
}

func (s *UpdateProjectMemberRequestUpdateCommand) SetEnv(v string) *UpdateProjectMemberRequestUpdateCommand {
	s.Env = &v
	return s
}

func (s *UpdateProjectMemberRequestUpdateCommand) SetUserList(v []*UpdateProjectMemberRequestUpdateCommandUserList) *UpdateProjectMemberRequestUpdateCommand {
	s.UserList = v
	return s
}

func (s *UpdateProjectMemberRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}

type UpdateProjectMemberRequestUpdateCommandUserList struct {
	// This parameter is required.
	RoleList []*int32 `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 30012011
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateProjectMemberRequestUpdateCommandUserList) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectMemberRequestUpdateCommandUserList) GoString() string {
	return s.String()
}

func (s *UpdateProjectMemberRequestUpdateCommandUserList) GetRoleList() []*int32 {
	return s.RoleList
}

func (s *UpdateProjectMemberRequestUpdateCommandUserList) GetUserId() *string {
	return s.UserId
}

func (s *UpdateProjectMemberRequestUpdateCommandUserList) SetRoleList(v []*int32) *UpdateProjectMemberRequestUpdateCommandUserList {
	s.RoleList = v
	return s
}

func (s *UpdateProjectMemberRequestUpdateCommandUserList) SetUserId(v string) *UpdateProjectMemberRequestUpdateCommandUserList {
	s.UserId = &v
	return s
}

func (s *UpdateProjectMemberRequestUpdateCommandUserList) Validate() error {
	return dara.Validate(s)
}

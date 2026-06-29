// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveProjectMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RemoveProjectMemberRequest
	GetId() *int64
	SetOpTenantId(v int64) *RemoveProjectMemberRequest
	GetOpTenantId() *int64
	SetRemoveCommand(v *RemoveProjectMemberRequestRemoveCommand) *RemoveProjectMemberRequest
	GetRemoveCommand() *RemoveProjectMemberRequestRemoveCommand
}

type RemoveProjectMemberRequest struct {
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 711833
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The command to remove a member.
	//
	// This parameter is required.
	RemoveCommand *RemoveProjectMemberRequestRemoveCommand `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty" type:"Struct"`
}

func (s RemoveProjectMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveProjectMemberRequest) GetId() *int64 {
	return s.Id
}

func (s *RemoveProjectMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveProjectMemberRequest) GetRemoveCommand() *RemoveProjectMemberRequestRemoveCommand {
	return s.RemoveCommand
}

func (s *RemoveProjectMemberRequest) SetId(v int64) *RemoveProjectMemberRequest {
	s.Id = &v
	return s
}

func (s *RemoveProjectMemberRequest) SetOpTenantId(v int64) *RemoveProjectMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveProjectMemberRequest) SetRemoveCommand(v *RemoveProjectMemberRequestRemoveCommand) *RemoveProjectMemberRequest {
	s.RemoveCommand = v
	return s
}

func (s *RemoveProjectMemberRequest) Validate() error {
	if s.RemoveCommand != nil {
		if err := s.RemoveCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveProjectMemberRequestRemoveCommand struct {
	// The environment identifier. Valid values: DEV and PROD.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The user IDs.
	UserIdList []*string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty" type:"Repeated"`
}

func (s RemoveProjectMemberRequestRemoveCommand) String() string {
	return dara.Prettify(s)
}

func (s RemoveProjectMemberRequestRemoveCommand) GoString() string {
	return s.String()
}

func (s *RemoveProjectMemberRequestRemoveCommand) GetEnv() *string {
	return s.Env
}

func (s *RemoveProjectMemberRequestRemoveCommand) GetUserIdList() []*string {
	return s.UserIdList
}

func (s *RemoveProjectMemberRequestRemoveCommand) SetEnv(v string) *RemoveProjectMemberRequestRemoveCommand {
	s.Env = &v
	return s
}

func (s *RemoveProjectMemberRequestRemoveCommand) SetUserIdList(v []*string) *RemoveProjectMemberRequestRemoveCommand {
	s.UserIdList = v
	return s
}

func (s *RemoveProjectMemberRequestRemoveCommand) Validate() error {
	return dara.Validate(s)
}

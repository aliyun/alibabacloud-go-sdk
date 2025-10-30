// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTenantMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveTenantMemberRequest
	GetOpTenantId() *int64
	SetRemoveCommand(v *RemoveTenantMemberRequestRemoveCommand) *RemoveTenantMemberRequest
	GetRemoveCommand() *RemoveTenantMemberRequestRemoveCommand
}

type RemoveTenantMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommand *RemoveTenantMemberRequestRemoveCommand `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty" type:"Struct"`
}

func (s RemoveTenantMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTenantMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveTenantMemberRequest) GetRemoveCommand() *RemoveTenantMemberRequestRemoveCommand {
	return s.RemoveCommand
}

func (s *RemoveTenantMemberRequest) SetOpTenantId(v int64) *RemoveTenantMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveTenantMemberRequest) SetRemoveCommand(v *RemoveTenantMemberRequestRemoveCommand) *RemoveTenantMemberRequest {
	s.RemoveCommand = v
	return s
}

func (s *RemoveTenantMemberRequest) Validate() error {
	if s.RemoveCommand != nil {
		if err := s.RemoveCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveTenantMemberRequestRemoveCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 123@xx.com
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s RemoveTenantMemberRequestRemoveCommand) String() string {
	return dara.Prettify(s)
}

func (s RemoveTenantMemberRequestRemoveCommand) GoString() string {
	return s.String()
}

func (s *RemoveTenantMemberRequestRemoveCommand) GetSourceId() *string {
	return s.SourceId
}

func (s *RemoveTenantMemberRequestRemoveCommand) SetSourceId(v string) *RemoveTenantMemberRequestRemoveCommand {
	s.SourceId = &v
	return s
}

func (s *RemoveTenantMemberRequestRemoveCommand) Validate() error {
	return dara.Validate(s)
}

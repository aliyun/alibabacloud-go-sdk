// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataServiceAppMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RemoveDataServiceAppMemberRequest
	GetOpTenantId() *int64
	SetRemoveCommand(v *RemoveDataServiceAppMemberRequestRemoveCommand) *RemoveDataServiceAppMemberRequest
	GetRemoveCommand() *RemoveDataServiceAppMemberRequestRemoveCommand
}

type RemoveDataServiceAppMemberRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	RemoveCommand *RemoveDataServiceAppMemberRequestRemoveCommand `json:"RemoveCommand,omitempty" xml:"RemoveCommand,omitempty" type:"Struct"`
}

func (s RemoveDataServiceAppMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataServiceAppMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveDataServiceAppMemberRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RemoveDataServiceAppMemberRequest) GetRemoveCommand() *RemoveDataServiceAppMemberRequestRemoveCommand {
	return s.RemoveCommand
}

func (s *RemoveDataServiceAppMemberRequest) SetOpTenantId(v int64) *RemoveDataServiceAppMemberRequest {
	s.OpTenantId = &v
	return s
}

func (s *RemoveDataServiceAppMemberRequest) SetRemoveCommand(v *RemoveDataServiceAppMemberRequestRemoveCommand) *RemoveDataServiceAppMemberRequest {
	s.RemoveCommand = v
	return s
}

func (s *RemoveDataServiceAppMemberRequest) Validate() error {
	if s.RemoveCommand != nil {
		if err := s.RemoveCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveDataServiceAppMemberRequestRemoveCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 200000000
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	MemberIds []*string `json:"MemberIds,omitempty" xml:"MemberIds,omitempty" type:"Repeated"`
}

func (s RemoveDataServiceAppMemberRequestRemoveCommand) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataServiceAppMemberRequestRemoveCommand) GoString() string {
	return s.String()
}

func (s *RemoveDataServiceAppMemberRequestRemoveCommand) GetAppId() *int32 {
	return s.AppId
}

func (s *RemoveDataServiceAppMemberRequestRemoveCommand) GetMemberIds() []*string {
	return s.MemberIds
}

func (s *RemoveDataServiceAppMemberRequestRemoveCommand) SetAppId(v int32) *RemoveDataServiceAppMemberRequestRemoveCommand {
	s.AppId = &v
	return s
}

func (s *RemoveDataServiceAppMemberRequestRemoveCommand) SetMemberIds(v []*string) *RemoveDataServiceAppMemberRequestRemoveCommand {
	s.MemberIds = v
	return s
}

func (s *RemoveDataServiceAppMemberRequestRemoveCommand) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityLevelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteSecurityLevelRequestDeleteCommand) *DeleteSecurityLevelRequest
	GetDeleteCommand() *DeleteSecurityLevelRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteSecurityLevelRequest
	GetOpTenantId() *int64
}

type DeleteSecurityLevelRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteSecurityLevelRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteSecurityLevelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityLevelRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityLevelRequest) GetDeleteCommand() *DeleteSecurityLevelRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteSecurityLevelRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteSecurityLevelRequest) SetDeleteCommand(v *DeleteSecurityLevelRequestDeleteCommand) *DeleteSecurityLevelRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteSecurityLevelRequest) SetOpTenantId(v int64) *DeleteSecurityLevelRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteSecurityLevelRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSecurityLevelRequestDeleteCommand struct {
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteSecurityLevelRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityLevelRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteSecurityLevelRequestDeleteCommand) GetIndex() *int64 {
	return s.Index
}

func (s *DeleteSecurityLevelRequestDeleteCommand) GetName() *string {
	return s.Name
}

func (s *DeleteSecurityLevelRequestDeleteCommand) SetIndex(v int64) *DeleteSecurityLevelRequestDeleteCommand {
	s.Index = &v
	return s
}

func (s *DeleteSecurityLevelRequestDeleteCommand) SetName(v string) *DeleteSecurityLevelRequestDeleteCommand {
	s.Name = &v
	return s
}

func (s *DeleteSecurityLevelRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteStandardRequestDeleteCommand) *DeleteStandardRequest
	GetDeleteCommand() *DeleteStandardRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteStandardRequest
	GetOpTenantId() *int64
}

type DeleteStandardRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteStandardRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardRequest) GetDeleteCommand() *DeleteStandardRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteStandardRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardRequest) SetDeleteCommand(v *DeleteStandardRequestDeleteCommand) *DeleteStandardRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteStandardRequest) SetOpTenantId(v int64) *DeleteStandardRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteStandardRequestDeleteCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteStandardRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteStandardRequestDeleteCommand) GetId() *int64 {
	return s.Id
}

func (s *DeleteStandardRequestDeleteCommand) SetId(v int64) *DeleteStandardRequestDeleteCommand {
	s.Id = &v
	return s
}

func (s *DeleteStandardRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}

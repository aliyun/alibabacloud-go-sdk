// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityClassifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteSecurityClassifyRequestDeleteCommand) *DeleteSecurityClassifyRequest
	GetDeleteCommand() *DeleteSecurityClassifyRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteSecurityClassifyRequest
	GetOpTenantId() *int64
}

type DeleteSecurityClassifyRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteSecurityClassifyRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteSecurityClassifyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyRequest) GetDeleteCommand() *DeleteSecurityClassifyRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteSecurityClassifyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteSecurityClassifyRequest) SetDeleteCommand(v *DeleteSecurityClassifyRequestDeleteCommand) *DeleteSecurityClassifyRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteSecurityClassifyRequest) SetOpTenantId(v int64) *DeleteSecurityClassifyRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteSecurityClassifyRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSecurityClassifyRequestDeleteCommand struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// /
	ParentPath *string `json:"ParentPath,omitempty" xml:"ParentPath,omitempty"`
}

func (s DeleteSecurityClassifyRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyRequestDeleteCommand) GetId() *int64 {
	return s.Id
}

func (s *DeleteSecurityClassifyRequestDeleteCommand) GetName() *string {
	return s.Name
}

func (s *DeleteSecurityClassifyRequestDeleteCommand) GetParentPath() *string {
	return s.ParentPath
}

func (s *DeleteSecurityClassifyRequestDeleteCommand) SetId(v int64) *DeleteSecurityClassifyRequestDeleteCommand {
	s.Id = &v
	return s
}

func (s *DeleteSecurityClassifyRequestDeleteCommand) SetName(v string) *DeleteSecurityClassifyRequestDeleteCommand {
	s.Name = &v
	return s
}

func (s *DeleteSecurityClassifyRequestDeleteCommand) SetParentPath(v string) *DeleteSecurityClassifyRequestDeleteCommand {
	s.ParentPath = &v
	return s
}

func (s *DeleteSecurityClassifyRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}

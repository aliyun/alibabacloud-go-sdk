// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordRunCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCodeCommand(v *GetOperationRecordRunCodeRequestCodeCommand) *GetOperationRecordRunCodeRequest
	GetCodeCommand() *GetOperationRecordRunCodeRequestCodeCommand
	SetOpTenantId(v int64) *GetOperationRecordRunCodeRequest
	GetOpTenantId() *int64
}

type GetOperationRecordRunCodeRequest struct {
	// The query command.
	//
	// This parameter is required.
	CodeCommand *GetOperationRecordRunCodeRequestCodeCommand `json:"CodeCommand,omitempty" xml:"CodeCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetOperationRecordRunCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordRunCodeRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRecordRunCodeRequest) GetCodeCommand() *GetOperationRecordRunCodeRequestCodeCommand {
	return s.CodeCommand
}

func (s *GetOperationRecordRunCodeRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetOperationRecordRunCodeRequest) SetCodeCommand(v *GetOperationRecordRunCodeRequestCodeCommand) *GetOperationRecordRunCodeRequest {
	s.CodeCommand = v
	return s
}

func (s *GetOperationRecordRunCodeRequest) SetOpTenantId(v int64) *GetOperationRecordRunCodeRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetOperationRecordRunCodeRequest) Validate() error {
	if s.CodeCommand != nil {
		if err := s.CodeCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOperationRecordRunCodeRequestCodeCommand struct {
	// The operation log ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 987654321
	OperationId *int64 `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetOperationRecordRunCodeRequestCodeCommand) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordRunCodeRequestCodeCommand) GoString() string {
	return s.String()
}

func (s *GetOperationRecordRunCodeRequestCodeCommand) GetOperationId() *int64 {
	return s.OperationId
}

func (s *GetOperationRecordRunCodeRequestCodeCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetOperationRecordRunCodeRequestCodeCommand) SetOperationId(v int64) *GetOperationRecordRunCodeRequestCodeCommand {
	s.OperationId = &v
	return s
}

func (s *GetOperationRecordRunCodeRequestCodeCommand) SetProjectId(v int64) *GetOperationRecordRunCodeRequestCodeCommand {
	s.ProjectId = &v
	return s
}

func (s *GetOperationRecordRunCodeRequestCodeCommand) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetailCommand(v *GetOperationRecordByIdRequestDetailCommand) *GetOperationRecordByIdRequest
	GetDetailCommand() *GetOperationRecordByIdRequestDetailCommand
	SetOpTenantId(v int64) *GetOperationRecordByIdRequest
	GetOpTenantId() *int64
}

type GetOperationRecordByIdRequest struct {
	// This parameter is required.
	DetailCommand *GetOperationRecordByIdRequestDetailCommand `json:"DetailCommand,omitempty" xml:"DetailCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetOperationRecordByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordByIdRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRecordByIdRequest) GetDetailCommand() *GetOperationRecordByIdRequestDetailCommand {
	return s.DetailCommand
}

func (s *GetOperationRecordByIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetOperationRecordByIdRequest) SetDetailCommand(v *GetOperationRecordByIdRequestDetailCommand) *GetOperationRecordByIdRequest {
	s.DetailCommand = v
	return s
}

func (s *GetOperationRecordByIdRequest) SetOpTenantId(v int64) *GetOperationRecordByIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetOperationRecordByIdRequest) Validate() error {
	if s.DetailCommand != nil {
		if err := s.DetailCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOperationRecordByIdRequestDetailCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 987654321
	OperationId *int64 `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetOperationRecordByIdRequestDetailCommand) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordByIdRequestDetailCommand) GoString() string {
	return s.String()
}

func (s *GetOperationRecordByIdRequestDetailCommand) GetOperationId() *int64 {
	return s.OperationId
}

func (s *GetOperationRecordByIdRequestDetailCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetOperationRecordByIdRequestDetailCommand) SetOperationId(v int64) *GetOperationRecordByIdRequestDetailCommand {
	s.OperationId = &v
	return s
}

func (s *GetOperationRecordByIdRequestDetailCommand) SetProjectId(v int64) *GetOperationRecordByIdRequestDetailCommand {
	s.ProjectId = &v
	return s
}

func (s *GetOperationRecordByIdRequestDetailCommand) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationRecordDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetOperationRecordDetailRequest
	GetOpTenantId() *int64
	SetRecordDetailCommand(v *GetOperationRecordDetailRequestRecordDetailCommand) *GetOperationRecordDetailRequest
	GetRecordDetailCommand() *GetOperationRecordDetailRequestRecordDetailCommand
}

type GetOperationRecordDetailRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The query command.
	//
	// This parameter is required.
	RecordDetailCommand *GetOperationRecordDetailRequestRecordDetailCommand `json:"RecordDetailCommand,omitempty" xml:"RecordDetailCommand,omitempty" type:"Struct"`
}

func (s GetOperationRecordDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordDetailRequest) GoString() string {
	return s.String()
}

func (s *GetOperationRecordDetailRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetOperationRecordDetailRequest) GetRecordDetailCommand() *GetOperationRecordDetailRequestRecordDetailCommand {
	return s.RecordDetailCommand
}

func (s *GetOperationRecordDetailRequest) SetOpTenantId(v int64) *GetOperationRecordDetailRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetOperationRecordDetailRequest) SetRecordDetailCommand(v *GetOperationRecordDetailRequestRecordDetailCommand) *GetOperationRecordDetailRequest {
	s.RecordDetailCommand = v
	return s
}

func (s *GetOperationRecordDetailRequest) Validate() error {
	if s.RecordDetailCommand != nil {
		if err := s.RecordDetailCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOperationRecordDetailRequestRecordDetailCommand struct {
	// The operation record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 987654321
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetOperationRecordDetailRequestRecordDetailCommand) String() string {
	return dara.Prettify(s)
}

func (s GetOperationRecordDetailRequestRecordDetailCommand) GoString() string {
	return s.String()
}

func (s *GetOperationRecordDetailRequestRecordDetailCommand) GetOperationId() *string {
	return s.OperationId
}

func (s *GetOperationRecordDetailRequestRecordDetailCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetOperationRecordDetailRequestRecordDetailCommand) SetOperationId(v string) *GetOperationRecordDetailRequestRecordDetailCommand {
	s.OperationId = &v
	return s
}

func (s *GetOperationRecordDetailRequestRecordDetailCommand) SetProjectId(v int64) *GetOperationRecordDetailRequestRecordDetailCommand {
	s.ProjectId = &v
	return s
}

func (s *GetOperationRecordDetailRequestRecordDetailCommand) Validate() error {
	return dara.Validate(s)
}

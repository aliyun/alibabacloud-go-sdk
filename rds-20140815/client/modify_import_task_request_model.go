// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyImportTaskRequest
	GetDBInstanceId() *string
	SetOperation(v string) *ModifyImportTaskRequest
	GetOperation() *string
	SetOwnerId(v int64) *ModifyImportTaskRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyImportTaskRequest
	GetRegionId() *string
	SetTaskId(v string) *ModifyImportTaskRequest
	GetTaskId() *string
}

type ModifyImportTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze63v2p3o3k****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CANCEL
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImportTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyImportTaskRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyImportTaskRequest) GetOperation() *string {
	return s.Operation
}

func (s *ModifyImportTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyImportTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyImportTaskRequest) SetDBInstanceId(v string) *ModifyImportTaskRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyImportTaskRequest) SetOperation(v string) *ModifyImportTaskRequest {
	s.Operation = &v
	return s
}

func (s *ModifyImportTaskRequest) SetOwnerId(v int64) *ModifyImportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyImportTaskRequest) SetRegionId(v string) *ModifyImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImportTaskRequest) SetTaskId(v string) *ModifyImportTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyImportTaskRequest) Validate() error {
	return dara.Validate(s)
}

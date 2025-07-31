// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAdHocTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *StopAdHocTaskRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *StopAdHocTaskRequest
	GetProjectId() *int64
	SetTaskId(v string) *StopAdHocTaskRequest
	GetTaskId() *string
}

type StopAdHocTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute_SQL_300000843_1611548758327
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopAdHocTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAdHocTaskRequest) GoString() string {
	return s.String()
}

func (s *StopAdHocTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *StopAdHocTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *StopAdHocTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StopAdHocTaskRequest) SetOpTenantId(v int64) *StopAdHocTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *StopAdHocTaskRequest) SetProjectId(v int64) *StopAdHocTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *StopAdHocTaskRequest) SetTaskId(v string) *StopAdHocTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopAdHocTaskRequest) Validate() error {
	return dara.Validate(s)
}

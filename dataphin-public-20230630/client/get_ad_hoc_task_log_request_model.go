// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdHocTaskLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int32) *GetAdHocTaskLogRequest
	GetOffset() *int32
	SetOpTenantId(v int64) *GetAdHocTaskLogRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetAdHocTaskLogRequest
	GetProjectId() *int64
	SetSubTaskId(v int32) *GetAdHocTaskLogRequest
	GetSubTaskId() *int32
	SetTaskId(v string) *GetAdHocTaskLogRequest
	GetTaskId() *string
}

type GetAdHocTaskLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1021
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
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
	// 0
	SubTaskId *int32 `json:"SubTaskId,omitempty" xml:"SubTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute_SQL_300000843_1611548758327
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAdHocTaskLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocTaskLogRequest) GoString() string {
	return s.String()
}

func (s *GetAdHocTaskLogRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *GetAdHocTaskLogRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAdHocTaskLogRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetAdHocTaskLogRequest) GetSubTaskId() *int32 {
	return s.SubTaskId
}

func (s *GetAdHocTaskLogRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAdHocTaskLogRequest) SetOffset(v int32) *GetAdHocTaskLogRequest {
	s.Offset = &v
	return s
}

func (s *GetAdHocTaskLogRequest) SetOpTenantId(v int64) *GetAdHocTaskLogRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAdHocTaskLogRequest) SetProjectId(v int64) *GetAdHocTaskLogRequest {
	s.ProjectId = &v
	return s
}

func (s *GetAdHocTaskLogRequest) SetSubTaskId(v int32) *GetAdHocTaskLogRequest {
	s.SubTaskId = &v
	return s
}

func (s *GetAdHocTaskLogRequest) SetTaskId(v string) *GetAdHocTaskLogRequest {
	s.TaskId = &v
	return s
}

func (s *GetAdHocTaskLogRequest) Validate() error {
	return dara.Validate(s)
}

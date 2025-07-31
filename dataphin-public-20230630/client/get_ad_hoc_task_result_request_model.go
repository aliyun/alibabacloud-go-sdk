// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdHocTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetAdHocTaskResultRequest
	GetOpTenantId() *int64
	SetProjectId(v int64) *GetAdHocTaskResultRequest
	GetProjectId() *int64
	SetSubTaskId(v int32) *GetAdHocTaskResultRequest
	GetSubTaskId() *int32
	SetTaskId(v string) *GetAdHocTaskResultRequest
	GetTaskId() *string
}

type GetAdHocTaskResultRequest struct {
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

func (s GetAdHocTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetAdHocTaskResultRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAdHocTaskResultRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetAdHocTaskResultRequest) GetSubTaskId() *int32 {
	return s.SubTaskId
}

func (s *GetAdHocTaskResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAdHocTaskResultRequest) SetOpTenantId(v int64) *GetAdHocTaskResultRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAdHocTaskResultRequest) SetProjectId(v int64) *GetAdHocTaskResultRequest {
	s.ProjectId = &v
	return s
}

func (s *GetAdHocTaskResultRequest) SetSubTaskId(v int32) *GetAdHocTaskResultRequest {
	s.SubTaskId = &v
	return s
}

func (s *GetAdHocTaskResultRequest) SetTaskId(v string) *GetAdHocTaskResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetAdHocTaskResultRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *GetStreamTaskRequest
	GetEnv() *string
	SetOpTenantId(v int64) *GetStreamTaskRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetStreamTaskRequest
	GetOpUserId() *string
	SetProjectId(v int64) *GetStreamTaskRequest
	GetProjectId() *int64
	SetTaskId(v string) *GetStreamTaskRequest
	GetTaskId() *string
}

type GetStreamTaskRequest struct {
	// The environment to query. Valid values: DEV for the development environment, and PROD for the production environment. If this parameter is not specified, the gateway defaults to PROD. The value is case-insensitive.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The ID of the Dataphin project to which the task belongs. This parameter is used for project permission verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7299592831043008
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Instance ID of the real-time task instance. Instance ID starts with t_ and is not a node ID or a file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t_8442437168737681408_20261014_8442447717177229312
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetStreamTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStreamTaskRequest) GoString() string {
	return s.String()
}

func (s *GetStreamTaskRequest) GetEnv() *string {
	return s.Env
}

func (s *GetStreamTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetStreamTaskRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetStreamTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetStreamTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetStreamTaskRequest) SetEnv(v string) *GetStreamTaskRequest {
	s.Env = &v
	return s
}

func (s *GetStreamTaskRequest) SetOpTenantId(v int64) *GetStreamTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetStreamTaskRequest) SetOpUserId(v string) *GetStreamTaskRequest {
	s.OpUserId = &v
	return s
}

func (s *GetStreamTaskRequest) SetProjectId(v int64) *GetStreamTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *GetStreamTaskRequest) SetTaskId(v string) *GetStreamTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetStreamTaskRequest) Validate() error {
	return dara.Validate(s)
}

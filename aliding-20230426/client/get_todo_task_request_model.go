// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTodoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetTodoTaskRequest
	GetTaskId() *string
	SetTenantContext(v *GetTodoTaskRequestTenantContext) *GetTodoTaskRequest
	GetTenantContext() *GetTodoTaskRequestTenantContext
}

type GetTodoTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// OPJpwtxxxx
	TaskId        *string                          `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantContext *GetTodoTaskRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetTodoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTodoTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTodoTaskRequest) GetTenantContext() *GetTodoTaskRequestTenantContext {
	return s.TenantContext
}

func (s *GetTodoTaskRequest) SetTaskId(v string) *GetTodoTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTodoTaskRequest) SetTenantContext(v *GetTodoTaskRequestTenantContext) *GetTodoTaskRequest {
	s.TenantContext = v
	return s
}

func (s *GetTodoTaskRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTodoTaskRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetTodoTaskRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetTodoTaskRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetTodoTaskRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTodoTaskRequestTenantContext) SetTenantId(v string) *GetTodoTaskRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetTodoTaskRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

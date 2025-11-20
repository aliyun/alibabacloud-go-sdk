// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskExecutorStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *UpdateTodoTaskExecutorStatusRequestTenantContext) *UpdateTodoTaskExecutorStatusRequest
	GetTenantContext() *UpdateTodoTaskExecutorStatusRequestTenantContext
	SetExecutorStatusList(v []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList) *UpdateTodoTaskExecutorStatusRequest
	GetExecutorStatusList() []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList
	SetOperatorId(v string) *UpdateTodoTaskExecutorStatusRequest
	GetOperatorId() *string
	SetTaskId(v string) *UpdateTodoTaskExecutorStatusRequest
	GetTaskId() *string
}

type UpdateTodoTaskExecutorStatusRequest struct {
	TenantContext      *UpdateTodoTaskExecutorStatusRequestTenantContext        `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	ExecutorStatusList []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList `json:"executorStatusList,omitempty" xml:"executorStatusList,omitempty" type:"Repeated"`
	// example:
	//
	// xxxx
	OperatorId *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15002141
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequest) GetTenantContext() *UpdateTodoTaskExecutorStatusRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateTodoTaskExecutorStatusRequest) GetExecutorStatusList() []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList {
	return s.ExecutorStatusList
}

func (s *UpdateTodoTaskExecutorStatusRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *UpdateTodoTaskExecutorStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetTenantContext(v *UpdateTodoTaskExecutorStatusRequestTenantContext) *UpdateTodoTaskExecutorStatusRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetExecutorStatusList(v []*UpdateTodoTaskExecutorStatusRequestExecutorStatusList) *UpdateTodoTaskExecutorStatusRequest {
	s.ExecutorStatusList = v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetOperatorId(v string) *UpdateTodoTaskExecutorStatusRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequest) SetTaskId(v string) *UpdateTodoTaskExecutorStatusRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	if s.ExecutorStatusList != nil {
		for _, item := range s.ExecutorStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTodoTaskExecutorStatusRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateTodoTaskExecutorStatusRequestTenantContext) SetTenantId(v string) *UpdateTodoTaskExecutorStatusRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

type UpdateTodoTaskExecutorStatusRequestExecutorStatusList struct {
	// example:
	//
	// userId
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
}

func (s UpdateTodoTaskExecutorStatusRequestExecutorStatusList) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskExecutorStatusRequestExecutorStatusList) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) GetId() *string {
	return s.Id
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) GetIsDone() *bool {
	return s.IsDone
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) SetId(v string) *UpdateTodoTaskExecutorStatusRequestExecutorStatusList {
	s.Id = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) SetIsDone(v bool) *UpdateTodoTaskExecutorStatusRequestExecutorStatusList {
	s.IsDone = &v
	return s
}

func (s *UpdateTodoTaskExecutorStatusRequestExecutorStatusList) Validate() error {
	return dara.Validate(s)
}

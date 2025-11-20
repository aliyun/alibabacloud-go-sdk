// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTodoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *UpdateTodoTaskRequestTenantContext) *UpdateTodoTaskRequest
	GetTenantContext() *UpdateTodoTaskRequestTenantContext
	SetDescription(v string) *UpdateTodoTaskRequest
	GetDescription() *string
	SetDone(v bool) *UpdateTodoTaskRequest
	GetDone() *bool
	SetDueTime(v int64) *UpdateTodoTaskRequest
	GetDueTime() *int64
	SetExecutorIds(v []*string) *UpdateTodoTaskRequest
	GetExecutorIds() []*string
	SetParticipantIds(v []*string) *UpdateTodoTaskRequest
	GetParticipantIds() []*string
	SetSubject(v string) *UpdateTodoTaskRequest
	GetSubject() *string
	SetTaskId(v string) *UpdateTodoTaskRequest
	GetTaskId() *string
}

type UpdateTodoTaskRequest struct {
	TenantContext *UpdateTodoTaskRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// 应用可以调用该接口发起一个钉钉待办任务，该待办事项会出现在钉钉客户端“待办”页面，需要注意的是，通过开放接口发起的待办，目前仅支持直接跳转ISV应用详情页（ISV在调该接口时需传入自身应用详情页链接）。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	Done *bool `json:"done,omitempty" xml:"done,omitempty"`
	// example:
	//
	// 1617675000000
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// example:
	//
	// []
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// example:
	//
	// []
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
	// example:
	//
	// 更新钉钉待办
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// taskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s UpdateTodoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequest) GetTenantContext() *UpdateTodoTaskRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateTodoTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTodoTaskRequest) GetDone() *bool {
	return s.Done
}

func (s *UpdateTodoTaskRequest) GetDueTime() *int64 {
	return s.DueTime
}

func (s *UpdateTodoTaskRequest) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *UpdateTodoTaskRequest) GetParticipantIds() []*string {
	return s.ParticipantIds
}

func (s *UpdateTodoTaskRequest) GetSubject() *string {
	return s.Subject
}

func (s *UpdateTodoTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateTodoTaskRequest) SetTenantContext(v *UpdateTodoTaskRequestTenantContext) *UpdateTodoTaskRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateTodoTaskRequest) SetDescription(v string) *UpdateTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDone(v bool) *UpdateTodoTaskRequest {
	s.Done = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetDueTime(v int64) *UpdateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetExecutorIds(v []*string) *UpdateTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *UpdateTodoTaskRequest) SetParticipantIds(v []*string) *UpdateTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *UpdateTodoTaskRequest) SetSubject(v string) *UpdateTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *UpdateTodoTaskRequest) SetTaskId(v string) *UpdateTodoTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateTodoTaskRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTodoTaskRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateTodoTaskRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateTodoTaskRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateTodoTaskRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateTodoTaskRequestTenantContext) SetTenantId(v string) *UpdateTodoTaskRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateTodoTaskRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

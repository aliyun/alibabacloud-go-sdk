// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTodoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePersonalTodoTaskRequest
	GetDescription() *string
	SetDueTime(v int64) *CreatePersonalTodoTaskRequest
	GetDueTime() *int64
	SetExecutorIds(v []*string) *CreatePersonalTodoTaskRequest
	GetExecutorIds() []*string
	SetNotifyConfigs(v *CreatePersonalTodoTaskRequestNotifyConfigs) *CreatePersonalTodoTaskRequest
	GetNotifyConfigs() *CreatePersonalTodoTaskRequestNotifyConfigs
	SetParticipantIds(v []*string) *CreatePersonalTodoTaskRequest
	GetParticipantIds() []*string
	SetReminderTimeStamp(v int64) *CreatePersonalTodoTaskRequest
	GetReminderTimeStamp() *int64
	SetSubject(v string) *CreatePersonalTodoTaskRequest
	GetSubject() *string
	SetTenantContext(v *CreatePersonalTodoTaskRequestTenantContext) *CreatePersonalTodoTaskRequest
	GetTenantContext() *CreatePersonalTodoTaskRequestTenantContext
}

type CreatePersonalTodoTaskRequest struct {
	// example:
	//
	// 待办备注信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1703750708595
	DueTime *int64 `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [ "012345" ]
	ExecutorIds   []*string                                   `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	NotifyConfigs *CreatePersonalTodoTaskRequestNotifyConfigs `json:"NotifyConfigs,omitempty" xml:"NotifyConfigs,omitempty" type:"Struct"`
	// example:
	//
	// [ "012345" ]
	ParticipantIds []*string `json:"ParticipantIds,omitempty" xml:"ParticipantIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1703750708595
	ReminderTimeStamp *int64 `json:"ReminderTimeStamp,omitempty" xml:"ReminderTimeStamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 待办标题
	Subject       *string                                     `json:"Subject,omitempty" xml:"Subject,omitempty"`
	TenantContext *CreatePersonalTodoTaskRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s CreatePersonalTodoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalTodoTaskRequest) GetDueTime() *int64 {
	return s.DueTime
}

func (s *CreatePersonalTodoTaskRequest) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *CreatePersonalTodoTaskRequest) GetNotifyConfigs() *CreatePersonalTodoTaskRequestNotifyConfigs {
	return s.NotifyConfigs
}

func (s *CreatePersonalTodoTaskRequest) GetParticipantIds() []*string {
	return s.ParticipantIds
}

func (s *CreatePersonalTodoTaskRequest) GetReminderTimeStamp() *int64 {
	return s.ReminderTimeStamp
}

func (s *CreatePersonalTodoTaskRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreatePersonalTodoTaskRequest) GetTenantContext() *CreatePersonalTodoTaskRequestTenantContext {
	return s.TenantContext
}

func (s *CreatePersonalTodoTaskRequest) SetDescription(v string) *CreatePersonalTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetDueTime(v int64) *CreatePersonalTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetExecutorIds(v []*string) *CreatePersonalTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetNotifyConfigs(v *CreatePersonalTodoTaskRequestNotifyConfigs) *CreatePersonalTodoTaskRequest {
	s.NotifyConfigs = v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetParticipantIds(v []*string) *CreatePersonalTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetReminderTimeStamp(v int64) *CreatePersonalTodoTaskRequest {
	s.ReminderTimeStamp = &v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetSubject(v string) *CreatePersonalTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *CreatePersonalTodoTaskRequest) SetTenantContext(v *CreatePersonalTodoTaskRequestTenantContext) *CreatePersonalTodoTaskRequest {
	s.TenantContext = v
	return s
}

func (s *CreatePersonalTodoTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePersonalTodoTaskRequestNotifyConfigs struct {
	// example:
	//
	// 1
	DingNotify *string `json:"DingNotify,omitempty" xml:"DingNotify,omitempty"`
}

func (s CreatePersonalTodoTaskRequestNotifyConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTodoTaskRequestNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskRequestNotifyConfigs) GetDingNotify() *string {
	return s.DingNotify
}

func (s *CreatePersonalTodoTaskRequestNotifyConfigs) SetDingNotify(v string) *CreatePersonalTodoTaskRequestNotifyConfigs {
	s.DingNotify = &v
	return s
}

func (s *CreatePersonalTodoTaskRequestNotifyConfigs) Validate() error {
	return dara.Validate(s)
}

type CreatePersonalTodoTaskRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalTodoTaskRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTodoTaskRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreatePersonalTodoTaskRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalTodoTaskRequestTenantContext) SetTenantId(v string) *CreatePersonalTodoTaskRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalTodoTaskRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDingtalkPersonalTodoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDingtalkPersonalTodoTaskRequest
	GetDescription() *string
	SetDueTime(v int64) *CreateDingtalkPersonalTodoTaskRequest
	GetDueTime() *int64
	SetExecutorIds(v []*string) *CreateDingtalkPersonalTodoTaskRequest
	GetExecutorIds() []*string
	SetNotifyConfigs(v *CreateDingtalkPersonalTodoTaskRequestNotifyConfigs) *CreateDingtalkPersonalTodoTaskRequest
	GetNotifyConfigs() *CreateDingtalkPersonalTodoTaskRequestNotifyConfigs
	SetParticipantIds(v []*string) *CreateDingtalkPersonalTodoTaskRequest
	GetParticipantIds() []*string
	SetSubject(v string) *CreateDingtalkPersonalTodoTaskRequest
	GetSubject() *string
	SetTenantContext(v *CreateDingtalkPersonalTodoTaskRequestTenantContext) *CreateDingtalkPersonalTodoTaskRequest
	GetTenantContext() *CreateDingtalkPersonalTodoTaskRequestTenantContext
	SetUserToken(v string) *CreateDingtalkPersonalTodoTaskRequest
	GetUserToken() *string
}

type CreateDingtalkPersonalTodoTaskRequest struct {
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
	ExecutorIds   []*string                                           `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	NotifyConfigs *CreateDingtalkPersonalTodoTaskRequestNotifyConfigs `json:"NotifyConfigs,omitempty" xml:"NotifyConfigs,omitempty" type:"Struct"`
	// example:
	//
	// [ "012345" ]
	ParticipantIds []*string `json:"ParticipantIds,omitempty" xml:"ParticipantIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 待办标题
	Subject       *string                                             `json:"Subject,omitempty" xml:"Subject,omitempty"`
	TenantContext *CreateDingtalkPersonalTodoTaskRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// 用户token
	UserToken *string `json:"UserToken,omitempty" xml:"UserToken,omitempty"`
}

func (s CreateDingtalkPersonalTodoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDingtalkPersonalTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDingtalkPersonalTodoTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDingtalkPersonalTodoTaskRequest) GetDueTime() *int64 {
	return s.DueTime
}

func (s *CreateDingtalkPersonalTodoTaskRequest) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *CreateDingtalkPersonalTodoTaskRequest) GetNotifyConfigs() *CreateDingtalkPersonalTodoTaskRequestNotifyConfigs {
	return s.NotifyConfigs
}

func (s *CreateDingtalkPersonalTodoTaskRequest) GetParticipantIds() []*string {
	return s.ParticipantIds
}

func (s *CreateDingtalkPersonalTodoTaskRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreateDingtalkPersonalTodoTaskRequest) GetTenantContext() *CreateDingtalkPersonalTodoTaskRequestTenantContext {
	return s.TenantContext
}

func (s *CreateDingtalkPersonalTodoTaskRequest) GetUserToken() *string {
	return s.UserToken
}

func (s *CreateDingtalkPersonalTodoTaskRequest) SetDescription(v string) *CreateDingtalkPersonalTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequest) SetDueTime(v int64) *CreateDingtalkPersonalTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequest) SetExecutorIds(v []*string) *CreateDingtalkPersonalTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequest) SetNotifyConfigs(v *CreateDingtalkPersonalTodoTaskRequestNotifyConfigs) *CreateDingtalkPersonalTodoTaskRequest {
	s.NotifyConfigs = v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequest) SetParticipantIds(v []*string) *CreateDingtalkPersonalTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequest) SetSubject(v string) *CreateDingtalkPersonalTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequest) SetTenantContext(v *CreateDingtalkPersonalTodoTaskRequestTenantContext) *CreateDingtalkPersonalTodoTaskRequest {
	s.TenantContext = v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequest) SetUserToken(v string) *CreateDingtalkPersonalTodoTaskRequest {
	s.UserToken = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDingtalkPersonalTodoTaskRequestNotifyConfigs struct {
	// example:
	//
	// 1
	DingNotify *string `json:"DingNotify,omitempty" xml:"DingNotify,omitempty"`
}

func (s CreateDingtalkPersonalTodoTaskRequestNotifyConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateDingtalkPersonalTodoTaskRequestNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateDingtalkPersonalTodoTaskRequestNotifyConfigs) GetDingNotify() *string {
	return s.DingNotify
}

func (s *CreateDingtalkPersonalTodoTaskRequestNotifyConfigs) SetDingNotify(v string) *CreateDingtalkPersonalTodoTaskRequestNotifyConfigs {
	s.DingNotify = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequestNotifyConfigs) Validate() error {
	return dara.Validate(s)
}

type CreateDingtalkPersonalTodoTaskRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateDingtalkPersonalTodoTaskRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateDingtalkPersonalTodoTaskRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateDingtalkPersonalTodoTaskRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateDingtalkPersonalTodoTaskRequestTenantContext) SetTenantId(v string) *CreateDingtalkPersonalTodoTaskRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateDingtalkPersonalTodoTaskRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

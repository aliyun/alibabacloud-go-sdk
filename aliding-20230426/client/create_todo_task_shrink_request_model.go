// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *CreateTodoTaskShrinkRequest
	GetTenantContextShrink() *string
	SetActionListShrink(v string) *CreateTodoTaskShrinkRequest
	GetActionListShrink() *string
	SetContentFieldListShrink(v string) *CreateTodoTaskShrinkRequest
	GetContentFieldListShrink() *string
	SetCreatorId(v string) *CreateTodoTaskShrinkRequest
	GetCreatorId() *string
	SetDescription(v string) *CreateTodoTaskShrinkRequest
	GetDescription() *string
	SetDetailUrlShrink(v string) *CreateTodoTaskShrinkRequest
	GetDetailUrlShrink() *string
	SetDueTime(v int64) *CreateTodoTaskShrinkRequest
	GetDueTime() *int64
	SetExecutorIdsShrink(v string) *CreateTodoTaskShrinkRequest
	GetExecutorIdsShrink() *string
	SetIsOnlyShowExecutor(v bool) *CreateTodoTaskShrinkRequest
	GetIsOnlyShowExecutor() *bool
	SetNotifyConfigsShrink(v string) *CreateTodoTaskShrinkRequest
	GetNotifyConfigsShrink() *string
	SetOperatorId(v string) *CreateTodoTaskShrinkRequest
	GetOperatorId() *string
	SetParticipantIdsShrink(v string) *CreateTodoTaskShrinkRequest
	GetParticipantIdsShrink() *string
	SetPriority(v int32) *CreateTodoTaskShrinkRequest
	GetPriority() *int32
	SetSourceId(v string) *CreateTodoTaskShrinkRequest
	GetSourceId() *string
	SetSubject(v string) *CreateTodoTaskShrinkRequest
	GetSubject() *string
}

type CreateTodoTaskShrinkRequest struct {
	TenantContextShrink    *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	ActionListShrink       *string `json:"actionList,omitempty" xml:"actionList,omitempty"`
	ContentFieldListShrink *string `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty"`
	// example:
	//
	// PUoiinWIpa2yH2ymhiiGiP6g
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 应用可以调用该接口发起一个钉钉待办任务，该待办事项会出现在钉钉客户端“待办”页面，需要注意的是，通过开放接口发起的待办，目前仅支持直接跳转ISV应用详情页（ISV在调该接口时需传入自身应用详情页链接）。
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrlShrink *string `json:"detailUrl,omitempty" xml:"detailUrl,omitempty"`
	// example:
	//
	// 1617675000000
	DueTime           *int64  `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIdsShrink *string `json:"executorIds,omitempty" xml:"executorIds,omitempty"`
	// example:
	//
	// true
	IsOnlyShowExecutor  *bool   `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	NotifyConfigsShrink *string `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty"`
	// example:
	//
	// 12345
	OperatorId           *string `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	ParticipantIdsShrink *string `json:"participantIds,omitempty" xml:"participantIds,omitempty"`
	// example:
	//
	// 20
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// example:
	//
	// isv_dingtalkTodo1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 接入钉钉待办
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
}

func (s CreateTodoTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateTodoTaskShrinkRequest) GetActionListShrink() *string {
	return s.ActionListShrink
}

func (s *CreateTodoTaskShrinkRequest) GetContentFieldListShrink() *string {
	return s.ContentFieldListShrink
}

func (s *CreateTodoTaskShrinkRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateTodoTaskShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTodoTaskShrinkRequest) GetDetailUrlShrink() *string {
	return s.DetailUrlShrink
}

func (s *CreateTodoTaskShrinkRequest) GetDueTime() *int64 {
	return s.DueTime
}

func (s *CreateTodoTaskShrinkRequest) GetExecutorIdsShrink() *string {
	return s.ExecutorIdsShrink
}

func (s *CreateTodoTaskShrinkRequest) GetIsOnlyShowExecutor() *bool {
	return s.IsOnlyShowExecutor
}

func (s *CreateTodoTaskShrinkRequest) GetNotifyConfigsShrink() *string {
	return s.NotifyConfigsShrink
}

func (s *CreateTodoTaskShrinkRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *CreateTodoTaskShrinkRequest) GetParticipantIdsShrink() *string {
	return s.ParticipantIdsShrink
}

func (s *CreateTodoTaskShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTodoTaskShrinkRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateTodoTaskShrinkRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreateTodoTaskShrinkRequest) SetTenantContextShrink(v string) *CreateTodoTaskShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetActionListShrink(v string) *CreateTodoTaskShrinkRequest {
	s.ActionListShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetContentFieldListShrink(v string) *CreateTodoTaskShrinkRequest {
	s.ContentFieldListShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetCreatorId(v string) *CreateTodoTaskShrinkRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetDescription(v string) *CreateTodoTaskShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetDetailUrlShrink(v string) *CreateTodoTaskShrinkRequest {
	s.DetailUrlShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetDueTime(v int64) *CreateTodoTaskShrinkRequest {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetExecutorIdsShrink(v string) *CreateTodoTaskShrinkRequest {
	s.ExecutorIdsShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetIsOnlyShowExecutor(v bool) *CreateTodoTaskShrinkRequest {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetNotifyConfigsShrink(v string) *CreateTodoTaskShrinkRequest {
	s.NotifyConfigsShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetOperatorId(v string) *CreateTodoTaskShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetParticipantIdsShrink(v string) *CreateTodoTaskShrinkRequest {
	s.ParticipantIdsShrink = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetPriority(v int32) *CreateTodoTaskShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetSourceId(v string) *CreateTodoTaskShrinkRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) SetSubject(v string) *CreateTodoTaskShrinkRequest {
	s.Subject = &v
	return s
}

func (s *CreateTodoTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

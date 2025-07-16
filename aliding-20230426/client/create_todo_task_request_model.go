// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTodoTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *CreateTodoTaskRequestTenantContext) *CreateTodoTaskRequest
	GetTenantContext() *CreateTodoTaskRequestTenantContext
	SetActionList(v []*CreateTodoTaskRequestActionList) *CreateTodoTaskRequest
	GetActionList() []*CreateTodoTaskRequestActionList
	SetContentFieldList(v []*CreateTodoTaskRequestContentFieldList) *CreateTodoTaskRequest
	GetContentFieldList() []*CreateTodoTaskRequestContentFieldList
	SetCreatorId(v string) *CreateTodoTaskRequest
	GetCreatorId() *string
	SetDescription(v string) *CreateTodoTaskRequest
	GetDescription() *string
	SetDetailUrl(v *CreateTodoTaskRequestDetailUrl) *CreateTodoTaskRequest
	GetDetailUrl() *CreateTodoTaskRequestDetailUrl
	SetDueTime(v int64) *CreateTodoTaskRequest
	GetDueTime() *int64
	SetExecutorIds(v []*string) *CreateTodoTaskRequest
	GetExecutorIds() []*string
	SetIsOnlyShowExecutor(v bool) *CreateTodoTaskRequest
	GetIsOnlyShowExecutor() *bool
	SetNotifyConfigs(v *CreateTodoTaskRequestNotifyConfigs) *CreateTodoTaskRequest
	GetNotifyConfigs() *CreateTodoTaskRequestNotifyConfigs
	SetOperatorId(v string) *CreateTodoTaskRequest
	GetOperatorId() *string
	SetParticipantIds(v []*string) *CreateTodoTaskRequest
	GetParticipantIds() []*string
	SetPriority(v int32) *CreateTodoTaskRequest
	GetPriority() *int32
	SetSourceId(v string) *CreateTodoTaskRequest
	GetSourceId() *string
	SetSubject(v string) *CreateTodoTaskRequest
	GetSubject() *string
}

type CreateTodoTaskRequest struct {
	TenantContext    *CreateTodoTaskRequestTenantContext      `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	ActionList       []*CreateTodoTaskRequestActionList       `json:"actionList,omitempty" xml:"actionList,omitempty" type:"Repeated"`
	ContentFieldList []*CreateTodoTaskRequestContentFieldList `json:"contentFieldList,omitempty" xml:"contentFieldList,omitempty" type:"Repeated"`
	// example:
	//
	// PUoiinWIpa2yH2ymhiiGiP6g
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// 应用可以调用该接口发起一个钉钉待办任务，该待办事项会出现在钉钉客户端“待办”页面，需要注意的是，通过开放接口发起的待办，目前仅支持直接跳转ISV应用详情页（ISV在调该接口时需传入自身应用详情页链接）。
	Description *string                         `json:"description,omitempty" xml:"description,omitempty"`
	DetailUrl   *CreateTodoTaskRequestDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// example:
	//
	// 1617675000000
	DueTime     *int64    `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	ExecutorIds []*string `json:"executorIds,omitempty" xml:"executorIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IsOnlyShowExecutor *bool                               `json:"isOnlyShowExecutor,omitempty" xml:"isOnlyShowExecutor,omitempty"`
	NotifyConfigs      *CreateTodoTaskRequestNotifyConfigs `json:"notifyConfigs,omitempty" xml:"notifyConfigs,omitempty" type:"Struct"`
	// example:
	//
	// 12345
	OperatorId     *string   `json:"operatorId,omitempty" xml:"operatorId,omitempty"`
	ParticipantIds []*string `json:"participantIds,omitempty" xml:"participantIds,omitempty" type:"Repeated"`
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

func (s CreateTodoTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequest) GetTenantContext() *CreateTodoTaskRequestTenantContext {
	return s.TenantContext
}

func (s *CreateTodoTaskRequest) GetActionList() []*CreateTodoTaskRequestActionList {
	return s.ActionList
}

func (s *CreateTodoTaskRequest) GetContentFieldList() []*CreateTodoTaskRequestContentFieldList {
	return s.ContentFieldList
}

func (s *CreateTodoTaskRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *CreateTodoTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTodoTaskRequest) GetDetailUrl() *CreateTodoTaskRequestDetailUrl {
	return s.DetailUrl
}

func (s *CreateTodoTaskRequest) GetDueTime() *int64 {
	return s.DueTime
}

func (s *CreateTodoTaskRequest) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *CreateTodoTaskRequest) GetIsOnlyShowExecutor() *bool {
	return s.IsOnlyShowExecutor
}

func (s *CreateTodoTaskRequest) GetNotifyConfigs() *CreateTodoTaskRequestNotifyConfigs {
	return s.NotifyConfigs
}

func (s *CreateTodoTaskRequest) GetOperatorId() *string {
	return s.OperatorId
}

func (s *CreateTodoTaskRequest) GetParticipantIds() []*string {
	return s.ParticipantIds
}

func (s *CreateTodoTaskRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateTodoTaskRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateTodoTaskRequest) GetSubject() *string {
	return s.Subject
}

func (s *CreateTodoTaskRequest) SetTenantContext(v *CreateTodoTaskRequestTenantContext) *CreateTodoTaskRequest {
	s.TenantContext = v
	return s
}

func (s *CreateTodoTaskRequest) SetActionList(v []*CreateTodoTaskRequestActionList) *CreateTodoTaskRequest {
	s.ActionList = v
	return s
}

func (s *CreateTodoTaskRequest) SetContentFieldList(v []*CreateTodoTaskRequestContentFieldList) *CreateTodoTaskRequest {
	s.ContentFieldList = v
	return s
}

func (s *CreateTodoTaskRequest) SetCreatorId(v string) *CreateTodoTaskRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetDescription(v string) *CreateTodoTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTodoTaskRequest) SetDetailUrl(v *CreateTodoTaskRequestDetailUrl) *CreateTodoTaskRequest {
	s.DetailUrl = v
	return s
}

func (s *CreateTodoTaskRequest) SetDueTime(v int64) *CreateTodoTaskRequest {
	s.DueTime = &v
	return s
}

func (s *CreateTodoTaskRequest) SetExecutorIds(v []*string) *CreateTodoTaskRequest {
	s.ExecutorIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetIsOnlyShowExecutor(v bool) *CreateTodoTaskRequest {
	s.IsOnlyShowExecutor = &v
	return s
}

func (s *CreateTodoTaskRequest) SetNotifyConfigs(v *CreateTodoTaskRequestNotifyConfigs) *CreateTodoTaskRequest {
	s.NotifyConfigs = v
	return s
}

func (s *CreateTodoTaskRequest) SetOperatorId(v string) *CreateTodoTaskRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetParticipantIds(v []*string) *CreateTodoTaskRequest {
	s.ParticipantIds = v
	return s
}

func (s *CreateTodoTaskRequest) SetPriority(v int32) *CreateTodoTaskRequest {
	s.Priority = &v
	return s
}

func (s *CreateTodoTaskRequest) SetSourceId(v string) *CreateTodoTaskRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTodoTaskRequest) SetSubject(v string) *CreateTodoTaskRequest {
	s.Subject = &v
	return s
}

func (s *CreateTodoTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTodoTaskRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateTodoTaskRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateTodoTaskRequestTenantContext) SetTenantId(v string) *CreateTodoTaskRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateTodoTaskRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

type CreateTodoTaskRequestActionList struct {
	ActionKey       *string                               `json:"actionKey,omitempty" xml:"actionKey,omitempty"`
	ActionType      *int32                                `json:"actionType,omitempty" xml:"actionType,omitempty"`
	ButtonStyleType *int32                                `json:"buttonStyleType,omitempty" xml:"buttonStyleType,omitempty"`
	Param           *CreateTodoTaskRequestActionListParam `json:"param,omitempty" xml:"param,omitempty" type:"Struct"`
	PcUrl           *string                               `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
	Title           *string                               `json:"title,omitempty" xml:"title,omitempty"`
	Url             *string                               `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateTodoTaskRequestActionList) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskRequestActionList) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestActionList) GetActionKey() *string {
	return s.ActionKey
}

func (s *CreateTodoTaskRequestActionList) GetActionType() *int32 {
	return s.ActionType
}

func (s *CreateTodoTaskRequestActionList) GetButtonStyleType() *int32 {
	return s.ButtonStyleType
}

func (s *CreateTodoTaskRequestActionList) GetParam() *CreateTodoTaskRequestActionListParam {
	return s.Param
}

func (s *CreateTodoTaskRequestActionList) GetPcUrl() *string {
	return s.PcUrl
}

func (s *CreateTodoTaskRequestActionList) GetTitle() *string {
	return s.Title
}

func (s *CreateTodoTaskRequestActionList) GetUrl() *string {
	return s.Url
}

func (s *CreateTodoTaskRequestActionList) SetActionKey(v string) *CreateTodoTaskRequestActionList {
	s.ActionKey = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetActionType(v int32) *CreateTodoTaskRequestActionList {
	s.ActionType = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetButtonStyleType(v int32) *CreateTodoTaskRequestActionList {
	s.ButtonStyleType = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetParam(v *CreateTodoTaskRequestActionListParam) *CreateTodoTaskRequestActionList {
	s.Param = v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetPcUrl(v string) *CreateTodoTaskRequestActionList {
	s.PcUrl = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetTitle(v string) *CreateTodoTaskRequestActionList {
	s.Title = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) SetUrl(v string) *CreateTodoTaskRequestActionList {
	s.Url = &v
	return s
}

func (s *CreateTodoTaskRequestActionList) Validate() error {
	return dara.Validate(s)
}

type CreateTodoTaskRequestActionListParam struct {
	Body   *string            `json:"body,omitempty" xml:"body,omitempty"`
	Header map[string]*string `json:"header,omitempty" xml:"header,omitempty"`
}

func (s CreateTodoTaskRequestActionListParam) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskRequestActionListParam) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestActionListParam) GetBody() *string {
	return s.Body
}

func (s *CreateTodoTaskRequestActionListParam) GetHeader() map[string]*string {
	return s.Header
}

func (s *CreateTodoTaskRequestActionListParam) SetBody(v string) *CreateTodoTaskRequestActionListParam {
	s.Body = &v
	return s
}

func (s *CreateTodoTaskRequestActionListParam) SetHeader(v map[string]*string) *CreateTodoTaskRequestActionListParam {
	s.Header = v
	return s
}

func (s *CreateTodoTaskRequestActionListParam) Validate() error {
	return dara.Validate(s)
}

type CreateTodoTaskRequestContentFieldList struct {
	// fieldKey
	//
	// example:
	//
	// fieldKey
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// fieldValue
	//
	// example:
	//
	// fieldValue
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s CreateTodoTaskRequestContentFieldList) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskRequestContentFieldList) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestContentFieldList) GetFieldKey() *string {
	return s.FieldKey
}

func (s *CreateTodoTaskRequestContentFieldList) GetFieldValue() *string {
	return s.FieldValue
}

func (s *CreateTodoTaskRequestContentFieldList) SetFieldKey(v string) *CreateTodoTaskRequestContentFieldList {
	s.FieldKey = &v
	return s
}

func (s *CreateTodoTaskRequestContentFieldList) SetFieldValue(v string) *CreateTodoTaskRequestContentFieldList {
	s.FieldValue = &v
	return s
}

func (s *CreateTodoTaskRequestContentFieldList) Validate() error {
	return dara.Validate(s)
}

type CreateTodoTaskRequestDetailUrl struct {
	// example:
	//
	// https://www.dingtalk.com
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s CreateTodoTaskRequestDetailUrl) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskRequestDetailUrl) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestDetailUrl) GetAppUrl() *string {
	return s.AppUrl
}

func (s *CreateTodoTaskRequestDetailUrl) GetPcUrl() *string {
	return s.PcUrl
}

func (s *CreateTodoTaskRequestDetailUrl) SetAppUrl(v string) *CreateTodoTaskRequestDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *CreateTodoTaskRequestDetailUrl) SetPcUrl(v string) *CreateTodoTaskRequestDetailUrl {
	s.PcUrl = &v
	return s
}

func (s *CreateTodoTaskRequestDetailUrl) Validate() error {
	return dara.Validate(s)
}

type CreateTodoTaskRequestNotifyConfigs struct {
	// example:
	//
	// 1
	DingNotify *string `json:"dingNotify,omitempty" xml:"dingNotify,omitempty"`
}

func (s CreateTodoTaskRequestNotifyConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateTodoTaskRequestNotifyConfigs) GoString() string {
	return s.String()
}

func (s *CreateTodoTaskRequestNotifyConfigs) GetDingNotify() *string {
	return s.DingNotify
}

func (s *CreateTodoTaskRequestNotifyConfigs) SetDingNotify(v string) *CreateTodoTaskRequestNotifyConfigs {
	s.DingNotify = &v
	return s
}

func (s *CreateTodoTaskRequestNotifyConfigs) Validate() error {
	return dara.Validate(s)
}

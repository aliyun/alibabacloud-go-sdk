// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMultipleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityContentState(v interface{}) *PushMultipleRequest
	GetActivityContentState() interface{}
	SetActivityEvent(v string) *PushMultipleRequest
	GetActivityEvent() *string
	SetAppId(v string) *PushMultipleRequest
	GetAppId() *string
	SetChannelId(v string) *PushMultipleRequest
	GetChannelId() *string
	SetClassification(v string) *PushMultipleRequest
	GetClassification() *string
	SetDeliveryType(v int64) *PushMultipleRequest
	GetDeliveryType() *int64
	SetDismissalDate(v int64) *PushMultipleRequest
	GetDismissalDate() *int64
	SetExpiredSeconds(v int64) *PushMultipleRequest
	GetExpiredSeconds() *int64
	SetExtendedParams(v string) *PushMultipleRequest
	GetExtendedParams() *string
	SetMiChannelId(v string) *PushMultipleRequest
	GetMiChannelId() *string
	SetNotifyType(v string) *PushMultipleRequest
	GetNotifyType() *string
	SetPushAction(v int64) *PushMultipleRequest
	GetPushAction() *int64
	SetSilent(v int64) *PushMultipleRequest
	GetSilent() *int64
	SetStrategyContent(v string) *PushMultipleRequest
	GetStrategyContent() *string
	SetStrategyType(v int32) *PushMultipleRequest
	GetStrategyType() *int32
	SetTargetMsg(v []*PushMultipleRequestTargetMsg) *PushMultipleRequest
	GetTargetMsg() []*PushMultipleRequestTargetMsg
	SetTaskName(v string) *PushMultipleRequest
	GetTaskName() *string
	SetTemplateName(v string) *PushMultipleRequest
	GetTemplateName() *string
	SetTenantId(v string) *PushMultipleRequest
	GetTenantId() *string
	SetThirdChannelCategory(v map[string]interface{}) *PushMultipleRequest
	GetThirdChannelCategory() map[string]interface{}
	SetTransparentMessagePayload(v interface{}) *PushMultipleRequest
	GetTransparentMessagePayload() interface{}
	SetTransparentMessageUrgency(v string) *PushMultipleRequest
	GetTransparentMessageUrgency() *string
	SetWorkspaceId(v string) *PushMultipleRequest
	GetWorkspaceId() *string
}

type PushMultipleRequest struct {
	ActivityContentState interface{} `json:"ActivityContentState,omitempty" xml:"ActivityContentState,omitempty"`
	ActivityEvent        *string     `json:"ActivityEvent,omitempty" xml:"ActivityEvent,omitempty"`
	// This parameter is required.
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId      *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// This parameter is required.
	DeliveryType  *int64 `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	DismissalDate *int64 `json:"DismissalDate,omitempty" xml:"DismissalDate,omitempty"`
	// This parameter is required.
	ExpiredSeconds  *int64  `json:"ExpiredSeconds,omitempty" xml:"ExpiredSeconds,omitempty"`
	ExtendedParams  *string `json:"ExtendedParams,omitempty" xml:"ExtendedParams,omitempty"`
	MiChannelId     *string `json:"MiChannelId,omitempty" xml:"MiChannelId,omitempty"`
	NotifyType      *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	PushAction      *int64  `json:"PushAction,omitempty" xml:"PushAction,omitempty"`
	Silent          *int64  `json:"Silent,omitempty" xml:"Silent,omitempty"`
	StrategyContent *string `json:"StrategyContent,omitempty" xml:"StrategyContent,omitempty"`
	StrategyType    *int32  `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// This parameter is required.
	TargetMsg []*PushMultipleRequestTargetMsg `json:"TargetMsg,omitempty" xml:"TargetMsg,omitempty" type:"Repeated"`
	TaskName  *string                         `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	TemplateName              *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantId                  *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdChannelCategory      map[string]interface{} `json:"ThirdChannelCategory,omitempty" xml:"ThirdChannelCategory,omitempty"`
	TransparentMessagePayload interface{}            `json:"TransparentMessagePayload,omitempty" xml:"TransparentMessagePayload,omitempty"`
	TransparentMessageUrgency *string                `json:"TransparentMessageUrgency,omitempty" xml:"TransparentMessageUrgency,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushMultipleRequest) String() string {
	return dara.Prettify(s)
}

func (s PushMultipleRequest) GoString() string {
	return s.String()
}

func (s *PushMultipleRequest) GetActivityContentState() interface{} {
	return s.ActivityContentState
}

func (s *PushMultipleRequest) GetActivityEvent() *string {
	return s.ActivityEvent
}

func (s *PushMultipleRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushMultipleRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *PushMultipleRequest) GetClassification() *string {
	return s.Classification
}

func (s *PushMultipleRequest) GetDeliveryType() *int64 {
	return s.DeliveryType
}

func (s *PushMultipleRequest) GetDismissalDate() *int64 {
	return s.DismissalDate
}

func (s *PushMultipleRequest) GetExpiredSeconds() *int64 {
	return s.ExpiredSeconds
}

func (s *PushMultipleRequest) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushMultipleRequest) GetMiChannelId() *string {
	return s.MiChannelId
}

func (s *PushMultipleRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *PushMultipleRequest) GetPushAction() *int64 {
	return s.PushAction
}

func (s *PushMultipleRequest) GetSilent() *int64 {
	return s.Silent
}

func (s *PushMultipleRequest) GetStrategyContent() *string {
	return s.StrategyContent
}

func (s *PushMultipleRequest) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *PushMultipleRequest) GetTargetMsg() []*PushMultipleRequestTargetMsg {
	return s.TargetMsg
}

func (s *PushMultipleRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *PushMultipleRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PushMultipleRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushMultipleRequest) GetThirdChannelCategory() map[string]interface{} {
	return s.ThirdChannelCategory
}

func (s *PushMultipleRequest) GetTransparentMessagePayload() interface{} {
	return s.TransparentMessagePayload
}

func (s *PushMultipleRequest) GetTransparentMessageUrgency() *string {
	return s.TransparentMessageUrgency
}

func (s *PushMultipleRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushMultipleRequest) SetActivityContentState(v interface{}) *PushMultipleRequest {
	s.ActivityContentState = v
	return s
}

func (s *PushMultipleRequest) SetActivityEvent(v string) *PushMultipleRequest {
	s.ActivityEvent = &v
	return s
}

func (s *PushMultipleRequest) SetAppId(v string) *PushMultipleRequest {
	s.AppId = &v
	return s
}

func (s *PushMultipleRequest) SetChannelId(v string) *PushMultipleRequest {
	s.ChannelId = &v
	return s
}

func (s *PushMultipleRequest) SetClassification(v string) *PushMultipleRequest {
	s.Classification = &v
	return s
}

func (s *PushMultipleRequest) SetDeliveryType(v int64) *PushMultipleRequest {
	s.DeliveryType = &v
	return s
}

func (s *PushMultipleRequest) SetDismissalDate(v int64) *PushMultipleRequest {
	s.DismissalDate = &v
	return s
}

func (s *PushMultipleRequest) SetExpiredSeconds(v int64) *PushMultipleRequest {
	s.ExpiredSeconds = &v
	return s
}

func (s *PushMultipleRequest) SetExtendedParams(v string) *PushMultipleRequest {
	s.ExtendedParams = &v
	return s
}

func (s *PushMultipleRequest) SetMiChannelId(v string) *PushMultipleRequest {
	s.MiChannelId = &v
	return s
}

func (s *PushMultipleRequest) SetNotifyType(v string) *PushMultipleRequest {
	s.NotifyType = &v
	return s
}

func (s *PushMultipleRequest) SetPushAction(v int64) *PushMultipleRequest {
	s.PushAction = &v
	return s
}

func (s *PushMultipleRequest) SetSilent(v int64) *PushMultipleRequest {
	s.Silent = &v
	return s
}

func (s *PushMultipleRequest) SetStrategyContent(v string) *PushMultipleRequest {
	s.StrategyContent = &v
	return s
}

func (s *PushMultipleRequest) SetStrategyType(v int32) *PushMultipleRequest {
	s.StrategyType = &v
	return s
}

func (s *PushMultipleRequest) SetTargetMsg(v []*PushMultipleRequestTargetMsg) *PushMultipleRequest {
	s.TargetMsg = v
	return s
}

func (s *PushMultipleRequest) SetTaskName(v string) *PushMultipleRequest {
	s.TaskName = &v
	return s
}

func (s *PushMultipleRequest) SetTemplateName(v string) *PushMultipleRequest {
	s.TemplateName = &v
	return s
}

func (s *PushMultipleRequest) SetTenantId(v string) *PushMultipleRequest {
	s.TenantId = &v
	return s
}

func (s *PushMultipleRequest) SetThirdChannelCategory(v map[string]interface{}) *PushMultipleRequest {
	s.ThirdChannelCategory = v
	return s
}

func (s *PushMultipleRequest) SetTransparentMessagePayload(v interface{}) *PushMultipleRequest {
	s.TransparentMessagePayload = v
	return s
}

func (s *PushMultipleRequest) SetTransparentMessageUrgency(v string) *PushMultipleRequest {
	s.TransparentMessageUrgency = &v
	return s
}

func (s *PushMultipleRequest) SetWorkspaceId(v string) *PushMultipleRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushMultipleRequest) Validate() error {
	return dara.Validate(s)
}

type PushMultipleRequestTargetMsg struct {
	ExtendedParams *string `json:"ExtendedParams,omitempty" xml:"ExtendedParams,omitempty"`
	// This parameter is required.
	MsgKey *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	// This parameter is required.
	Target           *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TemplateKeyValue *string `json:"TemplateKeyValue,omitempty" xml:"TemplateKeyValue,omitempty"`
}

func (s PushMultipleRequestTargetMsg) String() string {
	return dara.Prettify(s)
}

func (s PushMultipleRequestTargetMsg) GoString() string {
	return s.String()
}

func (s *PushMultipleRequestTargetMsg) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushMultipleRequestTargetMsg) GetMsgKey() *string {
	return s.MsgKey
}

func (s *PushMultipleRequestTargetMsg) GetTarget() *string {
	return s.Target
}

func (s *PushMultipleRequestTargetMsg) GetTemplateKeyValue() *string {
	return s.TemplateKeyValue
}

func (s *PushMultipleRequestTargetMsg) SetExtendedParams(v string) *PushMultipleRequestTargetMsg {
	s.ExtendedParams = &v
	return s
}

func (s *PushMultipleRequestTargetMsg) SetMsgKey(v string) *PushMultipleRequestTargetMsg {
	s.MsgKey = &v
	return s
}

func (s *PushMultipleRequestTargetMsg) SetTarget(v string) *PushMultipleRequestTargetMsg {
	s.Target = &v
	return s
}

func (s *PushMultipleRequestTargetMsg) SetTemplateKeyValue(v string) *PushMultipleRequestTargetMsg {
	s.TemplateKeyValue = &v
	return s
}

func (s *PushMultipleRequestTargetMsg) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMultipleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityContentState(v interface{}) *PushMultipleShrinkRequest
	GetActivityContentState() interface{}
	SetActivityEvent(v string) *PushMultipleShrinkRequest
	GetActivityEvent() *string
	SetAppId(v string) *PushMultipleShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *PushMultipleShrinkRequest
	GetChannelId() *string
	SetClassification(v string) *PushMultipleShrinkRequest
	GetClassification() *string
	SetDeliveryType(v int64) *PushMultipleShrinkRequest
	GetDeliveryType() *int64
	SetDismissalDate(v int64) *PushMultipleShrinkRequest
	GetDismissalDate() *int64
	SetExpiredSeconds(v int64) *PushMultipleShrinkRequest
	GetExpiredSeconds() *int64
	SetExtendedParams(v string) *PushMultipleShrinkRequest
	GetExtendedParams() *string
	SetMiChannelId(v string) *PushMultipleShrinkRequest
	GetMiChannelId() *string
	SetNotifyLevelShrink(v string) *PushMultipleShrinkRequest
	GetNotifyLevelShrink() *string
	SetNotifyType(v string) *PushMultipleShrinkRequest
	GetNotifyType() *string
	SetPushAction(v int64) *PushMultipleShrinkRequest
	GetPushAction() *int64
	SetSilent(v int64) *PushMultipleShrinkRequest
	GetSilent() *int64
	SetStrategyContent(v string) *PushMultipleShrinkRequest
	GetStrategyContent() *string
	SetStrategyType(v int32) *PushMultipleShrinkRequest
	GetStrategyType() *int32
	SetTargetMsg(v []*PushMultipleShrinkRequestTargetMsg) *PushMultipleShrinkRequest
	GetTargetMsg() []*PushMultipleShrinkRequestTargetMsg
	SetTaskName(v string) *PushMultipleShrinkRequest
	GetTaskName() *string
	SetTemplateName(v string) *PushMultipleShrinkRequest
	GetTemplateName() *string
	SetTenantId(v string) *PushMultipleShrinkRequest
	GetTenantId() *string
	SetThirdChannelCategoryShrink(v string) *PushMultipleShrinkRequest
	GetThirdChannelCategoryShrink() *string
	SetTransparentMessagePayload(v interface{}) *PushMultipleShrinkRequest
	GetTransparentMessagePayload() interface{}
	SetTransparentMessageUrgency(v string) *PushMultipleShrinkRequest
	GetTransparentMessageUrgency() *string
	SetWorkspaceId(v string) *PushMultipleShrinkRequest
	GetWorkspaceId() *string
}

type PushMultipleShrinkRequest struct {
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
	ExpiredSeconds    *int64  `json:"ExpiredSeconds,omitempty" xml:"ExpiredSeconds,omitempty"`
	ExtendedParams    *string `json:"ExtendedParams,omitempty" xml:"ExtendedParams,omitempty"`
	MiChannelId       *string `json:"MiChannelId,omitempty" xml:"MiChannelId,omitempty"`
	NotifyLevelShrink *string `json:"NotifyLevel,omitempty" xml:"NotifyLevel,omitempty"`
	NotifyType        *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	PushAction        *int64  `json:"PushAction,omitempty" xml:"PushAction,omitempty"`
	Silent            *int64  `json:"Silent,omitempty" xml:"Silent,omitempty"`
	StrategyContent   *string `json:"StrategyContent,omitempty" xml:"StrategyContent,omitempty"`
	StrategyType      *int32  `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// This parameter is required.
	TargetMsg []*PushMultipleShrinkRequestTargetMsg `json:"TargetMsg,omitempty" xml:"TargetMsg,omitempty" type:"Repeated"`
	TaskName  *string                               `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// This parameter is required.
	TemplateName               *string     `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantId                   *string     `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdChannelCategoryShrink *string     `json:"ThirdChannelCategory,omitempty" xml:"ThirdChannelCategory,omitempty"`
	TransparentMessagePayload  interface{} `json:"TransparentMessagePayload,omitempty" xml:"TransparentMessagePayload,omitempty"`
	TransparentMessageUrgency  *string     `json:"TransparentMessageUrgency,omitempty" xml:"TransparentMessageUrgency,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushMultipleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushMultipleShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushMultipleShrinkRequest) GetActivityContentState() interface{} {
	return s.ActivityContentState
}

func (s *PushMultipleShrinkRequest) GetActivityEvent() *string {
	return s.ActivityEvent
}

func (s *PushMultipleShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushMultipleShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *PushMultipleShrinkRequest) GetClassification() *string {
	return s.Classification
}

func (s *PushMultipleShrinkRequest) GetDeliveryType() *int64 {
	return s.DeliveryType
}

func (s *PushMultipleShrinkRequest) GetDismissalDate() *int64 {
	return s.DismissalDate
}

func (s *PushMultipleShrinkRequest) GetExpiredSeconds() *int64 {
	return s.ExpiredSeconds
}

func (s *PushMultipleShrinkRequest) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushMultipleShrinkRequest) GetMiChannelId() *string {
	return s.MiChannelId
}

func (s *PushMultipleShrinkRequest) GetNotifyLevelShrink() *string {
	return s.NotifyLevelShrink
}

func (s *PushMultipleShrinkRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *PushMultipleShrinkRequest) GetPushAction() *int64 {
	return s.PushAction
}

func (s *PushMultipleShrinkRequest) GetSilent() *int64 {
	return s.Silent
}

func (s *PushMultipleShrinkRequest) GetStrategyContent() *string {
	return s.StrategyContent
}

func (s *PushMultipleShrinkRequest) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *PushMultipleShrinkRequest) GetTargetMsg() []*PushMultipleShrinkRequestTargetMsg {
	return s.TargetMsg
}

func (s *PushMultipleShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *PushMultipleShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PushMultipleShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushMultipleShrinkRequest) GetThirdChannelCategoryShrink() *string {
	return s.ThirdChannelCategoryShrink
}

func (s *PushMultipleShrinkRequest) GetTransparentMessagePayload() interface{} {
	return s.TransparentMessagePayload
}

func (s *PushMultipleShrinkRequest) GetTransparentMessageUrgency() *string {
	return s.TransparentMessageUrgency
}

func (s *PushMultipleShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushMultipleShrinkRequest) SetActivityContentState(v interface{}) *PushMultipleShrinkRequest {
	s.ActivityContentState = v
	return s
}

func (s *PushMultipleShrinkRequest) SetActivityEvent(v string) *PushMultipleShrinkRequest {
	s.ActivityEvent = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetAppId(v string) *PushMultipleShrinkRequest {
	s.AppId = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetChannelId(v string) *PushMultipleShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetClassification(v string) *PushMultipleShrinkRequest {
	s.Classification = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetDeliveryType(v int64) *PushMultipleShrinkRequest {
	s.DeliveryType = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetDismissalDate(v int64) *PushMultipleShrinkRequest {
	s.DismissalDate = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetExpiredSeconds(v int64) *PushMultipleShrinkRequest {
	s.ExpiredSeconds = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetExtendedParams(v string) *PushMultipleShrinkRequest {
	s.ExtendedParams = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetMiChannelId(v string) *PushMultipleShrinkRequest {
	s.MiChannelId = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetNotifyLevelShrink(v string) *PushMultipleShrinkRequest {
	s.NotifyLevelShrink = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetNotifyType(v string) *PushMultipleShrinkRequest {
	s.NotifyType = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetPushAction(v int64) *PushMultipleShrinkRequest {
	s.PushAction = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetSilent(v int64) *PushMultipleShrinkRequest {
	s.Silent = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetStrategyContent(v string) *PushMultipleShrinkRequest {
	s.StrategyContent = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetStrategyType(v int32) *PushMultipleShrinkRequest {
	s.StrategyType = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetTargetMsg(v []*PushMultipleShrinkRequestTargetMsg) *PushMultipleShrinkRequest {
	s.TargetMsg = v
	return s
}

func (s *PushMultipleShrinkRequest) SetTaskName(v string) *PushMultipleShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetTemplateName(v string) *PushMultipleShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetTenantId(v string) *PushMultipleShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetThirdChannelCategoryShrink(v string) *PushMultipleShrinkRequest {
	s.ThirdChannelCategoryShrink = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetTransparentMessagePayload(v interface{}) *PushMultipleShrinkRequest {
	s.TransparentMessagePayload = v
	return s
}

func (s *PushMultipleShrinkRequest) SetTransparentMessageUrgency(v string) *PushMultipleShrinkRequest {
	s.TransparentMessageUrgency = &v
	return s
}

func (s *PushMultipleShrinkRequest) SetWorkspaceId(v string) *PushMultipleShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushMultipleShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type PushMultipleShrinkRequestTargetMsg struct {
	ExtendedParams *string `json:"ExtendedParams,omitempty" xml:"ExtendedParams,omitempty"`
	// This parameter is required.
	MsgKey *string `json:"MsgKey,omitempty" xml:"MsgKey,omitempty"`
	// This parameter is required.
	Target           *string `json:"Target,omitempty" xml:"Target,omitempty"`
	TemplateKeyValue *string `json:"TemplateKeyValue,omitempty" xml:"TemplateKeyValue,omitempty"`
}

func (s PushMultipleShrinkRequestTargetMsg) String() string {
	return dara.Prettify(s)
}

func (s PushMultipleShrinkRequestTargetMsg) GoString() string {
	return s.String()
}

func (s *PushMultipleShrinkRequestTargetMsg) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushMultipleShrinkRequestTargetMsg) GetMsgKey() *string {
	return s.MsgKey
}

func (s *PushMultipleShrinkRequestTargetMsg) GetTarget() *string {
	return s.Target
}

func (s *PushMultipleShrinkRequestTargetMsg) GetTemplateKeyValue() *string {
	return s.TemplateKeyValue
}

func (s *PushMultipleShrinkRequestTargetMsg) SetExtendedParams(v string) *PushMultipleShrinkRequestTargetMsg {
	s.ExtendedParams = &v
	return s
}

func (s *PushMultipleShrinkRequestTargetMsg) SetMsgKey(v string) *PushMultipleShrinkRequestTargetMsg {
	s.MsgKey = &v
	return s
}

func (s *PushMultipleShrinkRequestTargetMsg) SetTarget(v string) *PushMultipleShrinkRequestTargetMsg {
	s.Target = &v
	return s
}

func (s *PushMultipleShrinkRequestTargetMsg) SetTemplateKeyValue(v string) *PushMultipleShrinkRequestTargetMsg {
	s.TemplateKeyValue = &v
	return s
}

func (s *PushMultipleShrinkRequestTargetMsg) Validate() error {
	return dara.Validate(s)
}

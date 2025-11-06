// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityContentState(v interface{}) *PushTemplateShrinkRequest
	GetActivityContentState() interface{}
	SetActivityEvent(v string) *PushTemplateShrinkRequest
	GetActivityEvent() *string
	SetAppId(v string) *PushTemplateShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *PushTemplateShrinkRequest
	GetChannelId() *string
	SetClassification(v string) *PushTemplateShrinkRequest
	GetClassification() *string
	SetDeliveryType(v int64) *PushTemplateShrinkRequest
	GetDeliveryType() *int64
	SetDismissalDate(v int64) *PushTemplateShrinkRequest
	GetDismissalDate() *int64
	SetExpiredSeconds(v int64) *PushTemplateShrinkRequest
	GetExpiredSeconds() *int64
	SetExtendedParams(v string) *PushTemplateShrinkRequest
	GetExtendedParams() *string
	SetMiChannelId(v string) *PushTemplateShrinkRequest
	GetMiChannelId() *string
	SetNotifyLevelShrink(v string) *PushTemplateShrinkRequest
	GetNotifyLevelShrink() *string
	SetNotifyType(v string) *PushTemplateShrinkRequest
	GetNotifyType() *string
	SetPushAction(v int64) *PushTemplateShrinkRequest
	GetPushAction() *int64
	SetSilent(v int64) *PushTemplateShrinkRequest
	GetSilent() *int64
	SetSmsSignName(v string) *PushTemplateShrinkRequest
	GetSmsSignName() *string
	SetSmsStrategy(v int32) *PushTemplateShrinkRequest
	GetSmsStrategy() *int32
	SetSmsTemplateCode(v string) *PushTemplateShrinkRequest
	GetSmsTemplateCode() *string
	SetSmsTemplateParam(v string) *PushTemplateShrinkRequest
	GetSmsTemplateParam() *string
	SetStrategyContent(v string) *PushTemplateShrinkRequest
	GetStrategyContent() *string
	SetStrategyType(v int32) *PushTemplateShrinkRequest
	GetStrategyType() *int32
	SetTargetMsgkey(v string) *PushTemplateShrinkRequest
	GetTargetMsgkey() *string
	SetTaskName(v string) *PushTemplateShrinkRequest
	GetTaskName() *string
	SetTemplateKeyValue(v string) *PushTemplateShrinkRequest
	GetTemplateKeyValue() *string
	SetTemplateName(v string) *PushTemplateShrinkRequest
	GetTemplateName() *string
	SetTenantId(v string) *PushTemplateShrinkRequest
	GetTenantId() *string
	SetThirdChannelCategoryShrink(v string) *PushTemplateShrinkRequest
	GetThirdChannelCategoryShrink() *string
	SetTransparentMessagePayload(v interface{}) *PushTemplateShrinkRequest
	GetTransparentMessagePayload() interface{}
	SetTransparentMessageUrgency(v string) *PushTemplateShrinkRequest
	GetTransparentMessageUrgency() *string
	SetWorkspaceId(v string) *PushTemplateShrinkRequest
	GetWorkspaceId() *string
}

type PushTemplateShrinkRequest struct {
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
	SmsSignName       *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	SmsStrategy       *int32  `json:"SmsStrategy,omitempty" xml:"SmsStrategy,omitempty"`
	SmsTemplateCode   *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	SmsTemplateParam  *string `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	StrategyContent   *string `json:"StrategyContent,omitempty" xml:"StrategyContent,omitempty"`
	StrategyType      *int32  `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// This parameter is required.
	TargetMsgkey     *string `json:"TargetMsgkey,omitempty" xml:"TargetMsgkey,omitempty"`
	TaskName         *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateKeyValue *string `json:"TemplateKeyValue,omitempty" xml:"TemplateKeyValue,omitempty"`
	// This parameter is required.
	TemplateName               *string     `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantId                   *string     `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdChannelCategoryShrink *string     `json:"ThirdChannelCategory,omitempty" xml:"ThirdChannelCategory,omitempty"`
	TransparentMessagePayload  interface{} `json:"TransparentMessagePayload,omitempty" xml:"TransparentMessagePayload,omitempty"`
	TransparentMessageUrgency  *string     `json:"TransparentMessageUrgency,omitempty" xml:"TransparentMessageUrgency,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushTemplateShrinkRequest) GetActivityContentState() interface{} {
	return s.ActivityContentState
}

func (s *PushTemplateShrinkRequest) GetActivityEvent() *string {
	return s.ActivityEvent
}

func (s *PushTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushTemplateShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *PushTemplateShrinkRequest) GetClassification() *string {
	return s.Classification
}

func (s *PushTemplateShrinkRequest) GetDeliveryType() *int64 {
	return s.DeliveryType
}

func (s *PushTemplateShrinkRequest) GetDismissalDate() *int64 {
	return s.DismissalDate
}

func (s *PushTemplateShrinkRequest) GetExpiredSeconds() *int64 {
	return s.ExpiredSeconds
}

func (s *PushTemplateShrinkRequest) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushTemplateShrinkRequest) GetMiChannelId() *string {
	return s.MiChannelId
}

func (s *PushTemplateShrinkRequest) GetNotifyLevelShrink() *string {
	return s.NotifyLevelShrink
}

func (s *PushTemplateShrinkRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *PushTemplateShrinkRequest) GetPushAction() *int64 {
	return s.PushAction
}

func (s *PushTemplateShrinkRequest) GetSilent() *int64 {
	return s.Silent
}

func (s *PushTemplateShrinkRequest) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *PushTemplateShrinkRequest) GetSmsStrategy() *int32 {
	return s.SmsStrategy
}

func (s *PushTemplateShrinkRequest) GetSmsTemplateCode() *string {
	return s.SmsTemplateCode
}

func (s *PushTemplateShrinkRequest) GetSmsTemplateParam() *string {
	return s.SmsTemplateParam
}

func (s *PushTemplateShrinkRequest) GetStrategyContent() *string {
	return s.StrategyContent
}

func (s *PushTemplateShrinkRequest) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *PushTemplateShrinkRequest) GetTargetMsgkey() *string {
	return s.TargetMsgkey
}

func (s *PushTemplateShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *PushTemplateShrinkRequest) GetTemplateKeyValue() *string {
	return s.TemplateKeyValue
}

func (s *PushTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PushTemplateShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushTemplateShrinkRequest) GetThirdChannelCategoryShrink() *string {
	return s.ThirdChannelCategoryShrink
}

func (s *PushTemplateShrinkRequest) GetTransparentMessagePayload() interface{} {
	return s.TransparentMessagePayload
}

func (s *PushTemplateShrinkRequest) GetTransparentMessageUrgency() *string {
	return s.TransparentMessageUrgency
}

func (s *PushTemplateShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushTemplateShrinkRequest) SetActivityContentState(v interface{}) *PushTemplateShrinkRequest {
	s.ActivityContentState = v
	return s
}

func (s *PushTemplateShrinkRequest) SetActivityEvent(v string) *PushTemplateShrinkRequest {
	s.ActivityEvent = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetAppId(v string) *PushTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetChannelId(v string) *PushTemplateShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetClassification(v string) *PushTemplateShrinkRequest {
	s.Classification = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetDeliveryType(v int64) *PushTemplateShrinkRequest {
	s.DeliveryType = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetDismissalDate(v int64) *PushTemplateShrinkRequest {
	s.DismissalDate = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetExpiredSeconds(v int64) *PushTemplateShrinkRequest {
	s.ExpiredSeconds = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetExtendedParams(v string) *PushTemplateShrinkRequest {
	s.ExtendedParams = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetMiChannelId(v string) *PushTemplateShrinkRequest {
	s.MiChannelId = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetNotifyLevelShrink(v string) *PushTemplateShrinkRequest {
	s.NotifyLevelShrink = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetNotifyType(v string) *PushTemplateShrinkRequest {
	s.NotifyType = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetPushAction(v int64) *PushTemplateShrinkRequest {
	s.PushAction = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetSilent(v int64) *PushTemplateShrinkRequest {
	s.Silent = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetSmsSignName(v string) *PushTemplateShrinkRequest {
	s.SmsSignName = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetSmsStrategy(v int32) *PushTemplateShrinkRequest {
	s.SmsStrategy = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetSmsTemplateCode(v string) *PushTemplateShrinkRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetSmsTemplateParam(v string) *PushTemplateShrinkRequest {
	s.SmsTemplateParam = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetStrategyContent(v string) *PushTemplateShrinkRequest {
	s.StrategyContent = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetStrategyType(v int32) *PushTemplateShrinkRequest {
	s.StrategyType = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetTargetMsgkey(v string) *PushTemplateShrinkRequest {
	s.TargetMsgkey = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetTaskName(v string) *PushTemplateShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetTemplateKeyValue(v string) *PushTemplateShrinkRequest {
	s.TemplateKeyValue = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetTemplateName(v string) *PushTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetTenantId(v string) *PushTemplateShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetThirdChannelCategoryShrink(v string) *PushTemplateShrinkRequest {
	s.ThirdChannelCategoryShrink = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetTransparentMessagePayload(v interface{}) *PushTemplateShrinkRequest {
	s.TransparentMessagePayload = v
	return s
}

func (s *PushTemplateShrinkRequest) SetTransparentMessageUrgency(v string) *PushTemplateShrinkRequest {
	s.TransparentMessageUrgency = &v
	return s
}

func (s *PushTemplateShrinkRequest) SetWorkspaceId(v string) *PushTemplateShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}

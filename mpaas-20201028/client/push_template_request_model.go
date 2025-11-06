// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityContentState(v interface{}) *PushTemplateRequest
	GetActivityContentState() interface{}
	SetActivityEvent(v string) *PushTemplateRequest
	GetActivityEvent() *string
	SetAppId(v string) *PushTemplateRequest
	GetAppId() *string
	SetChannelId(v string) *PushTemplateRequest
	GetChannelId() *string
	SetClassification(v string) *PushTemplateRequest
	GetClassification() *string
	SetDeliveryType(v int64) *PushTemplateRequest
	GetDeliveryType() *int64
	SetDismissalDate(v int64) *PushTemplateRequest
	GetDismissalDate() *int64
	SetExpiredSeconds(v int64) *PushTemplateRequest
	GetExpiredSeconds() *int64
	SetExtendedParams(v string) *PushTemplateRequest
	GetExtendedParams() *string
	SetMiChannelId(v string) *PushTemplateRequest
	GetMiChannelId() *string
	SetNotifyLevel(v map[string]interface{}) *PushTemplateRequest
	GetNotifyLevel() map[string]interface{}
	SetNotifyType(v string) *PushTemplateRequest
	GetNotifyType() *string
	SetPushAction(v int64) *PushTemplateRequest
	GetPushAction() *int64
	SetSilent(v int64) *PushTemplateRequest
	GetSilent() *int64
	SetSmsSignName(v string) *PushTemplateRequest
	GetSmsSignName() *string
	SetSmsStrategy(v int32) *PushTemplateRequest
	GetSmsStrategy() *int32
	SetSmsTemplateCode(v string) *PushTemplateRequest
	GetSmsTemplateCode() *string
	SetSmsTemplateParam(v string) *PushTemplateRequest
	GetSmsTemplateParam() *string
	SetStrategyContent(v string) *PushTemplateRequest
	GetStrategyContent() *string
	SetStrategyType(v int32) *PushTemplateRequest
	GetStrategyType() *int32
	SetTargetMsgkey(v string) *PushTemplateRequest
	GetTargetMsgkey() *string
	SetTaskName(v string) *PushTemplateRequest
	GetTaskName() *string
	SetTemplateKeyValue(v string) *PushTemplateRequest
	GetTemplateKeyValue() *string
	SetTemplateName(v string) *PushTemplateRequest
	GetTemplateName() *string
	SetTenantId(v string) *PushTemplateRequest
	GetTenantId() *string
	SetThirdChannelCategory(v map[string]interface{}) *PushTemplateRequest
	GetThirdChannelCategory() map[string]interface{}
	SetTransparentMessagePayload(v interface{}) *PushTemplateRequest
	GetTransparentMessagePayload() interface{}
	SetTransparentMessageUrgency(v string) *PushTemplateRequest
	GetTransparentMessageUrgency() *string
	SetWorkspaceId(v string) *PushTemplateRequest
	GetWorkspaceId() *string
}

type PushTemplateRequest struct {
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
	ExpiredSeconds   *int64                 `json:"ExpiredSeconds,omitempty" xml:"ExpiredSeconds,omitempty"`
	ExtendedParams   *string                `json:"ExtendedParams,omitempty" xml:"ExtendedParams,omitempty"`
	MiChannelId      *string                `json:"MiChannelId,omitempty" xml:"MiChannelId,omitempty"`
	NotifyLevel      map[string]interface{} `json:"NotifyLevel,omitempty" xml:"NotifyLevel,omitempty"`
	NotifyType       *string                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	PushAction       *int64                 `json:"PushAction,omitempty" xml:"PushAction,omitempty"`
	Silent           *int64                 `json:"Silent,omitempty" xml:"Silent,omitempty"`
	SmsSignName      *string                `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	SmsStrategy      *int32                 `json:"SmsStrategy,omitempty" xml:"SmsStrategy,omitempty"`
	SmsTemplateCode  *string                `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	SmsTemplateParam *string                `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	StrategyContent  *string                `json:"StrategyContent,omitempty" xml:"StrategyContent,omitempty"`
	StrategyType     *int32                 `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// This parameter is required.
	TargetMsgkey     *string `json:"TargetMsgkey,omitempty" xml:"TargetMsgkey,omitempty"`
	TaskName         *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateKeyValue *string `json:"TemplateKeyValue,omitempty" xml:"TemplateKeyValue,omitempty"`
	// This parameter is required.
	TemplateName              *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantId                  *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdChannelCategory      map[string]interface{} `json:"ThirdChannelCategory,omitempty" xml:"ThirdChannelCategory,omitempty"`
	TransparentMessagePayload interface{}            `json:"TransparentMessagePayload,omitempty" xml:"TransparentMessagePayload,omitempty"`
	TransparentMessageUrgency *string                `json:"TransparentMessageUrgency,omitempty" xml:"TransparentMessageUrgency,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s PushTemplateRequest) GoString() string {
	return s.String()
}

func (s *PushTemplateRequest) GetActivityContentState() interface{} {
	return s.ActivityContentState
}

func (s *PushTemplateRequest) GetActivityEvent() *string {
	return s.ActivityEvent
}

func (s *PushTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushTemplateRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *PushTemplateRequest) GetClassification() *string {
	return s.Classification
}

func (s *PushTemplateRequest) GetDeliveryType() *int64 {
	return s.DeliveryType
}

func (s *PushTemplateRequest) GetDismissalDate() *int64 {
	return s.DismissalDate
}

func (s *PushTemplateRequest) GetExpiredSeconds() *int64 {
	return s.ExpiredSeconds
}

func (s *PushTemplateRequest) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushTemplateRequest) GetMiChannelId() *string {
	return s.MiChannelId
}

func (s *PushTemplateRequest) GetNotifyLevel() map[string]interface{} {
	return s.NotifyLevel
}

func (s *PushTemplateRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *PushTemplateRequest) GetPushAction() *int64 {
	return s.PushAction
}

func (s *PushTemplateRequest) GetSilent() *int64 {
	return s.Silent
}

func (s *PushTemplateRequest) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *PushTemplateRequest) GetSmsStrategy() *int32 {
	return s.SmsStrategy
}

func (s *PushTemplateRequest) GetSmsTemplateCode() *string {
	return s.SmsTemplateCode
}

func (s *PushTemplateRequest) GetSmsTemplateParam() *string {
	return s.SmsTemplateParam
}

func (s *PushTemplateRequest) GetStrategyContent() *string {
	return s.StrategyContent
}

func (s *PushTemplateRequest) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *PushTemplateRequest) GetTargetMsgkey() *string {
	return s.TargetMsgkey
}

func (s *PushTemplateRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *PushTemplateRequest) GetTemplateKeyValue() *string {
	return s.TemplateKeyValue
}

func (s *PushTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PushTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushTemplateRequest) GetThirdChannelCategory() map[string]interface{} {
	return s.ThirdChannelCategory
}

func (s *PushTemplateRequest) GetTransparentMessagePayload() interface{} {
	return s.TransparentMessagePayload
}

func (s *PushTemplateRequest) GetTransparentMessageUrgency() *string {
	return s.TransparentMessageUrgency
}

func (s *PushTemplateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushTemplateRequest) SetActivityContentState(v interface{}) *PushTemplateRequest {
	s.ActivityContentState = v
	return s
}

func (s *PushTemplateRequest) SetActivityEvent(v string) *PushTemplateRequest {
	s.ActivityEvent = &v
	return s
}

func (s *PushTemplateRequest) SetAppId(v string) *PushTemplateRequest {
	s.AppId = &v
	return s
}

func (s *PushTemplateRequest) SetChannelId(v string) *PushTemplateRequest {
	s.ChannelId = &v
	return s
}

func (s *PushTemplateRequest) SetClassification(v string) *PushTemplateRequest {
	s.Classification = &v
	return s
}

func (s *PushTemplateRequest) SetDeliveryType(v int64) *PushTemplateRequest {
	s.DeliveryType = &v
	return s
}

func (s *PushTemplateRequest) SetDismissalDate(v int64) *PushTemplateRequest {
	s.DismissalDate = &v
	return s
}

func (s *PushTemplateRequest) SetExpiredSeconds(v int64) *PushTemplateRequest {
	s.ExpiredSeconds = &v
	return s
}

func (s *PushTemplateRequest) SetExtendedParams(v string) *PushTemplateRequest {
	s.ExtendedParams = &v
	return s
}

func (s *PushTemplateRequest) SetMiChannelId(v string) *PushTemplateRequest {
	s.MiChannelId = &v
	return s
}

func (s *PushTemplateRequest) SetNotifyLevel(v map[string]interface{}) *PushTemplateRequest {
	s.NotifyLevel = v
	return s
}

func (s *PushTemplateRequest) SetNotifyType(v string) *PushTemplateRequest {
	s.NotifyType = &v
	return s
}

func (s *PushTemplateRequest) SetPushAction(v int64) *PushTemplateRequest {
	s.PushAction = &v
	return s
}

func (s *PushTemplateRequest) SetSilent(v int64) *PushTemplateRequest {
	s.Silent = &v
	return s
}

func (s *PushTemplateRequest) SetSmsSignName(v string) *PushTemplateRequest {
	s.SmsSignName = &v
	return s
}

func (s *PushTemplateRequest) SetSmsStrategy(v int32) *PushTemplateRequest {
	s.SmsStrategy = &v
	return s
}

func (s *PushTemplateRequest) SetSmsTemplateCode(v string) *PushTemplateRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *PushTemplateRequest) SetSmsTemplateParam(v string) *PushTemplateRequest {
	s.SmsTemplateParam = &v
	return s
}

func (s *PushTemplateRequest) SetStrategyContent(v string) *PushTemplateRequest {
	s.StrategyContent = &v
	return s
}

func (s *PushTemplateRequest) SetStrategyType(v int32) *PushTemplateRequest {
	s.StrategyType = &v
	return s
}

func (s *PushTemplateRequest) SetTargetMsgkey(v string) *PushTemplateRequest {
	s.TargetMsgkey = &v
	return s
}

func (s *PushTemplateRequest) SetTaskName(v string) *PushTemplateRequest {
	s.TaskName = &v
	return s
}

func (s *PushTemplateRequest) SetTemplateKeyValue(v string) *PushTemplateRequest {
	s.TemplateKeyValue = &v
	return s
}

func (s *PushTemplateRequest) SetTemplateName(v string) *PushTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *PushTemplateRequest) SetTenantId(v string) *PushTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *PushTemplateRequest) SetThirdChannelCategory(v map[string]interface{}) *PushTemplateRequest {
	s.ThirdChannelCategory = v
	return s
}

func (s *PushTemplateRequest) SetTransparentMessagePayload(v interface{}) *PushTemplateRequest {
	s.TransparentMessagePayload = v
	return s
}

func (s *PushTemplateRequest) SetTransparentMessageUrgency(v string) *PushTemplateRequest {
	s.TransparentMessageUrgency = &v
	return s
}

func (s *PushTemplateRequest) SetWorkspaceId(v string) *PushTemplateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushTemplateRequest) Validate() error {
	return dara.Validate(s)
}

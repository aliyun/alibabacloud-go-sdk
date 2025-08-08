// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushBroadcastShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidChannel(v int32) *PushBroadcastShrinkRequest
	GetAndroidChannel() *int32
	SetAppId(v string) *PushBroadcastShrinkRequest
	GetAppId() *string
	SetBindPeriod(v int32) *PushBroadcastShrinkRequest
	GetBindPeriod() *int32
	SetChannelId(v string) *PushBroadcastShrinkRequest
	GetChannelId() *string
	SetClassification(v string) *PushBroadcastShrinkRequest
	GetClassification() *string
	SetDeliveryType(v int64) *PushBroadcastShrinkRequest
	GetDeliveryType() *int64
	SetExpiredSeconds(v int64) *PushBroadcastShrinkRequest
	GetExpiredSeconds() *int64
	SetExtendedParams(v string) *PushBroadcastShrinkRequest
	GetExtendedParams() *string
	SetMiChannelId(v string) *PushBroadcastShrinkRequest
	GetMiChannelId() *string
	SetMsgkey(v string) *PushBroadcastShrinkRequest
	GetMsgkey() *string
	SetNotifyType(v string) *PushBroadcastShrinkRequest
	GetNotifyType() *string
	SetPushAction(v int64) *PushBroadcastShrinkRequest
	GetPushAction() *int64
	SetPushStatus(v int64) *PushBroadcastShrinkRequest
	GetPushStatus() *int64
	SetSilent(v int64) *PushBroadcastShrinkRequest
	GetSilent() *int64
	SetStrategyContent(v string) *PushBroadcastShrinkRequest
	GetStrategyContent() *string
	SetStrategyType(v int32) *PushBroadcastShrinkRequest
	GetStrategyType() *int32
	SetTaskName(v string) *PushBroadcastShrinkRequest
	GetTaskName() *string
	SetTemplateKeyValue(v string) *PushBroadcastShrinkRequest
	GetTemplateKeyValue() *string
	SetTemplateName(v string) *PushBroadcastShrinkRequest
	GetTemplateName() *string
	SetTenantId(v string) *PushBroadcastShrinkRequest
	GetTenantId() *string
	SetThirdChannelCategoryShrink(v string) *PushBroadcastShrinkRequest
	GetThirdChannelCategoryShrink() *string
	SetTransparentMessagePayload(v interface{}) *PushBroadcastShrinkRequest
	GetTransparentMessagePayload() interface{}
	SetTransparentMessageUrgency(v string) *PushBroadcastShrinkRequest
	GetTransparentMessageUrgency() *string
	SetUnBindPeriod(v int64) *PushBroadcastShrinkRequest
	GetUnBindPeriod() *int64
	SetWorkspaceId(v string) *PushBroadcastShrinkRequest
	GetWorkspaceId() *string
}

type PushBroadcastShrinkRequest struct {
	AndroidChannel *int32 `json:"AndroidChannel,omitempty" xml:"AndroidChannel,omitempty"`
	// This parameter is required.
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BindPeriod     *int32  `json:"BindPeriod,omitempty" xml:"BindPeriod,omitempty"`
	ChannelId      *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// This parameter is required.
	DeliveryType *int64 `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	// This parameter is required.
	ExpiredSeconds *int64  `json:"ExpiredSeconds,omitempty" xml:"ExpiredSeconds,omitempty"`
	ExtendedParams *string `json:"ExtendedParams,omitempty" xml:"ExtendedParams,omitempty"`
	MiChannelId    *string `json:"MiChannelId,omitempty" xml:"MiChannelId,omitempty"`
	// This parameter is required.
	Msgkey           *string `json:"Msgkey,omitempty" xml:"Msgkey,omitempty"`
	NotifyType       *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	PushAction       *int64  `json:"PushAction,omitempty" xml:"PushAction,omitempty"`
	PushStatus       *int64  `json:"PushStatus,omitempty" xml:"PushStatus,omitempty"`
	Silent           *int64  `json:"Silent,omitempty" xml:"Silent,omitempty"`
	StrategyContent  *string `json:"StrategyContent,omitempty" xml:"StrategyContent,omitempty"`
	StrategyType     *int32  `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	TaskName         *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateKeyValue *string `json:"TemplateKeyValue,omitempty" xml:"TemplateKeyValue,omitempty"`
	// This parameter is required.
	TemplateName               *string     `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantId                   *string     `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdChannelCategoryShrink *string     `json:"ThirdChannelCategory,omitempty" xml:"ThirdChannelCategory,omitempty"`
	TransparentMessagePayload  interface{} `json:"TransparentMessagePayload,omitempty" xml:"TransparentMessagePayload,omitempty"`
	TransparentMessageUrgency  *string     `json:"TransparentMessageUrgency,omitempty" xml:"TransparentMessageUrgency,omitempty"`
	UnBindPeriod               *int64      `json:"UnBindPeriod,omitempty" xml:"UnBindPeriod,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushBroadcastShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushBroadcastShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushBroadcastShrinkRequest) GetAndroidChannel() *int32 {
	return s.AndroidChannel
}

func (s *PushBroadcastShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushBroadcastShrinkRequest) GetBindPeriod() *int32 {
	return s.BindPeriod
}

func (s *PushBroadcastShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *PushBroadcastShrinkRequest) GetClassification() *string {
	return s.Classification
}

func (s *PushBroadcastShrinkRequest) GetDeliveryType() *int64 {
	return s.DeliveryType
}

func (s *PushBroadcastShrinkRequest) GetExpiredSeconds() *int64 {
	return s.ExpiredSeconds
}

func (s *PushBroadcastShrinkRequest) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushBroadcastShrinkRequest) GetMiChannelId() *string {
	return s.MiChannelId
}

func (s *PushBroadcastShrinkRequest) GetMsgkey() *string {
	return s.Msgkey
}

func (s *PushBroadcastShrinkRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *PushBroadcastShrinkRequest) GetPushAction() *int64 {
	return s.PushAction
}

func (s *PushBroadcastShrinkRequest) GetPushStatus() *int64 {
	return s.PushStatus
}

func (s *PushBroadcastShrinkRequest) GetSilent() *int64 {
	return s.Silent
}

func (s *PushBroadcastShrinkRequest) GetStrategyContent() *string {
	return s.StrategyContent
}

func (s *PushBroadcastShrinkRequest) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *PushBroadcastShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *PushBroadcastShrinkRequest) GetTemplateKeyValue() *string {
	return s.TemplateKeyValue
}

func (s *PushBroadcastShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PushBroadcastShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushBroadcastShrinkRequest) GetThirdChannelCategoryShrink() *string {
	return s.ThirdChannelCategoryShrink
}

func (s *PushBroadcastShrinkRequest) GetTransparentMessagePayload() interface{} {
	return s.TransparentMessagePayload
}

func (s *PushBroadcastShrinkRequest) GetTransparentMessageUrgency() *string {
	return s.TransparentMessageUrgency
}

func (s *PushBroadcastShrinkRequest) GetUnBindPeriod() *int64 {
	return s.UnBindPeriod
}

func (s *PushBroadcastShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushBroadcastShrinkRequest) SetAndroidChannel(v int32) *PushBroadcastShrinkRequest {
	s.AndroidChannel = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetAppId(v string) *PushBroadcastShrinkRequest {
	s.AppId = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetBindPeriod(v int32) *PushBroadcastShrinkRequest {
	s.BindPeriod = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetChannelId(v string) *PushBroadcastShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetClassification(v string) *PushBroadcastShrinkRequest {
	s.Classification = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetDeliveryType(v int64) *PushBroadcastShrinkRequest {
	s.DeliveryType = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetExpiredSeconds(v int64) *PushBroadcastShrinkRequest {
	s.ExpiredSeconds = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetExtendedParams(v string) *PushBroadcastShrinkRequest {
	s.ExtendedParams = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetMiChannelId(v string) *PushBroadcastShrinkRequest {
	s.MiChannelId = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetMsgkey(v string) *PushBroadcastShrinkRequest {
	s.Msgkey = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetNotifyType(v string) *PushBroadcastShrinkRequest {
	s.NotifyType = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetPushAction(v int64) *PushBroadcastShrinkRequest {
	s.PushAction = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetPushStatus(v int64) *PushBroadcastShrinkRequest {
	s.PushStatus = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetSilent(v int64) *PushBroadcastShrinkRequest {
	s.Silent = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetStrategyContent(v string) *PushBroadcastShrinkRequest {
	s.StrategyContent = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetStrategyType(v int32) *PushBroadcastShrinkRequest {
	s.StrategyType = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetTaskName(v string) *PushBroadcastShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetTemplateKeyValue(v string) *PushBroadcastShrinkRequest {
	s.TemplateKeyValue = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetTemplateName(v string) *PushBroadcastShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetTenantId(v string) *PushBroadcastShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetThirdChannelCategoryShrink(v string) *PushBroadcastShrinkRequest {
	s.ThirdChannelCategoryShrink = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetTransparentMessagePayload(v interface{}) *PushBroadcastShrinkRequest {
	s.TransparentMessagePayload = v
	return s
}

func (s *PushBroadcastShrinkRequest) SetTransparentMessageUrgency(v string) *PushBroadcastShrinkRequest {
	s.TransparentMessageUrgency = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetUnBindPeriod(v int64) *PushBroadcastShrinkRequest {
	s.UnBindPeriod = &v
	return s
}

func (s *PushBroadcastShrinkRequest) SetWorkspaceId(v string) *PushBroadcastShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushBroadcastShrinkRequest) Validate() error {
	return dara.Validate(s)
}

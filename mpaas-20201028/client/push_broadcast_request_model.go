// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushBroadcastRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidChannel(v int32) *PushBroadcastRequest
	GetAndroidChannel() *int32
	SetAppId(v string) *PushBroadcastRequest
	GetAppId() *string
	SetBindPeriod(v int32) *PushBroadcastRequest
	GetBindPeriod() *int32
	SetChannelId(v string) *PushBroadcastRequest
	GetChannelId() *string
	SetClassification(v string) *PushBroadcastRequest
	GetClassification() *string
	SetDeliveryType(v int64) *PushBroadcastRequest
	GetDeliveryType() *int64
	SetExpiredSeconds(v int64) *PushBroadcastRequest
	GetExpiredSeconds() *int64
	SetExtendedParams(v string) *PushBroadcastRequest
	GetExtendedParams() *string
	SetMiChannelId(v string) *PushBroadcastRequest
	GetMiChannelId() *string
	SetMsgkey(v string) *PushBroadcastRequest
	GetMsgkey() *string
	SetNotifyType(v string) *PushBroadcastRequest
	GetNotifyType() *string
	SetPushAction(v int64) *PushBroadcastRequest
	GetPushAction() *int64
	SetPushStatus(v int64) *PushBroadcastRequest
	GetPushStatus() *int64
	SetSilent(v int64) *PushBroadcastRequest
	GetSilent() *int64
	SetStrategyContent(v string) *PushBroadcastRequest
	GetStrategyContent() *string
	SetStrategyType(v int32) *PushBroadcastRequest
	GetStrategyType() *int32
	SetTaskName(v string) *PushBroadcastRequest
	GetTaskName() *string
	SetTemplateKeyValue(v string) *PushBroadcastRequest
	GetTemplateKeyValue() *string
	SetTemplateName(v string) *PushBroadcastRequest
	GetTemplateName() *string
	SetTenantId(v string) *PushBroadcastRequest
	GetTenantId() *string
	SetThirdChannelCategory(v map[string]interface{}) *PushBroadcastRequest
	GetThirdChannelCategory() map[string]interface{}
	SetTransparentMessagePayload(v interface{}) *PushBroadcastRequest
	GetTransparentMessagePayload() interface{}
	SetTransparentMessageUrgency(v string) *PushBroadcastRequest
	GetTransparentMessageUrgency() *string
	SetUnBindPeriod(v int64) *PushBroadcastRequest
	GetUnBindPeriod() *int64
	SetWorkspaceId(v string) *PushBroadcastRequest
	GetWorkspaceId() *string
}

type PushBroadcastRequest struct {
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
	TemplateName              *string                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TenantId                  *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdChannelCategory      map[string]interface{} `json:"ThirdChannelCategory,omitempty" xml:"ThirdChannelCategory,omitempty"`
	TransparentMessagePayload interface{}            `json:"TransparentMessagePayload,omitempty" xml:"TransparentMessagePayload,omitempty"`
	TransparentMessageUrgency *string                `json:"TransparentMessageUrgency,omitempty" xml:"TransparentMessageUrgency,omitempty"`
	UnBindPeriod              *int64                 `json:"UnBindPeriod,omitempty" xml:"UnBindPeriod,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushBroadcastRequest) String() string {
	return dara.Prettify(s)
}

func (s PushBroadcastRequest) GoString() string {
	return s.String()
}

func (s *PushBroadcastRequest) GetAndroidChannel() *int32 {
	return s.AndroidChannel
}

func (s *PushBroadcastRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushBroadcastRequest) GetBindPeriod() *int32 {
	return s.BindPeriod
}

func (s *PushBroadcastRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *PushBroadcastRequest) GetClassification() *string {
	return s.Classification
}

func (s *PushBroadcastRequest) GetDeliveryType() *int64 {
	return s.DeliveryType
}

func (s *PushBroadcastRequest) GetExpiredSeconds() *int64 {
	return s.ExpiredSeconds
}

func (s *PushBroadcastRequest) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushBroadcastRequest) GetMiChannelId() *string {
	return s.MiChannelId
}

func (s *PushBroadcastRequest) GetMsgkey() *string {
	return s.Msgkey
}

func (s *PushBroadcastRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *PushBroadcastRequest) GetPushAction() *int64 {
	return s.PushAction
}

func (s *PushBroadcastRequest) GetPushStatus() *int64 {
	return s.PushStatus
}

func (s *PushBroadcastRequest) GetSilent() *int64 {
	return s.Silent
}

func (s *PushBroadcastRequest) GetStrategyContent() *string {
	return s.StrategyContent
}

func (s *PushBroadcastRequest) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *PushBroadcastRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *PushBroadcastRequest) GetTemplateKeyValue() *string {
	return s.TemplateKeyValue
}

func (s *PushBroadcastRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *PushBroadcastRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushBroadcastRequest) GetThirdChannelCategory() map[string]interface{} {
	return s.ThirdChannelCategory
}

func (s *PushBroadcastRequest) GetTransparentMessagePayload() interface{} {
	return s.TransparentMessagePayload
}

func (s *PushBroadcastRequest) GetTransparentMessageUrgency() *string {
	return s.TransparentMessageUrgency
}

func (s *PushBroadcastRequest) GetUnBindPeriod() *int64 {
	return s.UnBindPeriod
}

func (s *PushBroadcastRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushBroadcastRequest) SetAndroidChannel(v int32) *PushBroadcastRequest {
	s.AndroidChannel = &v
	return s
}

func (s *PushBroadcastRequest) SetAppId(v string) *PushBroadcastRequest {
	s.AppId = &v
	return s
}

func (s *PushBroadcastRequest) SetBindPeriod(v int32) *PushBroadcastRequest {
	s.BindPeriod = &v
	return s
}

func (s *PushBroadcastRequest) SetChannelId(v string) *PushBroadcastRequest {
	s.ChannelId = &v
	return s
}

func (s *PushBroadcastRequest) SetClassification(v string) *PushBroadcastRequest {
	s.Classification = &v
	return s
}

func (s *PushBroadcastRequest) SetDeliveryType(v int64) *PushBroadcastRequest {
	s.DeliveryType = &v
	return s
}

func (s *PushBroadcastRequest) SetExpiredSeconds(v int64) *PushBroadcastRequest {
	s.ExpiredSeconds = &v
	return s
}

func (s *PushBroadcastRequest) SetExtendedParams(v string) *PushBroadcastRequest {
	s.ExtendedParams = &v
	return s
}

func (s *PushBroadcastRequest) SetMiChannelId(v string) *PushBroadcastRequest {
	s.MiChannelId = &v
	return s
}

func (s *PushBroadcastRequest) SetMsgkey(v string) *PushBroadcastRequest {
	s.Msgkey = &v
	return s
}

func (s *PushBroadcastRequest) SetNotifyType(v string) *PushBroadcastRequest {
	s.NotifyType = &v
	return s
}

func (s *PushBroadcastRequest) SetPushAction(v int64) *PushBroadcastRequest {
	s.PushAction = &v
	return s
}

func (s *PushBroadcastRequest) SetPushStatus(v int64) *PushBroadcastRequest {
	s.PushStatus = &v
	return s
}

func (s *PushBroadcastRequest) SetSilent(v int64) *PushBroadcastRequest {
	s.Silent = &v
	return s
}

func (s *PushBroadcastRequest) SetStrategyContent(v string) *PushBroadcastRequest {
	s.StrategyContent = &v
	return s
}

func (s *PushBroadcastRequest) SetStrategyType(v int32) *PushBroadcastRequest {
	s.StrategyType = &v
	return s
}

func (s *PushBroadcastRequest) SetTaskName(v string) *PushBroadcastRequest {
	s.TaskName = &v
	return s
}

func (s *PushBroadcastRequest) SetTemplateKeyValue(v string) *PushBroadcastRequest {
	s.TemplateKeyValue = &v
	return s
}

func (s *PushBroadcastRequest) SetTemplateName(v string) *PushBroadcastRequest {
	s.TemplateName = &v
	return s
}

func (s *PushBroadcastRequest) SetTenantId(v string) *PushBroadcastRequest {
	s.TenantId = &v
	return s
}

func (s *PushBroadcastRequest) SetThirdChannelCategory(v map[string]interface{}) *PushBroadcastRequest {
	s.ThirdChannelCategory = v
	return s
}

func (s *PushBroadcastRequest) SetTransparentMessagePayload(v interface{}) *PushBroadcastRequest {
	s.TransparentMessagePayload = v
	return s
}

func (s *PushBroadcastRequest) SetTransparentMessageUrgency(v string) *PushBroadcastRequest {
	s.TransparentMessageUrgency = &v
	return s
}

func (s *PushBroadcastRequest) SetUnBindPeriod(v int64) *PushBroadcastRequest {
	s.UnBindPeriod = &v
	return s
}

func (s *PushBroadcastRequest) SetWorkspaceId(v string) *PushBroadcastRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushBroadcastRequest) Validate() error {
	return dara.Validate(s)
}

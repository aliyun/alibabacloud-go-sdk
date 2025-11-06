// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushSimpleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityContentState(v interface{}) *PushSimpleShrinkRequest
	GetActivityContentState() interface{}
	SetActivityEvent(v string) *PushSimpleShrinkRequest
	GetActivityEvent() *string
	SetAppId(v string) *PushSimpleShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *PushSimpleShrinkRequest
	GetChannelId() *string
	SetClassification(v string) *PushSimpleShrinkRequest
	GetClassification() *string
	SetContent(v string) *PushSimpleShrinkRequest
	GetContent() *string
	SetDeliveryType(v int64) *PushSimpleShrinkRequest
	GetDeliveryType() *int64
	SetDismissalDate(v int64) *PushSimpleShrinkRequest
	GetDismissalDate() *int64
	SetExpiredSeconds(v int64) *PushSimpleShrinkRequest
	GetExpiredSeconds() *int64
	SetExtendedParams(v string) *PushSimpleShrinkRequest
	GetExtendedParams() *string
	SetIconUrls(v string) *PushSimpleShrinkRequest
	GetIconUrls() *string
	SetImageUrls(v string) *PushSimpleShrinkRequest
	GetImageUrls() *string
	SetMiChannelId(v string) *PushSimpleShrinkRequest
	GetMiChannelId() *string
	SetNotifyLevelShrink(v string) *PushSimpleShrinkRequest
	GetNotifyLevelShrink() *string
	SetNotifyType(v string) *PushSimpleShrinkRequest
	GetNotifyType() *string
	SetPushAction(v int64) *PushSimpleShrinkRequest
	GetPushAction() *int64
	SetPushStyle(v int32) *PushSimpleShrinkRequest
	GetPushStyle() *int32
	SetSilent(v int64) *PushSimpleShrinkRequest
	GetSilent() *int64
	SetSmsSignName(v string) *PushSimpleShrinkRequest
	GetSmsSignName() *string
	SetSmsStrategy(v int32) *PushSimpleShrinkRequest
	GetSmsStrategy() *int32
	SetSmsTemplateCode(v string) *PushSimpleShrinkRequest
	GetSmsTemplateCode() *string
	SetSmsTemplateParam(v string) *PushSimpleShrinkRequest
	GetSmsTemplateParam() *string
	SetStrategyContent(v string) *PushSimpleShrinkRequest
	GetStrategyContent() *string
	SetStrategyType(v int32) *PushSimpleShrinkRequest
	GetStrategyType() *int32
	SetTargetMsgkey(v string) *PushSimpleShrinkRequest
	GetTargetMsgkey() *string
	SetTaskName(v string) *PushSimpleShrinkRequest
	GetTaskName() *string
	SetTenantId(v string) *PushSimpleShrinkRequest
	GetTenantId() *string
	SetThirdChannelCategoryShrink(v string) *PushSimpleShrinkRequest
	GetThirdChannelCategoryShrink() *string
	SetTitle(v string) *PushSimpleShrinkRequest
	GetTitle() *string
	SetTransparentMessagePayload(v interface{}) *PushSimpleShrinkRequest
	GetTransparentMessagePayload() interface{}
	SetTransparentMessageUrgency(v string) *PushSimpleShrinkRequest
	GetTransparentMessageUrgency() *string
	SetUri(v string) *PushSimpleShrinkRequest
	GetUri() *string
	SetWorkspaceId(v string) *PushSimpleShrinkRequest
	GetWorkspaceId() *string
}

type PushSimpleShrinkRequest struct {
	ActivityContentState interface{} `json:"ActivityContentState,omitempty" xml:"ActivityContentState,omitempty"`
	ActivityEvent        *string     `json:"ActivityEvent,omitempty" xml:"ActivityEvent,omitempty"`
	// This parameter is required.
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId      *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	DeliveryType  *int64 `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	DismissalDate *int64 `json:"DismissalDate,omitempty" xml:"DismissalDate,omitempty"`
	// This parameter is required.
	ExpiredSeconds    *int64  `json:"ExpiredSeconds,omitempty" xml:"ExpiredSeconds,omitempty"`
	ExtendedParams    *string `json:"ExtendedParams,omitempty" xml:"ExtendedParams,omitempty"`
	IconUrls          *string `json:"IconUrls,omitempty" xml:"IconUrls,omitempty"`
	ImageUrls         *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	MiChannelId       *string `json:"MiChannelId,omitempty" xml:"MiChannelId,omitempty"`
	NotifyLevelShrink *string `json:"NotifyLevel,omitempty" xml:"NotifyLevel,omitempty"`
	NotifyType        *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	PushAction        *int64  `json:"PushAction,omitempty" xml:"PushAction,omitempty"`
	PushStyle         *int32  `json:"PushStyle,omitempty" xml:"PushStyle,omitempty"`
	Silent            *int64  `json:"Silent,omitempty" xml:"Silent,omitempty"`
	SmsSignName       *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	SmsStrategy       *int32  `json:"SmsStrategy,omitempty" xml:"SmsStrategy,omitempty"`
	SmsTemplateCode   *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	SmsTemplateParam  *string `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	StrategyContent   *string `json:"StrategyContent,omitempty" xml:"StrategyContent,omitempty"`
	StrategyType      *int32  `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// This parameter is required.
	TargetMsgkey               *string `json:"TargetMsgkey,omitempty" xml:"TargetMsgkey,omitempty"`
	TaskName                   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TenantId                   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdChannelCategoryShrink *string `json:"ThirdChannelCategory,omitempty" xml:"ThirdChannelCategory,omitempty"`
	// This parameter is required.
	Title                     *string     `json:"Title,omitempty" xml:"Title,omitempty"`
	TransparentMessagePayload interface{} `json:"TransparentMessagePayload,omitempty" xml:"TransparentMessagePayload,omitempty"`
	TransparentMessageUrgency *string     `json:"TransparentMessageUrgency,omitempty" xml:"TransparentMessageUrgency,omitempty"`
	Uri                       *string     `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushSimpleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PushSimpleShrinkRequest) GoString() string {
	return s.String()
}

func (s *PushSimpleShrinkRequest) GetActivityContentState() interface{} {
	return s.ActivityContentState
}

func (s *PushSimpleShrinkRequest) GetActivityEvent() *string {
	return s.ActivityEvent
}

func (s *PushSimpleShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushSimpleShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *PushSimpleShrinkRequest) GetClassification() *string {
	return s.Classification
}

func (s *PushSimpleShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *PushSimpleShrinkRequest) GetDeliveryType() *int64 {
	return s.DeliveryType
}

func (s *PushSimpleShrinkRequest) GetDismissalDate() *int64 {
	return s.DismissalDate
}

func (s *PushSimpleShrinkRequest) GetExpiredSeconds() *int64 {
	return s.ExpiredSeconds
}

func (s *PushSimpleShrinkRequest) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushSimpleShrinkRequest) GetIconUrls() *string {
	return s.IconUrls
}

func (s *PushSimpleShrinkRequest) GetImageUrls() *string {
	return s.ImageUrls
}

func (s *PushSimpleShrinkRequest) GetMiChannelId() *string {
	return s.MiChannelId
}

func (s *PushSimpleShrinkRequest) GetNotifyLevelShrink() *string {
	return s.NotifyLevelShrink
}

func (s *PushSimpleShrinkRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *PushSimpleShrinkRequest) GetPushAction() *int64 {
	return s.PushAction
}

func (s *PushSimpleShrinkRequest) GetPushStyle() *int32 {
	return s.PushStyle
}

func (s *PushSimpleShrinkRequest) GetSilent() *int64 {
	return s.Silent
}

func (s *PushSimpleShrinkRequest) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *PushSimpleShrinkRequest) GetSmsStrategy() *int32 {
	return s.SmsStrategy
}

func (s *PushSimpleShrinkRequest) GetSmsTemplateCode() *string {
	return s.SmsTemplateCode
}

func (s *PushSimpleShrinkRequest) GetSmsTemplateParam() *string {
	return s.SmsTemplateParam
}

func (s *PushSimpleShrinkRequest) GetStrategyContent() *string {
	return s.StrategyContent
}

func (s *PushSimpleShrinkRequest) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *PushSimpleShrinkRequest) GetTargetMsgkey() *string {
	return s.TargetMsgkey
}

func (s *PushSimpleShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *PushSimpleShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushSimpleShrinkRequest) GetThirdChannelCategoryShrink() *string {
	return s.ThirdChannelCategoryShrink
}

func (s *PushSimpleShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *PushSimpleShrinkRequest) GetTransparentMessagePayload() interface{} {
	return s.TransparentMessagePayload
}

func (s *PushSimpleShrinkRequest) GetTransparentMessageUrgency() *string {
	return s.TransparentMessageUrgency
}

func (s *PushSimpleShrinkRequest) GetUri() *string {
	return s.Uri
}

func (s *PushSimpleShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushSimpleShrinkRequest) SetActivityContentState(v interface{}) *PushSimpleShrinkRequest {
	s.ActivityContentState = v
	return s
}

func (s *PushSimpleShrinkRequest) SetActivityEvent(v string) *PushSimpleShrinkRequest {
	s.ActivityEvent = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetAppId(v string) *PushSimpleShrinkRequest {
	s.AppId = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetChannelId(v string) *PushSimpleShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetClassification(v string) *PushSimpleShrinkRequest {
	s.Classification = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetContent(v string) *PushSimpleShrinkRequest {
	s.Content = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetDeliveryType(v int64) *PushSimpleShrinkRequest {
	s.DeliveryType = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetDismissalDate(v int64) *PushSimpleShrinkRequest {
	s.DismissalDate = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetExpiredSeconds(v int64) *PushSimpleShrinkRequest {
	s.ExpiredSeconds = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetExtendedParams(v string) *PushSimpleShrinkRequest {
	s.ExtendedParams = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetIconUrls(v string) *PushSimpleShrinkRequest {
	s.IconUrls = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetImageUrls(v string) *PushSimpleShrinkRequest {
	s.ImageUrls = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetMiChannelId(v string) *PushSimpleShrinkRequest {
	s.MiChannelId = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetNotifyLevelShrink(v string) *PushSimpleShrinkRequest {
	s.NotifyLevelShrink = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetNotifyType(v string) *PushSimpleShrinkRequest {
	s.NotifyType = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetPushAction(v int64) *PushSimpleShrinkRequest {
	s.PushAction = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetPushStyle(v int32) *PushSimpleShrinkRequest {
	s.PushStyle = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetSilent(v int64) *PushSimpleShrinkRequest {
	s.Silent = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetSmsSignName(v string) *PushSimpleShrinkRequest {
	s.SmsSignName = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetSmsStrategy(v int32) *PushSimpleShrinkRequest {
	s.SmsStrategy = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetSmsTemplateCode(v string) *PushSimpleShrinkRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetSmsTemplateParam(v string) *PushSimpleShrinkRequest {
	s.SmsTemplateParam = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetStrategyContent(v string) *PushSimpleShrinkRequest {
	s.StrategyContent = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetStrategyType(v int32) *PushSimpleShrinkRequest {
	s.StrategyType = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetTargetMsgkey(v string) *PushSimpleShrinkRequest {
	s.TargetMsgkey = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetTaskName(v string) *PushSimpleShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetTenantId(v string) *PushSimpleShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetThirdChannelCategoryShrink(v string) *PushSimpleShrinkRequest {
	s.ThirdChannelCategoryShrink = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetTitle(v string) *PushSimpleShrinkRequest {
	s.Title = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetTransparentMessagePayload(v interface{}) *PushSimpleShrinkRequest {
	s.TransparentMessagePayload = v
	return s
}

func (s *PushSimpleShrinkRequest) SetTransparentMessageUrgency(v string) *PushSimpleShrinkRequest {
	s.TransparentMessageUrgency = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetUri(v string) *PushSimpleShrinkRequest {
	s.Uri = &v
	return s
}

func (s *PushSimpleShrinkRequest) SetWorkspaceId(v string) *PushSimpleShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushSimpleShrinkRequest) Validate() error {
	return dara.Validate(s)
}

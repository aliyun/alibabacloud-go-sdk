// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushSimpleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityContentState(v interface{}) *PushSimpleRequest
	GetActivityContentState() interface{}
	SetActivityEvent(v string) *PushSimpleRequest
	GetActivityEvent() *string
	SetAppId(v string) *PushSimpleRequest
	GetAppId() *string
	SetChannelId(v string) *PushSimpleRequest
	GetChannelId() *string
	SetClassification(v string) *PushSimpleRequest
	GetClassification() *string
	SetContent(v string) *PushSimpleRequest
	GetContent() *string
	SetDeliveryType(v int64) *PushSimpleRequest
	GetDeliveryType() *int64
	SetDismissalDate(v int64) *PushSimpleRequest
	GetDismissalDate() *int64
	SetExpiredSeconds(v int64) *PushSimpleRequest
	GetExpiredSeconds() *int64
	SetExtendedParams(v string) *PushSimpleRequest
	GetExtendedParams() *string
	SetIconUrls(v string) *PushSimpleRequest
	GetIconUrls() *string
	SetImageUrls(v string) *PushSimpleRequest
	GetImageUrls() *string
	SetMiChannelId(v string) *PushSimpleRequest
	GetMiChannelId() *string
	SetNotifyLevel(v map[string]interface{}) *PushSimpleRequest
	GetNotifyLevel() map[string]interface{}
	SetNotifyType(v string) *PushSimpleRequest
	GetNotifyType() *string
	SetPushAction(v int64) *PushSimpleRequest
	GetPushAction() *int64
	SetPushStyle(v int32) *PushSimpleRequest
	GetPushStyle() *int32
	SetSilent(v int64) *PushSimpleRequest
	GetSilent() *int64
	SetSmsSignName(v string) *PushSimpleRequest
	GetSmsSignName() *string
	SetSmsStrategy(v int32) *PushSimpleRequest
	GetSmsStrategy() *int32
	SetSmsTemplateCode(v string) *PushSimpleRequest
	GetSmsTemplateCode() *string
	SetSmsTemplateParam(v string) *PushSimpleRequest
	GetSmsTemplateParam() *string
	SetStrategyContent(v string) *PushSimpleRequest
	GetStrategyContent() *string
	SetStrategyType(v int32) *PushSimpleRequest
	GetStrategyType() *int32
	SetTargetMsgkey(v string) *PushSimpleRequest
	GetTargetMsgkey() *string
	SetTaskName(v string) *PushSimpleRequest
	GetTaskName() *string
	SetTenantId(v string) *PushSimpleRequest
	GetTenantId() *string
	SetThirdChannelCategory(v map[string]interface{}) *PushSimpleRequest
	GetThirdChannelCategory() map[string]interface{}
	SetTitle(v string) *PushSimpleRequest
	GetTitle() *string
	SetTransparentMessagePayload(v interface{}) *PushSimpleRequest
	GetTransparentMessagePayload() interface{}
	SetTransparentMessageUrgency(v string) *PushSimpleRequest
	GetTransparentMessageUrgency() *string
	SetUri(v string) *PushSimpleRequest
	GetUri() *string
	SetWorkspaceId(v string) *PushSimpleRequest
	GetWorkspaceId() *string
}

type PushSimpleRequest struct {
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
	ExpiredSeconds   *int64                 `json:"ExpiredSeconds,omitempty" xml:"ExpiredSeconds,omitempty"`
	ExtendedParams   *string                `json:"ExtendedParams,omitempty" xml:"ExtendedParams,omitempty"`
	IconUrls         *string                `json:"IconUrls,omitempty" xml:"IconUrls,omitempty"`
	ImageUrls        *string                `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	MiChannelId      *string                `json:"MiChannelId,omitempty" xml:"MiChannelId,omitempty"`
	NotifyLevel      map[string]interface{} `json:"NotifyLevel,omitempty" xml:"NotifyLevel,omitempty"`
	NotifyType       *string                `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	PushAction       *int64                 `json:"PushAction,omitempty" xml:"PushAction,omitempty"`
	PushStyle        *int32                 `json:"PushStyle,omitempty" xml:"PushStyle,omitempty"`
	Silent           *int64                 `json:"Silent,omitempty" xml:"Silent,omitempty"`
	SmsSignName      *string                `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	SmsStrategy      *int32                 `json:"SmsStrategy,omitempty" xml:"SmsStrategy,omitempty"`
	SmsTemplateCode  *string                `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
	SmsTemplateParam *string                `json:"SmsTemplateParam,omitempty" xml:"SmsTemplateParam,omitempty"`
	StrategyContent  *string                `json:"StrategyContent,omitempty" xml:"StrategyContent,omitempty"`
	StrategyType     *int32                 `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	// This parameter is required.
	TargetMsgkey         *string                `json:"TargetMsgkey,omitempty" xml:"TargetMsgkey,omitempty"`
	TaskName             *string                `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TenantId             *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ThirdChannelCategory map[string]interface{} `json:"ThirdChannelCategory,omitempty" xml:"ThirdChannelCategory,omitempty"`
	// This parameter is required.
	Title                     *string     `json:"Title,omitempty" xml:"Title,omitempty"`
	TransparentMessagePayload interface{} `json:"TransparentMessagePayload,omitempty" xml:"TransparentMessagePayload,omitempty"`
	TransparentMessageUrgency *string     `json:"TransparentMessageUrgency,omitempty" xml:"TransparentMessageUrgency,omitempty"`
	Uri                       *string     `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PushSimpleRequest) String() string {
	return dara.Prettify(s)
}

func (s PushSimpleRequest) GoString() string {
	return s.String()
}

func (s *PushSimpleRequest) GetActivityContentState() interface{} {
	return s.ActivityContentState
}

func (s *PushSimpleRequest) GetActivityEvent() *string {
	return s.ActivityEvent
}

func (s *PushSimpleRequest) GetAppId() *string {
	return s.AppId
}

func (s *PushSimpleRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *PushSimpleRequest) GetClassification() *string {
	return s.Classification
}

func (s *PushSimpleRequest) GetContent() *string {
	return s.Content
}

func (s *PushSimpleRequest) GetDeliveryType() *int64 {
	return s.DeliveryType
}

func (s *PushSimpleRequest) GetDismissalDate() *int64 {
	return s.DismissalDate
}

func (s *PushSimpleRequest) GetExpiredSeconds() *int64 {
	return s.ExpiredSeconds
}

func (s *PushSimpleRequest) GetExtendedParams() *string {
	return s.ExtendedParams
}

func (s *PushSimpleRequest) GetIconUrls() *string {
	return s.IconUrls
}

func (s *PushSimpleRequest) GetImageUrls() *string {
	return s.ImageUrls
}

func (s *PushSimpleRequest) GetMiChannelId() *string {
	return s.MiChannelId
}

func (s *PushSimpleRequest) GetNotifyLevel() map[string]interface{} {
	return s.NotifyLevel
}

func (s *PushSimpleRequest) GetNotifyType() *string {
	return s.NotifyType
}

func (s *PushSimpleRequest) GetPushAction() *int64 {
	return s.PushAction
}

func (s *PushSimpleRequest) GetPushStyle() *int32 {
	return s.PushStyle
}

func (s *PushSimpleRequest) GetSilent() *int64 {
	return s.Silent
}

func (s *PushSimpleRequest) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *PushSimpleRequest) GetSmsStrategy() *int32 {
	return s.SmsStrategy
}

func (s *PushSimpleRequest) GetSmsTemplateCode() *string {
	return s.SmsTemplateCode
}

func (s *PushSimpleRequest) GetSmsTemplateParam() *string {
	return s.SmsTemplateParam
}

func (s *PushSimpleRequest) GetStrategyContent() *string {
	return s.StrategyContent
}

func (s *PushSimpleRequest) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *PushSimpleRequest) GetTargetMsgkey() *string {
	return s.TargetMsgkey
}

func (s *PushSimpleRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *PushSimpleRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PushSimpleRequest) GetThirdChannelCategory() map[string]interface{} {
	return s.ThirdChannelCategory
}

func (s *PushSimpleRequest) GetTitle() *string {
	return s.Title
}

func (s *PushSimpleRequest) GetTransparentMessagePayload() interface{} {
	return s.TransparentMessagePayload
}

func (s *PushSimpleRequest) GetTransparentMessageUrgency() *string {
	return s.TransparentMessageUrgency
}

func (s *PushSimpleRequest) GetUri() *string {
	return s.Uri
}

func (s *PushSimpleRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PushSimpleRequest) SetActivityContentState(v interface{}) *PushSimpleRequest {
	s.ActivityContentState = v
	return s
}

func (s *PushSimpleRequest) SetActivityEvent(v string) *PushSimpleRequest {
	s.ActivityEvent = &v
	return s
}

func (s *PushSimpleRequest) SetAppId(v string) *PushSimpleRequest {
	s.AppId = &v
	return s
}

func (s *PushSimpleRequest) SetChannelId(v string) *PushSimpleRequest {
	s.ChannelId = &v
	return s
}

func (s *PushSimpleRequest) SetClassification(v string) *PushSimpleRequest {
	s.Classification = &v
	return s
}

func (s *PushSimpleRequest) SetContent(v string) *PushSimpleRequest {
	s.Content = &v
	return s
}

func (s *PushSimpleRequest) SetDeliveryType(v int64) *PushSimpleRequest {
	s.DeliveryType = &v
	return s
}

func (s *PushSimpleRequest) SetDismissalDate(v int64) *PushSimpleRequest {
	s.DismissalDate = &v
	return s
}

func (s *PushSimpleRequest) SetExpiredSeconds(v int64) *PushSimpleRequest {
	s.ExpiredSeconds = &v
	return s
}

func (s *PushSimpleRequest) SetExtendedParams(v string) *PushSimpleRequest {
	s.ExtendedParams = &v
	return s
}

func (s *PushSimpleRequest) SetIconUrls(v string) *PushSimpleRequest {
	s.IconUrls = &v
	return s
}

func (s *PushSimpleRequest) SetImageUrls(v string) *PushSimpleRequest {
	s.ImageUrls = &v
	return s
}

func (s *PushSimpleRequest) SetMiChannelId(v string) *PushSimpleRequest {
	s.MiChannelId = &v
	return s
}

func (s *PushSimpleRequest) SetNotifyLevel(v map[string]interface{}) *PushSimpleRequest {
	s.NotifyLevel = v
	return s
}

func (s *PushSimpleRequest) SetNotifyType(v string) *PushSimpleRequest {
	s.NotifyType = &v
	return s
}

func (s *PushSimpleRequest) SetPushAction(v int64) *PushSimpleRequest {
	s.PushAction = &v
	return s
}

func (s *PushSimpleRequest) SetPushStyle(v int32) *PushSimpleRequest {
	s.PushStyle = &v
	return s
}

func (s *PushSimpleRequest) SetSilent(v int64) *PushSimpleRequest {
	s.Silent = &v
	return s
}

func (s *PushSimpleRequest) SetSmsSignName(v string) *PushSimpleRequest {
	s.SmsSignName = &v
	return s
}

func (s *PushSimpleRequest) SetSmsStrategy(v int32) *PushSimpleRequest {
	s.SmsStrategy = &v
	return s
}

func (s *PushSimpleRequest) SetSmsTemplateCode(v string) *PushSimpleRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *PushSimpleRequest) SetSmsTemplateParam(v string) *PushSimpleRequest {
	s.SmsTemplateParam = &v
	return s
}

func (s *PushSimpleRequest) SetStrategyContent(v string) *PushSimpleRequest {
	s.StrategyContent = &v
	return s
}

func (s *PushSimpleRequest) SetStrategyType(v int32) *PushSimpleRequest {
	s.StrategyType = &v
	return s
}

func (s *PushSimpleRequest) SetTargetMsgkey(v string) *PushSimpleRequest {
	s.TargetMsgkey = &v
	return s
}

func (s *PushSimpleRequest) SetTaskName(v string) *PushSimpleRequest {
	s.TaskName = &v
	return s
}

func (s *PushSimpleRequest) SetTenantId(v string) *PushSimpleRequest {
	s.TenantId = &v
	return s
}

func (s *PushSimpleRequest) SetThirdChannelCategory(v map[string]interface{}) *PushSimpleRequest {
	s.ThirdChannelCategory = v
	return s
}

func (s *PushSimpleRequest) SetTitle(v string) *PushSimpleRequest {
	s.Title = &v
	return s
}

func (s *PushSimpleRequest) SetTransparentMessagePayload(v interface{}) *PushSimpleRequest {
	s.TransparentMessagePayload = v
	return s
}

func (s *PushSimpleRequest) SetTransparentMessageUrgency(v string) *PushSimpleRequest {
	s.TransparentMessageUrgency = &v
	return s
}

func (s *PushSimpleRequest) SetUri(v string) *PushSimpleRequest {
	s.Uri = &v
	return s
}

func (s *PushSimpleRequest) SetWorkspaceId(v string) *PushSimpleRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PushSimpleRequest) Validate() error {
	return dara.Validate(s)
}

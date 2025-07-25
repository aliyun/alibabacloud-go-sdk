// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmGlobalAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmGlobalAlertRequest
	GetAcceptLanguage() *string
	SetAlertConfig(v []*UpdateCloudGtmGlobalAlertRequestAlertConfig) *UpdateCloudGtmGlobalAlertRequest
	GetAlertConfig() []*UpdateCloudGtmGlobalAlertRequestAlertConfig
	SetAlertGroup(v []*string) *UpdateCloudGtmGlobalAlertRequest
	GetAlertGroup() []*string
	SetClientToken(v string) *UpdateCloudGtmGlobalAlertRequest
	GetClientToken() *string
}

type UpdateCloudGtmGlobalAlertRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The alert configurations.
	AlertConfig []*UpdateCloudGtmGlobalAlertRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	// The alert contact groups.
	AlertGroup []*string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateCloudGtmGlobalAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmGlobalAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmGlobalAlertRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmGlobalAlertRequest) GetAlertConfig() []*UpdateCloudGtmGlobalAlertRequestAlertConfig {
	return s.AlertConfig
}

func (s *UpdateCloudGtmGlobalAlertRequest) GetAlertGroup() []*string {
	return s.AlertGroup
}

func (s *UpdateCloudGtmGlobalAlertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmGlobalAlertRequest) SetAcceptLanguage(v string) *UpdateCloudGtmGlobalAlertRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequest) SetAlertConfig(v []*UpdateCloudGtmGlobalAlertRequestAlertConfig) *UpdateCloudGtmGlobalAlertRequest {
	s.AlertConfig = v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequest) SetAlertGroup(v []*string) *UpdateCloudGtmGlobalAlertRequest {
	s.AlertGroup = v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequest) SetClientToken(v string) *UpdateCloudGtmGlobalAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudGtmGlobalAlertRequestAlertConfig struct {
	// Specifies whether to configure DingTalk notifications. Valid values:
	//
	// 	- true: configures DingTalk notifications. DingTalk notifications are sent when alerts are triggered.
	//
	// 	- false: does not configure DingTalk notifications.
	//
	// example:
	//
	// false
	DingtalkNotice *bool `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	// Specifies whether to configure email notifications. Valid values:
	//
	// 	- true: configures email notifications. Emails are sent when alerts are triggered.
	//
	// 	- false｜null: does not configure email notifications.
	//
	// example:
	//
	// true
	EmailNotice *bool `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	// The type of the alert event. Valid values:
	//
	// 	- addr_alert: The address is unavailable.
	//
	// 	- addr_resume: The address becomes available.
	//
	// 	- addr_pool_unavailable: The address pool is unavailable.
	//
	// 	- addr_pool_available: The address pool becomes available.
	//
	// example:
	//
	// addr_alert
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// Specifies whether to configure text message notifications. Valid values:
	//
	// 	- true: configures text message notifications. Text messages are sent when alerts are triggered.
	//
	// 	- false｜null: does not configure text message notifications.
	//
	// Only the China site (aliyun.com) supports text message notifications.
	//
	// example:
	//
	// true
	SmsNotice *bool `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
}

func (s UpdateCloudGtmGlobalAlertRequestAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmGlobalAlertRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) GetDingtalkNotice() *bool {
	return s.DingtalkNotice
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) GetEmailNotice() *bool {
	return s.EmailNotice
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) GetNoticeType() *string {
	return s.NoticeType
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) GetSmsNotice() *bool {
	return s.SmsNotice
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) SetDingtalkNotice(v bool) *UpdateCloudGtmGlobalAlertRequestAlertConfig {
	s.DingtalkNotice = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) SetEmailNotice(v bool) *UpdateCloudGtmGlobalAlertRequestAlertConfig {
	s.EmailNotice = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) SetNoticeType(v string) *UpdateCloudGtmGlobalAlertRequestAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) SetSmsNotice(v bool) *UpdateCloudGtmGlobalAlertRequestAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) Validate() error {
	return dara.Validate(s)
}

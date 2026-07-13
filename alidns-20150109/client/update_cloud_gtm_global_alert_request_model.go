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
	// - `zh-CN`: Chinese
	//
	// - `en-US`: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// A list of alert configurations.
	AlertConfig []*UpdateCloudGtmGlobalAlertRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	// A list of alert notification groups.
	AlertGroup []*string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Repeated"`
	// A client-generated token to ensure request idempotence. This token must be unique for each request, contain only ASCII characters, and be no more than 64 characters in length.
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
	if s.AlertConfig != nil {
		for _, item := range s.AlertConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCloudGtmGlobalAlertRequestAlertConfig struct {
	// Whether to send a DingTalk notification when an alert is triggered. Valid values:
	//
	// - `true`: A DingTalk notification is sent.
	//
	// - `false`: Do not send a DingTalk notification.
	//
	// example:
	//
	// false
	DingtalkNotice *bool `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	// Whether to send an email notification when an alert is triggered. Valid values:
	//
	// - `true`: An email notification is sent.
	//
	// - `false` or `null`: Do not send an email notification.
	//
	// example:
	//
	// true
	EmailNotice *bool `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	// The alert event type. Valid values:
	//
	// - `addr_alert`: An address becomes unavailable.
	//
	// - `addr_resume`: An address becomes available.
	//
	// - `addr_pool_unavailable`: An address pool becomes unavailable.
	//
	// - `addr_pool_available`: An address pool becomes available.
	//
	// example:
	//
	// addr_alert
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// The alert threshold for queries per second (QPS).
	//
	// example:
	//
	// 10
	QpsThreshold *int64 `json:"QpsThreshold,omitempty" xml:"QpsThreshold,omitempty"`
	// Whether to send a text message notification when an alert is triggered. Valid values:
	//
	// - `true`: A text message notification is sent.
	//
	// - `false` or `null`: Do not send a text message notification.
	//
	// Text message notifications are available only on the China site.
	//
	// example:
	//
	// true
	SmsNotice *bool `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
	// The alert threshold.
	//
	// example:
	//
	// 100
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
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

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) GetQpsThreshold() *int64 {
	return s.QpsThreshold
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) GetSmsNotice() *bool {
	return s.SmsNotice
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) GetThreshold() *int32 {
	return s.Threshold
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

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) SetQpsThreshold(v int64) *UpdateCloudGtmGlobalAlertRequestAlertConfig {
	s.QpsThreshold = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) SetSmsNotice(v bool) *UpdateCloudGtmGlobalAlertRequestAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) SetThreshold(v int32) *UpdateCloudGtmGlobalAlertRequestAlertConfig {
	s.Threshold = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertRequestAlertConfig) Validate() error {
	return dara.Validate(s)
}

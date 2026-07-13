// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigAlertRequest
	GetAcceptLanguage() *string
	SetAlertConfig(v []*UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) *UpdateCloudGtmInstanceConfigAlertRequest
	GetAlertConfig() []*UpdateCloudGtmInstanceConfigAlertRequestAlertConfig
	SetAlertGroup(v []*string) *UpdateCloudGtmInstanceConfigAlertRequest
	GetAlertGroup() []*string
	SetAlertMode(v string) *UpdateCloudGtmInstanceConfigAlertRequest
	GetAlertMode() *string
	SetClientToken(v string) *UpdateCloudGtmInstanceConfigAlertRequest
	GetClientToken() *string
	SetConfigId(v string) *UpdateCloudGtmInstanceConfigAlertRequest
	GetConfigId() *string
	SetInstanceId(v string) *UpdateCloudGtmInstanceConfigAlertRequest
	GetInstanceId() *string
}

type UpdateCloudGtmInstanceConfigAlertRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese
	//
	// - en-US: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// A list of alert configurations.
	AlertConfig []*UpdateCloudGtmInstanceConfigAlertRequestAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	// A list of alert notification groups.
	AlertGroup []*string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty" type:"Repeated"`
	// The alert configuration mode for the instance. Valid values:
	//
	// - global: The instance inherits the global alert configuration.
	//
	// - instance_config: The instance uses a custom alert configuration.
	//
	// example:
	//
	// global
	AlertMode *string `json:"AlertMode,omitempty" xml:"AlertMode,omitempty"`
	// A client-generated token that ensures the idempotence of the request. The client must generate a unique value for this parameter. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the domain name instance configuration. A GTM instance can have two configurations for the same access domain name if you configure both A and AAAA records. The ConfigId uniquely identifies a configuration.
	//
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
	//
	// example:
	//
	// gtm-cn-zz11t58**0s
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) GetAlertConfig() []*UpdateCloudGtmInstanceConfigAlertRequestAlertConfig {
	return s.AlertConfig
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) GetAlertGroup() []*string {
	return s.AlertGroup
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) GetAlertMode() *string {
	return s.AlertMode
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigAlertRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) SetAlertConfig(v []*UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) *UpdateCloudGtmInstanceConfigAlertRequest {
	s.AlertConfig = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) SetAlertGroup(v []*string) *UpdateCloudGtmInstanceConfigAlertRequest {
	s.AlertGroup = v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) SetAlertMode(v string) *UpdateCloudGtmInstanceConfigAlertRequest {
	s.AlertMode = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) SetClientToken(v string) *UpdateCloudGtmInstanceConfigAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) SetConfigId(v string) *UpdateCloudGtmInstanceConfigAlertRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) SetInstanceId(v string) *UpdateCloudGtmInstanceConfigAlertRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequest) Validate() error {
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

type UpdateCloudGtmInstanceConfigAlertRequestAlertConfig struct {
	// Specifies whether to enable DingTalk notifications. Valid values:
	//
	// - true: Enabled. When an alert is triggered, a DingTalk message is sent.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	DingtalkNotice *bool `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	// Specifies whether to enable email notifications. Valid values:
	//
	// - true: Enabled. When an alert is triggered, an email is sent.
	//
	// - false: Disabled.
	//
	// example:
	//
	// true
	EmailNotice *bool `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	// The type of alert event. Valid values:
	//
	// - addr_alert: The address is unavailable.
	//
	// - addr_resume: The address has recovered.
	//
	// - addr_pool_unavailable: The address pool is unavailable.
	//
	// - addr_pool_available: The address pool has recovered.
	//
	// example:
	//
	// addr_alert
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// Specifies whether to enable text message notifications. Valid values:
	//
	// - true: Enabled. When an alert is triggered, a text message is sent.
	//
	// - false: Disabled.
	//
	// Note: Text message notifications are supported only on the China site (aliyun.com).
	//
	// example:
	//
	// true
	SmsNotice *bool `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) GetDingtalkNotice() *bool {
	return s.DingtalkNotice
}

func (s *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) GetEmailNotice() *bool {
	return s.EmailNotice
}

func (s *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) GetNoticeType() *string {
	return s.NoticeType
}

func (s *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) GetSmsNotice() *bool {
	return s.SmsNotice
}

func (s *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) SetDingtalkNotice(v bool) *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig {
	s.DingtalkNotice = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) SetEmailNotice(v bool) *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig {
	s.EmailNotice = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) SetNoticeType(v string) *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) SetSmsNotice(v bool) *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertRequestAlertConfig) Validate() error {
	return dara.Validate(s)
}

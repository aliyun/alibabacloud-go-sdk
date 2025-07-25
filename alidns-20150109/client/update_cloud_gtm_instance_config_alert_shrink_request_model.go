// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigAlertShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest
	GetAcceptLanguage() *string
	SetAlertConfigShrink(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest
	GetAlertConfigShrink() *string
	SetAlertGroupShrink(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest
	GetAlertGroupShrink() *string
	SetAlertMode(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest
	GetAlertMode() *string
	SetClientToken(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest
	GetClientToken() *string
	SetConfigId(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest
	GetConfigId() *string
	SetInstanceId(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest
	GetInstanceId() *string
}

type UpdateCloudGtmInstanceConfigAlertShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The alert configurations.
	AlertConfigShrink *string `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty"`
	// The alert contact groups.
	AlertGroupShrink *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The alert configuration mode of the instance. Valid values:
	//
	// 	- global: global alert configuration
	//
	// 	- instance_config: custom alert configuration
	//
	// example:
	//
	// global
	AlertMode *string `json:"AlertMode,omitempty" xml:"AlertMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration ID of the access domain name. Two configuration IDs exist when an A record and an AAAA record are configured for the access domain name that is bound to the GTM instance. This ID uniquely identifies a configuration.
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

func (s UpdateCloudGtmInstanceConfigAlertShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigAlertShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) GetAlertConfigShrink() *string {
	return s.AlertConfigShrink
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) GetAlertGroupShrink() *string {
	return s.AlertGroupShrink
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) GetAlertMode() *string {
	return s.AlertMode
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) SetAlertConfigShrink(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest {
	s.AlertConfigShrink = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) SetAlertGroupShrink(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest {
	s.AlertGroupShrink = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) SetAlertMode(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest {
	s.AlertMode = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) SetClientToken(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) SetConfigId(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) SetInstanceId(v string) *UpdateCloudGtmInstanceConfigAlertShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigAlertShrinkRequest) Validate() error {
	return dara.Validate(s)
}

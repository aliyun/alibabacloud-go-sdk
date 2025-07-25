// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmGlobalAlertShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmGlobalAlertShrinkRequest
	GetAcceptLanguage() *string
	SetAlertConfigShrink(v string) *UpdateCloudGtmGlobalAlertShrinkRequest
	GetAlertConfigShrink() *string
	SetAlertGroupShrink(v string) *UpdateCloudGtmGlobalAlertShrinkRequest
	GetAlertGroupShrink() *string
	SetClientToken(v string) *UpdateCloudGtmGlobalAlertShrinkRequest
	GetClientToken() *string
}

type UpdateCloudGtmGlobalAlertShrinkRequest struct {
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
	AlertConfigShrink *string `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty"`
	// The alert contact groups.
	AlertGroupShrink *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateCloudGtmGlobalAlertShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmGlobalAlertShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmGlobalAlertShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmGlobalAlertShrinkRequest) GetAlertConfigShrink() *string {
	return s.AlertConfigShrink
}

func (s *UpdateCloudGtmGlobalAlertShrinkRequest) GetAlertGroupShrink() *string {
	return s.AlertGroupShrink
}

func (s *UpdateCloudGtmGlobalAlertShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmGlobalAlertShrinkRequest) SetAcceptLanguage(v string) *UpdateCloudGtmGlobalAlertShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertShrinkRequest) SetAlertConfigShrink(v string) *UpdateCloudGtmGlobalAlertShrinkRequest {
	s.AlertConfigShrink = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertShrinkRequest) SetAlertGroupShrink(v string) *UpdateCloudGtmGlobalAlertShrinkRequest {
	s.AlertGroupShrink = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertShrinkRequest) SetClientToken(v string) *UpdateCloudGtmGlobalAlertShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmGlobalAlertShrinkRequest) Validate() error {
	return dara.Validate(s)
}

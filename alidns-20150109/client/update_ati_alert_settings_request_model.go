// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAtiAlertSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertConfig(v string) *UpdateAtiAlertSettingsRequest
	GetAlertConfig() *string
	SetAlertGroup(v string) *UpdateAtiAlertSettingsRequest
	GetAlertGroup() *string
	SetClientToken(v string) *UpdateAtiAlertSettingsRequest
	GetClientToken() *string
}

type UpdateAtiAlertSettingsRequest struct {
	// The list of alert configurations.
	//
	// example:
	//
	// "[{\\"NoticeType\\":\\"identity_cert_expiring\\",\\"SmsNotice\\":true,\\"EmailNotice\\":true,\\"DingtalkNotice\\":true},{\\"NoticeType\\":\\"server_cert_expiring\\",\\"SmsNotice\\":true,\\"EmailNotice\\":true,\\"DingtalkNotice\\":true}]"
	AlertConfig *string `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty"`
	// The list of alert notification groups.
	//
	// example:
	//
	// [\\"Cloud account alert contact\\"]
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s UpdateAtiAlertSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAtiAlertSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateAtiAlertSettingsRequest) GetAlertConfig() *string {
	return s.AlertConfig
}

func (s *UpdateAtiAlertSettingsRequest) GetAlertGroup() *string {
	return s.AlertGroup
}

func (s *UpdateAtiAlertSettingsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAtiAlertSettingsRequest) SetAlertConfig(v string) *UpdateAtiAlertSettingsRequest {
	s.AlertConfig = &v
	return s
}

func (s *UpdateAtiAlertSettingsRequest) SetAlertGroup(v string) *UpdateAtiAlertSettingsRequest {
	s.AlertGroup = &v
	return s
}

func (s *UpdateAtiAlertSettingsRequest) SetClientToken(v string) *UpdateAtiAlertSettingsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAtiAlertSettingsRequest) Validate() error {
	return dara.Validate(s)
}

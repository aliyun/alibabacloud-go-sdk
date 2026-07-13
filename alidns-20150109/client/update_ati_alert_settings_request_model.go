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
	// example:
	//
	// "[{\\"NoticeType\\":\\"identity_cert_expiring\\",\\"SmsNotice\\":true,\\"EmailNotice\\":true,\\"DingtalkNotice\\":true},{\\"NoticeType\\":\\"server_cert_expiring\\",\\"SmsNotice\\":true,\\"EmailNotice\\":true,\\"DingtalkNotice\\":true}]"
	AlertConfig *string `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty"`
	// example:
	//
	// [\\"云账号报警联系人\\"]
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
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

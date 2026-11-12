// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMyAgenticNotificationChannelConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *UpdateMyAgenticNotificationChannelConfigRequest
	GetChannel() *string
	SetConfig(v string) *UpdateMyAgenticNotificationChannelConfigRequest
	GetConfig() *string
	SetEnabled(v bool) *UpdateMyAgenticNotificationChannelConfigRequest
	GetEnabled() *bool
}

type UpdateMyAgenticNotificationChannelConfigRequest struct {
	// The notification channel type. Currently, only FEISHU (Lark) is supported. Specifying other values results in a parameter error.
	//
	// This parameter is required.
	//
	// example:
	//
	// FEISHU
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The channel configuration in the format of a JSON object string. This parameter is not required if you do not update the binding information. Supports mobile (phone number) or email. You can bind or rebind a Lark account by using one of them at a time. If the existing masked values of mobile or email are passed as ******, they are not modified. openId is a read-only field and cannot be specified.
	//
	// example:
	//
	// {"email":"user@example.com"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether to enable the notification channel. If this parameter is not specified, the original value is retained. When you only modify the switch, the Config parameter is not required, and the existing binding information remains unchanged.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s UpdateMyAgenticNotificationChannelConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMyAgenticNotificationChannelConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateMyAgenticNotificationChannelConfigRequest) GetChannel() *string {
	return s.Channel
}

func (s *UpdateMyAgenticNotificationChannelConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateMyAgenticNotificationChannelConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateMyAgenticNotificationChannelConfigRequest) SetChannel(v string) *UpdateMyAgenticNotificationChannelConfigRequest {
	s.Channel = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigRequest) SetConfig(v string) *UpdateMyAgenticNotificationChannelConfigRequest {
	s.Config = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigRequest) SetEnabled(v bool) *UpdateMyAgenticNotificationChannelConfigRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateMyAgenticNotificationChannelConfigRequest) Validate() error {
	return dara.Validate(s)
}

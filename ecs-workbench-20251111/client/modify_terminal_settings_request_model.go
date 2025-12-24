// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTerminalSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordlessLoginConfig(v *ModifyTerminalSettingsRequestPasswordlessLoginConfig) *ModifyTerminalSettingsRequest
	GetPasswordlessLoginConfig() *ModifyTerminalSettingsRequestPasswordlessLoginConfig
}

type ModifyTerminalSettingsRequest struct {
	// 免密登录配置
	//
	// This parameter is required.
	PasswordlessLoginConfig *ModifyTerminalSettingsRequestPasswordlessLoginConfig `json:"PasswordlessLoginConfig,omitempty" xml:"PasswordlessLoginConfig,omitempty" type:"Struct"`
}

func (s ModifyTerminalSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTerminalSettingsRequest) GoString() string {
	return s.String()
}

func (s *ModifyTerminalSettingsRequest) GetPasswordlessLoginConfig() *ModifyTerminalSettingsRequestPasswordlessLoginConfig {
	return s.PasswordlessLoginConfig
}

func (s *ModifyTerminalSettingsRequest) SetPasswordlessLoginConfig(v *ModifyTerminalSettingsRequestPasswordlessLoginConfig) *ModifyTerminalSettingsRequest {
	s.PasswordlessLoginConfig = v
	return s
}

func (s *ModifyTerminalSettingsRequest) Validate() error {
	if s.PasswordlessLoginConfig != nil {
		if err := s.PasswordlessLoginConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyTerminalSettingsRequestPasswordlessLoginConfig struct {
	// 免密功能开关
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s ModifyTerminalSettingsRequestPasswordlessLoginConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTerminalSettingsRequestPasswordlessLoginConfig) GoString() string {
	return s.String()
}

func (s *ModifyTerminalSettingsRequestPasswordlessLoginConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyTerminalSettingsRequestPasswordlessLoginConfig) SetEnabled(v bool) *ModifyTerminalSettingsRequestPasswordlessLoginConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyTerminalSettingsRequestPasswordlessLoginConfig) Validate() error {
	return dara.Validate(s)
}

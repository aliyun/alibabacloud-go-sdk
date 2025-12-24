// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTerminalSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordlessLoginConfigShrink(v string) *ModifyTerminalSettingsShrinkRequest
	GetPasswordlessLoginConfigShrink() *string
}

type ModifyTerminalSettingsShrinkRequest struct {
	// 免密登录配置
	//
	// This parameter is required.
	PasswordlessLoginConfigShrink *string `json:"PasswordlessLoginConfig,omitempty" xml:"PasswordlessLoginConfig,omitempty"`
}

func (s ModifyTerminalSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTerminalSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyTerminalSettingsShrinkRequest) GetPasswordlessLoginConfigShrink() *string {
	return s.PasswordlessLoginConfigShrink
}

func (s *ModifyTerminalSettingsShrinkRequest) SetPasswordlessLoginConfigShrink(v string) *ModifyTerminalSettingsShrinkRequest {
	s.PasswordlessLoginConfigShrink = &v
	return s
}

func (s *ModifyTerminalSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

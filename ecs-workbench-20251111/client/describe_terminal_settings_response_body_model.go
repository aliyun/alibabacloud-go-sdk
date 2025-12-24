// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTerminalSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordlessLoginConfig(v *DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig) *DescribeTerminalSettingsResponseBody
	GetPasswordlessLoginConfig() *DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig
	SetRequestId(v string) *DescribeTerminalSettingsResponseBody
	GetRequestId() *string
}

type DescribeTerminalSettingsResponseBody struct {
	PasswordlessLoginConfig *DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig `json:"PasswordlessLoginConfig,omitempty" xml:"PasswordlessLoginConfig,omitempty" type:"Struct"`
	// 请求ID
	//
	// example:
	//
	// 47348885-C929-489A-93D7-B2E99D50D77B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTerminalSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSettingsResponseBody) GetPasswordlessLoginConfig() *DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig {
	return s.PasswordlessLoginConfig
}

func (s *DescribeTerminalSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTerminalSettingsResponseBody) SetPasswordlessLoginConfig(v *DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig) *DescribeTerminalSettingsResponseBody {
	s.PasswordlessLoginConfig = v
	return s
}

func (s *DescribeTerminalSettingsResponseBody) SetRequestId(v string) *DescribeTerminalSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTerminalSettingsResponseBody) Validate() error {
	if s.PasswordlessLoginConfig != nil {
		if err := s.PasswordlessLoginConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig struct {
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig) SetEnabled(v bool) *DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig {
	s.Enabled = &v
	return s
}

func (s *DescribeTerminalSettingsResponseBodyPasswordlessLoginConfig) Validate() error {
	return dara.Validate(s)
}

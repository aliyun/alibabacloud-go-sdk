// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGatewayLogConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSlsConfig(v *GatewayLogConfigSlsConfig) *GatewayLogConfig
	GetSlsConfig() *GatewayLogConfigSlsConfig
}

type GatewayLogConfig struct {
	SlsConfig *GatewayLogConfigSlsConfig `json:"slsConfig,omitempty" xml:"slsConfig,omitempty" type:"Struct"`
}

func (s GatewayLogConfig) String() string {
	return dara.Prettify(s)
}

func (s GatewayLogConfig) GoString() string {
	return s.String()
}

func (s *GatewayLogConfig) GetSlsConfig() *GatewayLogConfigSlsConfig {
	return s.SlsConfig
}

func (s *GatewayLogConfig) SetSlsConfig(v *GatewayLogConfigSlsConfig) *GatewayLogConfig {
	s.SlsConfig = v
	return s
}

func (s *GatewayLogConfig) Validate() error {
	return dara.Validate(s)
}

type GatewayLogConfigSlsConfig struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s GatewayLogConfigSlsConfig) String() string {
	return dara.Prettify(s)
}

func (s GatewayLogConfigSlsConfig) GoString() string {
	return s.String()
}

func (s *GatewayLogConfigSlsConfig) GetEnable() *bool {
	return s.Enable
}

func (s *GatewayLogConfigSlsConfig) SetEnable(v bool) *GatewayLogConfigSlsConfig {
	s.Enable = &v
	return s
}

func (s *GatewayLogConfigSlsConfig) Validate() error {
	return dara.Validate(s)
}

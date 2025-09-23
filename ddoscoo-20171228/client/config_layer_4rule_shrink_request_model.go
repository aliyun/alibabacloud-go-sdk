// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *ConfigLayer4RuleShrinkRequest
	GetListeners() *string
	SetProxyEnable(v int64) *ConfigLayer4RuleShrinkRequest
	GetProxyEnable() *int64
	SetUsTimeoutShrink(v string) *ConfigLayer4RuleShrinkRequest
	GetUsTimeoutShrink() *string
}

type ConfigLayer4RuleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"xxxxxx-xxxxxx-xxxxxx-xxxxxxx","Protocol":"tcp","FrontendPort":80,"BackendPort":5,"RealServers":"1.1.1.1","2.2.2.2"}]
	Listeners       *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	ProxyEnable     *int64  `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
	UsTimeoutShrink *string `json:"UsTimeout,omitempty" xml:"UsTimeout,omitempty"`
}

func (s ConfigLayer4RuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleShrinkRequest) GetListeners() *string {
	return s.Listeners
}

func (s *ConfigLayer4RuleShrinkRequest) GetProxyEnable() *int64 {
	return s.ProxyEnable
}

func (s *ConfigLayer4RuleShrinkRequest) GetUsTimeoutShrink() *string {
	return s.UsTimeoutShrink
}

func (s *ConfigLayer4RuleShrinkRequest) SetListeners(v string) *ConfigLayer4RuleShrinkRequest {
	s.Listeners = &v
	return s
}

func (s *ConfigLayer4RuleShrinkRequest) SetProxyEnable(v int64) *ConfigLayer4RuleShrinkRequest {
	s.ProxyEnable = &v
	return s
}

func (s *ConfigLayer4RuleShrinkRequest) SetUsTimeoutShrink(v string) *ConfigLayer4RuleShrinkRequest {
	s.UsTimeoutShrink = &v
	return s
}

func (s *ConfigLayer4RuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *ConfigLayer4RuleRequest
	GetListeners() *string
	SetProxyEnable(v int64) *ConfigLayer4RuleRequest
	GetProxyEnable() *int64
	SetUsTimeout(v *ConfigLayer4RuleRequestUsTimeout) *ConfigLayer4RuleRequest
	GetUsTimeout() *ConfigLayer4RuleRequestUsTimeout
}

type ConfigLayer4RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"xxxxxx-xxxxxx-xxxxxx-xxxxxxx","Protocol":"tcp","FrontendPort":80,"BackendPort":5,"RealServers":"1.1.1.1","2.2.2.2"}]
	Listeners   *string                           `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	ProxyEnable *int64                            `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
	UsTimeout   *ConfigLayer4RuleRequestUsTimeout `json:"UsTimeout,omitempty" xml:"UsTimeout,omitempty" type:"Struct"`
}

func (s ConfigLayer4RuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleRequest) GetListeners() *string {
	return s.Listeners
}

func (s *ConfigLayer4RuleRequest) GetProxyEnable() *int64 {
	return s.ProxyEnable
}

func (s *ConfigLayer4RuleRequest) GetUsTimeout() *ConfigLayer4RuleRequestUsTimeout {
	return s.UsTimeout
}

func (s *ConfigLayer4RuleRequest) SetListeners(v string) *ConfigLayer4RuleRequest {
	s.Listeners = &v
	return s
}

func (s *ConfigLayer4RuleRequest) SetProxyEnable(v int64) *ConfigLayer4RuleRequest {
	s.ProxyEnable = &v
	return s
}

func (s *ConfigLayer4RuleRequest) SetUsTimeout(v *ConfigLayer4RuleRequestUsTimeout) *ConfigLayer4RuleRequest {
	s.UsTimeout = v
	return s
}

func (s *ConfigLayer4RuleRequest) Validate() error {
	return dara.Validate(s)
}

type ConfigLayer4RuleRequestUsTimeout struct {
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	RsTimeout      *int64 `json:"RsTimeout,omitempty" xml:"RsTimeout,omitempty"`
}

func (s ConfigLayer4RuleRequestUsTimeout) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RuleRequestUsTimeout) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RuleRequestUsTimeout) GetConnectTimeout() *int64 {
	return s.ConnectTimeout
}

func (s *ConfigLayer4RuleRequestUsTimeout) GetRsTimeout() *int64 {
	return s.RsTimeout
}

func (s *ConfigLayer4RuleRequestUsTimeout) SetConnectTimeout(v int64) *ConfigLayer4RuleRequestUsTimeout {
	s.ConnectTimeout = &v
	return s
}

func (s *ConfigLayer4RuleRequestUsTimeout) SetRsTimeout(v int64) *ConfigLayer4RuleRequestUsTimeout {
	s.RsTimeout = &v
	return s
}

func (s *ConfigLayer4RuleRequestUsTimeout) Validate() error {
	return dara.Validate(s)
}

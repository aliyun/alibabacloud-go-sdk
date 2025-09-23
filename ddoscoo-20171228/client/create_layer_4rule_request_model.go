// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayer4RuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *CreateLayer4RuleRequest
	GetListeners() *string
	SetProxyEnable(v int32) *CreateLayer4RuleRequest
	GetProxyEnable() *int32
	SetUsTimeout(v *CreateLayer4RuleRequestUsTimeout) *CreateLayer4RuleRequest
	GetUsTimeout() *CreateLayer4RuleRequestUsTimeout
}

type CreateLayer4RuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"xxxxxx-xxxxxx-xxxxxx-xxxxxxx","Protocol":"tcp","FrontendPort":80,"BackendPort":5,"RealServers":"1.1.1.1","2.2.2.2"}]
	Listeners   *string                           `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	ProxyEnable *int32                            `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
	UsTimeout   *CreateLayer4RuleRequestUsTimeout `json:"UsTimeout,omitempty" xml:"UsTimeout,omitempty" type:"Struct"`
}

func (s CreateLayer4RuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLayer4RuleRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleRequest) GetListeners() *string {
	return s.Listeners
}

func (s *CreateLayer4RuleRequest) GetProxyEnable() *int32 {
	return s.ProxyEnable
}

func (s *CreateLayer4RuleRequest) GetUsTimeout() *CreateLayer4RuleRequestUsTimeout {
	return s.UsTimeout
}

func (s *CreateLayer4RuleRequest) SetListeners(v string) *CreateLayer4RuleRequest {
	s.Listeners = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetProxyEnable(v int32) *CreateLayer4RuleRequest {
	s.ProxyEnable = &v
	return s
}

func (s *CreateLayer4RuleRequest) SetUsTimeout(v *CreateLayer4RuleRequestUsTimeout) *CreateLayer4RuleRequest {
	s.UsTimeout = v
	return s
}

func (s *CreateLayer4RuleRequest) Validate() error {
	return dara.Validate(s)
}

type CreateLayer4RuleRequestUsTimeout struct {
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	RsTimeout      *int64 `json:"RsTimeout,omitempty" xml:"RsTimeout,omitempty"`
}

func (s CreateLayer4RuleRequestUsTimeout) String() string {
	return dara.Prettify(s)
}

func (s CreateLayer4RuleRequestUsTimeout) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleRequestUsTimeout) GetConnectTimeout() *int64 {
	return s.ConnectTimeout
}

func (s *CreateLayer4RuleRequestUsTimeout) GetRsTimeout() *int64 {
	return s.RsTimeout
}

func (s *CreateLayer4RuleRequestUsTimeout) SetConnectTimeout(v int64) *CreateLayer4RuleRequestUsTimeout {
	s.ConnectTimeout = &v
	return s
}

func (s *CreateLayer4RuleRequestUsTimeout) SetRsTimeout(v int64) *CreateLayer4RuleRequestUsTimeout {
	s.RsTimeout = &v
	return s
}

func (s *CreateLayer4RuleRequestUsTimeout) Validate() error {
	return dara.Validate(s)
}

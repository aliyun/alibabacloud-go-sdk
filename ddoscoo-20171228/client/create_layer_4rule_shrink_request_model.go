// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayer4RuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *CreateLayer4RuleShrinkRequest
	GetListeners() *string
	SetProxyEnable(v int32) *CreateLayer4RuleShrinkRequest
	GetProxyEnable() *int32
	SetUsTimeoutShrink(v string) *CreateLayer4RuleShrinkRequest
	GetUsTimeoutShrink() *string
}

type CreateLayer4RuleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"xxxxxx-xxxxxx-xxxxxx-xxxxxxx","Protocol":"tcp","FrontendPort":80,"BackendPort":5,"RealServers":"1.1.1.1","2.2.2.2"}]
	Listeners       *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
	ProxyEnable     *int32  `json:"ProxyEnable,omitempty" xml:"ProxyEnable,omitempty"`
	UsTimeoutShrink *string `json:"UsTimeout,omitempty" xml:"UsTimeout,omitempty"`
}

func (s CreateLayer4RuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLayer4RuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLayer4RuleShrinkRequest) GetListeners() *string {
	return s.Listeners
}

func (s *CreateLayer4RuleShrinkRequest) GetProxyEnable() *int32 {
	return s.ProxyEnable
}

func (s *CreateLayer4RuleShrinkRequest) GetUsTimeoutShrink() *string {
	return s.UsTimeoutShrink
}

func (s *CreateLayer4RuleShrinkRequest) SetListeners(v string) *CreateLayer4RuleShrinkRequest {
	s.Listeners = &v
	return s
}

func (s *CreateLayer4RuleShrinkRequest) SetProxyEnable(v int32) *CreateLayer4RuleShrinkRequest {
	s.ProxyEnable = &v
	return s
}

func (s *CreateLayer4RuleShrinkRequest) SetUsTimeoutShrink(v string) *CreateLayer4RuleShrinkRequest {
	s.UsTimeoutShrink = &v
	return s
}

func (s *CreateLayer4RuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}

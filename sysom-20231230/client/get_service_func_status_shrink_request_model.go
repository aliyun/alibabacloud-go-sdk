// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceFuncStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *GetServiceFuncStatusShrinkRequest
	GetChannel() *string
	SetParamsShrink(v string) *GetServiceFuncStatusShrinkRequest
	GetParamsShrink() *string
	SetServiceName(v string) *GetServiceFuncStatusShrinkRequest
	GetServiceName() *string
}

type GetServiceFuncStatusShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// This parameter is required.
	ParamsShrink *string `json:"params,omitempty" xml:"params,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// livetrace
	ServiceName *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
}

func (s GetServiceFuncStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceFuncStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusShrinkRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetServiceFuncStatusShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *GetServiceFuncStatusShrinkRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceFuncStatusShrinkRequest) SetChannel(v string) *GetServiceFuncStatusShrinkRequest {
	s.Channel = &v
	return s
}

func (s *GetServiceFuncStatusShrinkRequest) SetParamsShrink(v string) *GetServiceFuncStatusShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *GetServiceFuncStatusShrinkRequest) SetServiceName(v string) *GetServiceFuncStatusShrinkRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceFuncStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}

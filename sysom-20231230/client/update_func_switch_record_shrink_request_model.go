// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFuncSwitchRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *UpdateFuncSwitchRecordShrinkRequest
	GetChannel() *string
	SetParamsShrink(v string) *UpdateFuncSwitchRecordShrinkRequest
	GetParamsShrink() *string
	SetServiceName(v string) *UpdateFuncSwitchRecordShrinkRequest
	GetServiceName() *string
}

type UpdateFuncSwitchRecordShrinkRequest struct {
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

func (s UpdateFuncSwitchRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFuncSwitchRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordShrinkRequest) GetChannel() *string {
	return s.Channel
}

func (s *UpdateFuncSwitchRecordShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *UpdateFuncSwitchRecordShrinkRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *UpdateFuncSwitchRecordShrinkRequest) SetChannel(v string) *UpdateFuncSwitchRecordShrinkRequest {
	s.Channel = &v
	return s
}

func (s *UpdateFuncSwitchRecordShrinkRequest) SetParamsShrink(v string) *UpdateFuncSwitchRecordShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *UpdateFuncSwitchRecordShrinkRequest) SetServiceName(v string) *UpdateFuncSwitchRecordShrinkRequest {
	s.ServiceName = &v
	return s
}

func (s *UpdateFuncSwitchRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}

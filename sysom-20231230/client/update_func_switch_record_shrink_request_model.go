// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFuncSwitchRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *UpdateFuncSwitchRecordShrinkRequest
	GetXDebugId() *string
	SetChannel(v string) *UpdateFuncSwitchRecordShrinkRequest
	GetChannel() *string
	SetParamsShrink(v string) *UpdateFuncSwitchRecordShrinkRequest
	GetParamsShrink() *string
	SetServiceName(v string) *UpdateFuncSwitchRecordShrinkRequest
	GetServiceName() *string
	SetXSysomInvokeSource(v string) *UpdateFuncSwitchRecordShrinkRequest
	GetXSysomInvokeSource() *string
}

type UpdateFuncSwitchRecordShrinkRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The diagnostic channel. Currently, this parameter is fixed to the ECS channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The diagnostic parameters. Different types of diagnostics require different diagnostic parameters. You can use this field to filter records whose parameters match specified values.
	//
	// This parameter is required.
	ParamsShrink *string `json:"params,omitempty" xml:"params,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// livetrace
	ServiceName        *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s UpdateFuncSwitchRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFuncSwitchRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateFuncSwitchRecordShrinkRequest) GetXDebugId() *string {
	return s.XDebugId
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

func (s *UpdateFuncSwitchRecordShrinkRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *UpdateFuncSwitchRecordShrinkRequest) SetXDebugId(v string) *UpdateFuncSwitchRecordShrinkRequest {
	s.XDebugId = &v
	return s
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

func (s *UpdateFuncSwitchRecordShrinkRequest) SetXSysomInvokeSource(v string) *UpdateFuncSwitchRecordShrinkRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *UpdateFuncSwitchRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}

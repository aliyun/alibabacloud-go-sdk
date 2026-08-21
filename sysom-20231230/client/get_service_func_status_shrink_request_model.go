// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceFuncStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *GetServiceFuncStatusShrinkRequest
	GetXDebugId() *string
	SetChannel(v string) *GetServiceFuncStatusShrinkRequest
	GetChannel() *string
	SetParamsShrink(v string) *GetServiceFuncStatusShrinkRequest
	GetParamsShrink() *string
	SetServiceName(v string) *GetServiceFuncStatusShrinkRequest
	GetServiceName() *string
	SetXSysomInvokeSource(v string) *GetServiceFuncStatusShrinkRequest
	GetXSysomInvokeSource() *string
}

type GetServiceFuncStatusShrinkRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The channel name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The diagnostic parameters. Different types of diagnostics require different diagnostic parameters. You can use this field to filter records whose parameters match the specified values.
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

func (s GetServiceFuncStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceFuncStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceFuncStatusShrinkRequest) GetXDebugId() *string {
	return s.XDebugId
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

func (s *GetServiceFuncStatusShrinkRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *GetServiceFuncStatusShrinkRequest) SetXDebugId(v string) *GetServiceFuncStatusShrinkRequest {
	s.XDebugId = &v
	return s
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

func (s *GetServiceFuncStatusShrinkRequest) SetXSysomInvokeSource(v string) *GetServiceFuncStatusShrinkRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *GetServiceFuncStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}

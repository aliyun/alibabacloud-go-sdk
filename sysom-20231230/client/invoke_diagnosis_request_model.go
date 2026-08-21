// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeDiagnosisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *InvokeDiagnosisRequest
	GetXDebugId() *string
	SetChannel(v string) *InvokeDiagnosisRequest
	GetChannel() *string
	SetParams(v string) *InvokeDiagnosisRequest
	GetParams() *string
	SetServiceName(v string) *InvokeDiagnosisRequest
	GetServiceName() *string
	SetXSysomInvokeSource(v string) *InvokeDiagnosisRequest
	GetXSysomInvokeSource() *string
}

type InvokeDiagnosisRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The diagnosis channel (currently fixed to the ECS channel).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The diagnosis parameters. Different diagnosis types require different parameters. Refer to the supplementary request parameter descriptions below for the parameters required by each diagnosis type.
	//
	// 	Notice: Pass a JSON-formatted string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "instance": "i-wz9gdv7qmdhusamc4dl01",
	//
	//     "uid": "xxxxxxxxxxxxxx",
	//
	//     "region": "cn-shenzhen"
	//
	// }
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// The diagnosis type. This parameter distinguishes between different types of diagnostics.
	//
	// This parameter is required.
	//
	// example:
	//
	// memgraph
	ServiceName        *string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s InvokeDiagnosisRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeDiagnosisRequest) GoString() string {
	return s.String()
}

func (s *InvokeDiagnosisRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *InvokeDiagnosisRequest) GetChannel() *string {
	return s.Channel
}

func (s *InvokeDiagnosisRequest) GetParams() *string {
	return s.Params
}

func (s *InvokeDiagnosisRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *InvokeDiagnosisRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *InvokeDiagnosisRequest) SetXDebugId(v string) *InvokeDiagnosisRequest {
	s.XDebugId = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetChannel(v string) *InvokeDiagnosisRequest {
	s.Channel = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetParams(v string) *InvokeDiagnosisRequest {
	s.Params = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetServiceName(v string) *InvokeDiagnosisRequest {
	s.ServiceName = &v
	return s
}

func (s *InvokeDiagnosisRequest) SetXSysomInvokeSource(v string) *InvokeDiagnosisRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *InvokeDiagnosisRequest) Validate() error {
	return dara.Validate(s)
}

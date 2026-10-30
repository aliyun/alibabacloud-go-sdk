// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkAgentRunParameters interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeName(v string) *SinkAgentRunParameters
	GetAgentRuntimeName() *string
	SetBody(v *SinkAgentRunParametersBody) *SinkAgentRunParameters
	GetBody() *SinkAgentRunParametersBody
	SetEndpointName(v string) *SinkAgentRunParameters
	GetEndpointName() *string
	SetRoleName(v string) *SinkAgentRunParameters
	GetRoleName() *string
	SetTimeout(v string) *SinkAgentRunParameters
	GetTimeout() *string
}

type SinkAgentRunParameters struct {
	// The name of the AgentRun agent runtime. You can obtain it from the Alibaba Cloud AgentRun console.
	//
	// example:
	//
	// agent-demo-example
	AgentRuntimeName *string `json:"AgentRuntimeName,omitempty" xml:"AgentRuntimeName,omitempty"`
	// The triplet configuration for the Body parameter template. It defines the data transformation method and supports two modes: ORIGINAL (pass-through) and TEMPLATE (template replacement).
	Body *SinkAgentRunParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The endpoint name of the AgentRun agent runtime, which is used to construct the invocation URL.
	//
	// example:
	//
	// default
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The name of the AgentRun agent role.
	//
	// example:
	//
	// example-agent-role
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The HTTP request timeout period in milliseconds. Default value: 30000.
	//
	// example:
	//
	// 30000
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s SinkAgentRunParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkAgentRunParameters) GoString() string {
	return s.String()
}

func (s *SinkAgentRunParameters) GetAgentRuntimeName() *string {
	return s.AgentRuntimeName
}

func (s *SinkAgentRunParameters) GetBody() *SinkAgentRunParametersBody {
	return s.Body
}

func (s *SinkAgentRunParameters) GetEndpointName() *string {
	return s.EndpointName
}

func (s *SinkAgentRunParameters) GetRoleName() *string {
	return s.RoleName
}

func (s *SinkAgentRunParameters) GetTimeout() *string {
	return s.Timeout
}

func (s *SinkAgentRunParameters) SetAgentRuntimeName(v string) *SinkAgentRunParameters {
	s.AgentRuntimeName = &v
	return s
}

func (s *SinkAgentRunParameters) SetBody(v *SinkAgentRunParametersBody) *SinkAgentRunParameters {
	s.Body = v
	return s
}

func (s *SinkAgentRunParameters) SetEndpointName(v string) *SinkAgentRunParameters {
	s.EndpointName = &v
	return s
}

func (s *SinkAgentRunParameters) SetRoleName(v string) *SinkAgentRunParameters {
	s.RoleName = &v
	return s
}

func (s *SinkAgentRunParameters) SetTimeout(v string) *SinkAgentRunParameters {
	s.Timeout = &v
	return s
}

func (s *SinkAgentRunParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SinkAgentRunParametersBody struct {
	// The transformation format. Default value: CONSTANT. Valid values: ORIGINAL, TEMPLATE. Example: `ORIGINAL`.
	//
	// example:
	//
	// ORIGINAL
	Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
	// The template pattern. This parameter is required only when Form is set to TEMPLATE. It supports ${variable} placeholder replacement. Example: `The value of ${key} is ${value}!`.
	//
	// example:
	//
	// The value of ${key} is ${value}!
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// The parameter value. In ORIGINAL mode, the value is $. In TEMPLATE mode, the value is a key-value mapping in JSON format, where the value is a JSONPath expression or a constant. Example: `$`.
	//
	// example:
	//
	// $
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkAgentRunParametersBody) String() string {
	return dara.Prettify(s)
}

func (s SinkAgentRunParametersBody) GoString() string {
	return s.String()
}

func (s *SinkAgentRunParametersBody) GetForm() *string {
	return s.Form
}

func (s *SinkAgentRunParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *SinkAgentRunParametersBody) GetValue() *string {
	return s.Value
}

func (s *SinkAgentRunParametersBody) SetForm(v string) *SinkAgentRunParametersBody {
	s.Form = &v
	return s
}

func (s *SinkAgentRunParametersBody) SetTemplate(v string) *SinkAgentRunParametersBody {
	s.Template = &v
	return s
}

func (s *SinkAgentRunParametersBody) SetValue(v string) *SinkAgentRunParametersBody {
	s.Value = &v
	return s
}

func (s *SinkAgentRunParametersBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunChatResultGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInferenceParameters(v map[string]interface{}) *RunChatResultGenerationRequest
	GetInferenceParameters() map[string]interface{}
	SetMessages(v []*RunChatResultGenerationRequestMessages) *RunChatResultGenerationRequest
	GetMessages() []*RunChatResultGenerationRequestMessages
	SetModelId(v string) *RunChatResultGenerationRequest
	GetModelId() *string
	SetSessionId(v string) *RunChatResultGenerationRequest
	GetSessionId() *string
	SetStream(v bool) *RunChatResultGenerationRequest
	GetStream() *bool
	SetTools(v []*RunChatResultGenerationRequestTools) *RunChatResultGenerationRequest
	GetTools() []*RunChatResultGenerationRequestTools
}

type RunChatResultGenerationRequest struct {
	// example:
	//
	// {"topP": 0.8}
	InferenceParameters map[string]interface{} `json:"inferenceParameters,omitempty" xml:"inferenceParameters,omitempty"`
	// This parameter is required.
	Messages []*RunChatResultGenerationRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 237645726354
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// false
	Stream *bool                                  `json:"stream,omitempty" xml:"stream,omitempty"`
	Tools  []*RunChatResultGenerationRequestTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s RunChatResultGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequest) GetInferenceParameters() map[string]interface{} {
	return s.InferenceParameters
}

func (s *RunChatResultGenerationRequest) GetMessages() []*RunChatResultGenerationRequestMessages {
	return s.Messages
}

func (s *RunChatResultGenerationRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunChatResultGenerationRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunChatResultGenerationRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunChatResultGenerationRequest) GetTools() []*RunChatResultGenerationRequestTools {
	return s.Tools
}

func (s *RunChatResultGenerationRequest) SetInferenceParameters(v map[string]interface{}) *RunChatResultGenerationRequest {
	s.InferenceParameters = v
	return s
}

func (s *RunChatResultGenerationRequest) SetMessages(v []*RunChatResultGenerationRequestMessages) *RunChatResultGenerationRequest {
	s.Messages = v
	return s
}

func (s *RunChatResultGenerationRequest) SetModelId(v string) *RunChatResultGenerationRequest {
	s.ModelId = &v
	return s
}

func (s *RunChatResultGenerationRequest) SetSessionId(v string) *RunChatResultGenerationRequest {
	s.SessionId = &v
	return s
}

func (s *RunChatResultGenerationRequest) SetStream(v bool) *RunChatResultGenerationRequest {
	s.Stream = &v
	return s
}

func (s *RunChatResultGenerationRequest) SetTools(v []*RunChatResultGenerationRequestTools) *RunChatResultGenerationRequest {
	s.Tools = v
	return s
}

func (s *RunChatResultGenerationRequest) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunChatResultGenerationRequestMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunChatResultGenerationRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationRequestMessages) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestMessages) GetContent() *string {
	return s.Content
}

func (s *RunChatResultGenerationRequestMessages) GetRole() *string {
	return s.Role
}

func (s *RunChatResultGenerationRequestMessages) SetContent(v string) *RunChatResultGenerationRequestMessages {
	s.Content = &v
	return s
}

func (s *RunChatResultGenerationRequestMessages) SetRole(v string) *RunChatResultGenerationRequestMessages {
	s.Role = &v
	return s
}

func (s *RunChatResultGenerationRequestMessages) Validate() error {
	return dara.Validate(s)
}

type RunChatResultGenerationRequestTools struct {
	Function *RunChatResultGenerationRequestToolsFunction `json:"function,omitempty" xml:"function,omitempty" type:"Struct"`
	// example:
	//
	// function
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RunChatResultGenerationRequestTools) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationRequestTools) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestTools) GetFunction() *RunChatResultGenerationRequestToolsFunction {
	return s.Function
}

func (s *RunChatResultGenerationRequestTools) GetType() *string {
	return s.Type
}

func (s *RunChatResultGenerationRequestTools) SetFunction(v *RunChatResultGenerationRequestToolsFunction) *RunChatResultGenerationRequestTools {
	s.Function = v
	return s
}

func (s *RunChatResultGenerationRequestTools) SetType(v string) *RunChatResultGenerationRequestTools {
	s.Type = &v
	return s
}

func (s *RunChatResultGenerationRequestTools) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunChatResultGenerationRequestToolsFunction struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// get_time
	Name       *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	Parameters *RunChatResultGenerationRequestToolsFunctionParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	Required   []*string                                              `json:"required,omitempty" xml:"required,omitempty" type:"Repeated"`
}

func (s RunChatResultGenerationRequestToolsFunction) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationRequestToolsFunction) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestToolsFunction) GetDescription() *string {
	return s.Description
}

func (s *RunChatResultGenerationRequestToolsFunction) GetName() *string {
	return s.Name
}

func (s *RunChatResultGenerationRequestToolsFunction) GetParameters() *RunChatResultGenerationRequestToolsFunctionParameters {
	return s.Parameters
}

func (s *RunChatResultGenerationRequestToolsFunction) GetRequired() []*string {
	return s.Required
}

func (s *RunChatResultGenerationRequestToolsFunction) SetDescription(v string) *RunChatResultGenerationRequestToolsFunction {
	s.Description = &v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) SetName(v string) *RunChatResultGenerationRequestToolsFunction {
	s.Name = &v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) SetParameters(v *RunChatResultGenerationRequestToolsFunctionParameters) *RunChatResultGenerationRequestToolsFunction {
	s.Parameters = v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) SetRequired(v []*string) *RunChatResultGenerationRequestToolsFunction {
	s.Required = v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunction) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunChatResultGenerationRequestToolsFunctionParameters struct {
	// example:
	//
	// {
	//
	//                             "location": {
	//
	//                                 "type": "string",
	//
	//                                 "description": "The city and state, e.g. San Francisco, CA"
	//
	//                             },
	//
	//                             "unit": {
	//
	//                                 "type": "string",
	//
	//                                 "enum": [
	//
	//                                     "celsius",
	//
	//                                     "fahrenheit"
	//
	//                                 ]
	//
	//                             }
	//
	//                         }
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// example:
	//
	// object
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RunChatResultGenerationRequestToolsFunctionParameters) String() string {
	return dara.Prettify(s)
}

func (s RunChatResultGenerationRequestToolsFunctionParameters) GoString() string {
	return s.String()
}

func (s *RunChatResultGenerationRequestToolsFunctionParameters) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *RunChatResultGenerationRequestToolsFunctionParameters) GetType() *string {
	return s.Type
}

func (s *RunChatResultGenerationRequestToolsFunctionParameters) SetProperties(v map[string]interface{}) *RunChatResultGenerationRequestToolsFunctionParameters {
	s.Properties = v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunctionParameters) SetType(v string) *RunChatResultGenerationRequestToolsFunctionParameters {
	s.Type = &v
	return s
}

func (s *RunChatResultGenerationRequestToolsFunctionParameters) Validate() error {
	return dara.Validate(s)
}

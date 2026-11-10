// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v []*InvokeAgentRequestContext) *InvokeAgentRequest
	GetContext() []*InvokeAgentRequestContext
	SetForwardedProps(v *InvokeAgentRequestForwardedProps) *InvokeAgentRequest
	GetForwardedProps() *InvokeAgentRequestForwardedProps
	SetMessages(v []*InvokeAgentRequestMessages) *InvokeAgentRequest
	GetMessages() []*InvokeAgentRequestMessages
	SetParentRunId(v string) *InvokeAgentRequest
	GetParentRunId() *string
	SetResume(v []*InvokeAgentRequestResume) *InvokeAgentRequest
	GetResume() []*InvokeAgentRequestResume
	SetRunId(v string) *InvokeAgentRequest
	GetRunId() *string
	SetState(v interface{}) *InvokeAgentRequest
	GetState() interface{}
	SetThreadId(v string) *InvokeAgentRequest
	GetThreadId() *string
	SetTools(v []*InvokeAgentRequestTools) *InvokeAgentRequest
	GetTools() []*InvokeAgentRequestTools
}

type InvokeAgentRequest struct {
	// example:
	//
	// [{"description":"cluster","value":"c-1"}]
	Context []*InvokeAgentRequestContext `json:"context,omitempty" xml:"context,omitempty" type:"Repeated"`
	// example:
	//
	// {"enableThinking":true}
	ForwardedProps *InvokeAgentRequestForwardedProps `json:"forwardedProps,omitempty" xml:"forwardedProps,omitempty" type:"Struct"`
	// example:
	//
	// [{"id":"m-1","role":"user","content":"你好"}]
	Messages []*InvokeAgentRequestMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// example:
	//
	// r-0
	ParentRunId *string `json:"parentRunId,omitempty" xml:"parentRunId,omitempty"`
	// example:
	//
	// [{"interruptId":"i-1","status":"resolved","payload":{"decision":"yes"}}]
	Resume []*InvokeAgentRequestResume `json:"resume,omitempty" xml:"resume,omitempty" type:"Repeated"`
	// example:
	//
	// run-1
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// {}
	State interface{} `json:"state,omitempty" xml:"state,omitempty"`
	// example:
	//
	// abcdef0123456789abcdef0123456789
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// []
	Tools []*InvokeAgentRequestTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s InvokeAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentRequest) GoString() string {
	return s.String()
}

func (s *InvokeAgentRequest) GetContext() []*InvokeAgentRequestContext {
	return s.Context
}

func (s *InvokeAgentRequest) GetForwardedProps() *InvokeAgentRequestForwardedProps {
	return s.ForwardedProps
}

func (s *InvokeAgentRequest) GetMessages() []*InvokeAgentRequestMessages {
	return s.Messages
}

func (s *InvokeAgentRequest) GetParentRunId() *string {
	return s.ParentRunId
}

func (s *InvokeAgentRequest) GetResume() []*InvokeAgentRequestResume {
	return s.Resume
}

func (s *InvokeAgentRequest) GetRunId() *string {
	return s.RunId
}

func (s *InvokeAgentRequest) GetState() interface{} {
	return s.State
}

func (s *InvokeAgentRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *InvokeAgentRequest) GetTools() []*InvokeAgentRequestTools {
	return s.Tools
}

func (s *InvokeAgentRequest) SetContext(v []*InvokeAgentRequestContext) *InvokeAgentRequest {
	s.Context = v
	return s
}

func (s *InvokeAgentRequest) SetForwardedProps(v *InvokeAgentRequestForwardedProps) *InvokeAgentRequest {
	s.ForwardedProps = v
	return s
}

func (s *InvokeAgentRequest) SetMessages(v []*InvokeAgentRequestMessages) *InvokeAgentRequest {
	s.Messages = v
	return s
}

func (s *InvokeAgentRequest) SetParentRunId(v string) *InvokeAgentRequest {
	s.ParentRunId = &v
	return s
}

func (s *InvokeAgentRequest) SetResume(v []*InvokeAgentRequestResume) *InvokeAgentRequest {
	s.Resume = v
	return s
}

func (s *InvokeAgentRequest) SetRunId(v string) *InvokeAgentRequest {
	s.RunId = &v
	return s
}

func (s *InvokeAgentRequest) SetState(v interface{}) *InvokeAgentRequest {
	s.State = v
	return s
}

func (s *InvokeAgentRequest) SetThreadId(v string) *InvokeAgentRequest {
	s.ThreadId = &v
	return s
}

func (s *InvokeAgentRequest) SetTools(v []*InvokeAgentRequestTools) *InvokeAgentRequest {
	s.Tools = v
	return s
}

func (s *InvokeAgentRequest) Validate() error {
	if s.Context != nil {
		for _, item := range s.Context {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ForwardedProps != nil {
		if err := s.ForwardedProps.Validate(); err != nil {
			return err
		}
	}
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Resume != nil {
		for _, item := range s.Resume {
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

type InvokeAgentRequestContext struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s InvokeAgentRequestContext) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentRequestContext) GoString() string {
	return s.String()
}

func (s *InvokeAgentRequestContext) GetDescription() *string {
	return s.Description
}

func (s *InvokeAgentRequestContext) GetValue() *string {
	return s.Value
}

func (s *InvokeAgentRequestContext) SetDescription(v string) *InvokeAgentRequestContext {
	s.Description = &v
	return s
}

func (s *InvokeAgentRequestContext) SetValue(v string) *InvokeAgentRequestContext {
	s.Value = &v
	return s
}

func (s *InvokeAgentRequestContext) Validate() error {
	return dara.Validate(s)
}

type InvokeAgentRequestForwardedProps struct {
	// example:
	//
	// true
	EnableThinking *bool `json:"enableThinking,omitempty" xml:"enableThinking,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// user
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s InvokeAgentRequestForwardedProps) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentRequestForwardedProps) GoString() string {
	return s.String()
}

func (s *InvokeAgentRequestForwardedProps) GetEnableThinking() *bool {
	return s.EnableThinking
}

func (s *InvokeAgentRequestForwardedProps) GetLanguage() *string {
	return s.Language
}

func (s *InvokeAgentRequestForwardedProps) GetSource() *string {
	return s.Source
}

func (s *InvokeAgentRequestForwardedProps) SetEnableThinking(v bool) *InvokeAgentRequestForwardedProps {
	s.EnableThinking = &v
	return s
}

func (s *InvokeAgentRequestForwardedProps) SetLanguage(v string) *InvokeAgentRequestForwardedProps {
	s.Language = &v
	return s
}

func (s *InvokeAgentRequestForwardedProps) SetSource(v string) *InvokeAgentRequestForwardedProps {
	s.Source = &v
	return s
}

func (s *InvokeAgentRequestForwardedProps) Validate() error {
	return dara.Validate(s)
}

type InvokeAgentRequestMessages struct {
	// example:
	//
	// hello
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// ""
	EncryptedValue *string `json:"encryptedValue,omitempty" xml:"encryptedValue,omitempty"`
	// example:
	//
	// 64ea4d94937a4c5c84e2b3b814d14696
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// n-1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// 64ea4d94937a4c5c84e2b3b814d14696
	SubagentRunId *string `json:"subagentRunId,omitempty" xml:"subagentRunId,omitempty"`
	// example:
	//
	// 64ea4d94937a4c5c84e2b3b814d14696
	ToolCallId *string                                `json:"toolCallId,omitempty" xml:"toolCallId,omitempty"`
	ToolCalls  []*InvokeAgentRequestMessagesToolCalls `json:"toolCalls,omitempty" xml:"toolCalls,omitempty" type:"Repeated"`
}

func (s InvokeAgentRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentRequestMessages) GoString() string {
	return s.String()
}

func (s *InvokeAgentRequestMessages) GetContent() *string {
	return s.Content
}

func (s *InvokeAgentRequestMessages) GetEncryptedValue() *string {
	return s.EncryptedValue
}

func (s *InvokeAgentRequestMessages) GetId() *string {
	return s.Id
}

func (s *InvokeAgentRequestMessages) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *InvokeAgentRequestMessages) GetName() *string {
	return s.Name
}

func (s *InvokeAgentRequestMessages) GetRole() *string {
	return s.Role
}

func (s *InvokeAgentRequestMessages) GetSubagentRunId() *string {
	return s.SubagentRunId
}

func (s *InvokeAgentRequestMessages) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *InvokeAgentRequestMessages) GetToolCalls() []*InvokeAgentRequestMessagesToolCalls {
	return s.ToolCalls
}

func (s *InvokeAgentRequestMessages) SetContent(v string) *InvokeAgentRequestMessages {
	s.Content = &v
	return s
}

func (s *InvokeAgentRequestMessages) SetEncryptedValue(v string) *InvokeAgentRequestMessages {
	s.EncryptedValue = &v
	return s
}

func (s *InvokeAgentRequestMessages) SetId(v string) *InvokeAgentRequestMessages {
	s.Id = &v
	return s
}

func (s *InvokeAgentRequestMessages) SetMetadata(v map[string]interface{}) *InvokeAgentRequestMessages {
	s.Metadata = v
	return s
}

func (s *InvokeAgentRequestMessages) SetName(v string) *InvokeAgentRequestMessages {
	s.Name = &v
	return s
}

func (s *InvokeAgentRequestMessages) SetRole(v string) *InvokeAgentRequestMessages {
	s.Role = &v
	return s
}

func (s *InvokeAgentRequestMessages) SetSubagentRunId(v string) *InvokeAgentRequestMessages {
	s.SubagentRunId = &v
	return s
}

func (s *InvokeAgentRequestMessages) SetToolCallId(v string) *InvokeAgentRequestMessages {
	s.ToolCallId = &v
	return s
}

func (s *InvokeAgentRequestMessages) SetToolCalls(v []*InvokeAgentRequestMessagesToolCalls) *InvokeAgentRequestMessages {
	s.ToolCalls = v
	return s
}

func (s *InvokeAgentRequestMessages) Validate() error {
	if s.ToolCalls != nil {
		for _, item := range s.ToolCalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InvokeAgentRequestMessagesToolCalls struct {
	// example:
	//
	// ""
	EncryptedValue *string `json:"encryptedValue,omitempty" xml:"encryptedValue,omitempty"`
	// example:
	//
	// {"name": "xxx", "arguments": "xxx"}
	Function *InvokeAgentRequestMessagesToolCallsFunction `json:"function,omitempty" xml:"function,omitempty" type:"Struct"`
	// example:
	//
	// 64ea4d94937a4c5c84e2b3b814d14696
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// function
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvokeAgentRequestMessagesToolCalls) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentRequestMessagesToolCalls) GoString() string {
	return s.String()
}

func (s *InvokeAgentRequestMessagesToolCalls) GetEncryptedValue() *string {
	return s.EncryptedValue
}

func (s *InvokeAgentRequestMessagesToolCalls) GetFunction() *InvokeAgentRequestMessagesToolCallsFunction {
	return s.Function
}

func (s *InvokeAgentRequestMessagesToolCalls) GetId() *string {
	return s.Id
}

func (s *InvokeAgentRequestMessagesToolCalls) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *InvokeAgentRequestMessagesToolCalls) GetType() *string {
	return s.Type
}

func (s *InvokeAgentRequestMessagesToolCalls) SetEncryptedValue(v string) *InvokeAgentRequestMessagesToolCalls {
	s.EncryptedValue = &v
	return s
}

func (s *InvokeAgentRequestMessagesToolCalls) SetFunction(v *InvokeAgentRequestMessagesToolCallsFunction) *InvokeAgentRequestMessagesToolCalls {
	s.Function = v
	return s
}

func (s *InvokeAgentRequestMessagesToolCalls) SetId(v string) *InvokeAgentRequestMessagesToolCalls {
	s.Id = &v
	return s
}

func (s *InvokeAgentRequestMessagesToolCalls) SetMetadata(v map[string]interface{}) *InvokeAgentRequestMessagesToolCalls {
	s.Metadata = v
	return s
}

func (s *InvokeAgentRequestMessagesToolCalls) SetType(v string) *InvokeAgentRequestMessagesToolCalls {
	s.Type = &v
	return s
}

func (s *InvokeAgentRequestMessagesToolCalls) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokeAgentRequestMessagesToolCallsFunction struct {
	// example:
	//
	// {}
	Arguments *string `json:"arguments,omitempty" xml:"arguments,omitempty"`
	// example:
	//
	// navy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s InvokeAgentRequestMessagesToolCallsFunction) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentRequestMessagesToolCallsFunction) GoString() string {
	return s.String()
}

func (s *InvokeAgentRequestMessagesToolCallsFunction) GetArguments() *string {
	return s.Arguments
}

func (s *InvokeAgentRequestMessagesToolCallsFunction) GetName() *string {
	return s.Name
}

func (s *InvokeAgentRequestMessagesToolCallsFunction) SetArguments(v string) *InvokeAgentRequestMessagesToolCallsFunction {
	s.Arguments = &v
	return s
}

func (s *InvokeAgentRequestMessagesToolCallsFunction) SetName(v string) *InvokeAgentRequestMessagesToolCallsFunction {
	s.Name = &v
	return s
}

func (s *InvokeAgentRequestMessagesToolCallsFunction) Validate() error {
	return dara.Validate(s)
}

type InvokeAgentRequestResume struct {
	// example:
	//
	// i-1
	InterruptId *string `json:"interruptId,omitempty" xml:"interruptId,omitempty"`
	// example:
	//
	// {"decision":"yes"}
	Payload interface{} `json:"payload,omitempty" xml:"payload,omitempty"`
	// example:
	//
	// resolved
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s InvokeAgentRequestResume) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentRequestResume) GoString() string {
	return s.String()
}

func (s *InvokeAgentRequestResume) GetInterruptId() *string {
	return s.InterruptId
}

func (s *InvokeAgentRequestResume) GetPayload() interface{} {
	return s.Payload
}

func (s *InvokeAgentRequestResume) GetStatus() *string {
	return s.Status
}

func (s *InvokeAgentRequestResume) SetInterruptId(v string) *InvokeAgentRequestResume {
	s.InterruptId = &v
	return s
}

func (s *InvokeAgentRequestResume) SetPayload(v interface{}) *InvokeAgentRequestResume {
	s.Payload = v
	return s
}

func (s *InvokeAgentRequestResume) SetStatus(v string) *InvokeAgentRequestResume {
	s.Status = &v
	return s
}

func (s *InvokeAgentRequestResume) Validate() error {
	return dara.Validate(s)
}

type InvokeAgentRequestTools struct {
	// This parameter is required.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cluster-qa-icloud-edge-ohttp
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// {\\"region\\": \\"cn-shanghai\\", \\"vpcId\\": \\"vpc-uf6bnhlxne8enjurxxjv4\\", \\"functionName\\": \\"filemgr-cn-shanghai-334a349cd2\\", \\"mountPointsServerAddr\\": \\"334a349cd2-umc45.cn-shanghai.nas.aliyuncs.com\\"}
	Parameters interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s InvokeAgentRequestTools) String() string {
	return dara.Prettify(s)
}

func (s InvokeAgentRequestTools) GoString() string {
	return s.String()
}

func (s *InvokeAgentRequestTools) GetDescription() *string {
	return s.Description
}

func (s *InvokeAgentRequestTools) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *InvokeAgentRequestTools) GetName() *string {
	return s.Name
}

func (s *InvokeAgentRequestTools) GetParameters() interface{} {
	return s.Parameters
}

func (s *InvokeAgentRequestTools) SetDescription(v string) *InvokeAgentRequestTools {
	s.Description = &v
	return s
}

func (s *InvokeAgentRequestTools) SetMetadata(v map[string]interface{}) *InvokeAgentRequestTools {
	s.Metadata = v
	return s
}

func (s *InvokeAgentRequestTools) SetName(v string) *InvokeAgentRequestTools {
	s.Name = &v
	return s
}

func (s *InvokeAgentRequestTools) SetParameters(v interface{}) *InvokeAgentRequestTools {
	s.Parameters = v
	return s
}

func (s *InvokeAgentRequestTools) Validate() error {
	return dara.Validate(s)
}

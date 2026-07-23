// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAguiMessage interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *AguiMessage
	GetContent() *string
	SetId(v string) *AguiMessage
	GetId() *string
	SetMetadata(v *AguiMessageMetadata) *AguiMessage
	GetMetadata() *AguiMessageMetadata
	SetReasoning(v string) *AguiMessage
	GetReasoning() *string
	SetRole(v string) *AguiMessage
	GetRole() *string
	SetToolCallId(v string) *AguiMessage
	GetToolCallId() *string
	SetToolCalls(v []*AguiMessageToolCalls) *AguiMessage
	GetToolCalls() []*AguiMessageToolCalls
}

type AguiMessage struct {
	// The text content of the message.
	//
	// example:
	//
	// 根据您的问题，我将查询过去7天的事件量...
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The unique identifier of the message.
	//
	// example:
	//
	// msg_123456_a1b2c3d4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The extension metadata.
	Metadata  *AguiMessageMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	Reasoning *string              `json:"Reasoning,omitempty" xml:"Reasoning,omitempty"`
	// The role of the message.
	//
	// example:
	//
	// assistant
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The associated tool invocation ID.
	//
	// example:
	//
	// call_xxx
	ToolCallId *string `json:"ToolCallId,omitempty" xml:"ToolCallId,omitempty"`
	// The tool invocation list.
	ToolCalls []*AguiMessageToolCalls `json:"ToolCalls,omitempty" xml:"ToolCalls,omitempty" type:"Repeated"`
}

func (s AguiMessage) String() string {
	return dara.Prettify(s)
}

func (s AguiMessage) GoString() string {
	return s.String()
}

func (s *AguiMessage) GetContent() *string {
	return s.Content
}

func (s *AguiMessage) GetId() *string {
	return s.Id
}

func (s *AguiMessage) GetMetadata() *AguiMessageMetadata {
	return s.Metadata
}

func (s *AguiMessage) GetReasoning() *string {
	return s.Reasoning
}

func (s *AguiMessage) GetRole() *string {
	return s.Role
}

func (s *AguiMessage) GetToolCallId() *string {
	return s.ToolCallId
}

func (s *AguiMessage) GetToolCalls() []*AguiMessageToolCalls {
	return s.ToolCalls
}

func (s *AguiMessage) SetContent(v string) *AguiMessage {
	s.Content = &v
	return s
}

func (s *AguiMessage) SetId(v string) *AguiMessage {
	s.Id = &v
	return s
}

func (s *AguiMessage) SetMetadata(v *AguiMessageMetadata) *AguiMessage {
	s.Metadata = v
	return s
}

func (s *AguiMessage) SetReasoning(v string) *AguiMessage {
	s.Reasoning = &v
	return s
}

func (s *AguiMessage) SetRole(v string) *AguiMessage {
	s.Role = &v
	return s
}

func (s *AguiMessage) SetToolCallId(v string) *AguiMessage {
	s.ToolCallId = &v
	return s
}

func (s *AguiMessage) SetToolCalls(v []*AguiMessageToolCalls) *AguiMessage {
	s.ToolCalls = v
	return s
}

func (s *AguiMessage) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
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

type AguiMessageMetadata struct {
	// The extension data.
	Attachments *AguiMessageMetadataAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Struct"`
}

func (s AguiMessageMetadata) String() string {
	return dara.Prettify(s)
}

func (s AguiMessageMetadata) GoString() string {
	return s.String()
}

func (s *AguiMessageMetadata) GetAttachments() *AguiMessageMetadataAttachments {
	return s.Attachments
}

func (s *AguiMessageMetadata) SetAttachments(v *AguiMessageMetadataAttachments) *AguiMessageMetadata {
	s.Attachments = v
	return s
}

func (s *AguiMessageMetadata) Validate() error {
	if s.Attachments != nil {
		if err := s.Attachments.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AguiMessageMetadataAttachments struct {
	// The name of the extension data.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:12345:eventhouse/system-rocketmq/namespace/rmq-cn-xxx/table/order
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the extension data.
	//
	// example:
	//
	// inner-resource/event-table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AguiMessageMetadataAttachments) String() string {
	return dara.Prettify(s)
}

func (s AguiMessageMetadataAttachments) GoString() string {
	return s.String()
}

func (s *AguiMessageMetadataAttachments) GetName() *string {
	return s.Name
}

func (s *AguiMessageMetadataAttachments) GetType() *string {
	return s.Type
}

func (s *AguiMessageMetadataAttachments) SetName(v string) *AguiMessageMetadataAttachments {
	s.Name = &v
	return s
}

func (s *AguiMessageMetadataAttachments) SetType(v string) *AguiMessageMetadataAttachments {
	s.Type = &v
	return s
}

func (s *AguiMessageMetadataAttachments) Validate() error {
	return dara.Validate(s)
}

type AguiMessageToolCalls struct {
	// The tool calling function.
	Function *AguiMessageToolCallsFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	// The tool calling ID.
	//
	// example:
	//
	// call_xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The tool calling type.
	//
	// example:
	//
	// function
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AguiMessageToolCalls) String() string {
	return dara.Prettify(s)
}

func (s AguiMessageToolCalls) GoString() string {
	return s.String()
}

func (s *AguiMessageToolCalls) GetFunction() *AguiMessageToolCallsFunction {
	return s.Function
}

func (s *AguiMessageToolCalls) GetId() *string {
	return s.Id
}

func (s *AguiMessageToolCalls) GetType() *string {
	return s.Type
}

func (s *AguiMessageToolCalls) SetFunction(v *AguiMessageToolCallsFunction) *AguiMessageToolCalls {
	s.Function = v
	return s
}

func (s *AguiMessageToolCalls) SetId(v string) *AguiMessageToolCalls {
	s.Id = &v
	return s
}

func (s *AguiMessageToolCalls) SetType(v string) *AguiMessageToolCalls {
	s.Type = &v
	return s
}

func (s *AguiMessageToolCalls) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AguiMessageToolCallsFunction struct {
	// The arguments of the tool calling function.
	//
	// example:
	//
	// {}
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// The name of the tool calling function.
	//
	// example:
	//
	// discoverMetadata
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s AguiMessageToolCallsFunction) String() string {
	return dara.Prettify(s)
}

func (s AguiMessageToolCallsFunction) GoString() string {
	return s.String()
}

func (s *AguiMessageToolCallsFunction) GetArguments() *string {
	return s.Arguments
}

func (s *AguiMessageToolCallsFunction) GetName() *string {
	return s.Name
}

func (s *AguiMessageToolCallsFunction) SetArguments(v string) *AguiMessageToolCallsFunction {
	s.Arguments = &v
	return s
}

func (s *AguiMessageToolCallsFunction) SetName(v string) *AguiMessageToolCallsFunction {
	s.Name = &v
	return s
}

func (s *AguiMessageToolCallsFunction) Validate() error {
	return dara.Validate(s)
}

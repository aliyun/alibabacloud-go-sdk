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
	SetRole(v string) *AguiMessage
	GetRole() *string
	SetToolCallId(v string) *AguiMessage
	GetToolCallId() *string
	SetToolCalls(v []*AguiMessageToolCalls) *AguiMessage
	GetToolCalls() []*AguiMessageToolCalls
}

type AguiMessage struct {
	Content    *string                 `json:"Content,omitempty" xml:"Content,omitempty"`
	Id         *string                 `json:"Id,omitempty" xml:"Id,omitempty"`
	Metadata   *AguiMessageMetadata    `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	Role       *string                 `json:"Role,omitempty" xml:"Role,omitempty"`
	ToolCallId *string                 `json:"ToolCallId,omitempty" xml:"ToolCallId,omitempty"`
	ToolCalls  []*AguiMessageToolCalls `json:"ToolCalls,omitempty" xml:"ToolCalls,omitempty" type:"Repeated"`
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
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	Function *AguiMessageToolCallsFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	Id       *string                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Type     *string                       `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

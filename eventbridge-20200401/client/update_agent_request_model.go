// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateAgentRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateAgentRequest
	GetDescription() *string
	SetMetadata(v *UpdateAgentRequestMetadata) *UpdateAgentRequest
	GetMetadata() *UpdateAgentRequestMetadata
	SetName(v string) *UpdateAgentRequest
	GetName() *string
	SetPrompt(v string) *UpdateAgentRequest
	GetPrompt() *string
}

type UpdateAgentRequest struct {
	// example:
	//
	// TF-CreateRule-1652253755-aa33f762-7e99-4aee-bd27-d3370afa5625
	ClientToken *string                     `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string                     `json:"Description,omitempty" xml:"Description,omitempty"`
	Metadata    *UpdateAgentRequestMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// example:
	//
	// my-agent
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s UpdateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAgentRequest) GetMetadata() *UpdateAgentRequestMetadata {
	return s.Metadata
}

func (s *UpdateAgentRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAgentRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateAgentRequest) SetClientToken(v string) *UpdateAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentRequest) SetDescription(v string) *UpdateAgentRequest {
	s.Description = &v
	return s
}

func (s *UpdateAgentRequest) SetMetadata(v *UpdateAgentRequestMetadata) *UpdateAgentRequest {
	s.Metadata = v
	return s
}

func (s *UpdateAgentRequest) SetName(v string) *UpdateAgentRequest {
	s.Name = &v
	return s
}

func (s *UpdateAgentRequest) SetPrompt(v string) *UpdateAgentRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateAgentRequest) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentRequestMetadata struct {
	Attachments []*UpdateAgentRequestMetadataAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
}

func (s UpdateAgentRequestMetadata) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequestMetadata) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequestMetadata) GetAttachments() []*UpdateAgentRequestMetadataAttachments {
	return s.Attachments
}

func (s *UpdateAgentRequestMetadata) SetAttachments(v []*UpdateAgentRequestMetadataAttachments) *UpdateAgentRequestMetadata {
	s.Attachments = v
	return s
}

func (s *UpdateAgentRequestMetadata) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateAgentRequestMetadataAttachments struct {
	// example:
	//
	// acs:eventbridge:cn-hangzhou:12345:eventhouse/system-rocketmq/namespace/rmq-cn-xxx/table/order
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// example:
	//
	// inner-resource/event-table
	MimeType *string `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
}

func (s UpdateAgentRequestMetadataAttachments) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequestMetadataAttachments) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequestMetadataAttachments) GetArn() *string {
	return s.Arn
}

func (s *UpdateAgentRequestMetadataAttachments) GetMimeType() *string {
	return s.MimeType
}

func (s *UpdateAgentRequestMetadataAttachments) SetArn(v string) *UpdateAgentRequestMetadataAttachments {
	s.Arn = &v
	return s
}

func (s *UpdateAgentRequestMetadataAttachments) SetMimeType(v string) *UpdateAgentRequestMetadataAttachments {
	s.MimeType = &v
	return s
}

func (s *UpdateAgentRequestMetadataAttachments) Validate() error {
	return dara.Validate(s)
}

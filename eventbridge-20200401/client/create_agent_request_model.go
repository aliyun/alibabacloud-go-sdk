// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateAgentRequest
	GetDescription() *string
	SetMetadata(v *CreateAgentRequestMetadata) *CreateAgentRequest
	GetMetadata() *CreateAgentRequestMetadata
	SetName(v string) *CreateAgentRequest
	GetName() *string
	SetPrompt(v string) *CreateAgentRequest
	GetPrompt() *string
}

type CreateAgentRequest struct {
	Description *string                     `json:"Description,omitempty" xml:"Description,omitempty"`
	Metadata    *CreateAgentRequestMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Struct"`
	// example:
	//
	// my-agent
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
}

func (s CreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentRequest) GetMetadata() *CreateAgentRequestMetadata {
	return s.Metadata
}

func (s *CreateAgentRequest) GetName() *string {
	return s.Name
}

func (s *CreateAgentRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateAgentRequest) SetDescription(v string) *CreateAgentRequest {
	s.Description = &v
	return s
}

func (s *CreateAgentRequest) SetMetadata(v *CreateAgentRequestMetadata) *CreateAgentRequest {
	s.Metadata = v
	return s
}

func (s *CreateAgentRequest) SetName(v string) *CreateAgentRequest {
	s.Name = &v
	return s
}

func (s *CreateAgentRequest) SetPrompt(v string) *CreateAgentRequest {
	s.Prompt = &v
	return s
}

func (s *CreateAgentRequest) Validate() error {
	if s.Metadata != nil {
		if err := s.Metadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentRequestMetadata struct {
	Attachments []*CreateAgentRequestMetadataAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
}

func (s CreateAgentRequestMetadata) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequestMetadata) GoString() string {
	return s.String()
}

func (s *CreateAgentRequestMetadata) GetAttachments() []*CreateAgentRequestMetadataAttachments {
	return s.Attachments
}

func (s *CreateAgentRequestMetadata) SetAttachments(v []*CreateAgentRequestMetadataAttachments) *CreateAgentRequestMetadata {
	s.Attachments = v
	return s
}

func (s *CreateAgentRequestMetadata) Validate() error {
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

type CreateAgentRequestMetadataAttachments struct {
	// example:
	//
	// acs:eventbridge:cn-hangzhou:12345:eventhouse/system-rocketmq/namespace/rmq-cn-XXX/table/order
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// example:
	//
	// inner-resource/event-table
	MimeType *string `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
}

func (s CreateAgentRequestMetadataAttachments) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequestMetadataAttachments) GoString() string {
	return s.String()
}

func (s *CreateAgentRequestMetadataAttachments) GetArn() *string {
	return s.Arn
}

func (s *CreateAgentRequestMetadataAttachments) GetMimeType() *string {
	return s.MimeType
}

func (s *CreateAgentRequestMetadataAttachments) SetArn(v string) *CreateAgentRequestMetadataAttachments {
	s.Arn = &v
	return s
}

func (s *CreateAgentRequestMetadataAttachments) SetMimeType(v string) *CreateAgentRequestMetadataAttachments {
	s.MimeType = &v
	return s
}

func (s *CreateAgentRequestMetadataAttachments) Validate() error {
	return dara.Validate(s)
}

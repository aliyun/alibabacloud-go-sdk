// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *CreateTicketRequest
	GetContactId() *string
	SetContext(v string) *CreateTicketRequest
	GetContext() *string
	SetCustomerId(v string) *CreateTicketRequest
	GetCustomerId() *string
	SetInstanceId(v string) *CreateTicketRequest
	GetInstanceId() *string
	SetSource(v string) *CreateTicketRequest
	GetSource() *string
	SetTemplateId(v string) *CreateTicketRequest
	GetTemplateId() *string
	SetTitle(v string) *CreateTicketRequest
	GetTitle() *string
}

type CreateTicketRequest struct {
	// example:
	//
	// job-38860977107324****
	ContactId *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	Context   *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// example:
	//
	// 51e155ce-3747-4f21-b402-13c69597b920
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// CHAT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// e9e4c76c-948d-4a6e-9ce2-9da0f5967a73
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) GetContactId() *string {
	return s.ContactId
}

func (s *CreateTicketRequest) GetContext() *string {
	return s.Context
}

func (s *CreateTicketRequest) GetCustomerId() *string {
	return s.CustomerId
}

func (s *CreateTicketRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTicketRequest) GetSource() *string {
	return s.Source
}

func (s *CreateTicketRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateTicketRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateTicketRequest) SetContactId(v string) *CreateTicketRequest {
	s.ContactId = &v
	return s
}

func (s *CreateTicketRequest) SetContext(v string) *CreateTicketRequest {
	s.Context = &v
	return s
}

func (s *CreateTicketRequest) SetCustomerId(v string) *CreateTicketRequest {
	s.CustomerId = &v
	return s
}

func (s *CreateTicketRequest) SetInstanceId(v string) *CreateTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTicketRequest) SetSource(v string) *CreateTicketRequest {
	s.Source = &v
	return s
}

func (s *CreateTicketRequest) SetTemplateId(v string) *CreateTicketRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketRequest) Validate() error {
	return dara.Validate(s)
}

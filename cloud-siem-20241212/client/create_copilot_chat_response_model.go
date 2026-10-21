// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCopilotChatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCopilotChatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCopilotChatResponse
	GetStatusCode() *int32
	SetId(v string) *CreateCopilotChatResponse
	GetId() *string
	SetEvent(v string) *CreateCopilotChatResponse
	GetEvent() *string
	SetBody(v string) *CreateCopilotChatResponse
	GetBody() *string
}

type CreateCopilotChatResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCopilotChatResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotChatResponse) GoString() string {
	return s.String()
}

func (s *CreateCopilotChatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCopilotChatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCopilotChatResponse) GetId() *string {
	return s.Id
}

func (s *CreateCopilotChatResponse) GetEvent() *string {
	return s.Event
}

func (s *CreateCopilotChatResponse) GetBody() *string {
	return s.Body
}

func (s *CreateCopilotChatResponse) SetHeaders(v map[string]*string) *CreateCopilotChatResponse {
	s.Headers = v
	return s
}

func (s *CreateCopilotChatResponse) SetStatusCode(v int32) *CreateCopilotChatResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCopilotChatResponse) SetId(v string) *CreateCopilotChatResponse {
	s.Id = &v
	return s
}

func (s *CreateCopilotChatResponse) SetEvent(v string) *CreateCopilotChatResponse {
	s.Event = &v
	return s
}

func (s *CreateCopilotChatResponse) SetBody(v string) *CreateCopilotChatResponse {
	s.Body = &v
	return s
}

func (s *CreateCopilotChatResponse) Validate() error {
	return dara.Validate(s)
}

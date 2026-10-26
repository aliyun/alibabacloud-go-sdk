// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConversationEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentConversationEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentConversationEventsResponse
	GetStatusCode() *int32
	SetId(v string) *GetAgentConversationEventsResponse
	GetId() *string
	SetEvent(v string) *GetAgentConversationEventsResponse
	GetEvent() *string
	SetBody(v string) *GetAgentConversationEventsResponse
	GetBody() *string
}

type GetAgentConversationEventsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string            `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string            `json:"event,omitempty" xml:"event,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentConversationEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConversationEventsResponse) GoString() string {
	return s.String()
}

func (s *GetAgentConversationEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentConversationEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentConversationEventsResponse) GetId() *string {
	return s.Id
}

func (s *GetAgentConversationEventsResponse) GetEvent() *string {
	return s.Event
}

func (s *GetAgentConversationEventsResponse) GetBody() *string {
	return s.Body
}

func (s *GetAgentConversationEventsResponse) SetHeaders(v map[string]*string) *GetAgentConversationEventsResponse {
	s.Headers = v
	return s
}

func (s *GetAgentConversationEventsResponse) SetStatusCode(v int32) *GetAgentConversationEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentConversationEventsResponse) SetId(v string) *GetAgentConversationEventsResponse {
	s.Id = &v
	return s
}

func (s *GetAgentConversationEventsResponse) SetEvent(v string) *GetAgentConversationEventsResponse {
	s.Event = &v
	return s
}

func (s *GetAgentConversationEventsResponse) SetBody(v string) *GetAgentConversationEventsResponse {
	s.Body = &v
	return s
}

func (s *GetAgentConversationEventsResponse) Validate() error {
	return dara.Validate(s)
}

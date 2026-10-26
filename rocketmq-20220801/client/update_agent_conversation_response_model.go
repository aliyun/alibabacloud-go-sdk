// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentConversationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentConversationResponseBody) *UpdateAgentConversationResponse
	GetBody() *UpdateAgentConversationResponseBody
}

type UpdateAgentConversationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConversationResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentConversationResponse) GetBody() *UpdateAgentConversationResponseBody {
	return s.Body
}

func (s *UpdateAgentConversationResponse) SetHeaders(v map[string]*string) *UpdateAgentConversationResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentConversationResponse) SetStatusCode(v int32) *UpdateAgentConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentConversationResponse) SetBody(v *UpdateAgentConversationResponseBody) *UpdateAgentConversationResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

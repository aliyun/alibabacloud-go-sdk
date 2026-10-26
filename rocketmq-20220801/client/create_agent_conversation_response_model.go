// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentConversationResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentConversationResponseBody) *CreateAgentConversationResponse
	GetBody() *CreateAgentConversationResponseBody
}

type CreateAgentConversationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentConversationResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentConversationResponse) GetBody() *CreateAgentConversationResponseBody {
	return s.Body
}

func (s *CreateAgentConversationResponse) SetHeaders(v map[string]*string) *CreateAgentConversationResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentConversationResponse) SetStatusCode(v int32) *CreateAgentConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentConversationResponse) SetBody(v *CreateAgentConversationResponseBody) *CreateAgentConversationResponse {
	s.Body = v
	return s
}

func (s *CreateAgentConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

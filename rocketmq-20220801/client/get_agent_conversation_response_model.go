// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentConversationResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentConversationResponseBody) *GetAgentConversationResponse
	GetBody() *GetAgentConversationResponseBody
}

type GetAgentConversationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConversationResponse) GoString() string {
	return s.String()
}

func (s *GetAgentConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentConversationResponse) GetBody() *GetAgentConversationResponseBody {
	return s.Body
}

func (s *GetAgentConversationResponse) SetHeaders(v map[string]*string) *GetAgentConversationResponse {
	s.Headers = v
	return s
}

func (s *GetAgentConversationResponse) SetStatusCode(v int32) *GetAgentConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentConversationResponse) SetBody(v *GetAgentConversationResponseBody) *GetAgentConversationResponse {
	s.Body = v
	return s
}

func (s *GetAgentConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

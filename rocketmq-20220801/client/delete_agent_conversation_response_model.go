// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAgentConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAgentConversationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAgentConversationResponseBody) *DeleteAgentConversationResponse
	GetBody() *DeleteAgentConversationResponseBody
}

type DeleteAgentConversationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentConversationResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAgentConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAgentConversationResponse) GetBody() *DeleteAgentConversationResponseBody {
	return s.Body
}

func (s *DeleteAgentConversationResponse) SetHeaders(v map[string]*string) *DeleteAgentConversationResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentConversationResponse) SetStatusCode(v int32) *DeleteAgentConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentConversationResponse) SetBody(v *DeleteAgentConversationResponseBody) *DeleteAgentConversationResponse {
	s.Body = v
	return s
}

func (s *DeleteAgentConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

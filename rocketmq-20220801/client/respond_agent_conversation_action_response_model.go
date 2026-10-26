// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRespondAgentConversationActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RespondAgentConversationActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RespondAgentConversationActionResponse
	GetStatusCode() *int32
	SetBody(v *RespondAgentConversationActionResponseBody) *RespondAgentConversationActionResponse
	GetBody() *RespondAgentConversationActionResponseBody
}

type RespondAgentConversationActionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RespondAgentConversationActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RespondAgentConversationActionResponse) String() string {
	return dara.Prettify(s)
}

func (s RespondAgentConversationActionResponse) GoString() string {
	return s.String()
}

func (s *RespondAgentConversationActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RespondAgentConversationActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RespondAgentConversationActionResponse) GetBody() *RespondAgentConversationActionResponseBody {
	return s.Body
}

func (s *RespondAgentConversationActionResponse) SetHeaders(v map[string]*string) *RespondAgentConversationActionResponse {
	s.Headers = v
	return s
}

func (s *RespondAgentConversationActionResponse) SetStatusCode(v int32) *RespondAgentConversationActionResponse {
	s.StatusCode = &v
	return s
}

func (s *RespondAgentConversationActionResponse) SetBody(v *RespondAgentConversationActionResponseBody) *RespondAgentConversationActionResponse {
	s.Body = v
	return s
}

func (s *RespondAgentConversationActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

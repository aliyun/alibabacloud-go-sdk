// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentConversationShareScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentConversationShareScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentConversationShareScopeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentConversationShareScopeResponseBody) *UpdateAgentConversationShareScopeResponse
	GetBody() *UpdateAgentConversationShareScopeResponseBody
}

type UpdateAgentConversationShareScopeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentConversationShareScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentConversationShareScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentConversationShareScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentConversationShareScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentConversationShareScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentConversationShareScopeResponse) GetBody() *UpdateAgentConversationShareScopeResponseBody {
	return s.Body
}

func (s *UpdateAgentConversationShareScopeResponse) SetHeaders(v map[string]*string) *UpdateAgentConversationShareScopeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentConversationShareScopeResponse) SetStatusCode(v int32) *UpdateAgentConversationShareScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentConversationShareScopeResponse) SetBody(v *UpdateAgentConversationShareScopeResponseBody) *UpdateAgentConversationShareScopeResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentConversationShareScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

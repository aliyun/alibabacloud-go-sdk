// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentIMChannelCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentIMChannelCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentIMChannelCredentialResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentIMChannelCredentialResponseBody) *UpdateAgentIMChannelCredentialResponse
	GetBody() *UpdateAgentIMChannelCredentialResponseBody
}

type UpdateAgentIMChannelCredentialResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentIMChannelCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentIMChannelCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelCredentialResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentIMChannelCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentIMChannelCredentialResponse) GetBody() *UpdateAgentIMChannelCredentialResponseBody {
	return s.Body
}

func (s *UpdateAgentIMChannelCredentialResponse) SetHeaders(v map[string]*string) *UpdateAgentIMChannelCredentialResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponse) SetStatusCode(v int32) *UpdateAgentIMChannelCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponse) SetBody(v *UpdateAgentIMChannelCredentialResponseBody) *UpdateAgentIMChannelCredentialResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentIMChannelCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

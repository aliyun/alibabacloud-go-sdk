// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentAuthModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentAuthModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentAuthModeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentAuthModeResponseBody) *UpdateAgentAuthModeResponse
	GetBody() *UpdateAgentAuthModeResponseBody
}

type UpdateAgentAuthModeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentAuthModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentAuthModeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentAuthModeResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentAuthModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentAuthModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentAuthModeResponse) GetBody() *UpdateAgentAuthModeResponseBody {
	return s.Body
}

func (s *UpdateAgentAuthModeResponse) SetHeaders(v map[string]*string) *UpdateAgentAuthModeResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentAuthModeResponse) SetStatusCode(v int32) *UpdateAgentAuthModeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentAuthModeResponse) SetBody(v *UpdateAgentAuthModeResponseBody) *UpdateAgentAuthModeResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentAuthModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

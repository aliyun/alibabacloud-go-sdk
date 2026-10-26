// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSessionTitleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentSessionTitleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentSessionTitleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentSessionTitleResponseBody) *UpdateAgentSessionTitleResponse
	GetBody() *UpdateAgentSessionTitleResponseBody
}

type UpdateAgentSessionTitleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentSessionTitleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentSessionTitleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSessionTitleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentSessionTitleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentSessionTitleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentSessionTitleResponse) GetBody() *UpdateAgentSessionTitleResponseBody {
	return s.Body
}

func (s *UpdateAgentSessionTitleResponse) SetHeaders(v map[string]*string) *UpdateAgentSessionTitleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentSessionTitleResponse) SetStatusCode(v int32) *UpdateAgentSessionTitleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentSessionTitleResponse) SetBody(v *UpdateAgentSessionTitleResponseBody) *UpdateAgentSessionTitleResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentSessionTitleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomAgentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomAgentResponseBody) *UpdateCustomAgentResponse
	GetBody() *UpdateCustomAgentResponseBody
}

type UpdateCustomAgentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomAgentResponse) GetBody() *UpdateCustomAgentResponseBody {
	return s.Body
}

func (s *UpdateCustomAgentResponse) SetHeaders(v map[string]*string) *UpdateCustomAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomAgentResponse) SetStatusCode(v int32) *UpdateCustomAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomAgentResponse) SetBody(v *UpdateCustomAgentResponseBody) *UpdateCustomAgentResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentSpaceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentSpaceResponseBody) *UpdateAgentSpaceResponse
	GetBody() *UpdateAgentSpaceResponseBody
}

type UpdateAgentSpaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSpaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentSpaceResponse) GetBody() *UpdateAgentSpaceResponseBody {
	return s.Body
}

func (s *UpdateAgentSpaceResponse) SetHeaders(v map[string]*string) *UpdateAgentSpaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentSpaceResponse) SetStatusCode(v int32) *UpdateAgentSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentSpaceResponse) SetBody(v *UpdateAgentSpaceResponseBody) *UpdateAgentSpaceResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentSpaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentSpaceResponseBody) *CreateAgentSpaceResponse
	GetBody() *CreateAgentSpaceResponseBody
}

type CreateAgentSpaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentSpaceResponse) GetBody() *CreateAgentSpaceResponseBody {
	return s.Body
}

func (s *CreateAgentSpaceResponse) SetHeaders(v map[string]*string) *CreateAgentSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentSpaceResponse) SetStatusCode(v int32) *CreateAgentSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentSpaceResponse) SetBody(v *CreateAgentSpaceResponseBody) *CreateAgentSpaceResponse {
	s.Body = v
	return s
}

func (s *CreateAgentSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

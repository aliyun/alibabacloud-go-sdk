// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpecVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentSpecVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentSpecVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentSpecVersionResponseBody) *CreateAgentSpecVersionResponse
	GetBody() *CreateAgentSpecVersionResponseBody
}

type CreateAgentSpecVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentSpecVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentSpecVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentSpecVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentSpecVersionResponse) GetBody() *CreateAgentSpecVersionResponseBody {
	return s.Body
}

func (s *CreateAgentSpecVersionResponse) SetHeaders(v map[string]*string) *CreateAgentSpecVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentSpecVersionResponse) SetStatusCode(v int32) *CreateAgentSpecVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentSpecVersionResponse) SetBody(v *CreateAgentSpecVersionResponseBody) *CreateAgentSpecVersionResponse {
	s.Body = v
	return s
}

func (s *CreateAgentSpecVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

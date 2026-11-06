// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentVersionResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentVersionResponseBody) *CreateAgentVersionResponse
	GetBody() *CreateAgentVersionResponseBody
}

type CreateAgentVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentVersionResponse) GetBody() *CreateAgentVersionResponseBody {
	return s.Body
}

func (s *CreateAgentVersionResponse) SetHeaders(v map[string]*string) *CreateAgentVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentVersionResponse) SetStatusCode(v int32) *CreateAgentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentVersionResponse) SetBody(v *CreateAgentVersionResponseBody) *CreateAgentVersionResponse {
	s.Body = v
	return s
}

func (s *CreateAgentVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

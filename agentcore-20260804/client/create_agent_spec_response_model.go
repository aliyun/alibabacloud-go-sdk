// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAgentSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAgentSpecResponse
	GetStatusCode() *int32
	SetBody(v *CreateAgentSpecResponseBody) *CreateAgentSpecResponse
	GetBody() *CreateAgentSpecResponseBody
}

type CreateAgentSpecResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAgentSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAgentSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSpecResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAgentSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAgentSpecResponse) GetBody() *CreateAgentSpecResponseBody {
	return s.Body
}

func (s *CreateAgentSpecResponse) SetHeaders(v map[string]*string) *CreateAgentSpecResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentSpecResponse) SetStatusCode(v int32) *CreateAgentSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentSpecResponse) SetBody(v *CreateAgentSpecResponseBody) *CreateAgentSpecResponse {
	s.Body = v
	return s
}

func (s *CreateAgentSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

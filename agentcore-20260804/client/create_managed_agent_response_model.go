// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManagedAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateManagedAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateManagedAgentResponse
	GetStatusCode() *int32
	SetBody(v *CreateManagedAgentResponseBody) *CreateManagedAgentResponse
	GetBody() *CreateManagedAgentResponseBody
}

type CreateManagedAgentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateManagedAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateManagedAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateManagedAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateManagedAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateManagedAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateManagedAgentResponse) GetBody() *CreateManagedAgentResponseBody {
	return s.Body
}

func (s *CreateManagedAgentResponse) SetHeaders(v map[string]*string) *CreateManagedAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateManagedAgentResponse) SetStatusCode(v int32) *CreateManagedAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateManagedAgentResponse) SetBody(v *CreateManagedAgentResponseBody) *CreateManagedAgentResponse {
	s.Body = v
	return s
}

func (s *CreateManagedAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

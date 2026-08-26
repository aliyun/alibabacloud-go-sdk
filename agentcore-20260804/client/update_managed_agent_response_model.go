// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateManagedAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateManagedAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateManagedAgentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateManagedAgentResponseBody) *UpdateManagedAgentResponse
	GetBody() *UpdateManagedAgentResponseBody
}

type UpdateManagedAgentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateManagedAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateManagedAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateManagedAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdateManagedAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateManagedAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateManagedAgentResponse) GetBody() *UpdateManagedAgentResponseBody {
	return s.Body
}

func (s *UpdateManagedAgentResponse) SetHeaders(v map[string]*string) *UpdateManagedAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdateManagedAgentResponse) SetStatusCode(v int32) *UpdateManagedAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateManagedAgentResponse) SetBody(v *UpdateManagedAgentResponseBody) *UpdateManagedAgentResponse {
	s.Body = v
	return s
}

func (s *UpdateManagedAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

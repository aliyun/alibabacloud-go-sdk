// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentSpecResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentSpecResponseBody) *UpdateAgentSpecResponse
	GetBody() *UpdateAgentSpecResponseBody
}

type UpdateAgentSpecResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentSpecResponse) GetBody() *UpdateAgentSpecResponseBody {
	return s.Body
}

func (s *UpdateAgentSpecResponse) SetHeaders(v map[string]*string) *UpdateAgentSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentSpecResponse) SetStatusCode(v int32) *UpdateAgentSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentSpecResponse) SetBody(v *UpdateAgentSpecResponseBody) *UpdateAgentSpecResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

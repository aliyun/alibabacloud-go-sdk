// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAgentStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAgentStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAgentStatusResponseBody) *UpdateAgentStatusResponse
	GetBody() *UpdateAgentStatusResponseBody
}

type UpdateAgentStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAgentStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAgentStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAgentStatusResponse) GetBody() *UpdateAgentStatusResponseBody {
	return s.Body
}

func (s *UpdateAgentStatusResponse) SetHeaders(v map[string]*string) *UpdateAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentStatusResponse) SetStatusCode(v int32) *UpdateAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentStatusResponse) SetBody(v *UpdateAgentStatusResponseBody) *UpdateAgentStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateAgentStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

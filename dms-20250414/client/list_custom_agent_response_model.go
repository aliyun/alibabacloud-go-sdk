// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomAgentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomAgentResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomAgentResponseBody) *ListCustomAgentResponse
	GetBody() *ListCustomAgentResponseBody
}

type ListCustomAgentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomAgentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponse) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomAgentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomAgentResponse) GetBody() *ListCustomAgentResponseBody {
	return s.Body
}

func (s *ListCustomAgentResponse) SetHeaders(v map[string]*string) *ListCustomAgentResponse {
	s.Headers = v
	return s
}

func (s *ListCustomAgentResponse) SetStatusCode(v int32) *ListCustomAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomAgentResponse) SetBody(v *ListCustomAgentResponseBody) *ListCustomAgentResponse {
	s.Body = v
	return s
}

func (s *ListCustomAgentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomAgentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomAgentsResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomAgentsResponseBody) *ListCustomAgentsResponse
	GetBody() *ListCustomAgentsResponseBody
}

type ListCustomAgentsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomAgentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomAgentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomAgentsResponse) GetBody() *ListCustomAgentsResponseBody {
	return s.Body
}

func (s *ListCustomAgentsResponse) SetHeaders(v map[string]*string) *ListCustomAgentsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomAgentsResponse) SetStatusCode(v int32) *ListCustomAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomAgentsResponse) SetBody(v *ListCustomAgentsResponseBody) *ListCustomAgentsResponse {
	s.Body = v
	return s
}

func (s *ListCustomAgentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

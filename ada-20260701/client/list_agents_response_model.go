// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentsResponseBody) *ListAgentsResponse
	GetBody() *ListAgentsResponseBody
}

type ListAgentsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentsResponse) GetBody() *ListAgentsResponseBody {
	return s.Body
}

func (s *ListAgentsResponse) SetHeaders(v map[string]*string) *ListAgentsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentsResponse) SetStatusCode(v int32) *ListAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentsResponse) SetBody(v *ListAgentsResponseBody) *ListAgentsResponse {
	s.Body = v
	return s
}

func (s *ListAgentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

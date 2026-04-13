// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSessionsResponse
	GetStatusCode() *int32
	SetBody(v *ListSessionsResponseBody) *ListSessionsResponse
	GetBody() *ListSessionsResponseBody
}

type ListSessionsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSessionsResponse) GetBody() *ListSessionsResponseBody {
	return s.Body
}

func (s *ListSessionsResponse) SetHeaders(v map[string]*string) *ListSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListSessionsResponse) SetStatusCode(v int32) *ListSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSessionsResponse) SetBody(v *ListSessionsResponseBody) *ListSessionsResponse {
	s.Body = v
	return s
}

func (s *ListSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskExecutionsResponseBody) *ListTaskExecutionsResponse
	GetBody() *ListTaskExecutionsResponseBody
}

type ListTaskExecutionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskExecutionsResponse) GetBody() *ListTaskExecutionsResponseBody {
	return s.Body
}

func (s *ListTaskExecutionsResponse) SetHeaders(v map[string]*string) *ListTaskExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskExecutionsResponse) SetStatusCode(v int32) *ListTaskExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskExecutionsResponse) SetBody(v *ListTaskExecutionsResponseBody) *ListTaskExecutionsResponse {
	s.Body = v
	return s
}

func (s *ListTaskExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

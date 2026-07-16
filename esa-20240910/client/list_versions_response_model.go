// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListVersionsResponseBody) *ListVersionsResponse
	GetBody() *ListVersionsResponseBody
}

type ListVersionsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVersionsResponse) GetBody() *ListVersionsResponseBody {
	return s.Body
}

func (s *ListVersionsResponse) SetHeaders(v map[string]*string) *ListVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListVersionsResponse) SetStatusCode(v int32) *ListVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVersionsResponse) SetBody(v *ListVersionsResponseBody) *ListVersionsResponse {
	s.Body = v
	return s
}

func (s *ListVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

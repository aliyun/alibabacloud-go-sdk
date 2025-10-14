// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListObjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListObjectsResponseBody) *ListObjectsResponse
	GetBody() *ListObjectsResponseBody
}

type ListObjectsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListObjectsResponse) GoString() string {
	return s.String()
}

func (s *ListObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListObjectsResponse) GetBody() *ListObjectsResponseBody {
	return s.Body
}

func (s *ListObjectsResponse) SetHeaders(v map[string]*string) *ListObjectsResponse {
	s.Headers = v
	return s
}

func (s *ListObjectsResponse) SetStatusCode(v int32) *ListObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListObjectsResponse) SetBody(v *ListObjectsResponseBody) *ListObjectsResponse {
	s.Body = v
	return s
}

func (s *ListObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

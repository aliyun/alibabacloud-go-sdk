// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogProjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListLogProjectsResponseBody) *ListLogProjectsResponse
	GetBody() *ListLogProjectsResponseBody
}

type ListLogProjectsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListLogProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogProjectsResponse) GetBody() *ListLogProjectsResponseBody {
	return s.Body
}

func (s *ListLogProjectsResponse) SetHeaders(v map[string]*string) *ListLogProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListLogProjectsResponse) SetStatusCode(v int32) *ListLogProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogProjectsResponse) SetBody(v *ListLogProjectsResponseBody) *ListLogProjectsResponse {
	s.Body = v
	return s
}

func (s *ListLogProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

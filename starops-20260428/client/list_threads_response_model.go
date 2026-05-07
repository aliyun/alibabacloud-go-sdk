// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListThreadsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListThreadsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListThreadsResponse
	GetStatusCode() *int32
	SetBody(v *ListThreadsResponseBody) *ListThreadsResponse
	GetBody() *ListThreadsResponseBody
}

type ListThreadsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListThreadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListThreadsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListThreadsResponse) GoString() string {
	return s.String()
}

func (s *ListThreadsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListThreadsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListThreadsResponse) GetBody() *ListThreadsResponseBody {
	return s.Body
}

func (s *ListThreadsResponse) SetHeaders(v map[string]*string) *ListThreadsResponse {
	s.Headers = v
	return s
}

func (s *ListThreadsResponse) SetStatusCode(v int32) *ListThreadsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListThreadsResponse) SetBody(v *ListThreadsResponseBody) *ListThreadsResponse {
	s.Body = v
	return s
}

func (s *ListThreadsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

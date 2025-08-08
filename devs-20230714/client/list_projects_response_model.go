// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectsResponseBody) *ListProjectsResponse
	GetBody() *ListProjectsResponseBody
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectsResponse) GetBody() *ListProjectsResponseBody {
	return s.Body
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

func (s *ListProjectsResponse) Validate() error {
	return dara.Validate(s)
}

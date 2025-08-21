// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRenderingProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRenderingProjectsResponse
	GetStatusCode() *int32
	SetBody(v *ListRenderingProjectsResponseBody) *ListRenderingProjectsResponse
	GetBody() *ListRenderingProjectsResponseBody
}

type ListRenderingProjectsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRenderingProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRenderingProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRenderingProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRenderingProjectsResponse) GetBody() *ListRenderingProjectsResponseBody {
	return s.Body
}

func (s *ListRenderingProjectsResponse) SetHeaders(v map[string]*string) *ListRenderingProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListRenderingProjectsResponse) SetStatusCode(v int32) *ListRenderingProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRenderingProjectsResponse) SetBody(v *ListRenderingProjectsResponseBody) *ListRenderingProjectsResponse {
	s.Body = v
	return s
}

func (s *ListRenderingProjectsResponse) Validate() error {
	return dara.Validate(s)
}

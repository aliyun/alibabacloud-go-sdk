// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectIdsResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectIdsResponseBody) *ListProjectIdsResponse
	GetBody() *ListProjectIdsResponseBody
}

type ListProjectIdsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectIdsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectIdsResponse) GetBody() *ListProjectIdsResponseBody {
	return s.Body
}

func (s *ListProjectIdsResponse) SetHeaders(v map[string]*string) *ListProjectIdsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectIdsResponse) SetStatusCode(v int32) *ListProjectIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectIdsResponse) SetBody(v *ListProjectIdsResponseBody) *ListProjectIdsResponse {
	s.Body = v
	return s
}

func (s *ListProjectIdsResponse) Validate() error {
	return dara.Validate(s)
}

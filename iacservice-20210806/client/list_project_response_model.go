// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProjectResponse
	GetStatusCode() *int32
	SetBody(v *ListProjectResponseBody) *ListProjectResponse
	GetBody() *ListProjectResponseBody
}

type ListProjectResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponse) GoString() string {
	return s.String()
}

func (s *ListProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProjectResponse) GetBody() *ListProjectResponseBody {
	return s.Body
}

func (s *ListProjectResponse) SetHeaders(v map[string]*string) *ListProjectResponse {
	s.Headers = v
	return s
}

func (s *ListProjectResponse) SetStatusCode(v int32) *ListProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectResponse) SetBody(v *ListProjectResponseBody) *ListProjectResponse {
	s.Body = v
	return s
}

func (s *ListProjectResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDAGVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDAGVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDAGVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDAGVersionsResponseBody) *ListDAGVersionsResponse
	GetBody() *ListDAGVersionsResponseBody
}

type ListDAGVersionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDAGVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDAGVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDAGVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListDAGVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDAGVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDAGVersionsResponse) GetBody() *ListDAGVersionsResponseBody {
	return s.Body
}

func (s *ListDAGVersionsResponse) SetHeaders(v map[string]*string) *ListDAGVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListDAGVersionsResponse) SetStatusCode(v int32) *ListDAGVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDAGVersionsResponse) SetBody(v *ListDAGVersionsResponseBody) *ListDAGVersionsResponse {
	s.Body = v
	return s
}

func (s *ListDAGVersionsResponse) Validate() error {
	return dara.Validate(s)
}

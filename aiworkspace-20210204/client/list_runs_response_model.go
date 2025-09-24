// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListRunsResponseBody) *ListRunsResponse
	GetBody() *ListRunsResponseBody
}

type ListRunsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRunsResponse) GoString() string {
	return s.String()
}

func (s *ListRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRunsResponse) GetBody() *ListRunsResponseBody {
	return s.Body
}

func (s *ListRunsResponse) SetHeaders(v map[string]*string) *ListRunsResponse {
	s.Headers = v
	return s
}

func (s *ListRunsResponse) SetStatusCode(v int32) *ListRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRunsResponse) SetBody(v *ListRunsResponseBody) *ListRunsResponse {
	s.Body = v
	return s
}

func (s *ListRunsResponse) Validate() error {
	return dara.Validate(s)
}

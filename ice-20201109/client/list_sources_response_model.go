// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListSourcesResponseBody) *ListSourcesResponse
	GetBody() *ListSourcesResponseBody
}

type ListSourcesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSourcesResponse) GetBody() *ListSourcesResponseBody {
	return s.Body
}

func (s *ListSourcesResponse) SetHeaders(v map[string]*string) *ListSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListSourcesResponse) SetStatusCode(v int32) *ListSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSourcesResponse) SetBody(v *ListSourcesResponseBody) *ListSourcesResponse {
	s.Body = v
	return s
}

func (s *ListSourcesResponse) Validate() error {
	return dara.Validate(s)
}

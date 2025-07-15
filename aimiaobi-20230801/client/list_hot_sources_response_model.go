// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListHotSourcesResponseBody) *ListHotSourcesResponse
	GetBody() *ListHotSourcesResponseBody
}

type ListHotSourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListHotSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotSourcesResponse) GetBody() *ListHotSourcesResponseBody {
	return s.Body
}

func (s *ListHotSourcesResponse) SetHeaders(v map[string]*string) *ListHotSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListHotSourcesResponse) SetStatusCode(v int32) *ListHotSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotSourcesResponse) SetBody(v *ListHotSourcesResponseBody) *ListHotSourcesResponse {
	s.Body = v
	return s
}

func (s *ListHotSourcesResponse) Validate() error {
	return dara.Validate(s)
}

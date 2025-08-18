// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOriginPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOriginPoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListOriginPoolsResponseBody) *ListOriginPoolsResponse
	GetBody() *ListOriginPoolsResponseBody
}

type ListOriginPoolsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOriginPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOriginPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOriginPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListOriginPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOriginPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOriginPoolsResponse) GetBody() *ListOriginPoolsResponseBody {
	return s.Body
}

func (s *ListOriginPoolsResponse) SetHeaders(v map[string]*string) *ListOriginPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListOriginPoolsResponse) SetStatusCode(v int32) *ListOriginPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOriginPoolsResponse) SetBody(v *ListOriginPoolsResponseBody) *ListOriginPoolsResponse {
	s.Body = v
	return s
}

func (s *ListOriginPoolsResponse) Validate() error {
	return dara.Validate(s)
}

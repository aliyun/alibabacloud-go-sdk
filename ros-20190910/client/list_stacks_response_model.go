// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStacksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStacksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStacksResponse
	GetStatusCode() *int32
	SetBody(v *ListStacksResponseBody) *ListStacksResponse
	GetBody() *ListStacksResponseBody
}

type ListStacksResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStacksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStacksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStacksResponse) GoString() string {
	return s.String()
}

func (s *ListStacksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStacksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStacksResponse) GetBody() *ListStacksResponseBody {
	return s.Body
}

func (s *ListStacksResponse) SetHeaders(v map[string]*string) *ListStacksResponse {
	s.Headers = v
	return s
}

func (s *ListStacksResponse) SetStatusCode(v int32) *ListStacksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStacksResponse) SetBody(v *ListStacksResponseBody) *ListStacksResponse {
	s.Body = v
	return s
}

func (s *ListStacksResponse) Validate() error {
	return dara.Validate(s)
}

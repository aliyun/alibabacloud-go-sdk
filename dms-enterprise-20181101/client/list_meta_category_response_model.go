// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetaCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetaCategoryResponse
	GetStatusCode() *int32
	SetBody(v *ListMetaCategoryResponseBody) *ListMetaCategoryResponse
	GetBody() *ListMetaCategoryResponseBody
}

type ListMetaCategoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetaCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetaCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListMetaCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetaCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetaCategoryResponse) GetBody() *ListMetaCategoryResponseBody {
	return s.Body
}

func (s *ListMetaCategoryResponse) SetHeaders(v map[string]*string) *ListMetaCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListMetaCategoryResponse) SetStatusCode(v int32) *ListMetaCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetaCategoryResponse) SetBody(v *ListMetaCategoryResponseBody) *ListMetaCategoryResponse {
	s.Body = v
	return s
}

func (s *ListMetaCategoryResponse) Validate() error {
	return dara.Validate(s)
}

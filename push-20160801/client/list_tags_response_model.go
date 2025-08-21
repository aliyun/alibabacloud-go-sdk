// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListTagsResponseBody) *ListTagsResponse
	GetBody() *ListTagsResponseBody
}

type ListTagsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponse) GoString() string {
	return s.String()
}

func (s *ListTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagsResponse) GetBody() *ListTagsResponseBody {
	return s.Body
}

func (s *ListTagsResponse) SetHeaders(v map[string]*string) *ListTagsResponse {
	s.Headers = v
	return s
}

func (s *ListTagsResponse) SetStatusCode(v int32) *ListTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

func (s *ListTagsResponse) Validate() error {
	return dara.Validate(s)
}

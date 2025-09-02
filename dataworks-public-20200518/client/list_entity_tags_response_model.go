// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntityTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEntityTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEntityTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListEntityTagsResponseBody) *ListEntityTagsResponse
	GetBody() *ListEntityTagsResponseBody
}

type ListEntityTagsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEntityTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEntityTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEntityTagsResponse) GoString() string {
	return s.String()
}

func (s *ListEntityTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEntityTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEntityTagsResponse) GetBody() *ListEntityTagsResponseBody {
	return s.Body
}

func (s *ListEntityTagsResponse) SetHeaders(v map[string]*string) *ListEntityTagsResponse {
	s.Headers = v
	return s
}

func (s *ListEntityTagsResponse) SetStatusCode(v int32) *ListEntityTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEntityTagsResponse) SetBody(v *ListEntityTagsResponseBody) *ListEntityTagsResponse {
	s.Body = v
	return s
}

func (s *ListEntityTagsResponse) Validate() error {
	return dara.Validate(s)
}

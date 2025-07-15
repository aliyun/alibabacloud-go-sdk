// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCallTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCallTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListCallTagsResponseBody) *ListCallTagsResponse
	GetBody() *ListCallTagsResponseBody
}

type ListCallTagsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCallTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCallTagsResponse) GoString() string {
	return s.String()
}

func (s *ListCallTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCallTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCallTagsResponse) GetBody() *ListCallTagsResponseBody {
	return s.Body
}

func (s *ListCallTagsResponse) SetHeaders(v map[string]*string) *ListCallTagsResponse {
	s.Headers = v
	return s
}

func (s *ListCallTagsResponse) SetStatusCode(v int32) *ListCallTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallTagsResponse) SetBody(v *ListCallTagsResponseBody) *ListCallTagsResponse {
	s.Body = v
	return s
}

func (s *ListCallTagsResponse) Validate() error {
	return dara.Validate(s)
}

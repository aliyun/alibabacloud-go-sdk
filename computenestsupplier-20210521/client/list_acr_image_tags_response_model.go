// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAcrImageTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAcrImageTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAcrImageTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListAcrImageTagsResponseBody) *ListAcrImageTagsResponse
	GetBody() *ListAcrImageTagsResponseBody
}

type ListAcrImageTagsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAcrImageTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAcrImageTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAcrImageTagsResponse) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAcrImageTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAcrImageTagsResponse) GetBody() *ListAcrImageTagsResponseBody {
	return s.Body
}

func (s *ListAcrImageTagsResponse) SetHeaders(v map[string]*string) *ListAcrImageTagsResponse {
	s.Headers = v
	return s
}

func (s *ListAcrImageTagsResponse) SetStatusCode(v int32) *ListAcrImageTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAcrImageTagsResponse) SetBody(v *ListAcrImageTagsResponseBody) *ListAcrImageTagsResponse {
	s.Body = v
	return s
}

func (s *ListAcrImageTagsResponse) Validate() error {
	return dara.Validate(s)
}

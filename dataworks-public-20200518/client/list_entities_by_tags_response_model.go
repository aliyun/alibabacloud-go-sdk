// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesByTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEntitiesByTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEntitiesByTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListEntitiesByTagsResponseBody) *ListEntitiesByTagsResponse
	GetBody() *ListEntitiesByTagsResponseBody
}

type ListEntitiesByTagsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEntitiesByTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEntitiesByTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesByTagsResponse) GoString() string {
	return s.String()
}

func (s *ListEntitiesByTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEntitiesByTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEntitiesByTagsResponse) GetBody() *ListEntitiesByTagsResponseBody {
	return s.Body
}

func (s *ListEntitiesByTagsResponse) SetHeaders(v map[string]*string) *ListEntitiesByTagsResponse {
	s.Headers = v
	return s
}

func (s *ListEntitiesByTagsResponse) SetStatusCode(v int32) *ListEntitiesByTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEntitiesByTagsResponse) SetBody(v *ListEntitiesByTagsResponseBody) *ListEntitiesByTagsResponse {
	s.Body = v
	return s
}

func (s *ListEntitiesByTagsResponse) Validate() error {
	return dara.Validate(s)
}

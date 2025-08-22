// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagDcdnResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagDcdnResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagDcdnResourcesResponse
	GetStatusCode() *int32
	SetBody(v *TagDcdnResourcesResponseBody) *TagDcdnResourcesResponse
	GetBody() *TagDcdnResourcesResponseBody
}

type TagDcdnResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagDcdnResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagDcdnResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s TagDcdnResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagDcdnResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagDcdnResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagDcdnResourcesResponse) GetBody() *TagDcdnResourcesResponseBody {
	return s.Body
}

func (s *TagDcdnResourcesResponse) SetHeaders(v map[string]*string) *TagDcdnResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagDcdnResourcesResponse) SetStatusCode(v int32) *TagDcdnResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagDcdnResourcesResponse) SetBody(v *TagDcdnResourcesResponseBody) *TagDcdnResourcesResponse {
	s.Body = v
	return s
}

func (s *TagDcdnResourcesResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagDcdnResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UntagDcdnResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UntagDcdnResourcesResponse
	GetStatusCode() *int32
	SetBody(v *UntagDcdnResourcesResponseBody) *UntagDcdnResourcesResponse
	GetBody() *UntagDcdnResourcesResponseBody
}

type UntagDcdnResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagDcdnResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagDcdnResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s UntagDcdnResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagDcdnResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UntagDcdnResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UntagDcdnResourcesResponse) GetBody() *UntagDcdnResourcesResponseBody {
	return s.Body
}

func (s *UntagDcdnResourcesResponse) SetHeaders(v map[string]*string) *UntagDcdnResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagDcdnResourcesResponse) SetStatusCode(v int32) *UntagDcdnResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagDcdnResourcesResponse) SetBody(v *UntagDcdnResourcesResponseBody) *UntagDcdnResourcesResponse {
	s.Body = v
	return s
}

func (s *UntagDcdnResourcesResponse) Validate() error {
	return dara.Validate(s)
}

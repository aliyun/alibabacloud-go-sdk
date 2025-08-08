// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShareResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListShareResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListShareResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListShareResourcesResponseBody) *ListShareResourcesResponse
	GetBody() *ListShareResourcesResponseBody
}

type ListShareResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListShareResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListShareResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListShareResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListShareResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListShareResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListShareResourcesResponse) GetBody() *ListShareResourcesResponseBody {
	return s.Body
}

func (s *ListShareResourcesResponse) SetHeaders(v map[string]*string) *ListShareResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListShareResourcesResponse) SetStatusCode(v int32) *ListShareResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShareResourcesResponse) SetBody(v *ListShareResourcesResponseBody) *ListShareResourcesResponse {
	s.Body = v
	return s
}

func (s *ListShareResourcesResponse) Validate() error {
	return dara.Validate(s)
}

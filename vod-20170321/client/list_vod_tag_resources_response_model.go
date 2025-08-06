// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVodTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVodTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListVodTagResourcesResponseBody) *ListVodTagResourcesResponse
	GetBody() *ListVodTagResourcesResponseBody
}

type ListVodTagResourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVodTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVodTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVodTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListVodTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVodTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVodTagResourcesResponse) GetBody() *ListVodTagResourcesResponseBody {
	return s.Body
}

func (s *ListVodTagResourcesResponse) SetHeaders(v map[string]*string) *ListVodTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListVodTagResourcesResponse) SetStatusCode(v int32) *ListVodTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVodTagResourcesResponse) SetBody(v *ListVodTagResourcesResponseBody) *ListVodTagResourcesResponse {
	s.Body = v
	return s
}

func (s *ListVodTagResourcesResponse) Validate() error {
	return dara.Validate(s)
}

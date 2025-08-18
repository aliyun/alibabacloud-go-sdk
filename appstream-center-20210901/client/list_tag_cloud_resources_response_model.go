// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagCloudResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagCloudResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagCloudResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListTagCloudResourcesResponseBody) *ListTagCloudResourcesResponse
	GetBody() *ListTagCloudResourcesResponseBody
}

type ListTagCloudResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagCloudResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagCloudResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagCloudResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagCloudResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagCloudResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagCloudResourcesResponse) GetBody() *ListTagCloudResourcesResponseBody {
	return s.Body
}

func (s *ListTagCloudResourcesResponse) SetHeaders(v map[string]*string) *ListTagCloudResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagCloudResourcesResponse) SetStatusCode(v int32) *ListTagCloudResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagCloudResourcesResponse) SetBody(v *ListTagCloudResourcesResponseBody) *ListTagCloudResourcesResponse {
	s.Body = v
	return s
}

func (s *ListTagCloudResourcesResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableResourcesResponseBody) *DescribeAvailableResourcesResponse
	GetBody() *DescribeAvailableResourcesResponseBody
}

type DescribeAvailableResourcesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableResourcesResponse) GetBody() *DescribeAvailableResourcesResponseBody {
	return s.Body
}

func (s *DescribeAvailableResourcesResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourcesResponse) SetStatusCode(v int32) *DescribeAvailableResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableResourcesResponse) SetBody(v *DescribeAvailableResourcesResponseBody) *DescribeAvailableResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableResourcesResponse) Validate() error {
	return dara.Validate(s)
}

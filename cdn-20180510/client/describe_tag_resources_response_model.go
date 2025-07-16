// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagResourcesResponseBody) *DescribeTagResourcesResponse
	GetBody() *DescribeTagResourcesResponseBody
}

type DescribeTagResourcesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagResourcesResponse) GetBody() *DescribeTagResourcesResponseBody {
	return s.Body
}

func (s *DescribeTagResourcesResponse) SetHeaders(v map[string]*string) *DescribeTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagResourcesResponse) SetStatusCode(v int32) *DescribeTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagResourcesResponse) SetBody(v *DescribeTagResourcesResponseBody) *DescribeTagResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeTagResourcesResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceSourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceSourcesResponseBody) *DescribeInstanceSourcesResponse
	GetBody() *DescribeInstanceSourcesResponseBody
}

type DescribeInstanceSourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceSourcesResponse) GetBody() *DescribeInstanceSourcesResponseBody {
	return s.Body
}

func (s *DescribeInstanceSourcesResponse) SetHeaders(v map[string]*string) *DescribeInstanceSourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSourcesResponse) SetStatusCode(v int32) *DescribeInstanceSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSourcesResponse) SetBody(v *DescribeInstanceSourcesResponseBody) *DescribeInstanceSourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceSourcesResponse) Validate() error {
	return dara.Validate(s)
}

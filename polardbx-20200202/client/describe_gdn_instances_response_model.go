// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGdnInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGdnInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGdnInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGdnInstancesResponseBody) *DescribeGdnInstancesResponse
	GetBody() *DescribeGdnInstancesResponseBody
}

type DescribeGdnInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGdnInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGdnInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGdnInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGdnInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGdnInstancesResponse) GetBody() *DescribeGdnInstancesResponseBody {
	return s.Body
}

func (s *DescribeGdnInstancesResponse) SetHeaders(v map[string]*string) *DescribeGdnInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGdnInstancesResponse) SetStatusCode(v int32) *DescribeGdnInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGdnInstancesResponse) SetBody(v *DescribeGdnInstancesResponseBody) *DescribeGdnInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeGdnInstancesResponse) Validate() error {
	return dara.Validate(s)
}

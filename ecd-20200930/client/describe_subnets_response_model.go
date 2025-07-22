// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubnetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSubnetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSubnetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSubnetsResponseBody) *DescribeSubnetsResponse
	GetBody() *DescribeSubnetsResponseBody
}

type DescribeSubnetsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSubnetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSubnetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubnetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubnetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSubnetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSubnetsResponse) GetBody() *DescribeSubnetsResponseBody {
	return s.Body
}

func (s *DescribeSubnetsResponse) SetHeaders(v map[string]*string) *DescribeSubnetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubnetsResponse) SetStatusCode(v int32) *DescribeSubnetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubnetsResponse) SetBody(v *DescribeSubnetsResponseBody) *DescribeSubnetsResponse {
	s.Body = v
	return s
}

func (s *DescribeSubnetsResponse) Validate() error {
	return dara.Validate(s)
}

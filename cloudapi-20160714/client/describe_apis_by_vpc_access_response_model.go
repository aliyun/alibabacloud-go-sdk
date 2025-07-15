// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisByVpcAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisByVpcAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisByVpcAccessResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisByVpcAccessResponseBody) *DescribeApisByVpcAccessResponse
	GetBody() *DescribeApisByVpcAccessResponseBody
}

type DescribeApisByVpcAccessResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisByVpcAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisByVpcAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisByVpcAccessResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisByVpcAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisByVpcAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisByVpcAccessResponse) GetBody() *DescribeApisByVpcAccessResponseBody {
	return s.Body
}

func (s *DescribeApisByVpcAccessResponse) SetHeaders(v map[string]*string) *DescribeApisByVpcAccessResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisByVpcAccessResponse) SetStatusCode(v int32) *DescribeApisByVpcAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisByVpcAccessResponse) SetBody(v *DescribeApisByVpcAccessResponseBody) *DescribeApisByVpcAccessResponse {
	s.Body = v
	return s
}

func (s *DescribeApisByVpcAccessResponse) Validate() error {
	return dara.Validate(s)
}

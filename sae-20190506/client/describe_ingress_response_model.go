// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIngressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIngressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIngressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIngressResponseBody) *DescribeIngressResponse
	GetBody() *DescribeIngressResponseBody
}

type DescribeIngressResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIngressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIngressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIngressResponse) GoString() string {
	return s.String()
}

func (s *DescribeIngressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIngressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIngressResponse) GetBody() *DescribeIngressResponseBody {
	return s.Body
}

func (s *DescribeIngressResponse) SetHeaders(v map[string]*string) *DescribeIngressResponse {
	s.Headers = v
	return s
}

func (s *DescribeIngressResponse) SetStatusCode(v int32) *DescribeIngressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIngressResponse) SetBody(v *DescribeIngressResponseBody) *DescribeIngressResponse {
	s.Body = v
	return s
}

func (s *DescribeIngressResponse) Validate() error {
	return dara.Validate(s)
}

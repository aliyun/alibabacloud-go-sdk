// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpLocationServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpLocationServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpLocationServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpLocationServiceResponseBody) *DescribeIpLocationServiceResponse
	GetBody() *DescribeIpLocationServiceResponseBody
}

type DescribeIpLocationServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpLocationServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpLocationServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpLocationServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpLocationServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpLocationServiceResponse) GetBody() *DescribeIpLocationServiceResponseBody {
	return s.Body
}

func (s *DescribeIpLocationServiceResponse) SetHeaders(v map[string]*string) *DescribeIpLocationServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpLocationServiceResponse) SetStatusCode(v int32) *DescribeIpLocationServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpLocationServiceResponse) SetBody(v *DescribeIpLocationServiceResponseBody) *DescribeIpLocationServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeIpLocationServiceResponse) Validate() error {
	return dara.Validate(s)
}

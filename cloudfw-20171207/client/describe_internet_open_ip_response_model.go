// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetOpenIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetOpenIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetOpenIpResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetOpenIpResponseBody) *DescribeInternetOpenIpResponse
	GetBody() *DescribeInternetOpenIpResponseBody
}

type DescribeInternetOpenIpResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetOpenIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetOpenIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetOpenIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetOpenIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetOpenIpResponse) GetBody() *DescribeInternetOpenIpResponseBody {
	return s.Body
}

func (s *DescribeInternetOpenIpResponse) SetHeaders(v map[string]*string) *DescribeInternetOpenIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetOpenIpResponse) SetStatusCode(v int32) *DescribeInternetOpenIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetOpenIpResponse) SetBody(v *DescribeInternetOpenIpResponseBody) *DescribeInternetOpenIpResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetOpenIpResponse) Validate() error {
	return dara.Validate(s)
}

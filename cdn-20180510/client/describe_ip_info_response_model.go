// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIpInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIpInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIpInfoResponseBody) *DescribeIpInfoResponse
	GetBody() *DescribeIpInfoResponseBody
}

type DescribeIpInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIpInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIpInfoResponse) GetBody() *DescribeIpInfoResponseBody {
	return s.Body
}

func (s *DescribeIpInfoResponse) SetHeaders(v map[string]*string) *DescribeIpInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpInfoResponse) SetStatusCode(v int32) *DescribeIpInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpInfoResponse) SetBody(v *DescribeIpInfoResponseBody) *DescribeIpInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeIpInfoResponse) Validate() error {
	return dara.Validate(s)
}

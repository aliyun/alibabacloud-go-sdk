// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserConnectTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserConnectTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserConnectTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserConnectTimeResponseBody) *DescribeUserConnectTimeResponse
	GetBody() *DescribeUserConnectTimeResponseBody
}

type DescribeUserConnectTimeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserConnectTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserConnectTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConnectTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserConnectTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserConnectTimeResponse) GetBody() *DescribeUserConnectTimeResponseBody {
	return s.Body
}

func (s *DescribeUserConnectTimeResponse) SetHeaders(v map[string]*string) *DescribeUserConnectTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserConnectTimeResponse) SetStatusCode(v int32) *DescribeUserConnectTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserConnectTimeResponse) SetBody(v *DescribeUserConnectTimeResponseBody) *DescribeUserConnectTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeUserConnectTimeResponse) Validate() error {
	return dara.Validate(s)
}

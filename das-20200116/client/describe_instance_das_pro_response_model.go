// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDasProResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceDasProResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceDasProResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceDasProResponseBody) *DescribeInstanceDasProResponse
	GetBody() *DescribeInstanceDasProResponseBody
}

type DescribeInstanceDasProResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceDasProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceDasProResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDasProResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDasProResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceDasProResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceDasProResponse) GetBody() *DescribeInstanceDasProResponseBody {
	return s.Body
}

func (s *DescribeInstanceDasProResponse) SetHeaders(v map[string]*string) *DescribeInstanceDasProResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDasProResponse) SetStatusCode(v int32) *DescribeInstanceDasProResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceDasProResponse) SetBody(v *DescribeInstanceDasProResponseBody) *DescribeInstanceDasProResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceDasProResponse) Validate() error {
	return dara.Validate(s)
}

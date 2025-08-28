// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAuthInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAuthInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAuthInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAuthInfoResponseBody) *DescribeInstanceAuthInfoResponse
	GetBody() *DescribeInstanceAuthInfoResponseBody
}

type DescribeInstanceAuthInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAuthInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAuthInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAuthInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAuthInfoResponse) GetBody() *DescribeInstanceAuthInfoResponseBody {
	return s.Body
}

func (s *DescribeInstanceAuthInfoResponse) SetHeaders(v map[string]*string) *DescribeInstanceAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAuthInfoResponse) SetStatusCode(v int32) *DescribeInstanceAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAuthInfoResponse) SetBody(v *DescribeInstanceAuthInfoResponseBody) *DescribeInstanceAuthInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAuthInfoResponse) Validate() error {
	return dara.Validate(s)
}

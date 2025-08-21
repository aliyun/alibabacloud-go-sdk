// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceSpecResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceSpecResponseBody) *DescribeInstanceSpecResponse
	GetBody() *DescribeInstanceSpecResponseBody
}

type DescribeInstanceSpecResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceSpecResponse) GetBody() *DescribeInstanceSpecResponseBody {
	return s.Body
}

func (s *DescribeInstanceSpecResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecResponse) SetStatusCode(v int32) *DescribeInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSpecResponse) SetBody(v *DescribeInstanceSpecResponseBody) *DescribeInstanceSpecResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceSpecResponse) Validate() error {
	return dara.Validate(s)
}

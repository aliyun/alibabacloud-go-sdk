// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceAccountsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceAccountsResponseBody) *DescribeInstanceAccountsResponse
	GetBody() *DescribeInstanceAccountsResponseBody
}

type DescribeInstanceAccountsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceAccountsResponse) GetBody() *DescribeInstanceAccountsResponseBody {
	return s.Body
}

func (s *DescribeInstanceAccountsResponse) SetHeaders(v map[string]*string) *DescribeInstanceAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAccountsResponse) SetStatusCode(v int32) *DescribeInstanceAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAccountsResponse) SetBody(v *DescribeInstanceAccountsResponseBody) *DescribeInstanceAccountsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceAccountsResponse) Validate() error {
	return dara.Validate(s)
}

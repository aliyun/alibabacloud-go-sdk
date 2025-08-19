// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifyTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifyTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifyTokenResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifyTokenResponseBody) *DescribeVerifyTokenResponse
	GetBody() *DescribeVerifyTokenResponseBody
}

type DescribeVerifyTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifyTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifyTokenResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifyTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifyTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifyTokenResponse) GetBody() *DescribeVerifyTokenResponseBody {
	return s.Body
}

func (s *DescribeVerifyTokenResponse) SetHeaders(v map[string]*string) *DescribeVerifyTokenResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifyTokenResponse) SetStatusCode(v int32) *DescribeVerifyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifyTokenResponse) SetBody(v *DescribeVerifyTokenResponseBody) *DescribeVerifyTokenResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifyTokenResponse) Validate() error {
	return dara.Validate(s)
}

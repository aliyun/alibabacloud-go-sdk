// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCardVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCardVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCardVerifyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCardVerifyResponseBody) *DescribeCardVerifyResponse
	GetBody() *DescribeCardVerifyResponseBody
}

type DescribeCardVerifyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCardVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCardVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCardVerifyResponse) GoString() string {
	return s.String()
}

func (s *DescribeCardVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCardVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCardVerifyResponse) GetBody() *DescribeCardVerifyResponseBody {
	return s.Body
}

func (s *DescribeCardVerifyResponse) SetHeaders(v map[string]*string) *DescribeCardVerifyResponse {
	s.Headers = v
	return s
}

func (s *DescribeCardVerifyResponse) SetStatusCode(v int32) *DescribeCardVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCardVerifyResponse) SetBody(v *DescribeCardVerifyResponseBody) *DescribeCardVerifyResponse {
	s.Body = v
	return s
}

func (s *DescribeCardVerifyResponse) Validate() error {
	return dara.Validate(s)
}

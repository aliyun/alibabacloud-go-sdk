// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListAuthorizeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhiteListAuthorizeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhiteListAuthorizeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhiteListAuthorizeResponseBody) *DescribeWhiteListAuthorizeResponse
	GetBody() *DescribeWhiteListAuthorizeResponseBody
}

type DescribeWhiteListAuthorizeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhiteListAuthorizeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhiteListAuthorizeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListAuthorizeResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListAuthorizeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhiteListAuthorizeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhiteListAuthorizeResponse) GetBody() *DescribeWhiteListAuthorizeResponseBody {
	return s.Body
}

func (s *DescribeWhiteListAuthorizeResponse) SetHeaders(v map[string]*string) *DescribeWhiteListAuthorizeResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteListAuthorizeResponse) SetStatusCode(v int32) *DescribeWhiteListAuthorizeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhiteListAuthorizeResponse) SetBody(v *DescribeWhiteListAuthorizeResponseBody) *DescribeWhiteListAuthorizeResponse {
	s.Body = v
	return s
}

func (s *DescribeWhiteListAuthorizeResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifySchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifySchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifySchemeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifySchemeResponseBody) *DescribeVerifySchemeResponse
	GetBody() *DescribeVerifySchemeResponseBody
}

type DescribeVerifySchemeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifySchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifySchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifySchemeResponse) GetBody() *DescribeVerifySchemeResponseBody {
	return s.Body
}

func (s *DescribeVerifySchemeResponse) SetHeaders(v map[string]*string) *DescribeVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifySchemeResponse) SetStatusCode(v int32) *DescribeVerifySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifySchemeResponse) SetBody(v *DescribeVerifySchemeResponseBody) *DescribeVerifySchemeResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifySchemeResponse) Validate() error {
	return dara.Validate(s)
}

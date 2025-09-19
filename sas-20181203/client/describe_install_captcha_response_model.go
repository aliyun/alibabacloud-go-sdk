// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstallCaptchaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstallCaptchaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstallCaptchaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstallCaptchaResponseBody) *DescribeInstallCaptchaResponse
	GetBody() *DescribeInstallCaptchaResponseBody
}

type DescribeInstallCaptchaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstallCaptchaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstallCaptchaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstallCaptchaResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstallCaptchaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstallCaptchaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstallCaptchaResponse) GetBody() *DescribeInstallCaptchaResponseBody {
	return s.Body
}

func (s *DescribeInstallCaptchaResponse) SetHeaders(v map[string]*string) *DescribeInstallCaptchaResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstallCaptchaResponse) SetStatusCode(v int32) *DescribeInstallCaptchaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstallCaptchaResponse) SetBody(v *DescribeInstallCaptchaResponseBody) *DescribeInstallCaptchaResponse {
	s.Body = v
	return s
}

func (s *DescribeInstallCaptchaResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserCertificateExpireCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserCertificateExpireCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserCertificateExpireCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserCertificateExpireCountResponseBody) *DescribeUserCertificateExpireCountResponse
	GetBody() *DescribeUserCertificateExpireCountResponseBody
}

type DescribeUserCertificateExpireCountResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserCertificateExpireCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserCertificateExpireCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserCertificateExpireCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserCertificateExpireCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserCertificateExpireCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserCertificateExpireCountResponse) GetBody() *DescribeUserCertificateExpireCountResponseBody {
	return s.Body
}

func (s *DescribeUserCertificateExpireCountResponse) SetHeaders(v map[string]*string) *DescribeUserCertificateExpireCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserCertificateExpireCountResponse) SetStatusCode(v int32) *DescribeUserCertificateExpireCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserCertificateExpireCountResponse) SetBody(v *DescribeUserCertificateExpireCountResponseBody) *DescribeUserCertificateExpireCountResponse {
	s.Body = v
	return s
}

func (s *DescribeUserCertificateExpireCountResponse) Validate() error {
	return dara.Validate(s)
}

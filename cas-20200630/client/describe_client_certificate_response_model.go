// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientCertificateResponseBody) *DescribeClientCertificateResponse
	GetBody() *DescribeClientCertificateResponseBody
}

type DescribeClientCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientCertificateResponse) GetBody() *DescribeClientCertificateResponseBody {
	return s.Body
}

func (s *DescribeClientCertificateResponse) SetHeaders(v map[string]*string) *DescribeClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientCertificateResponse) SetStatusCode(v int32) *DescribeClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientCertificateResponse) SetBody(v *DescribeClientCertificateResponseBody) *DescribeClientCertificateResponse {
	s.Body = v
	return s
}

func (s *DescribeClientCertificateResponse) Validate() error {
	return dara.Validate(s)
}

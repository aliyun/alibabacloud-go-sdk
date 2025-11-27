// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainCertificateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainCertificateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainCertificateInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainCertificateInfoResponseBody) *DescribeVsDomainCertificateInfoResponse
	GetBody() *DescribeVsDomainCertificateInfoResponseBody
}

type DescribeVsDomainCertificateInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainCertificateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainCertificateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainCertificateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainCertificateInfoResponse) GetBody() *DescribeVsDomainCertificateInfoResponseBody {
	return s.Body
}

func (s *DescribeVsDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeVsDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponse) SetStatusCode(v int32) *DescribeVsDomainCertificateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponse) SetBody(v *DescribeVsDomainCertificateInfoResponseBody) *DescribeVsDomainCertificateInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainCertificateInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

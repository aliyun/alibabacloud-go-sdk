// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainCertificateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainCertificateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainCertificateInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainCertificateInfoResponseBody) *DescribeVodDomainCertificateInfoResponse
	GetBody() *DescribeVodDomainCertificateInfoResponseBody
}

type DescribeVodDomainCertificateInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainCertificateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainCertificateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainCertificateInfoResponse) GetBody() *DescribeVodDomainCertificateInfoResponseBody {
	return s.Body
}

func (s *DescribeVodDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeVodDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponse) SetStatusCode(v int32) *DescribeVodDomainCertificateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponse) SetBody(v *DescribeVodDomainCertificateInfoResponseBody) *DescribeVodDomainCertificateInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

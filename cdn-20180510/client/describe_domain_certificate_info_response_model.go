// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCertificateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainCertificateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainCertificateInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainCertificateInfoResponseBody) *DescribeDomainCertificateInfoResponse
	GetBody() *DescribeDomainCertificateInfoResponseBody
}

type DescribeDomainCertificateInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainCertificateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainCertificateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainCertificateInfoResponse) GetBody() *DescribeDomainCertificateInfoResponseBody {
	return s.Body
}

func (s *DescribeDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCertificateInfoResponse) SetStatusCode(v int32) *DescribeDomainCertificateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponse) SetBody(v *DescribeDomainCertificateInfoResponseBody) *DescribeDomainCertificateInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainCertificateInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

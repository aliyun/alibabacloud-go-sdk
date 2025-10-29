// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainCertificateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainCertificateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainCertificateInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainCertificateInfoResponseBody) *DescribeLiveDomainCertificateInfoResponse
	GetBody() *DescribeLiveDomainCertificateInfoResponseBody
}

type DescribeLiveDomainCertificateInfoResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainCertificateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainCertificateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainCertificateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainCertificateInfoResponse) GetBody() *DescribeLiveDomainCertificateInfoResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainCertificateInfoResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainCertificateInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponse) SetStatusCode(v int32) *DescribeLiveDomainCertificateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponse) SetBody(v *DescribeLiveDomainCertificateInfoResponseBody) *DescribeLiveDomainCertificateInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

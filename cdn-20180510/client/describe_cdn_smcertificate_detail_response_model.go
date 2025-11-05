// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSMCertificateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnSMCertificateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnSMCertificateDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnSMCertificateDetailResponseBody) *DescribeCdnSMCertificateDetailResponse
	GetBody() *DescribeCdnSMCertificateDetailResponseBody
}

type DescribeCdnSMCertificateDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnSMCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnSMCertificateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSMCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnSMCertificateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnSMCertificateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnSMCertificateDetailResponse) GetBody() *DescribeCdnSMCertificateDetailResponseBody {
	return s.Body
}

func (s *DescribeCdnSMCertificateDetailResponse) SetHeaders(v map[string]*string) *DescribeCdnSMCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponse) SetStatusCode(v int32) *DescribeCdnSMCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponse) SetBody(v *DescribeCdnSMCertificateDetailResponseBody) *DescribeCdnSMCertificateDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnSMCertificateDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

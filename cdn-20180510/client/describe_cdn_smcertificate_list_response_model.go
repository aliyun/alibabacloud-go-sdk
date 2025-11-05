// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSMCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnSMCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnSMCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnSMCertificateListResponseBody) *DescribeCdnSMCertificateListResponse
	GetBody() *DescribeCdnSMCertificateListResponseBody
}

type DescribeCdnSMCertificateListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnSMCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnSMCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSMCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnSMCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnSMCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnSMCertificateListResponse) GetBody() *DescribeCdnSMCertificateListResponseBody {
	return s.Body
}

func (s *DescribeCdnSMCertificateListResponse) SetHeaders(v map[string]*string) *DescribeCdnSMCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnSMCertificateListResponse) SetStatusCode(v int32) *DescribeCdnSMCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnSMCertificateListResponse) SetBody(v *DescribeCdnSMCertificateListResponseBody) *DescribeCdnSMCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnSMCertificateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

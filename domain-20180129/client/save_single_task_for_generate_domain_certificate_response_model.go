// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForGenerateDomainCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveSingleTaskForGenerateDomainCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveSingleTaskForGenerateDomainCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SaveSingleTaskForGenerateDomainCertificateResponseBody) *SaveSingleTaskForGenerateDomainCertificateResponse
	GetBody() *SaveSingleTaskForGenerateDomainCertificateResponseBody
}

type SaveSingleTaskForGenerateDomainCertificateResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSingleTaskForGenerateDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveSingleTaskForGenerateDomainCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForGenerateDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponse) GetBody() *SaveSingleTaskForGenerateDomainCertificateResponseBody {
	return s.Body
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponse) SetHeaders(v map[string]*string) *SaveSingleTaskForGenerateDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponse) SetStatusCode(v int32) *SaveSingleTaskForGenerateDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponse) SetBody(v *SaveSingleTaskForGenerateDomainCertificateResponseBody) *SaveSingleTaskForGenerateDomainCertificateResponse {
	s.Body = v
	return s
}

func (s *SaveSingleTaskForGenerateDomainCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

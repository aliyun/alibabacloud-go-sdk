// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForGenerateDomainCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveBatchTaskForGenerateDomainCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveBatchTaskForGenerateDomainCertificateResponse
	GetStatusCode() *int32
	SetBody(v *SaveBatchTaskForGenerateDomainCertificateResponseBody) *SaveBatchTaskForGenerateDomainCertificateResponse
	GetBody() *SaveBatchTaskForGenerateDomainCertificateResponseBody
}

type SaveBatchTaskForGenerateDomainCertificateResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveBatchTaskForGenerateDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveBatchTaskForGenerateDomainCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForGenerateDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponse) GetBody() *SaveBatchTaskForGenerateDomainCertificateResponseBody {
	return s.Body
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponse) SetHeaders(v map[string]*string) *SaveBatchTaskForGenerateDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponse) SetStatusCode(v int32) *SaveBatchTaskForGenerateDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponse) SetBody(v *SaveBatchTaskForGenerateDomainCertificateResponseBody) *SaveBatchTaskForGenerateDomainCertificateResponse {
	s.Body = v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

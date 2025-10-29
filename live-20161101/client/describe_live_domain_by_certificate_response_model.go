// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainByCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainByCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainByCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainByCertificateResponseBody) *DescribeLiveDomainByCertificateResponse
	GetBody() *DescribeLiveDomainByCertificateResponseBody
}

type DescribeLiveDomainByCertificateResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainByCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainByCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainByCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainByCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainByCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainByCertificateResponse) GetBody() *DescribeLiveDomainByCertificateResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainByCertificateResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainByCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainByCertificateResponse) SetStatusCode(v int32) *DescribeLiveDomainByCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponse) SetBody(v *DescribeLiveDomainByCertificateResponseBody) *DescribeLiveDomainByCertificateResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainByCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

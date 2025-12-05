// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCertificateResponseBody) *DescribeCertificateResponse
	GetBody() *DescribeCertificateResponseBody
}

type DescribeCertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCertificateResponse) GetBody() *DescribeCertificateResponseBody {
	return s.Body
}

func (s *DescribeCertificateResponse) SetHeaders(v map[string]*string) *DescribeCertificateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificateResponse) SetStatusCode(v int32) *DescribeCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertificateResponse) SetBody(v *DescribeCertificateResponseBody) *DescribeCertificateResponse {
	s.Body = v
	return s
}

func (s *DescribeCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

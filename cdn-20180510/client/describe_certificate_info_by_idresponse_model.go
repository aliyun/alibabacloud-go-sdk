// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertificateInfoByIDResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCertificateInfoByIDResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCertificateInfoByIDResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCertificateInfoByIDResponseBody) *DescribeCertificateInfoByIDResponse
	GetBody() *DescribeCertificateInfoByIDResponseBody
}

type DescribeCertificateInfoByIDResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertificateInfoByIDResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertificateInfoByIDResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertificateInfoByIDResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificateInfoByIDResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCertificateInfoByIDResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCertificateInfoByIDResponse) GetBody() *DescribeCertificateInfoByIDResponseBody {
	return s.Body
}

func (s *DescribeCertificateInfoByIDResponse) SetHeaders(v map[string]*string) *DescribeCertificateInfoByIDResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificateInfoByIDResponse) SetStatusCode(v int32) *DescribeCertificateInfoByIDResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertificateInfoByIDResponse) SetBody(v *DescribeCertificateInfoByIDResponseBody) *DescribeCertificateInfoByIDResponse {
	s.Body = v
	return s
}

func (s *DescribeCertificateInfoByIDResponse) Validate() error {
	return dara.Validate(s)
}

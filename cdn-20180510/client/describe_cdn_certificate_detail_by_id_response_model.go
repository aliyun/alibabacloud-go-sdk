// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnCertificateDetailByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnCertificateDetailByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnCertificateDetailByIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnCertificateDetailByIdResponseBody) *DescribeCdnCertificateDetailByIdResponse
	GetBody() *DescribeCdnCertificateDetailByIdResponseBody
}

type DescribeCdnCertificateDetailByIdResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnCertificateDetailByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnCertificateDetailByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateDetailByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateDetailByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnCertificateDetailByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnCertificateDetailByIdResponse) GetBody() *DescribeCdnCertificateDetailByIdResponseBody {
	return s.Body
}

func (s *DescribeCdnCertificateDetailByIdResponse) SetHeaders(v map[string]*string) *DescribeCdnCertificateDetailByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnCertificateDetailByIdResponse) SetStatusCode(v int32) *DescribeCdnCertificateDetailByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnCertificateDetailByIdResponse) SetBody(v *DescribeCdnCertificateDetailByIdResponseBody) *DescribeCdnCertificateDetailByIdResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnCertificateDetailByIdResponse) Validate() error {
	return dara.Validate(s)
}

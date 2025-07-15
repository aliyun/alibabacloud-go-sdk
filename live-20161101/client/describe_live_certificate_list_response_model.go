// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveCertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveCertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveCertificateListResponseBody) *DescribeLiveCertificateListResponse
	GetBody() *DescribeLiveCertificateListResponseBody
}

type DescribeLiveCertificateListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveCertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveCertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveCertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveCertificateListResponse) GetBody() *DescribeLiveCertificateListResponseBody {
	return s.Body
}

func (s *DescribeLiveCertificateListResponse) SetHeaders(v map[string]*string) *DescribeLiveCertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveCertificateListResponse) SetStatusCode(v int32) *DescribeLiveCertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveCertificateListResponse) SetBody(v *DescribeLiveCertificateListResponseBody) *DescribeLiveCertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveCertificateListResponse) Validate() error {
	return dara.Validate(s)
}

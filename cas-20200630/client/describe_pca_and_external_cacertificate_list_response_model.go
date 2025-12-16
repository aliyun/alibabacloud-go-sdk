// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePcaAndExternalCACertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePcaAndExternalCACertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePcaAndExternalCACertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePcaAndExternalCACertificateListResponseBody) *DescribePcaAndExternalCACertificateListResponse
	GetBody() *DescribePcaAndExternalCACertificateListResponseBody
}

type DescribePcaAndExternalCACertificateListResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePcaAndExternalCACertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePcaAndExternalCACertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePcaAndExternalCACertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribePcaAndExternalCACertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePcaAndExternalCACertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePcaAndExternalCACertificateListResponse) GetBody() *DescribePcaAndExternalCACertificateListResponseBody {
	return s.Body
}

func (s *DescribePcaAndExternalCACertificateListResponse) SetHeaders(v map[string]*string) *DescribePcaAndExternalCACertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponse) SetStatusCode(v int32) *DescribePcaAndExternalCACertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponse) SetBody(v *DescribePcaAndExternalCACertificateListResponseBody) *DescribePcaAndExternalCACertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribePcaAndExternalCACertificateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificateCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCACertificateCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCACertificateCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCACertificateCountResponseBody) *DescribeCACertificateCountResponse
	GetBody() *DescribeCACertificateCountResponseBody
}

type DescribeCACertificateCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCACertificateCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCACertificateCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCACertificateCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCACertificateCountResponse) GetBody() *DescribeCACertificateCountResponseBody {
	return s.Body
}

func (s *DescribeCACertificateCountResponse) SetHeaders(v map[string]*string) *DescribeCACertificateCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeCACertificateCountResponse) SetStatusCode(v int32) *DescribeCACertificateCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCACertificateCountResponse) SetBody(v *DescribeCACertificateCountResponseBody) *DescribeCACertificateCountResponse {
	s.Body = v
	return s
}

func (s *DescribeCACertificateCountResponse) Validate() error {
	return dara.Validate(s)
}

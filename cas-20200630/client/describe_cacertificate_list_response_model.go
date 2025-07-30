// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCACertificateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCACertificateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCACertificateListResponseBody) *DescribeCACertificateListResponse
	GetBody() *DescribeCACertificateListResponseBody
}

type DescribeCACertificateListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCACertificateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCACertificateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCACertificateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCACertificateListResponse) GetBody() *DescribeCACertificateListResponseBody {
	return s.Body
}

func (s *DescribeCACertificateListResponse) SetHeaders(v map[string]*string) *DescribeCACertificateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCACertificateListResponse) SetStatusCode(v int32) *DescribeCACertificateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCACertificateListResponse) SetBody(v *DescribeCACertificateListResponseBody) *DescribeCACertificateListResponse {
	s.Body = v
	return s
}

func (s *DescribeCACertificateListResponse) Validate() error {
	return dara.Validate(s)
}

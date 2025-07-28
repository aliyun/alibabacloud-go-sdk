// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCertDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCertDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCertDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCertDetailResponseBody) *DescribeCertDetailResponse
	GetBody() *DescribeCertDetailResponseBody
}

type DescribeCertDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCertDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCertDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCertDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCertDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCertDetailResponse) GetBody() *DescribeCertDetailResponseBody {
	return s.Body
}

func (s *DescribeCertDetailResponse) SetHeaders(v map[string]*string) *DescribeCertDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertDetailResponse) SetStatusCode(v int32) *DescribeCertDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCertDetailResponse) SetBody(v *DescribeCertDetailResponseBody) *DescribeCertDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeCertDetailResponse) Validate() error {
	return dara.Validate(s)
}

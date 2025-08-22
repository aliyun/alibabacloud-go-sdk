// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnHttpsDomainListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnHttpsDomainListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnHttpsDomainListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnHttpsDomainListResponseBody) *DescribeDcdnHttpsDomainListResponse
	GetBody() *DescribeDcdnHttpsDomainListResponseBody
}

type DescribeDcdnHttpsDomainListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnHttpsDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnHttpsDomainListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnHttpsDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnHttpsDomainListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnHttpsDomainListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnHttpsDomainListResponse) GetBody() *DescribeDcdnHttpsDomainListResponseBody {
	return s.Body
}

func (s *DescribeDcdnHttpsDomainListResponse) SetHeaders(v map[string]*string) *DescribeDcdnHttpsDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponse) SetStatusCode(v int32) *DescribeDcdnHttpsDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponse) SetBody(v *DescribeDcdnHttpsDomainListResponseBody) *DescribeDcdnHttpsDomainListResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnHttpsDomainListResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnHttpsDomainListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnHttpsDomainListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnHttpsDomainListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnHttpsDomainListResponseBody) *DescribeCdnHttpsDomainListResponse
	GetBody() *DescribeCdnHttpsDomainListResponseBody
}

type DescribeCdnHttpsDomainListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnHttpsDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnHttpsDomainListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnHttpsDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnHttpsDomainListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnHttpsDomainListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnHttpsDomainListResponse) GetBody() *DescribeCdnHttpsDomainListResponseBody {
	return s.Body
}

func (s *DescribeCdnHttpsDomainListResponse) SetHeaders(v map[string]*string) *DescribeCdnHttpsDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnHttpsDomainListResponse) SetStatusCode(v int32) *DescribeCdnHttpsDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnHttpsDomainListResponse) SetBody(v *DescribeCdnHttpsDomainListResponseBody) *DescribeCdnHttpsDomainListResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnHttpsDomainListResponse) Validate() error {
	return dara.Validate(s)
}

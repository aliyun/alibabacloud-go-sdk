// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveHttpsDomainListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveHttpsDomainListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveHttpsDomainListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveHttpsDomainListResponseBody) *DescribeLiveHttpsDomainListResponse
	GetBody() *DescribeLiveHttpsDomainListResponseBody
}

type DescribeLiveHttpsDomainListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveHttpsDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveHttpsDomainListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveHttpsDomainListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveHttpsDomainListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveHttpsDomainListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveHttpsDomainListResponse) GetBody() *DescribeLiveHttpsDomainListResponseBody {
	return s.Body
}

func (s *DescribeLiveHttpsDomainListResponse) SetHeaders(v map[string]*string) *DescribeLiveHttpsDomainListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveHttpsDomainListResponse) SetStatusCode(v int32) *DescribeLiveHttpsDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveHttpsDomainListResponse) SetBody(v *DescribeLiveHttpsDomainListResponseBody) *DescribeLiveHttpsDomainListResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveHttpsDomainListResponse) Validate() error {
	return dara.Validate(s)
}

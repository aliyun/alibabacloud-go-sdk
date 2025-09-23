// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQpsWithCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainQpsWithCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainQpsWithCacheResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainQpsWithCacheResponseBody) *DescribeDomainQpsWithCacheResponse
	GetBody() *DescribeDomainQpsWithCacheResponseBody
}

type DescribeDomainQpsWithCacheResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainQpsWithCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainQpsWithCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQpsWithCacheResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainQpsWithCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainQpsWithCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainQpsWithCacheResponse) GetBody() *DescribeDomainQpsWithCacheResponseBody {
	return s.Body
}

func (s *DescribeDomainQpsWithCacheResponse) SetHeaders(v map[string]*string) *DescribeDomainQpsWithCacheResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetStatusCode(v int32) *DescribeDomainQpsWithCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) SetBody(v *DescribeDomainQpsWithCacheResponseBody) *DescribeDomainQpsWithCacheResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainQpsWithCacheResponse) Validate() error {
	return dara.Validate(s)
}

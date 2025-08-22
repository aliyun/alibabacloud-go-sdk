// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDeletedDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDeletedDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDeletedDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDeletedDomainsResponseBody) *DescribeDcdnDeletedDomainsResponse
	GetBody() *DescribeDcdnDeletedDomainsResponseBody
}

type DescribeDcdnDeletedDomainsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDeletedDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDeletedDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDeletedDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDeletedDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDeletedDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDeletedDomainsResponse) GetBody() *DescribeDcdnDeletedDomainsResponseBody {
	return s.Body
}

func (s *DescribeDcdnDeletedDomainsResponse) SetHeaders(v map[string]*string) *DescribeDcdnDeletedDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponse) SetStatusCode(v int32) *DescribeDcdnDeletedDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponse) SetBody(v *DescribeDcdnDeletedDomainsResponseBody) *DescribeDcdnDeletedDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDeletedDomainsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafDomainsResponseBody) *DescribeDcdnWafDomainsResponse
	GetBody() *DescribeDcdnWafDomainsResponseBody
}

type DescribeDcdnWafDomainsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafDomainsResponse) GetBody() *DescribeDcdnWafDomainsResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafDomainsResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafDomainsResponse) SetStatusCode(v int32) *DescribeDcdnWafDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafDomainsResponse) SetBody(v *DescribeDcdnWafDomainsResponseBody) *DescribeDcdnWafDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafDomainsResponse) Validate() error {
	return dara.Validate(s)
}

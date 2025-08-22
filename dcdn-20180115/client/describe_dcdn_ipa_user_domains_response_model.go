// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaUserDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnIpaUserDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnIpaUserDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnIpaUserDomainsResponseBody) *DescribeDcdnIpaUserDomainsResponse
	GetBody() *DescribeDcdnIpaUserDomainsResponseBody
}

type DescribeDcdnIpaUserDomainsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnIpaUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnIpaUserDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaUserDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnIpaUserDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnIpaUserDomainsResponse) GetBody() *DescribeDcdnIpaUserDomainsResponseBody {
	return s.Body
}

func (s *DescribeDcdnIpaUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpaUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponse) SetStatusCode(v int32) *DescribeDcdnIpaUserDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponse) SetBody(v *DescribeDcdnIpaUserDomainsResponseBody) *DescribeDcdnIpaUserDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnIpaUserDomainsResponse) Validate() error {
	return dara.Validate(s)
}

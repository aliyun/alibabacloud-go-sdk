// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserDomainsResponseBody) *DescribeDcdnUserDomainsResponse
	GetBody() *DescribeDcdnUserDomainsResponseBody
}

type DescribeDcdnUserDomainsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserDomainsResponse) GetBody() *DescribeDcdnUserDomainsResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserDomainsResponse) SetStatusCode(v int32) *DescribeDcdnUserDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserDomainsResponse) SetBody(v *DescribeDcdnUserDomainsResponseBody) *DescribeDcdnUserDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserDomainsResponse) Validate() error {
	return dara.Validate(s)
}

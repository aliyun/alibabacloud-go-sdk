// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPolicyDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafPolicyDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafPolicyDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafPolicyDomainsResponseBody) *DescribeDcdnWafPolicyDomainsResponse
	GetBody() *DescribeDcdnWafPolicyDomainsResponseBody
}

type DescribeDcdnWafPolicyDomainsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafPolicyDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafPolicyDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafPolicyDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafPolicyDomainsResponse) GetBody() *DescribeDcdnWafPolicyDomainsResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafPolicyDomainsResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafPolicyDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsResponse) SetStatusCode(v int32) *DescribeDcdnWafPolicyDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsResponse) SetBody(v *DescribeDcdnWafPolicyDomainsResponseBody) *DescribeDcdnWafPolicyDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPolicyValidDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafPolicyValidDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafPolicyValidDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafPolicyValidDomainsResponseBody) *DescribeDcdnWafPolicyValidDomainsResponse
	GetBody() *DescribeDcdnWafPolicyValidDomainsResponseBody
}

type DescribeDcdnWafPolicyValidDomainsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafPolicyValidDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafPolicyValidDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyValidDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyValidDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafPolicyValidDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafPolicyValidDomainsResponse) GetBody() *DescribeDcdnWafPolicyValidDomainsResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafPolicyValidDomainsResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafPolicyValidDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponse) SetStatusCode(v int32) *DescribeDcdnWafPolicyValidDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponse) SetBody(v *DescribeDcdnWafPolicyValidDomainsResponseBody) *DescribeDcdnWafPolicyValidDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsResponse) Validate() error {
	return dara.Validate(s)
}

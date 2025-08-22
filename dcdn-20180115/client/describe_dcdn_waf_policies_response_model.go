// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafPoliciesResponseBody) *DescribeDcdnWafPoliciesResponse
	GetBody() *DescribeDcdnWafPoliciesResponseBody
}

type DescribeDcdnWafPoliciesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafPoliciesResponse) GetBody() *DescribeDcdnWafPoliciesResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafPoliciesResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafPoliciesResponse) SetStatusCode(v int32) *DescribeDcdnWafPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafPoliciesResponse) SetBody(v *DescribeDcdnWafPoliciesResponseBody) *DescribeDcdnWafPoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafPoliciesResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDefaultRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafDefaultRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafDefaultRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafDefaultRulesResponseBody) *DescribeDcdnWafDefaultRulesResponse
	GetBody() *DescribeDcdnWafDefaultRulesResponseBody
}

type DescribeDcdnWafDefaultRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafDefaultRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafDefaultRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDefaultRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDefaultRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafDefaultRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafDefaultRulesResponse) GetBody() *DescribeDcdnWafDefaultRulesResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafDefaultRulesResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafDefaultRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponse) SetStatusCode(v int32) *DescribeDcdnWafDefaultRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponse) SetBody(v *DescribeDcdnWafDefaultRulesResponseBody) *DescribeDcdnWafDefaultRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafDefaultRulesResponse) Validate() error {
	return dara.Validate(s)
}

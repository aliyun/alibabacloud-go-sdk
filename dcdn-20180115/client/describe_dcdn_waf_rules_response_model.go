// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafRulesResponseBody) *DescribeDcdnWafRulesResponse
	GetBody() *DescribeDcdnWafRulesResponseBody
}

type DescribeDcdnWafRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafRulesResponse) GetBody() *DescribeDcdnWafRulesResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafRulesResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafRulesResponse) SetStatusCode(v int32) *DescribeDcdnWafRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafRulesResponse) SetBody(v *DescribeDcdnWafRulesResponseBody) *DescribeDcdnWafRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafRulesResponse) Validate() error {
	return dara.Validate(s)
}

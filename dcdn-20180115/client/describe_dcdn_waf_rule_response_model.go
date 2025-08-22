// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafRuleResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafRuleResponseBody) *DescribeDcdnWafRuleResponse
	GetBody() *DescribeDcdnWafRuleResponseBody
}

type DescribeDcdnWafRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafRuleResponse) GetBody() *DescribeDcdnWafRuleResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafRuleResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafRuleResponse) SetStatusCode(v int32) *DescribeDcdnWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafRuleResponse) SetBody(v *DescribeDcdnWafRuleResponseBody) *DescribeDcdnWafRuleResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafRuleResponse) Validate() error {
	return dara.Validate(s)
}

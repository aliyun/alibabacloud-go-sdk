// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafPolicyResponseBody) *DescribeDcdnWafPolicyResponse
	GetBody() *DescribeDcdnWafPolicyResponseBody
}

type DescribeDcdnWafPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafPolicyResponse) GetBody() *DescribeDcdnWafPolicyResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafPolicyResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafPolicyResponse) SetStatusCode(v int32) *DescribeDcdnWafPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafPolicyResponse) SetBody(v *DescribeDcdnWafPolicyResponseBody) *DescribeDcdnWafPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafPolicyResponse) Validate() error {
	return dara.Validate(s)
}

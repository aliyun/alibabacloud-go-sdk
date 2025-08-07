// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMaskingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMaskingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMaskingRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMaskingRulesResponseBody) *DescribeMaskingRulesResponse
	GetBody() *DescribeMaskingRulesResponseBody
}

type DescribeMaskingRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMaskingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMaskingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaskingRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMaskingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMaskingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMaskingRulesResponse) GetBody() *DescribeMaskingRulesResponseBody {
	return s.Body
}

func (s *DescribeMaskingRulesResponse) SetHeaders(v map[string]*string) *DescribeMaskingRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMaskingRulesResponse) SetStatusCode(v int32) *DescribeMaskingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMaskingRulesResponse) SetBody(v *DescribeMaskingRulesResponseBody) *DescribeMaskingRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeMaskingRulesResponse) Validate() error {
	return dara.Validate(s)
}

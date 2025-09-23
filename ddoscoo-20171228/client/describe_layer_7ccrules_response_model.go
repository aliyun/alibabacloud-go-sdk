// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer7CCRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLayer7CCRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLayer7CCRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLayer7CCRulesResponseBody) *DescribeLayer7CCRulesResponse
	GetBody() *DescribeLayer7CCRulesResponseBody
}

type DescribeLayer7CCRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLayer7CCRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLayer7CCRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer7CCRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer7CCRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLayer7CCRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLayer7CCRulesResponse) GetBody() *DescribeLayer7CCRulesResponseBody {
	return s.Body
}

func (s *DescribeLayer7CCRulesResponse) SetHeaders(v map[string]*string) *DescribeLayer7CCRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLayer7CCRulesResponse) SetStatusCode(v int32) *DescribeLayer7CCRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLayer7CCRulesResponse) SetBody(v *DescribeLayer7CCRulesResponseBody) *DescribeLayer7CCRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeLayer7CCRulesResponse) Validate() error {
	return dara.Validate(s)
}

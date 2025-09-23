// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLayer4RulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLayer4RulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLayer4RulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLayer4RulesResponseBody) *DescribeLayer4RulesResponse
	GetBody() *DescribeLayer4RulesResponseBody
}

type DescribeLayer4RulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLayer4RulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLayer4RulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLayer4RulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLayer4RulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLayer4RulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLayer4RulesResponse) GetBody() *DescribeLayer4RulesResponseBody {
	return s.Body
}

func (s *DescribeLayer4RulesResponse) SetHeaders(v map[string]*string) *DescribeLayer4RulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLayer4RulesResponse) SetStatusCode(v int32) *DescribeLayer4RulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLayer4RulesResponse) SetBody(v *DescribeLayer4RulesResponseBody) *DescribeLayer4RulesResponse {
	s.Body = v
	return s
}

func (s *DescribeLayer4RulesResponse) Validate() error {
	return dara.Validate(s)
}

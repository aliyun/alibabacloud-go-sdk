// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecRulesResponseBody) *DescribeApisecRulesResponse
	GetBody() *DescribeApisecRulesResponseBody
}

type DescribeApisecRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecRulesResponse) GetBody() *DescribeApisecRulesResponseBody {
	return s.Body
}

func (s *DescribeApisecRulesResponse) SetHeaders(v map[string]*string) *DescribeApisecRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecRulesResponse) SetStatusCode(v int32) *DescribeApisecRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecRulesResponse) SetBody(v *DescribeApisecRulesResponseBody) *DescribeApisecRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecRulesResponse) Validate() error {
	return dara.Validate(s)
}

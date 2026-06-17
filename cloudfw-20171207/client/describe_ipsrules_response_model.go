// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPSRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIPSRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIPSRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIPSRulesResponseBody) *DescribeIPSRulesResponse
	GetBody() *DescribeIPSRulesResponseBody
}

type DescribeIPSRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIPSRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIPSRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPSRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeIPSRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIPSRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIPSRulesResponse) GetBody() *DescribeIPSRulesResponseBody {
	return s.Body
}

func (s *DescribeIPSRulesResponse) SetHeaders(v map[string]*string) *DescribeIPSRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeIPSRulesResponse) SetStatusCode(v int32) *DescribeIPSRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIPSRulesResponse) SetBody(v *DescribeIPSRulesResponseBody) *DescribeIPSRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeIPSRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

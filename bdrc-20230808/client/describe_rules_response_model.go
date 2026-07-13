// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRulesResponseBody) *DescribeRulesResponse
	GetBody() *DescribeRulesResponseBody
}

type DescribeRulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRulesResponse) GetBody() *DescribeRulesResponseBody {
	return s.Body
}

func (s *DescribeRulesResponse) SetHeaders(v map[string]*string) *DescribeRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRulesResponse) SetStatusCode(v int32) *DescribeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRulesResponse) SetBody(v *DescribeRulesResponseBody) *DescribeRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

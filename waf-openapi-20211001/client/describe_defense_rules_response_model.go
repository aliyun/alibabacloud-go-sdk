// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseRulesResponseBody) *DescribeDefenseRulesResponse
	GetBody() *DescribeDefenseRulesResponseBody
}

type DescribeDefenseRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseRulesResponse) GetBody() *DescribeDefenseRulesResponseBody {
	return s.Body
}

func (s *DescribeDefenseRulesResponse) SetHeaders(v map[string]*string) *DescribeDefenseRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseRulesResponse) SetStatusCode(v int32) *DescribeDefenseRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseRulesResponse) SetBody(v *DescribeDefenseRulesResponseBody) *DescribeDefenseRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

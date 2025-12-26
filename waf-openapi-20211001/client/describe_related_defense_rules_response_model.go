// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRelatedDefenseRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRelatedDefenseRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRelatedDefenseRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRelatedDefenseRulesResponseBody) *DescribeRelatedDefenseRulesResponse
	GetBody() *DescribeRelatedDefenseRulesResponseBody
}

type DescribeRelatedDefenseRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRelatedDefenseRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRelatedDefenseRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRelatedDefenseRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRelatedDefenseRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRelatedDefenseRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRelatedDefenseRulesResponse) GetBody() *DescribeRelatedDefenseRulesResponseBody {
	return s.Body
}

func (s *DescribeRelatedDefenseRulesResponse) SetHeaders(v map[string]*string) *DescribeRelatedDefenseRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRelatedDefenseRulesResponse) SetStatusCode(v int32) *DescribeRelatedDefenseRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRelatedDefenseRulesResponse) SetBody(v *DescribeRelatedDefenseRulesResponseBody) *DescribeRelatedDefenseRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeRelatedDefenseRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

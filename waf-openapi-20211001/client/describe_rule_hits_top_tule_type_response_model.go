// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopTuleTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleHitsTopTuleTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleHitsTopTuleTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleHitsTopTuleTypeResponseBody) *DescribeRuleHitsTopTuleTypeResponse
	GetBody() *DescribeRuleHitsTopTuleTypeResponseBody
}

type DescribeRuleHitsTopTuleTypeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopTuleTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopTuleTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopTuleTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopTuleTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleHitsTopTuleTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleHitsTopTuleTypeResponse) GetBody() *DescribeRuleHitsTopTuleTypeResponseBody {
	return s.Body
}

func (s *DescribeRuleHitsTopTuleTypeResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopTuleTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponse) SetStatusCode(v int32) *DescribeRuleHitsTopTuleTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponse) SetBody(v *DescribeRuleHitsTopTuleTypeResponseBody) *DescribeRuleHitsTopTuleTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleHitsTopTuleTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

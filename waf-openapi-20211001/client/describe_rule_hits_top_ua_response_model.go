// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopUaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleHitsTopUaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleHitsTopUaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleHitsTopUaResponseBody) *DescribeRuleHitsTopUaResponse
	GetBody() *DescribeRuleHitsTopUaResponseBody
}

type DescribeRuleHitsTopUaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopUaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopUaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopUaResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleHitsTopUaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleHitsTopUaResponse) GetBody() *DescribeRuleHitsTopUaResponseBody {
	return s.Body
}

func (s *DescribeRuleHitsTopUaResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopUaResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopUaResponse) SetStatusCode(v int32) *DescribeRuleHitsTopUaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopUaResponse) SetBody(v *DescribeRuleHitsTopUaResponseBody) *DescribeRuleHitsTopUaResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleHitsTopUaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

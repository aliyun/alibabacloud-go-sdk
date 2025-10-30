// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHitRuleFluctuationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHitRuleFluctuationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHitRuleFluctuationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHitRuleFluctuationResponseBody) *DescribeHitRuleFluctuationResponse
	GetBody() *DescribeHitRuleFluctuationResponseBody
}

type DescribeHitRuleFluctuationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHitRuleFluctuationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHitRuleFluctuationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHitRuleFluctuationResponse) GoString() string {
	return s.String()
}

func (s *DescribeHitRuleFluctuationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHitRuleFluctuationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHitRuleFluctuationResponse) GetBody() *DescribeHitRuleFluctuationResponseBody {
	return s.Body
}

func (s *DescribeHitRuleFluctuationResponse) SetHeaders(v map[string]*string) *DescribeHitRuleFluctuationResponse {
	s.Headers = v
	return s
}

func (s *DescribeHitRuleFluctuationResponse) SetStatusCode(v int32) *DescribeHitRuleFluctuationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHitRuleFluctuationResponse) SetBody(v *DescribeHitRuleFluctuationResponseBody) *DescribeHitRuleFluctuationResponse {
	s.Body = v
	return s
}

func (s *DescribeHitRuleFluctuationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleHitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleHitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleHitResponseBody) *DescribeRuleHitResponse
	GetBody() *DescribeRuleHitResponseBody
}

type DescribeRuleHitResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleHitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleHitResponse) GetBody() *DescribeRuleHitResponseBody {
	return s.Body
}

func (s *DescribeRuleHitResponse) SetHeaders(v map[string]*string) *DescribeRuleHitResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitResponse) SetStatusCode(v int32) *DescribeRuleHitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitResponse) SetBody(v *DescribeRuleHitResponseBody) *DescribeRuleHitResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleHitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

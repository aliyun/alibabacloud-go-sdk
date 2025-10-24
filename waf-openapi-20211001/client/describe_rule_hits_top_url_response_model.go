// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleHitsTopUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleHitsTopUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleHitsTopUrlResponseBody) *DescribeRuleHitsTopUrlResponse
	GetBody() *DescribeRuleHitsTopUrlResponseBody
}

type DescribeRuleHitsTopUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleHitsTopUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleHitsTopUrlResponse) GetBody() *DescribeRuleHitsTopUrlResponseBody {
	return s.Body
}

func (s *DescribeRuleHitsTopUrlResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopUrlResponse) SetStatusCode(v int32) *DescribeRuleHitsTopUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopUrlResponse) SetBody(v *DescribeRuleHitsTopUrlResponseBody) *DescribeRuleHitsTopUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleHitsTopUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

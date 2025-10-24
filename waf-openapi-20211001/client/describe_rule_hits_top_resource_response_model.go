// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleHitsTopResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleHitsTopResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleHitsTopResourceResponseBody) *DescribeRuleHitsTopResourceResponse
	GetBody() *DescribeRuleHitsTopResourceResponseBody
}

type DescribeRuleHitsTopResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleHitsTopResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleHitsTopResourceResponse) GetBody() *DescribeRuleHitsTopResourceResponseBody {
	return s.Body
}

func (s *DescribeRuleHitsTopResourceResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopResourceResponse) SetStatusCode(v int32) *DescribeRuleHitsTopResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopResourceResponse) SetBody(v *DescribeRuleHitsTopResourceResponseBody) *DescribeRuleHitsTopResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleHitsTopResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

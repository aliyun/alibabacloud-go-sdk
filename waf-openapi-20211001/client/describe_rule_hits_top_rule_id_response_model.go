// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleHitsTopRuleIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleHitsTopRuleIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleHitsTopRuleIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleHitsTopRuleIdResponseBody) *DescribeRuleHitsTopRuleIdResponse
	GetBody() *DescribeRuleHitsTopRuleIdResponseBody
}

type DescribeRuleHitsTopRuleIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleHitsTopRuleIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleHitsTopRuleIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleHitsTopRuleIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleHitsTopRuleIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleHitsTopRuleIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleHitsTopRuleIdResponse) GetBody() *DescribeRuleHitsTopRuleIdResponseBody {
	return s.Body
}

func (s *DescribeRuleHitsTopRuleIdResponse) SetHeaders(v map[string]*string) *DescribeRuleHitsTopRuleIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponse) SetStatusCode(v int32) *DescribeRuleHitsTopRuleIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponse) SetBody(v *DescribeRuleHitsTopRuleIdResponseBody) *DescribeRuleHitsTopRuleIdResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleHitsTopRuleIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

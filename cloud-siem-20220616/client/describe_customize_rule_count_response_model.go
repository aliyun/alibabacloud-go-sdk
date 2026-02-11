// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizeRuleCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomizeRuleCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomizeRuleCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomizeRuleCountResponseBody) *DescribeCustomizeRuleCountResponse
	GetBody() *DescribeCustomizeRuleCountResponseBody
}

type DescribeCustomizeRuleCountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizeRuleCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizeRuleCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizeRuleCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizeRuleCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomizeRuleCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomizeRuleCountResponse) GetBody() *DescribeCustomizeRuleCountResponseBody {
	return s.Body
}

func (s *DescribeCustomizeRuleCountResponse) SetHeaders(v map[string]*string) *DescribeCustomizeRuleCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizeRuleCountResponse) SetStatusCode(v int32) *DescribeCustomizeRuleCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizeRuleCountResponse) SetBody(v *DescribeCustomizeRuleCountResponseBody) *DescribeCustomizeRuleCountResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomizeRuleCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuthRulePageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuthRulePageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuthRulePageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuthRulePageListResponseBody) *DescribeAuthRulePageListResponse
	GetBody() *DescribeAuthRulePageListResponseBody
}

type DescribeAuthRulePageListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuthRulePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuthRulePageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuthRulePageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthRulePageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuthRulePageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuthRulePageListResponse) GetBody() *DescribeAuthRulePageListResponseBody {
	return s.Body
}

func (s *DescribeAuthRulePageListResponse) SetHeaders(v map[string]*string) *DescribeAuthRulePageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthRulePageListResponse) SetStatusCode(v int32) *DescribeAuthRulePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuthRulePageListResponse) SetBody(v *DescribeAuthRulePageListResponseBody) *DescribeAuthRulePageListResponse {
	s.Body = v
	return s
}

func (s *DescribeAuthRulePageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

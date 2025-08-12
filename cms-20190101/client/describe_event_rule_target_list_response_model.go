// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventRuleTargetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventRuleTargetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventRuleTargetListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventRuleTargetListResponseBody) *DescribeEventRuleTargetListResponse
	GetBody() *DescribeEventRuleTargetListResponseBody
}

type DescribeEventRuleTargetListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventRuleTargetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventRuleTargetListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventRuleTargetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventRuleTargetListResponse) GetBody() *DescribeEventRuleTargetListResponseBody {
	return s.Body
}

func (s *DescribeEventRuleTargetListResponse) SetHeaders(v map[string]*string) *DescribeEventRuleTargetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventRuleTargetListResponse) SetStatusCode(v int32) *DescribeEventRuleTargetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventRuleTargetListResponse) SetBody(v *DescribeEventRuleTargetListResponseBody) *DescribeEventRuleTargetListResponse {
	s.Body = v
	return s
}

func (s *DescribeEventRuleTargetListResponse) Validate() error {
	return dara.Validate(s)
}

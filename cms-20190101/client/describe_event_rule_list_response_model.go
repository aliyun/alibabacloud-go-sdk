// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventRuleListResponseBody) *DescribeEventRuleListResponse
	GetBody() *DescribeEventRuleListResponseBody
}

type DescribeEventRuleListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventRuleListResponse) GetBody() *DescribeEventRuleListResponseBody {
	return s.Body
}

func (s *DescribeEventRuleListResponse) SetHeaders(v map[string]*string) *DescribeEventRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventRuleListResponse) SetStatusCode(v int32) *DescribeEventRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventRuleListResponse) SetBody(v *DescribeEventRuleListResponseBody) *DescribeEventRuleListResponse {
	s.Body = v
	return s
}

func (s *DescribeEventRuleListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

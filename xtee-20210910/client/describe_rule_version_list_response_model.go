// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRuleVersionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRuleVersionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRuleVersionListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRuleVersionListResponseBody) *DescribeRuleVersionListResponse
	GetBody() *DescribeRuleVersionListResponseBody
}

type DescribeRuleVersionListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRuleVersionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRuleVersionListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRuleVersionListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRuleVersionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRuleVersionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRuleVersionListResponse) GetBody() *DescribeRuleVersionListResponseBody {
	return s.Body
}

func (s *DescribeRuleVersionListResponse) SetHeaders(v map[string]*string) *DescribeRuleVersionListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRuleVersionListResponse) SetStatusCode(v int32) *DescribeRuleVersionListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRuleVersionListResponse) SetBody(v *DescribeRuleVersionListResponseBody) *DescribeRuleVersionListResponse {
	s.Body = v
	return s
}

func (s *DescribeRuleVersionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

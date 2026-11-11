// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLoginPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserLoginPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserLoginPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserLoginPoliciesResponseBody) *DescribeUserLoginPoliciesResponse
	GetBody() *DescribeUserLoginPoliciesResponseBody
}

type DescribeUserLoginPoliciesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserLoginPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserLoginPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLoginPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserLoginPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserLoginPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserLoginPoliciesResponse) GetBody() *DescribeUserLoginPoliciesResponseBody {
	return s.Body
}

func (s *DescribeUserLoginPoliciesResponse) SetHeaders(v map[string]*string) *DescribeUserLoginPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserLoginPoliciesResponse) SetStatusCode(v int32) *DescribeUserLoginPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserLoginPoliciesResponse) SetBody(v *DescribeUserLoginPoliciesResponseBody) *DescribeUserLoginPoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeUserLoginPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

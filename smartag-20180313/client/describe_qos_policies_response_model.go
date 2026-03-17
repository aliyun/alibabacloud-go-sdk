// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQosPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQosPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQosPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQosPoliciesResponseBody) *DescribeQosPoliciesResponse
	GetBody() *DescribeQosPoliciesResponseBody
}

type DescribeQosPoliciesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQosPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQosPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQosPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeQosPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQosPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQosPoliciesResponse) GetBody() *DescribeQosPoliciesResponseBody {
	return s.Body
}

func (s *DescribeQosPoliciesResponse) SetHeaders(v map[string]*string) *DescribeQosPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeQosPoliciesResponse) SetStatusCode(v int32) *DescribeQosPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQosPoliciesResponse) SetBody(v *DescribeQosPoliciesResponseBody) *DescribeQosPoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeQosPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

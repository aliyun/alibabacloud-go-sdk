// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomScenePoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomScenePoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomScenePoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomScenePoliciesResponseBody) *DescribeCustomScenePoliciesResponse
	GetBody() *DescribeCustomScenePoliciesResponseBody
}

type DescribeCustomScenePoliciesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomScenePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomScenePoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomScenePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomScenePoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomScenePoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomScenePoliciesResponse) GetBody() *DescribeCustomScenePoliciesResponseBody {
	return s.Body
}

func (s *DescribeCustomScenePoliciesResponse) SetHeaders(v map[string]*string) *DescribeCustomScenePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomScenePoliciesResponse) SetStatusCode(v int32) *DescribeCustomScenePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomScenePoliciesResponse) SetBody(v *DescribeCustomScenePoliciesResponseBody) *DescribeCustomScenePoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomScenePoliciesResponse) Validate() error {
	return dara.Validate(s)
}

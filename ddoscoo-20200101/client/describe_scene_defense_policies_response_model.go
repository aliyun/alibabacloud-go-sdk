// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSceneDefensePoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSceneDefensePoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSceneDefensePoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSceneDefensePoliciesResponseBody) *DescribeSceneDefensePoliciesResponse
	GetBody() *DescribeSceneDefensePoliciesResponseBody
}

type DescribeSceneDefensePoliciesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSceneDefensePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSceneDefensePoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSceneDefensePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSceneDefensePoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSceneDefensePoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSceneDefensePoliciesResponse) GetBody() *DescribeSceneDefensePoliciesResponseBody {
	return s.Body
}

func (s *DescribeSceneDefensePoliciesResponse) SetHeaders(v map[string]*string) *DescribeSceneDefensePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSceneDefensePoliciesResponse) SetStatusCode(v int32) *DescribeSceneDefensePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSceneDefensePoliciesResponse) SetBody(v *DescribeSceneDefensePoliciesResponseBody) *DescribeSceneDefensePoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeSceneDefensePoliciesResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinarySecurityPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBinarySecurityPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBinarySecurityPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBinarySecurityPoliciesResponseBody) *DescribeBinarySecurityPoliciesResponse
	GetBody() *DescribeBinarySecurityPoliciesResponseBody
}

type DescribeBinarySecurityPoliciesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBinarySecurityPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBinarySecurityPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinarySecurityPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DescribeBinarySecurityPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBinarySecurityPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBinarySecurityPoliciesResponse) GetBody() *DescribeBinarySecurityPoliciesResponseBody {
	return s.Body
}

func (s *DescribeBinarySecurityPoliciesResponse) SetHeaders(v map[string]*string) *DescribeBinarySecurityPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponse) SetStatusCode(v int32) *DescribeBinarySecurityPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponse) SetBody(v *DescribeBinarySecurityPoliciesResponseBody) *DescribeBinarySecurityPoliciesResponse {
	s.Body = v
	return s
}

func (s *DescribeBinarySecurityPoliciesResponse) Validate() error {
	return dara.Validate(s)
}

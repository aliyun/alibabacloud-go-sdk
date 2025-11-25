// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeControlPolicyDomainResolveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeControlPolicyDomainResolveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeControlPolicyDomainResolveResponse
	GetStatusCode() *int32
	SetBody(v *DescribeControlPolicyDomainResolveResponseBody) *DescribeControlPolicyDomainResolveResponse
	GetBody() *DescribeControlPolicyDomainResolveResponseBody
}

type DescribeControlPolicyDomainResolveResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeControlPolicyDomainResolveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeControlPolicyDomainResolveResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeControlPolicyDomainResolveResponse) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyDomainResolveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeControlPolicyDomainResolveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeControlPolicyDomainResolveResponse) GetBody() *DescribeControlPolicyDomainResolveResponseBody {
	return s.Body
}

func (s *DescribeControlPolicyDomainResolveResponse) SetHeaders(v map[string]*string) *DescribeControlPolicyDomainResolveResponse {
	s.Headers = v
	return s
}

func (s *DescribeControlPolicyDomainResolveResponse) SetStatusCode(v int32) *DescribeControlPolicyDomainResolveResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeControlPolicyDomainResolveResponse) SetBody(v *DescribeControlPolicyDomainResolveResponseBody) *DescribeControlPolicyDomainResolveResponse {
	s.Body = v
	return s
}

func (s *DescribeControlPolicyDomainResolveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

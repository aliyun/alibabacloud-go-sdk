// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclCheckQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclCheckQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclCheckQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclCheckQuotaResponseBody) *DescribeAclCheckQuotaResponse
	GetBody() *DescribeAclCheckQuotaResponseBody
}

type DescribeAclCheckQuotaResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclCheckQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclCheckQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclCheckQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclCheckQuotaResponse) GetBody() *DescribeAclCheckQuotaResponseBody {
	return s.Body
}

func (s *DescribeAclCheckQuotaResponse) SetHeaders(v map[string]*string) *DescribeAclCheckQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclCheckQuotaResponse) SetStatusCode(v int32) *DescribeAclCheckQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclCheckQuotaResponse) SetBody(v *DescribeAclCheckQuotaResponseBody) *DescribeAclCheckQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeAclCheckQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

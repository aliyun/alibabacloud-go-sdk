// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackProtectionCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAttackProtectionCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAttackProtectionCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAttackProtectionCountResponseBody) *DescribeAttackProtectionCountResponse
	GetBody() *DescribeAttackProtectionCountResponseBody
}

type DescribeAttackProtectionCountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAttackProtectionCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAttackProtectionCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackProtectionCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeAttackProtectionCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAttackProtectionCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAttackProtectionCountResponse) GetBody() *DescribeAttackProtectionCountResponseBody {
	return s.Body
}

func (s *DescribeAttackProtectionCountResponse) SetHeaders(v map[string]*string) *DescribeAttackProtectionCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeAttackProtectionCountResponse) SetStatusCode(v int32) *DescribeAttackProtectionCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAttackProtectionCountResponse) SetBody(v *DescribeAttackProtectionCountResponseBody) *DescribeAttackProtectionCountResponse {
	s.Body = v
	return s
}

func (s *DescribeAttackProtectionCountResponse) Validate() error {
	return dara.Validate(s)
}

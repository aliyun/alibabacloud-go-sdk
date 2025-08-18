// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHttpDDoSAttackProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHttpDDoSAttackProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHttpDDoSAttackProtectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHttpDDoSAttackProtectionResponseBody) *DescribeHttpDDoSAttackProtectionResponse
	GetBody() *DescribeHttpDDoSAttackProtectionResponseBody
}

type DescribeHttpDDoSAttackProtectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHttpDDoSAttackProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHttpDDoSAttackProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHttpDDoSAttackProtectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeHttpDDoSAttackProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHttpDDoSAttackProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHttpDDoSAttackProtectionResponse) GetBody() *DescribeHttpDDoSAttackProtectionResponseBody {
	return s.Body
}

func (s *DescribeHttpDDoSAttackProtectionResponse) SetHeaders(v map[string]*string) *DescribeHttpDDoSAttackProtectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponse) SetStatusCode(v int32) *DescribeHttpDDoSAttackProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponse) SetBody(v *DescribeHttpDDoSAttackProtectionResponseBody) *DescribeHttpDDoSAttackProtectionResponse {
	s.Body = v
	return s
}

func (s *DescribeHttpDDoSAttackProtectionResponse) Validate() error {
	return dara.Validate(s)
}

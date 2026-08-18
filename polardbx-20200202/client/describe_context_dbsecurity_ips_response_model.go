// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContextDBSecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContextDBSecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContextDBSecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContextDBSecurityIpsResponseBody) *DescribeContextDBSecurityIpsResponse
	GetBody() *DescribeContextDBSecurityIpsResponseBody
}

type DescribeContextDBSecurityIpsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContextDBSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContextDBSecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContextDBSecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContextDBSecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContextDBSecurityIpsResponse) GetBody() *DescribeContextDBSecurityIpsResponseBody {
	return s.Body
}

func (s *DescribeContextDBSecurityIpsResponse) SetHeaders(v map[string]*string) *DescribeContextDBSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContextDBSecurityIpsResponse) SetStatusCode(v int32) *DescribeContextDBSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContextDBSecurityIpsResponse) SetBody(v *DescribeContextDBSecurityIpsResponseBody) *DescribeContextDBSecurityIpsResponse {
	s.Body = v
	return s
}

func (s *DescribeContextDBSecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

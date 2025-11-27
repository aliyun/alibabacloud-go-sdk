// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountMaskingPrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountMaskingPrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountMaskingPrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountMaskingPrivilegeResponseBody) *DescribeAccountMaskingPrivilegeResponse
	GetBody() *DescribeAccountMaskingPrivilegeResponseBody
}

type DescribeAccountMaskingPrivilegeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountMaskingPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountMaskingPrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountMaskingPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountMaskingPrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountMaskingPrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountMaskingPrivilegeResponse) GetBody() *DescribeAccountMaskingPrivilegeResponseBody {
	return s.Body
}

func (s *DescribeAccountMaskingPrivilegeResponse) SetHeaders(v map[string]*string) *DescribeAccountMaskingPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountMaskingPrivilegeResponse) SetStatusCode(v int32) *DescribeAccountMaskingPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeResponse) SetBody(v *DescribeAccountMaskingPrivilegeResponseBody) *DescribeAccountMaskingPrivilegeResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountMaskingPrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

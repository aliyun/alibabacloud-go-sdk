// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionDBRolePrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEncryptionDBRolePrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEncryptionDBRolePrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEncryptionDBRolePrivilegeResponseBody) *DescribeEncryptionDBRolePrivilegeResponse
	GetBody() *DescribeEncryptionDBRolePrivilegeResponseBody
}

type DescribeEncryptionDBRolePrivilegeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEncryptionDBRolePrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEncryptionDBRolePrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionDBRolePrivilegeResponse) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionDBRolePrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEncryptionDBRolePrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEncryptionDBRolePrivilegeResponse) GetBody() *DescribeEncryptionDBRolePrivilegeResponseBody {
	return s.Body
}

func (s *DescribeEncryptionDBRolePrivilegeResponse) SetHeaders(v map[string]*string) *DescribeEncryptionDBRolePrivilegeResponse {
	s.Headers = v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponse) SetStatusCode(v int32) *DescribeEncryptionDBRolePrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponse) SetBody(v *DescribeEncryptionDBRolePrivilegeResponseBody) *DescribeEncryptionDBRolePrivilegeResponse {
	s.Body = v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEncryptionDBRolePrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEncryptionDBRolePrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEncryptionDBRolePrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEncryptionDBRolePrivilegeResponseBody) *ModifyEncryptionDBRolePrivilegeResponse
	GetBody() *ModifyEncryptionDBRolePrivilegeResponseBody
}

type ModifyEncryptionDBRolePrivilegeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEncryptionDBRolePrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEncryptionDBRolePrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEncryptionDBRolePrivilegeResponse) GoString() string {
	return s.String()
}

func (s *ModifyEncryptionDBRolePrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEncryptionDBRolePrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEncryptionDBRolePrivilegeResponse) GetBody() *ModifyEncryptionDBRolePrivilegeResponseBody {
	return s.Body
}

func (s *ModifyEncryptionDBRolePrivilegeResponse) SetHeaders(v map[string]*string) *ModifyEncryptionDBRolePrivilegeResponse {
	s.Headers = v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeResponse) SetStatusCode(v int32) *ModifyEncryptionDBRolePrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeResponse) SetBody(v *ModifyEncryptionDBRolePrivilegeResponseBody) *ModifyEncryptionDBRolePrivilegeResponse {
	s.Body = v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeResponse) Validate() error {
	return dara.Validate(s)
}

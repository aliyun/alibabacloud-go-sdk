// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEncryptionDBRolePrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddEncryptionDBRolePrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddEncryptionDBRolePrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *AddEncryptionDBRolePrivilegeResponseBody) *AddEncryptionDBRolePrivilegeResponse
	GetBody() *AddEncryptionDBRolePrivilegeResponseBody
}

type AddEncryptionDBRolePrivilegeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEncryptionDBRolePrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEncryptionDBRolePrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddEncryptionDBRolePrivilegeResponse) GoString() string {
	return s.String()
}

func (s *AddEncryptionDBRolePrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddEncryptionDBRolePrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddEncryptionDBRolePrivilegeResponse) GetBody() *AddEncryptionDBRolePrivilegeResponseBody {
	return s.Body
}

func (s *AddEncryptionDBRolePrivilegeResponse) SetHeaders(v map[string]*string) *AddEncryptionDBRolePrivilegeResponse {
	s.Headers = v
	return s
}

func (s *AddEncryptionDBRolePrivilegeResponse) SetStatusCode(v int32) *AddEncryptionDBRolePrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeResponse) SetBody(v *AddEncryptionDBRolePrivilegeResponseBody) *AddEncryptionDBRolePrivilegeResponse {
	s.Body = v
	return s
}

func (s *AddEncryptionDBRolePrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

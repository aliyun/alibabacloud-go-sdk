// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEncryptionDBRolePrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEncryptionDBRolePrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEncryptionDBRolePrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEncryptionDBRolePrivilegeResponseBody) *DeleteEncryptionDBRolePrivilegeResponse
	GetBody() *DeleteEncryptionDBRolePrivilegeResponseBody
}

type DeleteEncryptionDBRolePrivilegeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEncryptionDBRolePrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEncryptionDBRolePrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEncryptionDBRolePrivilegeResponse) GoString() string {
	return s.String()
}

func (s *DeleteEncryptionDBRolePrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEncryptionDBRolePrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEncryptionDBRolePrivilegeResponse) GetBody() *DeleteEncryptionDBRolePrivilegeResponseBody {
	return s.Body
}

func (s *DeleteEncryptionDBRolePrivilegeResponse) SetHeaders(v map[string]*string) *DeleteEncryptionDBRolePrivilegeResponse {
	s.Headers = v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeResponse) SetStatusCode(v int32) *DeleteEncryptionDBRolePrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeResponse) SetBody(v *DeleteEncryptionDBRolePrivilegeResponseBody) *DeleteEncryptionDBRolePrivilegeResponse {
	s.Body = v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

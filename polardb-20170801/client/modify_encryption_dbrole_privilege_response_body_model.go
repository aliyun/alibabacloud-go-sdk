// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEncryptionDBRolePrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyEncryptionDBRolePrivilegeResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *ModifyEncryptionDBRolePrivilegeResponseBody
	GetRequestId() *string
}

type ModifyEncryptionDBRolePrivilegeResponseBody struct {
	// example:
	//
	// pc-******************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEncryptionDBRolePrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEncryptionDBRolePrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEncryptionDBRolePrivilegeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyEncryptionDBRolePrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEncryptionDBRolePrivilegeResponseBody) SetDBClusterId(v string) *ModifyEncryptionDBRolePrivilegeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeResponseBody) SetRequestId(v string) *ModifyEncryptionDBRolePrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEncryptionDBRolePrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}

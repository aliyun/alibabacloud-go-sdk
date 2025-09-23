// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEncryptionDBRolePrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *AddEncryptionDBRolePrivilegeResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *AddEncryptionDBRolePrivilegeResponseBody
	GetRequestId() *string
}

type AddEncryptionDBRolePrivilegeResponseBody struct {
	// example:
	//
	// pc-bp10gr51qasnl****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6A2EE5B4-CC9F-46E1-A747-E43BC9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEncryptionDBRolePrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddEncryptionDBRolePrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *AddEncryptionDBRolePrivilegeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddEncryptionDBRolePrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddEncryptionDBRolePrivilegeResponseBody) SetDBClusterId(v string) *AddEncryptionDBRolePrivilegeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeResponseBody) SetRequestId(v string) *AddEncryptionDBRolePrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEncryptionDBRolePrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}

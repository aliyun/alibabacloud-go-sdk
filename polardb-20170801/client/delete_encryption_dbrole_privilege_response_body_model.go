// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEncryptionDBRolePrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteEncryptionDBRolePrivilegeResponseBody
	GetDBClusterId() *string
	SetMessage(v string) *DeleteEncryptionDBRolePrivilegeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEncryptionDBRolePrivilegeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEncryptionDBRolePrivilegeResponseBody
	GetSuccess() *bool
}

type DeleteEncryptionDBRolePrivilegeResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The message returned for the request.
	//
	// > If the request is successful, `Successful` is returned. If the request fails, an error message is returned, such as an error code.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEncryptionDBRolePrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEncryptionDBRolePrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEncryptionDBRolePrivilegeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteEncryptionDBRolePrivilegeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEncryptionDBRolePrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEncryptionDBRolePrivilegeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEncryptionDBRolePrivilegeResponseBody) SetDBClusterId(v string) *DeleteEncryptionDBRolePrivilegeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeResponseBody) SetMessage(v string) *DeleteEncryptionDBRolePrivilegeResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeResponseBody) SetRequestId(v string) *DeleteEncryptionDBRolePrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeResponseBody) SetSuccess(v bool) *DeleteEncryptionDBRolePrivilegeResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEncryptionDBRolePrivilegeResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountId(v string) *ModifyHostAccountRequest
	GetHostAccountId() *string
	SetHostAccountName(v string) *ModifyHostAccountRequest
	GetHostAccountName() *string
	SetHostShareKeyId(v string) *ModifyHostAccountRequest
	GetHostShareKeyId() *string
	SetInstanceId(v string) *ModifyHostAccountRequest
	GetInstanceId() *string
	SetPassPhrase(v string) *ModifyHostAccountRequest
	GetPassPhrase() *string
	SetPassword(v string) *ModifyHostAccountRequest
	GetPassword() *string
	SetPrivateKey(v string) *ModifyHostAccountRequest
	GetPrivateKey() *string
	SetPrivilegeType(v string) *ModifyHostAccountRequest
	GetPrivilegeType() *string
	SetRegionId(v string) *ModifyHostAccountRequest
	GetRegionId() *string
	SetRotationMode(v string) *ModifyHostAccountRequest
	GetRotationMode() *string
}

type ModifyHostAccountRequest struct {
	// Specifies the ID of the host account to be modified.
	//
	// > You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// Specifies the modified host account name, which can contain up to 128 characters.
	//
	// example:
	//
	// abc
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The host shared key ID.
	//
	// > You can obtain this ID by calling the [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) operation.
	//
	// example:
	//
	// 1
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// Specifies the ID of the Bastionhost instance where the host account to be modified resides.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the Bastionhost instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the modified security token of the host account\\"s private key.
	//
	// > This parameter takes effect when the host account protocol is SSH. This parameter is not required when the host account protocol is RDP.
	//
	// example:
	//
	// 123456
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// Specifies the modified password of the host account.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies the modified private key of the host account, which is a Base64-encoded string.
	//
	// > This parameter takes effect when the host account protocol is SSH. This parameter is not required when the host account protocol is RDP. You can call the [GetHostAccount](https://help.aliyun.com/document_detail/204391.html) operation to query the protocol used by the host account. You can configure both a password and a private key for a host account. When connecting to an asset, Bastionhost preferentially uses the private key for connection.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----
	//
	// ......
	//
	// -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// Account permission type. Valid values:
	//
	// - **Privileged**: privileged account
	//
	// - **Normal**: regular account
	//
	// > This parameter is supported only in V3.2.47 and later versions.
	//
	// example:
	//
	// Normal
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// Specifies the region ID of the Bastionhost instance where the host account to be queried resides.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Account password rotation mode. Valid values:
	//
	// - **Privileged**: Use a privileged account to change the password
	//
	// - **Self**: Do not use a privileged account to change the password
	//
	// > This parameter is supported only in V3.2.47 and later versions.
	//
	// example:
	//
	// Self
	RotationMode *string `json:"RotationMode,omitempty" xml:"RotationMode,omitempty"`
}

func (s ModifyHostAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostAccountRequest) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *ModifyHostAccountRequest) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ModifyHostAccountRequest) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *ModifyHostAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHostAccountRequest) GetPassPhrase() *string {
	return s.PassPhrase
}

func (s *ModifyHostAccountRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyHostAccountRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *ModifyHostAccountRequest) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *ModifyHostAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHostAccountRequest) GetRotationMode() *string {
	return s.RotationMode
}

func (s *ModifyHostAccountRequest) SetHostAccountId(v string) *ModifyHostAccountRequest {
	s.HostAccountId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetHostAccountName(v string) *ModifyHostAccountRequest {
	s.HostAccountName = &v
	return s
}

func (s *ModifyHostAccountRequest) SetHostShareKeyId(v string) *ModifyHostAccountRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetInstanceId(v string) *ModifyHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetPassPhrase(v string) *ModifyHostAccountRequest {
	s.PassPhrase = &v
	return s
}

func (s *ModifyHostAccountRequest) SetPassword(v string) *ModifyHostAccountRequest {
	s.Password = &v
	return s
}

func (s *ModifyHostAccountRequest) SetPrivateKey(v string) *ModifyHostAccountRequest {
	s.PrivateKey = &v
	return s
}

func (s *ModifyHostAccountRequest) SetPrivilegeType(v string) *ModifyHostAccountRequest {
	s.PrivilegeType = &v
	return s
}

func (s *ModifyHostAccountRequest) SetRegionId(v string) *ModifyHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetRotationMode(v string) *ModifyHostAccountRequest {
	s.RotationMode = &v
	return s
}

func (s *ModifyHostAccountRequest) Validate() error {
	return dara.Validate(s)
}

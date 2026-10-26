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
	SetProjectId(v int64) *ModifyHostAccountRequest
	GetProjectId() *int64
	SetRegionId(v string) *ModifyHostAccountRequest
	GetRegionId() *string
	SetRotationMode(v string) *ModifyHostAccountRequest
	GetRotationMode() *string
}

type ModifyHostAccountRequest struct {
	// The ID of the host account to modify.
	//
	// > You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The new name of the host account. The name can be up to 128 characters in length.
	//
	// example:
	//
	// abc
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the shared key for the host.
	//
	// > You can call the [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) operation to obtain this ID.
	//
	// example:
	//
	// 1
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the bastion host instance where the host account to modify resides.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the bastion host instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new passphrase of the private key for the host account.
	//
	// > This parameter takes effect only when the protocol of the host account is SSH. If the protocol is RDP, you do not need to configure this parameter.
	//
	// example:
	//
	// 123456
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// The new password of the host account.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The new private key of the host account. The private key is a Base64-encoded string.
	//
	// > This parameter takes effect only when the protocol of the host account is SSH. If the protocol is RDP, you do not need to configure this parameter. You can call the [GetHostAccount](https://help.aliyun.com/document_detail/204391.html) operation to query the protocol used by the host account. You can configure both a password and a private key for the host account. When connecting to an asset, the bastion host preferentially uses the private key.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----
	//
	// ......
	//
	// -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The privilege type of the account. Valid values:
	//
	// - **Privileged**: privileged account.
	//
	// - **Normal**: standard account.
	//
	// >Only supported in V3.2.47 and later.
	//
	// example:
	//
	// Normal
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The ID of the project to which the host of the host account belongs.
	//
	// example:
	//
	// 1
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host instance where the host account to query resides.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The password rotation mode of the account. Valid values:
	//
	// - **Privileged**: uses a privileged account to rotate the password.
	//
	// - **Self**: does not use a privileged account to rotate the password.
	//
	// >Only supported in V3.2.47 and later.
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

func (s *ModifyHostAccountRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *ModifyHostAccountRequest) SetProjectId(v int64) *ModifyHostAccountRequest {
	s.ProjectId = &v
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

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
	// The ID of the host account whose information you want to modify.
	//
	// > You can call the [ListHostAccounts](https://help.aliyun.com/document_detail/204372.html) operation to query the ID of the host account.
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
	// The ID of the shared key that is associated with the host.
	//
	// >  You can call the [ListHostShareKeys](https://help.aliyun.com/document_detail/462973.html) operation to query the shared key ID.
	//
	// example:
	//
	// 1
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the bastion host in which you want to modify the information about the host account.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The passphrase for the new private key of the host account.
	//
	// >  This parameter is valid only if the protocol used by the host is SSH. You do not need to configure this parameter if the protocol used by the host is Remote Desktop Protocol (RDP).
	//
	// example:
	//
	// ****
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// The new password of the host account.
	//
	// example:
	//
	// ****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The new private key of the host account. Specify a Base64-encoded string.
	//
	// >  This parameter takes effect only if the protocol used by the host is SSH. You do not need to configure this parameter if the protocol used by the host is Remote Desktop Protocol (RDP). You can call the [GetHostAccount](https://help.aliyun.com/document_detail/204391.html) operation to query the protocol used by the host. You can configure a password and a private key for the host account at the same time. If both a password and a private key are configured for the host account, Bastionhost preferentially uses the private key for logon.
	//
	// example:
	//
	// ****
	PrivateKey    *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The region ID of the bastion host in which you want to query the details of the host account.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

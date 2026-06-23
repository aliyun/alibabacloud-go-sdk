// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccountName(v string) *CreateHostAccountRequest
	GetHostAccountName() *string
	SetHostId(v string) *CreateHostAccountRequest
	GetHostId() *string
	SetHostShareKeyId(v string) *CreateHostAccountRequest
	GetHostShareKeyId() *string
	SetInstanceId(v string) *CreateHostAccountRequest
	GetInstanceId() *string
	SetPassPhrase(v string) *CreateHostAccountRequest
	GetPassPhrase() *string
	SetPassword(v string) *CreateHostAccountRequest
	GetPassword() *string
	SetPrivateKey(v string) *CreateHostAccountRequest
	GetPrivateKey() *string
	SetPrivilegeType(v string) *CreateHostAccountRequest
	GetPrivilegeType() *string
	SetProtocolName(v string) *CreateHostAccountRequest
	GetProtocolName() *string
	SetRegionId(v string) *CreateHostAccountRequest
	GetRegionId() *string
	SetRotationMode(v string) *CreateHostAccountRequest
	GetRotationMode() *string
}

type CreateHostAccountRequest struct {
	// The name of the new host account. The name can be up to 128 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// accountname
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host for which you want to create a host account.
	//
	// > Call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to obtain the host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the shared key for the host.
	//
	// example:
	//
	// 1
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the Bastionhost instance where you want to create the host account.
	//
	// > Call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The passphrase for the private key of the new host account.
	//
	// > You can set this parameter only when ProtocolName is set to SSH. You do not need to set this parameter if ProtocolName is set to RDP.
	//
	// example:
	//
	// 123456
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// The password of the new host account.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The private key of the new host account. The value is a Base64-encoded string.
	//
	// > This parameter is used only when ProtocolName is set to SSH. You do not need to set this parameter if ProtocolName is set to RDP. You can set both a password and a private key for the host account. When connecting to the asset, Bastionhost prioritizes the private key for the connection.
	//
	// example:
	//
	// LS0tLS1******RCBSU0tLQ==
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The permission type of the account. If you do not set this parameter, the default value is Normal.
	//
	// - **Privileged**: privileged account
	//
	// - **Normal**: normal account
	//
	// > This parameter is supported only in Bastionhost V3.2.47 and later.
	//
	// example:
	//
	// Normal
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The protocol of the new host account. <br>Valid values:<br>
	//
	// - SSH
	//
	// - RDP
	//
	// This parameter is required.
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// The region ID of the Bastionhost instance where you want to create the host account.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The password change mode for the account. If you do not set this parameter, the default value is Self.
	//
	// - **Privileged**: Use a privileged account to change the password.
	//
	// - **Self**: Do not use a privileged account to change the password.
	//
	// > This parameter is supported only in Bastionhost V3.2.47 and later.
	//
	// example:
	//
	// Self
	RotationMode *string `json:"RotationMode,omitempty" xml:"RotationMode,omitempty"`
}

func (s CreateHostAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHostAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateHostAccountRequest) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *CreateHostAccountRequest) GetHostId() *string {
	return s.HostId
}

func (s *CreateHostAccountRequest) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *CreateHostAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHostAccountRequest) GetPassPhrase() *string {
	return s.PassPhrase
}

func (s *CreateHostAccountRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateHostAccountRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *CreateHostAccountRequest) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *CreateHostAccountRequest) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *CreateHostAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHostAccountRequest) GetRotationMode() *string {
	return s.RotationMode
}

func (s *CreateHostAccountRequest) SetHostAccountName(v string) *CreateHostAccountRequest {
	s.HostAccountName = &v
	return s
}

func (s *CreateHostAccountRequest) SetHostId(v string) *CreateHostAccountRequest {
	s.HostId = &v
	return s
}

func (s *CreateHostAccountRequest) SetHostShareKeyId(v string) *CreateHostAccountRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *CreateHostAccountRequest) SetInstanceId(v string) *CreateHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostAccountRequest) SetPassPhrase(v string) *CreateHostAccountRequest {
	s.PassPhrase = &v
	return s
}

func (s *CreateHostAccountRequest) SetPassword(v string) *CreateHostAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateHostAccountRequest) SetPrivateKey(v string) *CreateHostAccountRequest {
	s.PrivateKey = &v
	return s
}

func (s *CreateHostAccountRequest) SetPrivilegeType(v string) *CreateHostAccountRequest {
	s.PrivilegeType = &v
	return s
}

func (s *CreateHostAccountRequest) SetProtocolName(v string) *CreateHostAccountRequest {
	s.ProtocolName = &v
	return s
}

func (s *CreateHostAccountRequest) SetRegionId(v string) *CreateHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHostAccountRequest) SetRotationMode(v string) *CreateHostAccountRequest {
	s.RotationMode = &v
	return s
}

func (s *CreateHostAccountRequest) Validate() error {
	return dara.Validate(s)
}

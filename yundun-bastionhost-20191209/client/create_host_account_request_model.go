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
	// The name of the host account. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host to which you want to add a host account.
	//
	// >  You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to query the ID of the host.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the shared key.
	//
	// example:
	//
	// 1
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the bastion host in which you want to add a host account to the host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The passphrase for the private key of the host account.
	//
	// > You can configure this parameter only if ProtocolName is set to SSH. You do not need to configure this parameter if ProtocolName is set to RDP.
	//
	// example:
	//
	// ****
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// The password of the host account.
	//
	// example:
	//
	// ****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The private key of the host account. Specify a Base64-encoded string.
	//
	// > This parameter is valid only if ProtocolName is set to SSH. You do not need to configure this parameter if ProtocolName is set to RDP. You can configure a password and a private key for the host account at the same time. If both a password and a private key are configured for the host account, Bastionhost preferentially uses the private key for logon.
	//
	// example:
	//
	// ****
	PrivateKey    *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The protocol of the host to which you want to add a host account.
	//
	// Valid values:
	//
	// 	- SSH
	//
	// 	- RDP
	//
	// This parameter is required.
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// The region ID of the bastion host in which you want to add a host account to the host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

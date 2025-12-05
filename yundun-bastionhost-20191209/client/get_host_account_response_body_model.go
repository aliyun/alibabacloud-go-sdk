// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccount(v *GetHostAccountResponseBodyHostAccount) *GetHostAccountResponseBody
	GetHostAccount() *GetHostAccountResponseBodyHostAccount
	SetRequestId(v string) *GetHostAccountResponseBody
	GetRequestId() *string
}

type GetHostAccountResponseBody struct {
	// The details of the host account that was queried.
	HostAccount *GetHostAccountResponseBodyHostAccount `json:"HostAccount,omitempty" xml:"HostAccount,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHostAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostAccountResponseBody) GetHostAccount() *GetHostAccountResponseBodyHostAccount {
	return s.HostAccount
}

func (s *GetHostAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHostAccountResponseBody) SetHostAccount(v *GetHostAccountResponseBodyHostAccount) *GetHostAccountResponseBody {
	s.HostAccount = v
	return s
}

func (s *GetHostAccountResponseBody) SetRequestId(v string) *GetHostAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostAccountResponseBody) Validate() error {
	if s.HostAccount != nil {
		if err := s.HostAccount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHostAccountResponseBodyHostAccount struct {
	// Indicates whether a password is configured for the host account. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	HasPassword *bool `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// The ID of the host account.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The name of the host account.
	//
	// example:
	//
	// abc
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host to which the host account belongs.
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
	// The name of the shared key.
	//
	// example:
	//
	// name
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	// The fingerprint of the private key.
	//
	// example:
	//
	// fe:ca:37:42:30:00:9d:95:e6:73:e5:b0:32:0a:**:**
	PrivateKeyFingerprint *string `json:"PrivateKeyFingerprint,omitempty" xml:"PrivateKeyFingerprint,omitempty"`
	PrivilegeType         *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The protocol that is used by the host. Valid values:
	//
	// 	- **SSH**
	//
	// 	- **RDP**
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	RotationMode *string `json:"RotationMode,omitempty" xml:"RotationMode,omitempty"`
}

func (s GetHostAccountResponseBodyHostAccount) String() string {
	return dara.Prettify(s)
}

func (s GetHostAccountResponseBodyHostAccount) GoString() string {
	return s.String()
}

func (s *GetHostAccountResponseBodyHostAccount) GetHasPassword() *bool {
	return s.HasPassword
}

func (s *GetHostAccountResponseBodyHostAccount) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *GetHostAccountResponseBodyHostAccount) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *GetHostAccountResponseBodyHostAccount) GetHostId() *string {
	return s.HostId
}

func (s *GetHostAccountResponseBodyHostAccount) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *GetHostAccountResponseBodyHostAccount) GetHostShareKeyName() *string {
	return s.HostShareKeyName
}

func (s *GetHostAccountResponseBodyHostAccount) GetPrivateKeyFingerprint() *string {
	return s.PrivateKeyFingerprint
}

func (s *GetHostAccountResponseBodyHostAccount) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *GetHostAccountResponseBodyHostAccount) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *GetHostAccountResponseBodyHostAccount) GetRotationMode() *string {
	return s.RotationMode
}

func (s *GetHostAccountResponseBodyHostAccount) SetHasPassword(v bool) *GetHostAccountResponseBodyHostAccount {
	s.HasPassword = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostAccountId(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostAccountId = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostAccountName(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostAccountName = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostId(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostId = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostShareKeyId(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostShareKeyId = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostShareKeyName(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostShareKeyName = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetPrivateKeyFingerprint(v string) *GetHostAccountResponseBodyHostAccount {
	s.PrivateKeyFingerprint = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetPrivilegeType(v string) *GetHostAccountResponseBodyHostAccount {
	s.PrivilegeType = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetProtocolName(v string) *GetHostAccountResponseBodyHostAccount {
	s.ProtocolName = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetRotationMode(v string) *GetHostAccountResponseBodyHostAccount {
	s.RotationMode = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) Validate() error {
	return dara.Validate(s)
}

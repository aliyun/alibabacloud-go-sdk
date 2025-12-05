// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHostAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccounts(v []*ListOperationHostAccountsResponseBodyHostAccounts) *ListOperationHostAccountsResponseBody
	GetHostAccounts() []*ListOperationHostAccountsResponseBodyHostAccounts
	SetRequestId(v string) *ListOperationHostAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOperationHostAccountsResponseBody
	GetTotalCount() *int64
}

type ListOperationHostAccountsResponseBody struct {
	// The host accounts returned.
	HostAccounts []*ListOperationHostAccountsResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of host accounts returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationHostAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHostAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationHostAccountsResponseBody) GetHostAccounts() []*ListOperationHostAccountsResponseBodyHostAccounts {
	return s.HostAccounts
}

func (s *ListOperationHostAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationHostAccountsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOperationHostAccountsResponseBody) SetHostAccounts(v []*ListOperationHostAccountsResponseBodyHostAccounts) *ListOperationHostAccountsResponseBody {
	s.HostAccounts = v
	return s
}

func (s *ListOperationHostAccountsResponseBody) SetRequestId(v string) *ListOperationHostAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationHostAccountsResponseBody) SetTotalCount(v int64) *ListOperationHostAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOperationHostAccountsResponseBody) Validate() error {
	if s.HostAccounts != nil {
		for _, item := range s.HostAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationHostAccountsResponseBodyHostAccounts struct {
	// Indicates whether a password is configured for the host account.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasPassword *bool `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// The host account ID.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The host account name.
	//
	// example:
	//
	// root
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The host ID.
	//
	// example:
	//
	// １
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the shared key that is associated with the host.
	//
	// example:
	//
	// 3
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The fingerprint of the private key for the host account.
	//
	// example:
	//
	// fe:ca:37:42:30:00:9d:95:e6:73:e5:b0:32:0a:**:**
	PrivateKeyFingerprint *string `json:"PrivateKeyFingerprint,omitempty" xml:"PrivateKeyFingerprint,omitempty"`
	// The protocol that is used by the host account.
	//
	// 	- **SSH**
	//
	// 	- **RDP**
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// Indicates whether the Secure File Transfer Protocol (SFTP) channels or the SSH channels are enabled for the host account that uses the SSH protocol.
	SSHConfig *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig `json:"SSHConfig,omitempty" xml:"SSHConfig,omitempty" type:"Struct"`
}

func (s ListOperationHostAccountsResponseBodyHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHostAccountsResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) GetHasPassword() *bool {
	return s.HasPassword
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) GetHostId() *string {
	return s.HostId
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) GetPrivateKeyFingerprint() *string {
	return s.PrivateKeyFingerprint
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) GetSSHConfig() *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig {
	return s.SSHConfig
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) SetHasPassword(v bool) *ListOperationHostAccountsResponseBodyHostAccounts {
	s.HasPassword = &v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) SetHostAccountId(v string) *ListOperationHostAccountsResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) SetHostAccountName(v string) *ListOperationHostAccountsResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) SetHostId(v string) *ListOperationHostAccountsResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) SetHostShareKeyId(v string) *ListOperationHostAccountsResponseBodyHostAccounts {
	s.HostShareKeyId = &v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) SetPrivateKeyFingerprint(v string) *ListOperationHostAccountsResponseBodyHostAccounts {
	s.PrivateKeyFingerprint = &v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) SetProtocolName(v string) *ListOperationHostAccountsResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) SetSSHConfig(v *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig) *ListOperationHostAccountsResponseBodyHostAccounts {
	s.SSHConfig = v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccounts) Validate() error {
	if s.SSHConfig != nil {
		if err := s.SSHConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOperationHostAccountsResponseBodyHostAccountsSSHConfig struct {
	// Indicates whether SFTP channels are enabled for the account.
	EnableSFTPChannel *bool `json:"EnableSFTPChannel,omitempty" xml:"EnableSFTPChannel,omitempty"`
	// Indicates whether SSH channels are enabled for the account.
	EnableSSHChannel *bool `json:"EnableSSHChannel,omitempty" xml:"EnableSSHChannel,omitempty"`
}

func (s ListOperationHostAccountsResponseBodyHostAccountsSSHConfig) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHostAccountsResponseBodyHostAccountsSSHConfig) GoString() string {
	return s.String()
}

func (s *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig) GetEnableSFTPChannel() *bool {
	return s.EnableSFTPChannel
}

func (s *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig) GetEnableSSHChannel() *bool {
	return s.EnableSSHChannel
}

func (s *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig) SetEnableSFTPChannel(v bool) *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig {
	s.EnableSFTPChannel = &v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig) SetEnableSSHChannel(v bool) *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig {
	s.EnableSSHChannel = &v
	return s
}

func (s *ListOperationHostAccountsResponseBodyHostAccountsSSHConfig) Validate() error {
	return dara.Validate(s)
}

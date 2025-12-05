// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccounts(v []*ListHostAccountsResponseBodyHostAccounts) *ListHostAccountsResponseBody
	GetHostAccounts() []*ListHostAccountsResponseBodyHostAccounts
	SetRequestId(v string) *ListHostAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHostAccountsResponseBody
	GetTotalCount() *int32
}

type ListHostAccountsResponseBody struct {
	// An array that consists of the queried host accounts.
	HostAccounts []*ListHostAccountsResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of host accounts that are queried.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsResponseBody) GetHostAccounts() []*ListHostAccountsResponseBodyHostAccounts {
	return s.HostAccounts
}

func (s *ListHostAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostAccountsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHostAccountsResponseBody) SetHostAccounts(v []*ListHostAccountsResponseBodyHostAccounts) *ListHostAccountsResponseBody {
	s.HostAccounts = v
	return s
}

func (s *ListHostAccountsResponseBody) SetRequestId(v string) *ListHostAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsResponseBody) SetTotalCount(v int32) *ListHostAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostAccountsResponseBody) Validate() error {
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

type ListHostAccountsResponseBodyHostAccounts struct {
	// Indicates whether a password is configured for the host account.
	//
	// Valid values:
	//
	// 	- true: A password is configured for the host account.
	//
	// 	- false: No passwords are configured for the host account.
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
	// The ID of the host.
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
	// The fingerprint of the private key for the host account.
	//
	// example:
	//
	// fe:ca:37:42:30:00:9d:95:e6:73:e5:b0:32:0a:**:**
	PrivateKeyFingerprint *string `json:"PrivateKeyFingerprint,omitempty" xml:"PrivateKeyFingerprint,omitempty"`
	PrivilegeType         *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The protocol that is used by the host.
	//
	// Valid values:
	//
	// 	- SSH
	//
	// 	- RDP
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	RotationMode *string `json:"RotationMode,omitempty" xml:"RotationMode,omitempty"`
}

func (s ListHostAccountsResponseBodyHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetHasPassword() *bool {
	return s.HasPassword
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetHostId() *string {
	return s.HostId
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetHostShareKeyId() *string {
	return s.HostShareKeyId
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetHostShareKeyName() *string {
	return s.HostShareKeyName
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetPrivateKeyFingerprint() *string {
	return s.PrivateKeyFingerprint
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListHostAccountsResponseBodyHostAccounts) GetRotationMode() *string {
	return s.RotationMode
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHasPassword(v bool) *ListHostAccountsResponseBodyHostAccounts {
	s.HasPassword = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostAccountId(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostShareKeyId(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostShareKeyId = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostShareKeyName(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostShareKeyName = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetPrivateKeyFingerprint(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.PrivateKeyFingerprint = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetPrivilegeType(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.PrivilegeType = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetRotationMode(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.RotationMode = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) Validate() error {
	return dara.Validate(s)
}

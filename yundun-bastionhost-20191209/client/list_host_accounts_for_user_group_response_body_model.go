// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsForUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccounts(v []*ListHostAccountsForUserGroupResponseBodyHostAccounts) *ListHostAccountsForUserGroupResponseBody
	GetHostAccounts() []*ListHostAccountsForUserGroupResponseBodyHostAccounts
	SetRequestId(v string) *ListHostAccountsForUserGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHostAccountsForUserGroupResponseBody
	GetTotalCount() *int32
}

type ListHostAccountsForUserGroupResponseBody struct {
	// The host accounts.
	HostAccounts []*ListHostAccountsForUserGroupResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the request.
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
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostAccountsForUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupResponseBody) GetHostAccounts() []*ListHostAccountsForUserGroupResponseBodyHostAccounts {
	return s.HostAccounts
}

func (s *ListHostAccountsForUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostAccountsForUserGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHostAccountsForUserGroupResponseBody) SetHostAccounts(v []*ListHostAccountsForUserGroupResponseBodyHostAccounts) *ListHostAccountsForUserGroupResponseBody {
	s.HostAccounts = v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBody) SetRequestId(v string) *ListHostAccountsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBody) SetTotalCount(v int32) *ListHostAccountsForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBody) Validate() error {
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

type ListHostAccountsForUserGroupResponseBodyHostAccounts struct {
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
	// host１
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The type of the host account.
	//
	// example:
	//
	// "Normal"
	HostAccountType *string `json:"HostAccountType,omitempty" xml:"HostAccountType,omitempty"`
	// The ID of the host to which the host account belongs.
	//
	// example:
	//
	// １
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// Indicates whether the host account is authorized for the current user group. Valid values:
	//
	// - **true**: Authorized.
	//
	// - **false**: Unauthorized.
	//
	// example:
	//
	// true
	IsAuthorized *bool `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	// The protocol name of the host account. Valid values:
	//
	// - **SSH**
	//
	// - **RDP**
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s ListHostAccountsForUserGroupResponseBodyHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForUserGroupResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) GetHostAccountType() *string {
	return s.HostAccountType
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) GetHostId() *string {
	return s.HostId
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) GetIsAuthorized() *bool {
	return s.IsAuthorized
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostAccountId(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostAccountType(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostAccountType = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetIsAuthorized(v bool) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.IsAuthorized = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) Validate() error {
	return dara.Validate(s)
}

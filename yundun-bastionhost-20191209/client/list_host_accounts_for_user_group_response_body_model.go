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
	// An array that consists of the queried host accounts.
	HostAccounts []*ListHostAccountsForUserGroupResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of host accounts that were queried.
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
	// The ID of the host for which the host accounts were queried.
	//
	// example:
	//
	// １
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// Indicates whether the user group is authorized to manage the host account. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	IsAuthorized *bool `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
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

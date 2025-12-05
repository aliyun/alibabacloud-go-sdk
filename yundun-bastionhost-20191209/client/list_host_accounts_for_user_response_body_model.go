// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccounts(v []*ListHostAccountsForUserResponseBodyHostAccounts) *ListHostAccountsForUserResponseBody
	GetHostAccounts() []*ListHostAccountsForUserResponseBodyHostAccounts
	SetRequestId(v string) *ListHostAccountsForUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHostAccountsForUserResponseBody
	GetTotalCount() *int32
}

type ListHostAccountsForUserResponseBody struct {
	// An array that consists of the queried host accounts.
	HostAccounts []*ListHostAccountsForUserResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
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

func (s ListHostAccountsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserResponseBody) GetHostAccounts() []*ListHostAccountsForUserResponseBodyHostAccounts {
	return s.HostAccounts
}

func (s *ListHostAccountsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostAccountsForUserResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHostAccountsForUserResponseBody) SetHostAccounts(v []*ListHostAccountsForUserResponseBodyHostAccounts) *ListHostAccountsForUserResponseBody {
	s.HostAccounts = v
	return s
}

func (s *ListHostAccountsForUserResponseBody) SetRequestId(v string) *ListHostAccountsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsForUserResponseBody) SetTotalCount(v int32) *ListHostAccountsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostAccountsForUserResponseBody) Validate() error {
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

type ListHostAccountsForUserResponseBodyHostAccounts struct {
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
	// root
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host for which the host accounts were queried.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// Indicates whether the user is authorized to manage the host account. Valid values:
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

func (s ListHostAccountsForUserResponseBodyHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForUserResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) GetHostId() *string {
	return s.HostId
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) GetIsAuthorized() *bool {
	return s.IsAuthorized
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetHostAccountId(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetIsAuthorized(v bool) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.IsAuthorized = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) Validate() error {
	return dara.Validate(s)
}

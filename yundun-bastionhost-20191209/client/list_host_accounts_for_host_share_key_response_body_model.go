// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostAccountsForHostShareKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostAccounts(v []*ListHostAccountsForHostShareKeyResponseBodyHostAccounts) *ListHostAccountsForHostShareKeyResponseBody
	GetHostAccounts() []*ListHostAccountsForHostShareKeyResponseBodyHostAccounts
	SetRequestId(v string) *ListHostAccountsForHostShareKeyResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListHostAccountsForHostShareKeyResponseBody
	GetTotalCount() *int64
}

type ListHostAccountsForHostShareKeyResponseBody struct {
	// An array that consists of the host accounts that are associated with the shared key.
	HostAccounts []*ListHostAccountsForHostShareKeyResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the host accounts that are associated with the shared key.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostAccountsForHostShareKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForHostShareKeyResponseBody) GetHostAccounts() []*ListHostAccountsForHostShareKeyResponseBodyHostAccounts {
	return s.HostAccounts
}

func (s *ListHostAccountsForHostShareKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHostAccountsForHostShareKeyResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListHostAccountsForHostShareKeyResponseBody) SetHostAccounts(v []*ListHostAccountsForHostShareKeyResponseBodyHostAccounts) *ListHostAccountsForHostShareKeyResponseBody {
	s.HostAccounts = v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBody) SetRequestId(v string) *ListHostAccountsForHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBody) SetTotalCount(v int64) *ListHostAccountsForHostShareKeyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBody) Validate() error {
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

type ListHostAccountsForHostShareKeyResponseBodyHostAccounts struct {
	// The name of the host account.
	//
	// example:
	//
	// root1234
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host.
	//
	// example:
	//
	// 1113
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the host account.
	//
	// example:
	//
	// 1235
	HostsAccountId *string `json:"HostsAccountId,omitempty" xml:"HostsAccountId,omitempty"`
	// The O\\&M protocol.
	//
	// example:
	//
	// SSH
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s ListHostAccountsForHostShareKeyResponseBodyHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListHostAccountsForHostShareKeyResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) GetHostId() *string {
	return s.HostId
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) GetHostsAccountId() *string {
	return s.HostsAccountId
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsForHostShareKeyResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsForHostShareKeyResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) SetHostsAccountId(v string) *ListHostAccountsForHostShareKeyResponseBodyHostAccounts {
	s.HostsAccountId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsForHostShareKeyResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) Validate() error {
	return dara.Validate(s)
}

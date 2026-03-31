// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMemberAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountInfos(v []*DescribeMemberAccountsResponseBodyAccountInfos) *DescribeMemberAccountsResponseBody
	GetAccountInfos() []*DescribeMemberAccountsResponseBodyAccountInfos
	SetRequestId(v string) *DescribeMemberAccountsResponseBody
	GetRequestId() *string
}

type DescribeMemberAccountsResponseBody struct {
	// The information about the member.
	AccountInfos []*DescribeMemberAccountsResponseBodyAccountInfos `json:"AccountInfos,omitempty" xml:"AccountInfos,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 50C4A80D-D46C-57E0-9A7D-03C0****4852
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMemberAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMemberAccountsResponseBody) GetAccountInfos() []*DescribeMemberAccountsResponseBodyAccountInfos {
	return s.AccountInfos
}

func (s *DescribeMemberAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMemberAccountsResponseBody) SetAccountInfos(v []*DescribeMemberAccountsResponseBodyAccountInfos) *DescribeMemberAccountsResponseBody {
	s.AccountInfos = v
	return s
}

func (s *DescribeMemberAccountsResponseBody) SetRequestId(v string) *DescribeMemberAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMemberAccountsResponseBody) Validate() error {
	if s.AccountInfos != nil {
		for _, item := range s.AccountInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMemberAccountsResponseBodyAccountInfos struct {
	// The ID of the member.
	//
	// example:
	//
	// 169************21
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member.
	//
	// example:
	//
	// ipflgmqqnbjg
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The status of the member.
	//
	// 	- **enabled**: managed.
	//
	// 	- **disabled**: not managed.
	//
	// 	- **disabling**: being deleted.
	//
	// example:
	//
	// enabled
	AccountStatus *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	// The description of the member.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the member was added.
	//
	// example:
	//
	// 1683367751000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s DescribeMemberAccountsResponseBodyAccountInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeMemberAccountsResponseBodyAccountInfos) GoString() string {
	return s.String()
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) GetDescription() *string {
	return s.Description
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetAccountId(v string) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.AccountId = &v
	return s
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetAccountName(v string) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.AccountName = &v
	return s
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetAccountStatus(v string) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.AccountStatus = &v
	return s
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetDescription(v string) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.Description = &v
	return s
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) SetGmtCreate(v int64) *DescribeMemberAccountsResponseBodyAccountInfos {
	s.GmtCreate = &v
	return s
}

func (s *DescribeMemberAccountsResponseBodyAccountInfos) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnrolledAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnrolledAccounts(v []*ListEnrolledAccountsResponseBodyEnrolledAccounts) *ListEnrolledAccountsResponseBody
	GetEnrolledAccounts() []*ListEnrolledAccountsResponseBodyEnrolledAccounts
	SetNextToken(v string) *ListEnrolledAccountsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEnrolledAccountsResponseBody
	GetRequestId() *string
}

type ListEnrolledAccountsResponseBody struct {
	// The enrolled accounts.
	EnrolledAccounts []*ListEnrolledAccountsResponseBodyEnrolledAccounts `json:"EnrolledAccounts,omitempty" xml:"EnrolledAccounts,omitempty" type:"Repeated"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh62Xr2DzYbz/SAfc*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 768F908D-A66A-5A5D-816C-20C93CBBFEE3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEnrolledAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnrolledAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsResponseBody) GetEnrolledAccounts() []*ListEnrolledAccountsResponseBodyEnrolledAccounts {
	return s.EnrolledAccounts
}

func (s *ListEnrolledAccountsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEnrolledAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnrolledAccountsResponseBody) SetEnrolledAccounts(v []*ListEnrolledAccountsResponseBodyEnrolledAccounts) *ListEnrolledAccountsResponseBody {
	s.EnrolledAccounts = v
	return s
}

func (s *ListEnrolledAccountsResponseBody) SetNextToken(v string) *ListEnrolledAccountsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEnrolledAccountsResponseBody) SetRequestId(v string) *ListEnrolledAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnrolledAccountsResponseBody) Validate() error {
	if s.EnrolledAccounts != nil {
		for _, item := range s.EnrolledAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEnrolledAccountsResponseBodyEnrolledAccounts struct {
	// The account ID.
	//
	// example:
	//
	// 19534534552*****
	AccountUid *int64 `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	// The ID of the baseline that is implemented.
	//
	// example:
	//
	// afb-bp1durvn3lgqe28v****
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// TestAccount
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// fd-5ESoku****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The ID of the settlement account.
	//
	// example:
	//
	// 13161210500*****
	PayerAccountUid *int64 `json:"PayerAccountUid,omitempty" xml:"PayerAccountUid,omitempty"`
	// The creation status. Valid values:
	//
	// 	- Pending: The account is pending to be created.
	//
	// 	- Running: The account is being created.
	//
	// 	- Finished: The account is created.
	//
	// 	- Failed: The account fails to be created.
	//
	// 	- Scheduling: The account is being scheduled.
	//
	// 	- ScheduleFailed: The account fails to be scheduled.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2021-11-01T02:38:27Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEnrolledAccountsResponseBodyEnrolledAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListEnrolledAccountsResponseBodyEnrolledAccounts) GoString() string {
	return s.String()
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) GetAccountUid() *int64 {
	return s.AccountUid
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) GetBaselineId() *string {
	return s.BaselineId
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) GetFolderId() *string {
	return s.FolderId
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) GetPayerAccountUid() *int64 {
	return s.PayerAccountUid
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) GetStatus() *string {
	return s.Status
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetAccountUid(v int64) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.AccountUid = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetBaselineId(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.BaselineId = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetCreateTime(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.CreateTime = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetDisplayName(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.DisplayName = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetFolderId(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.FolderId = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetPayerAccountUid(v int64) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.PayerAccountUid = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetStatus(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.Status = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) SetUpdateTime(v string) *ListEnrolledAccountsResponseBodyEnrolledAccounts {
	s.UpdateTime = &v
	return s
}

func (s *ListEnrolledAccountsResponseBodyEnrolledAccounts) Validate() error {
	return dara.Validate(s)
}

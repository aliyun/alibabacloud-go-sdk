// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLakebaseS3AccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListLakebaseS3AccountsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLakebaseS3AccountsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLakebaseS3AccountsResponseBody
	GetRequestId() *string
	SetS3Accounts(v []*ListLakebaseS3AccountsResponseBodyS3Accounts) *ListLakebaseS3AccountsResponseBody
	GetS3Accounts() []*ListLakebaseS3AccountsResponseBodyS3Accounts
	SetTotalCount(v int32) *ListLakebaseS3AccountsResponseBody
	GetTotalCount() *int32
}

type ListLakebaseS3AccountsResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of S3 accounts.
	S3Accounts []*ListLakebaseS3AccountsResponseBodyS3Accounts `json:"S3Accounts,omitempty" xml:"S3Accounts,omitempty" type:"Repeated"`
	// The total number of accounts.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLakebaseS3AccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLakebaseS3AccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLakebaseS3AccountsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLakebaseS3AccountsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLakebaseS3AccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLakebaseS3AccountsResponseBody) GetS3Accounts() []*ListLakebaseS3AccountsResponseBodyS3Accounts {
	return s.S3Accounts
}

func (s *ListLakebaseS3AccountsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLakebaseS3AccountsResponseBody) SetPageNumber(v int32) *ListLakebaseS3AccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLakebaseS3AccountsResponseBody) SetPageSize(v int32) *ListLakebaseS3AccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLakebaseS3AccountsResponseBody) SetRequestId(v string) *ListLakebaseS3AccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLakebaseS3AccountsResponseBody) SetS3Accounts(v []*ListLakebaseS3AccountsResponseBodyS3Accounts) *ListLakebaseS3AccountsResponseBody {
	s.S3Accounts = v
	return s
}

func (s *ListLakebaseS3AccountsResponseBody) SetTotalCount(v int32) *ListLakebaseS3AccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLakebaseS3AccountsResponseBody) Validate() error {
	if s.S3Accounts != nil {
		for _, item := range s.S3Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLakebaseS3AccountsResponseBodyS3Accounts struct {
	// The account type. Valid values:
	//
	// - default: the built-in default account.
	//
	// - user: a user-created account.
	//
	// example:
	//
	// default
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The access key of the S3 account.
	//
	// example:
	//
	// accname
	UserAccAk *string `json:"UserAccAk,omitempty" xml:"UserAccAk,omitempty"`
	// The secret key of the S3 account (displayed in masked format).
	//
	// example:
	//
	// password***
	UserAccSk *string `json:"UserAccSk,omitempty" xml:"UserAccSk,omitempty"`
}

func (s ListLakebaseS3AccountsResponseBodyS3Accounts) String() string {
	return dara.Prettify(s)
}

func (s ListLakebaseS3AccountsResponseBodyS3Accounts) GoString() string {
	return s.String()
}

func (s *ListLakebaseS3AccountsResponseBodyS3Accounts) GetAccountType() *string {
	return s.AccountType
}

func (s *ListLakebaseS3AccountsResponseBodyS3Accounts) GetUserAccAk() *string {
	return s.UserAccAk
}

func (s *ListLakebaseS3AccountsResponseBodyS3Accounts) GetUserAccSk() *string {
	return s.UserAccSk
}

func (s *ListLakebaseS3AccountsResponseBodyS3Accounts) SetAccountType(v string) *ListLakebaseS3AccountsResponseBodyS3Accounts {
	s.AccountType = &v
	return s
}

func (s *ListLakebaseS3AccountsResponseBodyS3Accounts) SetUserAccAk(v string) *ListLakebaseS3AccountsResponseBodyS3Accounts {
	s.UserAccAk = &v
	return s
}

func (s *ListLakebaseS3AccountsResponseBodyS3Accounts) SetUserAccSk(v string) *ListLakebaseS3AccountsResponseBodyS3Accounts {
	s.UserAccSk = &v
	return s
}

func (s *ListLakebaseS3AccountsResponseBodyS3Accounts) Validate() error {
	return dara.Validate(s)
}

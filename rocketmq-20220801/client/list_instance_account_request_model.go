// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountStatus(v string) *ListInstanceAccountRequest
	GetAccountStatus() *string
	SetAccountType(v string) *ListInstanceAccountRequest
	GetAccountType() *string
	SetPageNumber(v int32) *ListInstanceAccountRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstanceAccountRequest
	GetPageSize() *int32
	SetUsername(v string) *ListInstanceAccountRequest
	GetUsername() *string
}

type ListInstanceAccountRequest struct {
	// The status of the account.
	//
	// Valid values:
	//
	// 	- DISABLE
	//
	// 	- ENABLE
	//
	// example:
	//
	// ENABLE
	AccountStatus *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	// The account type.
	//
	//   - CUSTOMER
	//
	//   - DEFAULT
	//
	// example:
	//
	// CUSTOMER
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListInstanceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountRequest) GetAccountStatus() *string {
	return s.AccountStatus
}

func (s *ListInstanceAccountRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *ListInstanceAccountRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstanceAccountRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstanceAccountRequest) GetUsername() *string {
	return s.Username
}

func (s *ListInstanceAccountRequest) SetAccountStatus(v string) *ListInstanceAccountRequest {
	s.AccountStatus = &v
	return s
}

func (s *ListInstanceAccountRequest) SetAccountType(v string) *ListInstanceAccountRequest {
	s.AccountType = &v
	return s
}

func (s *ListInstanceAccountRequest) SetPageNumber(v int32) *ListInstanceAccountRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceAccountRequest) SetPageSize(v int32) *ListInstanceAccountRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceAccountRequest) SetUsername(v string) *ListInstanceAccountRequest {
	s.Username = &v
	return s
}

func (s *ListInstanceAccountRequest) Validate() error {
	return dara.Validate(s)
}

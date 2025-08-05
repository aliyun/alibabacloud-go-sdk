// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCrossAccountsResponseBody
	GetCode() *string
	SetCrossAccounts(v *DescribeCrossAccountsResponseBodyCrossAccounts) *DescribeCrossAccountsResponseBody
	GetCrossAccounts() *DescribeCrossAccountsResponseBodyCrossAccounts
	SetMessage(v string) *DescribeCrossAccountsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeCrossAccountsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCrossAccountsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCrossAccountsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCrossAccountsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeCrossAccountsResponseBody
	GetTotalCount() *int64
}

type DescribeCrossAccountsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the accounts used in cross-account backup.
	CrossAccounts *DescribeCrossAccountsResponseBodyCrossAccounts `json:"CrossAccounts,omitempty" xml:"CrossAccounts,omitempty" type:"Struct"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 22
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCrossAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCrossAccountsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCrossAccountsResponseBody) GetCrossAccounts() *DescribeCrossAccountsResponseBodyCrossAccounts {
	return s.CrossAccounts
}

func (s *DescribeCrossAccountsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCrossAccountsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCrossAccountsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCrossAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCrossAccountsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCrossAccountsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCrossAccountsResponseBody) SetCode(v string) *DescribeCrossAccountsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCrossAccountsResponseBody) SetCrossAccounts(v *DescribeCrossAccountsResponseBodyCrossAccounts) *DescribeCrossAccountsResponseBody {
	s.CrossAccounts = v
	return s
}

func (s *DescribeCrossAccountsResponseBody) SetMessage(v string) *DescribeCrossAccountsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCrossAccountsResponseBody) SetPageNumber(v int32) *DescribeCrossAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCrossAccountsResponseBody) SetPageSize(v int32) *DescribeCrossAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCrossAccountsResponseBody) SetRequestId(v string) *DescribeCrossAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCrossAccountsResponseBody) SetSuccess(v bool) *DescribeCrossAccountsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCrossAccountsResponseBody) SetTotalCount(v int64) *DescribeCrossAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCrossAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCrossAccountsResponseBodyCrossAccounts struct {
	CrossAccount []*DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount `json:"CrossAccount,omitempty" xml:"CrossAccount,omitempty" type:"Repeated"`
}

func (s DescribeCrossAccountsResponseBodyCrossAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossAccountsResponseBodyCrossAccounts) GoString() string {
	return s.String()
}

func (s *DescribeCrossAccountsResponseBodyCrossAccounts) GetCrossAccount() []*DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount {
	return s.CrossAccount
}

func (s *DescribeCrossAccountsResponseBodyCrossAccounts) SetCrossAccount(v []*DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) *DescribeCrossAccountsResponseBodyCrossAccounts {
	s.CrossAccount = v
	return s
}

func (s *DescribeCrossAccountsResponseBodyCrossAccounts) Validate() error {
	return dara.Validate(s)
}

type DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount struct {
	// The account alias. The value can be up to 32 bits in length.
	//
	// example:
	//
	// content
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The time when the account was created. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1654570439
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The name of the RAM role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 15897534xxxxx625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The ID of the backup type.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the current account.
	//
	// example:
	//
	// 184164xxxxx49795
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The time when the account information was updated. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1640157098
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) GoString() string {
	return s.String()
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) GetAlias() *string {
	return s.Alias
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) GetId() *int64 {
	return s.Id
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) SetAlias(v string) *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount {
	s.Alias = &v
	return s
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) SetCreatedTime(v int64) *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount {
	s.CreatedTime = &v
	return s
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) SetCrossAccountRoleName(v string) *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) SetCrossAccountUserId(v int64) *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) SetId(v int64) *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount {
	s.Id = &v
	return s
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) SetOwnerId(v int64) *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount {
	s.OwnerId = &v
	return s
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) SetUpdatedTime(v int64) *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeCrossAccountsResponseBodyCrossAccountsCrossAccount) Validate() error {
	return dara.Validate(s)
}

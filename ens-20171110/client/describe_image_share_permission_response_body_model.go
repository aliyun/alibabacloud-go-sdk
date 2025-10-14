// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSharePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v *DescribeImageSharePermissionResponseBodyAccounts) *DescribeImageSharePermissionResponseBody
	GetAccounts() *DescribeImageSharePermissionResponseBodyAccounts
	SetImageId(v string) *DescribeImageSharePermissionResponseBody
	GetImageId() *string
	SetPageNumber(v int32) *DescribeImageSharePermissionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeImageSharePermissionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeImageSharePermissionResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImageSharePermissionResponseBody
	GetTotalCount() *int32
}

type DescribeImageSharePermissionResponseBody struct {
	// The account information.
	Accounts *DescribeImageSharePermissionResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	// The ID of the image.
	//
	// example:
	//
	// m-5qkf6jv9a0tzd5ipwx5fi****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A44EE357-6174-5E37-A801-48F5790F9ACE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of Alibaba Cloud accounts with which you share the image.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageSharePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBody) GetAccounts() *DescribeImageSharePermissionResponseBodyAccounts {
	return s.Accounts
}

func (s *DescribeImageSharePermissionResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageSharePermissionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeImageSharePermissionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageSharePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageSharePermissionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageSharePermissionResponseBody) SetAccounts(v *DescribeImageSharePermissionResponseBodyAccounts) *DescribeImageSharePermissionResponseBody {
	s.Accounts = v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetImageId(v string) *DescribeImageSharePermissionResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetPageNumber(v int32) *DescribeImageSharePermissionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetPageSize(v int32) *DescribeImageSharePermissionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetRequestId(v string) *DescribeImageSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) SetTotalCount(v int32) *DescribeImageSharePermissionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBody) Validate() error {
	if s.Accounts != nil {
		if err := s.Accounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageSharePermissionResponseBodyAccounts struct {
	Account []*DescribeImageSharePermissionResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s DescribeImageSharePermissionResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBodyAccounts) GetAccount() []*DescribeImageSharePermissionResponseBodyAccountsAccount {
	return s.Account
}

func (s *DescribeImageSharePermissionResponseBodyAccounts) SetAccount(v []*DescribeImageSharePermissionResponseBodyAccountsAccount) *DescribeImageSharePermissionResponseBodyAccounts {
	s.Account = v
	return s
}

func (s *DescribeImageSharePermissionResponseBodyAccounts) Validate() error {
	if s.Account != nil {
		for _, item := range s.Account {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageSharePermissionResponseBodyAccountsAccount struct {
	// The Alibaba Cloud account with which you share the image.
	//
	// example:
	//
	// 1515285523xxxx
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
}

func (s DescribeImageSharePermissionResponseBodyAccountsAccount) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSharePermissionResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeImageSharePermissionResponseBodyAccountsAccount) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeImageSharePermissionResponseBodyAccountsAccount) SetAliyunUid(v string) *DescribeImageSharePermissionResponseBodyAccountsAccount {
	s.AliyunUid = &v
	return s
}

func (s *DescribeImageSharePermissionResponseBodyAccountsAccount) Validate() error {
	return dara.Validate(s)
}

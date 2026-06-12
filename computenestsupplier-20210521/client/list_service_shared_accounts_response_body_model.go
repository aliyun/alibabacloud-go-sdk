// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceSharedAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServiceSharedAccountsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceSharedAccountsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceSharedAccountsResponseBody
	GetRequestId() *string
	SetShareAccount(v []*ListServiceSharedAccountsResponseBodyShareAccount) *ListServiceSharedAccountsResponseBody
	GetShareAccount() []*ListServiceSharedAccountsResponseBodyShareAccount
	SetTotalCount(v int32) *ListServiceSharedAccountsResponseBody
	GetTotalCount() *int32
}

type ListServiceSharedAccountsResponseBody struct {
	// The number of entries returned on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA3AE512-6D30-549A-B52D-B9042CA8D515
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the shared accounts.
	ShareAccount []*ListServiceSharedAccountsResponseBodyShareAccount `json:"ShareAccount,omitempty" xml:"ShareAccount,omitempty" type:"Repeated"`
	// The total number of entries that meet the filter criteria.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceSharedAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSharedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceSharedAccountsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceSharedAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceSharedAccountsResponseBody) GetShareAccount() []*ListServiceSharedAccountsResponseBodyShareAccount {
	return s.ShareAccount
}

func (s *ListServiceSharedAccountsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServiceSharedAccountsResponseBody) SetMaxResults(v int32) *ListServiceSharedAccountsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBody) SetNextToken(v string) *ListServiceSharedAccountsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBody) SetRequestId(v string) *ListServiceSharedAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBody) SetShareAccount(v []*ListServiceSharedAccountsResponseBodyShareAccount) *ListServiceSharedAccountsResponseBody {
	s.ShareAccount = v
	return s
}

func (s *ListServiceSharedAccountsResponseBody) SetTotalCount(v int32) *ListServiceSharedAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBody) Validate() error {
	if s.ShareAccount != nil {
		for _, item := range s.ShareAccount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceSharedAccountsResponseBodyShareAccount struct {
	// The time when the sharing was created.
	//
	// example:
	//
	// 2021-12-28T02:47:46.000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The logo of the distributor.
	//
	// example:
	//
	// logo
	Logo *string `json:"Logo,omitempty" xml:"Logo,omitempty"`
	// The name of the distributor.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permission type. Valid values:
	//
	// - Deployable: The service is deployable.
	//
	// - Accessible: The service is accessible.
	//
	// example:
	//
	// Deployable
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-e10349089de34exxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The time when the sharing was last updated.
	//
	// example:
	//
	// 2023-02-13T02:16:03.756Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The Alibaba Cloud account ID of the user.
	//
	// example:
	//
	// 127383705xxxxxx
	UserAliUid *string `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
}

func (s ListServiceSharedAccountsResponseBodyShareAccount) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSharedAccountsResponseBodyShareAccount) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) GetLogo() *string {
	return s.Logo
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) GetName() *string {
	return s.Name
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) GetPermission() *string {
	return s.Permission
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) GetUserAliUid() *string {
	return s.UserAliUid
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetCreateTime(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.CreateTime = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetLogo(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.Logo = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetName(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.Name = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetPermission(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.Permission = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetServiceId(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.ServiceId = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetUpdateTime(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) SetUserAliUid(v string) *ListServiceSharedAccountsResponseBodyShareAccount {
	s.UserAliUid = &v
	return s
}

func (s *ListServiceSharedAccountsResponseBodyShareAccount) Validate() error {
	return dara.Validate(s)
}

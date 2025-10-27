// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccesskeyLeakListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyLeakList(v []*DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) *DescribeAccesskeyLeakListResponseBody
	GetAccessKeyLeakList() []*DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList
	SetAkLeakCount(v int32) *DescribeAccesskeyLeakListResponseBody
	GetAkLeakCount() *int32
	SetCurrentPage(v int32) *DescribeAccesskeyLeakListResponseBody
	GetCurrentPage() *int32
	SetGmtLast(v int64) *DescribeAccesskeyLeakListResponseBody
	GetGmtLast() *int64
	SetPageSize(v int32) *DescribeAccesskeyLeakListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccesskeyLeakListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAccesskeyLeakListResponseBody
	GetTotalCount() *int32
}

type DescribeAccesskeyLeakListResponseBody struct {
	// An array that consists of the details about AccessKey pair leaks.
	AccessKeyLeakList []*DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList `json:"AccessKeyLeakList,omitempty" xml:"AccessKeyLeakList,omitempty" type:"Repeated"`
	// The number of AccessKey pair leaks that are unhandled.
	//
	// example:
	//
	// 1
	AkLeakCount *int32 `json:"AkLeakCount,omitempty" xml:"AkLeakCount,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1612357897000
	GmtLast *int64 `json:"GmtLast,omitempty" xml:"GmtLast,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B37C9052-A73E-4707-A024-9247702852BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of AccessKey pair leaks.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccesskeyLeakListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccesskeyLeakListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccesskeyLeakListResponseBody) GetAccessKeyLeakList() []*DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	return s.AccessKeyLeakList
}

func (s *DescribeAccesskeyLeakListResponseBody) GetAkLeakCount() *int32 {
	return s.AkLeakCount
}

func (s *DescribeAccesskeyLeakListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAccesskeyLeakListResponseBody) GetGmtLast() *int64 {
	return s.GmtLast
}

func (s *DescribeAccesskeyLeakListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccesskeyLeakListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccesskeyLeakListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccesskeyLeakListResponseBody) SetAccessKeyLeakList(v []*DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) *DescribeAccesskeyLeakListResponseBody {
	s.AccessKeyLeakList = v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetAkLeakCount(v int32) *DescribeAccesskeyLeakListResponseBody {
	s.AkLeakCount = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetCurrentPage(v int32) *DescribeAccesskeyLeakListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetGmtLast(v int64) *DescribeAccesskeyLeakListResponseBody {
	s.GmtLast = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetPageSize(v int32) *DescribeAccesskeyLeakListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetRequestId(v string) *DescribeAccesskeyLeakListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) SetTotalCount(v int32) *DescribeAccesskeyLeakListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBody) Validate() error {
	if s.AccessKeyLeakList != nil {
		for _, item := range s.AccessKeyLeakList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList struct {
	// The ID of the AccessKey pair that is leaked.
	//
	// example:
	//
	// yourAccessKeyID
	AccesskeyId *string `json:"AccesskeyId,omitempty" xml:"AccesskeyId,omitempty"`
	// The name of the Alibaba Cloud account that is affected.
	//
	// example:
	//
	// testAccountName
	AliUserName *string `json:"AliUserName,omitempty" xml:"AliUserName,omitempty"`
	// The platform to which the asset belongs. The value is fixed as **Cloud platform**.
	//
	// example:
	//
	// Cloud platform
	Asset *string `json:"Asset,omitempty" xml:"Asset,omitempty"`
	// The time when the AccessKey pair leak is handled.
	//
	// example:
	//
	// 2020-12-03 21:23:38
	DealTime *string `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// The method to handle the AccessKey pair leak. Valid values:
	//
	// 	- **pending**: The AccessKey pair leak is unhandled.
	//
	// 	- **manual**: The AccessKey pair leak is manually handled.
	//
	// 	- **disable**: The AccessKey pair leak is disabled.
	//
	// 	- **add-whitelist**: The AccessKey pair leak is added to the whitelist.
	//
	// example:
	//
	// pending
	DealType *string `json:"DealType,omitempty" xml:"DealType,omitempty"`
	// The time when the AccessKey pair leak is first detected. The value of this parameter is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1612357897000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary key ID of the database.
	//
	// example:
	//
	// 389357
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the AccessKey pair leak is handled. Valid values:
	//
	// 	- **pending**: unhandled
	//
	// 	- **dealed**: handled
	//
	// example:
	//
	// pending
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the leak. The value is fixed as **AccessKey**.
	//
	// example:
	//
	// AccessKey
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the platform on which the AccessKey pair leak is detected.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The type of the account to which the leaked AccessKey pair belongs. Valid values:
	//
	// 	- **master**: Alibaba Cloud account
	//
	// 	- **ram**: RAM user
	//
	// example:
	//
	// master
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GoString() string {
	return s.String()
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetAccesskeyId() *string {
	return s.AccesskeyId
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetAliUserName() *string {
	return s.AliUserName
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetAsset() *string {
	return s.Asset
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetDealTime() *string {
	return s.DealTime
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetDealType() *string {
	return s.DealType
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetId() *int64 {
	return s.Id
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetStatus() *string {
	return s.Status
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetType() *string {
	return s.Type
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetUrl() *string {
	return s.Url
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) GetUserType() *string {
	return s.UserType
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetAccesskeyId(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.AccesskeyId = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetAliUserName(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.AliUserName = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetAsset(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Asset = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetDealTime(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.DealTime = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetDealType(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.DealType = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetGmtModified(v int64) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.GmtModified = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetId(v int64) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Id = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetStatus(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Status = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetType(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Type = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetUrl(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.Url = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) SetUserType(v string) *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList {
	s.UserType = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponseBodyAccessKeyLeakList) Validate() error {
	return dara.Validate(s)
}

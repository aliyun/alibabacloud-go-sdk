// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerLockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *ListServerLockResponseBody
	GetCurrentPageNum() *int32
	SetData(v []*ListServerLockResponseBodyData) *ListServerLockResponseBody
	GetData() []*ListServerLockResponseBodyData
	SetNextPage(v bool) *ListServerLockResponseBody
	GetNextPage() *bool
	SetPageSize(v int32) *ListServerLockResponseBody
	GetPageSize() *int32
	SetPrePage(v bool) *ListServerLockResponseBody
	GetPrePage() *bool
	SetRequestId(v string) *ListServerLockResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *ListServerLockResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *ListServerLockResponseBody
	GetTotalPageNum() *int32
}

type ListServerLockResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 0
	CurrentPageNum *int32 `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	// The response parameters.
	Data []*ListServerLockResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Indicates whether the current page is followed by a page.
	//
	// example:
	//
	// false
	NextPage *bool `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Indicates whether the current page is preceded by a page.
	//
	// example:
	//
	// false
	PrePage *bool `json:"PrePage,omitempty" xml:"PrePage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9DFCF6F8-243C-****-8035-4B12FEFD7D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListServerLockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServerLockResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerLockResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListServerLockResponseBody) GetData() []*ListServerLockResponseBodyData {
	return s.Data
}

func (s *ListServerLockResponseBody) GetNextPage() *bool {
	return s.NextPage
}

func (s *ListServerLockResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServerLockResponseBody) GetPrePage() *bool {
	return s.PrePage
}

func (s *ListServerLockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServerLockResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListServerLockResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListServerLockResponseBody) SetCurrentPageNum(v int32) *ListServerLockResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListServerLockResponseBody) SetData(v []*ListServerLockResponseBodyData) *ListServerLockResponseBody {
	s.Data = v
	return s
}

func (s *ListServerLockResponseBody) SetNextPage(v bool) *ListServerLockResponseBody {
	s.NextPage = &v
	return s
}

func (s *ListServerLockResponseBody) SetPageSize(v int32) *ListServerLockResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServerLockResponseBody) SetPrePage(v bool) *ListServerLockResponseBody {
	s.PrePage = &v
	return s
}

func (s *ListServerLockResponseBody) SetRequestId(v string) *ListServerLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerLockResponseBody) SetTotalItemNum(v int32) *ListServerLockResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListServerLockResponseBody) SetTotalPageNum(v int32) *ListServerLockResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListServerLockResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServerLockResponseBodyData struct {
	// The instance ID of the domain name.
	//
	// example:
	//
	// S20190N1DAI4****
	DomainInstanceId *string `json:"DomainInstanceId,omitempty" xml:"DomainInstanceId,omitempty"`
	// The domain name that has valid registry lock information.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2030-07-10 17:37:36
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-02-19 16:38:07
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the domain name was last modified.
	//
	// example:
	//
	// 2022-02-19 16:40:38
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The instance ID of the domain name for which the registry lock is enabled.
	//
	// example:
	//
	// S2021591IQ28****
	LockInstanceId *string `json:"LockInstanceId,omitempty" xml:"LockInstanceId,omitempty"`
	// The ID of the product to which the domain name with the registry lock enabled belongs.
	//
	// example:
	//
	// 1807**
	LockProductId *string `json:"LockProductId,omitempty" xml:"LockProductId,omitempty"`
	// The status of the registry lock.
	//
	// example:
	//
	// 2
	ServerLockStatus *string `json:"ServerLockStatus,omitempty" xml:"ServerLockStatus,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2021-07-10 17:37:36
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 121000000****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListServerLockResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListServerLockResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServerLockResponseBodyData) GetDomainInstanceId() *string {
	return s.DomainInstanceId
}

func (s *ListServerLockResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *ListServerLockResponseBodyData) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ListServerLockResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServerLockResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServerLockResponseBodyData) GetLockInstanceId() *string {
	return s.LockInstanceId
}

func (s *ListServerLockResponseBodyData) GetLockProductId() *string {
	return s.LockProductId
}

func (s *ListServerLockResponseBodyData) GetServerLockStatus() *string {
	return s.ServerLockStatus
}

func (s *ListServerLockResponseBodyData) GetStartDate() *string {
	return s.StartDate
}

func (s *ListServerLockResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListServerLockResponseBodyData) SetDomainInstanceId(v string) *ListServerLockResponseBodyData {
	s.DomainInstanceId = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetDomainName(v string) *ListServerLockResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetExpireDate(v string) *ListServerLockResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetGmtCreate(v string) *ListServerLockResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetGmtModified(v string) *ListServerLockResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetLockInstanceId(v string) *ListServerLockResponseBodyData {
	s.LockInstanceId = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetLockProductId(v string) *ListServerLockResponseBodyData {
	s.LockProductId = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetServerLockStatus(v string) *ListServerLockResponseBodyData {
	s.ServerLockStatus = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetStartDate(v string) *ListServerLockResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *ListServerLockResponseBodyData) SetUserId(v string) *ListServerLockResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListServerLockResponseBodyData) Validate() error {
	return dara.Validate(s)
}

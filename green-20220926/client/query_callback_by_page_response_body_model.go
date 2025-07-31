// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCallbackByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryCallbackByPageResponseBody
	GetCurrentPage() *int32
	SetItems(v []*QueryCallbackByPageResponseBodyItems) *QueryCallbackByPageResponseBody
	GetItems() []*QueryCallbackByPageResponseBodyItems
	SetPageSize(v int32) *QueryCallbackByPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryCallbackByPageResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *QueryCallbackByPageResponseBody
	GetTotalCount() *int64
}

type QueryCallbackByPageResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*QueryCallbackByPageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCallbackByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCallbackByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallbackByPageResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryCallbackByPageResponseBody) GetItems() []*QueryCallbackByPageResponseBodyItems {
	return s.Items
}

func (s *QueryCallbackByPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCallbackByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCallbackByPageResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryCallbackByPageResponseBody) SetCurrentPage(v int32) *QueryCallbackByPageResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryCallbackByPageResponseBody) SetItems(v []*QueryCallbackByPageResponseBodyItems) *QueryCallbackByPageResponseBody {
	s.Items = v
	return s
}

func (s *QueryCallbackByPageResponseBody) SetPageSize(v int32) *QueryCallbackByPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryCallbackByPageResponseBody) SetRequestId(v string) *QueryCallbackByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallbackByPageResponseBody) SetTotalCount(v int64) *QueryCallbackByPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryCallbackByPageResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryCallbackByPageResponseBodyItems struct {
	// example:
	//
	// SHA256
	CryptType *string `json:"CryptType,omitempty" xml:"CryptType,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-06-03 15:20:14
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1697
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// all
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Seed。
	//
	// example:
	//
	// cb6gYS8GXj4Vn4Y4FN0Y8R5M-1x46Mq
	Seed *string `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// UID。
	//
	// example:
	//
	// 12161*****398900
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// example:
	//
	// https://console.aliyun.com/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s QueryCallbackByPageResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s QueryCallbackByPageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *QueryCallbackByPageResponseBodyItems) GetCryptType() *string {
	return s.CryptType
}

func (s *QueryCallbackByPageResponseBodyItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryCallbackByPageResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryCallbackByPageResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *QueryCallbackByPageResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *QueryCallbackByPageResponseBodyItems) GetScope() *string {
	return s.Scope
}

func (s *QueryCallbackByPageResponseBodyItems) GetSeed() *string {
	return s.Seed
}

func (s *QueryCallbackByPageResponseBodyItems) GetUid() *string {
	return s.Uid
}

func (s *QueryCallbackByPageResponseBodyItems) GetUrl() *string {
	return s.Url
}

func (s *QueryCallbackByPageResponseBodyItems) SetCryptType(v string) *QueryCallbackByPageResponseBodyItems {
	s.CryptType = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetGmtCreate(v string) *QueryCallbackByPageResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetGmtModified(v string) *QueryCallbackByPageResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetId(v int64) *QueryCallbackByPageResponseBodyItems {
	s.Id = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetName(v string) *QueryCallbackByPageResponseBodyItems {
	s.Name = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetScope(v string) *QueryCallbackByPageResponseBodyItems {
	s.Scope = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetSeed(v string) *QueryCallbackByPageResponseBodyItems {
	s.Seed = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetUid(v string) *QueryCallbackByPageResponseBodyItems {
	s.Uid = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) SetUrl(v string) *QueryCallbackByPageResponseBodyItems {
	s.Url = &v
	return s
}

func (s *QueryCallbackByPageResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

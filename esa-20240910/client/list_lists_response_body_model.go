// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItemsUsage(v int64) *ListListsResponseBody
	GetItemsUsage() *int64
	SetLists(v []*ListListsResponseBodyLists) *ListListsResponseBody
	GetLists() []*ListListsResponseBodyLists
	SetPageNumber(v int32) *ListListsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListListsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListListsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListListsResponseBody
	GetTotalCount() *int32
	SetUsage(v int64) *ListListsResponseBody
	GetUsage() *int64
}

type ListListsResponseBody struct {
	// The total number of items across all lists.
	//
	// example:
	//
	// 5000
	ItemsUsage *int64 `json:"ItemsUsage,omitempty" xml:"ItemsUsage,omitempty"`
	// A paginated array of lists.
	Lists []*ListListsResponseBodyLists `json:"Lists,omitempty" xml:"Lists,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of filtered records.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of lists used by the account.
	//
	// example:
	//
	// 10
	Usage *int64 `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListListsResponseBody) GoString() string {
	return s.String()
}

func (s *ListListsResponseBody) GetItemsUsage() *int64 {
	return s.ItemsUsage
}

func (s *ListListsResponseBody) GetLists() []*ListListsResponseBodyLists {
	return s.Lists
}

func (s *ListListsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListListsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListListsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListListsResponseBody) GetUsage() *int64 {
	return s.Usage
}

func (s *ListListsResponseBody) SetItemsUsage(v int64) *ListListsResponseBody {
	s.ItemsUsage = &v
	return s
}

func (s *ListListsResponseBody) SetLists(v []*ListListsResponseBodyLists) *ListListsResponseBody {
	s.Lists = v
	return s
}

func (s *ListListsResponseBody) SetPageNumber(v int32) *ListListsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListListsResponseBody) SetPageSize(v int32) *ListListsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListListsResponseBody) SetRequestId(v string) *ListListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListListsResponseBody) SetTotalCount(v int32) *ListListsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListListsResponseBody) SetUsage(v int64) *ListListsResponseBody {
	s.Usage = &v
	return s
}

func (s *ListListsResponseBody) Validate() error {
	if s.Lists != nil {
		for _, item := range s.Lists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListListsResponseBodyLists struct {
	// The description of the list.
	//
	// example:
	//
	// a custom list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom list. To get this ID, call the [ListLists](https://help.aliyun.com/document_detail/2850217.html) operation.
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The kind of the list.
	//
	// example:
	//
	// ip
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The number of items in the list.
	//
	// example:
	//
	// 100
	Length *int64 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The name of the list.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the list was last updated.
	//
	// example:
	//
	// 2024-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListListsResponseBodyLists) String() string {
	return dara.Prettify(s)
}

func (s ListListsResponseBodyLists) GoString() string {
	return s.String()
}

func (s *ListListsResponseBodyLists) GetDescription() *string {
	return s.Description
}

func (s *ListListsResponseBodyLists) GetId() *int64 {
	return s.Id
}

func (s *ListListsResponseBodyLists) GetKind() *string {
	return s.Kind
}

func (s *ListListsResponseBodyLists) GetLength() *int64 {
	return s.Length
}

func (s *ListListsResponseBodyLists) GetName() *string {
	return s.Name
}

func (s *ListListsResponseBodyLists) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListListsResponseBodyLists) SetDescription(v string) *ListListsResponseBodyLists {
	s.Description = &v
	return s
}

func (s *ListListsResponseBodyLists) SetId(v int64) *ListListsResponseBodyLists {
	s.Id = &v
	return s
}

func (s *ListListsResponseBodyLists) SetKind(v string) *ListListsResponseBodyLists {
	s.Kind = &v
	return s
}

func (s *ListListsResponseBodyLists) SetLength(v int64) *ListListsResponseBodyLists {
	s.Length = &v
	return s
}

func (s *ListListsResponseBodyLists) SetName(v string) *ListListsResponseBodyLists {
	s.Name = &v
	return s
}

func (s *ListListsResponseBodyLists) SetUpdateTime(v string) *ListListsResponseBodyLists {
	s.UpdateTime = &v
	return s
}

func (s *ListListsResponseBodyLists) Validate() error {
	return dara.Validate(s)
}

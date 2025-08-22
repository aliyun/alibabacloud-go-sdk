// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDcdnKvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v []*ListDcdnKvResponseBodyKeys) *ListDcdnKvResponseBody
	GetKeys() []*ListDcdnKvResponseBodyKeys
	SetPageNumber(v int32) *ListDcdnKvResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDcdnKvResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDcdnKvResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDcdnKvResponseBody
	GetTotalCount() *int32
}

type ListDcdnKvResponseBody struct {
	// The keys obtained in this traversal.
	Keys []*ListDcdnKvResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The total number of pages returned.
	//
	// example:
	//
	// 100
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D61E4801-EAFF-4A63-AAE1-FBF6CE1CFD1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1024
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDcdnKvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDcdnKvResponseBody) GoString() string {
	return s.String()
}

func (s *ListDcdnKvResponseBody) GetKeys() []*ListDcdnKvResponseBodyKeys {
	return s.Keys
}

func (s *ListDcdnKvResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDcdnKvResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDcdnKvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDcdnKvResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDcdnKvResponseBody) SetKeys(v []*ListDcdnKvResponseBodyKeys) *ListDcdnKvResponseBody {
	s.Keys = v
	return s
}

func (s *ListDcdnKvResponseBody) SetPageNumber(v int32) *ListDcdnKvResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDcdnKvResponseBody) SetPageSize(v int32) *ListDcdnKvResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDcdnKvResponseBody) SetRequestId(v string) *ListDcdnKvResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDcdnKvResponseBody) SetTotalCount(v int32) *ListDcdnKvResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDcdnKvResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDcdnKvResponseBodyKeys struct {
	// The value of the key obtained in this traversal.
	//
	// example:
	//
	// Key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the key was updated.
	//
	// example:
	//
	// 2021-12-13T07:46:03Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDcdnKvResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s ListDcdnKvResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListDcdnKvResponseBodyKeys) GetName() *string {
	return s.Name
}

func (s *ListDcdnKvResponseBodyKeys) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDcdnKvResponseBodyKeys) SetName(v string) *ListDcdnKvResponseBodyKeys {
	s.Name = &v
	return s
}

func (s *ListDcdnKvResponseBodyKeys) SetUpdateTime(v string) *ListDcdnKvResponseBodyKeys {
	s.UpdateTime = &v
	return s
}

func (s *ListDcdnKvResponseBodyKeys) Validate() error {
	return dara.Validate(s)
}

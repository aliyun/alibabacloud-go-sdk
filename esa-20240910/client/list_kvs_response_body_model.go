// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKvsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v []*ListKvsResponseBodyKeys) *ListKvsResponseBody
	GetKeys() []*ListKvsResponseBodyKeys
	SetPageNumber(v int32) *ListKvsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKvsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListKvsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListKvsResponseBody
	GetTotalCount() *int32
}

type ListKvsResponseBody struct {
	// The key-value pairs.
	Keys []*ListKvsResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The total number of pages returned.
	//
	// example:
	//
	// 100
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1024
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKvsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKvsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKvsResponseBody) GetKeys() []*ListKvsResponseBodyKeys {
	return s.Keys
}

func (s *ListKvsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKvsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKvsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKvsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKvsResponseBody) SetKeys(v []*ListKvsResponseBodyKeys) *ListKvsResponseBody {
	s.Keys = v
	return s
}

func (s *ListKvsResponseBody) SetPageNumber(v int32) *ListKvsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKvsResponseBody) SetPageSize(v int32) *ListKvsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKvsResponseBody) SetRequestId(v string) *ListKvsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKvsResponseBody) SetTotalCount(v int32) *ListKvsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKvsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListKvsResponseBodyKeys struct {
	// The key name.
	//
	// example:
	//
	// Key1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the key was last updated.
	//
	// example:
	//
	// 2021-12-13T07:46:03Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListKvsResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s ListKvsResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListKvsResponseBodyKeys) GetName() *string {
	return s.Name
}

func (s *ListKvsResponseBodyKeys) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListKvsResponseBodyKeys) SetName(v string) *ListKvsResponseBodyKeys {
	s.Name = &v
	return s
}

func (s *ListKvsResponseBodyKeys) SetUpdateTime(v string) *ListKvsResponseBodyKeys {
	s.UpdateTime = &v
	return s
}

func (s *ListKvsResponseBodyKeys) Validate() error {
	return dara.Validate(s)
}

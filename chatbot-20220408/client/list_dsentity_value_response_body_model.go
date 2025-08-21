// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDSEntityValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntityValues(v []*ListDSEntityValueResponseBodyEntityValues) *ListDSEntityValueResponseBody
	GetEntityValues() []*ListDSEntityValueResponseBodyEntityValues
	SetPageNumber(v int32) *ListDSEntityValueResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDSEntityValueResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDSEntityValueResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDSEntityValueResponseBody
	GetTotalCount() *int32
}

type ListDSEntityValueResponseBody struct {
	EntityValues []*ListDSEntityValueResponseBodyEntityValues `json:"EntityValues,omitempty" xml:"EntityValues,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// n3fg34gbfj8adf2gj923
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDSEntityValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDSEntityValueResponseBody) GoString() string {
	return s.String()
}

func (s *ListDSEntityValueResponseBody) GetEntityValues() []*ListDSEntityValueResponseBodyEntityValues {
	return s.EntityValues
}

func (s *ListDSEntityValueResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDSEntityValueResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDSEntityValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDSEntityValueResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDSEntityValueResponseBody) SetEntityValues(v []*ListDSEntityValueResponseBodyEntityValues) *ListDSEntityValueResponseBody {
	s.EntityValues = v
	return s
}

func (s *ListDSEntityValueResponseBody) SetPageNumber(v int32) *ListDSEntityValueResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDSEntityValueResponseBody) SetPageSize(v int32) *ListDSEntityValueResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDSEntityValueResponseBody) SetRequestId(v string) *ListDSEntityValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDSEntityValueResponseBody) SetTotalCount(v int32) *ListDSEntityValueResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDSEntityValueResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDSEntityValueResponseBodyEntityValues struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 34313785463
	EntityId *int64 `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 3467858234534534532
	EntityValueId *int64 `json:"EntityValueId,omitempty" xml:"EntityValueId,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Synonyms   []*string `json:"Synonyms,omitempty" xml:"Synonyms,omitempty" type:"Repeated"`
}

func (s ListDSEntityValueResponseBodyEntityValues) String() string {
	return dara.Prettify(s)
}

func (s ListDSEntityValueResponseBodyEntityValues) GoString() string {
	return s.String()
}

func (s *ListDSEntityValueResponseBodyEntityValues) GetContent() *string {
	return s.Content
}

func (s *ListDSEntityValueResponseBodyEntityValues) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDSEntityValueResponseBodyEntityValues) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ListDSEntityValueResponseBodyEntityValues) GetEntityValueId() *int64 {
	return s.EntityValueId
}

func (s *ListDSEntityValueResponseBodyEntityValues) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListDSEntityValueResponseBodyEntityValues) GetSynonyms() []*string {
	return s.Synonyms
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetContent(v string) *ListDSEntityValueResponseBodyEntityValues {
	s.Content = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetCreateTime(v string) *ListDSEntityValueResponseBodyEntityValues {
	s.CreateTime = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetEntityId(v int64) *ListDSEntityValueResponseBodyEntityValues {
	s.EntityId = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetEntityValueId(v int64) *ListDSEntityValueResponseBodyEntityValues {
	s.EntityValueId = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetModifyTime(v string) *ListDSEntityValueResponseBodyEntityValues {
	s.ModifyTime = &v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) SetSynonyms(v []*string) *ListDSEntityValueResponseBodyEntityValues {
	s.Synonyms = v
	return s
}

func (s *ListDSEntityValueResponseBodyEntityValues) Validate() error {
	return dara.Validate(s)
}

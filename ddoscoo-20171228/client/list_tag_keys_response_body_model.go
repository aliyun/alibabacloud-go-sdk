// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListTagKeysResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListTagKeysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTagKeysResponseBody
	GetRequestId() *string
	SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody
	GetTagKeys() []*ListTagKeysResponseBodyTagKeys
	SetTotalCount(v int32) *ListTagKeysResponseBody
	GetTotalCount() *int32
}

type ListTagKeysResponseBody struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 97935DF1-0289-4AA2-9DD1-72377838B16B
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys   []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTagKeysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagKeysResponseBody) GetTagKeys() []*ListTagKeysResponseBodyTagKeys {
	return s.TagKeys
}

func (s *ListTagKeysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTagKeysResponseBody) SetCurrentPage(v int32) *ListTagKeysResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTagKeysResponseBodyTagKeys struct {
	// example:
	//
	// 1
	TagCount *int32 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// example:
	//
	// a
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) GetTagCount() *int32 {
	return s.TagCount
}

func (s *ListTagKeysResponseBodyTagKeys) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagCount(v int32) *ListTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) Validate() error {
	return dara.Validate(s)
}

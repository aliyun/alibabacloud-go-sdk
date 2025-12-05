// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTagKeysResponseBody
	GetPageNumber() *int32
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
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 50177258-E817-4D2F-A5C6-3FD7BC4806E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of tags.
	TagKeys []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The total number of tags returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
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

func (s *ListTagKeysResponseBody) SetPageNumber(v int32) *ListTagKeysResponseBody {
	s.PageNumber = &v
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
	if s.TagKeys != nil {
		for _, item := range s.TagKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagKeysResponseBodyTagKeys struct {
	// The total number of tag keys.
	//
	// example:
	//
	// 2
	TagCount *int32 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The name of the tag key.
	//
	// example:
	//
	// key2
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

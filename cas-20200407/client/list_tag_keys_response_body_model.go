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
	SetMaxResults(v int32) *ListTagKeysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagKeysResponseBody
	GetNextToken() *string
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
	// The page number of the current page displayed in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries returned in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. An empty value of NextToken indicates that there is no next page.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The maximum number of entries per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tag keys.
	TagKeys []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The total number of entries in the list.
	//
	// example:
	//
	// 10
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

func (s *ListTagKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
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

func (s *ListTagKeysResponseBody) SetMaxResults(v int32) *ListTagKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
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
	// 1
	TagCount *int32 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The tag key.
	//
	// example:
	//
	// ac-cus-tag-3
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

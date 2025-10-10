// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTagKeysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagKeysResponseBody
	GetRequestId() *string
	SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody
	GetTagKeys() []*ListTagKeysResponseBodyTagKeys
	SetTotalCount(v int32) *ListTagKeysResponseBody
	GetTotalCount() *int32
}

type ListTagKeysResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value is the token that determines the start point of the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag keys.
	TagKeys []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The total number of entries returned.
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

func (s *ListTagKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
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

func (s *ListTagKeysResponseBody) SetMaxResults(v int32) *ListTagKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
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
	// The type of the tag.
	//
	// Valid values: **Custom**, **System**, and **All**.
	//
	// Default value: **All**.
	//
	// example:
	//
	// System
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The tag that matches all filter conditions.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) GetCategory() *string {
	return s.Category
}

func (s *ListTagKeysResponseBodyTagKeys) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagKeysResponseBodyTagKeys) SetCategory(v string) *ListTagKeysResponseBodyTagKeys {
	s.Category = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) Validate() error {
	return dara.Validate(s)
}

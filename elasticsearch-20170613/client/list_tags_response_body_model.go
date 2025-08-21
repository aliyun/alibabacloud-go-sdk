// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTagsResponseBody
	GetRequestId() *string
	SetResult(v []*ListTagsResponseBodyResult) *ListTagsResponseBody
	GetResult() []*ListTagsResponseBodyResult
}

type ListTagsResponseBody struct {
	// The key of the tag.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListTagsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagsResponseBody) GetResult() []*ListTagsResponseBodyResult {
	return s.Result
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetResult(v []*ListTagsResponseBodyResult) *ListTagsResponseBody {
	s.Result = v
	return s
}

func (s *ListTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTagsResponseBodyResult struct {
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// dev
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyResult) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagsResponseBodyResult) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagsResponseBodyResult) SetTagKey(v string) *ListTagsResponseBodyResult {
	s.TagKey = &v
	return s
}

func (s *ListTagsResponseBodyResult) SetTagValue(v string) *ListTagsResponseBodyResult {
	s.TagValue = &v
	return s
}

func (s *ListTagsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

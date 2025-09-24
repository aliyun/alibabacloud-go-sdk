// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeSources(v []*CodeSourceItem) *ListCodeSourcesResponseBody
	GetCodeSources() []*CodeSourceItem
	SetRequestId(v string) *ListCodeSourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCodeSourcesResponseBody
	GetTotalCount() *int64
}

type ListCodeSourcesResponseBody struct {
	// The code sources.
	CodeSources []*CodeSourceItem `json:"CodeSources,omitempty" xml:"CodeSources,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of code sources that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCodeSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCodeSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCodeSourcesResponseBody) GetCodeSources() []*CodeSourceItem {
	return s.CodeSources
}

func (s *ListCodeSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCodeSourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCodeSourcesResponseBody) SetCodeSources(v []*CodeSourceItem) *ListCodeSourcesResponseBody {
	s.CodeSources = v
	return s
}

func (s *ListCodeSourcesResponseBody) SetRequestId(v string) *ListCodeSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCodeSourcesResponseBody) SetTotalCount(v int64) *ListCodeSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCodeSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogProjects(v []*string) *ListLogProjectsResponseBody
	GetLogProjects() []*string
	SetMaxResults(v int32) *ListLogProjectsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListLogProjectsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListLogProjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLogProjectsResponseBody
	GetTotalCount() *int32
}

type ListLogProjectsResponseBody struct {
	LogProjects []*string `json:"LogProjects,omitempty" xml:"LogProjects,omitempty" type:"Repeated"`
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLogProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogProjectsResponseBody) GetLogProjects() []*string {
	return s.LogProjects
}

func (s *ListLogProjectsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLogProjectsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLogProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogProjectsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLogProjectsResponseBody) SetLogProjects(v []*string) *ListLogProjectsResponseBody {
	s.LogProjects = v
	return s
}

func (s *ListLogProjectsResponseBody) SetMaxResults(v int32) *ListLogProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListLogProjectsResponseBody) SetNextToken(v string) *ListLogProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLogProjectsResponseBody) SetRequestId(v string) *ListLogProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogProjectsResponseBody) SetTotalCount(v int32) *ListLogProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLogProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

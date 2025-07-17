// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMemoriesResponseBody
	GetMaxResults() *int32
	SetMemories(v []*ListMemoriesResponseBodyMemories) *ListMemoriesResponseBody
	GetMemories() []*ListMemoriesResponseBodyMemories
	SetNextToken(v string) *ListMemoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMemoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMemoriesResponseBody
	GetTotalCount() *int32
	SetWorkspaceId(v string) *ListMemoriesResponseBody
	GetWorkspaceId() *string
}

type ListMemoriesResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32                              `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Memories   []*ListMemoriesResponseBodyMemories `json:"memories,omitempty" xml:"memories,omitempty" type:"Repeated"`
	// example:
	//
	// dc270401186b433f975d7e1faaa34e0e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 105
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// llm-us9hjmt32nysdm5v
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListMemoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoriesResponseBody) GetMemories() []*ListMemoriesResponseBodyMemories {
	return s.Memories
}

func (s *ListMemoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMemoriesResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMemoriesResponseBody) SetMaxResults(v int32) *ListMemoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemoriesResponseBody) SetMemories(v []*ListMemoriesResponseBodyMemories) *ListMemoriesResponseBody {
	s.Memories = v
	return s
}

func (s *ListMemoriesResponseBody) SetNextToken(v string) *ListMemoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemoriesResponseBody) SetRequestId(v string) *ListMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoriesResponseBody) SetTotalCount(v int32) *ListMemoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMemoriesResponseBody) SetWorkspaceId(v string) *ListMemoriesResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *ListMemoriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMemoriesResponseBodyMemories struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 3fc531f4519444beaafffa4538f60667
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
}

func (s ListMemoriesResponseBodyMemories) String() string {
	return dara.Prettify(s)
}

func (s ListMemoriesResponseBodyMemories) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponseBodyMemories) GetDescription() *string {
	return s.Description
}

func (s *ListMemoriesResponseBodyMemories) GetMemoryId() *string {
	return s.MemoryId
}

func (s *ListMemoriesResponseBodyMemories) SetDescription(v string) *ListMemoriesResponseBodyMemories {
	s.Description = &v
	return s
}

func (s *ListMemoriesResponseBodyMemories) SetMemoryId(v string) *ListMemoriesResponseBodyMemories {
	s.MemoryId = &v
	return s
}

func (s *ListMemoriesResponseBodyMemories) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMemoryNodesResponseBody
	GetMaxResults() *int32
	SetMemoryNodes(v []*ListMemoryNodesResponseBodyMemoryNodes) *ListMemoryNodesResponseBody
	GetMemoryNodes() []*ListMemoryNodesResponseBodyMemoryNodes
	SetNextToken(v string) *ListMemoryNodesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMemoryNodesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMemoryNodesResponseBody
	GetTotalCount() *int32
}

type ListMemoryNodesResponseBody struct {
	// The maximum number of results returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The array of memory nodes.
	MemoryNodes []*ListMemoryNodesResponseBodyMemoryNodes `json:"memoryNodes,omitempty" xml:"memoryNodes,omitempty" type:"Repeated"`
	// The token used for token-based pagination.
	//
	// example:
	//
	// dc270401186b433f975d7e1faaa34e0e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of memory nodes.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMemoryNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoryNodesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoryNodesResponseBody) GetMemoryNodes() []*ListMemoryNodesResponseBodyMemoryNodes {
	return s.MemoryNodes
}

func (s *ListMemoryNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoryNodesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMemoryNodesResponseBody) SetMaxResults(v int32) *ListMemoryNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryNodesResponseBody) SetMemoryNodes(v []*ListMemoryNodesResponseBodyMemoryNodes) *ListMemoryNodesResponseBody {
	s.MemoryNodes = v
	return s
}

func (s *ListMemoryNodesResponseBody) SetNextToken(v string) *ListMemoryNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemoryNodesResponseBody) SetRequestId(v string) *ListMemoryNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoryNodesResponseBody) SetTotalCount(v int32) *ListMemoryNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMemoryNodesResponseBody) Validate() error {
	if s.MemoryNodes != nil {
		for _, item := range s.MemoryNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMemoryNodesResponseBodyMemoryNodes struct {
	// The content of the memory node.
	//
	// example:
	//
	// 用户喜欢吃西红柿炒鸡蛋
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The memory node ID.
	//
	// example:
	//
	// 68de06c95368463a8be4a84efc872cc5
	MemoryNodeId *string `json:"memoryNodeId,omitempty" xml:"memoryNodeId,omitempty"`
}

func (s ListMemoryNodesResponseBodyMemoryNodes) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryNodesResponseBodyMemoryNodes) GoString() string {
	return s.String()
}

func (s *ListMemoryNodesResponseBodyMemoryNodes) GetContent() *string {
	return s.Content
}

func (s *ListMemoryNodesResponseBodyMemoryNodes) GetMemoryNodeId() *string {
	return s.MemoryNodeId
}

func (s *ListMemoryNodesResponseBodyMemoryNodes) SetContent(v string) *ListMemoryNodesResponseBodyMemoryNodes {
	s.Content = &v
	return s
}

func (s *ListMemoryNodesResponseBodyMemoryNodes) SetMemoryNodeId(v string) *ListMemoryNodesResponseBodyMemoryNodes {
	s.MemoryNodeId = &v
	return s
}

func (s *ListMemoryNodesResponseBodyMemoryNodes) Validate() error {
	return dara.Validate(s)
}

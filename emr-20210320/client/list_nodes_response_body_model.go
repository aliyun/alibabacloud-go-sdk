// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNodesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodesResponseBody
	GetNextToken() *string
	SetNodes(v []*Node) *ListNodesResponseBody
	GetNodes() []*Node
	SetRequestId(v string) *ListNodesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNodesResponseBody
	GetTotalCount() *int32
}

type ListNodesResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position at which the next read starts. If null is returned, the data has been read.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The node list.
	Nodes []*Node `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records in this request.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodesResponseBody) GetNodes() []*Node {
	return s.Nodes
}

func (s *ListNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNodesResponseBody) SetMaxResults(v int32) *ListNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNodesResponseBody) SetNextToken(v string) *ListNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodesResponseBody) SetNodes(v []*Node) *ListNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetTotalCount(v int32) *ListNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNodesResponseBody) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

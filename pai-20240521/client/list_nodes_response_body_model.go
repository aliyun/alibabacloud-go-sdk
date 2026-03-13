// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodes(v []*Node) *ListNodesResponseBody
	GetNodes() []*Node
	SetRequestId(v string) *ListNodesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNodesResponseBody
	GetTotalCount() *int32
}

type ListNodesResponseBody struct {
	Nodes      []*Node `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
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

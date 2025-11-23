// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowEdgesByConditionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEdges(v *ListTaskFlowEdgesByConditionResponseBodyEdges) *ListTaskFlowEdgesByConditionResponseBody
	GetEdges() *ListTaskFlowEdgesByConditionResponseBodyEdges
	SetErrorCode(v string) *ListTaskFlowEdgesByConditionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTaskFlowEdgesByConditionResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTaskFlowEdgesByConditionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskFlowEdgesByConditionResponseBody
	GetSuccess() *bool
}

type ListTaskFlowEdgesByConditionResponseBody struct {
	// The list of task flow edges.
	Edges *ListTaskFlowEdgesByConditionResponseBodyEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// D86249CD-422F-5ACF-85BA-9187C986AE0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTaskFlowEdgesByConditionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowEdgesByConditionResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskFlowEdgesByConditionResponseBody) GetEdges() *ListTaskFlowEdgesByConditionResponseBodyEdges {
	return s.Edges
}

func (s *ListTaskFlowEdgesByConditionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTaskFlowEdgesByConditionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTaskFlowEdgesByConditionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskFlowEdgesByConditionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskFlowEdgesByConditionResponseBody) SetEdges(v *ListTaskFlowEdgesByConditionResponseBodyEdges) *ListTaskFlowEdgesByConditionResponseBody {
	s.Edges = v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponseBody) SetErrorCode(v string) *ListTaskFlowEdgesByConditionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponseBody) SetErrorMessage(v string) *ListTaskFlowEdgesByConditionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponseBody) SetRequestId(v string) *ListTaskFlowEdgesByConditionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponseBody) SetSuccess(v bool) *ListTaskFlowEdgesByConditionResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponseBody) Validate() error {
	if s.Edges != nil {
		if err := s.Edges.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskFlowEdgesByConditionResponseBodyEdges struct {
	Edge []*ListTaskFlowEdgesByConditionResponseBodyEdgesEdge `json:"Edge,omitempty" xml:"Edge,omitempty" type:"Repeated"`
}

func (s ListTaskFlowEdgesByConditionResponseBodyEdges) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowEdgesByConditionResponseBodyEdges) GoString() string {
	return s.String()
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdges) GetEdge() []*ListTaskFlowEdgesByConditionResponseBodyEdgesEdge {
	return s.Edge
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdges) SetEdge(v []*ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) *ListTaskFlowEdgesByConditionResponseBodyEdges {
	s.Edge = v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdges) Validate() error {
	if s.Edge != nil {
		for _, item := range s.Edge {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskFlowEdgesByConditionResponseBodyEdgesEdge struct {
	// The ID of the task flow edge.
	//
	// example:
	//
	// 24***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the end node on the edge.
	//
	// example:
	//
	// 44***
	NodeEnd *int64 `json:"NodeEnd,omitempty" xml:"NodeEnd,omitempty"`
	// The ID of the start node on the edge.
	//
	// example:
	//
	// 44***
	NodeFrom *int64 `json:"NodeFrom,omitempty" xml:"NodeFrom,omitempty"`
}

func (s ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) GoString() string {
	return s.String()
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) GetId() *int64 {
	return s.Id
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) GetNodeEnd() *int64 {
	return s.NodeEnd
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) GetNodeFrom() *int64 {
	return s.NodeFrom
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) SetId(v int64) *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge {
	s.Id = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) SetNodeEnd(v int64) *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge {
	s.NodeEnd = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) SetNodeFrom(v int64) *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge {
	s.NodeFrom = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionResponseBodyEdgesEdge) Validate() error {
	return dara.Validate(s)
}

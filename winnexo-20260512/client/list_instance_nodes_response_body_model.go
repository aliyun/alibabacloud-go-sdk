// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstanceNodesResponseBody
	GetCode() *string
	SetInstanceId(v string) *ListInstanceNodesResponseBody
	GetInstanceId() *string
	SetMessage(v string) *ListInstanceNodesResponseBody
	GetMessage() *string
	SetNodes(v []*ListInstanceNodesResponseBodyNodes) *ListInstanceNodesResponseBody
	GetNodes() []*ListInstanceNodesResponseBodyNodes
	SetRequestId(v string) *ListInstanceNodesResponseBody
	GetRequestId() *string
	SetStateVersion(v int64) *ListInstanceNodesResponseBody
	GetStateVersion() *int64
}

type ListInstanceNodesResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The physical instance ID.
	//
	// example:
	//
	// plan_example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The visible nodes arranged in topological order. The number of nodes is determined by the actual dynamic graph.
	//
	// example:
	//
	// []
	Nodes []*ListInstanceNodesResponseBodyNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The state version.
	//
	// example:
	//
	// 2
	StateVersion *int64 `json:"stateVersion,omitempty" xml:"stateVersion,omitempty"`
}

func (s ListInstanceNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceNodesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstanceNodesResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceNodesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstanceNodesResponseBody) GetNodes() []*ListInstanceNodesResponseBodyNodes {
	return s.Nodes
}

func (s *ListInstanceNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceNodesResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *ListInstanceNodesResponseBody) SetCode(v string) *ListInstanceNodesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceNodesResponseBody) SetInstanceId(v string) *ListInstanceNodesResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceNodesResponseBody) SetMessage(v string) *ListInstanceNodesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceNodesResponseBody) SetNodes(v []*ListInstanceNodesResponseBodyNodes) *ListInstanceNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListInstanceNodesResponseBody) SetRequestId(v string) *ListInstanceNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceNodesResponseBody) SetStateVersion(v int64) *ListInstanceNodesResponseBody {
	s.StateVersion = &v
	return s
}

func (s *ListInstanceNodesResponseBody) Validate() error {
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

type ListInstanceNodesResponseBodyNodes struct {
	// The conclusion of the node result. This field is empty if there is no conclusion.
	//
	// example:
	//
	// Analysis completed
	Conclusion *string `json:"conclusion,omitempty" xml:"conclusion,omitempty"`
	// The total number of seconds from the first run to the desired state. This field is empty if the time is missing.
	//
	// example:
	//
	// 123
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The time when the node was first run. This field is empty for historical or unrun nodes.
	//
	// example:
	//
	// 2026-10-14T08:00:00Z
	EnteredAt *string `json:"enteredAt,omitempty" xml:"enteredAt,omitempty"`
	// The time when the node actually entered the desired state. This field is empty if the node has not finished.
	//
	// example:
	//
	// 2026-10-14T08:02:03Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// The display name of the node assignee. This field is empty if the assignee is unknown.
	//
	// example:
	//
	// Assistant Customer Success Manager (CSM)
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// The lane position of the node in the workflow graph.
	//
	// example:
	//
	// 0
	Lane *int64 `json:"lane,omitempty" xml:"lane,omitempty"`
	// The key of the node (nodeKey). This is not the session ID.
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The display name of the node.
	//
	// example:
	//
	// Query signed contract amount
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// The display sequence number of the node in the returned topology list, starting from 1.
	//
	// example:
	//
	// 3
	NodeOrder *int64 `json:"nodeOrder,omitempty" xml:"nodeOrder,omitempty"`
	// The actual node type, such as AGENT or HUMAN.
	//
	// example:
	//
	// AGENT
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// The level of the node in the workflow graph, starting from 0.
	//
	// example:
	//
	// 2
	Rank *int64 `json:"rank,omitempty" xml:"rank,omitempty"`
	// The original running status of the node, such as READY, RUNNING, WAITING, SUCCEEDED, or FAILED.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for waiting, such as HUMAN_INPUT, HUMAN_STEP, REVIEW, MANUAL_PAUSE, or UNBLOCK. This field is empty if the node is not waiting.
	//
	// example:
	//
	// REVIEW
	WaitType *string `json:"waitType,omitempty" xml:"waitType,omitempty"`
}

func (s ListInstanceNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListInstanceNodesResponseBodyNodes) GetConclusion() *string {
	return s.Conclusion
}

func (s *ListInstanceNodesResponseBodyNodes) GetDuration() *int64 {
	return s.Duration
}

func (s *ListInstanceNodesResponseBodyNodes) GetEnteredAt() *string {
	return s.EnteredAt
}

func (s *ListInstanceNodesResponseBodyNodes) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *ListInstanceNodesResponseBodyNodes) GetHandler() *string {
	return s.Handler
}

func (s *ListInstanceNodesResponseBodyNodes) GetLane() *int64 {
	return s.Lane
}

func (s *ListInstanceNodesResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ListInstanceNodesResponseBodyNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *ListInstanceNodesResponseBodyNodes) GetNodeOrder() *int64 {
	return s.NodeOrder
}

func (s *ListInstanceNodesResponseBodyNodes) GetNodeType() *string {
	return s.NodeType
}

func (s *ListInstanceNodesResponseBodyNodes) GetRank() *int64 {
	return s.Rank
}

func (s *ListInstanceNodesResponseBodyNodes) GetStatus() *string {
	return s.Status
}

func (s *ListInstanceNodesResponseBodyNodes) GetWaitType() *string {
	return s.WaitType
}

func (s *ListInstanceNodesResponseBodyNodes) SetConclusion(v string) *ListInstanceNodesResponseBodyNodes {
	s.Conclusion = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetDuration(v int64) *ListInstanceNodesResponseBodyNodes {
	s.Duration = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetEnteredAt(v string) *ListInstanceNodesResponseBodyNodes {
	s.EnteredAt = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetFinishedAt(v string) *ListInstanceNodesResponseBodyNodes {
	s.FinishedAt = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetHandler(v string) *ListInstanceNodesResponseBodyNodes {
	s.Handler = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetLane(v int64) *ListInstanceNodesResponseBodyNodes {
	s.Lane = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetNodeId(v string) *ListInstanceNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetNodeName(v string) *ListInstanceNodesResponseBodyNodes {
	s.NodeName = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetNodeOrder(v int64) *ListInstanceNodesResponseBodyNodes {
	s.NodeOrder = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetNodeType(v string) *ListInstanceNodesResponseBodyNodes {
	s.NodeType = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetRank(v int64) *ListInstanceNodesResponseBodyNodes {
	s.Rank = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetStatus(v string) *ListInstanceNodesResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) SetWaitType(v string) *ListInstanceNodesResponseBodyNodes {
	s.WaitType = &v
	return s
}

func (s *ListInstanceNodesResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}

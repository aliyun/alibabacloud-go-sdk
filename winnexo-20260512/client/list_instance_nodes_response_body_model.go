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
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 物理实例ID
	//
	// example:
	//
	// plan_example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 按拓扑顺序排列的可见节点，数量由实际动态图决定
	//
	// example:
	//
	// []
	Nodes []*ListInstanceNodesResponseBodyNodes `json:"nodes,omitempty" xml:"nodes,omitempty" type:"Repeated"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 状态版本
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
	// 节点结果结论，无结论时为空
	//
	// example:
	//
	// 分析完成
	Conclusion *string `json:"conclusion,omitempty" xml:"conclusion,omitempty"`
	// 从首次运行到终态的整秒数，缺时间时为空
	//
	// example:
	//
	// 123
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 节点首次运行时间，历史或未运行节点为空
	//
	// example:
	//
	// 2026-10-14T08:00:00Z
	EnteredAt *string `json:"enteredAt,omitempty" xml:"enteredAt,omitempty"`
	// 节点实际进入终态的时间，未结束时为空
	//
	// example:
	//
	// 2026-10-14T08:02:03Z
	FinishedAt *string `json:"finishedAt,omitempty" xml:"finishedAt,omitempty"`
	// 节点负责人展示名称，未知时为空
	//
	// example:
	//
	// 客户成功经理助理
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 节点在工作流图中的泳道位置
	//
	// example:
	//
	// 0
	Lane *int64 `json:"lane,omitempty" xml:"lane,omitempty"`
	// 所属节点key（nodeKey），不是会话ID
	//
	// example:
	//
	// query.amount
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 节点展示名称
	//
	// example:
	//
	// 查询签单金额
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// 节点在返回拓扑列表中的展示序号，从1开始
	//
	// example:
	//
	// 3
	NodeOrder *int64 `json:"nodeOrder,omitempty" xml:"nodeOrder,omitempty"`
	// 实际节点类型，例如AGENT或HUMAN
	//
	// example:
	//
	// AGENT
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// 节点在工作流图中的层级，从0开始
	//
	// example:
	//
	// 2
	Rank *int64 `json:"rank,omitempty" xml:"rank,omitempty"`
	// 原始节点运行状态，例如READY、RUNNING、WAITING、SUCCEEDED或FAILED
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 等待原因：HUMAN_INPUT、HUMAN_STEP、REVIEW、MANUAL_PAUSE或UNBLOCK；未等待时为空
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkflowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v string) *CreateWorkflowInstanceResponseBody
	GetBusinessId() *string
	SetCode(v string) *CreateWorkflowInstanceResponseBody
	GetCode() *string
	SetCollaborationGroupId(v string) *CreateWorkflowInstanceResponseBody
	GetCollaborationGroupId() *string
	SetCreatedAt(v string) *CreateWorkflowInstanceResponseBody
	GetCreatedAt() *string
	SetCurrentNode(v *CreateWorkflowInstanceResponseBodyCurrentNode) *CreateWorkflowInstanceResponseBody
	GetCurrentNode() *CreateWorkflowInstanceResponseBodyCurrentNode
	SetCurrentNodes(v []*CreateWorkflowInstanceResponseBodyCurrentNodes) *CreateWorkflowInstanceResponseBody
	GetCurrentNodes() []*CreateWorkflowInstanceResponseBodyCurrentNodes
	SetEndedAt(v string) *CreateWorkflowInstanceResponseBody
	GetEndedAt() *string
	SetInitiator(v string) *CreateWorkflowInstanceResponseBody
	GetInitiator() *string
	SetInstanceId(v string) *CreateWorkflowInstanceResponseBody
	GetInstanceId() *string
	SetMessage(v string) *CreateWorkflowInstanceResponseBody
	GetMessage() *string
	SetPauseScope(v string) *CreateWorkflowInstanceResponseBody
	GetPauseScope() *string
	SetPlanId(v string) *CreateWorkflowInstanceResponseBody
	GetPlanId() *string
	SetProgress(v *CreateWorkflowInstanceResponseBodyProgress) *CreateWorkflowInstanceResponseBody
	GetProgress() *CreateWorkflowInstanceResponseBodyProgress
	SetRequestId(v string) *CreateWorkflowInstanceResponseBody
	GetRequestId() *string
	SetReviewStatus(v string) *CreateWorkflowInstanceResponseBody
	GetReviewStatus() *string
	SetRevision(v int64) *CreateWorkflowInstanceResponseBody
	GetRevision() *int64
	SetRunId(v string) *CreateWorkflowInstanceResponseBody
	GetRunId() *string
	SetRuntimeStatus(v string) *CreateWorkflowInstanceResponseBody
	GetRuntimeStatus() *string
	SetStartedAt(v string) *CreateWorkflowInstanceResponseBody
	GetStartedAt() *string
	SetStateVersion(v int64) *CreateWorkflowInstanceResponseBody
	GetStateVersion() *int64
	SetStatus(v string) *CreateWorkflowInstanceResponseBody
	GetStatus() *string
	SetTemplateId(v string) *CreateWorkflowInstanceResponseBody
	GetTemplateId() *string
	SetTitle(v string) *CreateWorkflowInstanceResponseBody
	GetTitle() *string
	SetWorkItemId(v string) *CreateWorkflowInstanceResponseBody
	GetWorkItemId() *string
}

type CreateWorkflowInstanceResponseBody struct {
	// 空间内唯一外部需求编号
	//
	// example:
	//
	// REQ-1001
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// 业务状态码，来自统一 POP 响应包装
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 协作空间 ID
	//
	// example:
	//
	// cg_20261016_abc123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 逻辑实例创建时间，ISO8601
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 首个活动节点，终态为 null
	CurrentNode *CreateWorkflowInstanceResponseBodyCurrentNode `json:"currentNode,omitempty" xml:"currentNode,omitempty" type:"Struct"`
	// 全部活动节点，支持并行 DAG
	CurrentNodes []*CreateWorkflowInstanceResponseBodyCurrentNodes `json:"currentNodes,omitempty" xml:"currentNodes,omitempty" type:"Repeated"`
	// 当前计划终态时间，未结束为 null
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	EndedAt *string `json:"endedAt,omitempty" xml:"endedAt,omitempty"`
	// 外部发起人标签，不参与鉴权
	//
	// example:
	//
	// 外部发起人张三
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// 稳定逻辑实例 ID，工作项重启不改变
	//
	// example:
	//
	// wfi_20261016_abc123
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 用户可见错误描述，来自统一 POP 响应包装
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 原始计划暂停范围
	//
	// example:
	//
	// ALL
	PauseScope *string `json:"pauseScope,omitempty" xml:"pauseScope,omitempty"`
	// 当前物理计划修订 ID
	//
	// example:
	//
	// plan_20261016_abc123
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// 可见图进度
	Progress *CreateWorkflowInstanceResponseBodyProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 原始计划审核状态
	//
	// example:
	//
	// APPROVED
	ReviewStatus *string `json:"reviewStatus,omitempty" xml:"reviewStatus,omitempty"`
	// 当前物理计划修订号
	//
	// example:
	//
	// 3
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// 当前运行 ID，重启后会改变
	//
	// example:
	//
	// run_20261016_abc123
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// 原始 BUILDING/RUNNING/SUCCEEDED/FAILED/CANCELLED
	//
	// example:
	//
	// RUNNING
	RuntimeStatus *string `json:"runtimeStatus,omitempty" xml:"runtimeStatus,omitempty"`
	// 当前计划节点首次开始时间，未执行为 null
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// 运行状态版本
	//
	// example:
	//
	// 7
	StateVersion *int64 `json:"stateVersion,omitempty" xml:"stateVersion,omitempty"`
	// Running/Completed/Failed/Terminated
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 创建时绑定的工作流模板 ID
	//
	// example:
	//
	// wft_20261016_abc123
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// 当前需求标题
	//
	// example:
	//
	// 季度经营分析报告
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// 关联工作项 ID
	//
	// example:
	//
	// wi_20261016_abc123
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s CreateWorkflowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstanceResponseBody) GetBusinessId() *string {
	return s.BusinessId
}

func (s *CreateWorkflowInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkflowInstanceResponseBody) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CreateWorkflowInstanceResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateWorkflowInstanceResponseBody) GetCurrentNode() *CreateWorkflowInstanceResponseBodyCurrentNode {
	return s.CurrentNode
}

func (s *CreateWorkflowInstanceResponseBody) GetCurrentNodes() []*CreateWorkflowInstanceResponseBodyCurrentNodes {
	return s.CurrentNodes
}

func (s *CreateWorkflowInstanceResponseBody) GetEndedAt() *string {
	return s.EndedAt
}

func (s *CreateWorkflowInstanceResponseBody) GetInitiator() *string {
	return s.Initiator
}

func (s *CreateWorkflowInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWorkflowInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkflowInstanceResponseBody) GetPauseScope() *string {
	return s.PauseScope
}

func (s *CreateWorkflowInstanceResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *CreateWorkflowInstanceResponseBody) GetProgress() *CreateWorkflowInstanceResponseBodyProgress {
	return s.Progress
}

func (s *CreateWorkflowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkflowInstanceResponseBody) GetReviewStatus() *string {
	return s.ReviewStatus
}

func (s *CreateWorkflowInstanceResponseBody) GetRevision() *int64 {
	return s.Revision
}

func (s *CreateWorkflowInstanceResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *CreateWorkflowInstanceResponseBody) GetRuntimeStatus() *string {
	return s.RuntimeStatus
}

func (s *CreateWorkflowInstanceResponseBody) GetStartedAt() *string {
	return s.StartedAt
}

func (s *CreateWorkflowInstanceResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *CreateWorkflowInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateWorkflowInstanceResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateWorkflowInstanceResponseBody) GetTitle() *string {
	return s.Title
}

func (s *CreateWorkflowInstanceResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *CreateWorkflowInstanceResponseBody) SetBusinessId(v string) *CreateWorkflowInstanceResponseBody {
	s.BusinessId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetCode(v string) *CreateWorkflowInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetCollaborationGroupId(v string) *CreateWorkflowInstanceResponseBody {
	s.CollaborationGroupId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetCreatedAt(v string) *CreateWorkflowInstanceResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetCurrentNode(v *CreateWorkflowInstanceResponseBodyCurrentNode) *CreateWorkflowInstanceResponseBody {
	s.CurrentNode = v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetCurrentNodes(v []*CreateWorkflowInstanceResponseBodyCurrentNodes) *CreateWorkflowInstanceResponseBody {
	s.CurrentNodes = v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetEndedAt(v string) *CreateWorkflowInstanceResponseBody {
	s.EndedAt = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetInitiator(v string) *CreateWorkflowInstanceResponseBody {
	s.Initiator = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetInstanceId(v string) *CreateWorkflowInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetMessage(v string) *CreateWorkflowInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetPauseScope(v string) *CreateWorkflowInstanceResponseBody {
	s.PauseScope = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetPlanId(v string) *CreateWorkflowInstanceResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetProgress(v *CreateWorkflowInstanceResponseBodyProgress) *CreateWorkflowInstanceResponseBody {
	s.Progress = v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetRequestId(v string) *CreateWorkflowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetReviewStatus(v string) *CreateWorkflowInstanceResponseBody {
	s.ReviewStatus = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetRevision(v int64) *CreateWorkflowInstanceResponseBody {
	s.Revision = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetRunId(v string) *CreateWorkflowInstanceResponseBody {
	s.RunId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetRuntimeStatus(v string) *CreateWorkflowInstanceResponseBody {
	s.RuntimeStatus = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetStartedAt(v string) *CreateWorkflowInstanceResponseBody {
	s.StartedAt = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetStateVersion(v int64) *CreateWorkflowInstanceResponseBody {
	s.StateVersion = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetStatus(v string) *CreateWorkflowInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetTemplateId(v string) *CreateWorkflowInstanceResponseBody {
	s.TemplateId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetTitle(v string) *CreateWorkflowInstanceResponseBody {
	s.Title = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) SetWorkItemId(v string) *CreateWorkflowInstanceResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBody) Validate() error {
	if s.CurrentNode != nil {
		if err := s.CurrentNode.Validate(); err != nil {
			return err
		}
	}
	if s.CurrentNodes != nil {
		for _, item := range s.CurrentNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Progress != nil {
		if err := s.Progress.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkflowInstanceResponseBodyCurrentNode struct {
	// 当前处理人展示名
	//
	// example:
	//
	// 张三
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 处理人或数字员工标识
	//
	// example:
	//
	// user_10086
	HandlerId *string `json:"handlerId,omitempty" xml:"handlerId,omitempty"`
	// 真实运行节点 nodeKey
	//
	// example:
	//
	// node_key_confirm
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 节点展示名
	//
	// example:
	//
	// 需求确认
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// 真实节点类型 AGENT/HUMAN/CONDITION
	//
	// example:
	//
	// AGENT
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// 节点运行状态
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateWorkflowInstanceResponseBodyCurrentNode) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstanceResponseBodyCurrentNode) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) GetHandler() *string {
	return s.Handler
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) GetHandlerId() *string {
	return s.HandlerId
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) GetNodeName() *string {
	return s.NodeName
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) GetStatus() *string {
	return s.Status
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) SetHandler(v string) *CreateWorkflowInstanceResponseBodyCurrentNode {
	s.Handler = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) SetHandlerId(v string) *CreateWorkflowInstanceResponseBodyCurrentNode {
	s.HandlerId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) SetNodeId(v string) *CreateWorkflowInstanceResponseBodyCurrentNode {
	s.NodeId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) SetNodeName(v string) *CreateWorkflowInstanceResponseBodyCurrentNode {
	s.NodeName = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) SetNodeType(v string) *CreateWorkflowInstanceResponseBodyCurrentNode {
	s.NodeType = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) SetStatus(v string) *CreateWorkflowInstanceResponseBodyCurrentNode {
	s.Status = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNode) Validate() error {
	return dara.Validate(s)
}

type CreateWorkflowInstanceResponseBodyCurrentNodes struct {
	// 当前处理人展示名
	//
	// example:
	//
	// 张三
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// 处理人或数字员工标识
	//
	// example:
	//
	// user_10086
	HandlerId *string `json:"handlerId,omitempty" xml:"handlerId,omitempty"`
	// 真实运行节点 nodeKey
	//
	// example:
	//
	// node_key_confirm
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 节点展示名
	//
	// example:
	//
	// 需求确认
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// 真实节点类型 AGENT/HUMAN/CONDITION
	//
	// example:
	//
	// AGENT
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// 节点运行状态
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateWorkflowInstanceResponseBodyCurrentNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstanceResponseBodyCurrentNodes) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) GetHandler() *string {
	return s.Handler
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) GetHandlerId() *string {
	return s.HandlerId
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) GetStatus() *string {
	return s.Status
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) SetHandler(v string) *CreateWorkflowInstanceResponseBodyCurrentNodes {
	s.Handler = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) SetHandlerId(v string) *CreateWorkflowInstanceResponseBodyCurrentNodes {
	s.HandlerId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) SetNodeId(v string) *CreateWorkflowInstanceResponseBodyCurrentNodes {
	s.NodeId = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) SetNodeName(v string) *CreateWorkflowInstanceResponseBodyCurrentNodes {
	s.NodeName = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) SetNodeType(v string) *CreateWorkflowInstanceResponseBodyCurrentNodes {
	s.NodeType = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) SetStatus(v string) *CreateWorkflowInstanceResponseBodyCurrentNodes {
	s.Status = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyCurrentNodes) Validate() error {
	return dara.Validate(s)
}

type CreateWorkflowInstanceResponseBodyProgress struct {
	// 已成功或跳过的可见节点数
	//
	// example:
	//
	// 3
	CompletedCount *int64 `json:"completedCount,omitempty" xml:"completedCount,omitempty"`
	// 当前图可见节点总数，非固定六节点
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s CreateWorkflowInstanceResponseBodyProgress) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkflowInstanceResponseBodyProgress) GoString() string {
	return s.String()
}

func (s *CreateWorkflowInstanceResponseBodyProgress) GetCompletedCount() *int64 {
	return s.CompletedCount
}

func (s *CreateWorkflowInstanceResponseBodyProgress) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *CreateWorkflowInstanceResponseBodyProgress) SetCompletedCount(v int64) *CreateWorkflowInstanceResponseBodyProgress {
	s.CompletedCount = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyProgress) SetTotalCount(v int64) *CreateWorkflowInstanceResponseBodyProgress {
	s.TotalCount = &v
	return s
}

func (s *CreateWorkflowInstanceResponseBodyProgress) Validate() error {
	return dara.Validate(s)
}

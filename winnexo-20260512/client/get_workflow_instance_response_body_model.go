// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v string) *GetWorkflowInstanceResponseBody
	GetBusinessId() *string
	SetCode(v string) *GetWorkflowInstanceResponseBody
	GetCode() *string
	SetCollaborationGroupId(v string) *GetWorkflowInstanceResponseBody
	GetCollaborationGroupId() *string
	SetCreatedAt(v string) *GetWorkflowInstanceResponseBody
	GetCreatedAt() *string
	SetCurrentNode(v *GetWorkflowInstanceResponseBodyCurrentNode) *GetWorkflowInstanceResponseBody
	GetCurrentNode() *GetWorkflowInstanceResponseBodyCurrentNode
	SetCurrentNodes(v []*GetWorkflowInstanceResponseBodyCurrentNodes) *GetWorkflowInstanceResponseBody
	GetCurrentNodes() []*GetWorkflowInstanceResponseBodyCurrentNodes
	SetEndedAt(v string) *GetWorkflowInstanceResponseBody
	GetEndedAt() *string
	SetInitiator(v string) *GetWorkflowInstanceResponseBody
	GetInitiator() *string
	SetInstanceId(v string) *GetWorkflowInstanceResponseBody
	GetInstanceId() *string
	SetMessage(v string) *GetWorkflowInstanceResponseBody
	GetMessage() *string
	SetPauseScope(v string) *GetWorkflowInstanceResponseBody
	GetPauseScope() *string
	SetPlanId(v string) *GetWorkflowInstanceResponseBody
	GetPlanId() *string
	SetProgress(v *GetWorkflowInstanceResponseBodyProgress) *GetWorkflowInstanceResponseBody
	GetProgress() *GetWorkflowInstanceResponseBodyProgress
	SetRequestId(v string) *GetWorkflowInstanceResponseBody
	GetRequestId() *string
	SetReviewStatus(v string) *GetWorkflowInstanceResponseBody
	GetReviewStatus() *string
	SetRevision(v int64) *GetWorkflowInstanceResponseBody
	GetRevision() *int64
	SetRunId(v string) *GetWorkflowInstanceResponseBody
	GetRunId() *string
	SetRuntimeStatus(v string) *GetWorkflowInstanceResponseBody
	GetRuntimeStatus() *string
	SetStartedAt(v string) *GetWorkflowInstanceResponseBody
	GetStartedAt() *string
	SetStateVersion(v int64) *GetWorkflowInstanceResponseBody
	GetStateVersion() *int64
	SetStatus(v string) *GetWorkflowInstanceResponseBody
	GetStatus() *string
	SetTemplateId(v string) *GetWorkflowInstanceResponseBody
	GetTemplateId() *string
	SetTitle(v string) *GetWorkflowInstanceResponseBody
	GetTitle() *string
	SetWorkItemId(v string) *GetWorkflowInstanceResponseBody
	GetWorkItemId() *string
}

type GetWorkflowInstanceResponseBody struct {
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
	CurrentNode *GetWorkflowInstanceResponseBodyCurrentNode `json:"currentNode,omitempty" xml:"currentNode,omitempty" type:"Struct"`
	// 全部活动节点，支持并行 DAG
	CurrentNodes []*GetWorkflowInstanceResponseBodyCurrentNodes `json:"currentNodes,omitempty" xml:"currentNodes,omitempty" type:"Repeated"`
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
	Progress *GetWorkflowInstanceResponseBodyProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
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

func (s GetWorkflowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBody) GetBusinessId() *string {
	return s.BusinessId
}

func (s *GetWorkflowInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkflowInstanceResponseBody) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkflowInstanceResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetWorkflowInstanceResponseBody) GetCurrentNode() *GetWorkflowInstanceResponseBodyCurrentNode {
	return s.CurrentNode
}

func (s *GetWorkflowInstanceResponseBody) GetCurrentNodes() []*GetWorkflowInstanceResponseBodyCurrentNodes {
	return s.CurrentNodes
}

func (s *GetWorkflowInstanceResponseBody) GetEndedAt() *string {
	return s.EndedAt
}

func (s *GetWorkflowInstanceResponseBody) GetInitiator() *string {
	return s.Initiator
}

func (s *GetWorkflowInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkflowInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkflowInstanceResponseBody) GetPauseScope() *string {
	return s.PauseScope
}

func (s *GetWorkflowInstanceResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *GetWorkflowInstanceResponseBody) GetProgress() *GetWorkflowInstanceResponseBodyProgress {
	return s.Progress
}

func (s *GetWorkflowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowInstanceResponseBody) GetReviewStatus() *string {
	return s.ReviewStatus
}

func (s *GetWorkflowInstanceResponseBody) GetRevision() *int64 {
	return s.Revision
}

func (s *GetWorkflowInstanceResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *GetWorkflowInstanceResponseBody) GetRuntimeStatus() *string {
	return s.RuntimeStatus
}

func (s *GetWorkflowInstanceResponseBody) GetStartedAt() *string {
	return s.StartedAt
}

func (s *GetWorkflowInstanceResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *GetWorkflowInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetWorkflowInstanceResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetWorkflowInstanceResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetWorkflowInstanceResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkflowInstanceResponseBody) SetBusinessId(v string) *GetWorkflowInstanceResponseBody {
	s.BusinessId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetCode(v string) *GetWorkflowInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetCollaborationGroupId(v string) *GetWorkflowInstanceResponseBody {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetCreatedAt(v string) *GetWorkflowInstanceResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetCurrentNode(v *GetWorkflowInstanceResponseBodyCurrentNode) *GetWorkflowInstanceResponseBody {
	s.CurrentNode = v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetCurrentNodes(v []*GetWorkflowInstanceResponseBodyCurrentNodes) *GetWorkflowInstanceResponseBody {
	s.CurrentNodes = v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetEndedAt(v string) *GetWorkflowInstanceResponseBody {
	s.EndedAt = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetInitiator(v string) *GetWorkflowInstanceResponseBody {
	s.Initiator = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetInstanceId(v string) *GetWorkflowInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetMessage(v string) *GetWorkflowInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetPauseScope(v string) *GetWorkflowInstanceResponseBody {
	s.PauseScope = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetPlanId(v string) *GetWorkflowInstanceResponseBody {
	s.PlanId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetProgress(v *GetWorkflowInstanceResponseBodyProgress) *GetWorkflowInstanceResponseBody {
	s.Progress = v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetRequestId(v string) *GetWorkflowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetReviewStatus(v string) *GetWorkflowInstanceResponseBody {
	s.ReviewStatus = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetRevision(v int64) *GetWorkflowInstanceResponseBody {
	s.Revision = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetRunId(v string) *GetWorkflowInstanceResponseBody {
	s.RunId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetRuntimeStatus(v string) *GetWorkflowInstanceResponseBody {
	s.RuntimeStatus = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetStartedAt(v string) *GetWorkflowInstanceResponseBody {
	s.StartedAt = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetStateVersion(v int64) *GetWorkflowInstanceResponseBody {
	s.StateVersion = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetStatus(v string) *GetWorkflowInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetTemplateId(v string) *GetWorkflowInstanceResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetTitle(v string) *GetWorkflowInstanceResponseBody {
	s.Title = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) SetWorkItemId(v string) *GetWorkflowInstanceResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBody) Validate() error {
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

type GetWorkflowInstanceResponseBodyCurrentNode struct {
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

func (s GetWorkflowInstanceResponseBodyCurrentNode) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyCurrentNode) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) GetHandler() *string {
	return s.Handler
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) GetHandlerId() *string {
	return s.HandlerId
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) GetNodeId() *string {
	return s.NodeId
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) GetNodeName() *string {
	return s.NodeName
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) GetNodeType() *string {
	return s.NodeType
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) GetStatus() *string {
	return s.Status
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) SetHandler(v string) *GetWorkflowInstanceResponseBodyCurrentNode {
	s.Handler = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) SetHandlerId(v string) *GetWorkflowInstanceResponseBodyCurrentNode {
	s.HandlerId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) SetNodeId(v string) *GetWorkflowInstanceResponseBodyCurrentNode {
	s.NodeId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) SetNodeName(v string) *GetWorkflowInstanceResponseBodyCurrentNode {
	s.NodeName = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) SetNodeType(v string) *GetWorkflowInstanceResponseBodyCurrentNode {
	s.NodeType = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) SetStatus(v string) *GetWorkflowInstanceResponseBodyCurrentNode {
	s.Status = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNode) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowInstanceResponseBodyCurrentNodes struct {
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

func (s GetWorkflowInstanceResponseBodyCurrentNodes) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyCurrentNodes) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) GetHandler() *string {
	return s.Handler
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) GetHandlerId() *string {
	return s.HandlerId
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) GetNodeType() *string {
	return s.NodeType
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) GetStatus() *string {
	return s.Status
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) SetHandler(v string) *GetWorkflowInstanceResponseBodyCurrentNodes {
	s.Handler = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) SetHandlerId(v string) *GetWorkflowInstanceResponseBodyCurrentNodes {
	s.HandlerId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) SetNodeId(v string) *GetWorkflowInstanceResponseBodyCurrentNodes {
	s.NodeId = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) SetNodeName(v string) *GetWorkflowInstanceResponseBodyCurrentNodes {
	s.NodeName = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) SetNodeType(v string) *GetWorkflowInstanceResponseBodyCurrentNodes {
	s.NodeType = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) SetStatus(v string) *GetWorkflowInstanceResponseBodyCurrentNodes {
	s.Status = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyCurrentNodes) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowInstanceResponseBodyProgress struct {
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

func (s GetWorkflowInstanceResponseBodyProgress) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowInstanceResponseBodyProgress) GoString() string {
	return s.String()
}

func (s *GetWorkflowInstanceResponseBodyProgress) GetCompletedCount() *int64 {
	return s.CompletedCount
}

func (s *GetWorkflowInstanceResponseBodyProgress) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetWorkflowInstanceResponseBodyProgress) SetCompletedCount(v int64) *GetWorkflowInstanceResponseBodyProgress {
	s.CompletedCount = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyProgress) SetTotalCount(v int64) *GetWorkflowInstanceResponseBodyProgress {
	s.TotalCount = &v
	return s
}

func (s *GetWorkflowInstanceResponseBodyProgress) Validate() error {
	return dara.Validate(s)
}

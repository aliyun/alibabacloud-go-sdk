// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkflowInstancesResponseBody
	GetCode() *string
	SetItems(v []*ListWorkflowInstancesResponseBodyItems) *ListWorkflowInstancesResponseBody
	GetItems() []*ListWorkflowInstancesResponseBodyItems
	SetMessage(v string) *ListWorkflowInstancesResponseBody
	GetMessage() *string
	SetPage(v int64) *ListWorkflowInstancesResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListWorkflowInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListWorkflowInstancesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListWorkflowInstancesResponseBody
	GetTotal() *int64
}

type ListWorkflowInstancesResponseBody struct {
	// 业务状态码，来自统一 POP 响应包装
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 当前页实例列表
	Items []*ListWorkflowInstancesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// 用户可见错误描述，来自统一 POP 响应包装
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 当前页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页数量
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 筛选后可见实例总数
	//
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListWorkflowInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWorkflowInstancesResponseBody) GetItems() []*ListWorkflowInstancesResponseBodyItems {
	return s.Items
}

func (s *ListWorkflowInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkflowInstancesResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkflowInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkflowInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowInstancesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListWorkflowInstancesResponseBody) SetCode(v string) *ListWorkflowInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetItems(v []*ListWorkflowInstancesResponseBodyItems) *ListWorkflowInstancesResponseBody {
	s.Items = v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetMessage(v string) *ListWorkflowInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetPage(v int64) *ListWorkflowInstancesResponseBody {
	s.Page = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetPageSize(v int64) *ListWorkflowInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetRequestId(v string) *ListWorkflowInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetTotal(v int64) *ListWorkflowInstancesResponseBody {
	s.Total = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowInstancesResponseBodyItems struct {
	// 空间内唯一外部需求编号
	//
	// example:
	//
	// REQ-1001
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
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
	CurrentNode *ListWorkflowInstancesResponseBodyItemsCurrentNode `json:"currentNode,omitempty" xml:"currentNode,omitempty" type:"Struct"`
	// 全部活动节点，支持并行 DAG
	CurrentNodes []*ListWorkflowInstancesResponseBodyItemsCurrentNodes `json:"currentNodes,omitempty" xml:"currentNodes,omitempty" type:"Repeated"`
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
	Progress *ListWorkflowInstancesResponseBodyItemsProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
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

func (s ListWorkflowInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyItems) GetBusinessId() *string {
	return s.BusinessId
}

func (s *ListWorkflowInstancesResponseBodyItems) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkflowInstancesResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListWorkflowInstancesResponseBodyItems) GetCurrentNode() *ListWorkflowInstancesResponseBodyItemsCurrentNode {
	return s.CurrentNode
}

func (s *ListWorkflowInstancesResponseBodyItems) GetCurrentNodes() []*ListWorkflowInstancesResponseBodyItemsCurrentNodes {
	return s.CurrentNodes
}

func (s *ListWorkflowInstancesResponseBodyItems) GetEndedAt() *string {
	return s.EndedAt
}

func (s *ListWorkflowInstancesResponseBodyItems) GetInitiator() *string {
	return s.Initiator
}

func (s *ListWorkflowInstancesResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWorkflowInstancesResponseBodyItems) GetPauseScope() *string {
	return s.PauseScope
}

func (s *ListWorkflowInstancesResponseBodyItems) GetPlanId() *string {
	return s.PlanId
}

func (s *ListWorkflowInstancesResponseBodyItems) GetProgress() *ListWorkflowInstancesResponseBodyItemsProgress {
	return s.Progress
}

func (s *ListWorkflowInstancesResponseBodyItems) GetReviewStatus() *string {
	return s.ReviewStatus
}

func (s *ListWorkflowInstancesResponseBodyItems) GetRevision() *int64 {
	return s.Revision
}

func (s *ListWorkflowInstancesResponseBodyItems) GetRunId() *string {
	return s.RunId
}

func (s *ListWorkflowInstancesResponseBodyItems) GetRuntimeStatus() *string {
	return s.RuntimeStatus
}

func (s *ListWorkflowInstancesResponseBodyItems) GetStartedAt() *string {
	return s.StartedAt
}

func (s *ListWorkflowInstancesResponseBodyItems) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *ListWorkflowInstancesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowInstancesResponseBodyItems) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListWorkflowInstancesResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *ListWorkflowInstancesResponseBodyItems) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ListWorkflowInstancesResponseBodyItems) SetBusinessId(v string) *ListWorkflowInstancesResponseBodyItems {
	s.BusinessId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetCollaborationGroupId(v string) *ListWorkflowInstancesResponseBodyItems {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetCreatedAt(v string) *ListWorkflowInstancesResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetCurrentNode(v *ListWorkflowInstancesResponseBodyItemsCurrentNode) *ListWorkflowInstancesResponseBodyItems {
	s.CurrentNode = v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetCurrentNodes(v []*ListWorkflowInstancesResponseBodyItemsCurrentNodes) *ListWorkflowInstancesResponseBodyItems {
	s.CurrentNodes = v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetEndedAt(v string) *ListWorkflowInstancesResponseBodyItems {
	s.EndedAt = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetInitiator(v string) *ListWorkflowInstancesResponseBodyItems {
	s.Initiator = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetInstanceId(v string) *ListWorkflowInstancesResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetPauseScope(v string) *ListWorkflowInstancesResponseBodyItems {
	s.PauseScope = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetPlanId(v string) *ListWorkflowInstancesResponseBodyItems {
	s.PlanId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetProgress(v *ListWorkflowInstancesResponseBodyItemsProgress) *ListWorkflowInstancesResponseBodyItems {
	s.Progress = v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetReviewStatus(v string) *ListWorkflowInstancesResponseBodyItems {
	s.ReviewStatus = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetRevision(v int64) *ListWorkflowInstancesResponseBodyItems {
	s.Revision = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetRunId(v string) *ListWorkflowInstancesResponseBodyItems {
	s.RunId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetRuntimeStatus(v string) *ListWorkflowInstancesResponseBodyItems {
	s.RuntimeStatus = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetStartedAt(v string) *ListWorkflowInstancesResponseBodyItems {
	s.StartedAt = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetStateVersion(v int64) *ListWorkflowInstancesResponseBodyItems {
	s.StateVersion = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetStatus(v string) *ListWorkflowInstancesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetTemplateId(v string) *ListWorkflowInstancesResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetTitle(v string) *ListWorkflowInstancesResponseBodyItems {
	s.Title = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) SetWorkItemId(v string) *ListWorkflowInstancesResponseBodyItems {
	s.WorkItemId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItems) Validate() error {
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

type ListWorkflowInstancesResponseBodyItemsCurrentNode struct {
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

func (s ListWorkflowInstancesResponseBodyItemsCurrentNode) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyItemsCurrentNode) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) GetHandler() *string {
	return s.Handler
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) GetHandlerId() *string {
	return s.HandlerId
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) GetNodeId() *string {
	return s.NodeId
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) GetNodeName() *string {
	return s.NodeName
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) GetNodeType() *string {
	return s.NodeType
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) SetHandler(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNode {
	s.Handler = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) SetHandlerId(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNode {
	s.HandlerId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) SetNodeId(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNode {
	s.NodeId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) SetNodeName(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNode {
	s.NodeName = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) SetNodeType(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNode {
	s.NodeType = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) SetStatus(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNode {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNode) Validate() error {
	return dara.Validate(s)
}

type ListWorkflowInstancesResponseBodyItemsCurrentNodes struct {
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

func (s ListWorkflowInstancesResponseBodyItemsCurrentNodes) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyItemsCurrentNodes) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) GetHandler() *string {
	return s.Handler
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) GetHandlerId() *string {
	return s.HandlerId
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) GetNodeName() *string {
	return s.NodeName
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) GetNodeType() *string {
	return s.NodeType
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) SetHandler(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNodes {
	s.Handler = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) SetHandlerId(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNodes {
	s.HandlerId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) SetNodeId(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNodes {
	s.NodeId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) SetNodeName(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNodes {
	s.NodeName = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) SetNodeType(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNodes {
	s.NodeType = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) SetStatus(v string) *ListWorkflowInstancesResponseBodyItemsCurrentNodes {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsCurrentNodes) Validate() error {
	return dara.Validate(s)
}

type ListWorkflowInstancesResponseBodyItemsProgress struct {
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

func (s ListWorkflowInstancesResponseBodyItemsProgress) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyItemsProgress) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyItemsProgress) GetCompletedCount() *int64 {
	return s.CompletedCount
}

func (s *ListWorkflowInstancesResponseBodyItemsProgress) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWorkflowInstancesResponseBodyItemsProgress) SetCompletedCount(v int64) *ListWorkflowInstancesResponseBodyItemsProgress {
	s.CompletedCount = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsProgress) SetTotalCount(v int64) *ListWorkflowInstancesResponseBodyItemsProgress {
	s.TotalCount = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyItemsProgress) Validate() error {
	return dara.Validate(s)
}

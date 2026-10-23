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
	// The business status code from the unified POP response wrapper.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of instances on the current page.
	Items []*ListWorkflowInstancesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The user-visible error message from the unified POP response wrapper.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of visible instances after filtering.
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
	// The unique external requirement ID in the space.
	//
	// example:
	//
	// REQ-1001
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// The collaboration space ID.
	//
	// example:
	//
	// cg_20261016_abc123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The time when the logical instance was created. The time is in the ISO 8601 format.
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The first active node. This value is null if the instance is in a terminal state.
	CurrentNode *ListWorkflowInstancesResponseBodyItemsCurrentNode `json:"currentNode,omitempty" xml:"currentNode,omitempty" type:"Struct"`
	// All active nodes. Parallel DAGs are supported.
	CurrentNodes []*ListWorkflowInstancesResponseBodyItemsCurrentNodes `json:"currentNodes,omitempty" xml:"currentNodes,omitempty" type:"Repeated"`
	// The time when the current plan reached a terminal state. This value is null if the plan has not ended.
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	EndedAt *string `json:"endedAt,omitempty" xml:"endedAt,omitempty"`
	// The external initiator tag. This tag is not used for authentication.
	//
	// example:
	//
	// External initiator Zhang San
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// The stable logical instance ID. This ID does not change when the work item is restarted.
	//
	// example:
	//
	// wfi_20261016_abc123
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The original plan pause scope.
	//
	// example:
	//
	// ALL
	PauseScope *string `json:"pauseScope,omitempty" xml:"pauseScope,omitempty"`
	// The revision ID of the current physical plan.
	//
	// example:
	//
	// plan_20261016_abc123
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// The progress of the visible graph.
	Progress *ListWorkflowInstancesResponseBodyItemsProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// The original plan review status.
	//
	// example:
	//
	// APPROVED
	ReviewStatus *string `json:"reviewStatus,omitempty" xml:"reviewStatus,omitempty"`
	// The revision number of the current physical plan.
	//
	// example:
	//
	// 3
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// The current run ID. This ID changes after a restart.
	//
	// example:
	//
	// run_20261016_abc123
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The original runtime status. Valid values: BUILDING, RUNNING, SUCCEEDED, FAILED, and CANCELLED.
	//
	// example:
	//
	// RUNNING
	RuntimeStatus *string `json:"runtimeStatus,omitempty" xml:"runtimeStatus,omitempty"`
	// The time when the current plan node started for the first time. This value is null if the node has not been executed.
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	StartedAt *string `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The runtime status version.
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
	// The ID of the workflow template bound at creation.
	//
	// example:
	//
	// wft_20261016_abc123
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The current requirement title.
	//
	// example:
	//
	// Quarterly business analysis report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The ID of the associated work item.
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
	// The display name of the current handler.
	//
	// example:
	//
	// Zhang San
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// The identifier of the handler or digital employee.
	//
	// example:
	//
	// user_10086
	HandlerId *string `json:"handlerId,omitempty" xml:"handlerId,omitempty"`
	// The nodeKey of the actual running node.
	//
	// example:
	//
	// node_key_confirm
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The display name of the node.
	//
	// example:
	//
	// Requirement confirmation
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// The actual node type. Valid values: AGENT, HUMAN, and CONDITION.
	//
	// example:
	//
	// AGENT
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// The running status of the node.
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
	// The display name of the current handler.
	//
	// example:
	//
	// Zhang San
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// The identifier of the handler or digital employee.
	//
	// example:
	//
	// user_10086
	HandlerId *string `json:"handlerId,omitempty" xml:"handlerId,omitempty"`
	// The nodeKey of the actual running node.
	//
	// example:
	//
	// node_key_confirm
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The display name of the node.
	//
	// example:
	//
	// Requirement confirmation
	NodeName *string `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	// The actual node type. Valid values: AGENT, HUMAN, and CONDITION.
	//
	// example:
	//
	// AGENT
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// The running status of the node.
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
	// The number of visible nodes that have succeeded or been skipped.
	//
	// example:
	//
	// 3
	CompletedCount *int64 `json:"completedCount,omitempty" xml:"completedCount,omitempty"`
	// The total number of visible nodes in the current graph. This number is not fixed to six.
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

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
	// The unique external requirement number within the workspace.
	//
	// example:
	//
	// REQ-1001
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// The business status code from the unified POP response wrapper.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the collaboration workspace.
	//
	// example:
	//
	// cg_20261016_abc123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The creation time of the logical instance in ISO 8601 format.
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The first active node. This value is null when the instance reaches a terminal state.
	CurrentNode *GetWorkflowInstanceResponseBodyCurrentNode `json:"currentNode,omitempty" xml:"currentNode,omitempty" type:"Struct"`
	// All active nodes. Parallel DAGs are supported.
	CurrentNodes []*GetWorkflowInstanceResponseBodyCurrentNodes `json:"currentNodes,omitempty" xml:"currentNodes,omitempty" type:"Repeated"`
	// The time when the current plan reached a terminal state. This value is null if the plan has not ended.
	//
	// example:
	//
	// 2026-10-16T10:30:00+08:00
	EndedAt *string `json:"endedAt,omitempty" xml:"endedAt,omitempty"`
	// The external initiator label. This label is not used for authentication.
	//
	// example:
	//
	// External Initiator Zhang San
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// The stable logical instance ID. This ID does not change when the work item is restarted.
	//
	// example:
	//
	// wfi_20261016_abc123
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The user-visible error description from the unified POP response wrapper.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The raw plan pause scope.
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
	// The visible graph progress.
	Progress *GetWorkflowInstanceResponseBodyProgress `json:"progress,omitempty" xml:"progress,omitempty" type:"Struct"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The raw plan review status.
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
	// The raw runtime status. Valid values: BUILDING, RUNNING, SUCCEEDED, FAILED, and CANCELLED.
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
	// Quarterly Business Analysis Report
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The ID of the associated work item.
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
	// The display name of the current handler.
	//
	// example:
	//
	// Zhang San
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// The identifier of the handler or digital worker.
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
	// Requirement Confirmation
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
	// The display name of the current handler.
	//
	// example:
	//
	// Zhang San
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// The identifier of the handler or digital worker.
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
	// Requirement Confirmation
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

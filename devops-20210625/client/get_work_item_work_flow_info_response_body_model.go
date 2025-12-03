// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemWorkFlowInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetWorkItemWorkFlowInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetWorkItemWorkFlowInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetWorkItemWorkFlowInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkItemWorkFlowInfoResponseBody
	GetSuccess() *bool
	SetWorkflow(v *GetWorkItemWorkFlowInfoResponseBodyWorkflow) *GetWorkItemWorkFlowInfoResponseBody
	GetWorkflow() *GetWorkItemWorkFlowInfoResponseBodyWorkflow
}

type GetWorkItemWorkFlowInfoResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success  *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
	Workflow *GetWorkItemWorkFlowInfoResponseBodyWorkflow `json:"workflow,omitempty" xml:"workflow,omitempty" type:"Struct"`
}

func (s GetWorkItemWorkFlowInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetWorkItemWorkFlowInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetWorkItemWorkFlowInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkItemWorkFlowInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkItemWorkFlowInfoResponseBody) GetWorkflow() *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	return s.Workflow
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetErrorCode(v string) *GetWorkItemWorkFlowInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetErrorMessage(v string) *GetWorkItemWorkFlowInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetRequestId(v string) *GetWorkItemWorkFlowInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetSuccess(v bool) *GetWorkItemWorkFlowInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBody) SetWorkflow(v *GetWorkItemWorkFlowInfoResponseBodyWorkflow) *GetWorkItemWorkFlowInfoResponseBody {
	s.Workflow = v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBody) Validate() error {
	if s.Workflow != nil {
		if err := s.Workflow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkItemWorkFlowInfoResponseBodyWorkflow struct {
	// example:
	//
	// 19xx7043xxxxxxx914
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 100005
	DefaultStatusIdentifier *string `json:"defaultStatusIdentifier,omitempty" xml:"defaultStatusIdentifier,omitempty"`
	// example:
	//
	// 工作流的描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1640850318000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1640850318000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// e8b26xxxxx6e76aa20xxxxx23
	OwnerSpaceIdentifier *string `json:"ownerSpaceIdentifier,omitempty" xml:"ownerSpaceIdentifier,omitempty"`
	// example:
	//
	// null
	OwnerSpaceType *string `json:"ownerSpaceType,omitempty" xml:"ownerSpaceType,omitempty"`
	// example:
	//
	// Project
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// system
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// null
	StatusOrder     *string                                                       `json:"statusOrder,omitempty" xml:"statusOrder,omitempty"`
	Statuses        []*GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses        `json:"statuses,omitempty" xml:"statuses,omitempty" type:"Repeated"`
	WorkflowActions []*GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions `json:"workflowActions,omitempty" xml:"workflowActions,omitempty" type:"Repeated"`
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflow) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflow) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetCreator() *string {
	return s.Creator
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetDefaultStatusIdentifier() *string {
	return s.DefaultStatusIdentifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetDescription() *string {
	return s.Description
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetModifier() *string {
	return s.Modifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetName() *string {
	return s.Name
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetOwnerSpaceIdentifier() *string {
	return s.OwnerSpaceIdentifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetOwnerSpaceType() *string {
	return s.OwnerSpaceType
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetSource() *string {
	return s.Source
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetStatusOrder() *string {
	return s.StatusOrder
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetStatuses() []*GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	return s.Statuses
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) GetWorkflowActions() []*GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	return s.WorkflowActions
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetCreator(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Creator = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetDefaultStatusIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.DefaultStatusIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetDescription(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Description = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetGmtCreate(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetGmtModified(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.GmtModified = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetModifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Modifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetName(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Name = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetOwnerSpaceIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.OwnerSpaceIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetOwnerSpaceType(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.OwnerSpaceType = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetResourceType(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.ResourceType = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetSource(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Source = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetStatusOrder(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.StatusOrder = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetStatuses(v []*GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.Statuses = v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) SetWorkflowActions(v []*GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) *GetWorkItemWorkFlowInfoResponseBodyWorkflow {
	s.WorkflowActions = v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflow) Validate() error {
	if s.Statuses != nil {
		for _, item := range s.Statuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WorkflowActions != nil {
		for _, item := range s.WorkflowActions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses struct {
	// example:
	//
	// 19xx7043xxxxxxx914
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// xxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1613805843000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1613805843000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 156603
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 待处理
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Workitem
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// system
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// 1
	WorkflowStageIdentifier *string `json:"workflowStageIdentifier,omitempty" xml:"workflowStageIdentifier,omitempty"`
	// example:
	//
	// 确认阶段
	WorkflowStageName *string `json:"workflowStageName,omitempty" xml:"workflowStageName,omitempty"`
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetCreator() *string {
	return s.Creator
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetDescription() *string {
	return s.Description
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetModifier() *string {
	return s.Modifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetName() *string {
	return s.Name
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetSource() *string {
	return s.Source
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetWorkflowStageIdentifier() *string {
	return s.WorkflowStageIdentifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) GetWorkflowStageName() *string {
	return s.WorkflowStageName
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetCreator(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Creator = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetDescription(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Description = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetGmtCreate(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.GmtCreate = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetGmtModified(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.GmtModified = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Identifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetModifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Modifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetName(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Name = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetResourceType(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.ResourceType = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetSource(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.Source = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetWorkflowStageIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.WorkflowStageIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) SetWorkflowStageName(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses {
	s.WorkflowStageName = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowStatuses) Validate() error {
	return dara.Validate(s)
}

type GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions struct {
	// example:
	//
	// 16274887
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// xxx
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 100011
	NextWorkflowStatusIdentifier *string `json:"nextWorkflowStatusIdentifier,omitempty" xml:"nextWorkflowStatusIdentifier,omitempty"`
	// example:
	//
	// fd0xxxxxd00d355b05dxxxx26
	WorkflowIdentifier *string `json:"workflowIdentifier,omitempty" xml:"workflowIdentifier,omitempty"`
	// example:
	//
	// 100005
	WorkflowStatusIdentifier *string `json:"workflowStatusIdentifier,omitempty" xml:"workflowStatusIdentifier,omitempty"`
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) GoString() string {
	return s.String()
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) GetId() *int64 {
	return s.Id
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) GetName() *string {
	return s.Name
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) GetNextWorkflowStatusIdentifier() *string {
	return s.NextWorkflowStatusIdentifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) GetWorkflowIdentifier() *string {
	return s.WorkflowIdentifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) GetWorkflowStatusIdentifier() *string {
	return s.WorkflowStatusIdentifier
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetId(v int64) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.Id = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetName(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.Name = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetNextWorkflowStatusIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.NextWorkflowStatusIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetWorkflowIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.WorkflowIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) SetWorkflowStatusIdentifier(v string) *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions {
	s.WorkflowStatusIdentifier = &v
	return s
}

func (s *GetWorkItemWorkFlowInfoResponseBodyWorkflowWorkflowActions) Validate() error {
	return dara.Validate(s)
}

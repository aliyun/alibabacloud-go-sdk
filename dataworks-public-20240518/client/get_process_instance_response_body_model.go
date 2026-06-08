// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProcessInstance(v *GetProcessInstanceResponseBodyProcessInstance) *GetProcessInstanceResponseBody
	GetProcessInstance() *GetProcessInstanceResponseBodyProcessInstance
	SetRequestId(v string) *GetProcessInstanceResponseBody
	GetRequestId() *string
}

type GetProcessInstanceResponseBody struct {
	ProcessInstance *GetProcessInstanceResponseBodyProcessInstance `json:"ProcessInstance,omitempty" xml:"ProcessInstance,omitempty" type:"Struct"`
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProcessInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBody) GetProcessInstance() *GetProcessInstanceResponseBodyProcessInstance {
	return s.ProcessInstance
}

func (s *GetProcessInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProcessInstanceResponseBody) SetProcessInstance(v *GetProcessInstanceResponseBodyProcessInstance) *GetProcessInstanceResponseBody {
	s.ProcessInstance = v
	return s
}

func (s *GetProcessInstanceResponseBody) SetRequestId(v string) *GetProcessInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProcessInstanceResponseBody) Validate() error {
	if s.ProcessInstance != nil {
		if err := s.ProcessInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProcessInstanceResponseBodyProcessInstance struct {
	// example:
	//
	// 1107558004253538
	Applicator *string `json:"Applicator,omitempty" xml:"Applicator,omitempty"`
	// example:
	//
	// test_account
	ApplicatorName            *string                                                                 `json:"ApplicatorName,omitempty" xml:"ApplicatorName,omitempty"`
	ApprovalProcessDefinition *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition `json:"ApprovalProcessDefinition,omitempty" xml:"ApprovalProcessDefinition,omitempty" type:"Struct"`
	ApprovalTasks             []*GetProcessInstanceResponseBodyProcessInstanceApprovalTasks           `json:"ApprovalTasks,omitempty" xml:"ApprovalTasks,omitempty" type:"Repeated"`
	// example:
	//
	// S-400007:ODPS acl auth failed. odps table acl auth failed
	AuthErrorMessage *string `json:"AuthErrorMessage,omitempty" xml:"AuthErrorMessage,omitempty"`
	// example:
	//
	// 332066440109224007
	Id     *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 2026-05-25 10:20:18 CST
	StartTime interface{} `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetProcessInstanceResponseBodyProcessInstance) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponseBodyProcessInstance) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetApplicator() *string {
	return s.Applicator
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetApplicatorName() *string {
	return s.ApplicatorName
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetApprovalProcessDefinition() *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	return s.ApprovalProcessDefinition
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetApprovalTasks() []*GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	return s.ApprovalTasks
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetAuthErrorMessage() *string {
	return s.AuthErrorMessage
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetId() *string {
	return s.Id
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetReason() *string {
	return s.Reason
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetStartTime() interface{} {
	return s.StartTime
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetStatus() *string {
	return s.Status
}

func (s *GetProcessInstanceResponseBodyProcessInstance) GetTitle() *string {
	return s.Title
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetApplicator(v string) *GetProcessInstanceResponseBodyProcessInstance {
	s.Applicator = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetApplicatorName(v string) *GetProcessInstanceResponseBodyProcessInstance {
	s.ApplicatorName = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetApprovalProcessDefinition(v *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) *GetProcessInstanceResponseBodyProcessInstance {
	s.ApprovalProcessDefinition = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetApprovalTasks(v []*GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) *GetProcessInstanceResponseBodyProcessInstance {
	s.ApprovalTasks = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetAuthErrorMessage(v string) *GetProcessInstanceResponseBodyProcessInstance {
	s.AuthErrorMessage = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetId(v string) *GetProcessInstanceResponseBodyProcessInstance {
	s.Id = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetReason(v string) *GetProcessInstanceResponseBodyProcessInstance {
	s.Reason = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetStartTime(v interface{}) *GetProcessInstanceResponseBodyProcessInstance {
	s.StartTime = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetStatus(v string) *GetProcessInstanceResponseBodyProcessInstance {
	s.Status = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) SetTitle(v string) *GetProcessInstanceResponseBodyProcessInstance {
	s.Title = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstance) Validate() error {
	if s.ApprovalProcessDefinition != nil {
		if err := s.ApprovalProcessDefinition.Validate(); err != nil {
			return err
		}
	}
	if s.ApprovalTasks != nil {
		for _, item := range s.ApprovalTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition struct {
	ApprovalNodes []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes `json:"ApprovalNodes,omitempty" xml:"ApprovalNodes,omitempty" type:"Repeated"`
	Description   *string                                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 323861511451222099
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// SYSTEM_GENERATE_DEFAULT
	Name                 *string                                                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	NotificationServices []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices `json:"NotificationServices,omitempty" xml:"NotificationServices,omitempty" type:"Repeated"`
	RuleConditions       []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions       `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// example:
	//
	// Table
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GetApprovalNodes() []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes {
	return s.ApprovalNodes
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GetDescription() *string {
	return s.Description
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GetId() *string {
	return s.Id
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GetName() *string {
	return s.Name
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GetNotificationServices() []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices {
	return s.NotificationServices
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GetRuleConditions() []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions {
	return s.RuleConditions
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GetSubType() *string {
	return s.SubType
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) GetType() *string {
	return s.Type
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) SetApprovalNodes(v []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	s.ApprovalNodes = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) SetDescription(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	s.Description = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) SetEnabled(v bool) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	s.Enabled = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) SetId(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	s.Id = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) SetName(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	s.Name = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) SetNotificationServices(v []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	s.NotificationServices = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) SetRuleConditions(v []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	s.RuleConditions = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) SetSubType(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	s.SubType = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) SetType(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition {
	s.Type = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition) Validate() error {
	if s.ApprovalNodes != nil {
		for _, item := range s.ApprovalNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NotificationServices != nil {
		for _, item := range s.NotificationServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleConditions != nil {
		for _, item := range s.RuleConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes struct {
	// example:
	//
	// DataWorksProjectRole
	AccountType *string   `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Assignees   []*string `json:"Assignees,omitempty" xml:"Assignees,omitempty" type:"Repeated"`
	// example:
	//
	// none
	ExtensionProperties *string `json:"ExtensionProperties,omitempty" xml:"ExtensionProperties,omitempty"`
	// example:
	//
	// 7a809b6a-2a62-4c6c-9c23-c2a145e3877d
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// default-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) GetAccountType() *string {
	return s.AccountType
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) GetAssignees() []*string {
	return s.Assignees
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) GetExtensionProperties() *string {
	return s.ExtensionProperties
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) GetId() *string {
	return s.Id
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) GetName() *string {
	return s.Name
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) SetAccountType(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes {
	s.AccountType = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) SetAssignees(v []*string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes {
	s.Assignees = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) SetExtensionProperties(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes {
	s.ExtensionProperties = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) SetId(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes {
	s.Id = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) SetName(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes {
	s.Name = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes) Validate() error {
	return dara.Validate(s)
}

type GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices struct {
	// example:
	//
	// Mail
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// https://dingtalk
	Receiver *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) GetChannel() *string {
	return s.Channel
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) GetExtension() *string {
	return s.Extension
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) GetReceiver() *string {
	return s.Receiver
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) SetChannel(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices {
	s.Channel = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) SetExtension(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices {
	s.Extension = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) SetReceiver(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices {
	s.Receiver = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices) Validate() error {
	return dara.Validate(s)
}

type GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions struct {
	// example:
	//
	// ((#odpsProject==\\"PX_BEIJING_TEST\\"))
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// Deployment
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// odpsProject
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) GetExpression() *string {
	return s.Expression
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) GetScope() *string {
	return s.Scope
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) GetType() *string {
	return s.Type
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) SetExpression(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions {
	s.Expression = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) SetScope(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions {
	s.Scope = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) SetType(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions {
	s.Type = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions) Validate() error {
	return dara.Validate(s)
}

type GetProcessInstanceResponseBodyProcessInstanceApprovalTasks struct {
	ApprovalComment *string `json:"ApprovalComment,omitempty" xml:"ApprovalComment,omitempty"`
	// example:
	//
	// deny
	ApprovalDecision *string                                                                 `json:"ApprovalDecision,omitempty" xml:"ApprovalDecision,omitempty"`
	ApprovalNode     *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode `json:"ApprovalNode,omitempty" xml:"ApprovalNode,omitempty" type:"Struct"`
	// example:
	//
	// 207947399706614297
	Assignee     *string `json:"Assignee,omitempty" xml:"Assignee,omitempty"`
	AssigneeName *string `json:"AssigneeName,omitempty" xml:"AssigneeName,omitempty"`
	// example:
	//
	// 1715590800000
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// example:
	//
	// 1715587200000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// task_001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Aborted
	Status         *string                                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskCandidates []*GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates `json:"TaskCandidates,omitempty" xml:"TaskCandidates,omitempty" type:"Repeated"`
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetApprovalComment() *string {
	return s.ApprovalComment
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetApprovalDecision() *string {
	return s.ApprovalDecision
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetApprovalNode() *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode {
	return s.ApprovalNode
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetAssignee() *string {
	return s.Assignee
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetAssigneeName() *string {
	return s.AssigneeName
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetId() *string {
	return s.Id
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetStatus() *string {
	return s.Status
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) GetTaskCandidates() []*GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates {
	return s.TaskCandidates
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetApprovalComment(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.ApprovalComment = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetApprovalDecision(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.ApprovalDecision = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetApprovalNode(v *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.ApprovalNode = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetAssignee(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.Assignee = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetAssigneeName(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.AssigneeName = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetCompleteTime(v int64) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.CompleteTime = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetCreateTime(v int64) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.CreateTime = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetId(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.Id = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetStatus(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.Status = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) SetTaskCandidates(v []*GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks {
	s.TaskCandidates = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasks) Validate() error {
	if s.ApprovalNode != nil {
		if err := s.ApprovalNode.Validate(); err != nil {
			return err
		}
	}
	if s.TaskCandidates != nil {
		for _, item := range s.TaskCandidates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode struct {
	// example:
	//
	// DataWorksProjectRole
	AccountType *string   `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Assignees   []*string `json:"Assignees,omitempty" xml:"Assignees,omitempty" type:"Repeated"`
	// example:
	//
	// 7a809b6a-2a62-4c6c-9c23-c2a145e3877d
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// default-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) GetAccountType() *string {
	return s.AccountType
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) GetAssignees() []*string {
	return s.Assignees
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) GetId() *string {
	return s.Id
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) GetName() *string {
	return s.Name
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) SetAccountType(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode {
	s.AccountType = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) SetAssignees(v []*string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode {
	s.Assignees = v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) SetId(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode {
	s.Id = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) SetName(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode {
	s.Name = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode) Validate() error {
	return dara.Validate(s)
}

type GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates struct {
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// 207947397776614297
	MemberUserId *string `json:"MemberUserId,omitempty" xml:"MemberUserId,omitempty"`
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates) GetMemberName() *string {
	return s.MemberName
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates) GetMemberUserId() *string {
	return s.MemberUserId
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates) SetMemberName(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates {
	s.MemberName = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates) SetMemberUserId(v string) *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates {
	s.MemberUserId = &v
	return s
}

func (s *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksTaskCandidates) Validate() error {
	return dara.Validate(s)
}

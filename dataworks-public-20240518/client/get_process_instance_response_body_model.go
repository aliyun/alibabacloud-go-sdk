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
	// Details of the approval process instance.
	ProcessInstance *GetProcessInstanceResponseBodyProcessInstance `json:"ProcessInstance,omitempty" xml:"ProcessInstance,omitempty" type:"Struct"`
	// The request ID. Use this ID to locate logs and troubleshoot issues.
	//
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
	// The user ID of the applicant.
	//
	// example:
	//
	// 1107558004253538
	Applicator *string `json:"Applicator,omitempty" xml:"Applicator,omitempty"`
	// The username of the applicant\\"s Alibaba Cloud account.
	//
	// example:
	//
	// test_account
	ApplicatorName *string `json:"ApplicatorName,omitempty" xml:"ApplicatorName,omitempty"`
	// The approval policy applied to this process instance.
	ApprovalProcessDefinition *GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinition `json:"ApprovalProcessDefinition,omitempty" xml:"ApprovalProcessDefinition,omitempty" type:"Struct"`
	// The approval tasks.
	ApprovalTasks []*GetProcessInstanceResponseBodyProcessInstanceApprovalTasks `json:"ApprovalTasks,omitempty" xml:"ApprovalTasks,omitempty" type:"Repeated"`
	// The authorization failure message.
	//
	// **Note**: This parameter is returned only if the authorization fails.
	//
	// example:
	//
	// S-400007:ODPS acl auth failed. odps table acl auth failed
	AuthErrorMessage *string `json:"AuthErrorMessage,omitempty" xml:"AuthErrorMessage,omitempty"`
	// The process instance ID.
	//
	// example:
	//
	// 332066440109224007
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The reason for the request.
	//
	// example:
	//
	// 业务需要
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The time when the approval process started.
	//
	// example:
	//
	// 2026-05-25 10:20:18 CST
	StartTime interface{} `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the process instance. Valid values:
	//
	// - `Completed`: The request is approved.
	//
	// - `Running`: The request is in the approval process.
	//
	// - `Aborted`: The request is withdrawn.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the process instance.
	//
	// example:
	//
	// MaxCompute表权限申请
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
	// The approval nodes.
	ApprovalNodes []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionApprovalNodes `json:"ApprovalNodes,omitempty" xml:"ApprovalNodes,omitempty" type:"Repeated"`
	// The description of the approval policy.
	//
	// example:
	//
	// 流程定义描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the policy is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The approval policy ID.
	//
	// example:
	//
	// 323861511451222099
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the approval policy.
	//
	// example:
	//
	// SYSTEM_GENERATE_DEFAULT
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The notification services.
	NotificationServices []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionNotificationServices `json:"NotificationServices,omitempty" xml:"NotificationServices,omitempty" type:"Repeated"`
	// The rules that determine when the approval policy takes effect.
	RuleConditions []*GetProcessInstanceResponseBodyProcessInstanceApprovalProcessDefinitionRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// The subtype of the approval policy. Valid values:
	//
	// - `Table`
	//
	// - `Column`
	//
	// - `Database`
	//
	// - `Schema`
	//
	// - `Default`
	//
	// example:
	//
	// Table
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// The type of the approval policy. Valid values:
	//
	// - `MaxCompute`
	//
	// - `DataService`
	//
	// - `DlfV1` (Custom creation is not supported)
	//
	// - `Extension`
	//
	// - `Hologres`
	//
	// - `Emr` (Custom creation is not supported)
	//
	// - `DataAssetGovernance` (Custom creation is not supported)
	//
	// - `Lindorm` (Custom creation is not supported)
	//
	// - `StarRocks` (Custom creation is not supported)
	//
	// - `DlfNext` (Custom creation is not supported)
	//
	// - `DataWorks` (Custom creation is not supported)
	//
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
	// The type of the approver for the node. Valid values:
	//
	// - `DataWorksProjectRole`: A workspace role
	//
	// - `DataWorksProjectMember`: A workspace member
	//
	// - `TableAdministrator`: A table administrator
	//
	// - `TableOrProjectAdministrator`: A table or workspace administrator
	//
	// - `AliyunResourceOwner`: An Alibaba Cloud account
	//
	// - `MaxComputeRole`: A MaxCompute role
	//
	// - `DLFAdmin`: A DlfLegacy administrator
	//
	// - `DLFNextAdmin`: A DLFNext administrator
	//
	// - `TenantRole`: A tenant role
	//
	// - `EmrAdministrator`: An Emr administrator
	//
	// - `LindormAdministrator`: A Lindorm administrator
	//
	// - `AliyunRamUser`: A RAM user
	//
	// example:
	//
	// DataWorksProjectRole
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The specified approvers.
	//
	// The contents of this parameter depend on the `AccountType` value:
	//
	// - If `AccountType` is `DataWorksProjectMember`, this parameter contains the user IDs of workspace members.
	//
	// - If `AccountType` is `DataWorksProjectRole`, this parameter contains the codes of workspace roles.
	//
	// - If `AccountType` is `MaxComputeRole`, this parameter contains the MaxCompute roles.
	//
	// - If `AccountType` is `TenantRole`, this parameter contains the codes of tenant roles.
	//
	// - If `AccountType` is `AliyunRamUser`, this parameter contains the user IDs of RAM users.
	Assignees []*string `json:"Assignees,omitempty" xml:"Assignees,omitempty" type:"Repeated"`
	// The extended description of the approval node.
	//
	// example:
	//
	// none
	ExtensionProperties *string `json:"ExtensionProperties,omitempty" xml:"ExtensionProperties,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 7a809b6a-2a62-4c6c-9c23-c2a145e3877d
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node name.
	//
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
	// The notification channel. Valid values:
	//
	// - `Mail`
	//
	// - `Sms`
	//
	// - `DingRobot`
	//
	// - `Weixin`
	//
	// example:
	//
	// Mail
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Additional information in JSON format. For example, `{"atAll":"true"}` indicates whether to @all members.
	//
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// If `Channel` is set to `DingRobot` or `Weixin`, the value of this parameter must be the webhook URL.
	//
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
	// The expression of the rule condition. Format: `((#type==\\"typeValue\\"))`.
	//
	// example:
	//
	// ((#odpsProject==\\"PX_BEIJING_TEST\\"))
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The rule scope. Valid values:
	//
	// - `Deployment`: Determines whether the policy applies when a request is submitted.
	//
	// - `Running`: Determines whether to skip approval while the process instance runs. This value is supported only for MaxCompute approval policies.
	//
	// example:
	//
	// Deployment
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The type of the rule condition. Valid values:
	//
	// - `odpsProject`: Applies to a specific MaxCompute project.
	//
	// - `hologresInstanceId`: Applies to a specific Hologres instance.
	//
	// - `sensibleLevel`: Applies to a specific security level.
	//
	// - `tableGuid`: Applies to a specific table.
	//
	// - `projectId`: Applies to a specific workspace.
	//
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
	// The approval comment.
	//
	// example:
	//
	// 同意
	ApprovalComment *string `json:"ApprovalComment,omitempty" xml:"ApprovalComment,omitempty"`
	// The approval decision. Valid values:
	//
	// - `Agree`
	//
	// - `Deny`
	//
	// example:
	//
	// Deny
	ApprovalDecision *string `json:"ApprovalDecision,omitempty" xml:"ApprovalDecision,omitempty"`
	// The approval node from the corresponding approval policy.
	ApprovalNode *GetProcessInstanceResponseBodyProcessInstanceApprovalTasksApprovalNode `json:"ApprovalNode,omitempty" xml:"ApprovalNode,omitempty" type:"Struct"`
	// The user ID of the actual approver.
	//
	// example:
	//
	// 207947399706614297
	Assignee *string `json:"Assignee,omitempty" xml:"Assignee,omitempty"`
	// The name of the actual approver.
	//
	// example:
	//
	// 李四
	AssigneeName *string `json:"AssigneeName,omitempty" xml:"AssigneeName,omitempty"`
	// The time when the task was completed.
	//
	// example:
	//
	// 1715590800000
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 1715587200000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The approval task ID.
	//
	// example:
	//
	// task_001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the task. Valid values:
	//
	// - `Completed`: The task is complete.
	//
	// - `Pending`: The task is pending.
	//
	// - `Aborted`: The task is aborted.
	//
	// example:
	//
	// Aborted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The candidate approvers for the task.
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
	// The type of the approver for the node. Valid values:
	//
	// - `DataWorksProjectRole`: A workspace role
	//
	// - `DataWorksProjectMember`: A workspace member
	//
	// - `TableAdministrator`: A table administrator
	//
	// - `TableOrProjectAdministrator`: A table or workspace administrator
	//
	// - `AliyunResourceOwner`: An Alibaba Cloud account
	//
	// - `MaxComputeRole`: A MaxCompute role
	//
	// - `DLFAdmin`: A DlfLegacy administrator
	//
	// - `DLFNextAdmin`: A DLFNext administrator
	//
	// - `TenantRole`: A tenant role
	//
	// - `EmrAdministrator`: An Emr administrator
	//
	// - `LindormAdministrator`: A Lindorm administrator
	//
	// - `AliyunRamUser`: A RAM user
	//
	// example:
	//
	// DataWorksProjectRole
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The specified approvers.
	//
	// The contents of this parameter depend on the `AccountType` value:
	//
	// - If `AccountType` is `DataWorksProjectMember`, this parameter contains the user IDs of workspace members.
	//
	// - If `AccountType` is `DataWorksProjectRole`, this parameter contains the codes of workspace roles.
	//
	// - If `AccountType` is `MaxComputeRole`, this parameter contains the MaxCompute roles.
	//
	// - If `AccountType` is `TenantRole`, this parameter contains the codes of tenant roles.
	//
	// - If `AccountType` is `AliyunRamUser`, this parameter contains the user IDs of RAM users.
	Assignees []*string `json:"Assignees,omitempty" xml:"Assignees,omitempty" type:"Repeated"`
	// The node ID.
	//
	// example:
	//
	// 7a809b6a-2a62-4c6c-9c23-c2a145e3877d
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node name.
	//
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
	// The name of the approver.
	//
	// example:
	//
	// 李四
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The user ID of the approver.
	//
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

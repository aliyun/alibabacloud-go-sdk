// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProcessDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalNodes(v []*UpdateProcessDefinitionRequestApprovalNodes) *UpdateProcessDefinitionRequest
	GetApprovalNodes() []*UpdateProcessDefinitionRequestApprovalNodes
	SetClientToken(v string) *UpdateProcessDefinitionRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateProcessDefinitionRequest
	GetDescription() *string
	SetId(v string) *UpdateProcessDefinitionRequest
	GetId() *string
	SetName(v string) *UpdateProcessDefinitionRequest
	GetName() *string
	SetNotificationServices(v []*UpdateProcessDefinitionRequestNotificationServices) *UpdateProcessDefinitionRequest
	GetNotificationServices() []*UpdateProcessDefinitionRequestNotificationServices
	SetRuleConditions(v []*UpdateProcessDefinitionRequestRuleConditions) *UpdateProcessDefinitionRequest
	GetRuleConditions() []*UpdateProcessDefinitionRequestRuleConditions
}

type UpdateProcessDefinitionRequest struct {
	// A list of approval nodes. This parameter does not apply to system policies.
	ApprovalNodes []*UpdateProcessDefinitionRequestApprovalNodes `json:"ApprovalNodes,omitempty" xml:"ApprovalNodes,omitempty" type:"Repeated"`
	// An idempotent parameter. It ensures that retried requests do not result in duplicate operations.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the process definition.
	//
	// example:
	//
	// lwt_ide_simple 项目 MaxCompute 表审批策略
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the process definition.
	//
	// This parameter is required.
	//
	// example:
	//
	// 177554881536128
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the process definition.
	//
	// example:
	//
	// MaxCompute 表审批
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The notification service configurations.
	NotificationServices []*UpdateProcessDefinitionRequestNotificationServices `json:"NotificationServices,omitempty" xml:"NotificationServices,omitempty" type:"Repeated"`
	// A list of rule conditions. This parameter does not apply to system policies.
	RuleConditions []*UpdateProcessDefinitionRequestRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
}

func (s UpdateProcessDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionRequest) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionRequest) GetApprovalNodes() []*UpdateProcessDefinitionRequestApprovalNodes {
	return s.ApprovalNodes
}

func (s *UpdateProcessDefinitionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateProcessDefinitionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProcessDefinitionRequest) GetId() *string {
	return s.Id
}

func (s *UpdateProcessDefinitionRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProcessDefinitionRequest) GetNotificationServices() []*UpdateProcessDefinitionRequestNotificationServices {
	return s.NotificationServices
}

func (s *UpdateProcessDefinitionRequest) GetRuleConditions() []*UpdateProcessDefinitionRequestRuleConditions {
	return s.RuleConditions
}

func (s *UpdateProcessDefinitionRequest) SetApprovalNodes(v []*UpdateProcessDefinitionRequestApprovalNodes) *UpdateProcessDefinitionRequest {
	s.ApprovalNodes = v
	return s
}

func (s *UpdateProcessDefinitionRequest) SetClientToken(v string) *UpdateProcessDefinitionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateProcessDefinitionRequest) SetDescription(v string) *UpdateProcessDefinitionRequest {
	s.Description = &v
	return s
}

func (s *UpdateProcessDefinitionRequest) SetId(v string) *UpdateProcessDefinitionRequest {
	s.Id = &v
	return s
}

func (s *UpdateProcessDefinitionRequest) SetName(v string) *UpdateProcessDefinitionRequest {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionRequest) SetNotificationServices(v []*UpdateProcessDefinitionRequestNotificationServices) *UpdateProcessDefinitionRequest {
	s.NotificationServices = v
	return s
}

func (s *UpdateProcessDefinitionRequest) SetRuleConditions(v []*UpdateProcessDefinitionRequestRuleConditions) *UpdateProcessDefinitionRequest {
	s.RuleConditions = v
	return s
}

func (s *UpdateProcessDefinitionRequest) Validate() error {
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

type UpdateProcessDefinitionRequestApprovalNodes struct {
	// The approver type for the node. Valid values:
	//
	// - `DataWorksProjectRole`: A workspace role.
	//
	// - `DataWorksProjectMember`: A workspace member.
	//
	// - `TableAdministrator`: A table administrator.
	//
	// - `TableOrProjectAdministrator`: The administrator of the table or project.
	//
	// - `AliyunResourceOwner`: An Alibaba Cloud account.
	//
	// - `MaxComputeRole`: A MaxCompute administrator.
	//
	// - `DLFAdmin`: A DlfLegacy administrator.
	//
	// - `DLFNextAdmin`: A DLFNext administrator.
	//
	// - `TenantRole`: A tenant role.
	//
	// - `EmrAdministrator`: An EMR administrator.
	//
	// - `LindormAdministrator`: A Lindorm administrator.
	//
	// - `AliyunRamUser`: A RAM user.
	//
	// example:
	//
	// TableOrProjectAdministrator
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// Specifies the approvers. The required value depends on the `AccountType`:
	//
	// - If `AccountType` is `DataWorksProjectMember`, this parameter specifies the user IDs of workspace members.
	//
	// - If `AccountType` is `DataWorksProjectRole`, this parameter specifies the codes of workspace roles.
	//
	// - If `AccountType` is `MaxComputeRole`, this parameter specifies the MaxCompute roles.
	//
	// - If `AccountType` is `TenantRole`, this parameter specifies the codes of tenant roles.
	//
	// - If `AccountType` is `AliyunRamUser`, this parameter specifies the user IDs of RAM users.
	//
	// example:
	//
	// DataWorksProjectMember
	Assignees *string `json:"Assignees,omitempty" xml:"Assignees,omitempty"`
	// Additional properties that are required for specific `AccountType` values:
	//
	// - If `AccountType` is `DataWorksProjectMember`: The key is `projectId` and the value is the user ID of a workspace member. Use commas (,) to separate multiple user IDs.
	//
	// - If `AccountType` is `MaxComputeRole`: The key is the MaxCompute project name and the value is the role name in MaxCompute. Use commas (,) to separate multiple role names.
	ExtensionProperties map[string]interface{} `json:"ExtensionProperties,omitempty" xml:"ExtensionProperties,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// default-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateProcessDefinitionRequestApprovalNodes) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionRequestApprovalNodes) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionRequestApprovalNodes) GetAccountType() *string {
	return s.AccountType
}

func (s *UpdateProcessDefinitionRequestApprovalNodes) GetAssignees() *string {
	return s.Assignees
}

func (s *UpdateProcessDefinitionRequestApprovalNodes) GetExtensionProperties() map[string]interface{} {
	return s.ExtensionProperties
}

func (s *UpdateProcessDefinitionRequestApprovalNodes) GetName() *string {
	return s.Name
}

func (s *UpdateProcessDefinitionRequestApprovalNodes) SetAccountType(v string) *UpdateProcessDefinitionRequestApprovalNodes {
	s.AccountType = &v
	return s
}

func (s *UpdateProcessDefinitionRequestApprovalNodes) SetAssignees(v string) *UpdateProcessDefinitionRequestApprovalNodes {
	s.Assignees = &v
	return s
}

func (s *UpdateProcessDefinitionRequestApprovalNodes) SetExtensionProperties(v map[string]interface{}) *UpdateProcessDefinitionRequestApprovalNodes {
	s.ExtensionProperties = v
	return s
}

func (s *UpdateProcessDefinitionRequestApprovalNodes) SetName(v string) *UpdateProcessDefinitionRequestApprovalNodes {
	s.Name = &v
	return s
}

func (s *UpdateProcessDefinitionRequestApprovalNodes) Validate() error {
	return dara.Validate(s)
}

type UpdateProcessDefinitionRequestNotificationServices struct {
	// The notification channel. Valid values:
	//
	// - Mail
	//
	// - Sms
	//
	// - DingRobot
	//
	// - Weixin
	//
	// example:
	//
	// Mail
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Additional information in JSON format. For example, use {"atAll":"true"} to specify whether to notify all members.
	//
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The webhook URL. This parameter is required when `Channel` is set to `DingRobot` or `Weixin`.
	//
	// example:
	//
	// https://dingtalk.com
	Receiver *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
}

func (s UpdateProcessDefinitionRequestNotificationServices) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionRequestNotificationServices) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionRequestNotificationServices) GetChannel() *string {
	return s.Channel
}

func (s *UpdateProcessDefinitionRequestNotificationServices) GetExtension() *string {
	return s.Extension
}

func (s *UpdateProcessDefinitionRequestNotificationServices) GetReceiver() *string {
	return s.Receiver
}

func (s *UpdateProcessDefinitionRequestNotificationServices) SetChannel(v string) *UpdateProcessDefinitionRequestNotificationServices {
	s.Channel = &v
	return s
}

func (s *UpdateProcessDefinitionRequestNotificationServices) SetExtension(v string) *UpdateProcessDefinitionRequestNotificationServices {
	s.Extension = &v
	return s
}

func (s *UpdateProcessDefinitionRequestNotificationServices) SetReceiver(v string) *UpdateProcessDefinitionRequestNotificationServices {
	s.Receiver = &v
	return s
}

func (s *UpdateProcessDefinitionRequestNotificationServices) Validate() error {
	return dara.Validate(s)
}

type UpdateProcessDefinitionRequestRuleConditions struct {
	// The conditional expression. Format: ((#type==\\"typeValue\\")). For example: ((#odpsProject==\\"PX_BEIJING_TEST\\")).
	//
	// example:
	//
	// ((#odpsProject==\\"PX_BEIJING_TEST\\"))
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The phase in which the rule takes effect. Valid values:
	//
	// - **Deployment**: Determines whether the approval policy applies when an application is submitted.
	//
	// - **Running**: Determines whether to skip the approval during the approval process. This phase is supported only for MaxCompute.
	//
	// example:
	//
	// Deployment
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The type of the condition. Valid values:
	//
	// - `odpsProject`
	//
	// - `hologresInstanceId`
	//
	// - `sensibleLevel`
	//
	// - `tableGuid`
	//
	// - `projectId`
	//
	// example:
	//
	// odpsProject
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateProcessDefinitionRequestRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateProcessDefinitionRequestRuleConditions) GoString() string {
	return s.String()
}

func (s *UpdateProcessDefinitionRequestRuleConditions) GetExpression() *string {
	return s.Expression
}

func (s *UpdateProcessDefinitionRequestRuleConditions) GetScope() *string {
	return s.Scope
}

func (s *UpdateProcessDefinitionRequestRuleConditions) GetType() *string {
	return s.Type
}

func (s *UpdateProcessDefinitionRequestRuleConditions) SetExpression(v string) *UpdateProcessDefinitionRequestRuleConditions {
	s.Expression = &v
	return s
}

func (s *UpdateProcessDefinitionRequestRuleConditions) SetScope(v string) *UpdateProcessDefinitionRequestRuleConditions {
	s.Scope = &v
	return s
}

func (s *UpdateProcessDefinitionRequestRuleConditions) SetType(v string) *UpdateProcessDefinitionRequestRuleConditions {
	s.Type = &v
	return s
}

func (s *UpdateProcessDefinitionRequestRuleConditions) Validate() error {
	return dara.Validate(s)
}

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
	// The list of approval nodes. Note: System policies do not take effect.
	ApprovalNodes []*UpdateProcessDefinitionRequestApprovalNodes `json:"ApprovalNodes,omitempty" xml:"ApprovalNodes,omitempty" type:"Repeated"`
	// The idempotency token. Used to prevent duplicate operations caused by multiple calls.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the approval process definition.
	//
	// example:
	//
	// Approval policy for MaxCompute tables in the lwt_ide_simple project
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the approval process definition.
	//
	// This parameter is required.
	//
	// example:
	//
	// 177554881536128
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the approval process definition.
	//
	// example:
	//
	// MaxCompute table approval
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The notification service declaration.
	NotificationServices []*UpdateProcessDefinitionRequestNotificationServices `json:"NotificationServices,omitempty" xml:"NotificationServices,omitempty" type:"Repeated"`
	// The list of conditional rules. Note: System policies do not take effect.
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
	// The type of approver for the node. Valid values:
	//
	// - DataWorksProjectRole: Workspace role
	//
	// - DataWorksProjectMember: Workspace member
	//
	// - TableAdministrator: Table owner
	//
	// - TableOrProjectAdministrator: Table or workspace administrator
	//
	// - AliyunResourceOwner: Alibaba Cloud account
	//
	// - MaxComputeRole: MaxCompute administrator
	//
	// - DLFAdmin: DLF legacy administrator
	//
	// - DLFNextAdmin: DLFNext administrator
	//
	// - TenantRole: Tenant role
	//
	// - EmrAdministrator: EMR administrator
	//
	// - LindormAdministrator: Lindorm administrator
	//
	// - AliyunRamUser: RAM user
	//
	// example:
	//
	// TableOrProjectAdministrator
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The assignees. The semantics vary based on the AccountType:
	//
	// - DataWorksProjectMember: Specifies the user ID of the workspace member.
	//
	// - DataWorksProjectRole: Specifies the code of the workspace role.
	//
	// - MaxComputeRole: Specifies the MaxCompute role.
	//
	// - TenantRole: Specifies the code of the tenant role.
	//
	// - AliyunRamUser: Specifies the RAM user ID.
	//
	// example:
	//
	// DataWorksProjectMember
	Assignees *string `json:"Assignees,omitempty" xml:"Assignees,omitempty"`
	// The extension properties. Different additional declarations are required based on the AccountType:
	//
	// - `DataWorksProjectMember`: The key is projectId, and the value is the user ID of the workspace member. Separate multiple values with commas.
	//
	// - `MaxComputeRole`: The key is the MaxCompute project, and the value is the role name in MaxCompute. Separate multiple values with commas.
	//
	// example:
	//
	// {"449656": "role_project_admin,role_project_pe"}
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
	// The extended information in JSON format. For example, `{"atAll":"true"}` specifies whether to mention all users.
	//
	// example:
	//
	// {"atAll":"true"}
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The Webhook URL. This parameter is required when the channel is DingRobot or Weixin.
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
	// The conditional expression, such as `((#type==\\"typeValue\\"))` or `((#odpsProject==\\"PX_BEIJING_TEST\\"))`.
	//
	// example:
	//
	// ((#odpsProject==\\"PX_BEIJING_TEST\\"))
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The stage where the rule takes effect. Valid values:
	//
	// - **Deployment**: Used to determine whether the approval policy is matched when the application is submitted.
	//
	// - **Running**: Used to determine whether to skip approval while the approval process runs. [Supported only for MaxCompute]
	//
	// example:
	//
	// Deployment
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The condition type. Valid values:
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

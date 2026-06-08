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
	ApprovalNodes []*UpdateProcessDefinitionRequestApprovalNodes `json:"ApprovalNodes,omitempty" xml:"ApprovalNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 177554881536128
	Id                   *string                                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                 *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	NotificationServices []*UpdateProcessDefinitionRequestNotificationServices `json:"NotificationServices,omitempty" xml:"NotificationServices,omitempty" type:"Repeated"`
	RuleConditions       []*UpdateProcessDefinitionRequestRuleConditions       `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
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
	// example:
	//
	// TableOrProjectAdministrator
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// example:
	//
	// DataWorksProjectMember
	Assignees           *string                `json:"Assignees,omitempty" xml:"Assignees,omitempty"`
	ExtensionProperties map[string]interface{} `json:"ExtensionProperties,omitempty" xml:"ExtensionProperties,omitempty"`
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

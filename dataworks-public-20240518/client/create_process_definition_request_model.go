// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProcessDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalNodes(v []*CreateProcessDefinitionRequestApprovalNodes) *CreateProcessDefinitionRequest
	GetApprovalNodes() []*CreateProcessDefinitionRequestApprovalNodes
	SetClientToken(v string) *CreateProcessDefinitionRequest
	GetClientToken() *string
	SetDescription(v string) *CreateProcessDefinitionRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateProcessDefinitionRequest
	GetEnabled() *bool
	SetName(v string) *CreateProcessDefinitionRequest
	GetName() *string
	SetNotificationServices(v []*CreateProcessDefinitionRequestNotificationServices) *CreateProcessDefinitionRequest
	GetNotificationServices() []*CreateProcessDefinitionRequestNotificationServices
	SetRuleConditions(v []*CreateProcessDefinitionRequestRuleConditions) *CreateProcessDefinitionRequest
	GetRuleConditions() []*CreateProcessDefinitionRequestRuleConditions
	SetSubType(v string) *CreateProcessDefinitionRequest
	GetSubType() *string
	SetType(v string) *CreateProcessDefinitionRequest
	GetType() *string
}

type CreateProcessDefinitionRequest struct {
	// This parameter is required.
	ApprovalNodes []*CreateProcessDefinitionRequestApprovalNodes `json:"ApprovalNodes,omitempty" xml:"ApprovalNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 0000-ABCD-EFG****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Name                 *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	NotificationServices []*CreateProcessDefinitionRequestNotificationServices `json:"NotificationServices,omitempty" xml:"NotificationServices,omitempty" type:"Repeated"`
	// This parameter is required.
	RuleConditions []*CreateProcessDefinitionRequestRuleConditions `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// example:
	//
	// Table
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// Extension
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateProcessDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionRequest) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionRequest) GetApprovalNodes() []*CreateProcessDefinitionRequestApprovalNodes {
	return s.ApprovalNodes
}

func (s *CreateProcessDefinitionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProcessDefinitionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProcessDefinitionRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateProcessDefinitionRequest) GetName() *string {
	return s.Name
}

func (s *CreateProcessDefinitionRequest) GetNotificationServices() []*CreateProcessDefinitionRequestNotificationServices {
	return s.NotificationServices
}

func (s *CreateProcessDefinitionRequest) GetRuleConditions() []*CreateProcessDefinitionRequestRuleConditions {
	return s.RuleConditions
}

func (s *CreateProcessDefinitionRequest) GetSubType() *string {
	return s.SubType
}

func (s *CreateProcessDefinitionRequest) GetType() *string {
	return s.Type
}

func (s *CreateProcessDefinitionRequest) SetApprovalNodes(v []*CreateProcessDefinitionRequestApprovalNodes) *CreateProcessDefinitionRequest {
	s.ApprovalNodes = v
	return s
}

func (s *CreateProcessDefinitionRequest) SetClientToken(v string) *CreateProcessDefinitionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProcessDefinitionRequest) SetDescription(v string) *CreateProcessDefinitionRequest {
	s.Description = &v
	return s
}

func (s *CreateProcessDefinitionRequest) SetEnabled(v bool) *CreateProcessDefinitionRequest {
	s.Enabled = &v
	return s
}

func (s *CreateProcessDefinitionRequest) SetName(v string) *CreateProcessDefinitionRequest {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionRequest) SetNotificationServices(v []*CreateProcessDefinitionRequestNotificationServices) *CreateProcessDefinitionRequest {
	s.NotificationServices = v
	return s
}

func (s *CreateProcessDefinitionRequest) SetRuleConditions(v []*CreateProcessDefinitionRequestRuleConditions) *CreateProcessDefinitionRequest {
	s.RuleConditions = v
	return s
}

func (s *CreateProcessDefinitionRequest) SetSubType(v string) *CreateProcessDefinitionRequest {
	s.SubType = &v
	return s
}

func (s *CreateProcessDefinitionRequest) SetType(v string) *CreateProcessDefinitionRequest {
	s.Type = &v
	return s
}

func (s *CreateProcessDefinitionRequest) Validate() error {
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

type CreateProcessDefinitionRequestApprovalNodes struct {
	// example:
	//
	// TableOrProjectAdministrator
	AccountType         *string                `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Assignees           []*string              `json:"Assignees,omitempty" xml:"Assignees,omitempty" type:"Repeated"`
	ExtensionProperties map[string]interface{} `json:"ExtensionProperties,omitempty" xml:"ExtensionProperties,omitempty"`
	// example:
	//
	// default-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateProcessDefinitionRequestApprovalNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionRequestApprovalNodes) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionRequestApprovalNodes) GetAccountType() *string {
	return s.AccountType
}

func (s *CreateProcessDefinitionRequestApprovalNodes) GetAssignees() []*string {
	return s.Assignees
}

func (s *CreateProcessDefinitionRequestApprovalNodes) GetExtensionProperties() map[string]interface{} {
	return s.ExtensionProperties
}

func (s *CreateProcessDefinitionRequestApprovalNodes) GetName() *string {
	return s.Name
}

func (s *CreateProcessDefinitionRequestApprovalNodes) SetAccountType(v string) *CreateProcessDefinitionRequestApprovalNodes {
	s.AccountType = &v
	return s
}

func (s *CreateProcessDefinitionRequestApprovalNodes) SetAssignees(v []*string) *CreateProcessDefinitionRequestApprovalNodes {
	s.Assignees = v
	return s
}

func (s *CreateProcessDefinitionRequestApprovalNodes) SetExtensionProperties(v map[string]interface{}) *CreateProcessDefinitionRequestApprovalNodes {
	s.ExtensionProperties = v
	return s
}

func (s *CreateProcessDefinitionRequestApprovalNodes) SetName(v string) *CreateProcessDefinitionRequestApprovalNodes {
	s.Name = &v
	return s
}

func (s *CreateProcessDefinitionRequestApprovalNodes) Validate() error {
	return dara.Validate(s)
}

type CreateProcessDefinitionRequestNotificationServices struct {
	// example:
	//
	// DingRobot
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

func (s CreateProcessDefinitionRequestNotificationServices) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionRequestNotificationServices) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionRequestNotificationServices) GetChannel() *string {
	return s.Channel
}

func (s *CreateProcessDefinitionRequestNotificationServices) GetExtension() *string {
	return s.Extension
}

func (s *CreateProcessDefinitionRequestNotificationServices) GetReceiver() *string {
	return s.Receiver
}

func (s *CreateProcessDefinitionRequestNotificationServices) SetChannel(v string) *CreateProcessDefinitionRequestNotificationServices {
	s.Channel = &v
	return s
}

func (s *CreateProcessDefinitionRequestNotificationServices) SetExtension(v string) *CreateProcessDefinitionRequestNotificationServices {
	s.Extension = &v
	return s
}

func (s *CreateProcessDefinitionRequestNotificationServices) SetReceiver(v string) *CreateProcessDefinitionRequestNotificationServices {
	s.Receiver = &v
	return s
}

func (s *CreateProcessDefinitionRequestNotificationServices) Validate() error {
	return dara.Validate(s)
}

type CreateProcessDefinitionRequestRuleConditions struct {
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

func (s CreateProcessDefinitionRequestRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateProcessDefinitionRequestRuleConditions) GoString() string {
	return s.String()
}

func (s *CreateProcessDefinitionRequestRuleConditions) GetExpression() *string {
	return s.Expression
}

func (s *CreateProcessDefinitionRequestRuleConditions) GetScope() *string {
	return s.Scope
}

func (s *CreateProcessDefinitionRequestRuleConditions) GetType() *string {
	return s.Type
}

func (s *CreateProcessDefinitionRequestRuleConditions) SetExpression(v string) *CreateProcessDefinitionRequestRuleConditions {
	s.Expression = &v
	return s
}

func (s *CreateProcessDefinitionRequestRuleConditions) SetScope(v string) *CreateProcessDefinitionRequestRuleConditions {
	s.Scope = &v
	return s
}

func (s *CreateProcessDefinitionRequestRuleConditions) SetType(v string) *CreateProcessDefinitionRequestRuleConditions {
	s.Type = &v
	return s
}

func (s *CreateProcessDefinitionRequestRuleConditions) Validate() error {
	return dara.Validate(s)
}

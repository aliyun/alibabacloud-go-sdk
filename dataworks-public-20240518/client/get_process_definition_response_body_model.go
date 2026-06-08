// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessDefinitionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProcessDefinition(v *GetProcessDefinitionResponseBodyProcessDefinition) *GetProcessDefinitionResponseBody
	GetProcessDefinition() *GetProcessDefinitionResponseBodyProcessDefinition
	SetRequestId(v string) *GetProcessDefinitionResponseBody
	GetRequestId() *string
}

type GetProcessDefinitionResponseBody struct {
	ProcessDefinition *GetProcessDefinitionResponseBodyProcessDefinition `json:"ProcessDefinition,omitempty" xml:"ProcessDefinition,omitempty" type:"Struct"`
	// example:
	//
	// 0bc5df3a17***903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProcessDefinitionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBody) GetProcessDefinition() *GetProcessDefinitionResponseBodyProcessDefinition {
	return s.ProcessDefinition
}

func (s *GetProcessDefinitionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProcessDefinitionResponseBody) SetProcessDefinition(v *GetProcessDefinitionResponseBodyProcessDefinition) *GetProcessDefinitionResponseBody {
	s.ProcessDefinition = v
	return s
}

func (s *GetProcessDefinitionResponseBody) SetRequestId(v string) *GetProcessDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProcessDefinitionResponseBody) Validate() error {
	if s.ProcessDefinition != nil {
		if err := s.ProcessDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProcessDefinitionResponseBodyProcessDefinition struct {
	ApprovalNodes []*GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes `json:"ApprovalNodes,omitempty" xml:"ApprovalNodes,omitempty" type:"Repeated"`
	Description   *string                                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 210001039767
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsSystem             *bool                                                                    `json:"IsSystem,omitempty" xml:"IsSystem,omitempty"`
	Name                 *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	NotificationServices []*GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices `json:"NotificationServices,omitempty" xml:"NotificationServices,omitempty" type:"Repeated"`
	RuleConditions       []*GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions       `json:"RuleConditions,omitempty" xml:"RuleConditions,omitempty" type:"Repeated"`
	// example:
	//
	// Table
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProcessDefinitionResponseBodyProcessDefinition) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyProcessDefinition) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetApprovalNodes() []*GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes {
	return s.ApprovalNodes
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetDescription() *string {
	return s.Description
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetId() *string {
	return s.Id
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetIsSystem() *bool {
	return s.IsSystem
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetName() *string {
	return s.Name
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetNotificationServices() []*GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices {
	return s.NotificationServices
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetRuleConditions() []*GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions {
	return s.RuleConditions
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetSubType() *string {
	return s.SubType
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) GetType() *string {
	return s.Type
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetApprovalNodes(v []*GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.ApprovalNodes = v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetDescription(v string) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.Description = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetEnabled(v bool) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.Enabled = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetId(v string) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.Id = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetIsSystem(v bool) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.IsSystem = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetName(v string) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.Name = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetNotificationServices(v []*GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.NotificationServices = v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetRuleConditions(v []*GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.RuleConditions = v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetSubType(v string) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.SubType = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) SetType(v string) *GetProcessDefinitionResponseBodyProcessDefinition {
	s.Type = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinition) Validate() error {
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

type GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes struct {
	// example:
	//
	// TableOrProjectAdministrator
	AccountType         *string                `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Assignees           []*string              `json:"Assignees,omitempty" xml:"Assignees,omitempty" type:"Repeated"`
	ExtensionProperties map[string]interface{} `json:"ExtensionProperties,omitempty" xml:"ExtensionProperties,omitempty"`
	// example:
	//
	// 7a809b6a-2a62-4c6c-9c23-c2a145e3877d
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// default-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) GetAccountType() *string {
	return s.AccountType
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) GetAssignees() []*string {
	return s.Assignees
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) GetExtensionProperties() map[string]interface{} {
	return s.ExtensionProperties
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) GetId() *string {
	return s.Id
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) GetName() *string {
	return s.Name
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) SetAccountType(v string) *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes {
	s.AccountType = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) SetAssignees(v []*string) *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes {
	s.Assignees = v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) SetExtensionProperties(v map[string]interface{}) *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes {
	s.ExtensionProperties = v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) SetId(v string) *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes {
	s.Id = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) SetName(v string) *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes {
	s.Name = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionApprovalNodes) Validate() error {
	return dara.Validate(s)
}

type GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices struct {
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

func (s GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) GetChannel() *string {
	return s.Channel
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) GetExtension() *string {
	return s.Extension
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) GetReceiver() *string {
	return s.Receiver
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) SetChannel(v string) *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices {
	s.Channel = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) SetExtension(v string) *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices {
	s.Extension = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) SetReceiver(v string) *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices {
	s.Receiver = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionNotificationServices) Validate() error {
	return dara.Validate(s)
}

type GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions struct {
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

func (s GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) GetExpression() *string {
	return s.Expression
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) GetScope() *string {
	return s.Scope
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) GetType() *string {
	return s.Type
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) SetExpression(v string) *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions {
	s.Expression = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) SetScope(v string) *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions {
	s.Scope = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) SetType(v string) *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions {
	s.Type = &v
	return s
}

func (s *GetProcessDefinitionResponseBodyProcessDefinitionRuleConditions) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoRepairPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListAutoRepairPoliciesResponseBodyItems) *ListAutoRepairPoliciesResponseBody
	GetItems() []*ListAutoRepairPoliciesResponseBodyItems
}

type ListAutoRepairPoliciesResponseBody struct {
	// A list of auto-repair rules.
	Items []*ListAutoRepairPoliciesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListAutoRepairPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBody) GetItems() []*ListAutoRepairPoliciesResponseBodyItems {
	return s.Items
}

func (s *ListAutoRepairPoliciesResponseBody) SetItems(v []*ListAutoRepairPoliciesResponseBodyItems) *ListAutoRepairPoliciesResponseBody {
	s.Items = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBody) Validate() error {
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

type ListAutoRepairPoliciesResponseBodyItems struct {
	// The ID of the auto-repair rule.
	//
	// example:
	//
	// r-xxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the auto-repair rule.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The IDs of the resources that the auto-repair rule affects.
	ResourceIds []*string `json:"resource_ids,omitempty" xml:"resource_ids,omitempty" type:"Repeated"`
	// The resource sub-type that the auto-repair rule affects.
	//
	// example:
	//
	// ess
	ResourceSubType *string `json:"resource_sub_type,omitempty" xml:"resource_sub_type,omitempty"`
	// The resource type that the auto-repair rule affects.
	//
	// example:
	//
	// nodepool
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// A list of auto-repair sub-rules.
	Rules []*ListAutoRepairPoliciesResponseBodyItemsRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s ListAutoRepairPoliciesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBodyItems) GetId() *string {
	return s.Id
}

func (s *ListAutoRepairPoliciesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListAutoRepairPoliciesResponseBodyItems) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListAutoRepairPoliciesResponseBodyItems) GetResourceSubType() *string {
	return s.ResourceSubType
}

func (s *ListAutoRepairPoliciesResponseBodyItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAutoRepairPoliciesResponseBodyItems) GetRules() []*ListAutoRepairPoliciesResponseBodyItemsRules {
	return s.Rules
}

func (s *ListAutoRepairPoliciesResponseBodyItems) SetId(v string) *ListAutoRepairPoliciesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItems) SetName(v string) *ListAutoRepairPoliciesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItems) SetResourceIds(v []*string) *ListAutoRepairPoliciesResponseBodyItems {
	s.ResourceIds = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItems) SetResourceSubType(v string) *ListAutoRepairPoliciesResponseBodyItems {
	s.ResourceSubType = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItems) SetResourceType(v string) *ListAutoRepairPoliciesResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItems) SetRules(v []*ListAutoRepairPoliciesResponseBodyItemsRules) *ListAutoRepairPoliciesResponseBodyItems {
	s.Rules = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItems) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAutoRepairPoliciesResponseBodyItemsRules struct {
	// A list of identified incidents.
	Incidents []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidents `json:"incidents,omitempty" xml:"incidents,omitempty" type:"Repeated"`
	// The repair procedure.
	RepairProcedure []*ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure `json:"repair_procedure,omitempty" xml:"repair_procedure,omitempty" type:"Repeated"`
}

func (s ListAutoRepairPoliciesResponseBodyItemsRules) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBodyItemsRules) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRules) GetIncidents() []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidents {
	return s.Incidents
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRules) GetRepairProcedure() []*ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure {
	return s.RepairProcedure
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRules) SetIncidents(v []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) *ListAutoRepairPoliciesResponseBodyItemsRules {
	s.Incidents = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRules) SetRepairProcedure(v []*ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) *ListAutoRepairPoliciesResponseBodyItemsRules {
	s.RepairProcedure = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRules) Validate() error {
	if s.Incidents != nil {
		for _, item := range s.Incidents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RepairProcedure != nil {
		for _, item := range s.RepairProcedure {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAutoRepairPoliciesResponseBodyItemsRulesIncidents struct {
	// The incident name.
	//
	// example:
	//
	// Node.FaultNeedReboot.HOST
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The diagnosis type.
	//
	// example:
	//
	// system
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) GetName() *string {
	return s.Name
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) GetType() *string {
	return s.Type
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) SetName(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents {
	s.Name = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) SetType(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents {
	s.Type = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) Validate() error {
	return dara.Validate(s)
}

type ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure struct {
	// The configuration parameters for the procedure step.
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The manual intervention settings for this procedure step.
	Intervention *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention `json:"intervention,omitempty" xml:"intervention,omitempty" type:"Struct"`
	// The name of the procedure step.
	//
	// example:
	//
	// QuarantineGPU
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) GetIntervention() *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention {
	return s.Intervention
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) GetName() *string {
	return s.Name
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) SetConfig(v map[string]interface{}) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure {
	s.Config = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) SetIntervention(v *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure {
	s.Intervention = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) SetName(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure {
	s.Name = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure) Validate() error {
	if s.Intervention != nil {
		if err := s.Intervention.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention struct {
	// The configuration for the approval label. Applying this label to the node authorizes Container Service for Kubernetes (ACK) to execute the action for this repair step. After the step is complete, ACK automatically removes both the inquiry and approval labels. If the approval label is not applied promptly, the repair process will not proceed, and the node may remain in an unhealthy state.
	ApprovedLabel *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel `json:"approved_label,omitempty" xml:"approved_label,omitempty" type:"Struct"`
	// Determines whether manual approval is required for the repair step.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The configuration for the authorization inquiry label. When this repair step starts, Container Service for Kubernetes (ACK) applies this label to the node and pauses, awaiting approval before executing the step\\"s action.
	InquiringLabel *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel `json:"inquiring_label,omitempty" xml:"inquiring_label,omitempty" type:"Struct"`
	// The manual approval type.
	//
	// example:
	//
	// label
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) GetApprovedLabel() *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel {
	return s.ApprovedLabel
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) GetEnable() *bool {
	return s.Enable
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) GetInquiringLabel() *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel {
	return s.InquiringLabel
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) GetType() *string {
	return s.Type
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) SetApprovedLabel(v *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention {
	s.ApprovedLabel = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) SetEnable(v bool) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention {
	s.Enable = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) SetInquiringLabel(v *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention {
	s.InquiringLabel = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) SetType(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention {
	s.Type = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) Validate() error {
	if s.ApprovedLabel != nil {
		if err := s.ApprovedLabel.Validate(); err != nil {
			return err
		}
	}
	if s.InquiringLabel != nil {
		if err := s.InquiringLabel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel struct {
	// The key of the label.
	//
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the label.
	//
	// example:
	//
	// approved
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel) GetKey() *string {
	return s.Key
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel) GetValue() *string {
	return s.Value
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel) SetKey(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel {
	s.Key = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel) SetValue(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel {
	s.Value = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel) Validate() error {
	return dara.Validate(s)
}

type ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel struct {
	// The key of the label.
	//
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the label.
	//
	// example:
	//
	// inquiring
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel) GetKey() *string {
	return s.Key
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel) GetValue() *string {
	return s.Value
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel) SetKey(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel {
	s.Key = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel) SetValue(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel {
	s.Value = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel) Validate() error {
	return dara.Validate(s)
}

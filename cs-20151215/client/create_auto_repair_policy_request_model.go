// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoRepairPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateAutoRepairPolicyRequest
	GetName() *string
	SetResourceSubType(v string) *CreateAutoRepairPolicyRequest
	GetResourceSubType() *string
	SetResourceType(v string) *CreateAutoRepairPolicyRequest
	GetResourceType() *string
	SetRules(v []*CreateAutoRepairPolicyRequestRules) *CreateAutoRepairPolicyRequest
	GetRules() []*CreateAutoRepairPolicyRequestRules
}

type CreateAutoRepairPolicyRequest struct {
	// The name of the auto repair policy.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource subtype to which the auto repair policy applies.
	//
	// example:
	//
	// ess
	ResourceSubType *string `json:"resource_sub_type,omitempty" xml:"resource_sub_type,omitempty"`
	// The resource type to which the auto repair policy applies.
	//
	// example:
	//
	// nodepool
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The sub-rules for the auto repair policy.
	//
	// example:
	//
	// ["np-xxx"]
	Rules []*CreateAutoRepairPolicyRequestRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s CreateAutoRepairPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateAutoRepairPolicyRequest) GetResourceSubType() *string {
	return s.ResourceSubType
}

func (s *CreateAutoRepairPolicyRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateAutoRepairPolicyRequest) GetRules() []*CreateAutoRepairPolicyRequestRules {
	return s.Rules
}

func (s *CreateAutoRepairPolicyRequest) SetName(v string) *CreateAutoRepairPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateAutoRepairPolicyRequest) SetResourceSubType(v string) *CreateAutoRepairPolicyRequest {
	s.ResourceSubType = &v
	return s
}

func (s *CreateAutoRepairPolicyRequest) SetResourceType(v string) *CreateAutoRepairPolicyRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateAutoRepairPolicyRequest) SetRules(v []*CreateAutoRepairPolicyRequestRules) *CreateAutoRepairPolicyRequest {
	s.Rules = v
	return s
}

func (s *CreateAutoRepairPolicyRequest) Validate() error {
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

type CreateAutoRepairPolicyRequestRules struct {
	// The incidents that the rule detects.
	Incidents []*CreateAutoRepairPolicyRequestRulesIncidents `json:"incidents,omitempty" xml:"incidents,omitempty" type:"Repeated"`
	// The repair procedure.
	RepairProcedure []*CreateAutoRepairPolicyRequestRulesRepairProcedure `json:"repair_procedure,omitempty" xml:"repair_procedure,omitempty" type:"Repeated"`
}

func (s CreateAutoRepairPolicyRequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyRequestRules) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyRequestRules) GetIncidents() []*CreateAutoRepairPolicyRequestRulesIncidents {
	return s.Incidents
}

func (s *CreateAutoRepairPolicyRequestRules) GetRepairProcedure() []*CreateAutoRepairPolicyRequestRulesRepairProcedure {
	return s.RepairProcedure
}

func (s *CreateAutoRepairPolicyRequestRules) SetIncidents(v []*CreateAutoRepairPolicyRequestRulesIncidents) *CreateAutoRepairPolicyRequestRules {
	s.Incidents = v
	return s
}

func (s *CreateAutoRepairPolicyRequestRules) SetRepairProcedure(v []*CreateAutoRepairPolicyRequestRulesRepairProcedure) *CreateAutoRepairPolicyRequestRules {
	s.RepairProcedure = v
	return s
}

func (s *CreateAutoRepairPolicyRequestRules) Validate() error {
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

type CreateAutoRepairPolicyRequestRulesIncidents struct {
	// The incident name.
	//
	// example:
	//
	// Node.FaultNeedReboot.HOST
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The incident type.
	//
	// example:
	//
	// system
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAutoRepairPolicyRequestRulesIncidents) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyRequestRulesIncidents) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) GetName() *string {
	return s.Name
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) GetType() *string {
	return s.Type
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) SetName(v string) *CreateAutoRepairPolicyRequestRulesIncidents {
	s.Name = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) SetType(v string) *CreateAutoRepairPolicyRequestRulesIncidents {
	s.Type = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) Validate() error {
	return dara.Validate(s)
}

type CreateAutoRepairPolicyRequestRulesRepairProcedure struct {
	// Configuration parameters for the repair step.
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// Settings for manual intervention.
	Intervention *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention `json:"intervention,omitempty" xml:"intervention,omitempty" type:"Struct"`
	// The name of the repair step.
	//
	// example:
	//
	// Drain
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAutoRepairPolicyRequestRulesRepairProcedure) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyRequestRulesRepairProcedure) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedure) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedure) GetIntervention() *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	return s.Intervention
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedure) GetName() *string {
	return s.Name
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedure) SetConfig(v map[string]interface{}) *CreateAutoRepairPolicyRequestRulesRepairProcedure {
	s.Config = v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedure) SetIntervention(v *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) *CreateAutoRepairPolicyRequestRulesRepairProcedure {
	s.Intervention = v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedure) SetName(v string) *CreateAutoRepairPolicyRequestRulesRepairProcedure {
	s.Name = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedure) Validate() error {
	if s.Intervention != nil {
		if err := s.Intervention.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention struct {
	// The label that grants authorization for the repair step. To approve the step, add this label to the node. After the action is complete, ACK automatically removes both the inquiry and approval labels for this step. If this label is not added promptly, the repair procedure halts and the node remains impaired.
	ApprovedLabel *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel `json:"approved_label,omitempty" xml:"approved_label,omitempty" type:"Struct"`
	// Specifies whether to enable manual approval.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The label used to request authorization for the repair step. When this step begins, ACK applies this label to the node and waits for approval before performing the action.
	InquiringLabel *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel `json:"inquiring_label,omitempty" xml:"inquiring_label,omitempty" type:"Struct"`
	// The manual approval type.
	//
	// example:
	//
	// label
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) GetApprovedLabel() *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel {
	return s.ApprovedLabel
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) GetEnable() *bool {
	return s.Enable
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) GetInquiringLabel() *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel {
	return s.InquiringLabel
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) GetType() *string {
	return s.Type
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) SetApprovedLabel(v *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	s.ApprovedLabel = v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) SetEnable(v bool) *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	s.Enable = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) SetInquiringLabel(v *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	s.InquiringLabel = v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) SetType(v string) *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	s.Type = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention) Validate() error {
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

type CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel struct {
	// The label key.
	//
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The label value.
	//
	// example:
	//
	// approved
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) GetKey() *string {
	return s.Key
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) GetValue() *string {
	return s.Value
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) SetKey(v string) *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel {
	s.Key = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) SetValue(v string) *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel {
	s.Value = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) Validate() error {
	return dara.Validate(s)
}

type CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel struct {
	// The label key.
	//
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The label value.
	//
	// example:
	//
	// inquiring
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) GetKey() *string {
	return s.Key
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) GetValue() *string {
	return s.Value
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) SetKey(v string) *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel {
	s.Key = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) SetValue(v string) *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel {
	s.Value = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) Validate() error {
	return dara.Validate(s)
}

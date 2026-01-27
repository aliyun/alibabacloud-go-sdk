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
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// ess
	ResourceSubType *string `json:"resource_sub_type,omitempty" xml:"resource_sub_type,omitempty"`
	// example:
	//
	// nodepool
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
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
	Incidents       []*CreateAutoRepairPolicyRequestRulesIncidents       `json:"incidents,omitempty" xml:"incidents,omitempty" type:"Repeated"`
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
	Conditions []*CreateAutoRepairPolicyRequestRulesIncidentsConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Events     []*CreateAutoRepairPolicyRequestRulesIncidentsEvents     `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// example:
	//
	// Node.FaultNeedReboot.HOST
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

func (s *CreateAutoRepairPolicyRequestRulesIncidents) GetConditions() []*CreateAutoRepairPolicyRequestRulesIncidentsConditions {
	return s.Conditions
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) GetEvents() []*CreateAutoRepairPolicyRequestRulesIncidentsEvents {
	return s.Events
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) GetName() *string {
	return s.Name
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) GetType() *string {
	return s.Type
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) SetConditions(v []*CreateAutoRepairPolicyRequestRulesIncidentsConditions) *CreateAutoRepairPolicyRequestRulesIncidents {
	s.Conditions = v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesIncidents) SetEvents(v []*CreateAutoRepairPolicyRequestRulesIncidentsEvents) *CreateAutoRepairPolicyRequestRulesIncidents {
	s.Events = v
	return s
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
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAutoRepairPolicyRequestRulesIncidentsConditions struct {
	// example:
	//
	// xxx
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// False
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// KubeletReady
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAutoRepairPolicyRequestRulesIncidentsConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyRequestRulesIncidentsConditions) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsConditions) GetReason() *string {
	return s.Reason
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsConditions) GetStatus() *string {
	return s.Status
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsConditions) GetType() *string {
	return s.Type
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsConditions) SetReason(v string) *CreateAutoRepairPolicyRequestRulesIncidentsConditions {
	s.Reason = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsConditions) SetStatus(v string) *CreateAutoRepairPolicyRequestRulesIncidentsConditions {
	s.Status = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsConditions) SetType(v string) *CreateAutoRepairPolicyRequestRulesIncidentsConditions {
	s.Type = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsConditions) Validate() error {
	return dara.Validate(s)
}

type CreateAutoRepairPolicyRequestRulesIncidentsEvents struct {
	// example:
	//
	// xxx
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// xxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAutoRepairPolicyRequestRulesIncidentsEvents) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoRepairPolicyRequestRulesIncidentsEvents) GoString() string {
	return s.String()
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsEvents) GetReason() *string {
	return s.Reason
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsEvents) GetType() *string {
	return s.Type
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsEvents) SetReason(v string) *CreateAutoRepairPolicyRequestRulesIncidentsEvents {
	s.Reason = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsEvents) SetType(v string) *CreateAutoRepairPolicyRequestRulesIncidentsEvents {
	s.Type = &v
	return s
}

func (s *CreateAutoRepairPolicyRequestRulesIncidentsEvents) Validate() error {
	return dara.Validate(s)
}

type CreateAutoRepairPolicyRequestRulesRepairProcedure struct {
	Config       map[string]interface{}                                         `json:"config,omitempty" xml:"config,omitempty"`
	Intervention *CreateAutoRepairPolicyRequestRulesRepairProcedureIntervention `json:"intervention,omitempty" xml:"intervention,omitempty" type:"Struct"`
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
	ApprovedLabel *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel `json:"approved_label,omitempty" xml:"approved_label,omitempty" type:"Struct"`
	// example:
	//
	// true
	Enable         *bool                                                                        `json:"enable,omitempty" xml:"enable,omitempty"`
	InquiringLabel *CreateAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel `json:"inquiring_label,omitempty" xml:"inquiring_label,omitempty" type:"Struct"`
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
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
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
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
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

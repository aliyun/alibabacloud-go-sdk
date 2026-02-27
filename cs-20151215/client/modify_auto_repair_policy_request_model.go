// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoRepairPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ModifyAutoRepairPolicyRequest
	GetName() *string
	SetRules(v []*ModifyAutoRepairPolicyRequestRules) *ModifyAutoRepairPolicyRequest
	GetRules() []*ModifyAutoRepairPolicyRequestRules
}

type ModifyAutoRepairPolicyRequest struct {
	// example:
	//
	// test
	Name  *string                               `json:"name,omitempty" xml:"name,omitempty"`
	Rules []*ModifyAutoRepairPolicyRequestRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s ModifyAutoRepairPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRepairPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoRepairPolicyRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAutoRepairPolicyRequest) GetRules() []*ModifyAutoRepairPolicyRequestRules {
	return s.Rules
}

func (s *ModifyAutoRepairPolicyRequest) SetName(v string) *ModifyAutoRepairPolicyRequest {
	s.Name = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequest) SetRules(v []*ModifyAutoRepairPolicyRequestRules) *ModifyAutoRepairPolicyRequest {
	s.Rules = v
	return s
}

func (s *ModifyAutoRepairPolicyRequest) Validate() error {
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

type ModifyAutoRepairPolicyRequestRules struct {
	Incidents       []*ModifyAutoRepairPolicyRequestRulesIncidents       `json:"incidents,omitempty" xml:"incidents,omitempty" type:"Repeated"`
	RepairProcedure []*ModifyAutoRepairPolicyRequestRulesRepairProcedure `json:"repair_procedure,omitempty" xml:"repair_procedure,omitempty" type:"Repeated"`
}

func (s ModifyAutoRepairPolicyRequestRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRepairPolicyRequestRules) GoString() string {
	return s.String()
}

func (s *ModifyAutoRepairPolicyRequestRules) GetIncidents() []*ModifyAutoRepairPolicyRequestRulesIncidents {
	return s.Incidents
}

func (s *ModifyAutoRepairPolicyRequestRules) GetRepairProcedure() []*ModifyAutoRepairPolicyRequestRulesRepairProcedure {
	return s.RepairProcedure
}

func (s *ModifyAutoRepairPolicyRequestRules) SetIncidents(v []*ModifyAutoRepairPolicyRequestRulesIncidents) *ModifyAutoRepairPolicyRequestRules {
	s.Incidents = v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRules) SetRepairProcedure(v []*ModifyAutoRepairPolicyRequestRulesRepairProcedure) *ModifyAutoRepairPolicyRequestRules {
	s.RepairProcedure = v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRules) Validate() error {
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

type ModifyAutoRepairPolicyRequestRulesIncidents struct {
	// example:
	//
	// Node.FaultNeedReboot.HOST
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// system
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyAutoRepairPolicyRequestRulesIncidents) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRepairPolicyRequestRulesIncidents) GoString() string {
	return s.String()
}

func (s *ModifyAutoRepairPolicyRequestRulesIncidents) GetName() *string {
	return s.Name
}

func (s *ModifyAutoRepairPolicyRequestRulesIncidents) GetType() *string {
	return s.Type
}

func (s *ModifyAutoRepairPolicyRequestRulesIncidents) SetName(v string) *ModifyAutoRepairPolicyRequestRulesIncidents {
	s.Name = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesIncidents) SetType(v string) *ModifyAutoRepairPolicyRequestRulesIncidents {
	s.Type = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesIncidents) Validate() error {
	return dara.Validate(s)
}

type ModifyAutoRepairPolicyRequestRulesRepairProcedure struct {
	Config       map[string]interface{}                                         `json:"config,omitempty" xml:"config,omitempty"`
	Intervention *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention `json:"intervention,omitempty" xml:"intervention,omitempty" type:"Struct"`
	// example:
	//
	// QuarantineGPU
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModifyAutoRepairPolicyRequestRulesRepairProcedure) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRepairPolicyRequestRulesRepairProcedure) GoString() string {
	return s.String()
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedure) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedure) GetIntervention() *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	return s.Intervention
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedure) GetName() *string {
	return s.Name
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedure) SetConfig(v map[string]interface{}) *ModifyAutoRepairPolicyRequestRulesRepairProcedure {
	s.Config = v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedure) SetIntervention(v *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) *ModifyAutoRepairPolicyRequestRulesRepairProcedure {
	s.Intervention = v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedure) SetName(v string) *ModifyAutoRepairPolicyRequestRulesRepairProcedure {
	s.Name = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedure) Validate() error {
	if s.Intervention != nil {
		if err := s.Intervention.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention struct {
	ApprovedLabel *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel `json:"approved_label,omitempty" xml:"approved_label,omitempty" type:"Struct"`
	// example:
	//
	// true
	Enable         *bool                                                                        `json:"enable,omitempty" xml:"enable,omitempty"`
	InquiringLabel *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel `json:"inquiring_label,omitempty" xml:"inquiring_label,omitempty" type:"Struct"`
	// example:
	//
	// label
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) GoString() string {
	return s.String()
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) GetApprovedLabel() *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel {
	return s.ApprovedLabel
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) GetInquiringLabel() *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel {
	return s.InquiringLabel
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) GetType() *string {
	return s.Type
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) SetApprovedLabel(v *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	s.ApprovedLabel = v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) SetEnable(v bool) *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	s.Enable = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) SetInquiringLabel(v *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	s.InquiringLabel = v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) SetType(v string) *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention {
	s.Type = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureIntervention) Validate() error {
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

type ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel struct {
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// approved
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) GoString() string {
	return s.String()
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) GetKey() *string {
	return s.Key
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) GetValue() *string {
	return s.Value
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) SetKey(v string) *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel {
	s.Key = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) SetValue(v string) *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel {
	s.Value = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionApprovedLabel) Validate() error {
	return dara.Validate(s)
}

type ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel struct {
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// inquiring
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) GoString() string {
	return s.String()
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) GetKey() *string {
	return s.Key
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) GetValue() *string {
	return s.Value
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) SetKey(v string) *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel {
	s.Key = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) SetValue(v string) *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel {
	s.Value = &v
	return s
}

func (s *ModifyAutoRepairPolicyRequestRulesRepairProcedureInterventionInquiringLabel) Validate() error {
	return dara.Validate(s)
}

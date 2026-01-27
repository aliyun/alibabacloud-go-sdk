// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoRepairPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribeAutoRepairPolicyResponseBody
	GetId() *string
	SetName(v string) *DescribeAutoRepairPolicyResponseBody
	GetName() *string
	SetResourceIds(v []*string) *DescribeAutoRepairPolicyResponseBody
	GetResourceIds() []*string
	SetResourceSubType(v string) *DescribeAutoRepairPolicyResponseBody
	GetResourceSubType() *string
	SetResourceType(v string) *DescribeAutoRepairPolicyResponseBody
	GetResourceType() *string
	SetRules(v []*DescribeAutoRepairPolicyResponseBodyRules) *DescribeAutoRepairPolicyResponseBody
	GetRules() []*DescribeAutoRepairPolicyResponseBodyRules
}

type DescribeAutoRepairPolicyResponseBody struct {
	// example:
	//
	// r-xxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test
	Name        *string   `json:"name,omitempty" xml:"name,omitempty"`
	ResourceIds []*string `json:"resource_ids,omitempty" xml:"resource_ids,omitempty" type:"Repeated"`
	// example:
	//
	// ess
	ResourceSubType *string `json:"resource_sub_type,omitempty" xml:"resource_sub_type,omitempty"`
	// example:
	//
	// nodepool
	ResourceType *string                                      `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	Rules        []*DescribeAutoRepairPolicyResponseBodyRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s DescribeAutoRepairPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribeAutoRepairPolicyResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeAutoRepairPolicyResponseBody) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeAutoRepairPolicyResponseBody) GetResourceSubType() *string {
	return s.ResourceSubType
}

func (s *DescribeAutoRepairPolicyResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeAutoRepairPolicyResponseBody) GetRules() []*DescribeAutoRepairPolicyResponseBodyRules {
	return s.Rules
}

func (s *DescribeAutoRepairPolicyResponseBody) SetId(v string) *DescribeAutoRepairPolicyResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBody) SetName(v string) *DescribeAutoRepairPolicyResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBody) SetResourceIds(v []*string) *DescribeAutoRepairPolicyResponseBody {
	s.ResourceIds = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBody) SetResourceSubType(v string) *DescribeAutoRepairPolicyResponseBody {
	s.ResourceSubType = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBody) SetResourceType(v string) *DescribeAutoRepairPolicyResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBody) SetRules(v []*DescribeAutoRepairPolicyResponseBodyRules) *DescribeAutoRepairPolicyResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBody) Validate() error {
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

type DescribeAutoRepairPolicyResponseBodyRules struct {
	Incidents       []*DescribeAutoRepairPolicyResponseBodyRulesIncidents       `json:"incidents,omitempty" xml:"incidents,omitempty" type:"Repeated"`
	RepairProcedure []*DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure `json:"repair_procedure,omitempty" xml:"repair_procedure,omitempty" type:"Repeated"`
}

func (s DescribeAutoRepairPolicyResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBodyRules) GetIncidents() []*DescribeAutoRepairPolicyResponseBodyRulesIncidents {
	return s.Incidents
}

func (s *DescribeAutoRepairPolicyResponseBodyRules) GetRepairProcedure() []*DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure {
	return s.RepairProcedure
}

func (s *DescribeAutoRepairPolicyResponseBodyRules) SetIncidents(v []*DescribeAutoRepairPolicyResponseBodyRulesIncidents) *DescribeAutoRepairPolicyResponseBodyRules {
	s.Incidents = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRules) SetRepairProcedure(v []*DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) *DescribeAutoRepairPolicyResponseBodyRules {
	s.RepairProcedure = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRules) Validate() error {
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

type DescribeAutoRepairPolicyResponseBodyRulesIncidents struct {
	Conditions []*DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Events     []*DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents     `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// example:
	//
	// Node.FaultNeedReboot.HOST
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// system
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAutoRepairPolicyResponseBodyRulesIncidents) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBodyRulesIncidents) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) GetConditions() []*DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions {
	return s.Conditions
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) GetEvents() []*DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents {
	return s.Events
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) GetName() *string {
	return s.Name
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) GetType() *string {
	return s.Type
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) SetConditions(v []*DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) *DescribeAutoRepairPolicyResponseBodyRulesIncidents {
	s.Conditions = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) SetEvents(v []*DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents) *DescribeAutoRepairPolicyResponseBodyRulesIncidents {
	s.Events = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) SetName(v string) *DescribeAutoRepairPolicyResponseBodyRulesIncidents {
	s.Name = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) SetType(v string) *DescribeAutoRepairPolicyResponseBodyRulesIncidents {
	s.Type = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) Validate() error {
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

type DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions struct {
	// example:
	//
	// xxxx
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

func (s DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) GetReason() *string {
	return s.Reason
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) GetType() *string {
	return s.Type
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) SetReason(v string) *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions {
	s.Reason = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) SetStatus(v string) *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions {
	s.Status = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) SetType(v string) *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions {
	s.Type = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsConditions) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents struct {
	// example:
	//
	// xxxx
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// xxxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents) GetReason() *string {
	return s.Reason
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents) GetType() *string {
	return s.Type
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents) SetReason(v string) *DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents {
	s.Reason = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents) SetType(v string) *DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents {
	s.Type = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidentsEvents) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure struct {
	Config       map[string]interface{}                                                `json:"config,omitempty" xml:"config,omitempty"`
	Intervention *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention `json:"intervention,omitempty" xml:"intervention,omitempty" type:"Struct"`
	// example:
	//
	// QuarantineGPU
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) GetIntervention() *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention {
	return s.Intervention
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) GetName() *string {
	return s.Name
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) SetConfig(v map[string]interface{}) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure {
	s.Config = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) SetIntervention(v *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure {
	s.Intervention = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) SetName(v string) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure {
	s.Name = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure) Validate() error {
	if s.Intervention != nil {
		if err := s.Intervention.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention struct {
	ApprovedLabel *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel `json:"approved_label,omitempty" xml:"approved_label,omitempty" type:"Struct"`
	// example:
	//
	// true
	Enable         *string                                                                             `json:"enable,omitempty" xml:"enable,omitempty"`
	InquiringLabel *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel `json:"inquiring_label,omitempty" xml:"inquiring_label,omitempty" type:"Struct"`
	// example:
	//
	// label
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) GetApprovedLabel() *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel {
	return s.ApprovedLabel
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) GetEnable() *string {
	return s.Enable
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) GetInquiringLabel() *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel {
	return s.InquiringLabel
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) GetType() *string {
	return s.Type
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) SetApprovedLabel(v *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention {
	s.ApprovedLabel = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) SetEnable(v string) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention {
	s.Enable = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) SetInquiringLabel(v *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention {
	s.InquiringLabel = v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) SetType(v string) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention {
	s.Type = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) Validate() error {
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

type DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel struct {
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// approved
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel) GetKey() *string {
	return s.Key
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel) GetValue() *string {
	return s.Value
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel) SetKey(v string) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel {
	s.Key = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel) SetValue(v string) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel {
	s.Value = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel struct {
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// inquiring
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel) GetKey() *string {
	return s.Key
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel) GetValue() *string {
	return s.Value
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel) SetKey(v string) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel {
	s.Key = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel) SetValue(v string) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel {
	s.Value = &v
	return s
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel) Validate() error {
	return dara.Validate(s)
}

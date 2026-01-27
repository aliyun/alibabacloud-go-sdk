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
	// example:
	//
	// r-xxxxx
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
	ResourceType *string                                         `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	Rules        []*ListAutoRepairPoliciesResponseBodyItemsRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
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
	Incidents       []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidents       `json:"incidents,omitempty" xml:"incidents,omitempty" type:"Repeated"`
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
	Conditions []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Events     []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents     `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// example:
	//
	// Node.FaultNeedReboot.HOST
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) GetConditions() []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions {
	return s.Conditions
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) GetEvents() []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents {
	return s.Events
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) GetName() *string {
	return s.Name
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) GetType() *string {
	return s.Type
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) SetConditions(v []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents {
	s.Conditions = v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents) SetEvents(v []*ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents) *ListAutoRepairPoliciesResponseBodyItemsRulesIncidents {
	s.Events = v
	return s
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

type ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions struct {
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

func (s ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) GetReason() *string {
	return s.Reason
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) GetStatus() *string {
	return s.Status
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) GetType() *string {
	return s.Type
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) SetReason(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions {
	s.Reason = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) SetStatus(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions {
	s.Status = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) SetType(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions {
	s.Type = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsConditions) Validate() error {
	return dara.Validate(s)
}

type ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents struct {
	// example:
	//
	// xxx
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// xxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents) String() string {
	return dara.Prettify(s)
}

func (s ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents) GoString() string {
	return s.String()
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents) GetReason() *string {
	return s.Reason
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents) GetType() *string {
	return s.Type
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents) SetReason(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents {
	s.Reason = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents) SetType(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents {
	s.Type = &v
	return s
}

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesIncidentsEvents) Validate() error {
	return dara.Validate(s)
}

type ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedure struct {
	Config       map[string]interface{}                                                   `json:"config,omitempty" xml:"config,omitempty"`
	Intervention *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention `json:"intervention,omitempty" xml:"intervention,omitempty" type:"Struct"`
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
	ApprovedLabel *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionApprovedLabel `json:"approved_label,omitempty" xml:"approved_label,omitempty" type:"Struct"`
	// example:
	//
	// true
	Enable         *string                                                                                `json:"enable,omitempty" xml:"enable,omitempty"`
	InquiringLabel *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureInterventionInquiringLabel `json:"inquiring_label,omitempty" xml:"inquiring_label,omitempty" type:"Struct"`
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

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) GetEnable() *string {
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

func (s *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention) SetEnable(v string) *ListAutoRepairPoliciesResponseBodyItemsRulesRepairProcedureIntervention {
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
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
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
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
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

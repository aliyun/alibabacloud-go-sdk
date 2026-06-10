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
	// The ID of the auto-repair rule.
	//
	// example:
	//
	// r-xxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the auto-repair rule.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The IDs of resources affected by the auto-repair rule.
	ResourceIds []*string `json:"resource_ids,omitempty" xml:"resource_ids,omitempty" type:"Repeated"`
	// The subtype of the resource affected by the auto-repair rule.
	//
	// example:
	//
	// ess
	ResourceSubType *string `json:"resource_sub_type,omitempty" xml:"resource_sub_type,omitempty"`
	// The resource type affected by the auto-repair rule.
	//
	// example:
	//
	// nodepool
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The list of rules.
	Rules []*DescribeAutoRepairPolicyResponseBodyRules `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
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
	// The detected incidents that trigger the rule.
	Incidents []*DescribeAutoRepairPolicyResponseBodyRulesIncidents `json:"incidents,omitempty" xml:"incidents,omitempty" type:"Repeated"`
	// The repair procedure, which contains a list of repair actions.
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

func (s DescribeAutoRepairPolicyResponseBodyRulesIncidents) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoRepairPolicyResponseBodyRulesIncidents) GoString() string {
	return s.String()
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) GetName() *string {
	return s.Name
}

func (s *DescribeAutoRepairPolicyResponseBodyRulesIncidents) GetType() *string {
	return s.Type
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
	return dara.Validate(s)
}

type DescribeAutoRepairPolicyResponseBodyRulesRepairProcedure struct {
	// The configuration parameters for the repair action.
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The manual approval configuration.
	Intervention *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention `json:"intervention,omitempty" xml:"intervention,omitempty" type:"Struct"`
	// The name of the repair action.
	//
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
	// The label you add to a node to approve a repair action. When ACK detects this label, it proceeds with the current repair step. After the action is complete, ACK automatically removes both the inquiring and approved labels. If you do not add this label promptly, the repair procedure is paused, and the node may remain unhealthy.
	ApprovedLabel *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionApprovedLabel `json:"approved_label,omitempty" xml:"approved_label,omitempty" type:"Struct"`
	// Specifies whether to enable manual approval.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// When a repair procedure reaches this step, ACK applies this label to the affected node and pauses until you grant approval.
	InquiringLabel *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureInterventionInquiringLabel `json:"inquiring_label,omitempty" xml:"inquiring_label,omitempty" type:"Struct"`
	// The manual approval type.
	//
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

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) GetEnable() *bool {
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

func (s *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention) SetEnable(v bool) *DescribeAutoRepairPolicyResponseBodyRulesRepairProcedureIntervention {
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
	// The label\\"s `key`.
	//
	// example:
	//
	// k8s.aliyun.com/incident
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The label\\"s value.
	//
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentEscalationPolicyForModify interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *IncidentEscalationPolicyForModify
	GetDescription() *string
	SetEnable(v bool) *IncidentEscalationPolicyForModify
	GetEnable() *bool
	SetEscalationStageList(v []*IncidentEscalationStageForView) *IncidentEscalationPolicyForModify
	GetEscalationStageList() []*IncidentEscalationStageForView
	SetName(v string) *IncidentEscalationPolicyForModify
	GetName() *string
}

type IncidentEscalationPolicyForModify struct {
	Description         *string                           `json:"description,omitempty" xml:"description,omitempty"`
	Enable              *bool                             `json:"enable,omitempty" xml:"enable,omitempty"`
	EscalationStageList []*IncidentEscalationStageForView `json:"escalationStageList,omitempty" xml:"escalationStageList,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s IncidentEscalationPolicyForModify) String() string {
	return dara.Prettify(s)
}

func (s IncidentEscalationPolicyForModify) GoString() string {
	return s.String()
}

func (s *IncidentEscalationPolicyForModify) GetDescription() *string {
	return s.Description
}

func (s *IncidentEscalationPolicyForModify) GetEnable() *bool {
	return s.Enable
}

func (s *IncidentEscalationPolicyForModify) GetEscalationStageList() []*IncidentEscalationStageForView {
	return s.EscalationStageList
}

func (s *IncidentEscalationPolicyForModify) GetName() *string {
	return s.Name
}

func (s *IncidentEscalationPolicyForModify) SetDescription(v string) *IncidentEscalationPolicyForModify {
	s.Description = &v
	return s
}

func (s *IncidentEscalationPolicyForModify) SetEnable(v bool) *IncidentEscalationPolicyForModify {
	s.Enable = &v
	return s
}

func (s *IncidentEscalationPolicyForModify) SetEscalationStageList(v []*IncidentEscalationStageForView) *IncidentEscalationPolicyForModify {
	s.EscalationStageList = v
	return s
}

func (s *IncidentEscalationPolicyForModify) SetName(v string) *IncidentEscalationPolicyForModify {
	s.Name = &v
	return s
}

func (s *IncidentEscalationPolicyForModify) Validate() error {
	if s.EscalationStageList != nil {
		for _, item := range s.EscalationStageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

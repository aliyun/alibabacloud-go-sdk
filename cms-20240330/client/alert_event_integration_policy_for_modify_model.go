// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertEventIntegrationPolicyForModify interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEventIntegrationPolicyName(v string) *AlertEventIntegrationPolicyForModify
	GetAlertEventIntegrationPolicyName() *string
	SetDescription(v string) *AlertEventIntegrationPolicyForModify
	GetDescription() *string
	SetFilterSetting(v *FilterSetting) *AlertEventIntegrationPolicyForModify
	GetFilterSetting() *FilterSetting
	SetIntegrationSetting(v string) *AlertEventIntegrationPolicyForModify
	GetIntegrationSetting() *string
	SetTransformerSetting(v []*TransformAction) *AlertEventIntegrationPolicyForModify
	GetTransformerSetting() []*TransformAction
	SetType(v string) *AlertEventIntegrationPolicyForModify
	GetType() *string
}

type AlertEventIntegrationPolicyForModify struct {
	// This parameter is required.
	AlertEventIntegrationPolicyName *string            `json:"alertEventIntegrationPolicyName,omitempty" xml:"alertEventIntegrationPolicyName,omitempty"`
	Description                     *string            `json:"description,omitempty" xml:"description,omitempty"`
	FilterSetting                   *FilterSetting     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	IntegrationSetting              *string            `json:"integrationSetting,omitempty" xml:"integrationSetting,omitempty"`
	TransformerSetting              []*TransformAction `json:"transformerSetting,omitempty" xml:"transformerSetting,omitempty" type:"Repeated"`
	Type                            *string            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AlertEventIntegrationPolicyForModify) String() string {
	return dara.Prettify(s)
}

func (s AlertEventIntegrationPolicyForModify) GoString() string {
	return s.String()
}

func (s *AlertEventIntegrationPolicyForModify) GetAlertEventIntegrationPolicyName() *string {
	return s.AlertEventIntegrationPolicyName
}

func (s *AlertEventIntegrationPolicyForModify) GetDescription() *string {
	return s.Description
}

func (s *AlertEventIntegrationPolicyForModify) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *AlertEventIntegrationPolicyForModify) GetIntegrationSetting() *string {
	return s.IntegrationSetting
}

func (s *AlertEventIntegrationPolicyForModify) GetTransformerSetting() []*TransformAction {
	return s.TransformerSetting
}

func (s *AlertEventIntegrationPolicyForModify) GetType() *string {
	return s.Type
}

func (s *AlertEventIntegrationPolicyForModify) SetAlertEventIntegrationPolicyName(v string) *AlertEventIntegrationPolicyForModify {
	s.AlertEventIntegrationPolicyName = &v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetDescription(v string) *AlertEventIntegrationPolicyForModify {
	s.Description = &v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetFilterSetting(v *FilterSetting) *AlertEventIntegrationPolicyForModify {
	s.FilterSetting = v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetIntegrationSetting(v string) *AlertEventIntegrationPolicyForModify {
	s.IntegrationSetting = &v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetTransformerSetting(v []*TransformAction) *AlertEventIntegrationPolicyForModify {
	s.TransformerSetting = v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) SetType(v string) *AlertEventIntegrationPolicyForModify {
	s.Type = &v
	return s
}

func (s *AlertEventIntegrationPolicyForModify) Validate() error {
	if s.FilterSetting != nil {
		if err := s.FilterSetting.Validate(); err != nil {
			return err
		}
	}
	if s.TransformerSetting != nil {
		for _, item := range s.TransformerSetting {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

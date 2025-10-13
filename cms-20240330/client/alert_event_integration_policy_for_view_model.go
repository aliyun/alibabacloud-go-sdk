// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertEventIntegrationPolicyForView interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEventIntegrationPolicyId(v string) *AlertEventIntegrationPolicyForView
	GetAlertEventIntegrationPolicyId() *string
	SetAlertEventIntegrationPolicyName(v string) *AlertEventIntegrationPolicyForView
	GetAlertEventIntegrationPolicyName() *string
	SetCreateTime(v string) *AlertEventIntegrationPolicyForView
	GetCreateTime() *string
	SetDescription(v string) *AlertEventIntegrationPolicyForView
	GetDescription() *string
	SetEnable(v bool) *AlertEventIntegrationPolicyForView
	GetEnable() *bool
	SetFilterSetting(v *FilterSetting) *AlertEventIntegrationPolicyForView
	GetFilterSetting() *FilterSetting
	SetIntegrationSetting(v string) *AlertEventIntegrationPolicyForView
	GetIntegrationSetting() *string
	SetToken(v string) *AlertEventIntegrationPolicyForView
	GetToken() *string
	SetTransformerSetting(v []*TransformAction) *AlertEventIntegrationPolicyForView
	GetTransformerSetting() []*TransformAction
	SetType(v string) *AlertEventIntegrationPolicyForView
	GetType() *string
	SetUpdateTime(v string) *AlertEventIntegrationPolicyForView
	GetUpdateTime() *string
	SetUserId(v string) *AlertEventIntegrationPolicyForView
	GetUserId() *string
	SetWorkspace(v string) *AlertEventIntegrationPolicyForView
	GetWorkspace() *string
}

type AlertEventIntegrationPolicyForView struct {
	AlertEventIntegrationPolicyId *string `json:"alertEventIntegrationPolicyId,omitempty" xml:"alertEventIntegrationPolicyId,omitempty"`
	// This parameter is required.
	AlertEventIntegrationPolicyName *string            `json:"alertEventIntegrationPolicyName,omitempty" xml:"alertEventIntegrationPolicyName,omitempty"`
	CreateTime                      *string            `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description                     *string            `json:"description,omitempty" xml:"description,omitempty"`
	Enable                          *bool              `json:"enable,omitempty" xml:"enable,omitempty"`
	FilterSetting                   *FilterSetting     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	IntegrationSetting              *string            `json:"integrationSetting,omitempty" xml:"integrationSetting,omitempty"`
	Token                           *string            `json:"token,omitempty" xml:"token,omitempty"`
	TransformerSetting              []*TransformAction `json:"transformerSetting,omitempty" xml:"transformerSetting,omitempty" type:"Repeated"`
	Type                            *string            `json:"type,omitempty" xml:"type,omitempty"`
	UpdateTime                      *string            `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserId                          *string            `json:"userId,omitempty" xml:"userId,omitempty"`
	Workspace                       *string            `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s AlertEventIntegrationPolicyForView) String() string {
	return dara.Prettify(s)
}

func (s AlertEventIntegrationPolicyForView) GoString() string {
	return s.String()
}

func (s *AlertEventIntegrationPolicyForView) GetAlertEventIntegrationPolicyId() *string {
	return s.AlertEventIntegrationPolicyId
}

func (s *AlertEventIntegrationPolicyForView) GetAlertEventIntegrationPolicyName() *string {
	return s.AlertEventIntegrationPolicyName
}

func (s *AlertEventIntegrationPolicyForView) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AlertEventIntegrationPolicyForView) GetDescription() *string {
	return s.Description
}

func (s *AlertEventIntegrationPolicyForView) GetEnable() *bool {
	return s.Enable
}

func (s *AlertEventIntegrationPolicyForView) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *AlertEventIntegrationPolicyForView) GetIntegrationSetting() *string {
	return s.IntegrationSetting
}

func (s *AlertEventIntegrationPolicyForView) GetToken() *string {
	return s.Token
}

func (s *AlertEventIntegrationPolicyForView) GetTransformerSetting() []*TransformAction {
	return s.TransformerSetting
}

func (s *AlertEventIntegrationPolicyForView) GetType() *string {
	return s.Type
}

func (s *AlertEventIntegrationPolicyForView) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *AlertEventIntegrationPolicyForView) GetUserId() *string {
	return s.UserId
}

func (s *AlertEventIntegrationPolicyForView) GetWorkspace() *string {
	return s.Workspace
}

func (s *AlertEventIntegrationPolicyForView) SetAlertEventIntegrationPolicyId(v string) *AlertEventIntegrationPolicyForView {
	s.AlertEventIntegrationPolicyId = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetAlertEventIntegrationPolicyName(v string) *AlertEventIntegrationPolicyForView {
	s.AlertEventIntegrationPolicyName = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetCreateTime(v string) *AlertEventIntegrationPolicyForView {
	s.CreateTime = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetDescription(v string) *AlertEventIntegrationPolicyForView {
	s.Description = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetEnable(v bool) *AlertEventIntegrationPolicyForView {
	s.Enable = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetFilterSetting(v *FilterSetting) *AlertEventIntegrationPolicyForView {
	s.FilterSetting = v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetIntegrationSetting(v string) *AlertEventIntegrationPolicyForView {
	s.IntegrationSetting = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetToken(v string) *AlertEventIntegrationPolicyForView {
	s.Token = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetTransformerSetting(v []*TransformAction) *AlertEventIntegrationPolicyForView {
	s.TransformerSetting = v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetType(v string) *AlertEventIntegrationPolicyForView {
	s.Type = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetUpdateTime(v string) *AlertEventIntegrationPolicyForView {
	s.UpdateTime = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetUserId(v string) *AlertEventIntegrationPolicyForView {
	s.UserId = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) SetWorkspace(v string) *AlertEventIntegrationPolicyForView {
	s.Workspace = &v
	return s
}

func (s *AlertEventIntegrationPolicyForView) Validate() error {
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

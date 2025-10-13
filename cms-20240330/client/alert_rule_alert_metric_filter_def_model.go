// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleAlertMetricFilterDef interface {
	dara.Model
	String() string
	GoString() string
	SetDim(v string) *AlertRuleAlertMetricFilterDef
	GetDim() *string
	SetDimDisabled(v bool) *AlertRuleAlertMetricFilterDef
	GetDimDisabled() *bool
	SetDisplayNameCn(v string) *AlertRuleAlertMetricFilterDef
	GetDisplayNameCn() *string
	SetDisplayNameEn(v string) *AlertRuleAlertMetricFilterDef
	GetDisplayNameEn() *string
	SetHidden(v bool) *AlertRuleAlertMetricFilterDef
	GetHidden() *bool
	SetLabelDisabled(v bool) *AlertRuleAlertMetricFilterDef
	GetLabelDisabled() *bool
	SetOpt(v string) *AlertRuleAlertMetricFilterDef
	GetOpt() *string
	SetSupportedOpts(v []*AlertRuleAlertMetricFilterDefSupportedOpts) *AlertRuleAlertMetricFilterDef
	GetSupportedOpts() []*AlertRuleAlertMetricFilterDefSupportedOpts
}

type AlertRuleAlertMetricFilterDef struct {
	Dim           *string                                       `json:"dim,omitempty" xml:"dim,omitempty"`
	DimDisabled   *bool                                         `json:"dimDisabled,omitempty" xml:"dimDisabled,omitempty"`
	DisplayNameCn *string                                       `json:"displayNameCn,omitempty" xml:"displayNameCn,omitempty"`
	DisplayNameEn *string                                       `json:"displayNameEn,omitempty" xml:"displayNameEn,omitempty"`
	Hidden        *bool                                         `json:"hidden,omitempty" xml:"hidden,omitempty"`
	LabelDisabled *bool                                         `json:"labelDisabled,omitempty" xml:"labelDisabled,omitempty"`
	Opt           *string                                       `json:"opt,omitempty" xml:"opt,omitempty"`
	SupportedOpts []*AlertRuleAlertMetricFilterDefSupportedOpts `json:"supportedOpts,omitempty" xml:"supportedOpts,omitempty" type:"Repeated"`
}

func (s AlertRuleAlertMetricFilterDef) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAlertMetricFilterDef) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricFilterDef) GetDim() *string {
	return s.Dim
}

func (s *AlertRuleAlertMetricFilterDef) GetDimDisabled() *bool {
	return s.DimDisabled
}

func (s *AlertRuleAlertMetricFilterDef) GetDisplayNameCn() *string {
	return s.DisplayNameCn
}

func (s *AlertRuleAlertMetricFilterDef) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *AlertRuleAlertMetricFilterDef) GetHidden() *bool {
	return s.Hidden
}

func (s *AlertRuleAlertMetricFilterDef) GetLabelDisabled() *bool {
	return s.LabelDisabled
}

func (s *AlertRuleAlertMetricFilterDef) GetOpt() *string {
	return s.Opt
}

func (s *AlertRuleAlertMetricFilterDef) GetSupportedOpts() []*AlertRuleAlertMetricFilterDefSupportedOpts {
	return s.SupportedOpts
}

func (s *AlertRuleAlertMetricFilterDef) SetDim(v string) *AlertRuleAlertMetricFilterDef {
	s.Dim = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetDimDisabled(v bool) *AlertRuleAlertMetricFilterDef {
	s.DimDisabled = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetDisplayNameCn(v string) *AlertRuleAlertMetricFilterDef {
	s.DisplayNameCn = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetDisplayNameEn(v string) *AlertRuleAlertMetricFilterDef {
	s.DisplayNameEn = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetHidden(v bool) *AlertRuleAlertMetricFilterDef {
	s.Hidden = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetLabelDisabled(v bool) *AlertRuleAlertMetricFilterDef {
	s.LabelDisabled = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetOpt(v string) *AlertRuleAlertMetricFilterDef {
	s.Opt = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) SetSupportedOpts(v []*AlertRuleAlertMetricFilterDefSupportedOpts) *AlertRuleAlertMetricFilterDef {
	s.SupportedOpts = v
	return s
}

func (s *AlertRuleAlertMetricFilterDef) Validate() error {
	if s.SupportedOpts != nil {
		for _, item := range s.SupportedOpts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleAlertMetricFilterDefSupportedOpts struct {
	DisplayNameCn *string `json:"displayNameCn,omitempty" xml:"displayNameCn,omitempty"`
	DisplayNameEn *string `json:"displayNameEn,omitempty" xml:"displayNameEn,omitempty"`
	Value         *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleAlertMetricFilterDefSupportedOpts) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAlertMetricFilterDefSupportedOpts) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) GetDisplayNameCn() *string {
	return s.DisplayNameCn
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) GetDisplayNameEn() *string {
	return s.DisplayNameEn
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) GetValue() *string {
	return s.Value
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) SetDisplayNameCn(v string) *AlertRuleAlertMetricFilterDefSupportedOpts {
	s.DisplayNameCn = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) SetDisplayNameEn(v string) *AlertRuleAlertMetricFilterDefSupportedOpts {
	s.DisplayNameEn = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) SetValue(v string) *AlertRuleAlertMetricFilterDefSupportedOpts {
	s.Value = &v
	return s
}

func (s *AlertRuleAlertMetricFilterDefSupportedOpts) Validate() error {
	return dara.Validate(s)
}

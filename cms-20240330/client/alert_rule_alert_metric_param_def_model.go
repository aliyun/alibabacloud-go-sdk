// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleAlertMetricParamDef interface {
	dara.Model
	String() string
	GoString() string
	SetMaxWidth(v int32) *AlertRuleAlertMetricParamDef
	GetMaxWidth() *int32
	SetMinWidth(v int32) *AlertRuleAlertMetricParamDef
	GetMinWidth() *int32
	SetName(v string) *AlertRuleAlertMetricParamDef
	GetName() *string
	SetPlaceholderCn(v string) *AlertRuleAlertMetricParamDef
	GetPlaceholderCn() *string
	SetPlaceholderEn(v string) *AlertRuleAlertMetricParamDef
	GetPlaceholderEn() *string
	SetType(v string) *AlertRuleAlertMetricParamDef
	GetType() *string
	SetValue(v string) *AlertRuleAlertMetricParamDef
	GetValue() *string
	SetValues(v []*AlertRuleAlertMetricParamDefValues) *AlertRuleAlertMetricParamDef
	GetValues() []*AlertRuleAlertMetricParamDefValues
}

type AlertRuleAlertMetricParamDef struct {
	MaxWidth      *int32                                `json:"maxWidth,omitempty" xml:"maxWidth,omitempty"`
	MinWidth      *int32                                `json:"minWidth,omitempty" xml:"minWidth,omitempty"`
	Name          *string                               `json:"name,omitempty" xml:"name,omitempty"`
	PlaceholderCn *string                               `json:"placeholderCn,omitempty" xml:"placeholderCn,omitempty"`
	PlaceholderEn *string                               `json:"placeholderEn,omitempty" xml:"placeholderEn,omitempty"`
	Type          *string                               `json:"type,omitempty" xml:"type,omitempty"`
	Value         *string                               `json:"value,omitempty" xml:"value,omitempty"`
	Values        []*AlertRuleAlertMetricParamDefValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s AlertRuleAlertMetricParamDef) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAlertMetricParamDef) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricParamDef) GetMaxWidth() *int32 {
	return s.MaxWidth
}

func (s *AlertRuleAlertMetricParamDef) GetMinWidth() *int32 {
	return s.MinWidth
}

func (s *AlertRuleAlertMetricParamDef) GetName() *string {
	return s.Name
}

func (s *AlertRuleAlertMetricParamDef) GetPlaceholderCn() *string {
	return s.PlaceholderCn
}

func (s *AlertRuleAlertMetricParamDef) GetPlaceholderEn() *string {
	return s.PlaceholderEn
}

func (s *AlertRuleAlertMetricParamDef) GetType() *string {
	return s.Type
}

func (s *AlertRuleAlertMetricParamDef) GetValue() *string {
	return s.Value
}

func (s *AlertRuleAlertMetricParamDef) GetValues() []*AlertRuleAlertMetricParamDefValues {
	return s.Values
}

func (s *AlertRuleAlertMetricParamDef) SetMaxWidth(v int32) *AlertRuleAlertMetricParamDef {
	s.MaxWidth = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetMinWidth(v int32) *AlertRuleAlertMetricParamDef {
	s.MinWidth = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetName(v string) *AlertRuleAlertMetricParamDef {
	s.Name = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetPlaceholderCn(v string) *AlertRuleAlertMetricParamDef {
	s.PlaceholderCn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetPlaceholderEn(v string) *AlertRuleAlertMetricParamDef {
	s.PlaceholderEn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetType(v string) *AlertRuleAlertMetricParamDef {
	s.Type = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetValue(v string) *AlertRuleAlertMetricParamDef {
	s.Value = &v
	return s
}

func (s *AlertRuleAlertMetricParamDef) SetValues(v []*AlertRuleAlertMetricParamDefValues) *AlertRuleAlertMetricParamDef {
	s.Values = v
	return s
}

func (s *AlertRuleAlertMetricParamDef) Validate() error {
	return dara.Validate(s)
}

type AlertRuleAlertMetricParamDefValues struct {
	LabelCn *string `json:"labelCn,omitempty" xml:"labelCn,omitempty"`
	LabelEn *string `json:"labelEn,omitempty" xml:"labelEn,omitempty"`
	Value   *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AlertRuleAlertMetricParamDefValues) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleAlertMetricParamDefValues) GoString() string {
	return s.String()
}

func (s *AlertRuleAlertMetricParamDefValues) GetLabelCn() *string {
	return s.LabelCn
}

func (s *AlertRuleAlertMetricParamDefValues) GetLabelEn() *string {
	return s.LabelEn
}

func (s *AlertRuleAlertMetricParamDefValues) GetValue() *string {
	return s.Value
}

func (s *AlertRuleAlertMetricParamDefValues) SetLabelCn(v string) *AlertRuleAlertMetricParamDefValues {
	s.LabelCn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDefValues) SetLabelEn(v string) *AlertRuleAlertMetricParamDefValues {
	s.LabelEn = &v
	return s
}

func (s *AlertRuleAlertMetricParamDefValues) SetValue(v string) *AlertRuleAlertMetricParamDefValues {
	s.Value = &v
	return s
}

func (s *AlertRuleAlertMetricParamDefValues) Validate() error {
	return dara.Validate(s)
}

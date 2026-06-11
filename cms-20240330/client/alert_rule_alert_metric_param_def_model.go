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
	// The maximum width of the input control. This parameter is valid only for SELECT_PARAM and INPUT_PARAM.
	//
	// example:
	//
	// 200
	MaxWidth *int32 `json:"maxWidth,omitempty" xml:"maxWidth,omitempty"`
	// The minimum width of the input control. This parameter is valid only for SELECT_PARAM and INPUT_PARAM.
	//
	// example:
	//
	// 100
	MinWidth *int32 `json:"minWidth,omitempty" xml:"minWidth,omitempty"`
	// The name.
	//
	// example:
	//
	// env
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Chinese placeholder text displayed on the frontend. This parameter is valid only for INPUT_PARAM.
	//
	// example:
	//
	// 请输入值
	PlaceholderCn *string `json:"placeholderCn,omitempty" xml:"placeholderCn,omitempty"`
	// The English placeholder text displayed on the frontend. This parameter is valid only for INPUT_PARAM.
	//
	// example:
	//
	// Enter value
	PlaceholderEn *string `json:"placeholderEn,omitempty" xml:"placeholderEn,omitempty"`
	// ● TEXT_PARAM: A read-only text parameter defined by the backend. No user input control is displayed on the frontend.● INPUT_PARAM: An input box parameter.● SELECT_PARAM: A selection box parameter.
	//
	// example:
	//
	// TEXT_PARAM
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value.
	//
	// example:
	//
	// staging
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// The list of available options in the drop-down list. This parameter is valid only for SELECT_PARAM.
	Values []*AlertRuleAlertMetricParamDefValues `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
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
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AlertRuleAlertMetricParamDefValues struct {
	// The Chinese display name of the option.
	//
	// example:
	//
	// 测试环境
	LabelCn *string `json:"labelCn,omitempty" xml:"labelCn,omitempty"`
	// The English display name of the option.
	//
	// example:
	//
	// Staging
	LabelEn *string `json:"labelEn,omitempty" xml:"labelEn,omitempty"`
	// The value.
	//
	// example:
	//
	// staging
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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

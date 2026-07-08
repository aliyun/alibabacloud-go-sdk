// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWritingStyleTemplateField interface {
	dara.Model
	String() string
	GoString() string
	SetBuildIn(v bool) *WritingStyleTemplateField
	GetBuildIn() *bool
	SetCascadingFields(v []*WritingStyleTemplateField) *WritingStyleTemplateField
	GetCascadingFields() []*WritingStyleTemplateField
	SetEnums(v []*WritingStyleTemplateFieldEnums) *WritingStyleTemplateField
	GetEnums() []*WritingStyleTemplateFieldEnums
	SetInitialValue(v string) *WritingStyleTemplateField
	GetInitialValue() *string
	SetKey(v string) *WritingStyleTemplateField
	GetKey() *string
	SetMax(v float64) *WritingStyleTemplateField
	GetMax() *float64
	SetMaxItem(v int32) *WritingStyleTemplateField
	GetMaxItem() *int32
	SetMaxItemLength(v int32) *WritingStyleTemplateField
	GetMaxItemLength() *int32
	SetMaxLength(v int32) *WritingStyleTemplateField
	GetMaxLength() *int32
	SetMin(v float64) *WritingStyleTemplateField
	GetMin() *float64
	SetMinItemLength(v int32) *WritingStyleTemplateField
	GetMinItemLength() *int32
	SetMinLength(v int32) *WritingStyleTemplateField
	GetMinLength() *int32
	SetName(v string) *WritingStyleTemplateField
	GetName() *string
	SetRequired(v bool) *WritingStyleTemplateField
	GetRequired() *bool
	SetStyle(v *WritingStyleTemplateFieldStyle) *WritingStyleTemplateField
	GetStyle() *WritingStyleTemplateFieldStyle
}

type WritingStyleTemplateField struct {
	BuildIn         *bool                             `json:"BuildIn,omitempty" xml:"BuildIn,omitempty"`
	CascadingFields []*WritingStyleTemplateField      `json:"CascadingFields,omitempty" xml:"CascadingFields,omitempty" type:"Repeated"`
	Enums           []*WritingStyleTemplateFieldEnums `json:"Enums,omitempty" xml:"Enums,omitempty" type:"Repeated"`
	InitialValue    *string                           `json:"InitialValue,omitempty" xml:"InitialValue,omitempty"`
	Key             *string                           `json:"Key,omitempty" xml:"Key,omitempty"`
	Max             *float64                          `json:"Max,omitempty" xml:"Max,omitempty"`
	MaxItem         *int32                            `json:"MaxItem,omitempty" xml:"MaxItem,omitempty"`
	MaxItemLength   *int32                            `json:"MaxItemLength,omitempty" xml:"MaxItemLength,omitempty"`
	MaxLength       *int32                            `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	Min             *float64                          `json:"Min,omitempty" xml:"Min,omitempty"`
	MinItemLength   *int32                            `json:"MinItemLength,omitempty" xml:"MinItemLength,omitempty"`
	MinLength       *int32                            `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	Name            *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Required        *bool                             `json:"Required,omitempty" xml:"Required,omitempty"`
	Style           *WritingStyleTemplateFieldStyle   `json:"Style,omitempty" xml:"Style,omitempty" type:"Struct"`
}

func (s WritingStyleTemplateField) String() string {
	return dara.Prettify(s)
}

func (s WritingStyleTemplateField) GoString() string {
	return s.String()
}

func (s *WritingStyleTemplateField) GetBuildIn() *bool {
	return s.BuildIn
}

func (s *WritingStyleTemplateField) GetCascadingFields() []*WritingStyleTemplateField {
	return s.CascadingFields
}

func (s *WritingStyleTemplateField) GetEnums() []*WritingStyleTemplateFieldEnums {
	return s.Enums
}

func (s *WritingStyleTemplateField) GetInitialValue() *string {
	return s.InitialValue
}

func (s *WritingStyleTemplateField) GetKey() *string {
	return s.Key
}

func (s *WritingStyleTemplateField) GetMax() *float64 {
	return s.Max
}

func (s *WritingStyleTemplateField) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *WritingStyleTemplateField) GetMaxItemLength() *int32 {
	return s.MaxItemLength
}

func (s *WritingStyleTemplateField) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *WritingStyleTemplateField) GetMin() *float64 {
	return s.Min
}

func (s *WritingStyleTemplateField) GetMinItemLength() *int32 {
	return s.MinItemLength
}

func (s *WritingStyleTemplateField) GetMinLength() *int32 {
	return s.MinLength
}

func (s *WritingStyleTemplateField) GetName() *string {
	return s.Name
}

func (s *WritingStyleTemplateField) GetRequired() *bool {
	return s.Required
}

func (s *WritingStyleTemplateField) GetStyle() *WritingStyleTemplateFieldStyle {
	return s.Style
}

func (s *WritingStyleTemplateField) SetBuildIn(v bool) *WritingStyleTemplateField {
	s.BuildIn = &v
	return s
}

func (s *WritingStyleTemplateField) SetCascadingFields(v []*WritingStyleTemplateField) *WritingStyleTemplateField {
	s.CascadingFields = v
	return s
}

func (s *WritingStyleTemplateField) SetEnums(v []*WritingStyleTemplateFieldEnums) *WritingStyleTemplateField {
	s.Enums = v
	return s
}

func (s *WritingStyleTemplateField) SetInitialValue(v string) *WritingStyleTemplateField {
	s.InitialValue = &v
	return s
}

func (s *WritingStyleTemplateField) SetKey(v string) *WritingStyleTemplateField {
	s.Key = &v
	return s
}

func (s *WritingStyleTemplateField) SetMax(v float64) *WritingStyleTemplateField {
	s.Max = &v
	return s
}

func (s *WritingStyleTemplateField) SetMaxItem(v int32) *WritingStyleTemplateField {
	s.MaxItem = &v
	return s
}

func (s *WritingStyleTemplateField) SetMaxItemLength(v int32) *WritingStyleTemplateField {
	s.MaxItemLength = &v
	return s
}

func (s *WritingStyleTemplateField) SetMaxLength(v int32) *WritingStyleTemplateField {
	s.MaxLength = &v
	return s
}

func (s *WritingStyleTemplateField) SetMin(v float64) *WritingStyleTemplateField {
	s.Min = &v
	return s
}

func (s *WritingStyleTemplateField) SetMinItemLength(v int32) *WritingStyleTemplateField {
	s.MinItemLength = &v
	return s
}

func (s *WritingStyleTemplateField) SetMinLength(v int32) *WritingStyleTemplateField {
	s.MinLength = &v
	return s
}

func (s *WritingStyleTemplateField) SetName(v string) *WritingStyleTemplateField {
	s.Name = &v
	return s
}

func (s *WritingStyleTemplateField) SetRequired(v bool) *WritingStyleTemplateField {
	s.Required = &v
	return s
}

func (s *WritingStyleTemplateField) SetStyle(v *WritingStyleTemplateFieldStyle) *WritingStyleTemplateField {
	s.Style = v
	return s
}

func (s *WritingStyleTemplateField) Validate() error {
	if s.CascadingFields != nil {
		for _, item := range s.CascadingFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Enums != nil {
		for _, item := range s.Enums {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Style != nil {
		if err := s.Style.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WritingStyleTemplateFieldEnums struct {
	CascadingFields []*string `json:"CascadingFields,omitempty" xml:"CascadingFields,omitempty" type:"Repeated"`
	Key             *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s WritingStyleTemplateFieldEnums) String() string {
	return dara.Prettify(s)
}

func (s WritingStyleTemplateFieldEnums) GoString() string {
	return s.String()
}

func (s *WritingStyleTemplateFieldEnums) GetCascadingFields() []*string {
	return s.CascadingFields
}

func (s *WritingStyleTemplateFieldEnums) GetKey() *string {
	return s.Key
}

func (s *WritingStyleTemplateFieldEnums) GetName() *string {
	return s.Name
}

func (s *WritingStyleTemplateFieldEnums) SetCascadingFields(v []*string) *WritingStyleTemplateFieldEnums {
	s.CascadingFields = v
	return s
}

func (s *WritingStyleTemplateFieldEnums) SetKey(v string) *WritingStyleTemplateFieldEnums {
	s.Key = &v
	return s
}

func (s *WritingStyleTemplateFieldEnums) SetName(v string) *WritingStyleTemplateFieldEnums {
	s.Name = &v
	return s
}

func (s *WritingStyleTemplateFieldEnums) Validate() error {
	return dara.Validate(s)
}

type WritingStyleTemplateFieldStyle struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Format      *string `json:"Format,omitempty" xml:"Format,omitempty"`
	Placeholder *string `json:"Placeholder,omitempty" xml:"Placeholder,omitempty"`
	ShowTime    *bool   `json:"ShowTime,omitempty" xml:"ShowTime,omitempty"`
	Suffix      *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s WritingStyleTemplateFieldStyle) String() string {
	return dara.Prettify(s)
}

func (s WritingStyleTemplateFieldStyle) GoString() string {
	return s.String()
}

func (s *WritingStyleTemplateFieldStyle) GetDescription() *string {
	return s.Description
}

func (s *WritingStyleTemplateFieldStyle) GetFormat() *string {
	return s.Format
}

func (s *WritingStyleTemplateFieldStyle) GetPlaceholder() *string {
	return s.Placeholder
}

func (s *WritingStyleTemplateFieldStyle) GetShowTime() *bool {
	return s.ShowTime
}

func (s *WritingStyleTemplateFieldStyle) GetSuffix() *string {
	return s.Suffix
}

func (s *WritingStyleTemplateFieldStyle) GetType() *string {
	return s.Type
}

func (s *WritingStyleTemplateFieldStyle) SetDescription(v string) *WritingStyleTemplateFieldStyle {
	s.Description = &v
	return s
}

func (s *WritingStyleTemplateFieldStyle) SetFormat(v string) *WritingStyleTemplateFieldStyle {
	s.Format = &v
	return s
}

func (s *WritingStyleTemplateFieldStyle) SetPlaceholder(v string) *WritingStyleTemplateFieldStyle {
	s.Placeholder = &v
	return s
}

func (s *WritingStyleTemplateFieldStyle) SetShowTime(v bool) *WritingStyleTemplateFieldStyle {
	s.ShowTime = &v
	return s
}

func (s *WritingStyleTemplateFieldStyle) SetSuffix(v string) *WritingStyleTemplateFieldStyle {
	s.Suffix = &v
	return s
}

func (s *WritingStyleTemplateFieldStyle) SetType(v string) *WritingStyleTemplateFieldStyle {
	s.Type = &v
	return s
}

func (s *WritingStyleTemplateFieldStyle) Validate() error {
	return dara.Validate(s)
}

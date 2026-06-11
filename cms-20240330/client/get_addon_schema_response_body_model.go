// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFields(v []*GetAddonSchemaResponseBodyFields) *GetAddonSchemaResponseBody
	GetFields() []*GetAddonSchemaResponseBodyFields
	SetRequestId(v string) *GetAddonSchemaResponseBody
	GetRequestId() *string
	SetType(v string) *GetAddonSchemaResponseBody
	GetType() *string
}

type GetAddonSchemaResponseBody struct {
	// A list of data table fields.
	Fields []*GetAddonSchemaResponseBodyFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// E5B1D3D4-BB28-5996-8AD2-***********
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The type of the schema.
	//
	// example:
	//
	// common
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetAddonSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAddonSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddonSchemaResponseBody) GetFields() []*GetAddonSchemaResponseBodyFields {
	return s.Fields
}

func (s *GetAddonSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAddonSchemaResponseBody) GetType() *string {
	return s.Type
}

func (s *GetAddonSchemaResponseBody) SetFields(v []*GetAddonSchemaResponseBodyFields) *GetAddonSchemaResponseBody {
	s.Fields = v
	return s
}

func (s *GetAddonSchemaResponseBody) SetRequestId(v string) *GetAddonSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddonSchemaResponseBody) SetType(v string) *GetAddonSchemaResponseBody {
	s.Type = &v
	return s
}

func (s *GetAddonSchemaResponseBody) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAddonSchemaResponseBodyFields struct {
	// A list of display conditions.
	Conditions []*GetAddonSchemaResponseBodyFieldsConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The default value.
	//
	// example:
	//
	// mysql.default
	DefaultValue interface{} `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// The description.
	//
	// example:
	//
	// MySQL datasource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether the field is hidden. Valid values: \\`true\\` (hidden) and \\`false\\` (displayed).
	//
	// example:
	//
	// false
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// The type of the element.
	//
	// example:
	//
	// input
	Element *string `json:"element,omitempty" xml:"element,omitempty"`
	// The path of the field.
	//
	// example:
	//
	// mysql.host
	FieldPath *string `json:"fieldPath,omitempty" xml:"fieldPath,omitempty"`
	// The display name of the field.
	//
	// example:
	//
	// mysql
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// host
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The placeholder text.
	//
	// example:
	//
	// Please enter the metrics collection interval
	Placeholder *string `json:"placeholder,omitempty" xml:"placeholder,omitempty"`
	// Other properties of the component.
	Props *GetAddonSchemaResponseBodyFieldsProps `json:"props,omitempty" xml:"props,omitempty" type:"Struct"`
	// The type of the field.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The validation rule for the field.
	Validation *GetAddonSchemaResponseBodyFieldsValidation `json:"validation,omitempty" xml:"validation,omitempty" type:"Struct"`
}

func (s GetAddonSchemaResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s GetAddonSchemaResponseBodyFields) GoString() string {
	return s.String()
}

func (s *GetAddonSchemaResponseBodyFields) GetConditions() []*GetAddonSchemaResponseBodyFieldsConditions {
	return s.Conditions
}

func (s *GetAddonSchemaResponseBodyFields) GetDefaultValue() interface{} {
	return s.DefaultValue
}

func (s *GetAddonSchemaResponseBodyFields) GetDescription() *string {
	return s.Description
}

func (s *GetAddonSchemaResponseBodyFields) GetDisabled() *bool {
	return s.Disabled
}

func (s *GetAddonSchemaResponseBodyFields) GetElement() *string {
	return s.Element
}

func (s *GetAddonSchemaResponseBodyFields) GetFieldPath() *string {
	return s.FieldPath
}

func (s *GetAddonSchemaResponseBodyFields) GetLabel() *string {
	return s.Label
}

func (s *GetAddonSchemaResponseBodyFields) GetName() *string {
	return s.Name
}

func (s *GetAddonSchemaResponseBodyFields) GetPlaceholder() *string {
	return s.Placeholder
}

func (s *GetAddonSchemaResponseBodyFields) GetProps() *GetAddonSchemaResponseBodyFieldsProps {
	return s.Props
}

func (s *GetAddonSchemaResponseBodyFields) GetType() *string {
	return s.Type
}

func (s *GetAddonSchemaResponseBodyFields) GetValidation() *GetAddonSchemaResponseBodyFieldsValidation {
	return s.Validation
}

func (s *GetAddonSchemaResponseBodyFields) SetConditions(v []*GetAddonSchemaResponseBodyFieldsConditions) *GetAddonSchemaResponseBodyFields {
	s.Conditions = v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetDefaultValue(v interface{}) *GetAddonSchemaResponseBodyFields {
	s.DefaultValue = v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetDescription(v string) *GetAddonSchemaResponseBodyFields {
	s.Description = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetDisabled(v bool) *GetAddonSchemaResponseBodyFields {
	s.Disabled = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetElement(v string) *GetAddonSchemaResponseBodyFields {
	s.Element = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetFieldPath(v string) *GetAddonSchemaResponseBodyFields {
	s.FieldPath = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetLabel(v string) *GetAddonSchemaResponseBodyFields {
	s.Label = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetName(v string) *GetAddonSchemaResponseBodyFields {
	s.Name = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetPlaceholder(v string) *GetAddonSchemaResponseBodyFields {
	s.Placeholder = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetProps(v *GetAddonSchemaResponseBodyFieldsProps) *GetAddonSchemaResponseBodyFields {
	s.Props = v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetType(v string) *GetAddonSchemaResponseBodyFields {
	s.Type = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) SetValidation(v *GetAddonSchemaResponseBodyFieldsValidation) *GetAddonSchemaResponseBodyFields {
	s.Validation = v
	return s
}

func (s *GetAddonSchemaResponseBodyFields) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Props != nil {
		if err := s.Props.Validate(); err != nil {
			return err
		}
	}
	if s.Validation != nil {
		if err := s.Validation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAddonSchemaResponseBodyFieldsConditions struct {
	// The control mode of the field.
	//
	// example:
	//
	// show
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The name of the field.
	//
	// example:
	//
	// store.storageTarget
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The operator for the condition.
	//
	// example:
	//
	// ==
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
	// The target value for the condition.
	//
	// example:
	//
	// Custom
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAddonSchemaResponseBodyFieldsConditions) String() string {
	return dara.Prettify(s)
}

func (s GetAddonSchemaResponseBodyFieldsConditions) GoString() string {
	return s.String()
}

func (s *GetAddonSchemaResponseBodyFieldsConditions) GetAction() *string {
	return s.Action
}

func (s *GetAddonSchemaResponseBodyFieldsConditions) GetField() *string {
	return s.Field
}

func (s *GetAddonSchemaResponseBodyFieldsConditions) GetOp() *string {
	return s.Op
}

func (s *GetAddonSchemaResponseBodyFieldsConditions) GetValue() interface{} {
	return s.Value
}

func (s *GetAddonSchemaResponseBodyFieldsConditions) SetAction(v string) *GetAddonSchemaResponseBodyFieldsConditions {
	s.Action = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsConditions) SetField(v string) *GetAddonSchemaResponseBodyFieldsConditions {
	s.Field = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsConditions) SetOp(v string) *GetAddonSchemaResponseBodyFieldsConditions {
	s.Op = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsConditions) SetValue(v interface{}) *GetAddonSchemaResponseBodyFieldsConditions {
	s.Value = v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsConditions) Validate() error {
	return dara.Validate(s)
}

type GetAddonSchemaResponseBodyFieldsProps struct {
	// The information about the data source.
	DataSource []*GetAddonSchemaResponseBodyFieldsPropsDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Repeated"`
	// A collection of related data.
	Related []*string `json:"related,omitempty" xml:"related,omitempty" type:"Repeated"`
	// The mode.
	//
	// example:
	//
	// single
	SelectMode *string `json:"selectMode,omitempty" xml:"selectMode,omitempty"`
}

func (s GetAddonSchemaResponseBodyFieldsProps) String() string {
	return dara.Prettify(s)
}

func (s GetAddonSchemaResponseBodyFieldsProps) GoString() string {
	return s.String()
}

func (s *GetAddonSchemaResponseBodyFieldsProps) GetDataSource() []*GetAddonSchemaResponseBodyFieldsPropsDataSource {
	return s.DataSource
}

func (s *GetAddonSchemaResponseBodyFieldsProps) GetRelated() []*string {
	return s.Related
}

func (s *GetAddonSchemaResponseBodyFieldsProps) GetSelectMode() *string {
	return s.SelectMode
}

func (s *GetAddonSchemaResponseBodyFieldsProps) SetDataSource(v []*GetAddonSchemaResponseBodyFieldsPropsDataSource) *GetAddonSchemaResponseBodyFieldsProps {
	s.DataSource = v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsProps) SetRelated(v []*string) *GetAddonSchemaResponseBodyFieldsProps {
	s.Related = v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsProps) SetSelectMode(v string) *GetAddonSchemaResponseBodyFieldsProps {
	s.SelectMode = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsProps) Validate() error {
	if s.DataSource != nil {
		for _, item := range s.DataSource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAddonSchemaResponseBodyFieldsPropsDataSource struct {
	// The label.
	//
	// example:
	//
	// tagKey
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The value of the label.
	//
	// example:
	//
	// tagValue
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAddonSchemaResponseBodyFieldsPropsDataSource) String() string {
	return dara.Prettify(s)
}

func (s GetAddonSchemaResponseBodyFieldsPropsDataSource) GoString() string {
	return s.String()
}

func (s *GetAddonSchemaResponseBodyFieldsPropsDataSource) GetLabel() *string {
	return s.Label
}

func (s *GetAddonSchemaResponseBodyFieldsPropsDataSource) GetValue() *string {
	return s.Value
}

func (s *GetAddonSchemaResponseBodyFieldsPropsDataSource) SetLabel(v string) *GetAddonSchemaResponseBodyFieldsPropsDataSource {
	s.Label = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsPropsDataSource) SetValue(v string) *GetAddonSchemaResponseBodyFieldsPropsDataSource {
	s.Value = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsPropsDataSource) Validate() error {
	return dara.Validate(s)
}

type GetAddonSchemaResponseBodyFieldsValidation struct {
	// The maximum value. The value is inclusive.
	//
	// example:
	//
	// 2000
	Max *int32 `json:"max,omitempty" xml:"max,omitempty"`
	// The maximum length that is supported by the text control.
	//
	// example:
	//
	// 100
	MaxLength *int32 `json:"maxLength,omitempty" xml:"maxLength,omitempty"`
	// The details.
	//
	// example:
	//
	// length limit
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The minimum value.
	//
	// example:
	//
	// 0
	Min *int32 `json:"min,omitempty" xml:"min,omitempty"`
	// The minimum length.
	//
	// example:
	//
	// 16
	MinLength *int32 `json:"minLength,omitempty" xml:"minLength,omitempty"`
	// The regular expression.
	//
	// example:
	//
	// ^(http|https)://
	Regular *string `json:"regular,omitempty" xml:"regular,omitempty"`
	// Specifies whether the parameter is required.
	//
	// example:
	//
	// false
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
}

func (s GetAddonSchemaResponseBodyFieldsValidation) String() string {
	return dara.Prettify(s)
}

func (s GetAddonSchemaResponseBodyFieldsValidation) GoString() string {
	return s.String()
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) GetMax() *int32 {
	return s.Max
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) GetMessage() *string {
	return s.Message
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) GetMin() *int32 {
	return s.Min
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) GetMinLength() *int32 {
	return s.MinLength
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) GetRegular() *string {
	return s.Regular
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) GetRequired() *bool {
	return s.Required
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) SetMax(v int32) *GetAddonSchemaResponseBodyFieldsValidation {
	s.Max = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) SetMaxLength(v int32) *GetAddonSchemaResponseBodyFieldsValidation {
	s.MaxLength = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) SetMessage(v string) *GetAddonSchemaResponseBodyFieldsValidation {
	s.Message = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) SetMin(v int32) *GetAddonSchemaResponseBodyFieldsValidation {
	s.Min = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) SetMinLength(v int32) *GetAddonSchemaResponseBodyFieldsValidation {
	s.MinLength = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) SetRegular(v string) *GetAddonSchemaResponseBodyFieldsValidation {
	s.Regular = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) SetRequired(v bool) *GetAddonSchemaResponseBodyFieldsValidation {
	s.Required = &v
	return s
}

func (s *GetAddonSchemaResponseBodyFieldsValidation) Validate() error {
	return dara.Validate(s)
}

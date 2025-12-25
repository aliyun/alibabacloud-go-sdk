// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayFeaturesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGatewayFeaturesResponseBody
	GetCode() *string
	SetData(v *ListGatewayFeaturesResponseBodyData) *ListGatewayFeaturesResponseBody
	GetData() *ListGatewayFeaturesResponseBodyData
	SetMessage(v string) *ListGatewayFeaturesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayFeaturesResponseBody
	GetRequestId() *string
}

type ListGatewayFeaturesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListGatewayFeaturesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayFeaturesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayFeaturesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGatewayFeaturesResponseBody) GetData() *ListGatewayFeaturesResponseBodyData {
	return s.Data
}

func (s *ListGatewayFeaturesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayFeaturesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayFeaturesResponseBody) SetCode(v string) *ListGatewayFeaturesResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayFeaturesResponseBody) SetData(v *ListGatewayFeaturesResponseBodyData) *ListGatewayFeaturesResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayFeaturesResponseBody) SetMessage(v string) *ListGatewayFeaturesResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayFeaturesResponseBody) SetRequestId(v string) *ListGatewayFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayFeaturesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayFeaturesResponseBodyData struct {
	// The list of parameter configurations.
	Items []*ListGatewayFeaturesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListGatewayFeaturesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFeaturesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayFeaturesResponseBodyData) GetItems() []*ListGatewayFeaturesResponseBodyDataItems {
	return s.Items
}

func (s *ListGatewayFeaturesResponseBodyData) SetItems(v []*ListGatewayFeaturesResponseBodyDataItems) *ListGatewayFeaturesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewayFeaturesResponseBodyData) Validate() error {
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

type ListGatewayFeaturesResponseBodyDataItems struct {
	// The parameter definition.
	Definition *ListGatewayFeaturesResponseBodyDataItemsDefinition `json:"definition,omitempty" xml:"definition,omitempty" type:"Struct"`
	// The parameter value.
	//
	// example:
	//
	// "true"
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListGatewayFeaturesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFeaturesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListGatewayFeaturesResponseBodyDataItems) GetDefinition() *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	return s.Definition
}

func (s *ListGatewayFeaturesResponseBodyDataItems) GetValue() *string {
	return s.Value
}

func (s *ListGatewayFeaturesResponseBodyDataItems) SetDefinition(v *ListGatewayFeaturesResponseBodyDataItemsDefinition) *ListGatewayFeaturesResponseBodyDataItems {
	s.Definition = v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItems) SetValue(v string) *ListGatewayFeaturesResponseBodyDataItems {
	s.Value = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItems) Validate() error {
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayFeaturesResponseBodyDataItemsDefinition struct {
	// The default value of the parameter.
	//
	// example:
	//
	// "true"
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// The parameter description.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the parameter.
	//
	// example:
	//
	// EnableGzip
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The parameter group to which the parameter belongs. Valid values:
	//
	// 	- Telemetry: an observability parameter
	//
	// 	- Engine: an engine parameter
	//
	// example:
	//
	// Engine
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The input type of the parameter. Valid values:
	//
	// 	- Trigger
	//
	// 	- Input
	//
	// 	- SingleSelect
	//
	// 	- MultiSelect
	//
	// example:
	//
	// Trigger
	InputType *string `json:"inputType,omitempty" xml:"inputType,omitempty"`
	// The maximum length of the value. This parameter is valid when the value type is string.
	//
	// example:
	//
	// 64
	MaxLength *int32 `json:"maxLength,omitempty" xml:"maxLength,omitempty"`
	// The maximum value of the parameter. This parameter is valid when the value type is int32, int64, or float.
	//
	// example:
	//
	// 65535
	MaxValue *string `json:"maxValue,omitempty" xml:"maxValue,omitempty"`
	// The minimum length of the value. This parameter is valid when the value type is string.
	//
	// example:
	//
	// 8
	MinLength *int32 `json:"minLength,omitempty" xml:"minLength,omitempty"`
	// The minimum value of the parameter. This parameter is valid when the value type is int32, int64, or float.
	//
	// example:
	//
	// 100
	MinValue *string `json:"minValue,omitempty" xml:"minValue,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// enable-gzip
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the parameter is read-only.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"readOnly,omitempty" xml:"readOnly,omitempty"`
	// The regular expression that the parameter value must fulfill. This parameter is valid when the value type is string.
	//
	// example:
	//
	// [a-z].*
	Regex *string `json:"regex,omitempty" xml:"regex,omitempty"`
	// The list of options supported by the parameter value.
	ValueOptions []*ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions `json:"valueOptions,omitempty" xml:"valueOptions,omitempty" type:"Repeated"`
	// The value type of the parameter. Valid values:
	//
	// 	- bool: boolean
	//
	// 	- string
	//
	// 	- int32: integer
	//
	// 	- int64: long integer
	//
	// 	- json
	//
	// 	- array: JSON array
	//
	// 	- float: floating point
	//
	// example:
	//
	// bool
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
	// The value unit.
	//
	// example:
	//
	// byte
	ValueUnit *string `json:"valueUnit,omitempty" xml:"valueUnit,omitempty"`
}

func (s ListGatewayFeaturesResponseBodyDataItemsDefinition) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFeaturesResponseBodyDataItemsDefinition) GoString() string {
	return s.String()
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetDescription() *string {
	return s.Description
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetGroup() *string {
	return s.Group
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetInputType() *string {
	return s.InputType
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetMaxValue() *string {
	return s.MaxValue
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetMinLength() *int32 {
	return s.MinLength
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetMinValue() *string {
	return s.MinValue
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetName() *string {
	return s.Name
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetRegex() *string {
	return s.Regex
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetValueOptions() []*ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions {
	return s.ValueOptions
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetValueType() *string {
	return s.ValueType
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) GetValueUnit() *string {
	return s.ValueUnit
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetDefaultValue(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.DefaultValue = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetDescription(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.Description = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetDisplayName(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.DisplayName = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetGroup(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.Group = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetInputType(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.InputType = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetMaxLength(v int32) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.MaxLength = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetMaxValue(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.MaxValue = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetMinLength(v int32) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.MinLength = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetMinValue(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.MinValue = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetName(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.Name = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetReadOnly(v bool) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.ReadOnly = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetRegex(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.Regex = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetValueOptions(v []*ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.ValueOptions = v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetValueType(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.ValueType = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) SetValueUnit(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinition {
	s.ValueUnit = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinition) Validate() error {
	if s.ValueOptions != nil {
		for _, item := range s.ValueOptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions struct {
	// The key to pass the parameter.
	//
	// example:
	//
	// KEEP_UNCHANGED
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The display value.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
}

func (s ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions) GoString() string {
	return s.String()
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions) GetKey() *string {
	return s.Key
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions) GetLabel() *string {
	return s.Label
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions) SetKey(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions {
	s.Key = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions) SetLabel(v string) *ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions {
	s.Label = &v
	return s
}

func (s *ListGatewayFeaturesResponseBodyDataItemsDefinitionValueOptions) Validate() error {
	return dara.Validate(s)
}

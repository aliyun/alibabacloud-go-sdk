// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSchemaPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AddSchemaPropertyRequest
	GetInstanceId() *string
	SetProperty(v *AddSchemaPropertyRequestProperty) *AddSchemaPropertyRequest
	GetProperty() *AddSchemaPropertyRequestProperty
	SetRequestId(v string) *AddSchemaPropertyRequest
	GetRequestId() *string
	SetSchemaId(v string) *AddSchemaPropertyRequest
	GetSchemaId() *string
}

type AddSchemaPropertyRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b0eb2742-f37e-4c67-82d4-25c651c1xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Field
	Property *AddSchemaPropertyRequestProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// schema id
	//
	// This parameter is required.
	//
	// example:
	//
	// profile
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s AddSchemaPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSchemaPropertyRequest) GoString() string {
	return s.String()
}

func (s *AddSchemaPropertyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddSchemaPropertyRequest) GetProperty() *AddSchemaPropertyRequestProperty {
	return s.Property
}

func (s *AddSchemaPropertyRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSchemaPropertyRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *AddSchemaPropertyRequest) SetInstanceId(v string) *AddSchemaPropertyRequest {
	s.InstanceId = &v
	return s
}

func (s *AddSchemaPropertyRequest) SetProperty(v *AddSchemaPropertyRequestProperty) *AddSchemaPropertyRequest {
	s.Property = v
	return s
}

func (s *AddSchemaPropertyRequest) SetRequestId(v string) *AddSchemaPropertyRequest {
	s.RequestId = &v
	return s
}

func (s *AddSchemaPropertyRequest) SetSchemaId(v string) *AddSchemaPropertyRequest {
	s.SchemaId = &v
	return s
}

func (s *AddSchemaPropertyRequest) Validate() error {
	if s.Property != nil {
		if err := s.Property.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddSchemaPropertyRequestProperty struct {
	// Is array
	//
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// Extension attributes
	//
	// example:
	//
	// {\\"newName\\":\\"小桔充电-demo\\",\\"appId\\":\\"69FRKB4193W8BYP0\\"}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// Data type
	//
	// This parameter is required.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Description.
	//
	// example:
	//
	// -
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Is disabled
	//
	// example:
	//
	// False
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// Name
	//
	// example:
	//
	// name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// List display order
	//
	// example:
	//
	// 10
	DisplayOrder *int32 `json:"DisplayOrder,omitempty" xml:"DisplayOrder,omitempty"`
	// Editor type
	//
	// example:
	//
	// textbox
	EditorType *string `json:"EditorType,omitempty" xml:"EditorType,omitempty"`
	// Maximum length
	//
	// example:
	//
	// 100
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// Maximum numeric value
	//
	// example:
	//
	// 1
	Maximum *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// Minimum length
	//
	// example:
	//
	// 1
	MinLength *int32 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// Minimum numeric value
	//
	// example:
	//
	// 1
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// Name
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Regular expression validation rule
	//
	// example:
	//
	// *
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	// Regular expression validation error message
	//
	// example:
	//
	// 格式错误
	PatternErrorMessage *string `json:"PatternErrorMessage,omitempty" xml:"PatternErrorMessage,omitempty"`
	// Is read-only
	//
	// example:
	//
	// true
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// Is required
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s AddSchemaPropertyRequestProperty) String() string {
	return dara.Prettify(s)
}

func (s AddSchemaPropertyRequestProperty) GoString() string {
	return s.String()
}

func (s *AddSchemaPropertyRequestProperty) GetArray() *bool {
	return s.Array
}

func (s *AddSchemaPropertyRequestProperty) GetAttributes() *string {
	return s.Attributes
}

func (s *AddSchemaPropertyRequestProperty) GetDataType() *string {
	return s.DataType
}

func (s *AddSchemaPropertyRequestProperty) GetDescription() *string {
	return s.Description
}

func (s *AddSchemaPropertyRequestProperty) GetDisabled() *bool {
	return s.Disabled
}

func (s *AddSchemaPropertyRequestProperty) GetDisplayName() *string {
	return s.DisplayName
}

func (s *AddSchemaPropertyRequestProperty) GetDisplayOrder() *int32 {
	return s.DisplayOrder
}

func (s *AddSchemaPropertyRequestProperty) GetEditorType() *string {
	return s.EditorType
}

func (s *AddSchemaPropertyRequestProperty) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *AddSchemaPropertyRequestProperty) GetMaximum() *float64 {
	return s.Maximum
}

func (s *AddSchemaPropertyRequestProperty) GetMinLength() *int32 {
	return s.MinLength
}

func (s *AddSchemaPropertyRequestProperty) GetMinimum() *float64 {
	return s.Minimum
}

func (s *AddSchemaPropertyRequestProperty) GetName() *string {
	return s.Name
}

func (s *AddSchemaPropertyRequestProperty) GetPattern() *string {
	return s.Pattern
}

func (s *AddSchemaPropertyRequestProperty) GetPatternErrorMessage() *string {
	return s.PatternErrorMessage
}

func (s *AddSchemaPropertyRequestProperty) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *AddSchemaPropertyRequestProperty) GetRequired() *bool {
	return s.Required
}

func (s *AddSchemaPropertyRequestProperty) SetArray(v bool) *AddSchemaPropertyRequestProperty {
	s.Array = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetAttributes(v string) *AddSchemaPropertyRequestProperty {
	s.Attributes = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetDataType(v string) *AddSchemaPropertyRequestProperty {
	s.DataType = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetDescription(v string) *AddSchemaPropertyRequestProperty {
	s.Description = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetDisabled(v bool) *AddSchemaPropertyRequestProperty {
	s.Disabled = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetDisplayName(v string) *AddSchemaPropertyRequestProperty {
	s.DisplayName = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetDisplayOrder(v int32) *AddSchemaPropertyRequestProperty {
	s.DisplayOrder = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetEditorType(v string) *AddSchemaPropertyRequestProperty {
	s.EditorType = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetMaxLength(v int32) *AddSchemaPropertyRequestProperty {
	s.MaxLength = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetMaximum(v float64) *AddSchemaPropertyRequestProperty {
	s.Maximum = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetMinLength(v int32) *AddSchemaPropertyRequestProperty {
	s.MinLength = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetMinimum(v float64) *AddSchemaPropertyRequestProperty {
	s.Minimum = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetName(v string) *AddSchemaPropertyRequestProperty {
	s.Name = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetPattern(v string) *AddSchemaPropertyRequestProperty {
	s.Pattern = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetPatternErrorMessage(v string) *AddSchemaPropertyRequestProperty {
	s.PatternErrorMessage = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetReadOnly(v bool) *AddSchemaPropertyRequestProperty {
	s.ReadOnly = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) SetRequired(v bool) *AddSchemaPropertyRequestProperty {
	s.Required = &v
	return s
}

func (s *AddSchemaPropertyRequestProperty) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSchemaPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateSchemaPropertyRequest
	GetInstanceId() *string
	SetProperty(v *UpdateSchemaPropertyRequestProperty) *UpdateSchemaPropertyRequest
	GetProperty() *UpdateSchemaPropertyRequestProperty
	SetRequestId(v string) *UpdateSchemaPropertyRequest
	GetRequestId() *string
	SetSchemaId(v string) *UpdateSchemaPropertyRequest
	GetSchemaId() *string
}

type UpdateSchemaPropertyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b0eb2742-f37e-4c67-82d4-25c651c1xxxx
	InstanceId *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Property   *UpdateSchemaPropertyRequestProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Struct"`
	// example:
	//
	// BC976D32-AC4C-4E0F-8AA9-F4BC6C4E2B3E
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

func (s UpdateSchemaPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemaPropertyRequest) GoString() string {
	return s.String()
}

func (s *UpdateSchemaPropertyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateSchemaPropertyRequest) GetProperty() *UpdateSchemaPropertyRequestProperty {
	return s.Property
}

func (s *UpdateSchemaPropertyRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSchemaPropertyRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *UpdateSchemaPropertyRequest) SetInstanceId(v string) *UpdateSchemaPropertyRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSchemaPropertyRequest) SetProperty(v *UpdateSchemaPropertyRequestProperty) *UpdateSchemaPropertyRequest {
	s.Property = v
	return s
}

func (s *UpdateSchemaPropertyRequest) SetRequestId(v string) *UpdateSchemaPropertyRequest {
	s.RequestId = &v
	return s
}

func (s *UpdateSchemaPropertyRequest) SetSchemaId(v string) *UpdateSchemaPropertyRequest {
	s.SchemaId = &v
	return s
}

func (s *UpdateSchemaPropertyRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateSchemaPropertyRequestProperty struct {
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// {\\"instanceId\\":\\"4cc3f160-ca64-49ff-bc70-390a044a4e83\\",\\"appId\\":\\"1684145288664\\",\\"commodityCode\\":\\"dide_pre\\",\\"dide_pre_set\\":\\"version_ent\\"}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// -
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Disabled    *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 10
	DisplayOrder *int32 `json:"DisplayOrder,omitempty" xml:"DisplayOrder,omitempty"`
	// example:
	//
	// textbox
	EditorType *string `json:"EditorType,omitempty" xml:"EditorType,omitempty"`
	// example:
	//
	// 100
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// example:
	//
	// 11
	Maximum *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// example:
	//
	// 1
	MinLength *int32 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// example:
	//
	// 10800
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// (.*)
	Pattern             *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	PatternErrorMessage *string `json:"PatternErrorMessage,omitempty" xml:"PatternErrorMessage,omitempty"`
	// example:
	//
	// true
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s UpdateSchemaPropertyRequestProperty) String() string {
	return dara.Prettify(s)
}

func (s UpdateSchemaPropertyRequestProperty) GoString() string {
	return s.String()
}

func (s *UpdateSchemaPropertyRequestProperty) GetArray() *bool {
	return s.Array
}

func (s *UpdateSchemaPropertyRequestProperty) GetAttributes() *string {
	return s.Attributes
}

func (s *UpdateSchemaPropertyRequestProperty) GetDataType() *string {
	return s.DataType
}

func (s *UpdateSchemaPropertyRequestProperty) GetDescription() *string {
	return s.Description
}

func (s *UpdateSchemaPropertyRequestProperty) GetDisabled() *bool {
	return s.Disabled
}

func (s *UpdateSchemaPropertyRequestProperty) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateSchemaPropertyRequestProperty) GetDisplayOrder() *int32 {
	return s.DisplayOrder
}

func (s *UpdateSchemaPropertyRequestProperty) GetEditorType() *string {
	return s.EditorType
}

func (s *UpdateSchemaPropertyRequestProperty) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *UpdateSchemaPropertyRequestProperty) GetMaximum() *float64 {
	return s.Maximum
}

func (s *UpdateSchemaPropertyRequestProperty) GetMinLength() *int32 {
	return s.MinLength
}

func (s *UpdateSchemaPropertyRequestProperty) GetMinimum() *float64 {
	return s.Minimum
}

func (s *UpdateSchemaPropertyRequestProperty) GetName() *string {
	return s.Name
}

func (s *UpdateSchemaPropertyRequestProperty) GetPattern() *string {
	return s.Pattern
}

func (s *UpdateSchemaPropertyRequestProperty) GetPatternErrorMessage() *string {
	return s.PatternErrorMessage
}

func (s *UpdateSchemaPropertyRequestProperty) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateSchemaPropertyRequestProperty) GetRequired() *bool {
	return s.Required
}

func (s *UpdateSchemaPropertyRequestProperty) SetArray(v bool) *UpdateSchemaPropertyRequestProperty {
	s.Array = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetAttributes(v string) *UpdateSchemaPropertyRequestProperty {
	s.Attributes = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetDataType(v string) *UpdateSchemaPropertyRequestProperty {
	s.DataType = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetDescription(v string) *UpdateSchemaPropertyRequestProperty {
	s.Description = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetDisabled(v bool) *UpdateSchemaPropertyRequestProperty {
	s.Disabled = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetDisplayName(v string) *UpdateSchemaPropertyRequestProperty {
	s.DisplayName = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetDisplayOrder(v int32) *UpdateSchemaPropertyRequestProperty {
	s.DisplayOrder = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetEditorType(v string) *UpdateSchemaPropertyRequestProperty {
	s.EditorType = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetMaxLength(v int32) *UpdateSchemaPropertyRequestProperty {
	s.MaxLength = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetMaximum(v float64) *UpdateSchemaPropertyRequestProperty {
	s.Maximum = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetMinLength(v int32) *UpdateSchemaPropertyRequestProperty {
	s.MinLength = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetMinimum(v float64) *UpdateSchemaPropertyRequestProperty {
	s.Minimum = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetName(v string) *UpdateSchemaPropertyRequestProperty {
	s.Name = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetPattern(v string) *UpdateSchemaPropertyRequestProperty {
	s.Pattern = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetPatternErrorMessage(v string) *UpdateSchemaPropertyRequestProperty {
	s.PatternErrorMessage = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetReadOnly(v bool) *UpdateSchemaPropertyRequestProperty {
	s.ReadOnly = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) SetRequired(v bool) *UpdateSchemaPropertyRequestProperty {
	s.Required = &v
	return s
}

func (s *UpdateSchemaPropertyRequestProperty) Validate() error {
	return dara.Validate(s)
}

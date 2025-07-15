// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSchemaPropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *DataSchemaPropertiesValue
	GetDisplayName() *string
	SetDescription(v string) *DataSchemaPropertiesValue
	GetDescription() *string
	SetName(v string) *DataSchemaPropertiesValue
	GetName() *string
	SetDataType(v string) *DataSchemaPropertiesValue
	GetDataType() *string
	SetPattern(v string) *DataSchemaPropertiesValue
	GetPattern() *string
	SetPatternErrorMessage(v string) *DataSchemaPropertiesValue
	GetPatternErrorMessage() *string
	SetMinLength(v int32) *DataSchemaPropertiesValue
	GetMinLength() *int32
	SetMaxLength(v int32) *DataSchemaPropertiesValue
	GetMaxLength() *int32
	SetMinimum(v float64) *DataSchemaPropertiesValue
	GetMinimum() *float64
	SetMaximum(v float64) *DataSchemaPropertiesValue
	GetMaximum() *float64
	SetRequired(v bool) *DataSchemaPropertiesValue
	GetRequired() *bool
	SetSystem(v bool) *DataSchemaPropertiesValue
	GetSystem() *bool
	SetDisabled(v bool) *DataSchemaPropertiesValue
	GetDisabled() *bool
	SetArray(v bool) *DataSchemaPropertiesValue
	GetArray() *bool
	SetReadOnly(v bool) *DataSchemaPropertiesValue
	GetReadOnly() *bool
	SetEditorType(v string) *DataSchemaPropertiesValue
	GetEditorType() *string
	SetAttributes(v string) *DataSchemaPropertiesValue
	GetAttributes() *string
	SetDisplayOrder(v int32) *DataSchemaPropertiesValue
	GetDisplayOrder() *int32
	SetCreatedTime(v int64) *DataSchemaPropertiesValue
	GetCreatedTime() *int64
	SetUpdatedTime(v int64) *DataSchemaPropertiesValue
	GetUpdatedTime() *int64
	SetCreator(v string) *DataSchemaPropertiesValue
	GetCreator() *string
}

type DataSchemaPropertiesValue struct {
	// example:
	//
	// name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// -
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// string
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// *
	Pattern             *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	PatternErrorMessage *string `json:"PatternErrorMessage,omitempty" xml:"PatternErrorMessage,omitempty"`
	// example:
	//
	// 1
	MinLength *int32 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// example:
	//
	// 1
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// example:
	//
	// 1
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// example:
	//
	// 1
	Maximum *float64 `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// false
	System *bool `json:"System,omitempty" xml:"System,omitempty"`
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// example:
	//
	// textbox
	EditorType *string `json:"EditorType,omitempty" xml:"EditorType,omitempty"`
	// example:
	//
	// {}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// 1
	DisplayOrder *int32 `json:"DisplayOrder,omitempty" xml:"DisplayOrder,omitempty"`
	// example:
	//
	// 2020-10-14T09:53:53Z
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2020-10-14T09:53:53Z
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// example:
	//
	// tom
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
}

func (s DataSchemaPropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s DataSchemaPropertiesValue) GoString() string {
	return s.String()
}

func (s *DataSchemaPropertiesValue) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DataSchemaPropertiesValue) GetDescription() *string {
	return s.Description
}

func (s *DataSchemaPropertiesValue) GetName() *string {
	return s.Name
}

func (s *DataSchemaPropertiesValue) GetDataType() *string {
	return s.DataType
}

func (s *DataSchemaPropertiesValue) GetPattern() *string {
	return s.Pattern
}

func (s *DataSchemaPropertiesValue) GetPatternErrorMessage() *string {
	return s.PatternErrorMessage
}

func (s *DataSchemaPropertiesValue) GetMinLength() *int32 {
	return s.MinLength
}

func (s *DataSchemaPropertiesValue) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *DataSchemaPropertiesValue) GetMinimum() *float64 {
	return s.Minimum
}

func (s *DataSchemaPropertiesValue) GetMaximum() *float64 {
	return s.Maximum
}

func (s *DataSchemaPropertiesValue) GetRequired() *bool {
	return s.Required
}

func (s *DataSchemaPropertiesValue) GetSystem() *bool {
	return s.System
}

func (s *DataSchemaPropertiesValue) GetDisabled() *bool {
	return s.Disabled
}

func (s *DataSchemaPropertiesValue) GetArray() *bool {
	return s.Array
}

func (s *DataSchemaPropertiesValue) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DataSchemaPropertiesValue) GetEditorType() *string {
	return s.EditorType
}

func (s *DataSchemaPropertiesValue) GetAttributes() *string {
	return s.Attributes
}

func (s *DataSchemaPropertiesValue) GetDisplayOrder() *int32 {
	return s.DisplayOrder
}

func (s *DataSchemaPropertiesValue) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DataSchemaPropertiesValue) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DataSchemaPropertiesValue) GetCreator() *string {
	return s.Creator
}

func (s *DataSchemaPropertiesValue) SetDisplayName(v string) *DataSchemaPropertiesValue {
	s.DisplayName = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetDescription(v string) *DataSchemaPropertiesValue {
	s.Description = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetName(v string) *DataSchemaPropertiesValue {
	s.Name = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetDataType(v string) *DataSchemaPropertiesValue {
	s.DataType = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetPattern(v string) *DataSchemaPropertiesValue {
	s.Pattern = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetPatternErrorMessage(v string) *DataSchemaPropertiesValue {
	s.PatternErrorMessage = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetMinLength(v int32) *DataSchemaPropertiesValue {
	s.MinLength = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetMaxLength(v int32) *DataSchemaPropertiesValue {
	s.MaxLength = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetMinimum(v float64) *DataSchemaPropertiesValue {
	s.Minimum = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetMaximum(v float64) *DataSchemaPropertiesValue {
	s.Maximum = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetRequired(v bool) *DataSchemaPropertiesValue {
	s.Required = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetSystem(v bool) *DataSchemaPropertiesValue {
	s.System = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetDisabled(v bool) *DataSchemaPropertiesValue {
	s.Disabled = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetArray(v bool) *DataSchemaPropertiesValue {
	s.Array = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetReadOnly(v bool) *DataSchemaPropertiesValue {
	s.ReadOnly = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetEditorType(v string) *DataSchemaPropertiesValue {
	s.EditorType = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetAttributes(v string) *DataSchemaPropertiesValue {
	s.Attributes = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetDisplayOrder(v int32) *DataSchemaPropertiesValue {
	s.DisplayOrder = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetCreatedTime(v int64) *DataSchemaPropertiesValue {
	s.CreatedTime = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetUpdatedTime(v int64) *DataSchemaPropertiesValue {
	s.UpdatedTime = &v
	return s
}

func (s *DataSchemaPropertiesValue) SetCreator(v string) *DataSchemaPropertiesValue {
	s.Creator = &v
	return s
}

func (s *DataSchemaPropertiesValue) Validate() error {
	return dara.Validate(s)
}

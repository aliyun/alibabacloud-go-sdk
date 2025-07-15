// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataPropertiesValue interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *DataPropertiesValue
	GetDisplayName() *string
	SetDescription(v string) *DataPropertiesValue
	GetDescription() *string
	SetName(v string) *DataPropertiesValue
	GetName() *string
	SetDataType(v string) *DataPropertiesValue
	GetDataType() *string
	SetPattern(v string) *DataPropertiesValue
	GetPattern() *string
	SetPatternErrorMessage(v string) *DataPropertiesValue
	GetPatternErrorMessage() *string
	SetMinLength(v int32) *DataPropertiesValue
	GetMinLength() *int32
	SetMaxLength(v int32) *DataPropertiesValue
	GetMaxLength() *int32
	SetMinimum(v float64) *DataPropertiesValue
	GetMinimum() *float64
	SetMaximum(v float64) *DataPropertiesValue
	GetMaximum() *float64
	SetRequired(v bool) *DataPropertiesValue
	GetRequired() *bool
	SetSystem(v bool) *DataPropertiesValue
	GetSystem() *bool
	SetDisabled(v bool) *DataPropertiesValue
	GetDisabled() *bool
	SetArray(v bool) *DataPropertiesValue
	GetArray() *bool
	SetReadOnly(v bool) *DataPropertiesValue
	GetReadOnly() *bool
	SetEditorType(v string) *DataPropertiesValue
	GetEditorType() *string
	SetAttributes(v string) *DataPropertiesValue
	GetAttributes() *string
	SetDisplayOrder(v int32) *DataPropertiesValue
	GetDisplayOrder() *int32
	SetCreatedTime(v int64) *DataPropertiesValue
	GetCreatedTime() *int64
	SetUpdatedTime(v int64) *DataPropertiesValue
	GetUpdatedTime() *int64
	SetCreator(v string) *DataPropertiesValue
	GetCreator() *string
}

type DataPropertiesValue struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// name
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
	// ^
	Pattern             *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
	PatternErrorMessage *string `json:"PatternErrorMessage,omitempty" xml:"PatternErrorMessage,omitempty"`
	// example:
	//
	// 1
	MinLength *int32 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// example:
	//
	// 10
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// example:
	//
	// 1
	Minimum *float64 `json:"Minimum,omitempty" xml:"Minimum,omitempty"`
	// example:
	//
	// 10
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
	// 2021-07-14 10:48:43.0
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2021-07-14 10:48:43.0
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	// example:
	//
	// tom
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
}

func (s DataPropertiesValue) String() string {
	return dara.Prettify(s)
}

func (s DataPropertiesValue) GoString() string {
	return s.String()
}

func (s *DataPropertiesValue) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DataPropertiesValue) GetDescription() *string {
	return s.Description
}

func (s *DataPropertiesValue) GetName() *string {
	return s.Name
}

func (s *DataPropertiesValue) GetDataType() *string {
	return s.DataType
}

func (s *DataPropertiesValue) GetPattern() *string {
	return s.Pattern
}

func (s *DataPropertiesValue) GetPatternErrorMessage() *string {
	return s.PatternErrorMessage
}

func (s *DataPropertiesValue) GetMinLength() *int32 {
	return s.MinLength
}

func (s *DataPropertiesValue) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *DataPropertiesValue) GetMinimum() *float64 {
	return s.Minimum
}

func (s *DataPropertiesValue) GetMaximum() *float64 {
	return s.Maximum
}

func (s *DataPropertiesValue) GetRequired() *bool {
	return s.Required
}

func (s *DataPropertiesValue) GetSystem() *bool {
	return s.System
}

func (s *DataPropertiesValue) GetDisabled() *bool {
	return s.Disabled
}

func (s *DataPropertiesValue) GetArray() *bool {
	return s.Array
}

func (s *DataPropertiesValue) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *DataPropertiesValue) GetEditorType() *string {
	return s.EditorType
}

func (s *DataPropertiesValue) GetAttributes() *string {
	return s.Attributes
}

func (s *DataPropertiesValue) GetDisplayOrder() *int32 {
	return s.DisplayOrder
}

func (s *DataPropertiesValue) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DataPropertiesValue) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DataPropertiesValue) GetCreator() *string {
	return s.Creator
}

func (s *DataPropertiesValue) SetDisplayName(v string) *DataPropertiesValue {
	s.DisplayName = &v
	return s
}

func (s *DataPropertiesValue) SetDescription(v string) *DataPropertiesValue {
	s.Description = &v
	return s
}

func (s *DataPropertiesValue) SetName(v string) *DataPropertiesValue {
	s.Name = &v
	return s
}

func (s *DataPropertiesValue) SetDataType(v string) *DataPropertiesValue {
	s.DataType = &v
	return s
}

func (s *DataPropertiesValue) SetPattern(v string) *DataPropertiesValue {
	s.Pattern = &v
	return s
}

func (s *DataPropertiesValue) SetPatternErrorMessage(v string) *DataPropertiesValue {
	s.PatternErrorMessage = &v
	return s
}

func (s *DataPropertiesValue) SetMinLength(v int32) *DataPropertiesValue {
	s.MinLength = &v
	return s
}

func (s *DataPropertiesValue) SetMaxLength(v int32) *DataPropertiesValue {
	s.MaxLength = &v
	return s
}

func (s *DataPropertiesValue) SetMinimum(v float64) *DataPropertiesValue {
	s.Minimum = &v
	return s
}

func (s *DataPropertiesValue) SetMaximum(v float64) *DataPropertiesValue {
	s.Maximum = &v
	return s
}

func (s *DataPropertiesValue) SetRequired(v bool) *DataPropertiesValue {
	s.Required = &v
	return s
}

func (s *DataPropertiesValue) SetSystem(v bool) *DataPropertiesValue {
	s.System = &v
	return s
}

func (s *DataPropertiesValue) SetDisabled(v bool) *DataPropertiesValue {
	s.Disabled = &v
	return s
}

func (s *DataPropertiesValue) SetArray(v bool) *DataPropertiesValue {
	s.Array = &v
	return s
}

func (s *DataPropertiesValue) SetReadOnly(v bool) *DataPropertiesValue {
	s.ReadOnly = &v
	return s
}

func (s *DataPropertiesValue) SetEditorType(v string) *DataPropertiesValue {
	s.EditorType = &v
	return s
}

func (s *DataPropertiesValue) SetAttributes(v string) *DataPropertiesValue {
	s.Attributes = &v
	return s
}

func (s *DataPropertiesValue) SetDisplayOrder(v int32) *DataPropertiesValue {
	s.DisplayOrder = &v
	return s
}

func (s *DataPropertiesValue) SetCreatedTime(v int64) *DataPropertiesValue {
	s.CreatedTime = &v
	return s
}

func (s *DataPropertiesValue) SetUpdatedTime(v int64) *DataPropertiesValue {
	s.UpdatedTime = &v
	return s
}

func (s *DataPropertiesValue) SetCreator(v string) *DataPropertiesValue {
	s.Creator = &v
	return s
}

func (s *DataPropertiesValue) Validate() error {
	return dara.Validate(s)
}

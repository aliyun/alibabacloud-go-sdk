// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpsertQualityRuleRequest
	GetOpTenantId() *int64
	SetUpsertCommand(v *UpsertQualityRuleRequestUpsertCommand) *UpsertQualityRuleRequest
	GetUpsertCommand() *UpsertQualityRuleRequestUpsertCommand
}

type UpsertQualityRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpsertCommand *UpsertQualityRuleRequestUpsertCommand `json:"UpsertCommand,omitempty" xml:"UpsertCommand,omitempty" type:"Struct"`
}

func (s UpsertQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpsertQualityRuleRequest) GetUpsertCommand() *UpsertQualityRuleRequestUpsertCommand {
	return s.UpsertCommand
}

func (s *UpsertQualityRuleRequest) SetOpTenantId(v int64) *UpsertQualityRuleRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityRuleRequest) SetUpsertCommand(v *UpsertQualityRuleRequestUpsertCommand) *UpsertQualityRuleRequest {
	s.UpsertCommand = v
	return s
}

func (s *UpsertQualityRuleRequest) Validate() error {
	if s.UpsertCommand != nil {
		if err := s.UpsertCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityRuleRequestUpsertCommand struct {
	AttributeWithValueList []*UpsertQualityRuleRequestUpsertCommandAttributeWithValueList `json:"AttributeWithValueList,omitempty" xml:"AttributeWithValueList,omitempty" type:"Repeated"`
	// This parameter is required.
	CatalogList []*string `json:"CatalogList,omitempty" xml:"CatalogList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	Description        *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableErrorArchive *bool                                                    `json:"EnableErrorArchive,omitempty" xml:"EnableErrorArchive,omitempty"`
	FormPropertyList   []*UpsertQualityRuleRequestUpsertCommandFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STRONG
	Strength *string `json:"Strength,omitempty" xml:"Strength,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	TemplateType          *string                                                       `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	ValidateConditionList []*UpsertQualityRuleRequestUpsertCommandValidateConditionList `json:"ValidateConditionList,omitempty" xml:"ValidateConditionList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s UpsertQualityRuleRequestUpsertCommand) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleRequestUpsertCommand) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetAttributeWithValueList() []*UpsertQualityRuleRequestUpsertCommandAttributeWithValueList {
	return s.AttributeWithValueList
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetCatalogList() []*string {
	return s.CatalogList
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetDescription() *string {
	return s.Description
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetEnableErrorArchive() *bool {
	return s.EnableErrorArchive
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetFormPropertyList() []*UpsertQualityRuleRequestUpsertCommandFormPropertyList {
	return s.FormPropertyList
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetId() *int64 {
	return s.Id
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetName() *string {
	return s.Name
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetStrength() *string {
	return s.Strength
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetTemplateType() *string {
	return s.TemplateType
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetValidateConditionList() []*UpsertQualityRuleRequestUpsertCommandValidateConditionList {
	return s.ValidateConditionList
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetWatchId() *int64 {
	return s.WatchId
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetAttributeWithValueList(v []*UpsertQualityRuleRequestUpsertCommandAttributeWithValueList) *UpsertQualityRuleRequestUpsertCommand {
	s.AttributeWithValueList = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetCatalogList(v []*string) *UpsertQualityRuleRequestUpsertCommand {
	s.CatalogList = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetDescription(v string) *UpsertQualityRuleRequestUpsertCommand {
	s.Description = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetEnableErrorArchive(v bool) *UpsertQualityRuleRequestUpsertCommand {
	s.EnableErrorArchive = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetFormPropertyList(v []*UpsertQualityRuleRequestUpsertCommandFormPropertyList) *UpsertQualityRuleRequestUpsertCommand {
	s.FormPropertyList = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetId(v int64) *UpsertQualityRuleRequestUpsertCommand {
	s.Id = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetName(v string) *UpsertQualityRuleRequestUpsertCommand {
	s.Name = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetStrength(v string) *UpsertQualityRuleRequestUpsertCommand {
	s.Strength = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetTemplateId(v int64) *UpsertQualityRuleRequestUpsertCommand {
	s.TemplateId = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetTemplateType(v string) *UpsertQualityRuleRequestUpsertCommand {
	s.TemplateType = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetValidateConditionList(v []*UpsertQualityRuleRequestUpsertCommandValidateConditionList) *UpsertQualityRuleRequestUpsertCommand {
	s.ValidateConditionList = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetWatchId(v int64) *UpsertQualityRuleRequestUpsertCommand {
	s.WatchId = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) Validate() error {
	if s.AttributeWithValueList != nil {
		for _, item := range s.AttributeWithValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FormPropertyList != nil {
		for _, item := range s.FormPropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ValidateConditionList != nil {
		for _, item := range s.ValidateConditionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpsertQualityRuleRequestUpsertCommandAttributeWithValueList struct {
	AttributeInfo  *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo  `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty" type:"Struct"`
	AttributeValue *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty" type:"Struct"`
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueList) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueList) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueList) GetAttributeInfo() *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo {
	return s.AttributeInfo
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueList) GetAttributeValue() *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue {
	return s.AttributeValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueList) SetAttributeInfo(v *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueList {
	s.AttributeInfo = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueList) SetAttributeValue(v *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueList {
	s.AttributeValue = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueList) Validate() error {
	if s.AttributeInfo != nil {
		if err := s.AttributeInfo.Validate(); err != nil {
			return err
		}
	}
	if s.AttributeValue != nil {
		if err := s.AttributeValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Enabled     *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 711484689131
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// attr01
	Name        *string                                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	Required    *bool                                                                                `json:"Required,omitempty" xml:"Required,omitempty"`
	Searchable  *bool                                                                                `json:"Searchable,omitempty" xml:"Searchable,omitempty"`
	ValueConfig *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig `json:"ValueConfig,omitempty" xml:"ValueConfig,omitempty" type:"Struct"`
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) GetDescription() *string {
	return s.Description
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) GetId() *int64 {
	return s.Id
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) GetName() *string {
	return s.Name
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) GetRequired() *bool {
	return s.Required
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) GetSearchable() *bool {
	return s.Searchable
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) GetValueConfig() *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig {
	return s.ValueConfig
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) SetDescription(v string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo {
	s.Description = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) SetEnabled(v bool) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo {
	s.Enabled = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) SetId(v int64) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo {
	s.Id = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) SetName(v string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo {
	s.Name = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) SetRequired(v bool) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo {
	s.Required = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) SetSearchable(v bool) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo {
	s.Searchable = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) SetValueConfig(v *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo {
	s.ValueConfig = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo) Validate() error {
	if s.ValueConfig != nil {
		if err := s.ValueConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig struct {
	// example:
	//
	// STRING
	DataType     *string                                                                                          `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DefaultValue *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty" type:"Struct"`
	// example:
	//
	// 986992
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// CUSTOMIZED
	Type          *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	ValueEnumList []*string `json:"ValueEnumList,omitempty" xml:"ValueEnumList,omitempty" type:"Repeated"`
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) GetDataType() *string {
	return s.DataType
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) GetDefaultValue() *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	return s.DefaultValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) GetLength() *int32 {
	return s.Length
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) GetType() *string {
	return s.Type
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) GetValueEnumList() []*string {
	return s.ValueEnumList
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) SetDataType(v string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig {
	s.DataType = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) SetDefaultValue(v *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig {
	s.DefaultValue = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) SetLength(v int32) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig {
	s.Length = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) SetType(v string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig {
	s.Type = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) SetValueEnumList(v []*string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig {
	s.ValueEnumList = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfig) Validate() error {
	if s.DefaultValue != nil {
		if err := s.DefaultValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue struct {
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// example:
	//
	// 11
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// example:
	//
	// 1
	MinValue  *string   `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetMaxValue() *string {
	return s.MaxValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetMinValue() *string {
	return s.MinValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) GetValueList() []*string {
	return s.ValueList
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetIncludeMaxValue(v bool) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.IncludeMaxValue = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetIncludeMinValue(v bool) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.IncludeMinValue = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetMaxValue(v string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.MaxValue = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetMinValue(v string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.MinValue = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) SetValueList(v []*string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue {
	s.ValueList = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue) Validate() error {
	return dara.Validate(s)
}

type UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue struct {
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// example:
	//
	// 11
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// example:
	//
	// 1
	MinValue  *string   `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) GetMaxValue() *string {
	return s.MaxValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) GetMinValue() *string {
	return s.MinValue
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) GetValueList() []*string {
	return s.ValueList
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) SetIncludeMaxValue(v bool) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue {
	s.IncludeMaxValue = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) SetIncludeMinValue(v bool) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue {
	s.IncludeMinValue = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) SetMaxValue(v string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue {
	s.MaxValue = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) SetMinValue(v string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue {
	s.MinValue = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) SetValueList(v []*string) *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue {
	s.ValueList = v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeValue) Validate() error {
	return dara.Validate(s)
}

type UpsertQualityRuleRequestUpsertCommandFormPropertyList struct {
	// example:
	//
	// expression
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// col
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// abc
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpsertQualityRuleRequestUpsertCommandFormPropertyList) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleRequestUpsertCommandFormPropertyList) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleRequestUpsertCommandFormPropertyList) GetComponentType() *string {
	return s.ComponentType
}

func (s *UpsertQualityRuleRequestUpsertCommandFormPropertyList) GetName() *string {
	return s.Name
}

func (s *UpsertQualityRuleRequestUpsertCommandFormPropertyList) GetValue() *string {
	return s.Value
}

func (s *UpsertQualityRuleRequestUpsertCommandFormPropertyList) SetComponentType(v string) *UpsertQualityRuleRequestUpsertCommandFormPropertyList {
	s.ComponentType = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandFormPropertyList) SetName(v string) *UpsertQualityRuleRequestUpsertCommandFormPropertyList {
	s.Name = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandFormPropertyList) SetValue(v string) *UpsertQualityRuleRequestUpsertCommandFormPropertyList {
	s.Value = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandFormPropertyList) Validate() error {
	return dara.Validate(s)
}

type UpsertQualityRuleRequestUpsertCommandValidateConditionList struct {
	// example:
	//
	// 268
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// AND
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// RELATION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpsertQualityRuleRequestUpsertCommandValidateConditionList) String() string {
	return dara.Prettify(s)
}

func (s UpsertQualityRuleRequestUpsertCommandValidateConditionList) GoString() string {
	return s.String()
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) GetId() *string {
	return s.Id
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) GetMetric() *string {
	return s.Metric
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) GetOperator() *string {
	return s.Operator
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) GetParentId() *string {
	return s.ParentId
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) GetType() *string {
	return s.Type
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) GetValue() *string {
	return s.Value
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) SetId(v string) *UpsertQualityRuleRequestUpsertCommandValidateConditionList {
	s.Id = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) SetMetric(v string) *UpsertQualityRuleRequestUpsertCommandValidateConditionList {
	s.Metric = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) SetOperator(v string) *UpsertQualityRuleRequestUpsertCommandValidateConditionList {
	s.Operator = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) SetParentId(v string) *UpsertQualityRuleRequestUpsertCommandValidateConditionList {
	s.ParentId = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) SetType(v string) *UpsertQualityRuleRequestUpsertCommandValidateConditionList {
	s.Type = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) SetValue(v string) *UpsertQualityRuleRequestUpsertCommandValidateConditionList {
	s.Value = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommandValidateConditionList) Validate() error {
	return dara.Validate(s)
}

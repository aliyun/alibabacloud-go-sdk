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
	SetOpUserId(v string) *UpsertQualityRuleRequest
	GetOpUserId() *string
	SetUpsertCommand(v *UpsertQualityRuleRequestUpsertCommand) *UpsertQualityRuleRequest
	GetUpsertCommand() *UpsertQualityRuleRequestUpsertCommand
}

type UpsertQualityRuleRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update command.
	//
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

func (s *UpsertQualityRuleRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UpsertQualityRuleRequest) GetUpsertCommand() *UpsertQualityRuleRequestUpsertCommand {
	return s.UpsertCommand
}

func (s *UpsertQualityRuleRequest) SetOpTenantId(v int64) *UpsertQualityRuleRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpsertQualityRuleRequest) SetOpUserId(v string) *UpsertQualityRuleRequest {
	s.OpUserId = &v
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
	// The exception archive mode. Valid values:
	//
	// - ONLY_ERROR_FIELD: Archives only the exception fields.
	//
	// - FULL_RECORD: Archives the complete record.
	//
	// Default value: ONLY_ERROR_FIELD.
	//
	// example:
	//
	// ONLY_ERROR_FIELD
	ArchiveMode *string `json:"ArchiveMode,omitempty" xml:"ArchiveMode,omitempty"`
	// The exception archive storage type. Valid values:
	//
	// - FILE_SYSTEM: File system.
	//
	// - CUSTOM_TABLE: Custom table.
	//
	// Default value: FILE_SYSTEM.
	//
	// example:
	//
	// FILE_SYSTEM
	ArchiveStoreType *string `json:"ArchiveStoreType,omitempty" xml:"ArchiveStoreType,omitempty"`
	// The rule business property configuration.
	AttributeWithValueList []*UpsertQualityRuleRequestUpsertCommandAttributeWithValueList `json:"AttributeWithValueList,omitempty" xml:"AttributeWithValueList,omitempty" type:"Repeated"`
	// The rule catalog. Valid values:
	//
	// - CONSISTENT: consistency.
	//
	// - EFFECTIVE: validity.
	//
	// - TIMELINESE: timeliness.
	//
	// - ACCURATE: accuracy.
	//
	// - UNIQUENESS: uniqueness.
	//
	// - COMPLETENESS: completeness.
	//
	// - STABILITY: stability.
	//
	// - CUSTOM: custom.
	//
	// This parameter is required.
	CatalogList []*string `json:"CatalogList,omitempty" xml:"CatalogList,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable error archiving.
	EnableErrorArchive *bool `json:"EnableErrorArchive,omitempty" xml:"EnableErrorArchive,omitempty"`
	// The rule configuration key-value pairs. These are related to the templatetype. Different template types return different form key-value pair configurations.
	FormPropertyList []*UpsertQualityRuleRequestUpsertCommandFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
	// The rule ID. If this parameter is not empty, the operation updates the rule. If this parameter is empty, the operation creates a rule.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the quality rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The rule strength. Valid values:
	//
	// - STRONG
	//
	// - WEAK
	//
	// This parameter is required.
	//
	// example:
	//
	// STRONG
	Strength *string `json:"Strength,omitempty" xml:"Strength,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The templatetype. Valid values:
	//
	// - FIELD_NULL_VALUE_VALIDATE: field null value check.
	//
	// - FIELD_EMPTY_STRING_VALIDATE: field empty string check.
	//
	// - FIELD_UNIQUE_VALIDATE: field uniqueness check.
	//
	// - FIELD_GROUP_COUNT_VALIDATE: field unique value count check.
	//
	// - FIELD_DUPLICATE_VALUE_COUNT_VALIDATE: field duplicate value count check.
	//
	// - FUNCTION_TIME_COMPARE: time function comparison.
	//
	// - SINGLE_TABLE_TIME_COMPARE: single-table time field comparison.
	//
	// - DOUBLE_TABLE_TIME_COMPARE: two-table time field comparison.
	//
	// - FIELD_FORMAT_VALIDATE: field format check.
	//
	// - FIELD_LENGTH_VALIDATE: field length check.
	//
	// - FIELD_VALUE_RANGE_VALIDATE: field value range check.
	//
	// - CODE_TABLE_COMPARE: lookup table reference comparison.
	//
	// - STANDARD_CODE_TABLE_COMPARE: data standard lookup table reference comparison.
	//
	// - SINGLE_TABLE_FIELD_VALUE_COMPARE: single-table field value consistency comparison.
	//
	// - SINGLE_TABLE_FIELD_STATISTICAL_COMPARE: single-table field statistical value consistency comparison.
	//
	// - SINGLE_TABLE_FIELD_EXP_COMPARE: single-table field business logic consistency comparison.
	//
	// - DOUBLE_TABLE_FIELD_VALUE_COMPARE: two-table field value consistency comparison.
	//
	// - DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: two-table field statistical value consistency comparison.
	//
	// - CROSS_DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: cross-source two-table field statistical value consistency comparison.
	//
	// - DOUBLE_TABLE_FIELD_EXP_COMPARE: two-table field business logic consistency comparison.
	//
	// - TABLE_STABILITY_VALIDATE: table stability check.
	//
	// - TABLE_FLUCTUATION_VALIDATE: table fluctuation check.
	//
	// - FIELD_STABILITY_VALIDATE: field stability check.
	//
	// - FIELD_FLUCTUATION_VALIDATE: field fluctuation check.
	//
	// - CUSTOM_STATISTICAL_VALIDATE: custom statistical metric check.
	//
	// - CUSTOM_DATA_DETAILS_VALIDATE: custom data details check.
	//
	// - DATASOURCE_AVAILABLE_CHECK: datasource connectivity monitoring.
	//
	// - TABLE_SCHEMA_CHECK: table schema change monitoring.
	//
	// - REAL_TIME_OFFLINE_COMPARE: real-time and offline comparison.
	//
	// - REAL_TIME_STATISTICAL_VALIDATE: real-time statistical value monitoring.
	//
	// - REAL_TIME_MULTI_CHAIN_COMPARE: real-time multi-link comparison.
	//
	// This parameter is required.
	//
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The validation conditions.
	ValidateConditionList []*UpsertQualityRuleRequestUpsertCommandValidateConditionList `json:"ValidateConditionList,omitempty" xml:"ValidateConditionList,omitempty" type:"Repeated"`
	// The ID of the associated watch.
	//
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

func (s *UpsertQualityRuleRequestUpsertCommand) GetArchiveMode() *string {
	return s.ArchiveMode
}

func (s *UpsertQualityRuleRequestUpsertCommand) GetArchiveStoreType() *string {
	return s.ArchiveStoreType
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

func (s *UpsertQualityRuleRequestUpsertCommand) SetArchiveMode(v string) *UpsertQualityRuleRequestUpsertCommand {
	s.ArchiveMode = &v
	return s
}

func (s *UpsertQualityRuleRequestUpsertCommand) SetArchiveStoreType(v string) *UpsertQualityRuleRequestUpsertCommand {
	s.ArchiveStoreType = &v
	return s
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
	// The property details.
	AttributeInfo *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty" type:"Struct"`
	// The property value.
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
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the property is enabled.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The property ID.
	//
	// example:
	//
	// 711484689131
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The property name.
	//
	// example:
	//
	// attr01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the property is required.
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// Indicates whether the property is searchable.
	Searchable *bool `json:"Searchable,omitempty" xml:"Searchable,omitempty"`
	// The property value configuration details.
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
	// The property field data type. Valid values:
	//
	// - STRING: text.
	//
	// - BIGINT: integer.
	//
	// - DOUBLE: floating-point.
	//
	// - BOOLEAN: Boolean.
	//
	// - DATE: date.
	//
	// - DATETIME: datetime.
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The property default value.
	DefaultValue *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty" type:"Struct"`
	// The property field length. You can use this parameter to constrain the maximum length of text-type property values.
	//
	// example:
	//
	// 986992
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The property value input method. Valid values:
	//
	// - CUSTOMIZED: custom input.
	//
	// - SINGLE_ENUM: single-select dropdown.
	//
	// - MULTIPLE_ENUMS: multi-select dropdown.
	//
	// - RANGE: range interval.
	//
	// example:
	//
	// CUSTOMIZED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The property option values. This parameter applies only to properties whose input method is single-select dropdown or multi-select dropdown.
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
	// Indicates whether the maximum value is included.
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	// Indicates whether the minimum value is included.
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// The maximum value. This parameter applies to range interval properties.
	//
	// example:
	//
	// 11
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value. This parameter applies to range interval properties.
	//
	// example:
	//
	// 1
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The property value list. This parameter applies to properties whose input method is custom input, single-select dropdown, or multi-select dropdown.
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
	// Indicates whether the maximum value is included.
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	// Indicates whether the minimum value is included.
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// The maximum value. This parameter applies to range interval properties.
	//
	// example:
	//
	// 11
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value. This parameter applies to range interval properties.
	//
	// example:
	//
	// 1
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The property value list. This parameter applies to properties whose input method is custom input, single-select dropdown, or multi-select dropdown.
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
	// The control type.
	//
	// example:
	//
	// expression
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The property name.
	//
	// example:
	//
	// col
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The property value.
	//
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
	// The condition node ID.
	//
	// example:
	//
	// 268
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metric.
	//
	// example:
	//
	// test
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The operator. Valid values:
	//
	// - EQUAL
	//
	// - NOT_EQUAL
	//
	// - LARGER
	//
	// - SMALLER
	//
	// - LARGE_OR_EQUAL
	//
	// - SMALLER_OR_EQUAL
	//
	// - AND
	//
	// - OR
	//
	// example:
	//
	// AND
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The parent condition node ID.
	//
	// example:
	//
	// 123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The condition type. Valid values:
	//
	// - RELATION: relationship.
	//
	// - EXPRESSION: expression.
	//
	// example:
	//
	// RELATION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value.
	//
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

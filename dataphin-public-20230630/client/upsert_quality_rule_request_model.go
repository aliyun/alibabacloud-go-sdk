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
	// Tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The upsert command.
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
	// The rule business attribute configuration.
	AttributeWithValueList []*UpsertQualityRuleRequestUpsertCommandAttributeWithValueList `json:"AttributeWithValueList,omitempty" xml:"AttributeWithValueList,omitempty" type:"Repeated"`
	// The rule category. Valid values: CONSISTENT (Consistency), EFFECTIVE (Effectiveness), TIMELINESE (Timeliness), ACCURATE (Accuracy), UNIQUENESS (Uniqueness), COMPLETENESS (Completeness), STABILITY (Stability), CUSTOM (Custom).
	//
	// This parameter is required.
	CatalogList []*string `json:"CatalogList,omitempty" xml:"CatalogList,omitempty" type:"Repeated"`
	// The description of the quality rule.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable error archiving.
	EnableErrorArchive *bool `json:"EnableErrorArchive,omitempty" xml:"EnableErrorArchive,omitempty"`
	// The rule configuration key-value pairs. The configuration varies based on the template type. Different template types return different form key-value pair configurations.
	FormPropertyList []*UpsertQualityRuleRequestUpsertCommandFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
	// Rule ID. A non-empty value indicates a modification, and an empty value indicates a creation.
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
	// The rule strength. Valid values: STRONG, WEAK.
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
	// The template type. Valid values:
	//
	//   - FIELD_NULL_VALUE_VALIDATE: Field null value validation
	//
	//   - FIELD_EMPTY_STRING_VALIDATE: Field empty string validation
	//
	//   - FIELD_UNIQUE_VALIDATE: Field uniqueness validation
	//
	//   - FIELD_GROUP_COUNT_VALIDATE: Field unique value count validation
	//
	//   - FIELD_DUPLICATE_VALUE_COUNT_VALIDATE: Field duplicate value count validation
	//
	//   - FUNCTION_TIME_COMPARE: Time function comparison
	//
	//   - SINGLE_TABLE_TIME_COMPARE: Single-table time field comparison
	//
	//   - DOUBLE_TABLE_TIME_COMPARE: Cross-table time field comparison
	//
	//   - FIELD_FORMAT_VALIDATE: Field format validation
	//
	//   - FIELD_LENGTH_VALIDATE: Field length validation
	//
	//   - FIELD_VALUE_RANGE_VALIDATE: Field value range validation
	//
	//   - CODE_TABLE_COMPARE: Code table reference comparison
	//
	//   - STANDARD_CODE_TABLE_COMPARE: Data standard code table reference comparison
	//
	//   - SINGLE_TABLE_FIELD_VALUE_COMPARE: Single-table field value consistency comparison
	//
	//   - SINGLE_TABLE_FIELD_STATISTICAL_COMPARE: Single-table field statistical value consistency comparison
	//
	//   - SINGLE_TABLE_FIELD_EXP_COMPARE: Single-table field business logic consistency comparison
	//
	//   - DOUBLE_TABLE_FIELD_VALUE_COMPARE: Cross-table field value consistency comparison
	//
	//   - DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: Cross-table field statistical value consistency comparison
	//
	//   - CROSS_DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: Cross-source cross-table field statistical value consistency comparison
	//
	//   - DOUBLE_TABLE_FIELD_EXP_COMPARE: Cross-table field business logic consistency comparison
	//
	//   - TABLE_STABILITY_VALIDATE: Table stability validation
	//
	//   - TABLE_FLUCTUATION_VALIDATE: Table fluctuation validation
	//
	//   - FIELD_STABILITY_VALIDATE: Field stability validation
	//
	//   - FIELD_FLUCTUATION_VALIDATE: Field fluctuation validation
	//
	//   - CUSTOM_STATISTICAL_VALIDATE: Custom statistical metric validation
	//
	//   - CUSTOM_DATA_DETAILS_VALIDATE: Custom data details validation
	//
	//   - DATASOURCE_AVAILABLE_CHECK: Data source connectivity check
	//
	//   - TABLE_SCHEMA_CHECK: Table schema change monitoring
	//
	//   - REAL_TIME_OFFLINE_COMPARE: Real-time offline comparison
	//
	//   - REAL_TIME_STATISTICAL_VALIDATE: Real-time statistical value monitoring
	//
	//   - REAL_TIME_MULTI_CHAIN_COMPARE: Real-time multi-chain comparison, etc.
	//
	// This parameter is required.
	//
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The validation conditions.
	ValidateConditionList []*UpsertQualityRuleRequestUpsertCommandValidateConditionList `json:"ValidateConditionList,omitempty" xml:"ValidateConditionList,omitempty" type:"Repeated"`
	// The ID of the associated monitor.
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
	// The attribute details.
	AttributeInfo *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfo `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty" type:"Struct"`
	// The attribute value.
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
	// Specifies whether to enable the attribute.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The attribute ID.
	//
	// example:
	//
	// 711484689131
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The attribute name.
	//
	// example:
	//
	// attr01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether the attribute is required.
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// Specifies whether the attribute is searchable.
	Searchable *bool `json:"Searchable,omitempty" xml:"Searchable,omitempty"`
	// The attribute value configuration details.
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
	// The attribute field type. Valid values: STRING (Text), BIGINT (Integer), DOUBLE (Floating-point), BOOLEAN (Boolean), DATE (Date), DATETIME (Datetime).
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The attribute default value.
	DefaultValue *UpsertQualityRuleRequestUpsertCommandAttributeWithValueListAttributeInfoValueConfigDefaultValue `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty" type:"Struct"`
	// The attribute field length. Used to constrain the maximum length of text-type attribute values.
	//
	// example:
	//
	// 986992
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The attribute value input method. Valid values: CUSTOMIZED (Custom input), SINGLE_ENUM (Single-select dropdown), MULTIPLE_ENUMS (Multi-select dropdown), RANGE (Range interval).
	//
	// example:
	//
	// CUSTOMIZED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The attribute option values. Only applicable to attributes with the single-select dropdown or multi-select dropdown input method.
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
	// Specifies whether to include the maximum value.
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	// Specifies whether to include the minimum value.
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// The maximum value. Applicable to range interval attributes.
	//
	// example:
	//
	// 11
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value. Applicable to range interval attributes.
	//
	// example:
	//
	// 1
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The attribute value list. Applicable to attributes with the custom input, single-select dropdown, or multi-select dropdown input method.
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
	// Specifies whether to include the maximum value.
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	// Specifies whether to include the minimum value.
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// The maximum value. Applicable to range interval attributes.
	//
	// example:
	//
	// 11
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value. Applicable to range interval attributes.
	//
	// example:
	//
	// 1
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The attribute value list. Applicable to attributes with the custom input, single-select dropdown, or multi-select dropdown input method.
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
	// The component type.
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
	// The ID of the condition node.
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
	// The operator. Valid values: EQUAL, NOT_EQUAL, LARGER, SMALLER, LARGE_OR_EQUAL, SMALLER_OR_EQUAL, AND, OR.
	//
	// example:
	//
	// AND
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The ID of the parent condition node.
	//
	// example:
	//
	// 123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The condition type. Valid values: RELATION, EXPRESSION.
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

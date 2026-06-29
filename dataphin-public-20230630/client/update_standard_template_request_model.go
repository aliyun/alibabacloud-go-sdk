// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStandardTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UpdateStandardTemplateRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *UpdateStandardTemplateRequestUpdateCommand) *UpdateStandardTemplateRequest
	GetUpdateCommand() *UpdateStandardTemplateRequestUpdateCommand
}

type UpdateStandardTemplateRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The update command.
	//
	// This parameter is required.
	UpdateCommand *UpdateStandardTemplateRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s UpdateStandardTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateStandardTemplateRequest) GetUpdateCommand() *UpdateStandardTemplateRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *UpdateStandardTemplateRequest) SetOpTenantId(v int64) *UpdateStandardTemplateRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateStandardTemplateRequest) SetUpdateCommand(v *UpdateStandardTemplateRequestUpdateCommand) *UpdateStandardTemplateRequest {
	s.UpdateCommand = v
	return s
}

func (s *UpdateStandardTemplateRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommand struct {
	// The attribute configuration.
	//
	// This parameter is required.
	AttributesConfig *UpdateStandardTemplateRequestUpdateCommandAttributesConfig `json:"AttributesConfig,omitempty" xml:"AttributesConfig,omitempty" type:"Struct"`
	// The code of the standard template. The code must be globally unique. The code cannot be modified if the template is referenced.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration for automatic generation of standard codes.
	CodeRuleConfig *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig `json:"CodeRuleConfig,omitempty" xml:"CodeRuleConfig,omitempty" type:"Struct"`
	// The description of the standard template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the standard template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The list of maintainers.
	MaintainerList []*string `json:"MaintainerList,omitempty" xml:"MaintainerList,omitempty" type:"Repeated"`
	// The name of the standard template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The publish information of the standard template.
	PublishInfo *UpdateStandardTemplateRequestUpdateCommandPublishInfo `json:"PublishInfo,omitempty" xml:"PublishInfo,omitempty" type:"Struct"`
	// The version number. If this parameter is left empty or set to -1, the latest version is used.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommand) GetAttributesConfig() *UpdateStandardTemplateRequestUpdateCommandAttributesConfig {
	return s.AttributesConfig
}

func (s *UpdateStandardTemplateRequestUpdateCommand) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardTemplateRequestUpdateCommand) GetCodeRuleConfig() *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig {
	return s.CodeRuleConfig
}

func (s *UpdateStandardTemplateRequestUpdateCommand) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardTemplateRequestUpdateCommand) GetId() *int64 {
	return s.Id
}

func (s *UpdateStandardTemplateRequestUpdateCommand) GetMaintainerList() []*string {
	return s.MaintainerList
}

func (s *UpdateStandardTemplateRequestUpdateCommand) GetName() *string {
	return s.Name
}

func (s *UpdateStandardTemplateRequestUpdateCommand) GetPublishInfo() *UpdateStandardTemplateRequestUpdateCommandPublishInfo {
	return s.PublishInfo
}

func (s *UpdateStandardTemplateRequestUpdateCommand) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateStandardTemplateRequestUpdateCommand) SetAttributesConfig(v *UpdateStandardTemplateRequestUpdateCommandAttributesConfig) *UpdateStandardTemplateRequestUpdateCommand {
	s.AttributesConfig = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommand) SetCode(v string) *UpdateStandardTemplateRequestUpdateCommand {
	s.Code = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommand) SetCodeRuleConfig(v *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig) *UpdateStandardTemplateRequestUpdateCommand {
	s.CodeRuleConfig = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommand) SetDescription(v string) *UpdateStandardTemplateRequestUpdateCommand {
	s.Description = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommand) SetId(v int64) *UpdateStandardTemplateRequestUpdateCommand {
	s.Id = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommand) SetMaintainerList(v []*string) *UpdateStandardTemplateRequestUpdateCommand {
	s.MaintainerList = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommand) SetName(v string) *UpdateStandardTemplateRequestUpdateCommand {
	s.Name = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommand) SetPublishInfo(v *UpdateStandardTemplateRequestUpdateCommandPublishInfo) *UpdateStandardTemplateRequestUpdateCommand {
	s.PublishInfo = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommand) SetVersion(v int32) *UpdateStandardTemplateRequestUpdateCommand {
	s.Version = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommand) Validate() error {
	if s.AttributesConfig != nil {
		if err := s.AttributesConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CodeRuleConfig != nil {
		if err := s.CodeRuleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PublishInfo != nil {
		if err := s.PublishInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandAttributesConfig struct {
	// The list of attributes.
	//
	// This parameter is required.
	AttributeList []*UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Repeated"`
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfig) GetAttributeList() []*UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList {
	return s.AttributeList
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfig) SetAttributeList(v []*UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) *UpdateStandardTemplateRequestUpdateCommandAttributesConfig {
	s.AttributeList = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfig) Validate() error {
	if s.AttributeList != nil {
		for _, item := range s.AttributeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList struct {
	// The attribute code. This parameter is optional when a common attribute is referenced.
	//
	// example:
	//
	// test_attr
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The attribute name. This parameter is optional when a common attribute is referenced.
	//
	// example:
	//
	// attr1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The referenced attribute information.
	RefAttribute *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute `json:"RefAttribute,omitempty" xml:"RefAttribute,omitempty" type:"Struct"`
	// Specifies whether the attribute is required. This parameter is optional when a common attribute is referenced.
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// The attribute type. This parameter is optional when a common attribute is referenced. Valid values:
	//
	// - BIZ_ATTRIBUTE: business attribute.
	//
	// - TECH_ATTRIBUTE: technical attribute.
	//
	// - MANAGEMENT_ATTRIBUTE: management attribute.
	//
	// - QUALITY_ATTRIBUTE: quality attribute.
	//
	// - MASTER_DATA_ATTRIBUTE: master data attribute.
	//
	// - LIFECYCLE_ATTRIBUTE: lifecycle attribute.
	//
	// - SECURITY_ATTRIBUTE: security attribute.
	//
	// example:
	//
	// BIZ_ATTRIBUTE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value configuration. This parameter is optional when a common attribute is referenced.
	ValueConfig *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig `json:"ValueConfig,omitempty" xml:"ValueConfig,omitempty" type:"Struct"`
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) GetCode() *string {
	return s.Code
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) GetDescription() *string {
	return s.Description
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) GetName() *string {
	return s.Name
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) GetRefAttribute() *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute {
	return s.RefAttribute
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) GetRequired() *bool {
	return s.Required
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) GetType() *string {
	return s.Type
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) GetValueConfig() *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig {
	return s.ValueConfig
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) SetCode(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList {
	s.Code = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) SetDescription(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList {
	s.Description = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) SetName(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList {
	s.Name = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) SetRefAttribute(v *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList {
	s.RefAttribute = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) SetRequired(v bool) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList {
	s.Required = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) SetType(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList {
	s.Type = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) SetValueConfig(v *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList {
	s.ValueConfig = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeList) Validate() error {
	if s.RefAttribute != nil {
		if err := s.RefAttribute.Validate(); err != nil {
			return err
		}
	}
	if s.ValueConfig != nil {
		if err := s.ValueConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute struct {
	// The attribute source.
	//
	// This parameter is required.
	AttributeFromInfo *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo `json:"AttributeFromInfo,omitempty" xml:"AttributeFromInfo,omitempty" type:"Struct"`
	// The attribute ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	AttributeId *int64 `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute) GetAttributeFromInfo() *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo {
	return s.AttributeFromInfo
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute) SetAttributeFromInfo(v *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute {
	s.AttributeFromInfo = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute) SetAttributeId(v int64) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute {
	s.AttributeId = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute) Validate() error {
	if s.AttributeFromInfo != nil {
		if err := s.AttributeFromInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo struct {
	// The attribute source. Valid values:
	//
	// - SYSTEM: system attribute.
	//
	// - CUSTOM: custom attribute.
	//
	// - STANDARD: standard.
	//
	// This parameter is required.
	//
	// example:
	//
	// CUSTOM
	AttributeFrom *string `json:"AttributeFrom,omitempty" xml:"AttributeFrom,omitempty"`
	// The corresponding standard. This parameter takes effect when the attribute source is set to STANDARD.
	StandardReference *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference `json:"StandardReference,omitempty" xml:"StandardReference,omitempty" type:"Struct"`
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) GetAttributeFrom() *string {
	return s.AttributeFrom
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) GetStandardReference() *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference {
	return s.StandardReference
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) SetAttributeFrom(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo {
	s.AttributeFrom = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) SetStandardReference(v *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo {
	s.StandardReference = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) Validate() error {
	if s.StandardReference != nil {
		if err := s.StandardReference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference struct {
	// The standard ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// The version number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) GetStandardId() *int64 {
	return s.StandardId
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) SetStandardId(v int64) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference {
	s.StandardId = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) SetVersion(v int32) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference {
	s.Version = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig struct {
	// The data type of the attribute value. Valid values:
	//
	// - STRING: string.
	//
	// - BIGINT: numeric.
	//
	// - DOUBLE: floating-point.
	//
	// - DATE: date, accurate to the day.
	//
	// - DATETIME: date, accurate to milliseconds.
	//
	// - BOOLEAN: Boolean.
	//
	// This parameter is required.
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The default value.
	//
	// example:
	//
	// test
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The length of the attribute value. If this parameter is left empty or set to -1, the length is not limited. Typically, only string types have a length limit for attribute values.
	//
	// example:
	//
	// 1
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The attribute value type. Valid values:
	//
	// - CUSTOMIZED: custom input.
	//
	// - SINGLE_ENUM: single enumeration value.
	//
	// - MULTIPLE_ENUMS: multiple enumeration values.
	//
	// - RANGE: range value.
	//
	// This parameter is required.
	//
	// example:
	//
	// CUSTOMIZED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value range.
	//
	// This parameter is required.
	ValueRange *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange `json:"ValueRange,omitempty" xml:"ValueRange,omitempty" type:"Struct"`
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) GetDataType() *string {
	return s.DataType
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) GetLength() *int32 {
	return s.Length
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) GetType() *string {
	return s.Type
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) GetValueRange() *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange {
	return s.ValueRange
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) SetDataType(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig {
	s.DataType = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) SetDefaultValue(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig {
	s.DefaultValue = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) SetLength(v int32) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig {
	s.Length = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) SetType(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig {
	s.Type = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) SetValueRange(v *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig {
	s.ValueRange = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfig) Validate() error {
	if s.ValueRange != nil {
		if err := s.ValueRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange struct {
	// The value range. This parameter takes effect when the value source is set to DATAPHIN_ATTRIBUTE. Valid values:
	//
	// - BIZ_UNIT: data board.
	//
	// - PROJECT: project.
	//
	// - USER: user.
	//
	// - USER_GROUP: user group.
	//
	// example:
	//
	// DATAPHIN_ATTRIBUTE
	DataphinAttributeType *string `json:"DataphinAttributeType,omitempty" xml:"DataphinAttributeType,omitempty"`
	// The value range. This parameter takes effect when the value source is set to LOOKUP_TABLE.
	LookupTableReference *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference `json:"LookupTableReference,omitempty" xml:"LookupTableReference,omitempty" type:"Struct"`
	// The value range. This parameter takes effect when the value source is set to MIN_MAX.
	MinMaxValueConfig *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig `json:"MinMaxValueConfig,omitempty" xml:"MinMaxValueConfig,omitempty" type:"Struct"`
	// The value source. Valid values:
	//
	// - NONE: no constraint.
	//
	// - LIST: obtained from a list.
	//
	// - LOOKUP_TABLE: lookup table.
	//
	// - MIN_MAX: value between the minimum and maximum.
	//
	// - DATAPHIN_ATTRIBUTE: Dataphin system property.
	//
	// - BUILT_IN_DATA_TYPES: built-in data types.
	//
	// - BUILT_IN_DATA_CLASSIFICATION: built-in data categorization.
	//
	// - BUILT_IN_DATA_LEVEL: built-in data security classification.
	//
	// This parameter is required.
	//
	// example:
	//
	// NONE
	ValueConstraint *string `json:"ValueConstraint,omitempty" xml:"ValueConstraint,omitempty"`
	// The value range. This parameter takes effect when the value source is set to LIST.
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) GetDataphinAttributeType() *string {
	return s.DataphinAttributeType
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) GetLookupTableReference() *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference {
	return s.LookupTableReference
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) GetMinMaxValueConfig() *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	return s.MinMaxValueConfig
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) GetValueConstraint() *string {
	return s.ValueConstraint
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) GetValueList() []*string {
	return s.ValueList
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) SetDataphinAttributeType(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.DataphinAttributeType = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) SetLookupTableReference(v *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.LookupTableReference = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) SetMinMaxValueConfig(v *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.MinMaxValueConfig = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) SetValueConstraint(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.ValueConstraint = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) SetValueList(v []*string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.ValueList = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRange) Validate() error {
	if s.LookupTableReference != nil {
		if err := s.LookupTableReference.Validate(); err != nil {
			return err
		}
	}
	if s.MinMaxValueConfig != nil {
		if err := s.MinMaxValueConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference struct {
	// The referenced lookup table field.
	//
	// example:
	//
	// col1
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The ID of the lookup table.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	LookupTableId *int64 `json:"LookupTableId,omitempty" xml:"LookupTableId,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) GetColumn() *string {
	return s.Column
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) GetLookupTableId() *int64 {
	return s.LookupTableId
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) SetColumn(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference {
	s.Column = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) SetLookupTableId(v int64) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference {
	s.LookupTableId = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig struct {
	// Specifies whether the maximum value is included.
	//
	// This parameter is required.
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	// Specifies whether the minimum value is included.
	//
	// This parameter is required.
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// The maximum value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetMaxValue() *string {
	return s.MaxValue
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetMinValue() *string {
	return s.MinValue
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetIncludeMaxValue(v bool) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMaxValue = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetIncludeMinValue(v bool) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMinValue = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetMaxValue(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.MaxValue = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetMinValue(v string) *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.MinValue = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig struct {
	// The automatic generation configuration for standard code rules. This parameter takes effect when the generation method is set to AUTO_GENERATE.
	AutoConfig *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty" type:"Struct"`
	// The standard code generation method. Valid values:
	//
	// - CUSTOMIZED: custom.
	//
	// - AUTO_GENERATE: automatically generated based on standard code rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// CUSTOMIZED
	GenerateType *string `json:"GenerateType,omitempty" xml:"GenerateType,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig) GetAutoConfig() *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig {
	return s.AutoConfig
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig) GetGenerateType() *string {
	return s.GenerateType
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig) SetAutoConfig(v *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig {
	s.AutoConfig = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig) SetGenerateType(v string) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig {
	s.GenerateType = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig) Validate() error {
	if s.AutoConfig != nil {
		if err := s.AutoConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig struct {
	// The standard code rules.
	//
	// This parameter is required.
	CodeRuleList []*UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList `json:"CodeRuleList,omitempty" xml:"CodeRuleList,omitempty" type:"Repeated"`
	// Specifies whether strict validation is required.
	//
	// This parameter is required.
	NeedStrongValidate *bool `json:"NeedStrongValidate,omitempty" xml:"NeedStrongValidate,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig) GetCodeRuleList() []*UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList {
	return s.CodeRuleList
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig) GetNeedStrongValidate() *bool {
	return s.NeedStrongValidate
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig) SetCodeRuleList(v []*UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig {
	s.CodeRuleList = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig) SetNeedStrongValidate(v bool) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig {
	s.NeedStrongValidate = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig) Validate() error {
	if s.CodeRuleList != nil {
		for _, item := range s.CodeRuleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList struct {
	// The auto-increment sequence configuration.
	AutoIncrementSequenceConfig *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig `json:"AutoIncrementSequenceConfig,omitempty" xml:"AutoIncrementSequenceConfig,omitempty" type:"Struct"`
	// The position index of the code rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The type of the code rule. Valid values:
	//
	// - FIXED_STRING: fixed string.
	//
	// - AUTO_INCREMENT: auto-increment sequence.
	//
	// - STANDARD_SET_CODE: standard set code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FIXED_STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The format or value of the code rule.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) GetAutoIncrementSequenceConfig() *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	return s.AutoIncrementSequenceConfig
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) GetIndex() *int32 {
	return s.Index
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) GetType() *string {
	return s.Type
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) GetValue() *string {
	return s.Value
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) SetAutoIncrementSequenceConfig(v *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList {
	s.AutoIncrementSequenceConfig = v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) SetIndex(v int32) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList {
	s.Index = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) SetType(v string) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList {
	s.Type = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) SetValue(v string) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList {
	s.Value = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList) Validate() error {
	if s.AutoIncrementSequenceConfig != nil {
		if err := s.AutoIncrementSequenceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig struct {
	// The number of digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Digit *int32 `json:"Digit,omitempty" xml:"Digit,omitempty"`
	// Specifies whether to pad with zeros.
	//
	// This parameter is required.
	NeedPaddingZero *bool `json:"NeedPaddingZero,omitempty" xml:"NeedPaddingZero,omitempty"`
	// The start value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	StartValue *int64 `json:"StartValue,omitempty" xml:"StartValue,omitempty"`
	// The step size.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetDigit() *int32 {
	return s.Digit
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetNeedPaddingZero() *bool {
	return s.NeedPaddingZero
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetStartValue() *int64 {
	return s.StartValue
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetStep() *int32 {
	return s.Step
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetDigit(v int32) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.Digit = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetNeedPaddingZero(v bool) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.NeedPaddingZero = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetStartValue(v int64) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.StartValue = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetStep(v int32) *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.Step = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateStandardTemplateRequestUpdateCommandPublishInfo struct {
	// The publish comment.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s UpdateStandardTemplateRequestUpdateCommandPublishInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateStandardTemplateRequestUpdateCommandPublishInfo) GoString() string {
	return s.String()
}

func (s *UpdateStandardTemplateRequestUpdateCommandPublishInfo) GetComment() *string {
	return s.Comment
}

func (s *UpdateStandardTemplateRequestUpdateCommandPublishInfo) SetComment(v string) *UpdateStandardTemplateRequestUpdateCommandPublishInfo {
	s.Comment = &v
	return s
}

func (s *UpdateStandardTemplateRequestUpdateCommandPublishInfo) Validate() error {
	return dara.Validate(s)
}

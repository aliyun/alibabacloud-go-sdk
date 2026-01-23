// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStandardTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateStandardTemplateRequestCreateCommand) *CreateStandardTemplateRequest
	GetCreateCommand() *CreateStandardTemplateRequestCreateCommand
	SetOpTenantId(v int64) *CreateStandardTemplateRequest
	GetOpTenantId() *int64
}

type CreateStandardTemplateRequest struct {
	// This parameter is required.
	CreateCommand *CreateStandardTemplateRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateStandardTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequest) GetCreateCommand() *CreateStandardTemplateRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateStandardTemplateRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateStandardTemplateRequest) SetCreateCommand(v *CreateStandardTemplateRequestCreateCommand) *CreateStandardTemplateRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateStandardTemplateRequest) SetOpTenantId(v int64) *CreateStandardTemplateRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateStandardTemplateRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardTemplateRequestCreateCommand struct {
	// This parameter is required.
	AttributesConfig *CreateStandardTemplateRequestCreateCommandAttributesConfig `json:"AttributesConfig,omitempty" xml:"AttributesConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test01
	Code           *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	CodeRuleConfig *CreateStandardTemplateRequestCreateCommandCodeRuleConfig `json:"CodeRuleConfig,omitempty" xml:"CodeRuleConfig,omitempty" type:"Struct"`
	// example:
	//
	// test
	Description    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	MaintainerList []*string `json:"MaintainerList,omitempty" xml:"MaintainerList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 测试模板
	Name        *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	PublishInfo *CreateStandardTemplateRequestCreateCommandPublishInfo `json:"PublishInfo,omitempty" xml:"PublishInfo,omitempty" type:"Struct"`
}

func (s CreateStandardTemplateRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommand) GetAttributesConfig() *CreateStandardTemplateRequestCreateCommandAttributesConfig {
	return s.AttributesConfig
}

func (s *CreateStandardTemplateRequestCreateCommand) GetCode() *string {
	return s.Code
}

func (s *CreateStandardTemplateRequestCreateCommand) GetCodeRuleConfig() *CreateStandardTemplateRequestCreateCommandCodeRuleConfig {
	return s.CodeRuleConfig
}

func (s *CreateStandardTemplateRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardTemplateRequestCreateCommand) GetMaintainerList() []*string {
	return s.MaintainerList
}

func (s *CreateStandardTemplateRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateStandardTemplateRequestCreateCommand) GetPublishInfo() *CreateStandardTemplateRequestCreateCommandPublishInfo {
	return s.PublishInfo
}

func (s *CreateStandardTemplateRequestCreateCommand) SetAttributesConfig(v *CreateStandardTemplateRequestCreateCommandAttributesConfig) *CreateStandardTemplateRequestCreateCommand {
	s.AttributesConfig = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommand) SetCode(v string) *CreateStandardTemplateRequestCreateCommand {
	s.Code = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommand) SetCodeRuleConfig(v *CreateStandardTemplateRequestCreateCommandCodeRuleConfig) *CreateStandardTemplateRequestCreateCommand {
	s.CodeRuleConfig = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommand) SetDescription(v string) *CreateStandardTemplateRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommand) SetMaintainerList(v []*string) *CreateStandardTemplateRequestCreateCommand {
	s.MaintainerList = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommand) SetName(v string) *CreateStandardTemplateRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommand) SetPublishInfo(v *CreateStandardTemplateRequestCreateCommandPublishInfo) *CreateStandardTemplateRequestCreateCommand {
	s.PublishInfo = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommand) Validate() error {
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

type CreateStandardTemplateRequestCreateCommandAttributesConfig struct {
	// This parameter is required.
	AttributeList []*CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Repeated"`
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfig) GetAttributeList() []*CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList {
	return s.AttributeList
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfig) SetAttributeList(v []*CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) *CreateStandardTemplateRequestCreateCommandAttributesConfig {
	s.AttributeList = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfig) Validate() error {
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

type CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList struct {
	// example:
	//
	// test_attr
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// attr1
	Name         *string                                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	RefAttribute *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute `json:"RefAttribute,omitempty" xml:"RefAttribute,omitempty" type:"Struct"`
	Required     *bool                                                                                `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// BIZ_ATTRIBUTE
	Type        *string                                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	ValueConfig *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig `json:"ValueConfig,omitempty" xml:"ValueConfig,omitempty" type:"Struct"`
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) GetCode() *string {
	return s.Code
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) GetDescription() *string {
	return s.Description
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) GetName() *string {
	return s.Name
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) GetRefAttribute() *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute {
	return s.RefAttribute
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) GetRequired() *bool {
	return s.Required
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) GetType() *string {
	return s.Type
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) GetValueConfig() *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig {
	return s.ValueConfig
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) SetCode(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList {
	s.Code = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) SetDescription(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList {
	s.Description = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) SetName(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList {
	s.Name = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) SetRefAttribute(v *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList {
	s.RefAttribute = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) SetRequired(v bool) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList {
	s.Required = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) SetType(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList {
	s.Type = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) SetValueConfig(v *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList {
	s.ValueConfig = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeList) Validate() error {
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

type CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute struct {
	// This parameter is required.
	AttributeFromInfo *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo `json:"AttributeFromInfo,omitempty" xml:"AttributeFromInfo,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	AttributeId *int64 `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute) GetAttributeFromInfo() *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo {
	return s.AttributeFromInfo
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute) SetAttributeFromInfo(v *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute {
	s.AttributeFromInfo = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute) SetAttributeId(v int64) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute {
	s.AttributeId = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttribute) Validate() error {
	if s.AttributeFromInfo != nil {
		if err := s.AttributeFromInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// CUSTOM
	AttributeFrom     *string                                                                                                                `json:"AttributeFrom,omitempty" xml:"AttributeFrom,omitempty"`
	StandardReference *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference `json:"StandardReference,omitempty" xml:"StandardReference,omitempty" type:"Struct"`
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) GetAttributeFrom() *string {
	return s.AttributeFrom
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) GetStandardReference() *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference {
	return s.StandardReference
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) SetAttributeFrom(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo {
	s.AttributeFrom = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) SetStandardReference(v *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo {
	s.StandardReference = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo) Validate() error {
	if s.StandardReference != nil {
		if err := s.StandardReference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference struct {
	// This parameter is required.
	//
	// example:
	//
	// 22
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) GetStandardId() *int64 {
	return s.StandardId
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) GetVersion() *int32 {
	return s.Version
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) SetStandardId(v int64) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference {
	s.StandardId = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) SetVersion(v int32) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference {
	s.Version = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) Validate() error {
	return dara.Validate(s)
}

type CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// example:
	//
	// test
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// 1
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOMIZED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	ValueRange *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange `json:"ValueRange,omitempty" xml:"ValueRange,omitempty" type:"Struct"`
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) GetDataType() *string {
	return s.DataType
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) GetLength() *int32 {
	return s.Length
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) GetType() *string {
	return s.Type
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) GetValueRange() *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange {
	return s.ValueRange
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) SetDataType(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig {
	s.DataType = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) SetDefaultValue(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig {
	s.DefaultValue = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) SetLength(v int32) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig {
	s.Length = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) SetType(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig {
	s.Type = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) SetValueRange(v *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig {
	s.ValueRange = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfig) Validate() error {
	if s.ValueRange != nil {
		if err := s.ValueRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange struct {
	// example:
	//
	// DATAPHIN_ATTRIBUTE
	DataphinAttributeType *string                                                                                                           `json:"DataphinAttributeType,omitempty" xml:"DataphinAttributeType,omitempty"`
	LookupTableReference  *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference `json:"LookupTableReference,omitempty" xml:"LookupTableReference,omitempty" type:"Struct"`
	MinMaxValueConfig     *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig    `json:"MinMaxValueConfig,omitempty" xml:"MinMaxValueConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// NONE
	ValueConstraint *string   `json:"ValueConstraint,omitempty" xml:"ValueConstraint,omitempty"`
	ValueList       []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) GetDataphinAttributeType() *string {
	return s.DataphinAttributeType
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) GetLookupTableReference() *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference {
	return s.LookupTableReference
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) GetMinMaxValueConfig() *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	return s.MinMaxValueConfig
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) GetValueConstraint() *string {
	return s.ValueConstraint
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) GetValueList() []*string {
	return s.ValueList
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) SetDataphinAttributeType(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.DataphinAttributeType = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) SetLookupTableReference(v *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.LookupTableReference = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) SetMinMaxValueConfig(v *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.MinMaxValueConfig = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) SetValueConstraint(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.ValueConstraint = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) SetValueList(v []*string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange {
	s.ValueList = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRange) Validate() error {
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

type CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference struct {
	// example:
	//
	// col1
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	LookupTableId *int64 `json:"LookupTableId,omitempty" xml:"LookupTableId,omitempty"`
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) GetColumn() *string {
	return s.Column
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) GetLookupTableId() *int64 {
	return s.LookupTableId
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) SetColumn(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference {
	s.Column = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) SetLookupTableId(v int64) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference {
	s.LookupTableId = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) Validate() error {
	return dara.Validate(s)
}

type CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig struct {
	// This parameter is required.
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	// This parameter is required.
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetMaxValue() *string {
	return s.MaxValue
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetMinValue() *string {
	return s.MinValue
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetIncludeMaxValue(v bool) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMaxValue = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetIncludeMinValue(v bool) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMinValue = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetMaxValue(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.MaxValue = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetMinValue(v string) *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.MinValue = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) Validate() error {
	return dara.Validate(s)
}

type CreateStandardTemplateRequestCreateCommandCodeRuleConfig struct {
	AutoConfig *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// CUSTOMIZED
	GenerateType *string `json:"GenerateType,omitempty" xml:"GenerateType,omitempty"`
}

func (s CreateStandardTemplateRequestCreateCommandCodeRuleConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandCodeRuleConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfig) GetAutoConfig() *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig {
	return s.AutoConfig
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfig) GetGenerateType() *string {
	return s.GenerateType
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfig) SetAutoConfig(v *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig) *CreateStandardTemplateRequestCreateCommandCodeRuleConfig {
	s.AutoConfig = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfig) SetGenerateType(v string) *CreateStandardTemplateRequestCreateCommandCodeRuleConfig {
	s.GenerateType = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfig) Validate() error {
	if s.AutoConfig != nil {
		if err := s.AutoConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig struct {
	// This parameter is required.
	CodeRuleList []*CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList `json:"CodeRuleList,omitempty" xml:"CodeRuleList,omitempty" type:"Repeated"`
	// This parameter is required.
	NeedStrongValidate *bool `json:"NeedStrongValidate,omitempty" xml:"NeedStrongValidate,omitempty"`
}

func (s CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig) GetCodeRuleList() []*CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList {
	return s.CodeRuleList
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig) GetNeedStrongValidate() *bool {
	return s.NeedStrongValidate
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig) SetCodeRuleList(v []*CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig {
	s.CodeRuleList = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig) SetNeedStrongValidate(v bool) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig {
	s.NeedStrongValidate = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfig) Validate() error {
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

type CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList struct {
	AutoIncrementSequenceConfig *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig `json:"AutoIncrementSequenceConfig,omitempty" xml:"AutoIncrementSequenceConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FIXED_STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) GetAutoIncrementSequenceConfig() *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	return s.AutoIncrementSequenceConfig
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) GetIndex() *int32 {
	return s.Index
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) GetType() *string {
	return s.Type
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) GetValue() *string {
	return s.Value
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) SetAutoIncrementSequenceConfig(v *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList {
	s.AutoIncrementSequenceConfig = v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) SetIndex(v int32) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList {
	s.Index = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) SetType(v string) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList {
	s.Type = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) SetValue(v string) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList {
	s.Value = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleList) Validate() error {
	if s.AutoIncrementSequenceConfig != nil {
		if err := s.AutoIncrementSequenceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// 3
	Digit *int32 `json:"Digit,omitempty" xml:"Digit,omitempty"`
	// This parameter is required.
	NeedPaddingZero *bool `json:"NeedPaddingZero,omitempty" xml:"NeedPaddingZero,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	StartValue *int64 `json:"StartValue,omitempty" xml:"StartValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetDigit() *int32 {
	return s.Digit
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetNeedPaddingZero() *bool {
	return s.NeedPaddingZero
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetStartValue() *int64 {
	return s.StartValue
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetStep() *int32 {
	return s.Step
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetDigit(v int32) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.Digit = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetNeedPaddingZero(v bool) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.NeedPaddingZero = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetStartValue(v int64) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.StartValue = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetStep(v int32) *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.Step = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) Validate() error {
	return dara.Validate(s)
}

type CreateStandardTemplateRequestCreateCommandPublishInfo struct {
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s CreateStandardTemplateRequestCreateCommandPublishInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateStandardTemplateRequestCreateCommandPublishInfo) GoString() string {
	return s.String()
}

func (s *CreateStandardTemplateRequestCreateCommandPublishInfo) GetComment() *string {
	return s.Comment
}

func (s *CreateStandardTemplateRequestCreateCommandPublishInfo) SetComment(v string) *CreateStandardTemplateRequestCreateCommandPublishInfo {
	s.Comment = &v
	return s
}

func (s *CreateStandardTemplateRequestCreateCommandPublishInfo) Validate() error {
	return dara.Validate(s)
}

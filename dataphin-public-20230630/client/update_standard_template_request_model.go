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
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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
	// This parameter is required.
	AttributesConfig *UpdateStandardTemplateRequestUpdateCommandAttributesConfig `json:"AttributesConfig,omitempty" xml:"AttributesConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// test01
	Code           *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	CodeRuleConfig *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfig `json:"CodeRuleConfig,omitempty" xml:"CodeRuleConfig,omitempty" type:"Struct"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id             *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	MaintainerList []*string `json:"MaintainerList,omitempty" xml:"MaintainerList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 测试模板
	Name        *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	PublishInfo *UpdateStandardTemplateRequestUpdateCommandPublishInfo `json:"PublishInfo,omitempty" xml:"PublishInfo,omitempty" type:"Struct"`
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
	RefAttribute *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttribute `json:"RefAttribute,omitempty" xml:"RefAttribute,omitempty" type:"Struct"`
	Required     *bool                                                                                `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// BIZ_ATTRIBUTE
	Type        *string                                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// This parameter is required.
	AttributeFromInfo *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListRefAttributeAttributeFromInfo `json:"AttributeFromInfo,omitempty" xml:"AttributeFromInfo,omitempty" type:"Struct"`
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
	// This parameter is required.
	//
	// example:
	//
	// CUSTOM
	AttributeFrom     *string                                                                                                                `json:"AttributeFrom,omitempty" xml:"AttributeFrom,omitempty"`
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
	// example:
	//
	// DATAPHIN_ATTRIBUTE
	DataphinAttributeType *string                                                                                                           `json:"DataphinAttributeType,omitempty" xml:"DataphinAttributeType,omitempty"`
	LookupTableReference  *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeLookupTableReference `json:"LookupTableReference,omitempty" xml:"LookupTableReference,omitempty" type:"Struct"`
	MinMaxValueConfig     *UpdateStandardTemplateRequestUpdateCommandAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig    `json:"MinMaxValueConfig,omitempty" xml:"MinMaxValueConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// NONE
	ValueConstraint *string   `json:"ValueConstraint,omitempty" xml:"ValueConstraint,omitempty"`
	ValueList       []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
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
	AutoConfig *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfig `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty" type:"Struct"`
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
	// This parameter is required.
	CodeRuleList []*UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleList `json:"CodeRuleList,omitempty" xml:"CodeRuleList,omitempty" type:"Repeated"`
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
	AutoIncrementSequenceConfig *UpdateStandardTemplateRequestUpdateCommandCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig `json:"AutoIncrementSequenceConfig,omitempty" xml:"AutoIncrementSequenceConfig,omitempty" type:"Struct"`
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

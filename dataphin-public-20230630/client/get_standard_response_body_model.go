// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStandardResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetStandardResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStandardResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStandardResponseBody
	GetRequestId() *string
	SetStandardInfo(v *GetStandardResponseBodyStandardInfo) *GetStandardResponseBody
	GetStandardInfo() *GetStandardResponseBodyStandardInfo
	SetSuccess(v bool) *GetStandardResponseBody
	GetSuccess() *bool
}

type GetStandardResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The standard details.
	StandardInfo *GetStandardResponseBodyStandardInfo `json:"StandardInfo,omitempty" xml:"StandardInfo,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStandardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStandardResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStandardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStandardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStandardResponseBody) GetStandardInfo() *GetStandardResponseBodyStandardInfo {
	return s.StandardInfo
}

func (s *GetStandardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStandardResponseBody) SetCode(v string) *GetStandardResponseBody {
	s.Code = &v
	return s
}

func (s *GetStandardResponseBody) SetHttpStatusCode(v int32) *GetStandardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStandardResponseBody) SetMessage(v string) *GetStandardResponseBody {
	s.Message = &v
	return s
}

func (s *GetStandardResponseBody) SetRequestId(v string) *GetStandardResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardResponseBody) SetStandardInfo(v *GetStandardResponseBodyStandardInfo) *GetStandardResponseBody {
	s.StandardInfo = v
	return s
}

func (s *GetStandardResponseBody) SetSuccess(v bool) *GetStandardResponseBody {
	s.Success = &v
	return s
}

func (s *GetStandardResponseBody) Validate() error {
	if s.StandardInfo != nil {
		if err := s.StandardInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardResponseBodyStandardInfo struct {
	// The attribute value configurations.
	AttributeWithValueList []*GetStandardResponseBodyStandardInfoAttributeWithValueList `json:"AttributeWithValueList,omitempty" xml:"AttributeWithValueList,omitempty" type:"Repeated"`
	// The standard code.
	//
	// example:
	//
	// zz
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The creator.
	Creator *GetStandardResponseBodyStandardInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The effective period configuration.
	EffectiveTimeConfig *GetStandardResponseBodyStandardInfoEffectiveTimeConfig `json:"EffectiveTimeConfig,omitempty" xml:"EffectiveTimeConfig,omitempty" type:"Struct"`
	// The English name of the standard.
	//
	// example:
	//
	// test
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	// The lookup table.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last modifier.
	LastModifier *GetStandardResponseBodyStandardInfoLastModifier `json:"LastModifier,omitempty" xml:"LastModifier,omitempty" type:"Struct"`
	// The list of associated lookup tables.
	LookupTableRelations []*GetStandardResponseBodyStandardInfoLookupTableRelations `json:"LookupTableRelations,omitempty" xml:"LookupTableRelations,omitempty" type:"Repeated"`
	// The last modification time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The standard name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner.
	Owner *GetStandardResponseBodyStandardInfoOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// The stage to which the standard belongs.
	//
	// example:
	//
	// dev
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// The standard monitoring configuration.
	StandardGeneralMonitorConfig *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig `json:"StandardGeneralMonitorConfig,omitempty" xml:"StandardGeneralMonitorConfig,omitempty" type:"Struct"`
	// The list of associated standards.
	StandardRelations []*GetStandardResponseBodyStandardInfoStandardRelations `json:"StandardRelations,omitempty" xml:"StandardRelations,omitempty" type:"Repeated"`
	// The standard set on which the current standard depends.
	StandardSet *GetStandardResponseBodyStandardInfoStandardSet `json:"StandardSet,omitempty" xml:"StandardSet,omitempty" type:"Struct"`
	// The standard template on which the current standard depends.
	StandardTemplate *GetStandardResponseBodyStandardInfoStandardTemplate `json:"StandardTemplate,omitempty" xml:"StandardTemplate,omitempty" type:"Struct"`
	// The status of the standard.
	//
	// example:
	//
	// draft
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The standard type.
	//
	// example:
	//
	// zz
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetStandardResponseBodyStandardInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfo) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfo) GetAttributeWithValueList() []*GetStandardResponseBodyStandardInfoAttributeWithValueList {
	return s.AttributeWithValueList
}

func (s *GetStandardResponseBodyStandardInfo) GetCode() *string {
	return s.Code
}

func (s *GetStandardResponseBodyStandardInfo) GetCreator() *GetStandardResponseBodyStandardInfoCreator {
	return s.Creator
}

func (s *GetStandardResponseBodyStandardInfo) GetDescription() *string {
	return s.Description
}

func (s *GetStandardResponseBodyStandardInfo) GetEffectiveTimeConfig() *GetStandardResponseBodyStandardInfoEffectiveTimeConfig {
	return s.EffectiveTimeConfig
}

func (s *GetStandardResponseBodyStandardInfo) GetEnglishName() *string {
	return s.EnglishName
}

func (s *GetStandardResponseBodyStandardInfo) GetId() *int64 {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfo) GetLastModifier() *GetStandardResponseBodyStandardInfoLastModifier {
	return s.LastModifier
}

func (s *GetStandardResponseBodyStandardInfo) GetLookupTableRelations() []*GetStandardResponseBodyStandardInfoLookupTableRelations {
	return s.LookupTableRelations
}

func (s *GetStandardResponseBodyStandardInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetStandardResponseBodyStandardInfo) GetName() *string {
	return s.Name
}

func (s *GetStandardResponseBodyStandardInfo) GetOwner() *GetStandardResponseBodyStandardInfoOwner {
	return s.Owner
}

func (s *GetStandardResponseBodyStandardInfo) GetStage() *string {
	return s.Stage
}

func (s *GetStandardResponseBodyStandardInfo) GetStandardGeneralMonitorConfig() *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig {
	return s.StandardGeneralMonitorConfig
}

func (s *GetStandardResponseBodyStandardInfo) GetStandardRelations() []*GetStandardResponseBodyStandardInfoStandardRelations {
	return s.StandardRelations
}

func (s *GetStandardResponseBodyStandardInfo) GetStandardSet() *GetStandardResponseBodyStandardInfoStandardSet {
	return s.StandardSet
}

func (s *GetStandardResponseBodyStandardInfo) GetStandardTemplate() *GetStandardResponseBodyStandardInfoStandardTemplate {
	return s.StandardTemplate
}

func (s *GetStandardResponseBodyStandardInfo) GetStatus() *string {
	return s.Status
}

func (s *GetStandardResponseBodyStandardInfo) GetType() *string {
	return s.Type
}

func (s *GetStandardResponseBodyStandardInfo) GetVersion() *int32 {
	return s.Version
}

func (s *GetStandardResponseBodyStandardInfo) SetAttributeWithValueList(v []*GetStandardResponseBodyStandardInfoAttributeWithValueList) *GetStandardResponseBodyStandardInfo {
	s.AttributeWithValueList = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetCode(v string) *GetStandardResponseBodyStandardInfo {
	s.Code = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetCreator(v *GetStandardResponseBodyStandardInfoCreator) *GetStandardResponseBodyStandardInfo {
	s.Creator = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetDescription(v string) *GetStandardResponseBodyStandardInfo {
	s.Description = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetEffectiveTimeConfig(v *GetStandardResponseBodyStandardInfoEffectiveTimeConfig) *GetStandardResponseBodyStandardInfo {
	s.EffectiveTimeConfig = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetEnglishName(v string) *GetStandardResponseBodyStandardInfo {
	s.EnglishName = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetId(v int64) *GetStandardResponseBodyStandardInfo {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetLastModifier(v *GetStandardResponseBodyStandardInfoLastModifier) *GetStandardResponseBodyStandardInfo {
	s.LastModifier = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetLookupTableRelations(v []*GetStandardResponseBodyStandardInfoLookupTableRelations) *GetStandardResponseBodyStandardInfo {
	s.LookupTableRelations = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetModifyTime(v string) *GetStandardResponseBodyStandardInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetName(v string) *GetStandardResponseBodyStandardInfo {
	s.Name = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetOwner(v *GetStandardResponseBodyStandardInfoOwner) *GetStandardResponseBodyStandardInfo {
	s.Owner = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetStage(v string) *GetStandardResponseBodyStandardInfo {
	s.Stage = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetStandardGeneralMonitorConfig(v *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig) *GetStandardResponseBodyStandardInfo {
	s.StandardGeneralMonitorConfig = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetStandardRelations(v []*GetStandardResponseBodyStandardInfoStandardRelations) *GetStandardResponseBodyStandardInfo {
	s.StandardRelations = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetStandardSet(v *GetStandardResponseBodyStandardInfoStandardSet) *GetStandardResponseBodyStandardInfo {
	s.StandardSet = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetStandardTemplate(v *GetStandardResponseBodyStandardInfoStandardTemplate) *GetStandardResponseBodyStandardInfo {
	s.StandardTemplate = v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetStatus(v string) *GetStandardResponseBodyStandardInfo {
	s.Status = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetType(v string) *GetStandardResponseBodyStandardInfo {
	s.Type = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) SetVersion(v int32) *GetStandardResponseBodyStandardInfo {
	s.Version = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfo) Validate() error {
	if s.AttributeWithValueList != nil {
		for _, item := range s.AttributeWithValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.EffectiveTimeConfig != nil {
		if err := s.EffectiveTimeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LastModifier != nil {
		if err := s.LastModifier.Validate(); err != nil {
			return err
		}
	}
	if s.LookupTableRelations != nil {
		for _, item := range s.LookupTableRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	if s.StandardGeneralMonitorConfig != nil {
		if err := s.StandardGeneralMonitorConfig.Validate(); err != nil {
			return err
		}
	}
	if s.StandardRelations != nil {
		for _, item := range s.StandardRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StandardSet != nil {
		if err := s.StandardSet.Validate(); err != nil {
			return err
		}
	}
	if s.StandardTemplate != nil {
		if err := s.StandardTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardResponseBodyStandardInfoAttributeWithValueList struct {
	// The attribute details.
	Attribute *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Struct"`
	// The attribute value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueList) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueList) GetAttribute() *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	return s.Attribute
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueList) GetValue() *string {
	return s.Value
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueList) SetAttribute(v *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) *GetStandardResponseBodyStandardInfoAttributeWithValueList {
	s.Attribute = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueList) SetValue(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueList {
	s.Value = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueList) Validate() error {
	if s.Attribute != nil {
		if err := s.Attribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute struct {
	// The attribute code.
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
	// Indicates whether the monitoring configuration is enabled.
	EnableMonitorConfig *bool `json:"EnableMonitorConfig,omitempty" xml:"EnableMonitorConfig,omitempty"`
	// The attribute ID.
	//
	// example:
	//
	// 1011
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The monitoring configuration.
	MonitorConfig *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty" type:"Struct"`
	// The attribute name.
	//
	// example:
	//
	// attr1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The referenced attribute information.
	RefAttribute *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute `json:"RefAttribute,omitempty" xml:"RefAttribute,omitempty" type:"Struct"`
	// Indicates whether the attribute is required.
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// The attribute type. Valid values:
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
	// The value configuration.
	ValueConfig *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig `json:"ValueConfig,omitempty" xml:"ValueConfig,omitempty" type:"Struct"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetCode() *string {
	return s.Code
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetDescription() *string {
	return s.Description
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetEnableMonitorConfig() *bool {
	return s.EnableMonitorConfig
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetId() *int64 {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetMonitorConfig() *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig {
	return s.MonitorConfig
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetName() *string {
	return s.Name
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetRefAttribute() *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute {
	return s.RefAttribute
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetRequired() *bool {
	return s.Required
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetType() *string {
	return s.Type
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) GetValueConfig() *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig {
	return s.ValueConfig
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetCode(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.Code = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetDescription(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.Description = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetEnableMonitorConfig(v bool) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.EnableMonitorConfig = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetId(v int64) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetMonitorConfig(v *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.MonitorConfig = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetName(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.Name = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetRefAttribute(v *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.RefAttribute = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetRequired(v bool) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.Required = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetType(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.Type = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) SetValueConfig(v *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute {
	s.ValueConfig = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttribute) Validate() error {
	if s.MonitorConfig != nil {
		if err := s.MonitorConfig.Validate(); err != nil {
			return err
		}
	}
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

type GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig struct {
	// The field to check.
	//
	// example:
	//
	// column1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// Indicates whether the check is case-sensitive.
	IsCaseSensitive *bool `json:"IsCaseSensitive,omitempty" xml:"IsCaseSensitive,omitempty"`
	// The monitoring method. Valid values:
	//
	// - METADATA: metadata monitoring.
	//
	// - QUALITY: data quality monitoring.
	//
	// example:
	//
	// METADATA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) GetIsCaseSensitive() *bool {
	return s.IsCaseSensitive
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) GetType() *string {
	return s.Type
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) SetColumnName(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig {
	s.ColumnName = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) SetIsCaseSensitive(v bool) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig {
	s.IsCaseSensitive = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) SetType(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig {
	s.Type = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeMonitorConfig) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute struct {
	// The attribute source.
	AttributeFromInfo *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo `json:"AttributeFromInfo,omitempty" xml:"AttributeFromInfo,omitempty" type:"Struct"`
	// The attribute ID.
	//
	// example:
	//
	// 123
	AttributeId *int64 `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute) GetAttributeFromInfo() *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo {
	return s.AttributeFromInfo
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute) SetAttributeFromInfo(v *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute {
	s.AttributeFromInfo = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute) SetAttributeId(v int64) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute {
	s.AttributeId = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttribute) Validate() error {
	if s.AttributeFromInfo != nil {
		if err := s.AttributeFromInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo struct {
	// The attribute source. Valid values:
	//
	// - SYSTEM: system attribute.
	//
	// - CUSTOM: custom attribute.
	//
	// - STANDARD: standard.
	//
	// example:
	//
	// CUSTOM
	AttributeFrom *string `json:"AttributeFrom,omitempty" xml:"AttributeFrom,omitempty"`
	// The corresponding standard. This parameter takes effect when the attribute source is set to STANDARD.
	StandardReference *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference `json:"StandardReference,omitempty" xml:"StandardReference,omitempty" type:"Struct"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo) GetAttributeFrom() *string {
	return s.AttributeFrom
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo) GetStandardReference() *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference {
	return s.StandardReference
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo) SetAttributeFrom(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo {
	s.AttributeFrom = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo) SetStandardReference(v *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo {
	s.StandardReference = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfo) Validate() error {
	if s.StandardReference != nil {
		if err := s.StandardReference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference struct {
	// The standard ID.
	//
	// example:
	//
	// 22
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) GetStandardId() *int64 {
	return s.StandardId
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) GetVersion() *int32 {
	return s.Version
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) SetStandardId(v int64) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference {
	s.StandardId = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) SetVersion(v int32) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference {
	s.Version = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig struct {
	// The data type of the attribute value. Valid values:
	//
	// - STRING: string.
	//
	// - BIGINT: numeric type.
	//
	// - DOUBLE: floating-point type.
	//
	// - DATE: date, accurate to the day.
	//
	// - DATETIME: date, accurate to milliseconds.
	//
	// - BOOLEAN: Boolean.
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
	// The length of the attribute value. If empty or -1, the length is not limited. Typically, only string types have a length limit for attribute values.
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
	// example:
	//
	// CUSTOMIZED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value range.
	ValueRange *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange `json:"ValueRange,omitempty" xml:"ValueRange,omitempty" type:"Struct"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) GetDataType() *string {
	return s.DataType
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) GetLength() *int32 {
	return s.Length
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) GetType() *string {
	return s.Type
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) GetValueRange() *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange {
	return s.ValueRange
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) SetDataType(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig {
	s.DataType = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) SetDefaultValue(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig {
	s.DefaultValue = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) SetLength(v int32) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig {
	s.Length = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) SetType(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig {
	s.Type = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) SetValueRange(v *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig {
	s.ValueRange = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfig) Validate() error {
	if s.ValueRange != nil {
		if err := s.ValueRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange struct {
	// The value range. This parameter takes effect when the value source is set to DATAPHIN_ATTRIBUTE. Valid values:
	//
	// - BIZ_UNIT: business unit.
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
	LookupTableReference *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference `json:"LookupTableReference,omitempty" xml:"LookupTableReference,omitempty" type:"Struct"`
	// The value range. This parameter takes effect when the value source is set to MIN_MAX.
	MinMaxValueConfig *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig `json:"MinMaxValueConfig,omitempty" xml:"MinMaxValueConfig,omitempty" type:"Struct"`
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
	// example:
	//
	// NONE
	ValueConstraint *string `json:"ValueConstraint,omitempty" xml:"ValueConstraint,omitempty"`
	// The value range. This parameter takes effect when the value source is set to LIST.
	ValueList []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) GetDataphinAttributeType() *string {
	return s.DataphinAttributeType
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) GetLookupTableReference() *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference {
	return s.LookupTableReference
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) GetMinMaxValueConfig() *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	return s.MinMaxValueConfig
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) GetValueConstraint() *string {
	return s.ValueConstraint
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) GetValueList() []*string {
	return s.ValueList
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) SetDataphinAttributeType(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange {
	s.DataphinAttributeType = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) SetLookupTableReference(v *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange {
	s.LookupTableReference = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) SetMinMaxValueConfig(v *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange {
	s.MinMaxValueConfig = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) SetValueConstraint(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange {
	s.ValueConstraint = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) SetValueList(v []*string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange {
	s.ValueList = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRange) Validate() error {
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

type GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference struct {
	// The referenced lookup table field.
	//
	// example:
	//
	// col1
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The lookup table ID.
	//
	// example:
	//
	// 11
	LookupTableId *int64 `json:"LookupTableId,omitempty" xml:"LookupTableId,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) GetColumn() *string {
	return s.Column
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) GetLookupTableId() *int64 {
	return s.LookupTableId
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) SetColumn(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference {
	s.Column = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) SetLookupTableId(v int64) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference {
	s.LookupTableId = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig struct {
	// Indicates whether the maximum value is included.
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	// Indicates whether the minimum value is included.
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// The maximum value.
	//
	// example:
	//
	// 100
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum value.
	//
	// example:
	//
	// 0
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GetMaxValue() *string {
	return s.MaxValue
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GetMinValue() *string {
	return s.MinValue
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) SetIncludeMaxValue(v bool) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMaxValue = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) SetIncludeMinValue(v bool) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMinValue = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) SetMaxValue(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	s.MaxValue = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) SetMinValue(v string) *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	s.MinValue = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoCreator struct {
	// The user ID.
	//
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoCreator) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoCreator) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoCreator) GetId() *string {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfoCreator) GetName() *string {
	return s.Name
}

func (s *GetStandardResponseBodyStandardInfoCreator) SetId(v string) *GetStandardResponseBodyStandardInfoCreator {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoCreator) SetName(v string) *GetStandardResponseBodyStandardInfoCreator {
	s.Name = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoCreator) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoEffectiveTimeConfig struct {
	// The end time of the effective period.
	//
	// example:
	//
	// 2025-12-30 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the effective period.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The effective period type. Valid values:
	//
	// - FOREVER: permanent.
	//
	// - TIME_PERIOD: time period.
	//
	// example:
	//
	// TIME_PERIOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoEffectiveTimeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoEffectiveTimeConfig) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoEffectiveTimeConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStandardResponseBodyStandardInfoEffectiveTimeConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *GetStandardResponseBodyStandardInfoEffectiveTimeConfig) GetType() *string {
	return s.Type
}

func (s *GetStandardResponseBodyStandardInfoEffectiveTimeConfig) SetEndTime(v string) *GetStandardResponseBodyStandardInfoEffectiveTimeConfig {
	s.EndTime = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoEffectiveTimeConfig) SetStartTime(v string) *GetStandardResponseBodyStandardInfoEffectiveTimeConfig {
	s.StartTime = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoEffectiveTimeConfig) SetType(v string) *GetStandardResponseBodyStandardInfoEffectiveTimeConfig {
	s.Type = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoEffectiveTimeConfig) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoLastModifier struct {
	// The user ID.
	//
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoLastModifier) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoLastModifier) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoLastModifier) GetId() *string {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfoLastModifier) GetName() *string {
	return s.Name
}

func (s *GetStandardResponseBodyStandardInfoLastModifier) SetId(v string) *GetStandardResponseBodyStandardInfoLastModifier {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoLastModifier) SetName(v string) *GetStandardResponseBodyStandardInfoLastModifier {
	s.Name = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoLastModifier) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoLookupTableRelations struct {
	// The attribute name.
	//
	// example:
	//
	// zz
	AttributeCode *string `json:"AttributeCode,omitempty" xml:"AttributeCode,omitempty"`
	// The attribute ID.
	//
	// example:
	//
	// 1122
	AttributeId *int64 `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	// The attribute name.
	//
	// example:
	//
	// test
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// The lookup table code.
	//
	// example:
	//
	// test
	LookupTableCode *string `json:"LookupTableCode,omitempty" xml:"LookupTableCode,omitempty"`
	// The lookup table ID.
	//
	// example:
	//
	// 1121
	LookupTableId *int64 `json:"LookupTableId,omitempty" xml:"LookupTableId,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoLookupTableRelations) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoLookupTableRelations) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) GetAttributeCode() *string {
	return s.AttributeCode
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) GetAttributeName() *string {
	return s.AttributeName
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) GetLookupTableCode() *string {
	return s.LookupTableCode
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) GetLookupTableId() *int64 {
	return s.LookupTableId
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) SetAttributeCode(v string) *GetStandardResponseBodyStandardInfoLookupTableRelations {
	s.AttributeCode = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) SetAttributeId(v int64) *GetStandardResponseBodyStandardInfoLookupTableRelations {
	s.AttributeId = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) SetAttributeName(v string) *GetStandardResponseBodyStandardInfoLookupTableRelations {
	s.AttributeName = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) SetLookupTableCode(v string) *GetStandardResponseBodyStandardInfoLookupTableRelations {
	s.LookupTableCode = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) SetLookupTableId(v int64) *GetStandardResponseBodyStandardInfoLookupTableRelations {
	s.LookupTableId = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoLookupTableRelations) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoOwner struct {
	// The user ID.
	//
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username.
	//
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoOwner) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoOwner) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoOwner) GetId() *string {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfoOwner) GetName() *string {
	return s.Name
}

func (s *GetStandardResponseBodyStandardInfoOwner) SetId(v string) *GetStandardResponseBodyStandardInfoOwner {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoOwner) SetName(v string) *GetStandardResponseBodyStandardInfoOwner {
	s.Name = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoOwner) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig struct {
	// The list of standard monitoring configurations.
	StandardMonitorConfigList []*GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList `json:"StandardMonitorConfigList,omitempty" xml:"StandardMonitorConfigList,omitempty" type:"Repeated"`
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig) GetStandardMonitorConfigList() []*GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	return s.StandardMonitorConfigList
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig) SetStandardMonitorConfigList(v []*GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig {
	s.StandardMonitorConfigList = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfig) Validate() error {
	if s.StandardMonitorConfigList != nil {
		for _, item := range s.StandardMonitorConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList struct {
	// The associated attribute ID.
	//
	// example:
	//
	// 112
	AttributeId *int64 `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	// The monitoring configuration of the associated attribute.
	AttributeMonitorConfig *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig `json:"AttributeMonitorConfig,omitempty" xml:"AttributeMonitorConfig,omitempty" type:"Struct"`
	// The attribute name.
	//
	// example:
	//
	// teset
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// The rule description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The monitoring configuration ID. If empty, a new monitoring configuration is created. If an existing monitoring configuration ID is specified, the corresponding monitoring configuration is updated.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The method by which the monitoring configuration is added. Valid values:
	//
	// - BY_USER: manually added.
	//
	// - BY_SYSTEM_ATTRIBUTE: preset by system attribute.
	//
	// example:
	//
	// BY_SYSTEM_ATTRIBUTE
	MonitorFrom *string `json:"MonitorFrom,omitempty" xml:"MonitorFrom,omitempty"`
	// The rule template. This parameter is required when the monitoring type is QUALITY.
	QualityRuleTemplate *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate `json:"QualityRuleTemplate,omitempty" xml:"QualityRuleTemplate,omitempty" type:"Struct"`
	// The rule configurations. This parameter is required when the monitoring type is QUALITY.
	RuleConfigList []*GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList `json:"RuleConfigList,omitempty" xml:"RuleConfigList,omitempty" type:"Repeated"`
	// The rule name.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule subtype. Valid values:
	//
	// - BY_ATTRIBUTE: configured based on attributes.
	//
	// - CUSTOMIZED: custom configuration.
	//
	// This parameter is required when the monitoring type is QUALITY.
	//
	// example:
	//
	// CUSTOMIZED
	RuleSubType *string `json:"RuleSubType,omitempty" xml:"RuleSubType,omitempty"`
	// The rule validation configurations. This parameter is required when the monitoring type is QUALITY.
	RuleValidateConfigList []*GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList `json:"RuleValidateConfigList,omitempty" xml:"RuleValidateConfigList,omitempty" type:"Repeated"`
	// The monitoring type. Valid values:
	//
	// - METADATA: metadata monitoring.
	//
	// - QUALITY: data quality monitoring.
	//
	// example:
	//
	// METADATA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetAttributeMonitorConfig() *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	return s.AttributeMonitorConfig
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetAttributeName() *string {
	return s.AttributeName
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetDescription() *string {
	return s.Description
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetId() *int64 {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetMonitorFrom() *string {
	return s.MonitorFrom
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetQualityRuleTemplate() *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	return s.QualityRuleTemplate
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleConfigList() []*GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList {
	return s.RuleConfigList
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleName() *string {
	return s.RuleName
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleSubType() *string {
	return s.RuleSubType
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetRuleValidateConfigList() []*GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	return s.RuleValidateConfigList
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) GetType() *string {
	return s.Type
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetAttributeId(v int64) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.AttributeId = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetAttributeMonitorConfig(v *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.AttributeMonitorConfig = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetAttributeName(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.AttributeName = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetDescription(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.Description = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetId(v int64) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetMonitorFrom(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.MonitorFrom = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetQualityRuleTemplate(v *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.QualityRuleTemplate = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleConfigList(v []*GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleConfigList = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleName(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleName = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleSubType(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleSubType = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetRuleValidateConfigList(v []*GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.RuleValidateConfigList = v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) SetType(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList {
	s.Type = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigList) Validate() error {
	if s.AttributeMonitorConfig != nil {
		if err := s.AttributeMonitorConfig.Validate(); err != nil {
			return err
		}
	}
	if s.QualityRuleTemplate != nil {
		if err := s.QualityRuleTemplate.Validate(); err != nil {
			return err
		}
	}
	if s.RuleConfigList != nil {
		for _, item := range s.RuleConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleValidateConfigList != nil {
		for _, item := range s.RuleValidateConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig struct {
	// The field to check.
	//
	// example:
	//
	// column1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// Indicates whether the check is case-sensitive.
	IsCaseSensitive *bool `json:"IsCaseSensitive,omitempty" xml:"IsCaseSensitive,omitempty"`
	// The monitoring method. Valid values:
	//
	// - METADATA: metadata monitoring.
	//
	// - QUALITY: data quality monitoring.
	//
	// example:
	//
	// METADATA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GetIsCaseSensitive() *bool {
	return s.IsCaseSensitive
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) GetType() *string {
	return s.Type
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) SetColumnName(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	s.ColumnName = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) SetIsCaseSensitive(v bool) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	s.IsCaseSensitive = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) SetType(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig {
	s.Type = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListAttributeMonitorConfig) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate struct {
	// The template ID.
	//
	// example:
	//
	// 22
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The template name.
	//
	// example:
	//
	// CUSTOMIZED
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The template source. Valid values:
	//
	// - FROM_SYSTEM: system template.
	//
	// - CUSTOMIZED: custom template.
	//
	// example:
	//
	// CUSTOMIZED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GetId() *int64 {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GetName() *string {
	return s.Name
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) GetType() *string {
	return s.Type
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) SetId(v int64) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) SetName(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	s.Name = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) SetType(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate {
	s.Type = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListQualityRuleTemplate) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList struct {
	// The configuration item.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The configuration item value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) GetKey() *string {
	return s.Key
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) GetValue() *string {
	return s.Value
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) SetKey(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList {
	s.Key = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) SetValue(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList {
	s.Value = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleConfigList) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList struct {
	// The validation configuration ID. This ID is randomly generated by the business and must be unique.
	//
	// example:
	//
	// abc
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The metric. This parameter is required when the validation type is EXPRESSION.
	//
	// example:
	//
	// a
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The metric name. This parameter is required when the validation type is EXPRESSION.
	//
	// example:
	//
	// test
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The operator. If the validation type is EXPRESSION, valid values:
	//
	// - EQUAL: equal to.
	//
	// - NOT_EQUAL: not equal to.
	//
	// - LARGER: greater than.
	//
	// - LARGE_OR_EQUAL: greater than or equal to.
	//
	// - SMALLER: less than.
	//
	// - SMALLER_OR_EQUAL: less than or equal to.
	//
	// If the validation type is RELATION, valid values:
	//
	// - AND: and.
	//
	// - OR: or.
	//
	// example:
	//
	// AND
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The parent validation configuration ID. The parent rule validation type can only be RELATION.
	//
	// example:
	//
	// a
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The rule validation type. Valid values:
	//
	// - RELATION: relation.
	//
	// - EXPRESSION: expression.
	//
	// example:
	//
	// RELATION
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The comparison value.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetId() *string {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetMetric() *string {
	return s.Metric
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetMetricName() *string {
	return s.MetricName
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetOperator() *string {
	return s.Operator
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetParentId() *string {
	return s.ParentId
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetType() *string {
	return s.Type
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) GetValue() *string {
	return s.Value
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetId(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetMetric(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Metric = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetMetricName(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.MetricName = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetOperator(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Operator = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetParentId(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.ParentId = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetType(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Type = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) SetValue(v string) *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList {
	s.Value = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardGeneralMonitorConfigStandardMonitorConfigListRuleValidateConfigList) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoStandardRelations struct {
	// The association type. Valid values:
	//
	// - RELATIVE.
	//
	// example:
	//
	// RELATIVE
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The standard ID.
	//
	// example:
	//
	// 1121
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// The stage of the standard.
	//
	// example:
	//
	// DEV
	StandardStage *string `json:"StandardStage,omitempty" xml:"StandardStage,omitempty"`
	// The standard status.
	//
	// example:
	//
	// draft
	StandardStatus *string `json:"StandardStatus,omitempty" xml:"StandardStatus,omitempty"`
	// The standard version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoStandardRelations) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoStandardRelations) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) GetRelationType() *string {
	return s.RelationType
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) GetStandardId() *int64 {
	return s.StandardId
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) GetStandardStage() *string {
	return s.StandardStage
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) GetStandardStatus() *string {
	return s.StandardStatus
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) GetVersion() *int32 {
	return s.Version
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) SetRelationType(v string) *GetStandardResponseBodyStandardInfoStandardRelations {
	s.RelationType = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) SetStandardId(v int64) *GetStandardResponseBodyStandardInfoStandardRelations {
	s.StandardId = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) SetStandardStage(v string) *GetStandardResponseBodyStandardInfoStandardRelations {
	s.StandardStage = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) SetStandardStatus(v string) *GetStandardResponseBodyStandardInfoStandardRelations {
	s.StandardStatus = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) SetVersion(v int32) *GetStandardResponseBodyStandardInfoStandardRelations {
	s.Version = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardRelations) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoStandardSet struct {
	// The standard set code.
	//
	// example:
	//
	// cc
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The folder to which the standard set belongs.
	//
	// example:
	//
	// /dir1
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The standard set ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The standard set name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoStandardSet) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoStandardSet) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoStandardSet) GetCode() *string {
	return s.Code
}

func (s *GetStandardResponseBodyStandardInfoStandardSet) GetDirectory() *string {
	return s.Directory
}

func (s *GetStandardResponseBodyStandardInfoStandardSet) GetId() *int64 {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfoStandardSet) GetName() *string {
	return s.Name
}

func (s *GetStandardResponseBodyStandardInfoStandardSet) SetCode(v string) *GetStandardResponseBodyStandardInfoStandardSet {
	s.Code = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardSet) SetDirectory(v string) *GetStandardResponseBodyStandardInfoStandardSet {
	s.Directory = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardSet) SetId(v int64) *GetStandardResponseBodyStandardInfoStandardSet {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardSet) SetName(v string) *GetStandardResponseBodyStandardInfoStandardSet {
	s.Name = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardSet) Validate() error {
	return dara.Validate(s)
}

type GetStandardResponseBodyStandardInfoStandardTemplate struct {
	// The standard template code.
	//
	// example:
	//
	// 1121
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The standard template ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The standard template name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source of the standard template. Valid values:
	//
	// - CUSTOM: Custom standard template.
	//
	// - SYSTEM: System built-in standard template.
	//
	// example:
	//
	// SYSTEM
	TemplateFrom *string `json:"TemplateFrom,omitempty" xml:"TemplateFrom,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetStandardResponseBodyStandardInfoStandardTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetStandardResponseBodyStandardInfoStandardTemplate) GoString() string {
	return s.String()
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) GetCode() *string {
	return s.Code
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) GetId() *int64 {
	return s.Id
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) GetName() *string {
	return s.Name
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) GetTemplateFrom() *string {
	return s.TemplateFrom
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) GetVersion() *int32 {
	return s.Version
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) SetCode(v string) *GetStandardResponseBodyStandardInfoStandardTemplate {
	s.Code = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) SetId(v int64) *GetStandardResponseBodyStandardInfoStandardTemplate {
	s.Id = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) SetName(v string) *GetStandardResponseBodyStandardInfoStandardTemplate {
	s.Name = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) SetTemplateFrom(v string) *GetStandardResponseBodyStandardInfoStandardTemplate {
	s.TemplateFrom = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) SetVersion(v int32) *GetStandardResponseBodyStandardInfoStandardTemplate {
	s.Version = &v
	return s
}

func (s *GetStandardResponseBodyStandardInfoStandardTemplate) Validate() error {
	return dara.Validate(s)
}

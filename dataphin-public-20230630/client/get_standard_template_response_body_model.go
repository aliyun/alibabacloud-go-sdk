// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStandardTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStandardTemplateResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetStandardTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStandardTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStandardTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStandardTemplateResponseBody
	GetSuccess() *bool
	SetTemplateInfo(v *GetStandardTemplateResponseBodyTemplateInfo) *GetStandardTemplateResponseBody
	GetTemplateInfo() *GetStandardTemplateResponseBodyTemplateInfo
}

type GetStandardTemplateResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	TemplateInfo *GetStandardTemplateResponseBodyTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s GetStandardTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStandardTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStandardTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStandardTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStandardTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStandardTemplateResponseBody) GetTemplateInfo() *GetStandardTemplateResponseBodyTemplateInfo {
	return s.TemplateInfo
}

func (s *GetStandardTemplateResponseBody) SetCode(v string) *GetStandardTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetStandardTemplateResponseBody) SetHttpStatusCode(v int32) *GetStandardTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStandardTemplateResponseBody) SetMessage(v string) *GetStandardTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetStandardTemplateResponseBody) SetRequestId(v string) *GetStandardTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStandardTemplateResponseBody) SetSuccess(v bool) *GetStandardTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetStandardTemplateResponseBody) SetTemplateInfo(v *GetStandardTemplateResponseBodyTemplateInfo) *GetStandardTemplateResponseBody {
	s.TemplateInfo = v
	return s
}

func (s *GetStandardTemplateResponseBody) Validate() error {
	if s.TemplateInfo != nil {
		if err := s.TemplateInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardTemplateResponseBodyTemplateInfo struct {
	AttributesConfig *GetStandardTemplateResponseBodyTemplateInfoAttributesConfig `json:"AttributesConfig,omitempty" xml:"AttributesConfig,omitempty" type:"Struct"`
	// example:
	//
	// test01
	Code           *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	CodeRuleConfig *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig `json:"CodeRuleConfig,omitempty" xml:"CodeRuleConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator    *GetStandardTemplateResponseBodyTemplateInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 22
	Id             *int64                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	LastModifier   *GetStandardTemplateResponseBodyTemplateInfoLastModifier     `json:"LastModifier,omitempty" xml:"LastModifier,omitempty" type:"Struct"`
	MaintainerList []*GetStandardTemplateResponseBodyTemplateInfoMaintainerList `json:"MaintainerList,omitempty" xml:"MaintainerList,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 测试模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SYSTEM
	TemplateFrom *string `json:"TemplateFrom,omitempty" xml:"TemplateFrom,omitempty"`
	// uniqueId
	//
	// example:
	//
	// 1101
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfo) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetAttributesConfig() *GetStandardTemplateResponseBodyTemplateInfoAttributesConfig {
	return s.AttributesConfig
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetCode() *string {
	return s.Code
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetCodeRuleConfig() *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig {
	return s.CodeRuleConfig
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetCreator() *GetStandardTemplateResponseBodyTemplateInfoCreator {
	return s.Creator
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetDescription() *string {
	return s.Description
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetId() *int64 {
	return s.Id
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetLastModifier() *GetStandardTemplateResponseBodyTemplateInfoLastModifier {
	return s.LastModifier
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetMaintainerList() []*GetStandardTemplateResponseBodyTemplateInfoMaintainerList {
	return s.MaintainerList
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetName() *string {
	return s.Name
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetTemplateFrom() *string {
	return s.TemplateFrom
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetUniqueId() *string {
	return s.UniqueId
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) GetVersion() *int32 {
	return s.Version
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetAttributesConfig(v *GetStandardTemplateResponseBodyTemplateInfoAttributesConfig) *GetStandardTemplateResponseBodyTemplateInfo {
	s.AttributesConfig = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetCode(v string) *GetStandardTemplateResponseBodyTemplateInfo {
	s.Code = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetCodeRuleConfig(v *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig) *GetStandardTemplateResponseBodyTemplateInfo {
	s.CodeRuleConfig = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetCreateTime(v string) *GetStandardTemplateResponseBodyTemplateInfo {
	s.CreateTime = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetCreator(v *GetStandardTemplateResponseBodyTemplateInfoCreator) *GetStandardTemplateResponseBodyTemplateInfo {
	s.Creator = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetDescription(v string) *GetStandardTemplateResponseBodyTemplateInfo {
	s.Description = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetId(v int64) *GetStandardTemplateResponseBodyTemplateInfo {
	s.Id = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetLastModifier(v *GetStandardTemplateResponseBodyTemplateInfoLastModifier) *GetStandardTemplateResponseBodyTemplateInfo {
	s.LastModifier = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetMaintainerList(v []*GetStandardTemplateResponseBodyTemplateInfoMaintainerList) *GetStandardTemplateResponseBodyTemplateInfo {
	s.MaintainerList = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetModifyTime(v string) *GetStandardTemplateResponseBodyTemplateInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetName(v string) *GetStandardTemplateResponseBodyTemplateInfo {
	s.Name = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetTemplateFrom(v string) *GetStandardTemplateResponseBodyTemplateInfo {
	s.TemplateFrom = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetUniqueId(v string) *GetStandardTemplateResponseBodyTemplateInfo {
	s.UniqueId = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) SetVersion(v int32) *GetStandardTemplateResponseBodyTemplateInfo {
	s.Version = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfo) Validate() error {
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
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.LastModifier != nil {
		if err := s.LastModifier.Validate(); err != nil {
			return err
		}
	}
	if s.MaintainerList != nil {
		for _, item := range s.MaintainerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfig struct {
	AttributeList []*GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList `json:"AttributeList,omitempty" xml:"AttributeList,omitempty" type:"Repeated"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfig) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfig) GetAttributeList() []*GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	return s.AttributeList
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfig) SetAttributeList(v []*GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfig {
	s.AttributeList = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfig) Validate() error {
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

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList struct {
	// example:
	//
	// test_attr
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableMonitorConfig *bool   `json:"EnableMonitorConfig,omitempty" xml:"EnableMonitorConfig,omitempty"`
	// example:
	//
	// 1011
	Id            *int64                                                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	MonitorConfig *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty" type:"Struct"`
	// example:
	//
	// attr1
	Name         *string                                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	RefAttribute *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute `json:"RefAttribute,omitempty" xml:"RefAttribute,omitempty" type:"Struct"`
	Required     *bool                                                                                 `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// BIZ_ATTRIBUTE
	Type        *string                                                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	ValueConfig *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig `json:"ValueConfig,omitempty" xml:"ValueConfig,omitempty" type:"Struct"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetCode() *string {
	return s.Code
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetDescription() *string {
	return s.Description
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetEnableMonitorConfig() *bool {
	return s.EnableMonitorConfig
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetId() *int64 {
	return s.Id
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetMonitorConfig() *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig {
	return s.MonitorConfig
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetName() *string {
	return s.Name
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetRefAttribute() *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute {
	return s.RefAttribute
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetRequired() *bool {
	return s.Required
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetType() *string {
	return s.Type
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) GetValueConfig() *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig {
	return s.ValueConfig
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetCode(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.Code = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetDescription(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.Description = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetEnableMonitorConfig(v bool) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.EnableMonitorConfig = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetId(v int64) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.Id = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetMonitorConfig(v *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.MonitorConfig = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetName(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.Name = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetRefAttribute(v *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.RefAttribute = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetRequired(v bool) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.Required = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetType(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.Type = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) SetValueConfig(v *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList {
	s.ValueConfig = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeList) Validate() error {
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

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig struct {
	// example:
	//
	// column1
	ColumnName      *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	IsCaseSensitive *bool   `json:"IsCaseSensitive,omitempty" xml:"IsCaseSensitive,omitempty"`
	// example:
	//
	// METADATA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) GetIsCaseSensitive() *bool {
	return s.IsCaseSensitive
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) GetType() *string {
	return s.Type
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) SetColumnName(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig {
	s.ColumnName = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) SetIsCaseSensitive(v bool) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig {
	s.IsCaseSensitive = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) SetType(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig {
	s.Type = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListMonitorConfig) Validate() error {
	return dara.Validate(s)
}

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute struct {
	AttributeFromInfo *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo `json:"AttributeFromInfo,omitempty" xml:"AttributeFromInfo,omitempty" type:"Struct"`
	// example:
	//
	// 123
	AttributeId *int64 `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute) GetAttributeFromInfo() *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo {
	return s.AttributeFromInfo
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute) SetAttributeFromInfo(v *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute {
	s.AttributeFromInfo = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute) SetAttributeId(v int64) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute {
	s.AttributeId = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttribute) Validate() error {
	if s.AttributeFromInfo != nil {
		if err := s.AttributeFromInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo struct {
	// example:
	//
	// CUSTOM
	AttributeFrom     *string                                                                                                                 `json:"AttributeFrom,omitempty" xml:"AttributeFrom,omitempty"`
	StandardReference *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference `json:"StandardReference,omitempty" xml:"StandardReference,omitempty" type:"Struct"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo) GetAttributeFrom() *string {
	return s.AttributeFrom
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo) GetStandardReference() *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference {
	return s.StandardReference
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo) SetAttributeFrom(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo {
	s.AttributeFrom = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo) SetStandardReference(v *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo {
	s.StandardReference = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfo) Validate() error {
	if s.StandardReference != nil {
		if err := s.StandardReference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference struct {
	// example:
	//
	// 22
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) GetStandardId() *int64 {
	return s.StandardId
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) GetVersion() *int32 {
	return s.Version
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) SetStandardId(v int64) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference {
	s.StandardId = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) SetVersion(v int32) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference {
	s.Version = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListRefAttributeAttributeFromInfoStandardReference) Validate() error {
	return dara.Validate(s)
}

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig struct {
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
	// example:
	//
	// CUSTOMIZED
	Type       *string                                                                                        `json:"Type,omitempty" xml:"Type,omitempty"`
	ValueRange *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange `json:"ValueRange,omitempty" xml:"ValueRange,omitempty" type:"Struct"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) GetDataType() *string {
	return s.DataType
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) GetLength() *int32 {
	return s.Length
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) GetType() *string {
	return s.Type
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) GetValueRange() *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange {
	return s.ValueRange
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) SetDataType(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig {
	s.DataType = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) SetDefaultValue(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig {
	s.DefaultValue = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) SetLength(v int32) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig {
	s.Length = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) SetType(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig {
	s.Type = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) SetValueRange(v *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig {
	s.ValueRange = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfig) Validate() error {
	if s.ValueRange != nil {
		if err := s.ValueRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange struct {
	// example:
	//
	// DATAPHIN_ATTRIBUTE
	DataphinAttributeType *string                                                                                                            `json:"DataphinAttributeType,omitempty" xml:"DataphinAttributeType,omitempty"`
	LookupTableReference  *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference `json:"LookupTableReference,omitempty" xml:"LookupTableReference,omitempty" type:"Struct"`
	MinMaxValueConfig     *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig    `json:"MinMaxValueConfig,omitempty" xml:"MinMaxValueConfig,omitempty" type:"Struct"`
	// example:
	//
	// NONE
	ValueConstraint *string   `json:"ValueConstraint,omitempty" xml:"ValueConstraint,omitempty"`
	ValueList       []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) GetDataphinAttributeType() *string {
	return s.DataphinAttributeType
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) GetLookupTableReference() *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference {
	return s.LookupTableReference
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) GetMinMaxValueConfig() *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	return s.MinMaxValueConfig
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) GetValueConstraint() *string {
	return s.ValueConstraint
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) GetValueList() []*string {
	return s.ValueList
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) SetDataphinAttributeType(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange {
	s.DataphinAttributeType = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) SetLookupTableReference(v *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange {
	s.LookupTableReference = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) SetMinMaxValueConfig(v *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange {
	s.MinMaxValueConfig = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) SetValueConstraint(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange {
	s.ValueConstraint = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) SetValueList(v []*string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange {
	s.ValueList = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRange) Validate() error {
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

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference struct {
	// example:
	//
	// col1
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// example:
	//
	// 11
	LookupTableId *int64 `json:"LookupTableId,omitempty" xml:"LookupTableId,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) GetColumn() *string {
	return s.Column
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) GetLookupTableId() *int64 {
	return s.LookupTableId
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) SetColumn(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference {
	s.Column = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) SetLookupTableId(v int64) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference {
	s.LookupTableId = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeLookupTableReference) Validate() error {
	return dara.Validate(s)
}

type GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig struct {
	IncludeMaxValue *bool `json:"IncludeMaxValue,omitempty" xml:"IncludeMaxValue,omitempty"`
	IncludeMinValue *bool `json:"IncludeMinValue,omitempty" xml:"IncludeMinValue,omitempty"`
	// example:
	//
	// 100
	MaxValue *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// example:
	//
	// 0
	MinValue *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetMaxValue() *string {
	return s.MaxValue
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) GetMinValue() *string {
	return s.MinValue
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetIncludeMaxValue(v bool) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMaxValue = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetIncludeMinValue(v bool) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMinValue = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetMaxValue(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.MaxValue = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) SetMinValue(v string) *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig {
	s.MinValue = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoAttributesConfigAttributeListValueConfigValueRangeMinMaxValueConfig) Validate() error {
	return dara.Validate(s)
}

type GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig struct {
	AutoConfig *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig `json:"AutoConfig,omitempty" xml:"AutoConfig,omitempty" type:"Struct"`
	// example:
	//
	// CUSTOMIZED
	GenerateType *string `json:"GenerateType,omitempty" xml:"GenerateType,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig) GetAutoConfig() *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig {
	return s.AutoConfig
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig) GetGenerateType() *string {
	return s.GenerateType
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig) SetAutoConfig(v *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig {
	s.AutoConfig = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig) SetGenerateType(v string) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig {
	s.GenerateType = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfig) Validate() error {
	if s.AutoConfig != nil {
		if err := s.AutoConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig struct {
	CodeRuleList       []*GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList `json:"CodeRuleList,omitempty" xml:"CodeRuleList,omitempty" type:"Repeated"`
	NeedStrongValidate *bool                                                                              `json:"NeedStrongValidate,omitempty" xml:"NeedStrongValidate,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig) GetCodeRuleList() []*GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList {
	return s.CodeRuleList
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig) GetNeedStrongValidate() *bool {
	return s.NeedStrongValidate
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig) SetCodeRuleList(v []*GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig {
	s.CodeRuleList = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig) SetNeedStrongValidate(v bool) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig {
	s.NeedStrongValidate = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfig) Validate() error {
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

type GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList struct {
	AutoIncrementSequenceConfig *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig `json:"AutoIncrementSequenceConfig,omitempty" xml:"AutoIncrementSequenceConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// FIXED_STRING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) GetAutoIncrementSequenceConfig() *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	return s.AutoIncrementSequenceConfig
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) GetIndex() *int32 {
	return s.Index
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) GetType() *string {
	return s.Type
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) GetValue() *string {
	return s.Value
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) SetAutoIncrementSequenceConfig(v *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList {
	s.AutoIncrementSequenceConfig = v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) SetIndex(v int32) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList {
	s.Index = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) SetType(v string) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList {
	s.Type = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) SetValue(v string) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList {
	s.Value = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleList) Validate() error {
	if s.AutoIncrementSequenceConfig != nil {
		if err := s.AutoIncrementSequenceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig struct {
	// example:
	//
	// 3
	Digit           *int32 `json:"Digit,omitempty" xml:"Digit,omitempty"`
	NeedPaddingZero *bool  `json:"NeedPaddingZero,omitempty" xml:"NeedPaddingZero,omitempty"`
	// example:
	//
	// 1
	StartValue *int64 `json:"StartValue,omitempty" xml:"StartValue,omitempty"`
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetDigit() *int32 {
	return s.Digit
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetNeedPaddingZero() *bool {
	return s.NeedPaddingZero
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetStartValue() *int64 {
	return s.StartValue
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) GetStep() *int32 {
	return s.Step
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetDigit(v int32) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.Digit = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetNeedPaddingZero(v bool) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.NeedPaddingZero = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetStartValue(v int64) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.StartValue = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) SetStep(v int32) *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig {
	s.Step = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCodeRuleConfigAutoConfigCodeRuleListAutoIncrementSequenceConfig) Validate() error {
	return dara.Validate(s)
}

type GetStandardTemplateResponseBodyTemplateInfoCreator struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoCreator) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoCreator) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCreator) GetId() *string {
	return s.Id
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCreator) GetName() *string {
	return s.Name
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCreator) SetId(v string) *GetStandardTemplateResponseBodyTemplateInfoCreator {
	s.Id = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCreator) SetName(v string) *GetStandardTemplateResponseBodyTemplateInfoCreator {
	s.Name = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoCreator) Validate() error {
	return dara.Validate(s)
}

type GetStandardTemplateResponseBodyTemplateInfoLastModifier struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoLastModifier) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoLastModifier) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoLastModifier) GetId() *string {
	return s.Id
}

func (s *GetStandardTemplateResponseBodyTemplateInfoLastModifier) GetName() *string {
	return s.Name
}

func (s *GetStandardTemplateResponseBodyTemplateInfoLastModifier) SetId(v string) *GetStandardTemplateResponseBodyTemplateInfoLastModifier {
	s.Id = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoLastModifier) SetName(v string) *GetStandardTemplateResponseBodyTemplateInfoLastModifier {
	s.Name = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoLastModifier) Validate() error {
	return dara.Validate(s)
}

type GetStandardTemplateResponseBodyTemplateInfoMaintainerList struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetStandardTemplateResponseBodyTemplateInfoMaintainerList) String() string {
	return dara.Prettify(s)
}

func (s GetStandardTemplateResponseBodyTemplateInfoMaintainerList) GoString() string {
	return s.String()
}

func (s *GetStandardTemplateResponseBodyTemplateInfoMaintainerList) GetId() *string {
	return s.Id
}

func (s *GetStandardTemplateResponseBodyTemplateInfoMaintainerList) GetName() *string {
	return s.Name
}

func (s *GetStandardTemplateResponseBodyTemplateInfoMaintainerList) SetId(v string) *GetStandardTemplateResponseBodyTemplateInfoMaintainerList {
	s.Id = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoMaintainerList) SetName(v string) *GetStandardTemplateResponseBodyTemplateInfoMaintainerList {
	s.Name = &v
	return s
}

func (s *GetStandardTemplateResponseBodyTemplateInfoMaintainerList) Validate() error {
	return dara.Validate(s)
}

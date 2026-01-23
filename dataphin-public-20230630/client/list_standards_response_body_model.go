// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStandardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListStandardsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListStandardsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListStandardsResponseBody
	GetMessage() *string
	SetPageResult(v *ListStandardsResponseBodyPageResult) *ListStandardsResponseBody
	GetPageResult() *ListStandardsResponseBodyPageResult
	SetRequestId(v string) *ListStandardsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListStandardsResponseBody
	GetSuccess() *bool
}

type ListStandardsResponseBody struct {
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
	Message    *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListStandardsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListStandardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListStandardsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListStandardsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListStandardsResponseBody) GetPageResult() *ListStandardsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListStandardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStandardsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListStandardsResponseBody) SetCode(v string) *ListStandardsResponseBody {
	s.Code = &v
	return s
}

func (s *ListStandardsResponseBody) SetHttpStatusCode(v int32) *ListStandardsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListStandardsResponseBody) SetMessage(v string) *ListStandardsResponseBody {
	s.Message = &v
	return s
}

func (s *ListStandardsResponseBody) SetPageResult(v *ListStandardsResponseBodyPageResult) *ListStandardsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListStandardsResponseBody) SetRequestId(v string) *ListStandardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStandardsResponseBody) SetSuccess(v bool) *ListStandardsResponseBody {
	s.Success = &v
	return s
}

func (s *ListStandardsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStandardsResponseBodyPageResult struct {
	StandardList []*ListStandardsResponseBodyPageResultStandardList `json:"StandardList,omitempty" xml:"StandardList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStandardsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResult) GetStandardList() []*ListStandardsResponseBodyPageResultStandardList {
	return s.StandardList
}

func (s *ListStandardsResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListStandardsResponseBodyPageResult) SetStandardList(v []*ListStandardsResponseBodyPageResultStandardList) *ListStandardsResponseBodyPageResult {
	s.StandardList = v
	return s
}

func (s *ListStandardsResponseBodyPageResult) SetTotalCount(v int64) *ListStandardsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListStandardsResponseBodyPageResult) Validate() error {
	if s.StandardList != nil {
		for _, item := range s.StandardList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStandardsResponseBodyPageResultStandardList struct {
	AttributeWithValueList []*ListStandardsResponseBodyPageResultStandardListAttributeWithValueList `json:"AttributeWithValueList,omitempty" xml:"AttributeWithValueList,omitempty" type:"Repeated"`
	// example:
	//
	// zz
	Code    *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Creator *ListStandardsResponseBodyPageResultStandardListCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// test
	Description         *string                                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	EffectiveTimeConfig *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig `json:"EffectiveTimeConfig,omitempty" xml:"EffectiveTimeConfig,omitempty" type:"Struct"`
	// example:
	//
	// test
	EnglishName *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	// example:
	//
	// 1234
	Id           *int64                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	LastModifier *ListStandardsResponseBodyPageResultStandardListLastModifier `json:"LastModifier,omitempty" xml:"LastModifier,omitempty" type:"Struct"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test
	Name  *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner *ListStandardsResponseBodyPageResultStandardListOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// example:
	//
	// dev
	Stage            *string                                                          `json:"Stage,omitempty" xml:"Stage,omitempty"`
	StandardSet      *ListStandardsResponseBodyPageResultStandardListStandardSet      `json:"StandardSet,omitempty" xml:"StandardSet,omitempty" type:"Struct"`
	StandardTemplate *ListStandardsResponseBodyPageResultStandardListStandardTemplate `json:"StandardTemplate,omitempty" xml:"StandardTemplate,omitempty" type:"Struct"`
	// example:
	//
	// draft
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// zz
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardList) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardList) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetAttributeWithValueList() []*ListStandardsResponseBodyPageResultStandardListAttributeWithValueList {
	return s.AttributeWithValueList
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetCode() *string {
	return s.Code
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetCreator() *ListStandardsResponseBodyPageResultStandardListCreator {
	return s.Creator
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetDescription() *string {
	return s.Description
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetEffectiveTimeConfig() *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig {
	return s.EffectiveTimeConfig
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetEnglishName() *string {
	return s.EnglishName
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetId() *int64 {
	return s.Id
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetLastModifier() *ListStandardsResponseBodyPageResultStandardListLastModifier {
	return s.LastModifier
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetName() *string {
	return s.Name
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetOwner() *ListStandardsResponseBodyPageResultStandardListOwner {
	return s.Owner
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetStage() *string {
	return s.Stage
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetStandardSet() *ListStandardsResponseBodyPageResultStandardListStandardSet {
	return s.StandardSet
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetStandardTemplate() *ListStandardsResponseBodyPageResultStandardListStandardTemplate {
	return s.StandardTemplate
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetStatus() *string {
	return s.Status
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetType() *string {
	return s.Type
}

func (s *ListStandardsResponseBodyPageResultStandardList) GetVersion() *int32 {
	return s.Version
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetAttributeWithValueList(v []*ListStandardsResponseBodyPageResultStandardListAttributeWithValueList) *ListStandardsResponseBodyPageResultStandardList {
	s.AttributeWithValueList = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetCode(v string) *ListStandardsResponseBodyPageResultStandardList {
	s.Code = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetCreator(v *ListStandardsResponseBodyPageResultStandardListCreator) *ListStandardsResponseBodyPageResultStandardList {
	s.Creator = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetDescription(v string) *ListStandardsResponseBodyPageResultStandardList {
	s.Description = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetEffectiveTimeConfig(v *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) *ListStandardsResponseBodyPageResultStandardList {
	s.EffectiveTimeConfig = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetEnglishName(v string) *ListStandardsResponseBodyPageResultStandardList {
	s.EnglishName = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetId(v int64) *ListStandardsResponseBodyPageResultStandardList {
	s.Id = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetLastModifier(v *ListStandardsResponseBodyPageResultStandardListLastModifier) *ListStandardsResponseBodyPageResultStandardList {
	s.LastModifier = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetModifyTime(v string) *ListStandardsResponseBodyPageResultStandardList {
	s.ModifyTime = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetName(v string) *ListStandardsResponseBodyPageResultStandardList {
	s.Name = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetOwner(v *ListStandardsResponseBodyPageResultStandardListOwner) *ListStandardsResponseBodyPageResultStandardList {
	s.Owner = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetStage(v string) *ListStandardsResponseBodyPageResultStandardList {
	s.Stage = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetStandardSet(v *ListStandardsResponseBodyPageResultStandardListStandardSet) *ListStandardsResponseBodyPageResultStandardList {
	s.StandardSet = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetStandardTemplate(v *ListStandardsResponseBodyPageResultStandardListStandardTemplate) *ListStandardsResponseBodyPageResultStandardList {
	s.StandardTemplate = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetStatus(v string) *ListStandardsResponseBodyPageResultStandardList {
	s.Status = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetType(v string) *ListStandardsResponseBodyPageResultStandardList {
	s.Type = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) SetVersion(v int32) *ListStandardsResponseBodyPageResultStandardList {
	s.Version = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardList) Validate() error {
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
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
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

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueList struct {
	Attribute *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute `json:"Attribute,omitempty" xml:"Attribute,omitempty" type:"Struct"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueList) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueList) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueList) GetAttribute() *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	return s.Attribute
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueList) GetValue() *string {
	return s.Value
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueList) SetAttribute(v *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueList {
	s.Attribute = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueList) SetValue(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueList {
	s.Value = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueList) Validate() error {
	if s.Attribute != nil {
		if err := s.Attribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute struct {
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
	Id            *int64                                                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	MonitorConfig *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty" type:"Struct"`
	// example:
	//
	// attr1
	Name         *string                                                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RefAttribute *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute `json:"RefAttribute,omitempty" xml:"RefAttribute,omitempty" type:"Struct"`
	Required     *bool                                                                                       `json:"Required,omitempty" xml:"Required,omitempty"`
	// example:
	//
	// BIZ_ATTRIBUTE
	Type        *string                                                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
	ValueConfig *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig `json:"ValueConfig,omitempty" xml:"ValueConfig,omitempty" type:"Struct"`
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetCode() *string {
	return s.Code
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetDescription() *string {
	return s.Description
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetEnableMonitorConfig() *bool {
	return s.EnableMonitorConfig
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetId() *int64 {
	return s.Id
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetMonitorConfig() *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig {
	return s.MonitorConfig
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetName() *string {
	return s.Name
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetRefAttribute() *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute {
	return s.RefAttribute
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetRequired() *bool {
	return s.Required
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetType() *string {
	return s.Type
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) GetValueConfig() *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig {
	return s.ValueConfig
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetCode(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.Code = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetDescription(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.Description = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetEnableMonitorConfig(v bool) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.EnableMonitorConfig = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetId(v int64) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.Id = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetMonitorConfig(v *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.MonitorConfig = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetName(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.Name = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetRefAttribute(v *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.RefAttribute = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetRequired(v bool) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.Required = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetType(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.Type = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) SetValueConfig(v *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute {
	s.ValueConfig = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttribute) Validate() error {
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

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig struct {
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

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) GetIsCaseSensitive() *bool {
	return s.IsCaseSensitive
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) GetType() *string {
	return s.Type
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) SetColumnName(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig {
	s.ColumnName = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) SetIsCaseSensitive(v bool) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig {
	s.IsCaseSensitive = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) SetType(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig {
	s.Type = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeMonitorConfig) Validate() error {
	return dara.Validate(s)
}

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute struct {
	AttributeFromInfo *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo `json:"AttributeFromInfo,omitempty" xml:"AttributeFromInfo,omitempty" type:"Struct"`
	// example:
	//
	// 123
	AttributeId *int64 `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute) GetAttributeFromInfo() *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo {
	return s.AttributeFromInfo
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute) GetAttributeId() *int64 {
	return s.AttributeId
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute) SetAttributeFromInfo(v *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute {
	s.AttributeFromInfo = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute) SetAttributeId(v int64) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute {
	s.AttributeId = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttribute) Validate() error {
	if s.AttributeFromInfo != nil {
		if err := s.AttributeFromInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo struct {
	// example:
	//
	// CUSTOM
	AttributeFrom     *string                                                                                                                       `json:"AttributeFrom,omitempty" xml:"AttributeFrom,omitempty"`
	StandardReference *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference `json:"StandardReference,omitempty" xml:"StandardReference,omitempty" type:"Struct"`
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo) GetAttributeFrom() *string {
	return s.AttributeFrom
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo) GetStandardReference() *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference {
	return s.StandardReference
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo) SetAttributeFrom(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo {
	s.AttributeFrom = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo) SetStandardReference(v *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo {
	s.StandardReference = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfo) Validate() error {
	if s.StandardReference != nil {
		if err := s.StandardReference.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference struct {
	// example:
	//
	// 22
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) GetStandardId() *int64 {
	return s.StandardId
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) GetVersion() *int32 {
	return s.Version
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) SetStandardId(v int64) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference {
	s.StandardId = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) SetVersion(v int32) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference {
	s.Version = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeRefAttributeAttributeFromInfoStandardReference) Validate() error {
	return dara.Validate(s)
}

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig struct {
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
	Type       *string                                                                                              `json:"Type,omitempty" xml:"Type,omitempty"`
	ValueRange *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange `json:"ValueRange,omitempty" xml:"ValueRange,omitempty" type:"Struct"`
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) GetDataType() *string {
	return s.DataType
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) GetLength() *int32 {
	return s.Length
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) GetType() *string {
	return s.Type
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) GetValueRange() *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange {
	return s.ValueRange
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) SetDataType(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig {
	s.DataType = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) SetDefaultValue(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig {
	s.DefaultValue = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) SetLength(v int32) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig {
	s.Length = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) SetType(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig {
	s.Type = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) SetValueRange(v *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig {
	s.ValueRange = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfig) Validate() error {
	if s.ValueRange != nil {
		if err := s.ValueRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange struct {
	// example:
	//
	// DATAPHIN_ATTRIBUTE
	DataphinAttributeType *string                                                                                                                  `json:"DataphinAttributeType,omitempty" xml:"DataphinAttributeType,omitempty"`
	LookupTableReference  *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference `json:"LookupTableReference,omitempty" xml:"LookupTableReference,omitempty" type:"Struct"`
	MinMaxValueConfig     *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig    `json:"MinMaxValueConfig,omitempty" xml:"MinMaxValueConfig,omitempty" type:"Struct"`
	// example:
	//
	// NONE
	ValueConstraint *string   `json:"ValueConstraint,omitempty" xml:"ValueConstraint,omitempty"`
	ValueList       []*string `json:"ValueList,omitempty" xml:"ValueList,omitempty" type:"Repeated"`
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) GetDataphinAttributeType() *string {
	return s.DataphinAttributeType
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) GetLookupTableReference() *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference {
	return s.LookupTableReference
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) GetMinMaxValueConfig() *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	return s.MinMaxValueConfig
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) GetValueConstraint() *string {
	return s.ValueConstraint
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) GetValueList() []*string {
	return s.ValueList
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) SetDataphinAttributeType(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange {
	s.DataphinAttributeType = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) SetLookupTableReference(v *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange {
	s.LookupTableReference = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) SetMinMaxValueConfig(v *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange {
	s.MinMaxValueConfig = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) SetValueConstraint(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange {
	s.ValueConstraint = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) SetValueList(v []*string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange {
	s.ValueList = v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRange) Validate() error {
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

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference struct {
	// example:
	//
	// col1
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// example:
	//
	// 11
	LookupTableId *int64 `json:"LookupTableId,omitempty" xml:"LookupTableId,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) GetColumn() *string {
	return s.Column
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) GetLookupTableId() *int64 {
	return s.LookupTableId
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) SetColumn(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference {
	s.Column = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) SetLookupTableId(v int64) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference {
	s.LookupTableId = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeLookupTableReference) Validate() error {
	return dara.Validate(s)
}

type ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig struct {
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

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GetIncludeMaxValue() *bool {
	return s.IncludeMaxValue
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GetIncludeMinValue() *bool {
	return s.IncludeMinValue
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GetMaxValue() *string {
	return s.MaxValue
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) GetMinValue() *string {
	return s.MinValue
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) SetIncludeMaxValue(v bool) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMaxValue = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) SetIncludeMinValue(v bool) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	s.IncludeMinValue = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) SetMaxValue(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	s.MaxValue = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) SetMinValue(v string) *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig {
	s.MinValue = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListAttributeWithValueListAttributeValueConfigValueRangeMinMaxValueConfig) Validate() error {
	return dara.Validate(s)
}

type ListStandardsResponseBodyPageResultStandardListCreator struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListCreator) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListCreator) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListCreator) GetId() *string {
	return s.Id
}

func (s *ListStandardsResponseBodyPageResultStandardListCreator) GetName() *string {
	return s.Name
}

func (s *ListStandardsResponseBodyPageResultStandardListCreator) SetId(v string) *ListStandardsResponseBodyPageResultStandardListCreator {
	s.Id = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListCreator) SetName(v string) *ListStandardsResponseBodyPageResultStandardListCreator {
	s.Name = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListCreator) Validate() error {
	return dara.Validate(s)
}

type ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig struct {
	// example:
	//
	// 2025-12-30 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// TIME_PERIOD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) GetType() *string {
	return s.Type
}

func (s *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) SetEndTime(v string) *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig {
	s.EndTime = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) SetStartTime(v string) *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig {
	s.StartTime = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) SetType(v string) *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig {
	s.Type = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListEffectiveTimeConfig) Validate() error {
	return dara.Validate(s)
}

type ListStandardsResponseBodyPageResultStandardListLastModifier struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListLastModifier) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListLastModifier) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListLastModifier) GetId() *string {
	return s.Id
}

func (s *ListStandardsResponseBodyPageResultStandardListLastModifier) GetName() *string {
	return s.Name
}

func (s *ListStandardsResponseBodyPageResultStandardListLastModifier) SetId(v string) *ListStandardsResponseBodyPageResultStandardListLastModifier {
	s.Id = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListLastModifier) SetName(v string) *ListStandardsResponseBodyPageResultStandardListLastModifier {
	s.Name = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListLastModifier) Validate() error {
	return dara.Validate(s)
}

type ListStandardsResponseBodyPageResultStandardListOwner struct {
	// example:
	//
	// 300000913
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// susan
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListOwner) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListOwner) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListOwner) GetId() *string {
	return s.Id
}

func (s *ListStandardsResponseBodyPageResultStandardListOwner) GetName() *string {
	return s.Name
}

func (s *ListStandardsResponseBodyPageResultStandardListOwner) SetId(v string) *ListStandardsResponseBodyPageResultStandardListOwner {
	s.Id = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListOwner) SetName(v string) *ListStandardsResponseBodyPageResultStandardListOwner {
	s.Name = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListOwner) Validate() error {
	return dara.Validate(s)
}

type ListStandardsResponseBodyPageResultStandardListStandardSet struct {
	// example:
	//
	// cc
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// /dir1
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListStandardSet) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListStandardSet) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardSet) GetCode() *string {
	return s.Code
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardSet) GetDirectory() *string {
	return s.Directory
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardSet) GetId() *int64 {
	return s.Id
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardSet) GetName() *string {
	return s.Name
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardSet) SetCode(v string) *ListStandardsResponseBodyPageResultStandardListStandardSet {
	s.Code = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardSet) SetDirectory(v string) *ListStandardsResponseBodyPageResultStandardListStandardSet {
	s.Directory = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardSet) SetId(v int64) *ListStandardsResponseBodyPageResultStandardListStandardSet {
	s.Id = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardSet) SetName(v string) *ListStandardsResponseBodyPageResultStandardListStandardSet {
	s.Name = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardSet) Validate() error {
	return dara.Validate(s)
}

type ListStandardsResponseBodyPageResultStandardListStandardTemplate struct {
	// example:
	//
	// 1121
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// SYSTEM
	TemplateFrom *string `json:"TemplateFrom,omitempty" xml:"TemplateFrom,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListStandardsResponseBodyPageResultStandardListStandardTemplate) String() string {
	return dara.Prettify(s)
}

func (s ListStandardsResponseBodyPageResultStandardListStandardTemplate) GoString() string {
	return s.String()
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) GetCode() *string {
	return s.Code
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) GetId() *int64 {
	return s.Id
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) GetName() *string {
	return s.Name
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) GetTemplateFrom() *string {
	return s.TemplateFrom
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) GetVersion() *int32 {
	return s.Version
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) SetCode(v string) *ListStandardsResponseBodyPageResultStandardListStandardTemplate {
	s.Code = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) SetId(v int64) *ListStandardsResponseBodyPageResultStandardListStandardTemplate {
	s.Id = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) SetName(v string) *ListStandardsResponseBodyPageResultStandardListStandardTemplate {
	s.Name = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) SetTemplateFrom(v string) *ListStandardsResponseBodyPageResultStandardListStandardTemplate {
	s.TemplateFrom = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) SetVersion(v int32) *ListStandardsResponseBodyPageResultStandardListStandardTemplate {
	s.Version = &v
	return s
}

func (s *ListStandardsResponseBodyPageResultStandardListStandardTemplate) Validate() error {
	return dara.Validate(s)
}

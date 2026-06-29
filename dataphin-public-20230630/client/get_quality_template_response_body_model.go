// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityTemplateResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetQualityTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityTemplateResponseBody
	GetMessage() *string
	SetQualityTemplateInfo(v *GetQualityTemplateResponseBodyQualityTemplateInfo) *GetQualityTemplateResponseBody
	GetQualityTemplateInfo() *GetQualityTemplateResponseBodyQualityTemplateInfo
	SetRequestId(v string) *GetQualityTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityTemplateResponseBody
	GetSuccess() *bool
}

type GetQualityTemplateResponseBody struct {
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
	// The details of the template object.
	QualityTemplateInfo *GetQualityTemplateResponseBodyQualityTemplateInfo `json:"QualityTemplateInfo,omitempty" xml:"QualityTemplateInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityTemplateResponseBody) GetQualityTemplateInfo() *GetQualityTemplateResponseBodyQualityTemplateInfo {
	return s.QualityTemplateInfo
}

func (s *GetQualityTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityTemplateResponseBody) SetCode(v string) *GetQualityTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityTemplateResponseBody) SetHttpStatusCode(v int32) *GetQualityTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityTemplateResponseBody) SetMessage(v string) *GetQualityTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityTemplateResponseBody) SetQualityTemplateInfo(v *GetQualityTemplateResponseBodyQualityTemplateInfo) *GetQualityTemplateResponseBody {
	s.QualityTemplateInfo = v
	return s
}

func (s *GetQualityTemplateResponseBody) SetRequestId(v string) *GetQualityTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityTemplateResponseBody) SetSuccess(v bool) *GetQualityTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityTemplateResponseBody) Validate() error {
	if s.QualityTemplateInfo != nil {
		if err := s.QualityTemplateInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityTemplateResponseBodyQualityTemplateInfo struct {
	// The template category. Valid values:
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
	// example:
	//
	// CONSISTENT
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// The display name of the template category.
	//
	// example:
	//
	// 一致性
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The template description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key-value pairs of the rule configuration.
	FormPropertyList []*GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
	// The template ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the template is a system template.
	IsSystemTemplate *bool `json:"IsSystemTemplate,omitempty" xml:"IsSystemTemplate,omitempty"`
	// The user ID of the last modifier.
	//
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The name of the last modifier.
	//
	// example:
	//
	// test
	ModifierName *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The template name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the owner.
	//
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the owner.
	//
	// example:
	//
	// test
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// Indicates whether all data source types are supported.
	SupportAllDataSourceType *bool `json:"SupportAllDataSourceType,omitempty" xml:"SupportAllDataSourceType,omitempty"`
	// The list of supported data source types, such as MySQL, Oracle, Microsoft SQL Server, MaxCompute, and Hive.
	SupportDataSourceTypeList []*string `json:"SupportDataSourceTypeList,omitempty" xml:"SupportDataSourceTypeList,omitempty" type:"Repeated"`
	// The templatetype. Valid values:
	//
	// - FIELD_NULL_VALUE_VALIDATE: field null value check.
	//
	// - FIELD_EMPTY_STRING_VALIDATE: field empty character string check.
	//
	// - FIELD_UNIQUE_VALIDATE: field uniqueness check.
	//
	// - FIELD_GROUP_COUNT_VALIDATE: field unique value count check.
	//
	// - FIELD_DUPLICATE_VALUE_COUNT_VALIDATE: field duplicate value count check.
	//
	// - FUNCTION_TIME_COMPARE: time function comparison.
	//
	// - SINGLE_TABLE_TIME_COMPARE: non-partitioned table time field comparison.
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
	// - SINGLE_TABLE_FIELD_VALUE_COMPARE: non-partitioned table field value consistency comparison.
	//
	// - SINGLE_TABLE_FIELD_STATISTICAL_COMPARE: non-partitioned table field statistical value consistency comparison.
	//
	// - SINGLE_TABLE_FIELD_EXP_COMPARE: non-partitioned table field business logic consistency comparison.
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
	// - DATASOURCE_AVAILABLE_CHECK: data source connectivity monitoring.
	//
	// - TABLE_SCHEMA_CHECK: table schema change monitoring.
	//
	// - REAL_TIME_OFFLINE_COMPARE: real-time and offline comparison.
	//
	// - REAL_TIME_STATISTICAL_VALIDATE: real-time statistical value monitoring.
	//
	// - REAL_TIME_MULTI_CHAIN_COMPARE: real-time multi-link comparison.
	//
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The display name of the templatetype.
	//
	// example:
	//
	// 字段空值校验
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetQualityTemplateResponseBodyQualityTemplateInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityTemplateResponseBodyQualityTemplateInfo) GoString() string {
	return s.String()
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetCatalog() *string {
	return s.Catalog
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetCatalogName() *string {
	return s.CatalogName
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetDescription() *string {
	return s.Description
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetFormPropertyList() []*GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList {
	return s.FormPropertyList
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetId() *int64 {
	return s.Id
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetIsSystemTemplate() *bool {
	return s.IsSystemTemplate
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetModifierName() *string {
	return s.ModifierName
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetSupportAllDataSourceType() *bool {
	return s.SupportAllDataSourceType
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetSupportDataSourceTypeList() []*string {
	return s.SupportDataSourceTypeList
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) GetTypeName() *string {
	return s.TypeName
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetCatalog(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.Catalog = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetCatalogName(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.CatalogName = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetCreateTime(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.CreateTime = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetCreator(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.Creator = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetCreatorName(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.CreatorName = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetDescription(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.Description = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetFormPropertyList(v []*GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.FormPropertyList = v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetId(v int64) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.Id = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetIsSystemTemplate(v bool) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.IsSystemTemplate = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetModifier(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.Modifier = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetModifierName(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.ModifierName = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetModifyTime(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetName(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.Name = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetOwner(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.Owner = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetOwnerName(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.OwnerName = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetSupportAllDataSourceType(v bool) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.SupportAllDataSourceType = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetSupportDataSourceTypeList(v []*string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.SupportDataSourceTypeList = v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetType(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.Type = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) SetTypeName(v string) *GetQualityTemplateResponseBodyQualityTemplateInfo {
	s.TypeName = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfo) Validate() error {
	if s.FormPropertyList != nil {
		for _, item := range s.FormPropertyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList struct {
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

func (s GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) GoString() string {
	return s.String()
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) GetComponentType() *string {
	return s.ComponentType
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) GetName() *string {
	return s.Name
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) GetValue() *string {
	return s.Value
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) SetComponentType(v string) *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList {
	s.ComponentType = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) SetName(v string) *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList {
	s.Name = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) SetValue(v string) *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList {
	s.Value = &v
	return s
}

func (s *GetQualityTemplateResponseBodyQualityTemplateInfoFormPropertyList) Validate() error {
	return dara.Validate(s)
}

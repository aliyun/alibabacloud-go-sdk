// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListQualityTemplatesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListQualityTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListQualityTemplatesResponseBody
	GetMessage() *string
	SetPageResult(v *ListQualityTemplatesResponseBodyPageResult) *ListQualityTemplatesResponseBody
	GetPageResult() *ListQualityTemplatesResponseBodyPageResult
	SetRequestId(v string) *ListQualityTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListQualityTemplatesResponseBody
	GetSuccess() *bool
}

type ListQualityTemplatesResponseBody struct {
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
	// The paged query result.
	PageResult *ListQualityTemplatesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListQualityTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListQualityTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListQualityTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListQualityTemplatesResponseBody) GetPageResult() *ListQualityTemplatesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListQualityTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQualityTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListQualityTemplatesResponseBody) SetCode(v string) *ListQualityTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListQualityTemplatesResponseBody) SetHttpStatusCode(v int32) *ListQualityTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListQualityTemplatesResponseBody) SetMessage(v string) *ListQualityTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListQualityTemplatesResponseBody) SetPageResult(v *ListQualityTemplatesResponseBodyPageResult) *ListQualityTemplatesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListQualityTemplatesResponseBody) SetRequestId(v string) *ListQualityTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQualityTemplatesResponseBody) SetSuccess(v bool) *ListQualityTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListQualityTemplatesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityTemplatesResponseBodyPageResult struct {
	// The list of quality templates.
	QualityTemplateList []*ListQualityTemplatesResponseBodyPageResultQualityTemplateList `json:"QualityTemplateList,omitempty" xml:"QualityTemplateList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQualityTemplatesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesResponseBodyPageResult) GetQualityTemplateList() []*ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	return s.QualityTemplateList
}

func (s *ListQualityTemplatesResponseBodyPageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListQualityTemplatesResponseBodyPageResult) SetQualityTemplateList(v []*ListQualityTemplatesResponseBodyPageResultQualityTemplateList) *ListQualityTemplatesResponseBodyPageResult {
	s.QualityTemplateList = v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResult) SetTotalCount(v int64) *ListQualityTemplatesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResult) Validate() error {
	if s.QualityTemplateList != nil {
		for _, item := range s.QualityTemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListQualityTemplatesResponseBodyPageResultQualityTemplateList struct {
	// The template category. Valid values:
	//
	// - CONSISTENT: consistency
	//
	// - EFFECTIVE: validity
	//
	// - TIMELINESE: timeliness
	//
	// - ACCURATE: accuracy
	//
	// - UNIQUENESS: uniqueness
	//
	// - COMPLETENESS: completeness
	//
	// - STABILITY: stability
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
	// The rule configuration key-value pairs.
	FormPropertyList []*ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
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
	// The supported data source types, such as MySQL, Oracle, Microsoft SQL Server, MaxCompute, and Hive.
	SupportDataSourceTypeList []*string `json:"SupportDataSourceTypeList,omitempty" xml:"SupportDataSourceTypeList,omitempty" type:"Repeated"`
	// The templatetype. Valid values:
	//
	// - FIELD_NULL_VALUE_VALIDATE: field null value check
	//
	// - FIELD_EMPTY_STRING_VALIDATE: field empty character string check
	//
	// - FIELD_UNIQUE_VALIDATE: field uniqueness check
	//
	// - FIELD_GROUP_COUNT_VALIDATE: field unique value count check
	//
	// - FIELD_DUPLICATE_VALUE_COUNT_VALIDATE: field duplicate value count check
	//
	// - FUNCTION_TIME_COMPARE: time function comparison
	//
	// - SINGLE_TABLE_TIME_COMPARE: non-partitioned table time field comparison
	//
	// - DOUBLE_TABLE_TIME_COMPARE: two-table time field comparison
	//
	// - FIELD_FORMAT_VALIDATE: field format check
	//
	// - FIELD_LENGTH_VALIDATE: field length check
	//
	// - FIELD_VALUE_RANGE_VALIDATE: field value range check
	//
	// - CODE_TABLE_COMPARE: lookup table reference comparison
	//
	// - STANDARD_CODE_TABLE_COMPARE: data standard lookup table reference comparison
	//
	// - SINGLE_TABLE_FIELD_VALUE_COMPARE: non-partitioned table field value consistency comparison
	//
	// - SINGLE_TABLE_FIELD_STATISTICAL_COMPARE: non-partitioned table field statistical value consistency comparison
	//
	// - SINGLE_TABLE_FIELD_EXP_COMPARE: non-partitioned table field business logic consistency comparison
	//
	// - DOUBLE_TABLE_FIELD_VALUE_COMPARE: two-table field value consistency comparison
	//
	// - DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: two-table field statistical value consistency comparison
	//
	// - CROSS_DOUBLE_TABLE_FIELD_STATISTICAL_COMPARE: cross-source two-table field statistical value consistency comparison
	//
	// - DOUBLE_TABLE_FIELD_EXP_COMPARE: two-table field business logic consistency comparison
	//
	// - TABLE_STABILITY_VALIDATE: table stability check
	//
	// - TABLE_FLUCTUATION_VALIDATE: table fluctuation check
	//
	// - FIELD_STABILITY_VALIDATE: field stability check
	//
	// - FIELD_FLUCTUATION_VALIDATE: field fluctuation check
	//
	// - CUSTOM_STATISTICAL_VALIDATE: custom statistical metric check
	//
	// - CUSTOM_DATA_DETAILS_VALIDATE: custom data details check
	//
	// - DATASOURCE_AVAILABLE_CHECK: data source connectivity monitoring
	//
	// - TABLE_SCHEMA_CHECK: table schema change monitoring
	//
	// - REAL_TIME_OFFLINE_COMPARE: real-time and offline comparison
	//
	// - REAL_TIME_STATISTICAL_VALIDATE: real-time statistical value monitoring
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

func (s ListQualityTemplatesResponseBodyPageResultQualityTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetCatalog() *string {
	return s.Catalog
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetCreator() *string {
	return s.Creator
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetDescription() *string {
	return s.Description
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetFormPropertyList() []*ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList {
	return s.FormPropertyList
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetId() *int64 {
	return s.Id
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetIsSystemTemplate() *bool {
	return s.IsSystemTemplate
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetModifier() *string {
	return s.Modifier
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetModifierName() *string {
	return s.ModifierName
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetName() *string {
	return s.Name
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetOwner() *string {
	return s.Owner
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetSupportAllDataSourceType() *bool {
	return s.SupportAllDataSourceType
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetSupportDataSourceTypeList() []*string {
	return s.SupportDataSourceTypeList
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetType() *string {
	return s.Type
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) GetTypeName() *string {
	return s.TypeName
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetCatalog(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.Catalog = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetCatalogName(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.CatalogName = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetCreateTime(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.CreateTime = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetCreator(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.Creator = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetCreatorName(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.CreatorName = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetDescription(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.Description = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetFormPropertyList(v []*ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.FormPropertyList = v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetId(v int64) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.Id = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetIsSystemTemplate(v bool) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.IsSystemTemplate = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetModifier(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.Modifier = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetModifierName(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.ModifierName = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetModifyTime(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.ModifyTime = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetName(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.Name = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetOwner(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.Owner = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetOwnerName(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.OwnerName = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetSupportAllDataSourceType(v bool) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.SupportAllDataSourceType = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetSupportDataSourceTypeList(v []*string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.SupportDataSourceTypeList = v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetType(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.Type = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) SetTypeName(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateList {
	s.TypeName = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateList) Validate() error {
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

type ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList struct {
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

func (s ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) GetComponentType() *string {
	return s.ComponentType
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) GetName() *string {
	return s.Name
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) GetValue() *string {
	return s.Value
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) SetComponentType(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList {
	s.ComponentType = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) SetName(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList {
	s.Name = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) SetValue(v string) *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList {
	s.Value = &v
	return s
}

func (s *ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList) Validate() error {
	return dara.Validate(s)
}

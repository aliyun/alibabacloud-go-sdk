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
	Message    *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListQualityTemplatesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	QualityTemplateList []*ListQualityTemplatesResponseBodyPageResultQualityTemplateList `json:"QualityTemplateList,omitempty" xml:"QualityTemplateList,omitempty" type:"Repeated"`
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
	// example:
	//
	// CONSISTENT
	Catalog *string `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// example:
	//
	// 一致性
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// test
	Description      *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	FormPropertyList []*ListQualityTemplatesResponseBodyPageResultQualityTemplateListFormPropertyList `json:"FormPropertyList,omitempty" xml:"FormPropertyList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id               *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsSystemTemplate *bool  `json:"IsSystemTemplate,omitempty" xml:"IsSystemTemplate,omitempty"`
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// test
	ModifierName *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// test
	OwnerName                 *string   `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	SupportAllDataSourceType  *bool     `json:"SupportAllDataSourceType,omitempty" xml:"SupportAllDataSourceType,omitempty"`
	SupportDataSourceTypeList []*string `json:"SupportDataSourceTypeList,omitempty" xml:"SupportDataSourceTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// FIELD_NULL_VALUE_VALIDATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// example:
	//
	// expression
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// col
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

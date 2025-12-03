// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemAllFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListWorkItemAllFieldsResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListWorkItemAllFieldsResponseBody
	GetErrorMsg() *string
	SetFields(v []*ListWorkItemAllFieldsResponseBodyFields) *ListWorkItemAllFieldsResponseBody
	GetFields() []*ListWorkItemAllFieldsResponseBodyFields
	SetRequestId(v string) *ListWorkItemAllFieldsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkItemAllFieldsResponseBody
	GetSuccess() *bool
}

type ListWorkItemAllFieldsResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string                                    `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Fields   []*ListWorkItemAllFieldsResponseBodyFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListWorkItemAllFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemAllFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListWorkItemAllFieldsResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListWorkItemAllFieldsResponseBody) GetFields() []*ListWorkItemAllFieldsResponseBodyFields {
	return s.Fields
}

func (s *ListWorkItemAllFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkItemAllFieldsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkItemAllFieldsResponseBody) SetErrorCode(v string) *ListWorkItemAllFieldsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBody) SetErrorMsg(v string) *ListWorkItemAllFieldsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBody) SetFields(v []*ListWorkItemAllFieldsResponseBodyFields) *ListWorkItemAllFieldsResponseBody {
	s.Fields = v
	return s
}

func (s *ListWorkItemAllFieldsResponseBody) SetRequestId(v string) *ListWorkItemAllFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBody) SetSuccess(v bool) *ListWorkItemAllFieldsResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBody) Validate() error {
	if s.Fields != nil {
		for _, item := range s.Fields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkItemAllFieldsResponseBodyFields struct {
	// example:
	//
	// 例：date
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 123
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// example:
	//
	// 字段的具体信息
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// list
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// 1623916393000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1623916393000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// false
	IsRequired *bool `json:"isRequired,omitempty" xml:"isRequired,omitempty"`
	// example:
	//
	// false
	IsShowWhenCreate *bool `json:"isShowWhenCreate,omitempty" xml:"isShowWhenCreate,omitempty"`
	// example:
	//
	// false
	IsSystemRequired *bool `json:"isSystemRequired,omitempty" xml:"isSystemRequired,omitempty"`
	// example:
	//
	// null
	LinkWithService *string `json:"linkWithService,omitempty" xml:"linkWithService,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 负责人
	Name    *string                                           `json:"name,omitempty" xml:"name,omitempty"`
	Options []*ListWorkItemAllFieldsResponseBodyFieldsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// example:
	//
	// 例：Workitem
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// NativeField
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListWorkItemAllFieldsResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemAllFieldsResponseBodyFields) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetClassName() *string {
	return s.ClassName
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetCreator() *string {
	return s.Creator
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetDescription() *string {
	return s.Description
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetFormat() *string {
	return s.Format
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetIsShowWhenCreate() *bool {
	return s.IsShowWhenCreate
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetIsSystemRequired() *bool {
	return s.IsSystemRequired
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetLinkWithService() *string {
	return s.LinkWithService
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetModifier() *string {
	return s.Modifier
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetName() *string {
	return s.Name
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetOptions() []*ListWorkItemAllFieldsResponseBodyFieldsOptions {
	return s.Options
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListWorkItemAllFieldsResponseBodyFields) GetType() *string {
	return s.Type
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetClassName(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.ClassName = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetCreator(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Creator = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetDefaultValue(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.DefaultValue = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetDescription(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Description = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetFormat(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Format = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetGmtCreate(v int64) *ListWorkItemAllFieldsResponseBodyFields {
	s.GmtCreate = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetGmtModified(v int64) *ListWorkItemAllFieldsResponseBodyFields {
	s.GmtModified = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetIdentifier(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Identifier = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetIsRequired(v bool) *ListWorkItemAllFieldsResponseBodyFields {
	s.IsRequired = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetIsShowWhenCreate(v bool) *ListWorkItemAllFieldsResponseBodyFields {
	s.IsShowWhenCreate = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetIsSystemRequired(v bool) *ListWorkItemAllFieldsResponseBodyFields {
	s.IsSystemRequired = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetLinkWithService(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.LinkWithService = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetModifier(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Modifier = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetName(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Name = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetOptions(v []*ListWorkItemAllFieldsResponseBodyFieldsOptions) *ListWorkItemAllFieldsResponseBodyFields {
	s.Options = v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetResourceType(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.ResourceType = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) SetType(v string) *ListWorkItemAllFieldsResponseBodyFields {
	s.Type = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFields) Validate() error {
	if s.Options != nil {
		for _, item := range s.Options {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkItemAllFieldsResponseBodyFieldsOptions struct {
	// example:
	//
	// 重复的缺陷
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// example:
	//
	// 重复的缺陷
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 1
	Level *int64 `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// 重复的缺陷
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// null
	ValueEn *string `json:"valueEn,omitempty" xml:"valueEn,omitempty"`
}

func (s ListWorkItemAllFieldsResponseBodyFieldsOptions) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemAllFieldsResponseBodyFieldsOptions) GoString() string {
	return s.String()
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) GetDisplayValue() *string {
	return s.DisplayValue
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) GetLevel() *int64 {
	return s.Level
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) GetPosition() *int64 {
	return s.Position
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) GetValue() *string {
	return s.Value
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) GetValueEn() *string {
	return s.ValueEn
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetDisplayValue(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.DisplayValue = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetFieldIdentifier(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.FieldIdentifier = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetIdentifier(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.Identifier = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetLevel(v int64) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.Level = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetPosition(v int64) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.Position = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetValue(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.Value = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) SetValueEn(v string) *ListWorkItemAllFieldsResponseBodyFieldsOptions {
	s.ValueEn = &v
	return s
}

func (s *ListWorkItemAllFieldsResponseBodyFieldsOptions) Validate() error {
	return dara.Validate(s)
}

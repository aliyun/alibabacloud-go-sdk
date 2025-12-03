// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTestCaseFieldsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListTestCaseFieldsResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListTestCaseFieldsResponseBody
	GetErrorMsg() *string
	SetFields(v []*ListTestCaseFieldsResponseBodyFields) *ListTestCaseFieldsResponseBody
	GetFields() []*ListTestCaseFieldsResponseBodyFields
	SetRequestId(v string) *ListTestCaseFieldsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTestCaseFieldsResponseBody
	GetSuccess() *bool
}

type ListTestCaseFieldsResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string                                 `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Fields   []*ListTestCaseFieldsResponseBodyFields `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTestCaseFieldsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTestCaseFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTestCaseFieldsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTestCaseFieldsResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListTestCaseFieldsResponseBody) GetFields() []*ListTestCaseFieldsResponseBodyFields {
	return s.Fields
}

func (s *ListTestCaseFieldsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTestCaseFieldsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTestCaseFieldsResponseBody) SetErrorCode(v string) *ListTestCaseFieldsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTestCaseFieldsResponseBody) SetErrorMsg(v string) *ListTestCaseFieldsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListTestCaseFieldsResponseBody) SetFields(v []*ListTestCaseFieldsResponseBodyFields) *ListTestCaseFieldsResponseBody {
	s.Fields = v
	return s
}

func (s *ListTestCaseFieldsResponseBody) SetRequestId(v string) *ListTestCaseFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTestCaseFieldsResponseBody) SetSuccess(v bool) *ListTestCaseFieldsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTestCaseFieldsResponseBody) Validate() error {
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

type ListTestCaseFieldsResponseBodyFields struct {
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
	Name    *string                                        `json:"name,omitempty" xml:"name,omitempty"`
	Options []*ListTestCaseFieldsResponseBodyFieldsOptions `json:"options,omitempty" xml:"options,omitempty" type:"Repeated"`
	// example:
	//
	// 例：Workitem
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// NativeField
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTestCaseFieldsResponseBodyFields) String() string {
	return dara.Prettify(s)
}

func (s ListTestCaseFieldsResponseBodyFields) GoString() string {
	return s.String()
}

func (s *ListTestCaseFieldsResponseBodyFields) GetClassName() *string {
	return s.ClassName
}

func (s *ListTestCaseFieldsResponseBodyFields) GetCreator() *string {
	return s.Creator
}

func (s *ListTestCaseFieldsResponseBodyFields) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListTestCaseFieldsResponseBodyFields) GetDescription() *string {
	return s.Description
}

func (s *ListTestCaseFieldsResponseBodyFields) GetFormat() *string {
	return s.Format
}

func (s *ListTestCaseFieldsResponseBodyFields) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListTestCaseFieldsResponseBodyFields) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListTestCaseFieldsResponseBodyFields) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListTestCaseFieldsResponseBodyFields) GetIsRequired() *bool {
	return s.IsRequired
}

func (s *ListTestCaseFieldsResponseBodyFields) GetIsShowWhenCreate() *bool {
	return s.IsShowWhenCreate
}

func (s *ListTestCaseFieldsResponseBodyFields) GetIsSystemRequired() *bool {
	return s.IsSystemRequired
}

func (s *ListTestCaseFieldsResponseBodyFields) GetLinkWithService() *string {
	return s.LinkWithService
}

func (s *ListTestCaseFieldsResponseBodyFields) GetModifier() *string {
	return s.Modifier
}

func (s *ListTestCaseFieldsResponseBodyFields) GetName() *string {
	return s.Name
}

func (s *ListTestCaseFieldsResponseBodyFields) GetOptions() []*ListTestCaseFieldsResponseBodyFieldsOptions {
	return s.Options
}

func (s *ListTestCaseFieldsResponseBodyFields) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTestCaseFieldsResponseBodyFields) GetType() *string {
	return s.Type
}

func (s *ListTestCaseFieldsResponseBodyFields) SetClassName(v string) *ListTestCaseFieldsResponseBodyFields {
	s.ClassName = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetCreator(v string) *ListTestCaseFieldsResponseBodyFields {
	s.Creator = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetDefaultValue(v string) *ListTestCaseFieldsResponseBodyFields {
	s.DefaultValue = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetDescription(v string) *ListTestCaseFieldsResponseBodyFields {
	s.Description = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetFormat(v string) *ListTestCaseFieldsResponseBodyFields {
	s.Format = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetGmtCreate(v int64) *ListTestCaseFieldsResponseBodyFields {
	s.GmtCreate = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetGmtModified(v int64) *ListTestCaseFieldsResponseBodyFields {
	s.GmtModified = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetIdentifier(v string) *ListTestCaseFieldsResponseBodyFields {
	s.Identifier = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetIsRequired(v bool) *ListTestCaseFieldsResponseBodyFields {
	s.IsRequired = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetIsShowWhenCreate(v bool) *ListTestCaseFieldsResponseBodyFields {
	s.IsShowWhenCreate = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetIsSystemRequired(v bool) *ListTestCaseFieldsResponseBodyFields {
	s.IsSystemRequired = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetLinkWithService(v string) *ListTestCaseFieldsResponseBodyFields {
	s.LinkWithService = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetModifier(v string) *ListTestCaseFieldsResponseBodyFields {
	s.Modifier = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetName(v string) *ListTestCaseFieldsResponseBodyFields {
	s.Name = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetOptions(v []*ListTestCaseFieldsResponseBodyFieldsOptions) *ListTestCaseFieldsResponseBodyFields {
	s.Options = v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetResourceType(v string) *ListTestCaseFieldsResponseBodyFields {
	s.ResourceType = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) SetType(v string) *ListTestCaseFieldsResponseBodyFields {
	s.Type = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFields) Validate() error {
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

type ListTestCaseFieldsResponseBodyFieldsOptions struct {
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
	// null
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

func (s ListTestCaseFieldsResponseBodyFieldsOptions) String() string {
	return dara.Prettify(s)
}

func (s ListTestCaseFieldsResponseBodyFieldsOptions) GoString() string {
	return s.String()
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) GetDisplayValue() *string {
	return s.DisplayValue
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) GetLevel() *int64 {
	return s.Level
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) GetPosition() *int64 {
	return s.Position
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) GetValue() *string {
	return s.Value
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) GetValueEn() *string {
	return s.ValueEn
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) SetDisplayValue(v string) *ListTestCaseFieldsResponseBodyFieldsOptions {
	s.DisplayValue = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) SetFieldIdentifier(v string) *ListTestCaseFieldsResponseBodyFieldsOptions {
	s.FieldIdentifier = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) SetIdentifier(v string) *ListTestCaseFieldsResponseBodyFieldsOptions {
	s.Identifier = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) SetLevel(v int64) *ListTestCaseFieldsResponseBodyFieldsOptions {
	s.Level = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) SetPosition(v int64) *ListTestCaseFieldsResponseBodyFieldsOptions {
	s.Position = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) SetValue(v string) *ListTestCaseFieldsResponseBodyFieldsOptions {
	s.Value = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) SetValueEn(v string) *ListTestCaseFieldsResponseBodyFieldsOptions {
	s.ValueEn = &v
	return s
}

func (s *ListTestCaseFieldsResponseBodyFieldsOptions) Validate() error {
	return dara.Validate(s)
}

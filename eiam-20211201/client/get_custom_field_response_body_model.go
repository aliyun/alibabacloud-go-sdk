// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomField(v *GetCustomFieldResponseBodyCustomField) *GetCustomFieldResponseBody
	GetCustomField() *GetCustomFieldResponseBodyCustomField
	SetRequestId(v string) *GetCustomFieldResponseBody
	GetRequestId() *string
}

type GetCustomFieldResponseBody struct {
	CustomField *GetCustomFieldResponseBodyCustomField `json:"CustomField,omitempty" xml:"CustomField,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCustomFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomFieldResponseBody) GetCustomField() *GetCustomFieldResponseBodyCustomField {
	return s.CustomField
}

func (s *GetCustomFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomFieldResponseBody) SetCustomField(v *GetCustomFieldResponseBodyCustomField) *GetCustomFieldResponseBody {
	s.CustomField = v
	return s
}

func (s *GetCustomFieldResponseBody) SetRequestId(v string) *GetCustomFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomFieldResponseBody) Validate() error {
	if s.CustomField != nil {
		if err := s.CustomField.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomFieldResponseBodyCustomField struct {
	// 扩展字段创建时间，Unix时间戳格式，单位为毫秒
	//
	// example:
	//
	// 17642960730
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 字段默认值
	//
	// example:
	//
	// test
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// 扩展字段描述
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否加密，默认false
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// 字段所属实体类型
	//
	// example:
	//
	// user
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// 字段值配置项
	FieldDataConfig *GetCustomFieldResponseBodyCustomFieldFieldDataConfig `json:"FieldDataConfig,omitempty" xml:"FieldDataConfig,omitempty" type:"Struct"`
	// 数据类型，枚举值：string、number、boolean
	//
	// example:
	//
	// string
	FieldDataType *string `json:"FieldDataType,omitempty" xml:"FieldDataType,omitempty"`
	// 字段展示名
	//
	// example:
	//
	// userId
	FieldDisplayName *string `json:"FieldDisplayName,omitempty" xml:"FieldDisplayName,omitempty"`
	// 字段展示类型，枚举值，select、checkbox、input、number
	//
	// example:
	//
	// input
	FieldDisplayType *string `json:"FieldDisplayType,omitempty" xml:"FieldDisplayType,omitempty"`
	// 字段ID
	//
	// example:
	//
	// ufd_ncvy5trszg3zajaal5iofauy2q
	FieldId *string `json:"FieldId,omitempty" xml:"FieldId,omitempty"`
	// 字段标识
	//
	// example:
	//
	// userId
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// 实例ID
	//
	// example:
	//
	// idaas_z4pwq7v5ankdimdelzo2zbmzo4
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 是否必填，默认false
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 扩展字段状态, 枚举类型：enabled，disabled
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 是否唯一，默认false
	//
	// example:
	//
	// false
	Unique *bool `json:"Unique,omitempty" xml:"Unique,omitempty"`
	// 扩展字段最近一次更新时间, Unix时间戳格式，单位为毫秒
	//
	// example:
	//
	// 17642960730
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 用户端(portal侧)权限，hide、read_only、read_write，默认read_only
	//
	// example:
	//
	// read_only
	UserPermission *string `json:"UserPermission,omitempty" xml:"UserPermission,omitempty"`
}

func (s GetCustomFieldResponseBodyCustomField) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldResponseBodyCustomField) GoString() string {
	return s.String()
}

func (s *GetCustomFieldResponseBodyCustomField) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCustomFieldResponseBodyCustomField) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetCustomFieldResponseBodyCustomField) GetDescription() *string {
	return s.Description
}

func (s *GetCustomFieldResponseBodyCustomField) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *GetCustomFieldResponseBodyCustomField) GetEntityType() *string {
	return s.EntityType
}

func (s *GetCustomFieldResponseBodyCustomField) GetFieldDataConfig() *GetCustomFieldResponseBodyCustomFieldFieldDataConfig {
	return s.FieldDataConfig
}

func (s *GetCustomFieldResponseBodyCustomField) GetFieldDataType() *string {
	return s.FieldDataType
}

func (s *GetCustomFieldResponseBodyCustomField) GetFieldDisplayName() *string {
	return s.FieldDisplayName
}

func (s *GetCustomFieldResponseBodyCustomField) GetFieldDisplayType() *string {
	return s.FieldDisplayType
}

func (s *GetCustomFieldResponseBodyCustomField) GetFieldId() *string {
	return s.FieldId
}

func (s *GetCustomFieldResponseBodyCustomField) GetFieldName() *string {
	return s.FieldName
}

func (s *GetCustomFieldResponseBodyCustomField) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCustomFieldResponseBodyCustomField) GetRequired() *bool {
	return s.Required
}

func (s *GetCustomFieldResponseBodyCustomField) GetStatus() *string {
	return s.Status
}

func (s *GetCustomFieldResponseBodyCustomField) GetUnique() *bool {
	return s.Unique
}

func (s *GetCustomFieldResponseBodyCustomField) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetCustomFieldResponseBodyCustomField) GetUserPermission() *string {
	return s.UserPermission
}

func (s *GetCustomFieldResponseBodyCustomField) SetCreateTime(v int64) *GetCustomFieldResponseBodyCustomField {
	s.CreateTime = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetDefaultValue(v string) *GetCustomFieldResponseBodyCustomField {
	s.DefaultValue = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetDescription(v string) *GetCustomFieldResponseBodyCustomField {
	s.Description = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetEncrypted(v bool) *GetCustomFieldResponseBodyCustomField {
	s.Encrypted = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetEntityType(v string) *GetCustomFieldResponseBodyCustomField {
	s.EntityType = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetFieldDataConfig(v *GetCustomFieldResponseBodyCustomFieldFieldDataConfig) *GetCustomFieldResponseBodyCustomField {
	s.FieldDataConfig = v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetFieldDataType(v string) *GetCustomFieldResponseBodyCustomField {
	s.FieldDataType = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetFieldDisplayName(v string) *GetCustomFieldResponseBodyCustomField {
	s.FieldDisplayName = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetFieldDisplayType(v string) *GetCustomFieldResponseBodyCustomField {
	s.FieldDisplayType = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetFieldId(v string) *GetCustomFieldResponseBodyCustomField {
	s.FieldId = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetFieldName(v string) *GetCustomFieldResponseBodyCustomField {
	s.FieldName = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetInstanceId(v string) *GetCustomFieldResponseBodyCustomField {
	s.InstanceId = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetRequired(v bool) *GetCustomFieldResponseBodyCustomField {
	s.Required = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetStatus(v string) *GetCustomFieldResponseBodyCustomField {
	s.Status = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetUnique(v bool) *GetCustomFieldResponseBodyCustomField {
	s.Unique = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetUpdateTime(v int64) *GetCustomFieldResponseBodyCustomField {
	s.UpdateTime = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) SetUserPermission(v string) *GetCustomFieldResponseBodyCustomField {
	s.UserPermission = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomField) Validate() error {
	if s.FieldDataConfig != nil {
		if err := s.FieldDataConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomFieldResponseBodyCustomFieldFieldDataConfig struct {
	// 字段配置项列表，displayName、value、status
	Items []*GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s GetCustomFieldResponseBodyCustomFieldFieldDataConfig) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldResponseBodyCustomFieldFieldDataConfig) GoString() string {
	return s.String()
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfig) GetItems() []*GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems {
	return s.Items
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfig) SetItems(v []*GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) *GetCustomFieldResponseBodyCustomFieldFieldDataConfig {
	s.Items = v
	return s
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfig) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems struct {
	// 配置项展示名
	//
	// example:
	//
	// hobby
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 配置项状态, 枚举值，enabled、disabled
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 配置项值
	//
	// example:
	//
	// game
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) GoString() string {
	return s.String()
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) GetStatus() *string {
	return s.Status
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) GetValue() *string {
	return s.Value
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) SetDisplayName(v string) *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems {
	s.DisplayName = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) SetStatus(v string) *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems {
	s.Status = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) SetValue(v string) *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems {
	s.Value = &v
	return s
}

func (s *GetCustomFieldResponseBodyCustomFieldFieldDataConfigItems) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v string) *CreateCustomFieldRequest
	GetDefaultValue() *string
	SetDescription(v string) *CreateCustomFieldRequest
	GetDescription() *string
	SetEncrypted(v bool) *CreateCustomFieldRequest
	GetEncrypted() *bool
	SetEntityType(v string) *CreateCustomFieldRequest
	GetEntityType() *string
	SetFieldDataConfig(v *CreateCustomFieldRequestFieldDataConfig) *CreateCustomFieldRequest
	GetFieldDataConfig() *CreateCustomFieldRequestFieldDataConfig
	SetFieldDataType(v string) *CreateCustomFieldRequest
	GetFieldDataType() *string
	SetFieldDisplayName(v string) *CreateCustomFieldRequest
	GetFieldDisplayName() *string
	SetFieldDisplayType(v string) *CreateCustomFieldRequest
	GetFieldDisplayType() *string
	SetFieldName(v string) *CreateCustomFieldRequest
	GetFieldName() *string
	SetInstanceId(v string) *CreateCustomFieldRequest
	GetInstanceId() *string
	SetRequired(v bool) *CreateCustomFieldRequest
	GetRequired() *bool
	SetUnique(v bool) *CreateCustomFieldRequest
	GetUnique() *bool
	SetUserPermission(v string) *CreateCustomFieldRequest
	GetUserPermission() *string
}

type CreateCustomFieldRequest struct {
	// 字段默认值，必须与数据类型一致
	//
	// example:
	//
	// string
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// 对字段的描述信息
	//
	// example:
	//
	// 字段测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 是否加密，默认false
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// 字段归属实体。实体包括账户、组、组织
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// 字段值配置项，必须与数据类型一致
	FieldDataConfig *CreateCustomFieldRequestFieldDataConfig `json:"FieldDataConfig,omitempty" xml:"FieldDataConfig,omitempty" type:"Struct"`
	// 数据类型，枚举值：string、number、boolean
	//
	// This parameter is required.
	//
	// example:
	//
	// string
	FieldDataType *string `json:"FieldDataType,omitempty" xml:"FieldDataType,omitempty"`
	// 字段展示名，长度不超过128字符
	//
	// This parameter is required.
	//
	// example:
	//
	// name_001
	FieldDisplayName *string `json:"FieldDisplayName,omitempty" xml:"FieldDisplayName,omitempty"`
	// 字段展示类型，枚举值，select、checkbox、input
	//
	// This parameter is required.
	//
	// example:
	//
	// input
	FieldDisplayType *string `json:"FieldDisplayType,omitempty" xml:"FieldDisplayType,omitempty"`
	// 字段标识，英文字母、下划线
	//
	// This parameter is required.
	//
	// example:
	//
	// field_001
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 是否必填，默认false
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// 是否唯一，默认false
	//
	// example:
	//
	// false
	Unique *bool `json:"Unique,omitempty" xml:"Unique,omitempty"`
	// 用户端(portal侧)权限，hide、read_only、read_write，默认read_only
	//
	// example:
	//
	// false
	UserPermission *string `json:"UserPermission,omitempty" xml:"UserPermission,omitempty"`
}

func (s CreateCustomFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomFieldRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomFieldRequest) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *CreateCustomFieldRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomFieldRequest) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateCustomFieldRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *CreateCustomFieldRequest) GetFieldDataConfig() *CreateCustomFieldRequestFieldDataConfig {
	return s.FieldDataConfig
}

func (s *CreateCustomFieldRequest) GetFieldDataType() *string {
	return s.FieldDataType
}

func (s *CreateCustomFieldRequest) GetFieldDisplayName() *string {
	return s.FieldDisplayName
}

func (s *CreateCustomFieldRequest) GetFieldDisplayType() *string {
	return s.FieldDisplayType
}

func (s *CreateCustomFieldRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateCustomFieldRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomFieldRequest) GetRequired() *bool {
	return s.Required
}

func (s *CreateCustomFieldRequest) GetUnique() *bool {
	return s.Unique
}

func (s *CreateCustomFieldRequest) GetUserPermission() *string {
	return s.UserPermission
}

func (s *CreateCustomFieldRequest) SetDefaultValue(v string) *CreateCustomFieldRequest {
	s.DefaultValue = &v
	return s
}

func (s *CreateCustomFieldRequest) SetDescription(v string) *CreateCustomFieldRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomFieldRequest) SetEncrypted(v bool) *CreateCustomFieldRequest {
	s.Encrypted = &v
	return s
}

func (s *CreateCustomFieldRequest) SetEntityType(v string) *CreateCustomFieldRequest {
	s.EntityType = &v
	return s
}

func (s *CreateCustomFieldRequest) SetFieldDataConfig(v *CreateCustomFieldRequestFieldDataConfig) *CreateCustomFieldRequest {
	s.FieldDataConfig = v
	return s
}

func (s *CreateCustomFieldRequest) SetFieldDataType(v string) *CreateCustomFieldRequest {
	s.FieldDataType = &v
	return s
}

func (s *CreateCustomFieldRequest) SetFieldDisplayName(v string) *CreateCustomFieldRequest {
	s.FieldDisplayName = &v
	return s
}

func (s *CreateCustomFieldRequest) SetFieldDisplayType(v string) *CreateCustomFieldRequest {
	s.FieldDisplayType = &v
	return s
}

func (s *CreateCustomFieldRequest) SetFieldName(v string) *CreateCustomFieldRequest {
	s.FieldName = &v
	return s
}

func (s *CreateCustomFieldRequest) SetInstanceId(v string) *CreateCustomFieldRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomFieldRequest) SetRequired(v bool) *CreateCustomFieldRequest {
	s.Required = &v
	return s
}

func (s *CreateCustomFieldRequest) SetUnique(v bool) *CreateCustomFieldRequest {
	s.Unique = &v
	return s
}

func (s *CreateCustomFieldRequest) SetUserPermission(v string) *CreateCustomFieldRequest {
	s.UserPermission = &v
	return s
}

func (s *CreateCustomFieldRequest) Validate() error {
	if s.FieldDataConfig != nil {
		if err := s.FieldDataConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomFieldRequestFieldDataConfig struct {
	// 字段值配置项，必须与数据类型一致，只能新增数据项，不可删除，项字段：displayName、value、status
	//
	// example:
	//
	// string
	Items []*CreateCustomFieldRequestFieldDataConfigItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s CreateCustomFieldRequestFieldDataConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomFieldRequestFieldDataConfig) GoString() string {
	return s.String()
}

func (s *CreateCustomFieldRequestFieldDataConfig) GetItems() []*CreateCustomFieldRequestFieldDataConfigItems {
	return s.Items
}

func (s *CreateCustomFieldRequestFieldDataConfig) SetItems(v []*CreateCustomFieldRequestFieldDataConfigItems) *CreateCustomFieldRequestFieldDataConfig {
	s.Items = v
	return s
}

func (s *CreateCustomFieldRequestFieldDataConfig) Validate() error {
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

type CreateCustomFieldRequestFieldDataConfigItems struct {
	// 配置项展示名
	//
	// example:
	//
	// string
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 配置项状态，枚举值，enabled、disabled
	//
	// example:
	//
	// string
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 配置项展示值
	//
	// example:
	//
	// string
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCustomFieldRequestFieldDataConfigItems) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomFieldRequestFieldDataConfigItems) GoString() string {
	return s.String()
}

func (s *CreateCustomFieldRequestFieldDataConfigItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateCustomFieldRequestFieldDataConfigItems) GetStatus() *string {
	return s.Status
}

func (s *CreateCustomFieldRequestFieldDataConfigItems) GetValue() *string {
	return s.Value
}

func (s *CreateCustomFieldRequestFieldDataConfigItems) SetDisplayName(v string) *CreateCustomFieldRequestFieldDataConfigItems {
	s.DisplayName = &v
	return s
}

func (s *CreateCustomFieldRequestFieldDataConfigItems) SetStatus(v string) *CreateCustomFieldRequestFieldDataConfigItems {
	s.Status = &v
	return s
}

func (s *CreateCustomFieldRequestFieldDataConfigItems) SetValue(v string) *CreateCustomFieldRequestFieldDataConfigItems {
	s.Value = &v
	return s
}

func (s *CreateCustomFieldRequestFieldDataConfigItems) Validate() error {
	return dara.Validate(s)
}

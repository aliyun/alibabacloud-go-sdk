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
	// The default value of the field. If the data type has configuration items, the default value must be one of the enabled configuration items. The maximum length is 1,024 characters.
	//
	// example:
	//
	// string
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The description of the field. The maximum length is 512 characters.
	//
	// example:
	//
	// Test field
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the field is encrypted. If you set this parameter to true, the underlying system encrypts and stores the data value.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The entity to which the field belongs. Valid values:
	//
	// - user: account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The configuration items for the field value.
	FieldDataConfig *CreateCustomFieldRequestFieldDataConfig `json:"FieldDataConfig,omitempty" xml:"FieldDataConfig,omitempty" type:"Struct"`
	// The data type of the field. Valid values:
	//
	// - string: string.
	//
	// - number: number. The maximum length is 32 characters. Positive integers and decimals are supported.
	//
	// - boolean: Boolean value.
	//
	// This parameter is required.
	//
	// example:
	//
	// string
	FieldDataType *string `json:"FieldDataType,omitempty" xml:"FieldDataType,omitempty"`
	// The field display name. The maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// name_001
	FieldDisplayName *string `json:"FieldDisplayName,omitempty" xml:"FieldDisplayName,omitempty"`
	// The field display type. Valid values:
	//
	// - input: text box. The supported data types are string and number.
	//
	// - select: drop-down list. The supported data types are string and boolean.
	//
	// - checkbox: check box. The supported data type is string.
	//
	// This parameter is required.
	//
	// example:
	//
	// input
	FieldDisplayType *string `json:"FieldDisplayType,omitempty" xml:"FieldDisplayType,omitempty"`
	// The field identifier. The maximum length is 40 characters. The value can contain lowercase English letters and underscores, and cannot start with an underscore.
	//
	// This parameter is required.
	//
	// example:
	//
	// field_001
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the field is required.
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// Specifies whether the field value is unique. If you set this parameter to true, the value of this field must be unique within the corresponding entity type and cannot be duplicated.
	//
	// example:
	//
	// false
	Unique *bool `json:"Unique,omitempty" xml:"Unique,omitempty"`
	// The field permissions on the user portal. Valid values:
	//
	// - hide: invisible on the user portal
	//
	// - read_only: visible but not editable on the user portal
	//
	// - read_write: visible and editable on the user portal
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
	// The list of field configuration items. The maximum number of items is 100.
	//
	// example:
	//
	// [{"DisplayName":"Option 1","Value":"option1","Status":"enabled"}]
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
	// The display name of the configuration item. The maximum length is 128 characters.
	//
	// example:
	//
	// string
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The status of the configuration item. Valid values:
	//
	// - enabled: enabled.
	//
	// - disabled: disabled.
	//
	// If a configuration item is disabled, it is unavailable when you create or update the field value of the entity.
	//
	// example:
	//
	// string
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The value of the configuration item. The maximum length is 64 characters.
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

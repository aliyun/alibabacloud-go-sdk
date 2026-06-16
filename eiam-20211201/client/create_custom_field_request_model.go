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
	// The default value of the field. If the field has configuration items, the default value must be one of the enabled configuration items. The default value can be up to 1024 characters in length.
	//
	// example:
	//
	// string
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The description of the field. The description can be up to 512 characters in length.
	//
	// example:
	//
	// Field test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether to encrypt the field value. If you set this parameter to true, the system encrypts the data value before storing it.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The entity to which the field belongs. Valid value:
	//
	// - user: an account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The configuration items of the field value.
	FieldDataConfig *CreateCustomFieldRequestFieldDataConfig `json:"FieldDataConfig,omitempty" xml:"FieldDataConfig,omitempty" type:"Struct"`
	// The data type of the field. Valid values:
	//
	// - string: a string.
	//
	// - number: a number. The number can be up to 32 digits in length and can be a positive integer or a decimal.
	//
	// - boolean: a Boolean value.
	//
	// This parameter is required.
	//
	// example:
	//
	// string
	FieldDataType *string `json:"FieldDataType,omitempty" xml:"FieldDataType,omitempty"`
	// The display name of the field. The display name can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// name_001
	FieldDisplayName *string `json:"FieldDisplayName,omitempty" xml:"FieldDisplayName,omitempty"`
	// The display type of the field. Valid values:
	//
	// - input: a text box. This display type supports the string and number data types.
	//
	// - select: a drop-down list. This display type supports the string and Boolean data types.
	//
	// - checkbox: a check box. This display type supports the string data type.
	//
	// This parameter is required.
	//
	// example:
	//
	// input
	FieldDisplayType *string `json:"FieldDisplayType,omitempty" xml:"FieldDisplayType,omitempty"`
	// The name of the field. The name can be up to 40 characters in length and can contain lowercase letters and underscores (_). It cannot start with an underscore (_).
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
	// Indicates whether the field is required.
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
	// Indicates whether the field value is unique. If you set this parameter to true, the value of this field must be unique for the specified entity type.
	//
	// example:
	//
	// false
	Unique *bool `json:"Unique,omitempty" xml:"Unique,omitempty"`
	// The permission on the field in the portal. Valid values:
	//
	// - hide: The field is not visible in the portal.
	//
	// - read_only: The field is visible but cannot be modified in the portal.
	//
	// - read_write: The field is visible and can be modified in the portal.
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
	// A list of field configuration items. The list can contain up to 100 items.
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
	// The display name of the configuration item. The display name can be up to 128 characters in length.
	//
	// example:
	//
	// string
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The status of the configuration item. Valid values:
	//
	// - enabled: The configuration item is enabled.
	//
	// - disabled: The configuration item is disabled.
	//
	// If a configuration item is disabled, it is unavailable when you create or update the field value for an entity.
	//
	// example:
	//
	// string
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The value of the configuration item. The value can be up to 64 characters in length.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomFields(v []*UpdateUserRequestCustomFields) *UpdateUserRequest
	GetCustomFields() []*UpdateUserRequestCustomFields
	SetDisplayName(v string) *UpdateUserRequest
	GetDisplayName() *string
	SetEmail(v string) *UpdateUserRequest
	GetEmail() *string
	SetEmailVerified(v bool) *UpdateUserRequest
	GetEmailVerified() *bool
	SetInstanceId(v string) *UpdateUserRequest
	GetInstanceId() *string
	SetPhoneNumber(v string) *UpdateUserRequest
	GetPhoneNumber() *string
	SetPhoneNumberVerified(v bool) *UpdateUserRequest
	GetPhoneNumberVerified() *bool
	SetPhoneRegion(v string) *UpdateUserRequest
	GetPhoneRegion() *string
	SetUserId(v string) *UpdateUserRequest
	GetUserId() *string
	SetUsername(v string) *UpdateUserRequest
	GetUsername() *string
}

type UpdateUserRequest struct {
	// The list of custom field objects.
	CustomFields []*UpdateUserRequestCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// The display name of the account. The name can be up to 256 characters in length.
	//
	// example:
	//
	// test_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address. The prefix of the email address can contain uppercase letters, lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Specifies whether the email address is verified. This parameter is required if an email address is specified. If no special business requirements exist, set this parameter to true.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"EmailVerified,omitempty" xml:"EmailVerified,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The phone number. The value must be 6 to 15 digits in length.
	//
	// example:
	//
	// 156xxxxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Specifies whether the phone number is verified as a trusted phone number. This parameter is required if a phone number is specified. If no special business requirements exist, set this parameter to true.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The phone region code. Example: 86 for the Chinese mainland, without the 00 or + prefix. This parameter is required if a phone number is specified.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The account name. The name can contain letters, digits, underscores (_), periods (.), at signs (@), and hyphens (-). The name can be up to 256 characters in length.
	//
	// example:
	//
	// username_test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetCustomFields() []*UpdateUserRequestCustomFields {
	return s.CustomFields
}

func (s *UpdateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *UpdateUserRequest) GetEmailVerified() *bool {
	return s.EmailVerified
}

func (s *UpdateUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateUserRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *UpdateUserRequest) GetPhoneNumberVerified() *bool {
	return s.PhoneNumberVerified
}

func (s *UpdateUserRequest) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *UpdateUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserRequest) GetUsername() *string {
	return s.Username
}

func (s *UpdateUserRequest) SetCustomFields(v []*UpdateUserRequestCustomFields) *UpdateUserRequest {
	s.CustomFields = v
	return s
}

func (s *UpdateUserRequest) SetDisplayName(v string) *UpdateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetEmail(v string) *UpdateUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserRequest) SetEmailVerified(v bool) *UpdateUserRequest {
	s.EmailVerified = &v
	return s
}

func (s *UpdateUserRequest) SetInstanceId(v string) *UpdateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserRequest) SetPhoneNumber(v string) *UpdateUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateUserRequest) SetPhoneNumberVerified(v bool) *UpdateUserRequest {
	s.PhoneNumberVerified = &v
	return s
}

func (s *UpdateUserRequest) SetPhoneRegion(v string) *UpdateUserRequest {
	s.PhoneRegion = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) SetUsername(v string) *UpdateUserRequest {
	s.Username = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	if s.CustomFields != nil {
		for _, item := range s.CustomFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateUserRequestCustomFields struct {
	// The identifier of the custom field. Create the custom field in advance. For more information, refer to the custom fields module in the console.
	//
	// example:
	//
	// nick_name
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the custom field. The value must comply with the property constraints of the corresponding custom field.
	//
	// example:
	//
	// test_value
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The operation type for the custom field. Valid values:
	//
	// - add: adds a custom field value to the account.
	//
	// - replace: replaces an existing custom field value of the account. If the custom field value does not exist, the operation is converted to an add operation.
	//
	// - remove: removes the custom field value from the account.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s UpdateUserRequestCustomFields) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequestCustomFields) GoString() string {
	return s.String()
}

func (s *UpdateUserRequestCustomFields) GetFieldName() *string {
	return s.FieldName
}

func (s *UpdateUserRequestCustomFields) GetFieldValue() *string {
	return s.FieldValue
}

func (s *UpdateUserRequestCustomFields) GetOperation() *string {
	return s.Operation
}

func (s *UpdateUserRequestCustomFields) SetFieldName(v string) *UpdateUserRequestCustomFields {
	s.FieldName = &v
	return s
}

func (s *UpdateUserRequestCustomFields) SetFieldValue(v string) *UpdateUserRequestCustomFields {
	s.FieldValue = &v
	return s
}

func (s *UpdateUserRequestCustomFields) SetOperation(v string) *UpdateUserRequestCustomFields {
	s.Operation = &v
	return s
}

func (s *UpdateUserRequestCustomFields) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomFields(v []*PatchUserRequestCustomFields) *PatchUserRequest
	GetCustomFields() []*PatchUserRequestCustomFields
	SetDisplayName(v string) *PatchUserRequest
	GetDisplayName() *string
	SetEmail(v string) *PatchUserRequest
	GetEmail() *string
	SetEmailVerified(v bool) *PatchUserRequest
	GetEmailVerified() *bool
	SetPhoneNumber(v string) *PatchUserRequest
	GetPhoneNumber() *string
	SetPhoneNumberVerified(v bool) *PatchUserRequest
	GetPhoneNumberVerified() *bool
	SetPhoneRegion(v string) *PatchUserRequest
	GetPhoneRegion() *string
	SetUsername(v string) *PatchUserRequest
	GetUsername() *string
}

type PatchUserRequest struct {
	// The extended fields of the account.
	CustomFields []*PatchUserRequestCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// The display name of the account.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Specifies whether the email address is verified. This field is required if an email address is specified. If you have no special requirement, set this parameter to true.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	// The mobile number.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// Specifies whether the mobile number is verified. This field is required if a mobile number is specified. If you have no special requirement, set this parameter to true.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +. This parameter is required if a mobile number is specified.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	// The name of the account.
	//
	// example:
	//
	// name001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s PatchUserRequest) String() string {
	return dara.Prettify(s)
}

func (s PatchUserRequest) GoString() string {
	return s.String()
}

func (s *PatchUserRequest) GetCustomFields() []*PatchUserRequestCustomFields {
	return s.CustomFields
}

func (s *PatchUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *PatchUserRequest) GetEmail() *string {
	return s.Email
}

func (s *PatchUserRequest) GetEmailVerified() *bool {
	return s.EmailVerified
}

func (s *PatchUserRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *PatchUserRequest) GetPhoneNumberVerified() *bool {
	return s.PhoneNumberVerified
}

func (s *PatchUserRequest) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *PatchUserRequest) GetUsername() *string {
	return s.Username
}

func (s *PatchUserRequest) SetCustomFields(v []*PatchUserRequestCustomFields) *PatchUserRequest {
	s.CustomFields = v
	return s
}

func (s *PatchUserRequest) SetDisplayName(v string) *PatchUserRequest {
	s.DisplayName = &v
	return s
}

func (s *PatchUserRequest) SetEmail(v string) *PatchUserRequest {
	s.Email = &v
	return s
}

func (s *PatchUserRequest) SetEmailVerified(v bool) *PatchUserRequest {
	s.EmailVerified = &v
	return s
}

func (s *PatchUserRequest) SetPhoneNumber(v string) *PatchUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *PatchUserRequest) SetPhoneNumberVerified(v bool) *PatchUserRequest {
	s.PhoneNumberVerified = &v
	return s
}

func (s *PatchUserRequest) SetPhoneRegion(v string) *PatchUserRequest {
	s.PhoneRegion = &v
	return s
}

func (s *PatchUserRequest) SetUsername(v string) *PatchUserRequest {
	s.Username = &v
	return s
}

func (s *PatchUserRequest) Validate() error {
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

type PatchUserRequestCustomFields struct {
	// The name of the extended field. For more information about the type and valid values of the extended field, see the detailed description of the extended field in the IDaaS console.
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// The value of the extended field.
	//
	// example:
	//
	// test_value
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// The operation to be performed on the field. Valid values:
	//
	// 	- add
	//
	// 	- replace If you leave the value of the extended field empty, the replace operation is converted to an add operation.
	//
	// 	- remove
	//
	// example:
	//
	// replace
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// Deprecated
	//
	// The type of the operation. This parameter is deprecated. Replace it with operation.
	//
	// example:
	//
	// replace
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
}

func (s PatchUserRequestCustomFields) String() string {
	return dara.Prettify(s)
}

func (s PatchUserRequestCustomFields) GoString() string {
	return s.String()
}

func (s *PatchUserRequestCustomFields) GetFieldName() *string {
	return s.FieldName
}

func (s *PatchUserRequestCustomFields) GetFieldValue() *string {
	return s.FieldValue
}

func (s *PatchUserRequestCustomFields) GetOperation() *string {
	return s.Operation
}

func (s *PatchUserRequestCustomFields) GetOperator() *string {
	return s.Operator
}

func (s *PatchUserRequestCustomFields) SetFieldName(v string) *PatchUserRequestCustomFields {
	s.FieldName = &v
	return s
}

func (s *PatchUserRequestCustomFields) SetFieldValue(v string) *PatchUserRequestCustomFields {
	s.FieldValue = &v
	return s
}

func (s *PatchUserRequestCustomFields) SetOperation(v string) *PatchUserRequestCustomFields {
	s.Operation = &v
	return s
}

func (s *PatchUserRequestCustomFields) SetOperator(v string) *PatchUserRequestCustomFields {
	s.Operator = &v
	return s
}

func (s *PatchUserRequestCustomFields) Validate() error {
	return dara.Validate(s)
}

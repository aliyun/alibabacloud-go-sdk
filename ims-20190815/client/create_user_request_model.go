// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComments(v string) *CreateUserRequest
	GetComments() *string
	SetDisplayName(v string) *CreateUserRequest
	GetDisplayName() *string
	SetEmail(v string) *CreateUserRequest
	GetEmail() *string
	SetMobilePhone(v string) *CreateUserRequest
	GetMobilePhone() *string
	SetTag(v []*CreateUserRequestTag) *CreateUserRequest
	GetTag() []*CreateUserRequestTag
	SetUserPrincipalName(v string) *CreateUserRequest
	GetUserPrincipalName() *string
}

type CreateUserRequest struct {
	// The description.
	//
	// The description must be 1 to 128 characters in length.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The display name of the RAM user.
	//
	// The name must be 1 to 24 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile number of the RAM user.
	//
	// Format: Country code-Mobile phone number.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*CreateUserRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The logon name of the Resource Access Management (RAM) user.
	//
	// The name is in the format of `<username>@<AccountAlias>.onaliyun.com`. `<username>` indicates the name of the RAM user. `<AccountAlias>.onaliyun.com` indicates the default domain name. For more information about how to query the default domain name, see [GetDefaultDomain](https://help.aliyun.com/document_detail/186720.html).
	//
	// The value of `UserPrincipalName` must be `1 to 128` characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The value of `<username>` must be `1 to 64` characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetComments() *string {
	return s.Comments
}

func (s *CreateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateUserRequest) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *CreateUserRequest) GetTag() []*CreateUserRequestTag {
	return s.Tag
}

func (s *CreateUserRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *CreateUserRequest) SetComments(v string) *CreateUserRequest {
	s.Comments = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetMobilePhone(v string) *CreateUserRequest {
	s.MobilePhone = &v
	return s
}

func (s *CreateUserRequest) SetTag(v []*CreateUserRequestTag) *CreateUserRequest {
	s.Tag = v
	return s
}

func (s *CreateUserRequest) SetUserPrincipalName(v string) *CreateUserRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	return dara.Validate(s)
}

type CreateUserRequestTag struct {
	// The key of the tag.
	//
	// The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// operator
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// The tag value can be an empty string. The tag value can be up to 128 characters in length. The tag value cannot start with `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// alice
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateUserRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequestTag) GoString() string {
	return s.String()
}

func (s *CreateUserRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateUserRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateUserRequestTag) SetKey(v string) *CreateUserRequestTag {
	s.Key = &v
	return s
}

func (s *CreateUserRequestTag) SetValue(v string) *CreateUserRequestTag {
	s.Value = &v
	return s
}

func (s *CreateUserRequestTag) Validate() error {
	return dara.Validate(s)
}

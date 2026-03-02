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
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// This parameter is required.
	DisplayName *string                 `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string                 `json:"Email,omitempty" xml:"Email,omitempty"`
	MobilePhone *string                 `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	Tag         []*CreateUserRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateUserRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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

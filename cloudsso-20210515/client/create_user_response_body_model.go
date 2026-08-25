// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserResponseBody
	GetRequestId() *string
	SetUser(v *CreateUserResponseBodyUser) *CreateUserResponseBody
	GetUser() *CreateUserResponseBodyUser
}

type CreateUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6F90F3D-4502-5877-B80B-97476F6AE2CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user information.
	User *CreateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s CreateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserResponseBody) GetUser() *CreateUserResponseBodyUser {
	return s.User
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetUser(v *CreateUserResponseBodyUser) *CreateUserResponseBody {
	s.User = v
	return s
}

func (s *CreateUserResponseBody) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserResponseBodyUser struct {
	// The time when the user was created (UTC).
	//
	// example:
	//
	// 2021-10-26T03:03:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the user.
	//
	// example:
	//
	// This is a user.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// Alice
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// Alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The first name of the user.
	//
	// example:
	//
	// Alice
	FirstName *string `json:"FirstName,omitempty" xml:"FirstName,omitempty"`
	// The last name of the user.
	//
	// example:
	//
	// Lee
	LastName *string `json:"LastName,omitempty" xml:"LastName,omitempty"`
	// The type of the user. Valid values:
	//
	// - Manual: Manually created.
	//
	// - Synchronized: Synchronized from an external source.
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. Valid values:
	//
	// - Enabled: Enabled.
	//
	// - Disabled: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*CreateUserResponseBodyUserTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the user was last modified (UTC).
	//
	// example:
	//
	// 2021-10-26T03:03:42Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// Alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUser) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateUserResponseBodyUser) GetDescription() *string {
	return s.Description
}

func (s *CreateUserResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *CreateUserResponseBodyUser) GetFirstName() *string {
	return s.FirstName
}

func (s *CreateUserResponseBodyUser) GetLastName() *string {
	return s.LastName
}

func (s *CreateUserResponseBodyUser) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *CreateUserResponseBodyUser) GetStatus() *string {
	return s.Status
}

func (s *CreateUserResponseBodyUser) GetTags() []*CreateUserResponseBodyUserTags {
	return s.Tags
}

func (s *CreateUserResponseBodyUser) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *CreateUserResponseBodyUser) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserResponseBodyUser) SetCreateTime(v string) *CreateUserResponseBodyUser {
	s.CreateTime = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetDescription(v string) *CreateUserResponseBodyUser {
	s.Description = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetDisplayName(v string) *CreateUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetEmail(v string) *CreateUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetFirstName(v string) *CreateUserResponseBodyUser {
	s.FirstName = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetLastName(v string) *CreateUserResponseBodyUser {
	s.LastName = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetProvisionType(v string) *CreateUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetStatus(v string) *CreateUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetTags(v []*CreateUserResponseBodyUserTags) *CreateUserResponseBodyUser {
	s.Tags = v
	return s
}

func (s *CreateUserResponseBodyUser) SetUpdateTime(v string) *CreateUserResponseBodyUser {
	s.UpdateTime = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetUserId(v string) *CreateUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetUserName(v string) *CreateUserResponseBodyUser {
	s.UserName = &v
	return s
}

func (s *CreateUserResponseBodyUser) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateUserResponseBodyUserTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateUserResponseBodyUserTags) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBodyUserTags) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUserTags) GetKey() *string {
	return s.Key
}

func (s *CreateUserResponseBodyUserTags) GetValue() *string {
	return s.Value
}

func (s *CreateUserResponseBodyUserTags) SetKey(v string) *CreateUserResponseBodyUserTags {
	s.Key = &v
	return s
}

func (s *CreateUserResponseBodyUserTags) SetValue(v string) *CreateUserResponseBodyUserTags {
	s.Value = &v
	return s
}

func (s *CreateUserResponseBodyUserTags) Validate() error {
	return dara.Validate(s)
}

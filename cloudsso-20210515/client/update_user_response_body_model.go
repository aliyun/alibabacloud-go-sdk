// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserResponseBody
	GetRequestId() *string
	SetUser(v *UpdateUserResponseBodyUser) *UpdateUserResponseBody
	GetUser() *UpdateUserResponseBodyUser
}

type UpdateUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F44F02EC-70D1-5E51-8E8E-FA9AC4EF952A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the user.
	User *UpdateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s UpdateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserResponseBody) GetUser() *UpdateUserResponseBodyUser {
	return s.User
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetUser(v *UpdateUserResponseBodyUser) *UpdateUserResponseBody {
	s.User = v
	return s
}

func (s *UpdateUserResponseBody) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserResponseBodyUser struct {
	// The time when the user was created. The value is displayed in UTC.
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
	// AliceLee@example.com
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
	// - Manual: The user is manually created.
	//
	// - Synchronized: The user is synchronized from an external identity provider (IdP).
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. Valid values:
	//
	// - Enabled: The logon of the user is enabled.
	//
	// - Disabled: The logon of the user is disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the information about the user was modified. The value is displayed in UTC.
	//
	// example:
	//
	// 2021-10-26T07:32:32Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the user.
	//
	// example:
	//
	// Alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBodyUser) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateUserResponseBodyUser) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *UpdateUserResponseBodyUser) GetFirstName() *string {
	return s.FirstName
}

func (s *UpdateUserResponseBodyUser) GetLastName() *string {
	return s.LastName
}

func (s *UpdateUserResponseBodyUser) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *UpdateUserResponseBodyUser) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserResponseBodyUser) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserResponseBodyUser) GetUserName() *string {
	return s.UserName
}

func (s *UpdateUserResponseBodyUser) SetCreateTime(v string) *UpdateUserResponseBodyUser {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetDescription(v string) *UpdateUserResponseBodyUser {
	s.Description = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetDisplayName(v string) *UpdateUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetEmail(v string) *UpdateUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetFirstName(v string) *UpdateUserResponseBodyUser {
	s.FirstName = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetLastName(v string) *UpdateUserResponseBodyUser {
	s.LastName = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetProvisionType(v string) *UpdateUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetStatus(v string) *UpdateUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUpdateTime(v string) *UpdateUserResponseBodyUser {
	s.UpdateTime = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUserId(v string) *UpdateUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUserName(v string) *UpdateUserResponseBodyUser {
	s.UserName = &v
	return s
}

func (s *UpdateUserResponseBodyUser) Validate() error {
	return dara.Validate(s)
}

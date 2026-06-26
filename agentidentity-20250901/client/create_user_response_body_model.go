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
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	User      *CreateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBodyUser) GoString() string {
	return s.String()
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

func (s *CreateUserResponseBodyUser) GetStatus() *string {
	return s.Status
}

func (s *CreateUserResponseBodyUser) GetType() *string {
	return s.Type
}

func (s *CreateUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *CreateUserResponseBodyUser) GetUserName() *string {
	return s.UserName
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

func (s *CreateUserResponseBodyUser) SetStatus(v string) *CreateUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetType(v string) *CreateUserResponseBodyUser {
	s.Type = &v
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
	return dara.Validate(s)
}

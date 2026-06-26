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
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	User      *UpdateUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
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
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime  *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *UpdateUserResponseBodyUser) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserResponseBodyUser) GetType() *string {
	return s.Type
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

func (s *UpdateUserResponseBodyUser) SetStatus(v string) *UpdateUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetType(v string) *UpdateUserResponseBodyUser {
	s.Type = &v
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

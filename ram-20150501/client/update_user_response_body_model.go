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
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
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
	// The description of the RAM user.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The point in time when the RAM user was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// xiaoq****
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// >  This parameter can be returned only on the China site (aliyun.com).
	//
	// example:
	//
	// xiaoq****@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// >  This parameter can be returned only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1860000****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The point in time when the information about the RAM user was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-02-11T03:15:21Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 122748924538****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the RAM user.
	//
	// example:
	//
	// xiaoq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBodyUser) GetComments() *string {
	return s.Comments
}

func (s *UpdateUserResponseBodyUser) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdateUserResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *UpdateUserResponseBodyUser) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *UpdateUserResponseBodyUser) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserResponseBodyUser) GetUserName() *string {
	return s.UserName
}

func (s *UpdateUserResponseBodyUser) SetComments(v string) *UpdateUserResponseBodyUser {
	s.Comments = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetCreateDate(v string) *UpdateUserResponseBodyUser {
	s.CreateDate = &v
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

func (s *UpdateUserResponseBodyUser) SetMobilePhone(v string) *UpdateUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetUpdateDate(v string) *UpdateUserResponseBodyUser {
	s.UpdateDate = &v
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
	SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody
	GetUser() *GetUserResponseBodyUser
}

type GetUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF5189484043
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
	User *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) GetUser() *GetUserResponseBodyUser {
	return s.User
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserResponseBodyUser struct {
	// The description of the RAM user.
	//
	// example:
	//
	// Cloud computing engineer
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
	// alice
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// >  This parameter can be returned only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The point in time when the RAM user last logged on to the Alibaba Cloud Management Console by using the password. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	LastLoginDate *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
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
	// 222748924538****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the RAM user.
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) GetComments() *string {
	return s.Comments
}

func (s *GetUserResponseBodyUser) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetUserResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBodyUser) GetLastLoginDate() *string {
	return s.LastLoginDate
}

func (s *GetUserResponseBodyUser) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *GetUserResponseBodyUser) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *GetUserResponseBodyUser) GetUserName() *string {
	return s.UserName
}

func (s *GetUserResponseBodyUser) SetComments(v string) *GetUserResponseBodyUser {
	s.Comments = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCreateDate(v string) *GetUserResponseBodyUser {
	s.CreateDate = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDisplayName(v string) *GetUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmail(v string) *GetUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLastLoginDate(v string) *GetUserResponseBodyUser {
	s.LastLoginDate = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMobilePhone(v string) *GetUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUpdateDate(v string) *GetUserResponseBodyUser {
	s.UpdateDate = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserName(v string) *GetUserResponseBodyUser {
	s.UserName = &v
	return s
}

func (s *GetUserResponseBodyUser) Validate() error {
	return dara.Validate(s)
}

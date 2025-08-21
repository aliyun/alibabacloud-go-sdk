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
	// 1B56DD42-6962-4F89-A19C-079EED1F0FE3
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
	return dara.Validate(s)
}

type UpdateUserResponseBodyUser struct {
	// The description.
	//
	// example:
	//
	// This is a cloud computing engineer.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the RAM user was created.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the RAM user.
	//
	// example:
	//
	// new
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The last time when the RAM user logged on to the Alibaba Cloud Management Console.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
	LastLoginDate *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
	// The mobile phone number of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	MobilePhone *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	// The source of the RAM user. Valid values:
	//
	// 	- Manual: The RAM user is manually created in the RAM console.
	//
	// 	- SCIM: The RAM user is mapped by using System for Cross-domain Identity Management (SCIM).
	//
	// 	- CloudSSO: The RAM user is mapped from a CloudSSO user.
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The time when the information about the RAM user was updated.
	//
	// example:
	//
	// 2020-10-13T09:19:49Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The ID of the RAM user.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// example:
	//
	// new@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
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

func (s *UpdateUserResponseBodyUser) GetLastLoginDate() *string {
	return s.LastLoginDate
}

func (s *UpdateUserResponseBodyUser) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *UpdateUserResponseBodyUser) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *UpdateUserResponseBodyUser) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserResponseBodyUser) GetUserPrincipalName() *string {
	return s.UserPrincipalName
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

func (s *UpdateUserResponseBodyUser) SetLastLoginDate(v string) *UpdateUserResponseBodyUser {
	s.LastLoginDate = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetMobilePhone(v string) *UpdateUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *UpdateUserResponseBodyUser) SetProvisionType(v string) *UpdateUserResponseBodyUser {
	s.ProvisionType = &v
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

func (s *UpdateUserResponseBodyUser) SetUserPrincipalName(v string) *UpdateUserResponseBodyUser {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateUserResponseBodyUser) Validate() error {
	return dara.Validate(s)
}

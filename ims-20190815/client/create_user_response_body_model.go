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
	// 2BB8C44A-2862-4922-AD43-03924749173B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user.
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
	return dara.Validate(s)
}

type CreateUserResponseBodyUser struct {
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
	// The tag value.
	Tags *CreateUserResponseBodyUserTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The time when the information about the RAM user was updated.
	//
	// example:
	//
	// 2020-10-12T09:12:00Z
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
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUser) GetComments() *string {
	return s.Comments
}

func (s *CreateUserResponseBodyUser) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateUserResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *CreateUserResponseBodyUser) GetLastLoginDate() *string {
	return s.LastLoginDate
}

func (s *CreateUserResponseBodyUser) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *CreateUserResponseBodyUser) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *CreateUserResponseBodyUser) GetTags() *CreateUserResponseBodyUserTags {
	return s.Tags
}

func (s *CreateUserResponseBodyUser) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *CreateUserResponseBodyUser) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *CreateUserResponseBodyUser) SetComments(v string) *CreateUserResponseBodyUser {
	s.Comments = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetCreateDate(v string) *CreateUserResponseBodyUser {
	s.CreateDate = &v
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

func (s *CreateUserResponseBodyUser) SetLastLoginDate(v string) *CreateUserResponseBodyUser {
	s.LastLoginDate = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetMobilePhone(v string) *CreateUserResponseBodyUser {
	s.MobilePhone = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetProvisionType(v string) *CreateUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetTags(v *CreateUserResponseBodyUserTags) *CreateUserResponseBodyUser {
	s.Tags = v
	return s
}

func (s *CreateUserResponseBodyUser) SetUpdateDate(v string) *CreateUserResponseBodyUser {
	s.UpdateDate = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetUserId(v string) *CreateUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseBodyUser) SetUserPrincipalName(v string) *CreateUserResponseBodyUser {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateUserResponseBodyUser) Validate() error {
	return dara.Validate(s)
}

type CreateUserResponseBodyUserTags struct {
	Tag []*CreateUserResponseBodyUserTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateUserResponseBodyUserTags) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBodyUserTags) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUserTags) GetTag() []*CreateUserResponseBodyUserTagsTag {
	return s.Tag
}

func (s *CreateUserResponseBodyUserTags) SetTag(v []*CreateUserResponseBodyUserTagsTag) *CreateUserResponseBodyUserTags {
	s.Tag = v
	return s
}

func (s *CreateUserResponseBodyUserTags) Validate() error {
	return dara.Validate(s)
}

type CreateUserResponseBodyUserTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// operator
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// alice
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateUserResponseBodyUserTagsTag) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBodyUserTagsTag) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyUserTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateUserResponseBodyUserTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateUserResponseBodyUserTagsTag) SetTagKey(v string) *CreateUserResponseBodyUserTagsTag {
	s.TagKey = &v
	return s
}

func (s *CreateUserResponseBodyUserTagsTag) SetTagValue(v string) *CreateUserResponseBodyUserTagsTag {
	s.TagValue = &v
	return s
}

func (s *CreateUserResponseBodyUserTagsTag) Validate() error {
	return dara.Validate(s)
}

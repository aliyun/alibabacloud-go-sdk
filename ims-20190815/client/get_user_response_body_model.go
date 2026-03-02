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
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	User      *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
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
	Comments          *string                      `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate        *string                      `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName       *string                      `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email             *string                      `json:"Email,omitempty" xml:"Email,omitempty"`
	LastLoginDate     *string                      `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
	MobilePhone       *string                      `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	ProvisionType     *string                      `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	Tags              *GetUserResponseBodyUserTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UpdateDate        *string                      `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	UserId            *string                      `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName          *string                      `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPrincipalName *string                      `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
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

func (s *GetUserResponseBodyUser) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *GetUserResponseBodyUser) GetTags() *GetUserResponseBodyUserTags {
	return s.Tags
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

func (s *GetUserResponseBodyUser) GetUserPrincipalName() *string {
	return s.UserPrincipalName
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

func (s *GetUserResponseBodyUser) SetProvisionType(v string) *GetUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *GetUserResponseBodyUser) SetTags(v *GetUserResponseBodyUserTags) *GetUserResponseBodyUser {
	s.Tags = v
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

func (s *GetUserResponseBodyUser) SetUserPrincipalName(v string) *GetUserResponseBodyUser {
	s.UserPrincipalName = &v
	return s
}

func (s *GetUserResponseBodyUser) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserResponseBodyUserTags struct {
	Tag []*GetUserResponseBodyUserTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUserTags) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUserTags) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserTags) GetTag() []*GetUserResponseBodyUserTagsTag {
	return s.Tag
}

func (s *GetUserResponseBodyUserTags) SetTag(v []*GetUserResponseBodyUserTagsTag) *GetUserResponseBodyUserTags {
	s.Tag = v
	return s
}

func (s *GetUserResponseBodyUserTags) Validate() error {
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

type GetUserResponseBodyUserTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetUserResponseBodyUserTagsTag) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUserTagsTag) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *GetUserResponseBodyUserTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *GetUserResponseBodyUserTagsTag) SetTagKey(v string) *GetUserResponseBodyUserTagsTag {
	s.TagKey = &v
	return s
}

func (s *GetUserResponseBodyUserTagsTag) SetTagValue(v string) *GetUserResponseBodyUserTagsTag {
	s.TagValue = &v
	return s
}

func (s *GetUserResponseBodyUserTagsTag) Validate() error {
	return dara.Validate(s)
}

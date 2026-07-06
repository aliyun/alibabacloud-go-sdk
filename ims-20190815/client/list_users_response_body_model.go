// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListUsersResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListUsersResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody
	GetUsers() *ListUsersResponseBodyUsers
}

type ListUsersResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The parameter that is used to obtain the truncated part. It takes effect only when `IsTruncated` is set to `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B450CA1-36E8-4AA2-8461-86B42BF4CC4E
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     *ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListUsersResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetUsers() *ListUsersResponseBodyUsers {
	return s.Users
}

func (s *ListUsersResponseBody) SetIsTruncated(v bool) *ListUsersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersResponseBody) SetMarker(v string) *ListUsersResponseBody {
	s.Marker = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v *ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	if s.Users != nil {
		if err := s.Users.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUsersResponseBodyUsers struct {
	User []*ListUsersResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) GetUser() []*ListUsersResponseBodyUsersUser {
	return s.User
}

func (s *ListUsersResponseBodyUsers) SetUser(v []*ListUsersResponseBodyUsersUser) *ListUsersResponseBodyUsers {
	s.User = v
	return s
}

func (s *ListUsersResponseBodyUsers) Validate() error {
	if s.User != nil {
		for _, item := range s.User {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersResponseBodyUsersUser struct {
	Comments          *string                             `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateDate        *string                             `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DisplayName       *string                             `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email             *string                             `json:"Email,omitempty" xml:"Email,omitempty"`
	LastLoginDate     *string                             `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty"`
	MobilePhone       *string                             `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	ProvisionType     *string                             `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	Status            *string                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags              *ListUsersResponseBodyUsersUserTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UpdateDate        *string                             `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	UserId            *string                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserPrincipalName *string                             `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListUsersResponseBodyUsersUser) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUser) GetComments() *string {
	return s.Comments
}

func (s *ListUsersResponseBodyUsersUser) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListUsersResponseBodyUsersUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyUsersUser) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyUsersUser) GetLastLoginDate() *string {
	return s.LastLoginDate
}

func (s *ListUsersResponseBodyUsersUser) GetMobilePhone() *string {
	return s.MobilePhone
}

func (s *ListUsersResponseBodyUsersUser) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *ListUsersResponseBodyUsersUser) GetStatus() *string {
	return s.Status
}

func (s *ListUsersResponseBodyUsersUser) GetTags() *ListUsersResponseBodyUsersUserTags {
	return s.Tags
}

func (s *ListUsersResponseBodyUsersUser) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListUsersResponseBodyUsersUser) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersResponseBodyUsersUser) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *ListUsersResponseBodyUsersUser) SetComments(v string) *ListUsersResponseBodyUsersUser {
	s.Comments = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetCreateDate(v string) *ListUsersResponseBodyUsersUser {
	s.CreateDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetDisplayName(v string) *ListUsersResponseBodyUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetEmail(v string) *ListUsersResponseBodyUsersUser {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetLastLoginDate(v string) *ListUsersResponseBodyUsersUser {
	s.LastLoginDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetMobilePhone(v string) *ListUsersResponseBodyUsersUser {
	s.MobilePhone = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetProvisionType(v string) *ListUsersResponseBodyUsersUser {
	s.ProvisionType = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetStatus(v string) *ListUsersResponseBodyUsersUser {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetTags(v *ListUsersResponseBodyUsersUserTags) *ListUsersResponseBodyUsersUser {
	s.Tags = v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUpdateDate(v string) *ListUsersResponseBodyUsersUser {
	s.UpdateDate = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUserId(v string) *ListUsersResponseBodyUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) SetUserPrincipalName(v string) *ListUsersResponseBodyUsersUser {
	s.UserPrincipalName = &v
	return s
}

func (s *ListUsersResponseBodyUsersUser) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUsersResponseBodyUsersUserTags struct {
	Tag []*ListUsersResponseBodyUsersUserTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsersUserTags) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsersUserTags) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUserTags) GetTag() []*ListUsersResponseBodyUsersUserTagsTag {
	return s.Tag
}

func (s *ListUsersResponseBodyUsersUserTags) SetTag(v []*ListUsersResponseBodyUsersUserTagsTag) *ListUsersResponseBodyUsersUserTags {
	s.Tag = v
	return s
}

func (s *ListUsersResponseBodyUsersUserTags) Validate() error {
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

type ListUsersResponseBodyUsersUserTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListUsersResponseBodyUsersUserTagsTag) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsersUserTagsTag) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersUserTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListUsersResponseBodyUsersUserTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListUsersResponseBodyUsersUserTagsTag) SetTagKey(v string) *ListUsersResponseBodyUsersUserTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserTagsTag) SetTagValue(v string) *ListUsersResponseBodyUsersUserTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListUsersResponseBodyUsersUserTagsTag) Validate() error {
	return dara.Validate(s)
}

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
	// EE42D2C4-A30A-59B7-ACEB-6D22FB44174A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user information.
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
	// The time when the user was created (in UTC).
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
	// The user identifier information of the external identity provider.
	ExternalId *GetUserResponseBodyUserExternalId `json:"ExternalId,omitempty" xml:"ExternalId,omitempty" type:"Struct"`
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
	// - Synchronized: Synchronized from an external identity provider.
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
	Tags []*GetUserResponseBodyUserTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the user was last modified (in UTC).
	//
	// example:
	//
	// 2021-10-26T06:43:55Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
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

func (s GetUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserResponseBodyUser) GetDescription() *string {
	return s.Description
}

func (s *GetUserResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserResponseBodyUser) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBodyUser) GetExternalId() *GetUserResponseBodyUserExternalId {
	return s.ExternalId
}

func (s *GetUserResponseBodyUser) GetFirstName() *string {
	return s.FirstName
}

func (s *GetUserResponseBodyUser) GetLastName() *string {
	return s.LastName
}

func (s *GetUserResponseBodyUser) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *GetUserResponseBodyUser) GetStatus() *string {
	return s.Status
}

func (s *GetUserResponseBodyUser) GetTags() []*GetUserResponseBodyUserTags {
	return s.Tags
}

func (s *GetUserResponseBodyUser) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUserResponseBodyUser) GetUserId() *string {
	return s.UserId
}

func (s *GetUserResponseBodyUser) GetUserName() *string {
	return s.UserName
}

func (s *GetUserResponseBodyUser) SetCreateTime(v string) *GetUserResponseBodyUser {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDescription(v string) *GetUserResponseBodyUser {
	s.Description = &v
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

func (s *GetUserResponseBodyUser) SetExternalId(v *GetUserResponseBodyUserExternalId) *GetUserResponseBodyUser {
	s.ExternalId = v
	return s
}

func (s *GetUserResponseBodyUser) SetFirstName(v string) *GetUserResponseBodyUser {
	s.FirstName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLastName(v string) *GetUserResponseBodyUser {
	s.LastName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetProvisionType(v string) *GetUserResponseBodyUser {
	s.ProvisionType = &v
	return s
}

func (s *GetUserResponseBodyUser) SetStatus(v string) *GetUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *GetUserResponseBodyUser) SetTags(v []*GetUserResponseBodyUserTags) *GetUserResponseBodyUser {
	s.Tags = v
	return s
}

func (s *GetUserResponseBodyUser) SetUpdateTime(v string) *GetUserResponseBodyUser {
	s.UpdateTime = &v
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
	if s.ExternalId != nil {
		if err := s.ExternalId.Validate(); err != nil {
			return err
		}
	}
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

type GetUserResponseBodyUserExternalId struct {
	// The user identifier of the external identity provider.
	//
	// example:
	//
	// c73******a5fdd5
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The external identity synchronization channel. Currently, only SCIM synchronization is supported.
	//
	// example:
	//
	// SCIM
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
}

func (s GetUserResponseBodyUserExternalId) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUserExternalId) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserExternalId) GetId() *string {
	return s.Id
}

func (s *GetUserResponseBodyUserExternalId) GetIssuer() *string {
	return s.Issuer
}

func (s *GetUserResponseBodyUserExternalId) SetId(v string) *GetUserResponseBodyUserExternalId {
	s.Id = &v
	return s
}

func (s *GetUserResponseBodyUserExternalId) SetIssuer(v string) *GetUserResponseBodyUserExternalId {
	s.Issuer = &v
	return s
}

func (s *GetUserResponseBodyUserExternalId) Validate() error {
	return dara.Validate(s)
}

type GetUserResponseBodyUserTags struct {
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

func (s GetUserResponseBodyUserTags) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyUserTags) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserTags) GetKey() *string {
	return s.Key
}

func (s *GetUserResponseBodyUserTags) GetValue() *string {
	return s.Value
}

func (s *GetUserResponseBodyUserTags) SetKey(v string) *GetUserResponseBodyUserTags {
	s.Key = &v
	return s
}

func (s *GetUserResponseBodyUserTags) SetValue(v string) *GetUserResponseBodyUserTags {
	s.Value = &v
	return s
}

func (s *GetUserResponseBodyUserTags) Validate() error {
	return dara.Validate(s)
}

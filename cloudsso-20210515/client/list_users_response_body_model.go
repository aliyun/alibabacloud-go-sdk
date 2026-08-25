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
	SetMaxResults(v int32) *ListUsersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUsersResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListUsersResponseBody
	GetTotalCounts() *int32
	SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody
	GetUsers() []*ListUsersResponseBodyUsers
}

type ListUsersResponseBody struct {
	// Indicates whether the results are truncated. Valid values:
	//
	// - true: The results are truncated.
	//
	// - false: The results are not truncated.
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results.
	//
	// > This parameter is returned only when `IsTruncated` is `true`.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 734D9AAC-9A8E-5DF6-A633-ADE70FF2A9B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that match the request parameters.
	//
	// example:
	//
	// 2
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
	// The user list.
	Users []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
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

func (s *ListUsersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUsersResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListUsersResponseBody) GetUsers() []*ListUsersResponseBodyUsers {
	return s.Users
}

func (s *ListUsersResponseBody) SetIsTruncated(v bool) *ListUsersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersResponseBody) SetMaxResults(v int32) *ListUsersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUsersResponseBody) SetNextToken(v string) *ListUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCounts(v int32) *ListUsersResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

func (s *ListUsersResponseBody) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersResponseBodyUsers struct {
	// The time when the user was created (UTC).
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
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
	// AliceLee
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// AliceLee@example.onmicrosoft.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The identifier information of the user from the external identity provider.
	ExternalId *ListUsersResponseBodyUsersExternalId `json:"ExternalId,omitempty" xml:"ExternalId,omitempty" type:"Struct"`
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
	// - Synchronized: The user is synchronized from an external identity provider.
	//
	// example:
	//
	// Synchronized
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The status of the user. Valid values:
	//
	// - Enabled: The user is enabled.
	//
	// - Disabled: The user is disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*ListUsersResponseBodyUsersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the user was last modified (UTC).
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// u-00bikzkuzbb58luh****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// AliceLee@example.onmicrosoft.com
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUsersResponseBodyUsers) GetDescription() *string {
	return s.Description
}

func (s *ListUsersResponseBodyUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUsersResponseBodyUsers) GetEmail() *string {
	return s.Email
}

func (s *ListUsersResponseBodyUsers) GetExternalId() *ListUsersResponseBodyUsersExternalId {
	return s.ExternalId
}

func (s *ListUsersResponseBodyUsers) GetFirstName() *string {
	return s.FirstName
}

func (s *ListUsersResponseBodyUsers) GetLastName() *string {
	return s.LastName
}

func (s *ListUsersResponseBodyUsers) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *ListUsersResponseBodyUsers) GetStatus() *string {
	return s.Status
}

func (s *ListUsersResponseBodyUsers) GetTags() []*ListUsersResponseBodyUsersTags {
	return s.Tags
}

func (s *ListUsersResponseBodyUsers) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListUsersResponseBodyUsers) GetUserId() *string {
	return s.UserId
}

func (s *ListUsersResponseBodyUsers) GetUserName() *string {
	return s.UserName
}

func (s *ListUsersResponseBodyUsers) SetCreateTime(v string) *ListUsersResponseBodyUsers {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDescription(v string) *ListUsersResponseBodyUsers {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDisplayName(v string) *ListUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetExternalId(v *ListUsersResponseBodyUsersExternalId) *ListUsersResponseBodyUsers {
	s.ExternalId = v
	return s
}

func (s *ListUsersResponseBodyUsers) SetFirstName(v string) *ListUsersResponseBodyUsers {
	s.FirstName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetLastName(v string) *ListUsersResponseBodyUsers {
	s.LastName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetProvisionType(v string) *ListUsersResponseBodyUsers {
	s.ProvisionType = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetStatus(v string) *ListUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetTags(v []*ListUsersResponseBodyUsersTags) *ListUsersResponseBodyUsers {
	s.Tags = v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUpdateTime(v string) *ListUsersResponseBodyUsers {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserId(v string) *ListUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserName(v string) *ListUsersResponseBodyUsers {
	s.UserName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) Validate() error {
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

type ListUsersResponseBodyUsersExternalId struct {
	// The user identifier from the external identity provider.
	//
	// example:
	//
	// c73******a5fdd5
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The external identity synchronization channel. Currently, only SCIM-based user synchronization is supported.
	//
	// example:
	//
	// SCIM
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
}

func (s ListUsersResponseBodyUsersExternalId) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsersExternalId) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersExternalId) GetId() *string {
	return s.Id
}

func (s *ListUsersResponseBodyUsersExternalId) GetIssuer() *string {
	return s.Issuer
}

func (s *ListUsersResponseBodyUsersExternalId) SetId(v string) *ListUsersResponseBodyUsersExternalId {
	s.Id = &v
	return s
}

func (s *ListUsersResponseBodyUsersExternalId) SetIssuer(v string) *ListUsersResponseBodyUsersExternalId {
	s.Issuer = &v
	return s
}

func (s *ListUsersResponseBodyUsersExternalId) Validate() error {
	return dara.Validate(s)
}

type ListUsersResponseBodyUsersTags struct {
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

func (s ListUsersResponseBodyUsersTags) String() string {
	return dara.Prettify(s)
}

func (s ListUsersResponseBodyUsersTags) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsersTags) GetKey() *string {
	return s.Key
}

func (s *ListUsersResponseBodyUsersTags) GetValue() *string {
	return s.Value
}

func (s *ListUsersResponseBodyUsersTags) SetKey(v string) *ListUsersResponseBodyUsersTags {
	s.Key = &v
	return s
}

func (s *ListUsersResponseBodyUsersTags) SetValue(v string) *ListUsersResponseBodyUsersTags {
	s.Value = &v
	return s
}

func (s *ListUsersResponseBodyUsersTags) Validate() error {
	return dara.Validate(s)
}

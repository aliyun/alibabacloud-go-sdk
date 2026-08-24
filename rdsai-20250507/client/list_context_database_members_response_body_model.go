// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListContextDatabaseMembersResponseBody
	GetMaxResults() *int32
	SetMembers(v []*ListContextDatabaseMembersResponseBodyMembers) *ListContextDatabaseMembersResponseBody
	GetMembers() []*ListContextDatabaseMembersResponseBodyMembers
	SetNextToken(v string) *ListContextDatabaseMembersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListContextDatabaseMembersResponseBody
	GetRequestId() *string
}

type ListContextDatabaseMembersResponseBody struct {
	// The maximum number of entries per page. This field is empty.
	//
	// example:
	//
	// (null)
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The list of members.
	Members []*ListContextDatabaseMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The pagination token for the next page. This field is empty.
	//
	// example:
	//
	// (null)
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListContextDatabaseMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseMembersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextDatabaseMembersResponseBody) GetMembers() []*ListContextDatabaseMembersResponseBodyMembers {
	return s.Members
}

func (s *ListContextDatabaseMembersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextDatabaseMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseMembersResponseBody) SetMaxResults(v int32) *ListContextDatabaseMembersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBody) SetMembers(v []*ListContextDatabaseMembersResponseBodyMembers) *ListContextDatabaseMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListContextDatabaseMembersResponseBody) SetNextToken(v string) *ListContextDatabaseMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBody) SetRequestId(v string) *ListContextDatabaseMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBody) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseMembersResponseBodyMembers struct {
	// The time when the member was created.
	//
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The list of API keys.
	Keys []*ListContextDatabaseMembersResponseBodyMembersKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The member ID.
	//
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The member name.
	//
	// example:
	//
	// Alice
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The member role.
	//
	// example:
	//
	// admin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The member status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListContextDatabaseMembersResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseMembersResponseBodyMembers) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListContextDatabaseMembersResponseBodyMembers) GetKeys() []*ListContextDatabaseMembersResponseBodyMembersKeys {
	return s.Keys
}

func (s *ListContextDatabaseMembersResponseBodyMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *ListContextDatabaseMembersResponseBodyMembers) GetMemberName() *string {
	return s.MemberName
}

func (s *ListContextDatabaseMembersResponseBodyMembers) GetRole() *string {
	return s.Role
}

func (s *ListContextDatabaseMembersResponseBodyMembers) GetStatus() *string {
	return s.Status
}

func (s *ListContextDatabaseMembersResponseBodyMembers) SetCreatedAt(v string) *ListContextDatabaseMembersResponseBodyMembers {
	s.CreatedAt = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembers) SetKeys(v []*ListContextDatabaseMembersResponseBodyMembersKeys) *ListContextDatabaseMembersResponseBodyMembers {
	s.Keys = v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembers) SetMemberId(v string) *ListContextDatabaseMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembers) SetMemberName(v string) *ListContextDatabaseMembersResponseBodyMembers {
	s.MemberName = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembers) SetRole(v string) *ListContextDatabaseMembersResponseBodyMembers {
	s.Role = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembers) SetStatus(v string) *ListContextDatabaseMembersResponseBodyMembers {
	s.Status = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembers) Validate() error {
	if s.Keys != nil {
		for _, item := range s.Keys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContextDatabaseMembersResponseBodyMembersKeys struct {
	// The time when the member was created.
	//
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The API key description.
	//
	// example:
	//
	// data pipeline key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This field is empty.
	//
	// example:
	//
	// (null)
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// The suffix of the API key.
	//
	// example:
	//
	// 33631c
	KeyDisplaySuffix *string `json:"KeyDisplaySuffix,omitempty" xml:"KeyDisplaySuffix,omitempty"`
	// The key ID.
	//
	// example:
	//
	// 1
	KeyId *int64 `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The prefix of the API key.
	//
	// example:
	//
	// ctxdb-
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	// The time when the key was last used. This field is populated after the key has been authenticated and used. This field is empty for keys that have never been used.
	//
	// example:
	//
	// 2026-07-15T08:30:00Z
	LastUsedAt *string `json:"LastUsedAt,omitempty" xml:"LastUsedAt,omitempty"`
	// The API key name.
	//
	// example:
	//
	// my-key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This field is empty.
	//
	// example:
	//
	// (null)
	RevokedAt *string `json:"RevokedAt,omitempty" xml:"RevokedAt,omitempty"`
	// The API key status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListContextDatabaseMembersResponseBodyMembersKeys) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseMembersResponseBodyMembersKeys) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetDescription() *string {
	return s.Description
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetKeyDisplaySuffix() *string {
	return s.KeyDisplaySuffix
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetKeyId() *int64 {
	return s.KeyId
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetLastUsedAt() *string {
	return s.LastUsedAt
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetName() *string {
	return s.Name
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) GetStatus() *string {
	return s.Status
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetCreatedAt(v string) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.CreatedAt = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetDescription(v string) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.Description = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetExpiresAt(v string) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.ExpiresAt = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetKeyDisplaySuffix(v string) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.KeyDisplaySuffix = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetKeyId(v int64) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.KeyId = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetKeyPrefix(v string) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.KeyPrefix = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetLastUsedAt(v string) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.LastUsedAt = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetName(v string) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.Name = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetRevokedAt(v string) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.RevokedAt = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) SetStatus(v string) *ListContextDatabaseMembersResponseBodyMembersKeys {
	s.Status = &v
	return s
}

func (s *ListContextDatabaseMembersResponseBodyMembersKeys) Validate() error {
	return dara.Validate(s)
}

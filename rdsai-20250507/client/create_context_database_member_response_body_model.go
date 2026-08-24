// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateContextDatabaseMemberResponseBody
	GetApiKey() *string
	SetMember(v *CreateContextDatabaseMemberResponseBodyMember) *CreateContextDatabaseMemberResponseBody
	GetMember() *CreateContextDatabaseMemberResponseBodyMember
	SetRequestId(v string) *CreateContextDatabaseMemberResponseBody
	GetRequestId() *string
}

type CreateContextDatabaseMemberResponseBody struct {
	// The plaintext API key. This field is returned only when GenerateInitialKey is set to true. The plaintext is returned only once. Store it securely.
	//
	// example:
	//
	// ctxdb-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The member information.
	Member *CreateContextDatabaseMemberResponseBodyMember `json:"Member,omitempty" xml:"Member,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContextDatabaseMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseMemberResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseMemberResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateContextDatabaseMemberResponseBody) GetMember() *CreateContextDatabaseMemberResponseBodyMember {
	return s.Member
}

func (s *CreateContextDatabaseMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContextDatabaseMemberResponseBody) SetApiKey(v string) *CreateContextDatabaseMemberResponseBody {
	s.ApiKey = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBody) SetMember(v *CreateContextDatabaseMemberResponseBodyMember) *CreateContextDatabaseMemberResponseBody {
	s.Member = v
	return s
}

func (s *CreateContextDatabaseMemberResponseBody) SetRequestId(v string) *CreateContextDatabaseMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBody) Validate() error {
	if s.Member != nil {
		if err := s.Member.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateContextDatabaseMemberResponseBodyMember struct {
	// The time when the member was created.
	//
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The list of API keys.
	Keys []*CreateContextDatabaseMemberResponseBodyMemberKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
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

func (s CreateContextDatabaseMemberResponseBodyMember) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseMemberResponseBodyMember) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseMemberResponseBodyMember) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateContextDatabaseMemberResponseBodyMember) GetKeys() []*CreateContextDatabaseMemberResponseBodyMemberKeys {
	return s.Keys
}

func (s *CreateContextDatabaseMemberResponseBodyMember) GetMemberId() *string {
	return s.MemberId
}

func (s *CreateContextDatabaseMemberResponseBodyMember) GetMemberName() *string {
	return s.MemberName
}

func (s *CreateContextDatabaseMemberResponseBodyMember) GetRole() *string {
	return s.Role
}

func (s *CreateContextDatabaseMemberResponseBodyMember) GetStatus() *string {
	return s.Status
}

func (s *CreateContextDatabaseMemberResponseBodyMember) SetCreatedAt(v string) *CreateContextDatabaseMemberResponseBodyMember {
	s.CreatedAt = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMember) SetKeys(v []*CreateContextDatabaseMemberResponseBodyMemberKeys) *CreateContextDatabaseMemberResponseBodyMember {
	s.Keys = v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMember) SetMemberId(v string) *CreateContextDatabaseMemberResponseBodyMember {
	s.MemberId = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMember) SetMemberName(v string) *CreateContextDatabaseMemberResponseBodyMember {
	s.MemberName = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMember) SetRole(v string) *CreateContextDatabaseMemberResponseBodyMember {
	s.Role = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMember) SetStatus(v string) *CreateContextDatabaseMemberResponseBodyMember {
	s.Status = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMember) Validate() error {
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

type CreateContextDatabaseMemberResponseBodyMemberKeys struct {
	// The time when the member was created.
	//
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The API key description. This field is not used.
	//
	// example:
	//
	// 111
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
	// The key ID, which is generated by the server. This ID is used to locate the API key when it is revoked.
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
	// This field is empty.
	//
	// example:
	//
	// (null)
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

func (s CreateContextDatabaseMemberResponseBodyMemberKeys) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseMemberResponseBodyMemberKeys) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetDescription() *string {
	return s.Description
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetKeyDisplaySuffix() *string {
	return s.KeyDisplaySuffix
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetKeyId() *int64 {
	return s.KeyId
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetLastUsedAt() *string {
	return s.LastUsedAt
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetName() *string {
	return s.Name
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) GetStatus() *string {
	return s.Status
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetCreatedAt(v string) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.CreatedAt = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetDescription(v string) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.Description = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetExpiresAt(v string) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.ExpiresAt = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetKeyDisplaySuffix(v string) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.KeyDisplaySuffix = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetKeyId(v int64) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.KeyId = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetKeyPrefix(v string) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.KeyPrefix = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetLastUsedAt(v string) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.LastUsedAt = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetName(v string) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.Name = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetRevokedAt(v string) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.RevokedAt = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) SetStatus(v string) *CreateContextDatabaseMemberResponseBodyMemberKeys {
	s.Status = &v
	return s
}

func (s *CreateContextDatabaseMemberResponseBodyMemberKeys) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *DeleteContextDatabaseMemberResponseBody
	GetCreatedAt() *string
	SetKeys(v []*DeleteContextDatabaseMemberResponseBodyKeys) *DeleteContextDatabaseMemberResponseBody
	GetKeys() []*DeleteContextDatabaseMemberResponseBodyKeys
	SetMemberId(v string) *DeleteContextDatabaseMemberResponseBody
	GetMemberId() *string
	SetMemberName(v string) *DeleteContextDatabaseMemberResponseBody
	GetMemberName() *string
	SetRequestId(v string) *DeleteContextDatabaseMemberResponseBody
	GetRequestId() *string
	SetRole(v string) *DeleteContextDatabaseMemberResponseBody
	GetRole() *string
	SetStatus(v string) *DeleteContextDatabaseMemberResponseBody
	GetStatus() *string
}

type DeleteContextDatabaseMemberResponseBody struct {
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string                                        `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Keys      []*DeleteContextDatabaseMemberResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// Alice
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// admin
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// deleted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteContextDatabaseMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseMemberResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteContextDatabaseMemberResponseBody) GetKeys() []*DeleteContextDatabaseMemberResponseBodyKeys {
	return s.Keys
}

func (s *DeleteContextDatabaseMemberResponseBody) GetMemberId() *string {
	return s.MemberId
}

func (s *DeleteContextDatabaseMemberResponseBody) GetMemberName() *string {
	return s.MemberName
}

func (s *DeleteContextDatabaseMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContextDatabaseMemberResponseBody) GetRole() *string {
	return s.Role
}

func (s *DeleteContextDatabaseMemberResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteContextDatabaseMemberResponseBody) SetCreatedAt(v string) *DeleteContextDatabaseMemberResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBody) SetKeys(v []*DeleteContextDatabaseMemberResponseBodyKeys) *DeleteContextDatabaseMemberResponseBody {
	s.Keys = v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBody) SetMemberId(v string) *DeleteContextDatabaseMemberResponseBody {
	s.MemberId = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBody) SetMemberName(v string) *DeleteContextDatabaseMemberResponseBody {
	s.MemberName = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBody) SetRequestId(v string) *DeleteContextDatabaseMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBody) SetRole(v string) *DeleteContextDatabaseMemberResponseBody {
	s.Role = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBody) SetStatus(v string) *DeleteContextDatabaseMemberResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBody) Validate() error {
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

type DeleteContextDatabaseMemberResponseBodyKeys struct {
	// example:
	//
	// (null)
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// (null)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// (null)
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// example:
	//
	// (null)
	KeyDisplaySuffix *string `json:"KeyDisplaySuffix,omitempty" xml:"KeyDisplaySuffix,omitempty"`
	// example:
	//
	// (null)
	KeyId *int64 `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// example:
	//
	// (null)
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	// example:
	//
	// (null)
	LastUsedAt *string `json:"LastUsedAt,omitempty" xml:"LastUsedAt,omitempty"`
	// example:
	//
	// (null)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// (null)
	RevokedAt *string `json:"RevokedAt,omitempty" xml:"RevokedAt,omitempty"`
	// example:
	//
	// (null)
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteContextDatabaseMemberResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseMemberResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetDescription() *string {
	return s.Description
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetKeyDisplaySuffix() *string {
	return s.KeyDisplaySuffix
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetKeyId() *int64 {
	return s.KeyId
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetLastUsedAt() *string {
	return s.LastUsedAt
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetName() *string {
	return s.Name
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) GetStatus() *string {
	return s.Status
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetCreatedAt(v string) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.CreatedAt = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetDescription(v string) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.Description = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetExpiresAt(v string) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.ExpiresAt = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetKeyDisplaySuffix(v string) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.KeyDisplaySuffix = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetKeyId(v int64) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.KeyId = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetKeyPrefix(v string) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.KeyPrefix = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetLastUsedAt(v string) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.LastUsedAt = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetName(v string) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.Name = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetRevokedAt(v string) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.RevokedAt = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) SetStatus(v string) *DeleteContextDatabaseMemberResponseBodyKeys {
	s.Status = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponseBodyKeys) Validate() error {
	return dara.Validate(s)
}

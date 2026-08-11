// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *UpdateContextDatabaseMemberResponseBody
	GetCreatedAt() *string
	SetKeys(v []*UpdateContextDatabaseMemberResponseBodyKeys) *UpdateContextDatabaseMemberResponseBody
	GetKeys() []*UpdateContextDatabaseMemberResponseBodyKeys
	SetMemberId(v string) *UpdateContextDatabaseMemberResponseBody
	GetMemberId() *string
	SetMemberName(v string) *UpdateContextDatabaseMemberResponseBody
	GetMemberName() *string
	SetRequestId(v string) *UpdateContextDatabaseMemberResponseBody
	GetRequestId() *string
	SetRole(v string) *UpdateContextDatabaseMemberResponseBody
	GetRole() *string
	SetStatus(v string) *UpdateContextDatabaseMemberResponseBody
	GetStatus() *string
}

type UpdateContextDatabaseMemberResponseBody struct {
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string                                        `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Keys      []*UpdateContextDatabaseMemberResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
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
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateContextDatabaseMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseMemberResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseMemberResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateContextDatabaseMemberResponseBody) GetKeys() []*UpdateContextDatabaseMemberResponseBodyKeys {
	return s.Keys
}

func (s *UpdateContextDatabaseMemberResponseBody) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateContextDatabaseMemberResponseBody) GetMemberName() *string {
	return s.MemberName
}

func (s *UpdateContextDatabaseMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextDatabaseMemberResponseBody) GetRole() *string {
	return s.Role
}

func (s *UpdateContextDatabaseMemberResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateContextDatabaseMemberResponseBody) SetCreatedAt(v string) *UpdateContextDatabaseMemberResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBody) SetKeys(v []*UpdateContextDatabaseMemberResponseBodyKeys) *UpdateContextDatabaseMemberResponseBody {
	s.Keys = v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBody) SetMemberId(v string) *UpdateContextDatabaseMemberResponseBody {
	s.MemberId = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBody) SetMemberName(v string) *UpdateContextDatabaseMemberResponseBody {
	s.MemberName = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBody) SetRequestId(v string) *UpdateContextDatabaseMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBody) SetRole(v string) *UpdateContextDatabaseMemberResponseBody {
	s.Role = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBody) SetStatus(v string) *UpdateContextDatabaseMemberResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBody) Validate() error {
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

type UpdateContextDatabaseMemberResponseBodyKeys struct {
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// data pipeline key
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// (null)
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// example:
	//
	// 33631c
	KeyDisplaySuffix *string `json:"KeyDisplaySuffix,omitempty" xml:"KeyDisplaySuffix,omitempty"`
	// example:
	//
	// 1
	KeyId *int64 `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// example:
	//
	// ctxdb-
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	// example:
	//
	// 2026-07-15T08:30:00Z
	LastUsedAt *string `json:"LastUsedAt,omitempty" xml:"LastUsedAt,omitempty"`
	// example:
	//
	// my-key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// (null)
	RevokedAt *string `json:"RevokedAt,omitempty" xml:"RevokedAt,omitempty"`
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateContextDatabaseMemberResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseMemberResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetDescription() *string {
	return s.Description
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetKeyDisplaySuffix() *string {
	return s.KeyDisplaySuffix
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetKeyId() *int64 {
	return s.KeyId
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetLastUsedAt() *string {
	return s.LastUsedAt
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetName() *string {
	return s.Name
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) GetStatus() *string {
	return s.Status
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetCreatedAt(v string) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.CreatedAt = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetDescription(v string) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.Description = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetExpiresAt(v string) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.ExpiresAt = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetKeyDisplaySuffix(v string) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.KeyDisplaySuffix = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetKeyId(v int64) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.KeyId = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetKeyPrefix(v string) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.KeyPrefix = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetLastUsedAt(v string) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.LastUsedAt = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetName(v string) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.Name = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetRevokedAt(v string) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.RevokedAt = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) SetStatus(v string) *UpdateContextDatabaseMemberResponseBodyKeys {
	s.Status = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponseBodyKeys) Validate() error {
	return dara.Validate(s)
}

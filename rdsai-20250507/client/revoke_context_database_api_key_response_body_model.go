// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeContextDatabaseApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetCreatedAt() *string
	SetDescription(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetDescription() *string
	SetExpiresAt(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetExpiresAt() *string
	SetKeyDisplaySuffix(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetKeyDisplaySuffix() *string
	SetKeyId(v int64) *RevokeContextDatabaseApiKeyResponseBody
	GetKeyId() *int64
	SetKeyPrefix(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetKeyPrefix() *string
	SetLastUsedAt(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetLastUsedAt() *string
	SetName(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetName() *string
	SetRequestId(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetRequestId() *string
	SetRevokedAt(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetRevokedAt() *string
	SetStatus(v string) *RevokeContextDatabaseApiKeyResponseBody
	GetStatus() *string
}

type RevokeContextDatabaseApiKeyResponseBody struct {
	// The time when the API key was created.
	//
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The description of the API key.
	//
	// example:
	//
	// for nightly cron
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// A reserved field. This field is currently empty.
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
	// 1024
	KeyId *int64 `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The prefix of the API key.
	//
	// example:
	//
	// ctxdb-
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	// The time when the API key was last used.
	//
	// example:
	//
	// 2026-06-01T08:30:12Z
	LastUsedAt *string `json:"LastUsedAt,omitempty" xml:"LastUsedAt,omitempty"`
	// The name of the API key.
	//
	// example:
	//
	// my-key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the API key was revoked.
	//
	// example:
	//
	// 2026-08-07T10:15:30Z
	RevokedAt *string `json:"RevokedAt,omitempty" xml:"RevokedAt,omitempty"`
	// The status of the API key. After revocation, the value is revoked.
	//
	// example:
	//
	// revoked
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s RevokeContextDatabaseApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeContextDatabaseApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetKeyDisplaySuffix() *string {
	return s.KeyDisplaySuffix
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetKeyId() *int64 {
	return s.KeyId
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetLastUsedAt() *string {
	return s.LastUsedAt
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetName() *string {
	return s.Name
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *RevokeContextDatabaseApiKeyResponseBody) GetStatus() *string {
	return s.Status
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetCreatedAt(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetDescription(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.Description = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetExpiresAt(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetKeyDisplaySuffix(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.KeyDisplaySuffix = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetKeyId(v int64) *RevokeContextDatabaseApiKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetKeyPrefix(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.KeyPrefix = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetLastUsedAt(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.LastUsedAt = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetName(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.Name = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetRequestId(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetRevokedAt(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.RevokedAt = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) SetStatus(v string) *RevokeContextDatabaseApiKeyResponseBody {
	s.Status = &v
	return s
}

func (s *RevokeContextDatabaseApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetCreatedAt() *string
	SetDescription(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetDescription() *string
	SetExpiresAt(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetExpiresAt() *string
	SetKeyDisplaySuffix(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetKeyDisplaySuffix() *string
	SetKeyId(v int64) *UpdateContextDatabaseApiKeyResponseBody
	GetKeyId() *int64
	SetKeyPrefix(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetKeyPrefix() *string
	SetLastUsedAt(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetLastUsedAt() *string
	SetName(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetRequestId() *string
	SetRevokedAt(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetRevokedAt() *string
	SetStatus(v string) *UpdateContextDatabaseApiKeyResponseBody
	GetStatus() *string
}

type UpdateContextDatabaseApiKeyResponseBody struct {
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// for nightly cron
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
	// 1024
	KeyId *int64 `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// example:
	//
	// ctxdb-
	KeyPrefix *string `json:"KeyPrefix,omitempty" xml:"KeyPrefix,omitempty"`
	// example:
	//
	// 2026-06-01T08:30:12Z
	LastUsedAt *string `json:"LastUsedAt,omitempty" xml:"LastUsedAt,omitempty"`
	// example:
	//
	// my-key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// (null)
	RevokedAt *string `json:"RevokedAt,omitempty" xml:"RevokedAt,omitempty"`
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateContextDatabaseApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetKeyDisplaySuffix() *string {
	return s.KeyDisplaySuffix
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetKeyId() *int64 {
	return s.KeyId
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetLastUsedAt() *string {
	return s.LastUsedAt
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *UpdateContextDatabaseApiKeyResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetCreatedAt(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetDescription(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetExpiresAt(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.ExpiresAt = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetKeyDisplaySuffix(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.KeyDisplaySuffix = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetKeyId(v int64) *UpdateContextDatabaseApiKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetKeyPrefix(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.KeyPrefix = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetLastUsedAt(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.LastUsedAt = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetName(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetRequestId(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetRevokedAt(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.RevokedAt = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) SetStatus(v string) *UpdateContextDatabaseApiKeyResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateContextDatabaseApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

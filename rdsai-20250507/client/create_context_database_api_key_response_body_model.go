// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateContextDatabaseApiKeyResponseBody
	GetApiKey() *string
	SetKey(v *CreateContextDatabaseApiKeyResponseBodyKey) *CreateContextDatabaseApiKeyResponseBody
	GetKey() *CreateContextDatabaseApiKeyResponseBodyKey
	SetRequestId(v string) *CreateContextDatabaseApiKeyResponseBody
	GetRequestId() *string
}

type CreateContextDatabaseApiKeyResponseBody struct {
	// example:
	//
	// ctxdb-*****
	ApiKey *string                                     `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	Key    *CreateContextDatabaseApiKeyResponseBodyKey `json:"Key,omitempty" xml:"Key,omitempty" type:"Struct"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContextDatabaseApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseApiKeyResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateContextDatabaseApiKeyResponseBody) GetKey() *CreateContextDatabaseApiKeyResponseBodyKey {
	return s.Key
}

func (s *CreateContextDatabaseApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContextDatabaseApiKeyResponseBody) SetApiKey(v string) *CreateContextDatabaseApiKeyResponseBody {
	s.ApiKey = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBody) SetKey(v *CreateContextDatabaseApiKeyResponseBodyKey) *CreateContextDatabaseApiKeyResponseBody {
	s.Key = v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBody) SetRequestId(v string) *CreateContextDatabaseApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBody) Validate() error {
	if s.Key != nil {
		if err := s.Key.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateContextDatabaseApiKeyResponseBodyKey struct {
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// 111
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
	// (null)
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

func (s CreateContextDatabaseApiKeyResponseBodyKey) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseApiKeyResponseBodyKey) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetDescription() *string {
	return s.Description
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetKeyDisplaySuffix() *string {
	return s.KeyDisplaySuffix
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetKeyId() *int64 {
	return s.KeyId
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetLastUsedAt() *string {
	return s.LastUsedAt
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetName() *string {
	return s.Name
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) GetStatus() *string {
	return s.Status
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetCreatedAt(v string) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.CreatedAt = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetDescription(v string) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.Description = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetExpiresAt(v string) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.ExpiresAt = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetKeyDisplaySuffix(v string) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.KeyDisplaySuffix = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetKeyId(v int64) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.KeyId = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetKeyPrefix(v string) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.KeyPrefix = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetLastUsedAt(v string) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.LastUsedAt = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetName(v string) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.Name = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetRevokedAt(v string) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.RevokedAt = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) SetStatus(v string) *CreateContextDatabaseApiKeyResponseBodyKey {
	s.Status = &v
	return s
}

func (s *CreateContextDatabaseApiKeyResponseBodyKey) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextDatabaseApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v []*ListContextDatabaseApiKeysResponseBodyKeys) *ListContextDatabaseApiKeysResponseBody
	GetKeys() []*ListContextDatabaseApiKeysResponseBodyKeys
	SetMaxResults(v int32) *ListContextDatabaseApiKeysResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListContextDatabaseApiKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListContextDatabaseApiKeysResponseBody
	GetRequestId() *string
}

type ListContextDatabaseApiKeysResponseBody struct {
	Keys []*ListContextDatabaseApiKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// example:
	//
	// (null)
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// (null)
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListContextDatabaseApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseApiKeysResponseBody) GetKeys() []*ListContextDatabaseApiKeysResponseBodyKeys {
	return s.Keys
}

func (s *ListContextDatabaseApiKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextDatabaseApiKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextDatabaseApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContextDatabaseApiKeysResponseBody) SetKeys(v []*ListContextDatabaseApiKeysResponseBodyKeys) *ListContextDatabaseApiKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBody) SetMaxResults(v int32) *ListContextDatabaseApiKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBody) SetNextToken(v string) *ListContextDatabaseApiKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBody) SetRequestId(v string) *ListContextDatabaseApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBody) Validate() error {
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

type ListContextDatabaseApiKeysResponseBodyKeys struct {
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
	// (null)
	RevokedAt *string `json:"RevokedAt,omitempty" xml:"RevokedAt,omitempty"`
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListContextDatabaseApiKeysResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s ListContextDatabaseApiKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetDescription() *string {
	return s.Description
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetKeyDisplaySuffix() *string {
	return s.KeyDisplaySuffix
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetKeyId() *int64 {
	return s.KeyId
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetKeyPrefix() *string {
	return s.KeyPrefix
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetLastUsedAt() *string {
	return s.LastUsedAt
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetName() *string {
	return s.Name
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetRevokedAt() *string {
	return s.RevokedAt
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) GetStatus() *string {
	return s.Status
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetCreatedAt(v string) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.CreatedAt = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetDescription(v string) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.Description = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetExpiresAt(v string) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.ExpiresAt = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetKeyDisplaySuffix(v string) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.KeyDisplaySuffix = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetKeyId(v int64) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.KeyId = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetKeyPrefix(v string) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.KeyPrefix = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetLastUsedAt(v string) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.LastUsedAt = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetName(v string) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.Name = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetRevokedAt(v string) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.RevokedAt = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) SetStatus(v string) *ListContextDatabaseApiKeysResponseBodyKeys {
	s.Status = &v
	return s
}

func (s *ListContextDatabaseApiKeysResponseBodyKeys) Validate() error {
	return dara.Validate(s)
}

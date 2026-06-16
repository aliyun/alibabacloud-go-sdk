// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsumerAPIKey interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *ConsumerAPIKey
	GetActive() *bool
	SetConsumerApiKeyId(v string) *ConsumerAPIKey
	GetConsumerApiKeyId() *string
	SetCreatedAt(v string) *ConsumerAPIKey
	GetCreatedAt() *string
	SetDescription(v string) *ConsumerAPIKey
	GetDescription() *string
	SetLastUpdatedAt(v string) *ConsumerAPIKey
	GetLastUpdatedAt() *string
	SetMaskedKey(v string) *ConsumerAPIKey
	GetMaskedKey() *string
	SetModelConnectionId(v string) *ConsumerAPIKey
	GetModelConnectionId() *string
}

type ConsumerAPIKey struct {
	// Specifies if the key is enabled (true) or disabled (false).
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The unique identifier of the consumer API key.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789abc
	ConsumerApiKeyId *string `json:"consumerApiKeyId,omitempty" xml:"consumerApiKeyId,omitempty"`
	// The creation time of the consumer API key, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// A user-defined description for the consumer API key.
	//
	// example:
	//
	// 用于生产环境的API密钥
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The last update time of the consumer API key, in ISO 8601 format.
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// The masked API key, showing only the first and last few characters.
	//
	// example:
	//
	// sk-****1234
	MaskedKey *string `json:"maskedKey,omitempty" xml:"maskedKey,omitempty"`
	// The identifier of the associated model connection.
	//
	// example:
	//
	// mc-1234567890abcdef
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
}

func (s ConsumerAPIKey) String() string {
	return dara.Prettify(s)
}

func (s ConsumerAPIKey) GoString() string {
	return s.String()
}

func (s *ConsumerAPIKey) GetActive() *bool {
	return s.Active
}

func (s *ConsumerAPIKey) GetConsumerApiKeyId() *string {
	return s.ConsumerApiKeyId
}

func (s *ConsumerAPIKey) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ConsumerAPIKey) GetDescription() *string {
	return s.Description
}

func (s *ConsumerAPIKey) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *ConsumerAPIKey) GetMaskedKey() *string {
	return s.MaskedKey
}

func (s *ConsumerAPIKey) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *ConsumerAPIKey) SetActive(v bool) *ConsumerAPIKey {
	s.Active = &v
	return s
}

func (s *ConsumerAPIKey) SetConsumerApiKeyId(v string) *ConsumerAPIKey {
	s.ConsumerApiKeyId = &v
	return s
}

func (s *ConsumerAPIKey) SetCreatedAt(v string) *ConsumerAPIKey {
	s.CreatedAt = &v
	return s
}

func (s *ConsumerAPIKey) SetDescription(v string) *ConsumerAPIKey {
	s.Description = &v
	return s
}

func (s *ConsumerAPIKey) SetLastUpdatedAt(v string) *ConsumerAPIKey {
	s.LastUpdatedAt = &v
	return s
}

func (s *ConsumerAPIKey) SetMaskedKey(v string) *ConsumerAPIKey {
	s.MaskedKey = &v
	return s
}

func (s *ConsumerAPIKey) SetModelConnectionId(v string) *ConsumerAPIKey {
	s.ModelConnectionId = &v
	return s
}

func (s *ConsumerAPIKey) Validate() error {
	return dara.Validate(s)
}

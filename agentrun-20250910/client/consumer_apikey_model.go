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
	// 密钥是否启用，true表示启用，false表示禁用
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// 消费者API密钥的唯一标识符
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789abc
	ConsumerApiKeyId *string `json:"consumerApiKeyId,omitempty" xml:"consumerApiKeyId,omitempty"`
	// 消费者API密钥的创建时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 消费者API密钥的描述信息
	//
	// example:
	//
	// 用于生产环境的API密钥
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 消费者API密钥最后一次更新的时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// API密钥的掩码展示形式，仅显示前后几位字符
	//
	// example:
	//
	// sk-****1234
	MaskedKey *string `json:"maskedKey,omitempty" xml:"maskedKey,omitempty"`
	// 关联的模型连接标识符
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

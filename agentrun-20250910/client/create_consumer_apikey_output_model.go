// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAPIKeyOutput interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *CreateConsumerAPIKeyOutput
	GetActive() *bool
	SetApiKey(v string) *CreateConsumerAPIKeyOutput
	GetApiKey() *string
	SetConsumerApiKeyId(v string) *CreateConsumerAPIKeyOutput
	GetConsumerApiKeyId() *string
	SetCreatedAt(v string) *CreateConsumerAPIKeyOutput
	GetCreatedAt() *string
	SetDescription(v string) *CreateConsumerAPIKeyOutput
	GetDescription() *string
	SetLastUpdatedAt(v string) *CreateConsumerAPIKeyOutput
	GetLastUpdatedAt() *string
	SetMaskedKey(v string) *CreateConsumerAPIKeyOutput
	GetMaskedKey() *string
	SetModelConnectionId(v string) *CreateConsumerAPIKeyOutput
	GetModelConnectionId() *string
}

type CreateConsumerAPIKeyOutput struct {
	// 密钥是否启用
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// 完整的API密钥明文，仅在创建时返回一次，请妥善保存
	//
	// example:
	//
	// sk-xxxxxxxxxxxxxxxxxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// 消费者API密钥的唯一标识符
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789abc
	ConsumerApiKeyId *string `json:"consumerApiKeyId,omitempty" xml:"consumerApiKeyId,omitempty"`
	// 创建时间，采用ISO 8601格式
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
	// 更新时间，采用ISO 8601格式
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	// API密钥的掩码展示形式
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

func (s CreateConsumerAPIKeyOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAPIKeyOutput) GoString() string {
	return s.String()
}

func (s *CreateConsumerAPIKeyOutput) GetActive() *bool {
	return s.Active
}

func (s *CreateConsumerAPIKeyOutput) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateConsumerAPIKeyOutput) GetConsumerApiKeyId() *string {
	return s.ConsumerApiKeyId
}

func (s *CreateConsumerAPIKeyOutput) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateConsumerAPIKeyOutput) GetDescription() *string {
	return s.Description
}

func (s *CreateConsumerAPIKeyOutput) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *CreateConsumerAPIKeyOutput) GetMaskedKey() *string {
	return s.MaskedKey
}

func (s *CreateConsumerAPIKeyOutput) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *CreateConsumerAPIKeyOutput) SetActive(v bool) *CreateConsumerAPIKeyOutput {
	s.Active = &v
	return s
}

func (s *CreateConsumerAPIKeyOutput) SetApiKey(v string) *CreateConsumerAPIKeyOutput {
	s.ApiKey = &v
	return s
}

func (s *CreateConsumerAPIKeyOutput) SetConsumerApiKeyId(v string) *CreateConsumerAPIKeyOutput {
	s.ConsumerApiKeyId = &v
	return s
}

func (s *CreateConsumerAPIKeyOutput) SetCreatedAt(v string) *CreateConsumerAPIKeyOutput {
	s.CreatedAt = &v
	return s
}

func (s *CreateConsumerAPIKeyOutput) SetDescription(v string) *CreateConsumerAPIKeyOutput {
	s.Description = &v
	return s
}

func (s *CreateConsumerAPIKeyOutput) SetLastUpdatedAt(v string) *CreateConsumerAPIKeyOutput {
	s.LastUpdatedAt = &v
	return s
}

func (s *CreateConsumerAPIKeyOutput) SetMaskedKey(v string) *CreateConsumerAPIKeyOutput {
	s.MaskedKey = &v
	return s
}

func (s *CreateConsumerAPIKeyOutput) SetModelConnectionId(v string) *CreateConsumerAPIKeyOutput {
	s.ModelConnectionId = &v
	return s
}

func (s *CreateConsumerAPIKeyOutput) Validate() error {
	return dara.Validate(s)
}

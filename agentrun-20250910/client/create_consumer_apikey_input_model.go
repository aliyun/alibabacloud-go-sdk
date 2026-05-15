// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerAPIKeyInput interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateConsumerAPIKeyInput
	GetApiKey() *string
	SetDescription(v string) *CreateConsumerAPIKeyInput
	GetDescription() *string
	SetModelConnectionId(v string) *CreateConsumerAPIKeyInput
	GetModelConnectionId() *string
}

type CreateConsumerAPIKeyInput struct {
	// 用户自定义的完整API密钥；为空时由服务自动生成
	//
	// example:
	//
	// sk-xxxxxxxxxxxxxxxxxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// 消费者API密钥的描述信息，用于说明该密钥的用途
	//
	// example:
	//
	// 用于生产环境的API密钥
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 关联的模型连接标识符
	//
	// This parameter is required.
	//
	// example:
	//
	// mc-1234567890abcdef
	ModelConnectionId *string `json:"modelConnectionId,omitempty" xml:"modelConnectionId,omitempty"`
}

func (s CreateConsumerAPIKeyInput) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerAPIKeyInput) GoString() string {
	return s.String()
}

func (s *CreateConsumerAPIKeyInput) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateConsumerAPIKeyInput) GetDescription() *string {
	return s.Description
}

func (s *CreateConsumerAPIKeyInput) GetModelConnectionId() *string {
	return s.ModelConnectionId
}

func (s *CreateConsumerAPIKeyInput) SetApiKey(v string) *CreateConsumerAPIKeyInput {
	s.ApiKey = &v
	return s
}

func (s *CreateConsumerAPIKeyInput) SetDescription(v string) *CreateConsumerAPIKeyInput {
	s.Description = &v
	return s
}

func (s *CreateConsumerAPIKeyInput) SetModelConnectionId(v string) *CreateConsumerAPIKeyInput {
	s.ModelConnectionId = &v
	return s
}

func (s *CreateConsumerAPIKeyInput) Validate() error {
	return dara.Validate(s)
}

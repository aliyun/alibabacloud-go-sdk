// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelConnectionConsumerAPIKey interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v string) *ModelConnectionConsumerAPIKey
	GetApiKeyId() *string
	SetValue(v string) *ModelConnectionConsumerAPIKey
	GetValue() *string
}

type ModelConnectionConsumerAPIKey struct {
	// 消费者API密钥记录的唯一标识
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789abc
	ApiKeyId *string `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ModelConnectionConsumerAPIKey) String() string {
	return dara.Prettify(s)
}

func (s ModelConnectionConsumerAPIKey) GoString() string {
	return s.String()
}

func (s *ModelConnectionConsumerAPIKey) GetApiKeyId() *string {
	return s.ApiKeyId
}

func (s *ModelConnectionConsumerAPIKey) GetValue() *string {
	return s.Value
}

func (s *ModelConnectionConsumerAPIKey) SetApiKeyId(v string) *ModelConnectionConsumerAPIKey {
	s.ApiKeyId = &v
	return s
}

func (s *ModelConnectionConsumerAPIKey) SetValue(v string) *ModelConnectionConsumerAPIKey {
	s.Value = &v
	return s
}

func (s *ModelConnectionConsumerAPIKey) Validate() error {
	return dara.Validate(s)
}

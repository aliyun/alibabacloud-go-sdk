// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetSecretValueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecretsListShrink(v string) *BatchGetSecretValueShrinkRequest
	GetSecretsListShrink() *string
}

type BatchGetSecretValueShrinkRequest struct {
	// The list of secret information. You can query up to 20 different secrets at a time.
	SecretsListShrink *string `json:"SecretsList,omitempty" xml:"SecretsList,omitempty"`
}

func (s BatchGetSecretValueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetSecretValueShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchGetSecretValueShrinkRequest) GetSecretsListShrink() *string {
	return s.SecretsListShrink
}

func (s *BatchGetSecretValueShrinkRequest) SetSecretsListShrink(v string) *BatchGetSecretValueShrinkRequest {
	s.SecretsListShrink = &v
	return s
}

func (s *BatchGetSecretValueShrinkRequest) Validate() error {
	return dara.Validate(s)
}

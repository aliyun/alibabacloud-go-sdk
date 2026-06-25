// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *DeleteSecretRequest
	GetNamespaceId() *string
	SetSecretId(v int64) *DeleteSecretRequest
	GetSecretId() *int64
}

type DeleteSecretRequest struct {
	// The ID of the namespace. If the secret is in the default namespace, you need to specify only the region ID, for example, `cn-beijing`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The ID of the secret to delete. You can obtain the ID by calling the [ListSecrets](https://help.aliyun.com/document_detail/466929.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s DeleteSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DeleteSecretRequest) GetSecretId() *int64 {
	return s.SecretId
}

func (s *DeleteSecretRequest) SetNamespaceId(v string) *DeleteSecretRequest {
	s.NamespaceId = &v
	return s
}

func (s *DeleteSecretRequest) SetSecretId(v int64) *DeleteSecretRequest {
	s.SecretId = &v
	return s
}

func (s *DeleteSecretRequest) Validate() error {
	return dara.Validate(s)
}

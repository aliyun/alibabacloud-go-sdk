// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *UpdateSecretShrinkRequest
	GetNamespaceId() *string
	SetSecretDataShrink(v string) *UpdateSecretShrinkRequest
	GetSecretDataShrink() *string
	SetSecretId(v int64) *UpdateSecretShrinkRequest
	GetSecretId() *int64
}

type UpdateSecretShrinkRequest struct {
	// The ID of the namespace where the Secret resides. If the namespace is the default namespace, you need to only enter the region ID, such as `cn-beijing`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The Secret data.
	//
	// This parameter is required.
	SecretDataShrink *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s UpdateSecretShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretShrinkRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateSecretShrinkRequest) GetSecretDataShrink() *string {
	return s.SecretDataShrink
}

func (s *UpdateSecretShrinkRequest) GetSecretId() *int64 {
	return s.SecretId
}

func (s *UpdateSecretShrinkRequest) SetNamespaceId(v string) *UpdateSecretShrinkRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateSecretShrinkRequest) SetSecretDataShrink(v string) *UpdateSecretShrinkRequest {
	s.SecretDataShrink = &v
	return s
}

func (s *UpdateSecretShrinkRequest) SetSecretId(v int64) *UpdateSecretShrinkRequest {
	s.SecretId = &v
	return s
}

func (s *UpdateSecretShrinkRequest) Validate() error {
	return dara.Validate(s)
}

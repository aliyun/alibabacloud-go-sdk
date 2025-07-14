// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *CreateSecretShrinkRequest
	GetNamespaceId() *string
	SetSecretDataShrink(v string) *CreateSecretShrinkRequest
	GetSecretDataShrink() *string
	SetSecretName(v string) *CreateSecretShrinkRequest
	GetSecretName() *string
	SetSecretType(v string) *CreateSecretShrinkRequest
	GetSecretType() *string
}

type CreateSecretShrinkRequest struct {
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
	// registry-auth-acree
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kubernetes.io/dockerconfigjson
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
}

func (s CreateSecretShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretShrinkRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateSecretShrinkRequest) GetSecretDataShrink() *string {
	return s.SecretDataShrink
}

func (s *CreateSecretShrinkRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *CreateSecretShrinkRequest) GetSecretType() *string {
	return s.SecretType
}

func (s *CreateSecretShrinkRequest) SetNamespaceId(v string) *CreateSecretShrinkRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetSecretDataShrink(v string) *CreateSecretShrinkRequest {
	s.SecretDataShrink = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetSecretName(v string) *CreateSecretShrinkRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretShrinkRequest) SetSecretType(v string) *CreateSecretShrinkRequest {
	s.SecretType = &v
	return s
}

func (s *CreateSecretShrinkRequest) Validate() error {
	return dara.Validate(s)
}

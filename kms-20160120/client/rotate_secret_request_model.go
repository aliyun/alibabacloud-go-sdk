// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRotateSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecretName(v string) *RotateSecretRequest
	GetSecretName() *string
	SetVersionId(v string) *RotateSecretRequest
	GetVersionId() *string
}

type RotateSecretRequest struct {
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// RdsSecret/Mysql5.4/MyCred
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret after the secret is rotated.
	//
	// >  The version number is used to ensure the idempotence of the request. Secrets Manager uses this version number to prevent your application from creating the same version of the secret when the application retries a request. If a version number already exists, Secrets Manager ignores the request for rotation and returns a success message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 000000123
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s RotateSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s RotateSecretRequest) GoString() string {
	return s.String()
}

func (s *RotateSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *RotateSecretRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *RotateSecretRequest) SetSecretName(v string) *RotateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *RotateSecretRequest) SetVersionId(v string) *RotateSecretRequest {
	s.VersionId = &v
	return s
}

func (s *RotateSecretRequest) Validate() error {
	return dara.Validate(s)
}

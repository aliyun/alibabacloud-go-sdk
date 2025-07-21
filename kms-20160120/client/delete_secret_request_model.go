// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceDeleteWithoutRecovery(v string) *DeleteSecretRequest
	GetForceDeleteWithoutRecovery() *string
	SetRecoveryWindowInDays(v string) *DeleteSecretRequest
	GetRecoveryWindowInDays() *string
	SetSecretName(v string) *DeleteSecretRequest
	GetSecretName() *string
}

type DeleteSecretRequest struct {
	// Specifies whether to forcibly delete the secret. If this parameter is set to true, the secret cannot be recovered.
	//
	// Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default value)
	//
	// example:
	//
	// false
	ForceDeleteWithoutRecovery *string `json:"ForceDeleteWithoutRecovery,omitempty" xml:"ForceDeleteWithoutRecovery,omitempty"`
	// Specifies the recovery period of the secret if you do not forcibly delete it. Default value: 30. Unit: Days.
	//
	// example:
	//
	// 10
	RecoveryWindowInDays *string `json:"RecoveryWindowInDays,omitempty" xml:"RecoveryWindowInDays,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DeleteSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretRequest) GetForceDeleteWithoutRecovery() *string {
	return s.ForceDeleteWithoutRecovery
}

func (s *DeleteSecretRequest) GetRecoveryWindowInDays() *string {
	return s.RecoveryWindowInDays
}

func (s *DeleteSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *DeleteSecretRequest) SetForceDeleteWithoutRecovery(v string) *DeleteSecretRequest {
	s.ForceDeleteWithoutRecovery = &v
	return s
}

func (s *DeleteSecretRequest) SetRecoveryWindowInDays(v string) *DeleteSecretRequest {
	s.RecoveryWindowInDays = &v
	return s
}

func (s *DeleteSecretRequest) SetSecretName(v string) *DeleteSecretRequest {
	s.SecretName = &v
	return s
}

func (s *DeleteSecretRequest) Validate() error {
	return dara.Validate(s)
}

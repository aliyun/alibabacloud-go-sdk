// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecretName(v string) *RestoreSecretRequest
	GetSecretName() *string
}

type RestoreSecretRequest struct {
	// The name of the secret you want to restore.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s RestoreSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreSecretRequest) GoString() string {
	return s.String()
}

func (s *RestoreSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *RestoreSecretRequest) SetSecretName(v string) *RestoreSecretRequest {
	s.SecretName = &v
	return s
}

func (s *RestoreSecretRequest) Validate() error {
	return dara.Validate(s)
}

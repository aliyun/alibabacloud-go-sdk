// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecretData(v string) *UpdateSecretRequest
	GetSecretData() *string
}

type UpdateSecretRequest struct {
	SecretData *string `json:"secretData,omitempty" xml:"secretData,omitempty"`
}

func (s UpdateSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretRequest) GetSecretData() *string {
	return s.SecretData
}

func (s *UpdateSecretRequest) SetSecretData(v string) *UpdateSecretRequest {
	s.SecretData = &v
	return s
}

func (s *UpdateSecretRequest) Validate() error {
	return dara.Validate(s)
}

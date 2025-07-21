// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretRotationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSecretRotationPolicyResponseBody
	GetRequestId() *string
	SetSecretName(v string) *UpdateSecretRotationPolicyResponseBody
	GetSecretName() *string
}

type UpdateSecretRotationPolicyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 2c124f6f-4210-499f-b88a-69f54004d2d8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// RdsSecret/Mysql5.4/MyCred
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretRotationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretRotationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretRotationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecretRotationPolicyResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *UpdateSecretRotationPolicyResponseBody) SetRequestId(v string) *UpdateSecretRotationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretRotationPolicyResponseBody) SetSecretName(v string) *UpdateSecretRotationPolicyResponseBody {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretRotationPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

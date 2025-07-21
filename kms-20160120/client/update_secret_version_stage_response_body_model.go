// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretVersionStageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSecretVersionStageResponseBody
	GetRequestId() *string
	SetSecretName(v string) *UpdateSecretVersionStageResponseBody
	GetSecretName() *string
}

type UpdateSecretVersionStageResponseBody struct {
	// The name of the secret.
	//
	// example:
	//
	// 8cad259f-4d77-40ec-bbd7-b9c47a423bb9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version to which you want to apply the specified stage label.
	//
	// > 	- You must specify at least one of the RemoveFromVersion and MoveToVersion parameters.
	//
	// > 	- If the VersionStage parameter is set to ACSCurrent or ACSPrevious, this parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UpdateSecretVersionStageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretVersionStageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecretVersionStageResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *UpdateSecretVersionStageResponseBody) SetRequestId(v string) *UpdateSecretVersionStageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretVersionStageResponseBody) SetSecretName(v string) *UpdateSecretVersionStageResponseBody {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretVersionStageResponseBody) Validate() error {
	return dara.Validate(s)
}

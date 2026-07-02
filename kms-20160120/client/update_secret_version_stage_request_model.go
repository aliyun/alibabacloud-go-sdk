// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretVersionStageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMoveToVersion(v string) *UpdateSecretVersionStageRequest
	GetMoveToVersion() *string
	SetRemoveFromVersion(v string) *UpdateSecretVersionStageRequest
	GetRemoveFromVersion() *string
	SetSecretName(v string) *UpdateSecretVersionStageRequest
	GetSecretName() *string
	SetVersionStage(v string) *UpdateSecretVersionStageRequest
	GetVersionStage() *string
}

type UpdateSecretVersionStageRequest struct {
	// The version number of the secret. This parameter specifies that the version stage set by VersionStage is attached to this version.
	//
	// > - You must specify at least one of RemoveFromVersion and MoveToVersion.
	//
	// >
	//
	// > - If you set VersionStage to ACSCurrent or ACSPrevious, you must specify this parameter.
	//
	// example:
	//
	// 002
	MoveToVersion *string `json:"MoveToVersion,omitempty" xml:"MoveToVersion,omitempty"`
	// The version number of the secret. This parameter specifies that the version stage set by VersionStage is removed from this version.
	//
	// > You must specify at least one of RemoveFromVersion and MoveToVersion.
	//
	// example:
	//
	// 001
	RemoveFromVersion *string `json:"RemoveFromVersion,omitempty" xml:"RemoveFromVersion,omitempty"`
	// The name or Alibaba Cloud Resource Name (ARN) of the secret.
	//
	// > To access a secret in a different Alibaba Cloud account, you must specify the ARN of the secret. The ARN is in the format of `acs:kms:${region}:${account}:secret/${secret-name}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version stage of the secret.
	//
	// **Scenario 1: Add a version stage to a specified secret version.**
	//
	// Specify this parameter and MoveToVersion. Do not specify RemoveFromVersion. This parameter can be set to ACSCurrent, ACSPrevious, or a custom stage.
	//
	// **Scenario 2: Remove a version stage from a specified secret version.**
	//
	// Specify this parameter and RemoveFromVersion. Do not specify MoveToVersion. This parameter must be set to a custom stage.
	//
	// > ACSCurrent and ACSPrevious are system-reserved stages. You cannot directly remove them. You can only remove them from one secret version and attach them to another.
	//
	// **Scenario 3: Remove a version stage from a specified secret version and attach it to another secret version.**
	//
	// Specify this parameter, MoveToVersion, and RemoveFromVersion. This parameter can be set to ACSCurrent, ACSPrevious, or a custom stage.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACSCurrent
	VersionStage *string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty"`
}

func (s UpdateSecretVersionStageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretVersionStageRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageRequest) GetMoveToVersion() *string {
	return s.MoveToVersion
}

func (s *UpdateSecretVersionStageRequest) GetRemoveFromVersion() *string {
	return s.RemoveFromVersion
}

func (s *UpdateSecretVersionStageRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *UpdateSecretVersionStageRequest) GetVersionStage() *string {
	return s.VersionStage
}

func (s *UpdateSecretVersionStageRequest) SetMoveToVersion(v string) *UpdateSecretVersionStageRequest {
	s.MoveToVersion = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetRemoveFromVersion(v string) *UpdateSecretVersionStageRequest {
	s.RemoveFromVersion = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetSecretName(v string) *UpdateSecretVersionStageRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetVersionStage(v string) *UpdateSecretVersionStageRequest {
	s.VersionStage = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) Validate() error {
	return dara.Validate(s)
}

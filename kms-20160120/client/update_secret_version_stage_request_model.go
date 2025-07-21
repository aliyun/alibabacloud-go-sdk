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
	// The version from which you want to remove the specified stage label.
	//
	// >  You must specify at least one of the RemoveFromVersion and MoveToVersion parameters.
	//
	// example:
	//
	// 002
	MoveToVersion *string `json:"MoveToVersion,omitempty" xml:"MoveToVersion,omitempty"`
	// The specified stage label. Valid values:
	//
	// 	- ACSCurrent
	//
	// 	- ACSPrevious
	//
	// 	- Custom stage label
	//
	// example:
	//
	// 001
	RemoveFromVersion *string `json:"RemoveFromVersion,omitempty" xml:"RemoveFromVersion,omitempty"`
	// The operation that you want to perform. Set the value to **UpdateSecretVersionStage**.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The name of the secret.
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

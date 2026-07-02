// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutSecretValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSecretData(v string) *PutSecretValueRequest
	GetSecretData() *string
	SetSecretDataType(v string) *PutSecretValueRequest
	GetSecretDataType() *string
	SetSecretName(v string) *PutSecretValueRequest
	GetSecretName() *string
	SetVersionId(v string) *PutSecretValueRequest
	GetVersionId() *string
	SetVersionStages(v string) *PutSecretValueRequest
	GetVersionStages() *string
}

type PutSecretValueRequest struct {
	// The secret value. The value is encrypted and stored in the specified new version.
	//
	// This parameter is required.
	//
	// example:
	//
	// importantdata
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// - text (default)
	//
	// - binary
	//
	// example:
	//
	// text
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The name or Alibaba Cloud Resource Name (ARN) of the secret.
	//
	// > When you access a secret in another Alibaba Cloud account, you must specify the ARN of the secret. The ARN of a secret is in the format of `acs:kms:${region}:${account}:secret/${secret-name}`.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret. The value must be unique in the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that are used to mark the new version. If you do not specify this parameter, KMS marks the new version with ACSCurrent.
	//
	// example:
	//
	// ["ACSCurrent","ACSNext"]
	VersionStages *string `json:"VersionStages,omitempty" xml:"VersionStages,omitempty"`
}

func (s PutSecretValueRequest) String() string {
	return dara.Prettify(s)
}

func (s PutSecretValueRequest) GoString() string {
	return s.String()
}

func (s *PutSecretValueRequest) GetSecretData() *string {
	return s.SecretData
}

func (s *PutSecretValueRequest) GetSecretDataType() *string {
	return s.SecretDataType
}

func (s *PutSecretValueRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *PutSecretValueRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PutSecretValueRequest) GetVersionStages() *string {
	return s.VersionStages
}

func (s *PutSecretValueRequest) SetSecretData(v string) *PutSecretValueRequest {
	s.SecretData = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretDataType(v string) *PutSecretValueRequest {
	s.SecretDataType = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretName(v string) *PutSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *PutSecretValueRequest) SetVersionId(v string) *PutSecretValueRequest {
	s.VersionId = &v
	return s
}

func (s *PutSecretValueRequest) SetVersionStages(v string) *PutSecretValueRequest {
	s.VersionStages = &v
	return s
}

func (s *PutSecretValueRequest) Validate() error {
	return dara.Validate(s)
}

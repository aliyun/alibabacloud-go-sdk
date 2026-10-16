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
	// The secret value. The secret value is encrypted and stored in the specified new version.
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
	// > When you access a secret that belongs to another Alibaba Cloud account, you must specify the ARN of the secret. The ARN of a secret is in the `acs:kms:${region}:${account}:secret/${secret-name}` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version ID of the secret. The version ID must be unique within the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// v3
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The version stages that are attached to the secret version when it is stored. If you do not specify this parameter, KMS attaches ACSCurrent to the new version by default.
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

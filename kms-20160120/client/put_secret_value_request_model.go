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
	// The secret value. The value is encrypted and then stored in the new version.
	//
	// This parameter is required.
	//
	// example:
	//
	// importantdata
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The type of the secret value. Valid values:
	//
	// 	- text: This is the default value.
	//
	// 	- binary
	//
	// example:
	//
	// text
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	// The name of the secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The new version of the secret value. Version numbers must be unique in each secret.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00000000000000000000000000000000203
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The stage labels that are used to mark the new version. If you do not specify this parameter, Secrets Manager marks the new version with ACSCurrent.
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

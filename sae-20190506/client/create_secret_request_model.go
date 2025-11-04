// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *CreateSecretRequest
	GetNamespaceId() *string
	SetSecretData(v *CreateSecretRequestSecretData) *CreateSecretRequest
	GetSecretData() *CreateSecretRequestSecretData
	SetSecretName(v string) *CreateSecretRequest
	GetSecretName() *string
	SetSecretType(v string) *CreateSecretRequest
	GetSecretType() *string
}

type CreateSecretRequest struct {
	// The ID of the namespace where the Secret resides. If the namespace is the default namespace, you need to only enter the region ID, such as `cn-beijing`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The Secret data.
	//
	// This parameter is required.
	SecretData *CreateSecretRequestSecretData `json:"SecretData,omitempty" xml:"SecretData,omitempty" type:"Struct"`
	// The Secret name. The name can contain digits, letters, and underscores (_). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry-auth-acree
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The supported Secret type. Valid values:
	//
	// 	- **kubernetes.io/dockerconfigjson**: the Secret for the username and password of the image repository. The Secret is used for authentication when images are pulled during application deployment.
	//
	// Valid values:
	//
	// 	- Opaque
	//
	// 	- kubernetes.io/dockerconfigjson
	//
	// 	- kubernetes.io/tls
	//
	// This parameter is required.
	//
	// example:
	//
	// kubernetes.io/dockerconfigjson
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
}

func (s CreateSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateSecretRequest) GetSecretData() *CreateSecretRequestSecretData {
	return s.SecretData
}

func (s *CreateSecretRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *CreateSecretRequest) GetSecretType() *string {
	return s.SecretType
}

func (s *CreateSecretRequest) SetNamespaceId(v string) *CreateSecretRequest {
	s.NamespaceId = &v
	return s
}

func (s *CreateSecretRequest) SetSecretData(v *CreateSecretRequestSecretData) *CreateSecretRequest {
	s.SecretData = v
	return s
}

func (s *CreateSecretRequest) SetSecretName(v string) *CreateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretRequest) SetSecretType(v string) *CreateSecretRequest {
	s.SecretType = &v
	return s
}

func (s *CreateSecretRequest) Validate() error {
	if s.SecretData != nil {
		if err := s.SecretData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecretRequestSecretData struct {
	// The information about the key-value pairs of the Secret. This parameter is required. The following formats are supported:
	//
	// {"Data":"{"k1":"v1", "k2":"v2"}"}
	//
	// k specifies a key and v specifies a value.
	//
	// This parameter is required.
	//
	// example:
	//
	// {".dockerconfigjson":"eyJhdXRocyI6eyJyZWdpc3RyeS12cGMuY24tYmVpamluZy5hbGl5dW5jcy5jb20iOnsidXNlcm5hbWUiOiJ1c2VybmFtZSIsInBhc3N3b3JkIjoicGFzc3dvcmQiLCJhdXRoIjoiZFhObGNtNWhiV1U2Y0dGemMzZHZjbVE9In0sInJlZ2lzdHJ5LmNuLWJlaWppbmcuYWxpeXVuY3MuY29tIjp7InVzZXJuYW1lIjoidXNlcm5hbWUiLCJwYXNzd29yZCI6InBhc3N3b3JkIiwiYXV0aCI6ImRYTmxjbTVoYldVNmNHRnpjM2R2Y21RPSJ9fX0="}
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
}

func (s CreateSecretRequestSecretData) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretRequestSecretData) GoString() string {
	return s.String()
}

func (s *CreateSecretRequestSecretData) GetSecretData() *string {
	return s.SecretData
}

func (s *CreateSecretRequestSecretData) SetSecretData(v string) *CreateSecretRequestSecretData {
	s.SecretData = &v
	return s
}

func (s *CreateSecretRequestSecretData) Validate() error {
	return dara.Validate(s)
}

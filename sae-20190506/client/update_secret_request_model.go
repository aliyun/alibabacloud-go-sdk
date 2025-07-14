// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *UpdateSecretRequest
	GetNamespaceId() *string
	SetSecretData(v *UpdateSecretRequestSecretData) *UpdateSecretRequest
	GetSecretData() *UpdateSecretRequestSecretData
	SetSecretId(v int64) *UpdateSecretRequest
	GetSecretId() *int64
}

type UpdateSecretRequest struct {
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
	SecretData *UpdateSecretRequestSecretData `json:"SecretData,omitempty" xml:"SecretData,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s UpdateSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateSecretRequest) GetSecretData() *UpdateSecretRequestSecretData {
	return s.SecretData
}

func (s *UpdateSecretRequest) GetSecretId() *int64 {
	return s.SecretId
}

func (s *UpdateSecretRequest) SetNamespaceId(v string) *UpdateSecretRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateSecretRequest) SetSecretData(v *UpdateSecretRequestSecretData) *UpdateSecretRequest {
	s.SecretData = v
	return s
}

func (s *UpdateSecretRequest) SetSecretId(v int64) *UpdateSecretRequest {
	s.SecretId = &v
	return s
}

func (s *UpdateSecretRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateSecretRequestSecretData struct {
	// The information about the key-value pairs of the Secret. This parameter is required. The following formats are supported:
	//
	// {"Data":"{"k1":"v1", "k2":"v2"}"}
	//
	// k specifies a key and v specifies a value. For more information, see [Manage a Kubernetes Secret](https://help.aliyun.com/document_detail/463383.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {".dockerconfigjson":"eyJhdXRocyI6eyJyZWdpc3RyeS12cGMuY24tYmVpamluZy5hbGl5dW5jcy5jb20iOnsidXNlcm5hbWUiOiJ1c2VybmFtZSIsInBhc3N3b3JkIjoicGFzc3dvcmQiLCJhdXRoIjoiZFhObGNtNWhiV1U2Y0dGemMzZHZjbVE9In0sInJlZ2lzdHJ5LmNuLWJlaWppbmcuYWxpeXVuY3MuY29tIjp7InVzZXJuYW1lIjoidXNlcm5hbWUiLCJwYXNzd29yZCI6InBhc3N3b3JkIiwiYXV0aCI6ImRYTmxjbTVoYldVNmNHRnpjM2R2Y21RPSJ9fX0="}
	SecretData *string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
}

func (s UpdateSecretRequestSecretData) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretRequestSecretData) GoString() string {
	return s.String()
}

func (s *UpdateSecretRequestSecretData) GetSecretData() *string {
	return s.SecretData
}

func (s *UpdateSecretRequestSecretData) SetSecretData(v string) *UpdateSecretRequestSecretData {
	s.SecretData = &v
	return s
}

func (s *UpdateSecretRequestSecretData) Validate() error {
	return dara.Validate(s)
}

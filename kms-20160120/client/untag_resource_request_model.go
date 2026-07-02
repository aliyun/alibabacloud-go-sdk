// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUntagResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *UntagResourceRequest
	GetCertificateId() *string
	SetKeyId(v string) *UntagResourceRequest
	GetKeyId() *string
	SetSecretName(v string) *UntagResourceRequest
	GetSecretName() *string
	SetTagKeys(v string) *UntagResourceRequest
	GetTagKeys() *string
}

type UntagResourceRequest struct {
	// The ID of the certificate.
	//
	// > You must specify one and only one of the KeyId, SecretName, and CertificateId parameters.
	//
	// example:
	//
	// 770dbe42-e146-43d1-a55a-1355db86****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The ID of the key. This is the globally unique identifier (GUID) of the master key (CMK).
	//
	// > You must specify one and only one of the KeyId, SecretName, and CertificateId parameters.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddf****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the credential.
	//
	// > You must specify one and only one of the KeyId, SecretName, and CertificateId parameters.
	//
	// example:
	//
	// MyDbC****
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// One or more tag keys. Separate multiple tag keys with commas (,).<br> You need to specify only tag keys, not tag values.<br> The tag key can be 1 to 128 bytes in length.<br><br>
	//
	// This parameter is required.
	//
	// example:
	//
	// ["tagkey1","tagkey2"]
	TagKeys *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourceRequest) GoString() string {
	return s.String()
}

func (s *UntagResourceRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *UntagResourceRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *UntagResourceRequest) GetSecretName() *string {
	return s.SecretName
}

func (s *UntagResourceRequest) GetTagKeys() *string {
	return s.TagKeys
}

func (s *UntagResourceRequest) SetCertificateId(v string) *UntagResourceRequest {
	s.CertificateId = &v
	return s
}

func (s *UntagResourceRequest) SetKeyId(v string) *UntagResourceRequest {
	s.KeyId = &v
	return s
}

func (s *UntagResourceRequest) SetSecretName(v string) *UntagResourceRequest {
	s.SecretName = &v
	return s
}

func (s *UntagResourceRequest) SetTagKeys(v string) *UntagResourceRequest {
	s.TagKeys = &v
	return s
}

func (s *UntagResourceRequest) Validate() error {
	return dara.Validate(s)
}
